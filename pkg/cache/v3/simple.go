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
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/log"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
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

	// GetResourcesAndTTL returns all resources of the type indicted by
	// typeURL, together with their TTL.
	GetResourcesAndTTL(typeURL string) map[string]types.ResourceWithTTL

	// GetResources returns all resources of the type indicted by
	// typeURL. This is identical to GetResourcesAndTTL, except that
	// the TTL is omitted.
	GetResources(typeURL string) map[string]types.Resource

	// ConstructVersionMap is a hint that a delta watch will soon make a
	// call to GetVersionMap. The snapshot should construct an internal
	// opaque version string for each collection of resource types.
	ConstructVersionMap() error

	// GetVersionMap returns a map of resource name to resource version for
	// all the resources of type indicated by typeURL.
	GetVersionMap(typeURL string) map[string]string
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
	snapshot := cache.snapshots[node]
	if info, ok := cache.status[node]; ok {
		info.mu.Lock()
		for id, watch := range info.watches {
			// Respond with the current version regardless of whether the version has changed.
			version := snapshot.GetVersion(watch.Request.TypeUrl)
			resources := snapshot.GetResourcesAndTTL(watch.Request.TypeUrl)

			// TODO(snowp): Construct this once per type instead of once per watch.
			resourcesWithTTL := map[string]types.ResourceWithTTL{}
			for k, v := range resources {
				if v.TTL != nil {
					resourcesWithTTL[k] = v
				}
			}

			if len(resourcesWithTTL) == 0 {
				continue
			}
			cache.log.Debugf("respond open watch %d%v with heartbeat for version %q", id, watch.Request.ResourceNames, version)
			err := cache.respond(ctx, watch.Request, watch.StreamState, watch.Response, resourcesWithTTL, version, true)
			if err != nil {
				cache.log.Errorf("received error when attempting to respond to watches: %v", err)
			}

			// The watch must be deleted and we must rely on the client to ack this response to create a new watch.
			delete(info.watches, id)
		}
		info.mu.Unlock()
	}
}

// SetSnapshotCacheContext updates a snapshot for a node.
func (cache *snapshotCache) SetSnapshot(ctx context.Context, node string, snapshot ResourceSnapshot) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// update the existing entry
	cache.snapshots[node] = snapshot

	// trigger existing watches for which version changed
	if info, ok := cache.status[node]; ok {
		info.mu.Lock()
		defer info.mu.Unlock()
		for id, watch := range info.watches {
			version := snapshot.GetVersion(watch.Request.TypeUrl)
			if version != watch.Request.VersionInfo {
				cache.log.Debugf("respond open watch %d %s%v with new version %q", id, watch.Request.TypeUrl, watch.Request.ResourceNames, version)

				resources := snapshot.GetResourcesAndTTL(watch.Request.TypeUrl)
				err := cache.respond(ctx, watch.Request, watch.StreamState, watch.Response, resources, version, false)
				if err != nil {
					return err
				}

				// discard the watch
				delete(info.watches, id)
			}
		}

		// We only calculate version hashes when using delta. We don't
		// want to do this when using SOTW so we can avoid unnecessary
		// computational cost if not using delta.
		if len(info.deltaWatches) > 0 {
			err := snapshot.ConstructVersionMap()
			if err != nil {
				return err
			}
		}

		// process our delta watches
		for id, watch := range info.deltaWatches {
			res, err := cache.respondDelta(
				ctx,
				snapshot,
				watch.Request,
				watch.Response,
				watch.StreamState,
			)
			if err != nil {
				return err
			}
			// If we detect a nil response here, that means there has been no state change
			// so we don't want to respond or remove any existing resource watches
			if res != nil {
				delete(info.deltaWatches, id)
			}
		}
	}

	return nil
}

// GetSnapshots gets the snapshot for a node, and returns an error if not found.
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

// superset checks that all resources are listed in the names set.
func superset(names map[string]struct{}, resources map[string]types.ResourceWithTTL) error {
	for resourceName := range resources {
		if _, exists := names[resourceName]; !exists {
			return fmt.Errorf("%q not listed", resourceName)
		}
	}
	return nil
}

