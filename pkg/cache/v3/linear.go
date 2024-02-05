// Copyright 2020 Envoyproxy Authors
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
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/log"
)

// LinearCache supports collections of opaque resources. This cache has a
// single collection indexed by resource names and manages resource versions
// internally. It implements the cache interface for a single type URL and
// should be combined with other caches via type URL muxing. It can be used to
// supply EDS entries, for example, uniformly across a fleet of proxies.
type LinearCache struct {
	// Type URL specific to the cache.
	typeURL string
	// Collection of resources indexed by name.
	resources map[string]types.Resource
	// Watches open by clients, indexed by resource name. Whenever resources
	// are changed, the watch is triggered.
	watches map[string]map[int64]ResponseWatch
	// Set of watches for all resources in the collection, indexed by watch id.
	// watch id is unique for sotw watches and is used to index them without requiring
	// the watch itself to be hashable, as well as making logs easier to correlate.
	watchAll map[int64]ResponseWatch
	// Continuously incremented counter used to index sotw watches.
	sotwWatchCount int64
	// Set of delta watches. A delta watch always contain the list of subscribed resources
	// together with its current version
	// version and versionPrefix fields are ignored for delta watches, because we always generate the resource version.
	deltaWatches map[int64]DeltaResponseWatch
	// Continuously incremented counter used to index delta watches.
	deltaWatchCount int64
	// versionMap holds the current hash map of all resources in the cache when delta watches are present.
	// versionMap is only to be used with delta xDS.
	versionMap map[string]string
	// Continuously incremented version.
	version uint64
	// Version prefix to be sent to the clients
	versionPrefix string
	// Versions for each resource by name.
	versionVector map[string]string

	log log.Logger

	mu sync.RWMutex
}

var _ Cache = &LinearCache{}

// Options for modifying the behavior of the linear cache.
type LinearCacheOption func(*LinearCache)

// WithVersionPrefix sets a version prefix of the form "prefixN" in the version info.
// Version prefix can be used to distinguish replicated instances of the cache, in case
// a client re-connects to another instance.
func WithVersionPrefix(prefix string) LinearCacheOption {
	return func(cache *LinearCache) {
		cache.versionPrefix = prefix
	}
}

// WithInitialResources initializes the initial set of resources.
func WithInitialResources(resources map[string]types.Resource) LinearCacheOption {
	return func(cache *LinearCache) {
		cache.resources = resources
	}
}

func WithLogger(log log.Logger) LinearCacheOption {
	return func(cache *LinearCache) {
		cache.log = log
	}
}

// NewLinearCache creates a new cache. See the comments on the struct definition.
func NewLinearCache(typeURL string, opts ...LinearCacheOption) *LinearCache {
	out := &LinearCache{
		typeURL:       typeURL,
		resources:     make(map[string]types.Resource),
		watches:       make(map[string]map[int64]ResponseWatch),
		watchAll:      make(map[int64]ResponseWatch),
		deltaWatches:  make(map[int64]DeltaResponseWatch),
		versionMap:    nil,
		version:       0,
		versionVector: make(map[string]string),
		log:           log.NewDefaultLogger(),
	}
	for _, opt := range opts {
		opt(out)
	}
	for name := range out.resources {
		out.versionVector[name] = out.getVersion()
	}
	return out
}

