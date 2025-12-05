// Copyright 2018 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package cache

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"maps"
	"slices"
	"sync"
	"sync/atomic"
	"time"

	"github.com/envoyproxy/go-control-plane/pkg/cache/internal"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/log"
)

// ResourceSnapshot is an abstract snapshot of a collection of resources that
// can be stored in a SnapshotCache. This enables applications to use the
// SnapshotCache watch machinery with their own resource types. Most
// applications will use Snapshot.
type ResourceSnapshot interface {
	// GetVersion should return the current version of the resource indicated
	// by typeURL. The version string that is returned is opaque and should
	// only be compared for equality.
	GetVersion(typeURL string) string

	// GetResources returns all resources of the type indicted by typeURL.
	// This is called by the cache to build responses.
	GetTypeSnapshot(typeURL string) types.TypeSnapshot
}

// SnapshotCache is a snapshot-based cache that maintains a single versioned
// snapshot of responses per node. SnapshotCache consistently replies with the
// latest snapshot. For the protocol to work correctly in ADS mode, EDS/RDS
// requests are responded only when all resources in the snapshot xDS response
// are named as part of the request. It is expected that the CDS response names
// all EDS clusters, and the LDS response names all RDS routes in a snapshot,
// to ensure that Envoy makes the request for all EDS clusters or RDS routes
// eventually.
//
// SnapshotCache can operate as a REST or regular xDS backend. The snapshot
// can be partial, e.g. only include RDS or EDS resources.
type SnapshotCache interface {
	Cache

	// SetSnapshot sets a response snapshot for a node. For ADS, the snapshots
	// should have distinct versions and be internally consistent (e.g. all
	// referenced resources must be included in the snapshot).
	//
	// This method will cause the server to respond to all open watches, for which
	// the version differs from the snapshot version.
	SetSnapshot(ctx context.Context, node string, snapshot ResourceSnapshot) error

	// GetSnapshots gets the snapshot for a node.
	GetSnapshot(node string) (ResourceSnapshot, error)

	// ClearSnapshot removes all status and snapshot information associated with a node.
	ClearSnapshot(node string)

	// GetStatusInfo retrieves status information for a node ID.
	GetStatusInfo(string) StatusInfo

	// GetStatusKeys retrieves node IDs for all statuses.
	GetStatusKeys() []string
}

type snapshotCache struct {
	// watchCount and deltaWatchCount are atomic counters incremented for each watch respectively. They need to
	// be the first fields in the struct to guarantee 64-bit alignment,
	// which is a requirement for atomic operations on 64-bit operands to work on
	// 32-bit machines.
	watchCount      int64
	deltaWatchCount int64

	log log.Logger

	// ads flag to hold responses until all resources are named
	ads bool

	// snapshots are cached resources indexed by node IDs
	snapshots map[string]ResourceSnapshot

	// status information for all nodes indexed by node IDs
	status map[string]*statusInfo

	// hash is the hashing function for Envoy nodes
	hash NodeHash

	mu sync.RWMutex
}

// missingRequestResource is returned when the request is specifically dropped due to the resources in the request not matching the snapshot content.
// This error is not returned to the user.
// TODO(valerian-roche): remove this check which is very likely no longer needed
type missingRequestResource struct {
	resources []string
}

func (e missingRequestResource) Error() string {
	return fmt.Sprintf("missing resources in request: %v", e.resources)
}

// NewSnapshotCache initializes a simple cache.
//
// ADS flag forces a delay in responding to streaming requests until all
// resources are explicitly named in the request. This avoids the problem of a
// partial request over a single stream for a subset of resources which would
// require generating a fresh version for acknowledgement. ADS flag requires
// snapshot consistency. For non-ADS case (and fetch), multiple partial
// requests are sent across multiple streams and re-using the snapshot version
// is OK.
//
// Logger is optional.
func NewSnapshotCache(ads bool, hash NodeHash, logger log.Logger) SnapshotCache {
	return newSnapshotCache(ads, hash, logger)
}