// CreateWatch returns a watch for an xDS request.
func (cache *snapshotCache) CreateWatch(request *Request, streamState stream.StreamState, value chan Response) func() {
	nodeID := cache.hash.ID(request.Node)

	cache.mu.Lock()
	defer cache.mu.Unlock()

	info, ok := cache.status[nodeID]
	if !ok {
		info = newStatusInfo(request.Node)
		cache.status[nodeID] = info
	}

	// update last watch request time
	info.mu.Lock()
	info.lastWatchRequestTime = time.Now()
	info.mu.Unlock()

	var version string

	snapshot, exists := cache.snapshots[nodeID]
	if exists {
		version = snapshot.GetVersion(request.TypeUrl)
	}

	// Main cases here:
	//  * the node id is in the cache
	//    * and the version is not matching -> update to be sent
	//    * and the version is matching, but new resources are present in the watch (if explicitly subscribed or switched to wildcard) -> update to be sent
	//    * and the version is matching, with no new resources -> leave a watch opened
	//  * the node id does not exist in the snapshot cache -> leave a watch opened

	if exists && request.VersionInfo != version {
		// Version has changed, send an update
		resources := snapshot.GetResourcesAndTTL(request.TypeUrl)
		if err := cache.respond(context.Background(), request, streamState, value, resources, version, false); err != nil {
			cache.log.Errorf("failed to send a response for %s%v to nodeID %q: %s", request.TypeUrl,
				streamState.GetSubscribedResourceNames(), nodeID, err)
		}
		return nil
	}

	if exists {
		// Version has not changed.
		// We reply immediately if and only if there are resources now requested that were not returned previously.
		clientKnownResourceNames := streamState.GetResourceVersions()
		snapshotResources := snapshot.GetResourcesAndTTL(request.TypeUrl)
		missingResourceFound := false
		if !streamState.IsWildcard() {
			for r := range streamState.GetSubscribedResourceNames() {
				if _, ok := clientKnownResourceNames[r]; !ok {
					if _, existsInSnapshot := snapshotResources[r]; existsInSnapshot {
						missingResourceFound = true
						break
					}
				}
			}
		} else {
			// The stream might have been updated to wildcard, ensure all resources have already been returned
			for r := range snapshot.GetResourcesAndTTL(request.TypeUrl) {
				if _, ok := clientKnownResourceNames[r]; !ok {
					missingResourceFound = true
					break
				}
			}
		}

		cache.log.Debugf("nodeID %q with watch (wildcard: %t, subscribed %s%v) and known %v. Found missing resource %t", nodeID,
			request.TypeUrl, streamState.IsWildcard(), streamState.GetSubscribedResourceNames(), clientKnownResourceNames, missingResourceFound)

		if missingResourceFound {
			if err := cache.respond(context.Background(), request, streamState, value, snapshotResources, version, false); err != nil {
				cache.log.Errorf("failed to send a response for %s%v to nodeID %q: %s", request.TypeUrl,
					streamState.GetSubscribedResourceNames(), nodeID, err)
			}
			return nil
		}
	}

	// if we haven't sent a response already, leave a watch opened
	watchID := cache.nextWatchID()
	cache.log.Debugf("open watch %d for %s%v from nodeID %q, version %q", watchID, request.TypeUrl, streamState.GetSubscribedResourceNames(), nodeID, request.VersionInfo)
	info.mu.Lock()
	info.watches[watchID] = ResponseWatch{Request: request, StreamState: streamState, Response: value}
	info.mu.Unlock()
	return cache.cancelWatch(nodeID, watchID)
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

// Respond to a watch with the snapshot value. The value channel should have capacity not to block.
// TODO(kuat) do not respond always, see issue https://github.com/envoyproxy/go-control-plane/issues/46
func (cache *snapshotCache) respond(ctx context.Context, request *Request, streamState stream.StreamState, value chan Response, resources map[string]types.ResourceWithTTL, version string, heartbeat bool) error {
	// for ADS, the request names must match the snapshot names
	// if they do not, then the watch is never responded, and it is expected that envoy makes another request
	if !streamState.IsWildcard() && cache.ads {
		if err := superset(streamState.GetSubscribedResourceNames(), resources); err != nil {
			cache.log.Warnf("ADS mode: not responding to request: %v", err)
			return nil
		}
	}

	cache.log.Debugf("respond %s%v version %q with version %q", request.TypeUrl, streamState.GetSubscribedResourceNames(), request.VersionInfo, version)

	select {
	case value <- createResponse(ctx, request, streamState, resources, version, heartbeat):
		return nil
	case <-ctx.Done():
		return context.Canceled
	}
}

func createResponse(ctx context.Context, request *Request, streamState stream.StreamState, resources map[string]types.ResourceWithTTL, version string, heartbeat bool) Response {
	var filtered []types.ResourceWithTTL
	var resourceVersions map[string]string

	// Reply only with the requested resources. Envoy may ask each resource
	// individually in a separate stream. It is ok to reply with the same version
	// on separate streams since requests do not share their response versions.
	if !streamState.IsWildcard() {
		filtered = make([]types.ResourceWithTTL, 0, len(streamState.GetSubscribedResourceNames()))
		resourceVersions = make(map[string]string, len(streamState.GetSubscribedResourceNames()))
		for resourceName := range streamState.GetSubscribedResourceNames() {
			if resource, ok := resources[resourceName]; ok {
				filtered = append(filtered, resource)
				resourceVersions[resourceName] = version
			}
		}
	} else {
		filtered = make([]types.ResourceWithTTL, 0, len(resources))
		resourceVersions = make(map[string]string, len(resources))
		for name, resource := range resources {
			filtered = append(filtered, resource)
			resourceVersions[name] = version
		}
	}

	return &RawResponse{
		Request:        request,
		Version:        version,
		Resources:      filtered,
		NextVersionMap: resourceVersions,
		Heartbeat:      heartbeat,
		Ctx:            ctx,
	}
}

// CreateDeltaWatch returns a watch for a delta xDS request which implements the Simple SnapshotCache.
func (cache *snapshotCache) CreateDeltaWatch(request *DeltaRequest, state stream.StreamState, value chan DeltaResponse) func() {
	nodeID := cache.hash.ID(request.Node)
	t := request.GetTypeUrl()

	cache.mu.Lock()
	defer cache.mu.Unlock()

	info, ok := cache.status[nodeID]
	if !ok {
		info = newStatusInfo(request.Node)
		cache.status[nodeID] = info
	}

	// update last watch request time
	info.setLastDeltaWatchRequestTime(time.Now())

	// find the current cache snapshot for the provided node
	snapshot, exists := cache.snapshots[nodeID]

	// There are three different cases that leads to a delayed watch trigger:
	// - no snapshot exists for the requested nodeID
	// - a snapshot exists, but we failed to initialize its version map
	// - we attempted to issue a response, but the caller is already up to date
	delayedResponse := !exists
	if exists {
		err := snapshot.ConstructVersionMap()
		if err != nil {
			cache.log.Errorf("failed to compute version for snapshot resources inline: %s", err)
		}
		response, err := cache.respondDelta(context.Background(), snapshot, request, value, state)
		if err != nil {
			cache.log.Errorf("failed to respond with delta response: %s", err)
		}

		delayedResponse = response == nil
	}

	if delayedResponse {
		watchID := cache.nextDeltaWatchID()

		if exists {
			cache.log.Infof("open delta watch ID:%d for %s Resources:%v from nodeID: %q,  version %q", watchID, t, state.GetSubscribedResourceNames(), nodeID, snapshot.GetVersion(t))
		} else {
			cache.log.Infof("open delta watch ID:%d for %s Resources:%v from nodeID: %q", watchID, t, state.GetSubscribedResourceNames(), nodeID)
		}

		info.setDeltaResponseWatch(watchID, DeltaResponseWatch{Request: request, Response: value, StreamState: state})
		return cache.cancelDeltaWatch(nodeID, watchID)
	}

	return nil
}

// Respond to a delta watch with the provided snapshot value. If the response is nil, there has been no state change.
func (cache *snapshotCache) respondDelta(ctx context.Context, snapshot ResourceSnapshot, request *DeltaRequest, value chan DeltaResponse, state stream.StreamState) (*RawDeltaResponse, error) {
	resp := createDeltaResponse(ctx, request, state, resourceContainer{
		resourceMap:   snapshot.GetResources(request.TypeUrl),
		versionMap:    snapshot.GetVersionMap(request.TypeUrl),
		systemVersion: snapshot.GetVersion(request.TypeUrl),
	})

	// Only send a response if there were changes
	// We want to respond immediately for the first wildcard request in a stream, even if the response is empty
	// otherwise, envoy won't complete initialization
	if len(resp.Resources) > 0 || len(resp.RemovedResources) > 0 || (state.IsWildcard() && state.IsFirst()) {
		if cache.log != nil {
			cache.log.Debugf("node: %s, sending delta response with resources: %v removed resources %v wildcard: %t",
				request.GetNode().GetId(), resp.Resources, resp.RemovedResources, state.IsWildcard())
		}
		select {
		case value <- resp:
			return resp, nil
		case <-ctx.Done():
			return resp, context.Canceled
		}
	}
	return nil, nil
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
	nodeID := cache.hash.ID(request.Node)

	cache.mu.RLock()
	defer cache.mu.RUnlock()

	if snapshot, exists := cache.snapshots[nodeID]; exists {
		// Respond only if the request version is distinct from the current snapshot state.
		// It might be beneficial to hold the request since Envoy will re-attempt the refresh.
		version := snapshot.GetVersion(request.TypeUrl)
		if request.VersionInfo == version {
			cache.log.Warnf("skip fetch: version up to date")
			return nil, &types.SkipFetchError{}
		}

		// This code is duplicated from sotw/server.go
		state := stream.NewStreamState(len(request.ResourceNames) == 0, nil)
		wantedResources := make(map[string]struct{}, len(request.ResourceNames))
		for _, resourceName := range request.ResourceNames {
			if resourceName == "*" {
				state.SetWildcard(true)
				continue
			}
			wantedResources[resourceName] = struct{}{}
		}
		state.SetSubscribedResourceNames(wantedResources)

		resources := snapshot.GetResourcesAndTTL(request.TypeUrl)
		out := createResponse(ctx, request, state, resources, version, false)
		return out, nil
	}

	return nil, fmt.Errorf("missing snapshot for %q", nodeID)
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
