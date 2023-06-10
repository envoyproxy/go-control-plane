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

type watches = map[chan Response]Request

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
	watches map[string]watches
	// Set of watches for all resources in the collection
	watchAll watches
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
	versionVector map[string]uint64

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
		for name := range resources {
			cache.versionVector[name] = 0
		}
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
		watches:       make(map[string]watches),
		watchAll:      make(watches),
		deltaWatches:  make(map[int64]DeltaResponseWatch),
		versionMap:    nil,
		version:       0,
		versionVector: make(map[string]uint64),
	}
	for _, opt := range opts {
		opt(out)
	}
	return out
}

func (cache *LinearCache) respond(req Request, value chan Response, staleResources map[string]struct{}) {
	var resources []resourceWithTTLAndName
	// TODO: optimize the resources slice creations across different clients
	if len(staleResources) == 0 {
		resources = make([]resourceWithTTLAndName, 0, len(cache.resources))
		for name, resource := range cache.resources {
			resources = append(resources, ttlResource{name: name, resource: resource})
		}
	} else {
		resources = make([]resourceWithTTLAndName, 0, len(staleResources))
		for name := range staleResources {
			resource := cache.resources[name]
			if resource != nil {
				resources = append(resources, ttlResource{name: name, resource: resource})
			}
		}
	}
	value <- &rawResponse{
		request:   req,
		resources: resources,
		version:   cache.getVersion(),
		ctx:       context.Background(),
	}
}

func (cache *LinearCache) notifyAll(modified map[string]struct{}) {
	type staleWatch struct {
		staleResources map[string]struct{}
		req            Request
	}

	// de-duplicate watches that need to be responded
	notifyList := make(map[chan Response]staleWatch)

	for name := range modified {
		for w, req := range cache.watches[name] {
			if watch, ok := notifyList[w]; ok {
				watch.staleResources[name] = struct{}{}
			} else {
				notifyList[w] = staleWatch{
					staleResources: map[string]struct{}{name: {}},
					req:            req,
				}
			}
		}
		delete(cache.watches, name)
	}
	for value, staleWatch := range notifyList {
		cache.respond(staleWatch.req, value, staleWatch.staleResources)
	}
	for value, req := range cache.watchAll {
		cache.respond(req, value, nil)
	}
	cache.watchAll = make(watches)

	// Building the version map has a very high cost when using SetResources to do full updates.
	// As it is only used with delta watches, it is only maintained when applicable.
	if cache.versionMap != nil {
		err := cache.updateVersionMap(modified)
		if err != nil {
			cache.log.Errorf("failed to update version map: %v", err)
		}

		for id, watch := range cache.deltaWatches {
			if !watch.subscriptionState.WatchesResources(modified) {
				continue
			}

			res := cache.respondDelta(watch.Request, watch.Response, watch.subscriptionState)
			if res != nil {
				delete(cache.deltaWatches, id)
			}
		}
	}
}

func (cache *LinearCache) respondDelta(request DeltaRequest, value chan DeltaResponse, clientState SubscriptionState) *rawDeltaResponse {
	resp := createDeltaResponse(context.Background(), request, clientState, resourceContainer{
		resourceMap:   cache.resources,
		versionMap:    cache.versionMap,
		systemVersion: cache.getVersion(),
	})

	// Only send a response if there were changes
	if len(resp.resources) > 0 || len(resp.removedResources) > 0 {
		if cache.log != nil {
			cache.log.Debugf("[linear cache] node: %s, sending delta response for typeURL %s with resources: %v removed resources: %v with wildcard: %t",
				request.GetNode().GetId(), request.GetTypeUrl(), getResourceNames(resp.resources), resp.removedResources, clientState.IsWildcard())
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
	cache.versionVector[name] = cache.version
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

	modified := make(map[string]struct{}, len(toUpdate)+len(toDelete))
	for name, resource := range toUpdate {
		cache.versionVector[name] = cache.version
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
		cache.versionVector[name] = cache.version
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

func (cache *LinearCache) CreateWatch(request Request, clientState SubscriptionState, value chan Response) (func(), error) {
	if request.GetTypeUrl() != cache.typeURL {
		value <- nil
		return nil, fmt.Errorf("request type %s does not match cache type %s", request.GetTypeUrl(), cache.typeURL)
	}
	// If the version is not up to date, check whether any requested resource has
	// been updated between the last version and the current version. This avoids the problem
	// of sending empty updates whenever an irrelevant resource changes.
	stale := false
	var staleResources map[string]struct{} // empty means all

	// strip version prefix if it is present
	var lastVersion uint64
	var err error
	if strings.HasPrefix(request.GetVersionInfo(), cache.versionPrefix) {
		lastVersion, err = strconv.ParseUint(request.GetVersionInfo()[len(cache.versionPrefix):], 0, 64)
	} else {
		err = errors.New("mis-matched version prefix")
	}

	cache.mu.Lock()
	defer cache.mu.Unlock()

	if err != nil {
		// Version was not set or cannot be parsed
		// It is ignored and all resources are returned
		stale = true
		staleResources = clientState.GetSubscribedResources()
	} else if clientState.IsWildcard() {
		// For wildcard watches, all resources are sent if the version no longer matches
		// As we expect at least one resource to have changed
		stale = lastVersion != cache.version
	} else {
		// For non-wildcard watches, compare the provided version to the version of subscribed objects
		// If at least one version is changed, we send back a response
		staleResources = make(map[string]struct{})
		for name := range clientState.GetSubscribedResources() {
			if _, ok := clientState.GetKnownResources()[name]; !ok {
				// A resource is newly requested. Even if the version of the request matches or is higher than the object
				// we must return it
				staleResources[name] = struct{}{}
				continue
			}
			// When a resource is removed, its version defaults 0 and it is not considered stale.
			if lastVersion < cache.versionVector[name] {
				stale = true
				staleResources[name] = struct{}{}
			}
		}
	}
	if stale {
		cache.respond(request, value, staleResources)
		return nil, nil
	}
	// Create open watches since versions are up to date.
	if clientState.IsWildcard() {
		cache.watchAll[value] = request
		return func() {
			cache.mu.Lock()
			defer cache.mu.Unlock()
			delete(cache.watchAll, value)
		}, nil
	}
	for name := range clientState.GetSubscribedResources() {
		set, exists := cache.watches[name]
		if !exists {
			set = make(watches)
			cache.watches[name] = set
		}
		set[value] = request
	}
	return func() {
		cache.mu.Lock()
		defer cache.mu.Unlock()
		for name := range clientState.GetSubscribedResources() {
			set, exists := cache.watches[name]
			if exists {
				delete(set, value)
			}
			if len(set) == 0 {
				delete(cache.watches, name)
			}
		}
	}, nil
}

func (cache *LinearCache) CreateDeltaWatch(request DeltaRequest, clientState SubscriptionState, value chan DeltaResponse) (func(), error) {
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
	response := cache.respondDelta(request, value, clientState)

	// if respondDelta returns nil this means that there is no change in any resource version
	// create a new watch accordingly
	if response == nil {
		watchID := cache.nextDeltaWatchID()
		if cache.log != nil {
			cache.log.Infof("[linear cache] open delta watch ID:%d for %s Resources:%v, system version %q", watchID,
				cache.typeURL, clientState.GetSubscribedResources(), cache.getVersion())
		}

		cache.deltaWatches[watchID] = DeltaResponseWatch{Request: request, Response: value, subscriptionState: clientState}

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

func (cache *LinearCache) Fetch(context.Context, FetchRequest) (Response, error) {
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