func (cache *LinearCache) computeSotwResponse(watch ResponseWatch, ignoreReturnedResources bool) *RawResponse {
	var changedResources []string
	var removedResources []string

	knownVersions := watch.subscription.ReturnedResources()
	if ignoreReturnedResources {
		// The response will include all resources, with no regards of resources potentially already returned.
		knownVersions = make(map[string]string)
	}

	if watch.subscription.IsWildcard() {
		for resourceName, version := range cache.versionVector {
			knownVersion, ok := knownVersions[resourceName]
			if !ok {
				// This resource is not yet known by the client (new resource added in the cache or newly subscribed).
				changedResources = append(changedResources, resourceName)
			} else if knownVersion != version {
				// The client knows an outdated version.
				changedResources = append(changedResources, resourceName)
			}
		}

		// Negative check to identify resources that have been removed in the cache.
		// Sotw does not support returning "deletions", but in the case of full state resources
		// a response must then be returned.
		for resourceName := range knownVersions {
			if _, ok := cache.versionVector[resourceName]; !ok {
				removedResources = append(removedResources, resourceName)
			}
		}
	} else {
		for resourceName := range watch.subscription.SubscribedResources() {
			version, exists := cache.versionVector[resourceName]
			knownVersion, known := knownVersions[resourceName]
			if !exists {
				if known {
					// This resource was removed from the cache. If the type requires full state
					// we need to return a response.
					removedResources = append(removedResources, resourceName)
				}
				continue
			}

			if !known {
				// This resource is not yet known by the client (new resource added in the cache or newly subscribed).
				changedResources = append(changedResources, resourceName)
			} else if knownVersion != version {
				// The client knows an outdated version.
				changedResources = append(changedResources, resourceName)
			}
		}

		for resourceName := range knownVersions {
			// If the subscription no longer watches a resource,
			// we mark it as unknown on the client side to ensure it will be resent to the client if subscribing again later on.
			if _, ok := watch.subscription.SubscribedResources()[resourceName]; !ok {
				removedResources = append(removedResources, resourceName)
			}
		}
	}

	if len(changedResources) == 0 && len(removedResources) == 0 && !ignoreReturnedResources {
		// Nothing changed.
		return nil
	}

	returnedVersions := make(map[string]string, len(watch.subscription.ReturnedResources()))
	// Clone the current returned versions. The cache should not alter the subscription
	for resourceName, version := range watch.subscription.ReturnedResources() {
		returnedVersions[resourceName] = version
	}

	cacheVersion := cache.getVersion()
	var resources []types.ResourceWithTTL

	switch {
	// Depending on the type, the response will only include changed resources or all of them
	case !ResourceRequiresFullStateInSotw(cache.typeURL):
		// changedResources is already filtered based on the subscription.
		// TODO(valerian-roche): if the only change is a removal in the subscription,
		// or a watched resource getting deleted, this might send an empty reply.
		// While this does not violate the protocol, we might want to avoid it.
		resources = make([]types.ResourceWithTTL, 0, len(changedResources))
		for _, resourceName := range changedResources {
			resources = append(resources, types.ResourceWithTTL{Resource: cache.resources[resourceName]})
			returnedVersions[resourceName] = cache.versionVector[resourceName]
		}
	case watch.subscription.IsWildcard():
		// Include all resources for the type.
		resources = make([]types.ResourceWithTTL, 0, len(cache.resources))
		for resourceName, resource := range cache.resources {
			resources = append(resources, types.ResourceWithTTL{Resource: resource})
			returnedVersions[resourceName] = cache.versionVector[resourceName]
		}
	default:
		// Include all resources matching the subscription, with no concern on whether
		// it has been updated or not.
		requestedResources := watch.subscription.SubscribedResources()
		// The linear cache could be very large (e.g. containing all potential CLAs)
		// Therefore drives on the subscription requested resources.
		resources = make([]types.ResourceWithTTL, 0, len(requestedResources))
		for resourceName := range requestedResources {
			resource, ok := cache.resources[resourceName]
			if !ok {
				continue
			}
			resources = append(resources, types.ResourceWithTTL{Resource: resource})
			returnedVersions[resourceName] = cache.versionVector[resourceName]
		}
	}

	// Cleanup resources no longer existing in the cache or no longer subscribed.
	// In sotw we cannot return those if not full state,
	// but this ensures we detect unsubscription then resubscription.
	for _, resourceName := range removedResources {
		delete(returnedVersions, resourceName)
	}

	if !ignoreReturnedResources && !ResourceRequiresFullStateInSotw(cache.typeURL) && len(resources) == 0 {
		// If the request is not the initial one, and the type for not require full updates,
		// do not return if noting is to be set
		// For full-state resources an empty response does have a semantic meaning
		return nil
	}

	return &RawResponse{
		Request:           watch.Request,
		Resources:         resources,
		ReturnedResources: returnedVersions,
		Version:           cacheVersion,
		Ctx:               context.Background(),
	}
}