func newSnapshotCache(ads bool, hash NodeHash, logger log.Logger) *snapshotCache {
	if logger == nil {
		logger = log.NewDefaultLogger()
	}

	cache := &snapshotCache{
		log:       logger,
		ads:       ads,
		snapshots: make(map[string]ResourceSnapshot),
		status:    make(map[string]*statusInfo),
		hash:      hash,
	}

	return cache
}

// NewSnapshotCacheWithHeartbeating initializes a simple cache that sends periodic heartbeat
// responses for resources with a TTL.
//
// ADS flag forces a delay in responding to streaming requests until all
// resources are explicitly named in the request. This avoids the problem of a
// partial request over a single stream for a subset of resources which would
// require generating a fresh version for acknowledgement. ADS flag requires
// snapshot consistency. For non-ADS case (and fetch), multiple partial
// requests are sent across multiple streams and re-using the snapshot version
// is OK.
//
// Logger is optional.
//
// The context provides a way to cancel the heartbeating routine, while the heartbeatInterval
// parameter controls how often heartbeating occurs.
func NewSnapshotCacheWithHeartbeating(ctx context.Context, ads bool, hash NodeHash, logger log.Logger, heartbeatInterval time.Duration) SnapshotCache {
	cache := newSnapshotCache(ads, hash, logger)
	go func() {
		t := time.NewTicker(heartbeatInterval)

		for {
			select {
			case <-t.C:
				cache.mu.Lock()
				for node := range cache.status {
					// TODO(snowp): Omit heartbeats if a real response has been sent recently.
					cache.sendHeartbeats(ctx, node)
				}
				cache.mu.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}()
	return cache
}

func (cache *snapshotCache) sendHeartbeats(ctx context.Context, node string) {
	snapshot, ok := cache.snapshots[node]
	if !ok {
		return
	}

	if info, ok := cache.status[node]; ok {
		info.mu.Lock()
		for id, watch := range info.watches {
			typeSnapshot := snapshot.GetTypeSnapshot(watch.Request.GetTypeUrl())
			resources := typeSnapshot.GetResources()
			if len(resources) == 0 {
				// No resources for this type.
				continue
			}

			resourcesToReturn := map[string]*internal.CachedResource{}
			addResource := func(res *internal.CachedResource) {
				if !res.HasTTL() {
					return
				}
				resourcesToReturn[res.Name] = res
			}

			if watch.subscription.IsWildcard() {
				resourcesToReturn = make(map[string]*internal.CachedResource, len(resources))
				for _, res := range resources {
					// Include wildcard-eligible resources for wildcard subscriptions
					if !res.OnDemandOnly() {
						addResource(res)
					}
				}
			}
			for name := range watch.subscription.SubscribedResources() {
				if res, ok := resources[name]; ok {
					addResource(res)
				}
			}

			if len(resourcesToReturn) == 0 {
				continue
			}

			resp := &RawResponse{
				Request:   watch.Request,
				Version:   typeSnapshot.GetVersion(),
				resources: slices.Collect(maps.Values(resourcesToReturn)),
				// Do not alter it. Those TTLs do not touch what's actually in watches.
				returnedResources: watch.subscription.ReturnedResources(),
				Heartbeat:         true,
				Ctx:               ctx,
			}

			cache.log.Debugf("respond open watch %d %v with heartbeat for version %q", id, watch.Request.GetResourceNames(), resp.Version)
			err := cache.respond(ctx, watch, resp)
			if err != nil {
				cache.log.Errorf("received error when attempting to respond to watches: %v", err)
			} else {
				// The watch must be deleted and we must rely on the client to ack this response to create a new watch.
				delete(info.watches, id)
			}
		}
		info.mu.Unlock()
	}
}

// SetSnapshotCache updates a snapshot for a node.
func (cache *snapshotCache) SetSnapshot(ctx context.Context, node string, snapshot ResourceSnapshot) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.log.Debugf("setting snapshot for node %s", node)
	// update the existing entry
	cache.snapshots[node] = snapshot

	// trigger existing watches for which version changed
	if info, ok := cache.status[node]; ok {
		info.mu.Lock()
		defer info.mu.Unlock()

		// Respond to SOTW watches for the node.
		if err := cache.respondSOTWWatches(ctx, info, snapshot); err != nil {
			return err
		}

		// Respond to delta watches for the node.
		return cache.respondDeltaWatches(ctx, info, snapshot)
	}

	return nil
}

func (cache *snapshotCache) respondSOTWWatches(ctx context.Context, info *statusInfo, snapshot ResourceSnapshot) error {
	// responder callback for SOTW watches
	respond := func(watch ResponseWatch, id int64) error {
		version := snapshot.GetVersion(watch.Request.GetTypeUrl())
		if version == watch.Request.GetVersionInfo() {
			// Snapshot did not change, no reply.
			return nil
		}

		cache.log.Debugf("consider open watch %d %s %v with new version %q", id, watch.Request.GetTypeUrl(), watch.Request.GetResourceNames(), version)
		resp, err := createResponse(snapshot, watch, cache.ads)
		if errors.As(err, &missingRequestResource{}) {
			return nil
		}
		if err != nil {
			return err
		}
		if resp != nil {
			cache.log.Debugf("respond open watch %d %s %v with new version %q", id, watch.Request.GetTypeUrl(), watch.Request.GetResourceNames(), version)
			err := cache.respond(ctx, watch, resp)
			if err != nil {
				return err
			}
			// discard the watch
			delete(info.watches, id)
		}
		// If we did not reply we just keep the watch
		return nil
	}

	// If ADS is enabled we need to order response watches so we guarantee
	// sending them in the correct order. Go's default implementation
	// of maps are randomized order when ranged over.
	if cache.ads {
		info.orderResponseWatches()
		for _, key := range info.orderedWatches {
			err := respond(info.watches[key.ID], key.ID)
			if err != nil {
				return err
			}
		}
	} else {
		for id, watch := range info.watches {
			err := respond(watch, id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (cache *snapshotCache) respondDeltaWatches(ctx context.Context, info *statusInfo, snapshot ResourceSnapshot) error {
	// We only calculate version hashes when using delta. We don't
	// want to do this when using SOTW so we can avoid unnecessary
	// computational cost if not using delta.
	if len(info.deltaWatches) == 0 {
		return nil
	}

	replyWatch := func(watch DeltaResponseWatch, id int64) error {
		resp, err := createDeltaResponse(snapshot, watch, false)
		if err != nil {
			return fmt.Errorf("creating response: %w", err)
		}

		// If we receive a nil response here, that means there has been no state change
		// so we don't want to respond or remove any existing resource watches
		if resp == nil {
			return nil
		}

		err = cache.respondDelta(ctx, watch, resp)
		if err != nil {
			return err
		}
		delete(info.deltaWatches, id)
		return nil
	}

	// If ADS is enabled we need to order response delta watches so we guarantee
	// sending them in the correct order. Go's default implementation
	// of maps are randomized order when ranged over.
	if cache.ads {
		info.orderResponseDeltaWatches()
		for _, key := range info.orderedDeltaWatches {
			watch := info.deltaWatches[key.ID]
			err := replyWatch(watch, key.ID)
			if err != nil {
				return err
			}
		}
	} else {
		for id, watch := range info.deltaWatches {
			err := replyWatch(watch, id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetSnapshot gets the snapshot for a node, and returns an error if not found.
func (cache *snapshotCache) GetSnapshot(node string) (ResourceSnapshot, error) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	snap, ok := cache.snapshots[node]
	if !ok {
		return nil, fmt.Errorf("no snapshot found for node %s", node)
	}
	return snap, nil
}

// ClearSnapshot clears snapshot and info for a node.
func (cache *snapshotCache) ClearSnapshot(node string) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	delete(cache.snapshots, node)
	delete(cache.status, node)
}

// CreateWatch returns a watch for an xDS request.  A nil function may be
// returned if an error occurs.
func (cache *snapshotCache) CreateWatch(request *Request, sub Subscription, value chan Response) (func(), error) {
	nodeID := cache.hash.ID(request.GetNode())

	cache.mu.Lock()
	defer cache.mu.Unlock()

	info, ok := cache.status[nodeID]
	if !ok {
		info = newStatusInfo(request.GetNode())
		cache.status[nodeID] = info
	}

	// update last watch request time
	info.setLastWatchRequestTime(time.Now())

	createWatch := func(watch ResponseWatch) func() {
		watchID := cache.nextWatchID()
		cache.log.Debugf("open watch %d for %s %v from nodeID %q, version %q", watchID, request.GetTypeUrl(), sub.SubscribedResources(), nodeID, request.GetVersionInfo())
		info.mu.Lock()
		info.watches[watchID] = watch
		info.mu.Unlock()
		return cache.cancelWatch(nodeID, watchID)
	}

	watch := ResponseWatch{Request: request, Response: value, subscription: sub, fullStateResponses: ResourceRequiresFullStateInSotw(request.GetTypeUrl())}

	snapshot, exists := cache.snapshots[nodeID]
	if !exists {
		return createWatch(watch), nil
	}

	resp, err := createResponse(snapshot, watch, cache.ads)
	// Specific legacy case. We explicitly drop the request (and therefore do not reply or track the watch) while keeping the stream opened.
	// TODO(valerian-roche): this is likely unneeded now, to be cleaned
	if errors.As(err, &missingRequestResource{}) {
		cache.log.Warnf("ADS mode: not responding to request %s %v: %v", request.GetTypeUrl(), request.GetResourceNames(), err)
		return func() {}, nil
	}
	if err != nil {
		return func() {}, fmt.Errorf("failed to create response: %w", err)
	}
	if resp != nil {
		if err := cache.respond(context.Background(), watch, resp); err != nil {
			cache.log.Errorf("failed to send a response for %s%v to nodeID %q: %s", request.GetTypeUrl(),
				sub.SubscribedResources(), nodeID, err)
			return nil, fmt.Errorf("failed to send the response: %w", err)
		}
		return func() {}, nil
	}

	return createWatch(watch), nil
}

func (cache *snapshotCache) nextWatchID() int64 {
	return atomic.AddInt64(&cache.watchCount, 1)
}

// cancellation function for cleaning stale watches
func (cache *snapshotCache) cancelWatch(nodeID string, watchID int64) func() {
	return func() {
		// uses the cache mutex
		cache.mu.RLock()
		defer cache.mu.RUnlock()
		if info, ok := cache.status[nodeID]; ok {
			info.mu.Lock()
			delete(info.watches, watchID)
			info.mu.Unlock()
		}
	}
}

// difference returns the names present in resources but not in names.
func difference[T any](resources iter.Seq[string], names map[string]T) []string {
	var diff []string
	for resourceName := range resources {
		if _, exists := names[resourceName]; !exists {
			diff = append(diff, resourceName)
		}
	}
	return diff
}

// createResponse evaluates the provided watch against the given snapshot to build the response to return.
// It may return a nil response to indicate the watch is up-to-date for the given snapshot.
// It is currently inefficient as not evaluating known resources intrisic versions, but only the snapshot one.
// Further work may be performed to optimize this.
func createResponse(snapshot ResourceSnapshot, watch ResponseWatch, ads bool) (*RawResponse, error) {
	typeURL := watch.Request.TypeUrl
	// Use version from snapshot and not typeSnapshot to support negative case (i.e. no resources are set).
	version := snapshot.GetVersion(typeURL)
	resources := snapshot.GetTypeSnapshot(typeURL).GetResources()

	// for ADS, the request names must match the snapshot names
	// if they do not, then the watch is never responded, and it is expected that envoy makes another request
	if !watch.subscription.IsWildcard() && ads {
		if missing := difference(maps.Keys(resources), watch.subscription.SubscribedResources()); len(missing) > 0 {
			return nil, missingRequestResource{missing}
		}
	}

	// This implementation can seem more complex than needed, as it does not blindly rely on the request version.
	// This allows for a more generic implemenentation when considering wildcard + subscribed, or partial replies.
	// In other context a lot could be simplified as

	reqVersion := watch.Request.VersionInfo

	subscribedResources := watch.subscription.SubscribedResources()
	knownResources := watch.subscription.ReturnedResources()

	// Only populated if version has not changed.
	var changedResources map[string]struct{} // Use a map to merge wildcard and subscription.
	var deletedResources []string

	if version == reqVersion {
		// Check if a resource was not previously returned (e.g. if the watch is newly wildcard).
		if watch.subscription.IsWildcard() {
			changedResources = make(map[string]struct{}, len(resources))
			for name, res := range resources {
				// Include wildcard-eligible resources for wildcard subscriptions
				if res.OnDemandOnly() {
					continue
				}

				if _, known := knownResources[name]; !known {
					changedResources[name] = struct{}{}
				}
			}
		} else {
			changedResources = make(map[string]struct{}, len(subscribedResources))
		}

		for name := range subscribedResources {
			_, exist := resources[name]
			if !exist {
				continue
			}
			_, known := knownResources[name]
			if known {
				continue
			}
			changedResources[name] = struct{}{}
		}

		if len(changedResources) == 0 && !watch.sendFullStateResponses() {
			// If full state responses are needed we need to trigger if only deletions occurred,
			// otherwise we can just bail out.
			return nil, nil
		}

		for name := range knownResources {
			res, exist := resources[name]
			if !exist {
				deletedResources = append(deletedResources, name)
				continue
			}

			if watch.subscription.IsWildcard() && !res.OnDemandOnly() {
				continue
			}

			if _, explicitlySubscribed := subscribedResources[name]; !explicitlySubscribed {
				deletedResources = append(deletedResources, name)
			}
		}

		if len(changedResources) == 0 && len(deletedResources) == 0 {
			// Nothing's changed
			return nil, nil
		}
	}

	// Now compute the response.
	var resourcesToReturn []*internal.CachedResource
	var returnedResources map[string]string
	if version != reqVersion || watch.sendFullStateResponses() {
		// Return all resources, with no regard to known version.
		if watch.subscription.IsWildcard() {
			resourcesToReturn = make([]*internal.CachedResource, 0, len(resources))
			returnedResources = make(map[string]string, len(resources))
			for name, resource := range resources {
				// Include wildcard-eligible resources for wildcard subscriptions
				if !resource.OnDemandOnly() {
					resourcesToReturn = append(resourcesToReturn, resource)
					returnedResources[name] = version
				}
			}
		} else {
			resourcesToReturn = make([]*internal.CachedResource, 0, len(subscribedResources))
			returnedResources = make(map[string]string, len(subscribedResources))
		}
		for name := range subscribedResources {
			resource, ok := resources[name]
			if !ok {
				continue
			}
			if _, alreadyAdded := returnedResources[name]; !alreadyAdded {
				resourcesToReturn = append(resourcesToReturn, resource)
				returnedResources[name] = version
			}
		}
	} else {
		// Same version and not full state, only return newly subscribed resources.
		returnedResources = maps.Clone(knownResources)
		for name := range changedResources {
			resource, ok := resources[name]
			if !ok {
				// Should never occur.
				continue
			}
			resourcesToReturn = append(resourcesToReturn, resource)
			returnedResources[name] = version
		}
		for _, name := range deletedResources {
			// Cleanup resources no longer subscribed to make sure we resend them if re-subscribed later.
			delete(returnedResources, name)
		}
	}

	return &RawResponse{
		Request:           watch.Request,
		Version:           version,
		resources:         resourcesToReturn,
		returnedResources: returnedResources,
	}, nil
}

func (cache *snapshotCache) respond(ctx context.Context, watch ResponseWatch, response *RawResponse) error {
	request := watch.Request
	cache.log.Debugf("respond %s (requested %v) version %q with version %q and resources %v", request.GetTypeUrl(), request.GetResourceNames(), request.GetVersionInfo(), response.Version, slices.Collect(maps.Keys(response.GetReturnedResources())))
	response.Ctx = ctx

	select {
	case watch.Response <- response:
		return nil
	case <-ctx.Done():
		return context.Canceled
	}
}

// CreateDeltaWatch returns a watch for a delta xDS request which implements the Simple SnapshotCache.
func (cache *snapshotCache) CreateDeltaWatch(request *DeltaRequest, sub Subscription, value chan DeltaResponse) (func(), error) {
	nodeID := cache.hash.ID(request.GetNode())
	t := request.GetTypeUrl()

	cache.mu.Lock()
	defer cache.mu.Unlock()

	info, ok := cache.status[nodeID]
	if !ok {
		info = newStatusInfo(request.GetNode())
		cache.status[nodeID] = info
	}

	// update last watch request time
	info.setLastDeltaWatchRequestTime(time.Now())

	watch := DeltaResponseWatch{Request: request, Response: value, subscription: sub}

	// find the current cache snapshot for the provided node
	snapshot, exists := cache.snapshots[nodeID]
	if exists {
		resp, err := createDeltaResponse(snapshot, watch, sub.IsWildcard() && request.ResponseNonce == "")
		if err != nil {
			return nil, fmt.Errorf("building response: %w", err)
		}

		if resp != nil {
			err := cache.respondDelta(context.Background(), watch, resp)
			if err != nil {
				cache.log.Errorf("failed to respond with delta response: %s", err)
				return func() {}, fmt.Errorf("responding: %w", err)
			}
			return func() {}, nil
		}

		// We did not reply, fallthrough to watch tracking
	}

	// There are two different cases that leads to a delayed watch trigger:
	// - no snapshot exists for the requested nodeID
	// - we attempted to issue a response, but the caller is already up to date
	watchID := cache.nextDeltaWatchID()
	if exists {
		cache.log.Infof("open delta watch ID:%d for %s Resources:%v from nodeID: %q,  version %q", watchID, t, sub.SubscribedResources(), nodeID, snapshot.GetVersion(t))
	} else {
		cache.log.Infof("open delta watch ID:%d for %s Resources:%v from nodeID: %q", watchID, t, sub.SubscribedResources(), nodeID)
	}

	info.setDeltaResponseWatch(watchID, watch)
	return cache.cancelDeltaWatch(nodeID, watchID), nil
}

func createDeltaResponse(snapshot ResourceSnapshot, watch DeltaResponseWatch, replyIfEmpty bool) (*RawDeltaResponse, error) {
	typeURL := watch.Request.TypeUrl

	resources := snapshot.GetTypeSnapshot(typeURL).GetResources()
	subscribed := watch.subscription.SubscribedResources()

	var resourcesToReturn []*internal.CachedResource
	var deletedResources []string
	returnedResources := maps.Clone(watch.subscription.ReturnedResources())

	if watch.subscription.IsWildcard() {
		resourcesToReturn = make([]*internal.CachedResource, 0, len(resources)+len(subscribed))
	} else {
		resourcesToReturn = make([]*internal.CachedResource, 0, len(subscribed))
	}

	// Check if a resource was not previously returned or has changed version.
	addIfChanged := func(res *internal.CachedResource) error {
		name := res.Name
		resVersion, err := res.GetResourceVersion()
		if err != nil {
			return err
		}
		knownVersion, known := returnedResources[name]
		if known && knownVersion == resVersion {
			return nil
		}
		resourcesToReturn = append(resourcesToReturn, res)
		returnedResources[name] = resVersion
		return nil
	}

	if watch.subscription.IsWildcard() {
		for _, res := range resources {
			// Include wildcard-eligible resources for wildcard subscriptions
			if !res.OnDemandOnly() {
				if err := addIfChanged(res); err != nil {
					return nil, fmt.Errorf("failed to compute resource version for %s: %w", res.Name, err)
				}
			}
		}
	}

	for name := range subscribed {
		res, exist := resources[name]
		if !exist {
			continue
		}
		if err := addIfChanged(res); err != nil {
			return nil, fmt.Errorf("failed to compute resource version for %s: %w", res.Name, err)
		}
	}

	for name := range returnedResources {
		res, exist := resources[name]
		if !exist {
			deletedResources = append(deletedResources, name)
			delete(returnedResources, name)
			continue
		}

		if watch.subscription.IsWildcard() && !res.OnDemandOnly() {
			continue
		}

		if _, explicitlySubscribed := subscribed[name]; !explicitlySubscribed {
			deletedResources = append(deletedResources, name)
			delete(returnedResources, name)
		}
	}

	if len(resourcesToReturn) == 0 && len(deletedResources) == 0 && !replyIfEmpty {
		// Nothing's changed
		return nil, nil
	}

	return &RawDeltaResponse{
		DeltaRequest:      watch.Request,
		SystemVersionInfo: snapshot.GetVersion(typeURL),
		resources:         resourcesToReturn,
		removedResources:  deletedResources,
		returnedResources: returnedResources,
	}, nil
}

// Respond to a delta watch with the provided snapshot value. If the response is nil, there has been no state change.
func (cache *snapshotCache) respondDelta(ctx context.Context, watch DeltaResponseWatch, resp *RawDeltaResponse) error {
	cache.log.Debugf("node: %s, sending delta response for typeURL %s with resources: %v removed resources: %v with wildcard: %t",
		watch.Request.GetNode().GetId(), watch.Request.GetTypeUrl(), getCachedResourceNames(resp.resources), resp.removedResources, watch.subscription.IsWildcard())
	resp.Ctx = ctx
	select {
	case watch.Response <- resp:
		return nil
	case <-ctx.Done():
		return context.Canceled
	}
}

func (cache *snapshotCache) nextDeltaWatchID() int64 {
	return atomic.AddInt64(&cache.deltaWatchCount, 1)
}

// cancellation function for cleaning stale delta watches
func (cache *snapshotCache) cancelDeltaWatch(nodeID string, watchID int64) func() {
	return func() {
		cache.mu.RLock()
		defer cache.mu.RUnlock()
		if info, ok := cache.status[nodeID]; ok {
			info.mu.Lock()
			delete(info.deltaWatches, watchID)
			info.mu.Unlock()
		}
	}
}

// Fetch implements the cache fetch function.
// Fetch is called on multiple streams, so responding to individual names with the same version works.
func (cache *snapshotCache) Fetch(ctx context.Context, request *Request) (Response, error) {
	nodeID := cache.hash.ID(request.GetNode())

	cache.mu.RLock()
	defer cache.mu.RUnlock()

	snapshot, exists := cache.snapshots[nodeID]
	if !exists {
		return nil, fmt.Errorf("missing snapshot for %q", nodeID)
	}

	// Respond only if the request version is distinct from the current snapshot state.
	// It might be beneficial to hold the request since Envoy will re-attempt the refresh.
	version := snapshot.GetVersion(request.GetTypeUrl())
	if request.GetVersionInfo() == version {
		cache.log.Warnf("skip fetch: version up to date")
		return nil, &types.SkipFetchError{}
	}

	resources := snapshot.GetTypeSnapshot(request.GetTypeUrl()).GetResources()
	var resourcesToReturn []*internal.CachedResource
	if len(request.ResourceNames) == 0 {
		resourcesToReturn = slices.Collect(maps.Values(resources))
	} else {
		for _, name := range request.ResourceNames {
			res, ok := resources[name]
			if !ok {
				continue
			}
			resourcesToReturn = append(resourcesToReturn, res)
		}
	}

	return &RawResponse{
		Request:   request,
		Version:   version,
		resources: resourcesToReturn,
		Ctx:       ctx,
	}, nil
}

// GetStatusInfo retrieves the status info for the node.
func (cache *snapshotCache) GetStatusInfo(node string) StatusInfo {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	info, exists := cache.status[node]
	if !exists {
		cache.log.Warnf("node does not exist")
		return nil
	}

	return info
}

// GetStatusKeys retrieves all node IDs in the status map.
func (cache *snapshotCache) GetStatusKeys() []string {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	out := make([]string, 0, len(cache.status))
	for id := range cache.status {
		out = append(out, id)
	}

	return out
}