func (cache *LinearCache) notifyAll(modified map[string]struct{}) {
	// Gather the list of non-wildcard watches impacted by the modified resources.
	watches := make(map[int64]ResponseWatch)
	for name := range modified {
		for watchID, watch := range cache.watches[name] {
			watches[watchID] = watch
		}
	}
	for watchID, watch := range watches {
		response := cache.computeSotwResponse(watch, false)
		if response != nil {
			watch.Response <- response
			cache.removeWatch(watchID, watch.subscription)
		} else {
			cache.log.Warnf("[Linear cache] Watch %d detected as triggered did not get notified", watchID)
		}
	}

	for watchID, watch := range cache.watchAll {
		response := cache.computeSotwResponse(watch, false)
		if response != nil {
			watch.Response <- response
			delete(cache.watchAll, watchID)
		} else {
			cache.log.Warnf("[Linear cache] Watch %d detected as triggered did not get notified", watchID)
		}
	}

	// Building the version map has a very high cost when using SetResources to do full updates.
	// As it is only used with delta watches, it is only maintained when applicable.
	if cache.versionMap != nil {
		err := cache.updateVersionMap(modified)
		if err != nil {
			cache.log.Errorf("failed to update version map: %v", err)
		}

		for id, watch := range cache.deltaWatches {
			if !watch.subscription.WatchesResources(modified) {
				continue
			}

			res := cache.respondDelta(watch.Request, watch.Response, watch.subscription)
			if res != nil {
				delete(cache.deltaWatches, id)
			}
		}
	}
}

func (cache *LinearCache) respondDelta(request *DeltaRequest, value chan DeltaResponse, sub Subscription) *RawDeltaResponse {
	resp := createDeltaResponse(context.Background(), request, sub, resourceContainer{
		resourceMap:   cache.resources,
		versionMap:    cache.versionMap,
		systemVersion: cache.getVersion(),
	})

	// Only send a response if there were changes
	// We want to respond immediately for the first wildcard request in a stream, even if the response is empty
	// otherwise, envoy won't complete initialization
	if len(resp.Resources) > 0 || len(resp.RemovedResources) > 0 || (sub.IsWildcard() && request.ResponseNonce == "") {
		if cache.log != nil {
			cache.log.Debugf("[linear cache] node: %s, sending delta response for typeURL %s with resources: %v removed resources: %v with wildcard: %t",
				request.GetNode().GetId(), request.GetTypeUrl(), GetResourceNames(resp.Resources), resp.RemovedResources, sub.IsWildcard())
		}
		value <- resp
		return resp
	}
	return nil
}

// UpdateResource updates a resource in the collection.
func (cache *LinearCache) UpdateResource(name string, res types.Resource) error {
	if res == nil {
		return errors.New("nil resource")
	}
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.version++
	cache.versionVector[name] = cache.getVersion()
	cache.resources[name] = res

	// TODO: batch watch closures to prevent rapid updates
	cache.notifyAll(map[string]struct{}{name: {}})

	return nil
}

// DeleteResource removes a resource in the collection.
func (cache *LinearCache) DeleteResource(name string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.version++
	delete(cache.versionVector, name)
	delete(cache.resources, name)

	// TODO: batch watch closures to prevent rapid updates
	cache.notifyAll(map[string]struct{}{name: {}})
	return nil
}

// UpdateResources updates/deletes a list of resources in the cache.
// Calling UpdateResources instead of iterating on UpdateResource and DeleteResource
// is significantly more efficient when using delta or wildcard watches.
func (cache *LinearCache) UpdateResources(toUpdate map[string]types.Resource, toDelete []string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.version++
	cacheVersion := cache.getVersion()

	modified := make(map[string]struct{}, len(toUpdate)+len(toDelete))
	for name, resource := range toUpdate {
		cache.versionVector[name] = cacheVersion
		cache.resources[name] = resource
		modified[name] = struct{}{}
	}
	for _, name := range toDelete {
		delete(cache.versionVector, name)
		delete(cache.resources, name)
		modified[name] = struct{}{}
	}

	cache.notifyAll(modified)

	return nil
}

// SetResources replaces current resources with a new set of resources.
// This function is useful for wildcard xDS subscriptions.
// This way watches that are subscribed to all resources are triggered only once regardless of how many resources are changed.
func (cache *LinearCache) SetResources(resources map[string]types.Resource) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.version++
	cacheVersion := cache.getVersion()

	modified := map[string]struct{}{}
	// Collect deleted resource names.
	for name := range cache.resources {
		if _, found := resources[name]; !found {
			delete(cache.versionVector, name)
			modified[name] = struct{}{}
		}
	}

	cache.resources = resources

	// Collect changed resource names.
	// We assume all resources passed to SetResources are changed.
	// Otherwise we would have to do proto.Equal on resources which is pretty expensive operation
	for name := range resources {
		cache.versionVector[name] = cacheVersion
		modified[name] = struct{}{}
	}

	cache.notifyAll(modified)
}

// GetResources returns current resources stored in the cache
func (cache *LinearCache) GetResources() map[string]types.Resource {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	// create a copy of our internal storage to avoid data races
	// involving mutations of our backing map
	resources := make(map[string]types.Resource, len(cache.resources))
	for k, v := range cache.resources {
		resources[k] = v
	}
	return resources
}

func (cache *LinearCache) CreateWatch(request *Request, sub Subscription, value chan Response) (func(), error) {
	if request.GetTypeUrl() != cache.typeURL {
		return nil, fmt.Errorf("request type %s does not match cache type %s", request.GetTypeUrl(), cache.typeURL)
	}

	// If the request does not include a version the client considers it has no current state.
	// In this case we will always reply to allow proper initialization of dependencies in the client.
	ignoreCurrentSubscriptionResources := request.GetVersionInfo() == ""
	if !strings.HasPrefix(request.GetVersionInfo(), cache.versionPrefix) {
		// If the version of the request does not match the cache prefix, we will send a response in all cases to match the legacy behavior.
		ignoreCurrentSubscriptionResources = true
		cache.log.Debugf("[linear cache] received watch with version %s not matching the cache prefix %s. Will return all known resources", request.GetVersionInfo(), cache.versionPrefix)
	}

	// A major difference between delta and sotw is the ability to not resend everything when connecting to a new control-plane
	// In delta the request provides the version of the resources it does know, even if the request is wildcard or does request more resources
	// In sotw the request only provides the global version of the control-plane, and there is no way for the control-plane to know if resources have
	// been added since in the requested resources. In the context of generalized wildcard, even wildcard could be new, and taking the assumption
	// that wildcard implies that the client already knows all resources at the given version is no longer true.
	// We could optimize the reconnection case here if:
	//  - we take the assumption that clients will not start requesting wildcard while providing a version. We could then ignore requests providing the resources.
	//  - we use the version as some form of hash of resources known, and we can then consider it as a way to correctly verify whether all resources are unchanged.
	// For now it is not done as:
	//  - for the first case, while the protocol documentation does not explicitly mention the case, it does not mark it impossible and explicitly references unsubscribing from wildcard.
	//  - for the second one we could likely do it with little difficulty if need be, but if users rely on the current monotonic version it could impact their callbacks implementations.
	watch := ResponseWatch{Request: request, Response: value, subscription: sub}

	cache.mu.Lock()
	defer cache.mu.Unlock()

	response := cache.computeSotwResponse(watch, ignoreCurrentSubscriptionResources)
	if response != nil {
		cache.log.Debugf("replying to the watch with resources %v (subscription values %v, known %v)", response.GetReturnedResources(), sub.SubscribedResources(), sub.ReturnedResources())
		watch.Response <- response
		return func() {}, nil
	}

	watchID := cache.nextSotwWatchID()
	// Create open watches since versions are up to date.
	if sub.IsWildcard() {
		cache.log.Infof("[linear cache] open watch %d for %s all resources, system version %q", watchID, cache.typeURL, cache.getVersion())
		cache.watchAll[watchID] = watch
		return func() {
			cache.mu.Lock()
			defer cache.mu.Unlock()
			delete(cache.watchAll, watchID)
		}, nil
	}

	cache.log.Infof("[linear cache] open watch %d for %s resources %v, system version %q", watchID, cache.typeURL, sub.SubscribedResources(), cache.getVersion())
	for name := range sub.SubscribedResources() {
		set, exists := cache.watches[name]
		if !exists {
			set = make(map[int64]ResponseWatch)
			cache.watches[name] = set
		}
		set[watchID] = watch
	}
	return func() {
		cache.mu.Lock()
		defer cache.mu.Unlock()
		cache.removeWatch(watchID, watch.subscription)
	}, nil
}

func (cache *LinearCache) nextSotwWatchID() int64 {
	next := atomic.AddInt64(&cache.sotwWatchCount, 1)
	if next < 0 {
		panic("watch id count overflow")
	}
	return next
}

// Must be called under lock
func (cache *LinearCache) removeWatch(watchID int64, sub Subscription) {
	// Make sure we clean the watch for ALL resources it might be associated with,
	// as the channel will no longer be listened to
	for resource := range sub.SubscribedResources() {
		resourceWatches := cache.watches[resource]
		delete(resourceWatches, watchID)
		if len(resourceWatches) == 0 {
			delete(cache.watches, resource)
		}
	}
}

func (cache *LinearCache) CreateDeltaWatch(request *DeltaRequest, sub Subscription, value chan DeltaResponse) (func(), error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if cache.versionMap == nil {
		// If we had no previously open delta watches, we need to build the version map for the first time.
		// The version map will not be destroyed when the last delta watch is removed.
		// This avoids constantly rebuilding when only a few delta watches are open.
		modified := map[string]struct{}{}
		for name := range cache.resources {
			modified[name] = struct{}{}
		}
		err := cache.updateVersionMap(modified)
		if err != nil && cache.log != nil {
			cache.log.Errorf("failed to update version map: %v", err)
		}
	}
	response := cache.respondDelta(request, value, sub)

	// if respondDelta returns nil this means that there is no change in any resource version
	// create a new watch accordingly
	if response == nil {
		watchID := cache.nextDeltaWatchID()
		if cache.log != nil {
			cache.log.Infof("[linear cache] open delta watch ID:%d for %s Resources:%v, system version %q", watchID,
				cache.typeURL, sub.SubscribedResources(), cache.getVersion())
		}

		cache.deltaWatches[watchID] = DeltaResponseWatch{Request: request, Response: value, subscription: sub}

		return cache.cancelDeltaWatch(watchID), nil
	}

	return nil, nil
}

func (cache *LinearCache) updateVersionMap(modified map[string]struct{}) error {
	if cache.versionMap == nil {
		cache.versionMap = make(map[string]string, len(modified))
	}
	for name := range modified {
		r, ok := cache.resources[name]
		if !ok {
			// The resource was deleted
			delete(cache.versionMap, name)
			continue
		}
		// hash our version in here and build the version map
		marshaledResource, err := MarshalResource(r)
		if err != nil {
			return err
		}
		v := HashResource(marshaledResource)
		if v == "" {
			return errors.New("failed to build resource version")
		}

		cache.versionMap[name] = v
	}
	return nil
}

func (cache *LinearCache) getVersion() string {
	return cache.versionPrefix + strconv.FormatUint(cache.version, 10)
}

// cancellation function for cleaning stale watches
func (cache *LinearCache) cancelDeltaWatch(watchID int64) func() {
	return func() {
		cache.mu.Lock()
		defer cache.mu.Unlock()
		delete(cache.deltaWatches, watchID)
	}
}

func (cache *LinearCache) nextDeltaWatchID() int64 {
	return atomic.AddInt64(&cache.deltaWatchCount, 1)
}

func (cache *LinearCache) Fetch(context.Context, *Request) (Response, error) {
	return nil, errors.New("not implemented")
}

// Number of resources currently on the cache.
// As GetResources is building a clone it is expensive to get metrics otherwise.
func (cache *LinearCache) NumResources() int {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	return len(cache.resources)
}

// Number of active watches for a resource name.
func (cache *LinearCache) NumWatches(name string) int {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	return len(cache.watches[name]) + len(cache.watchAll)
}

// Number of active delta watches.
func (cache *LinearCache) NumDeltaWatches() int {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	return len(cache.deltaWatches)
}
