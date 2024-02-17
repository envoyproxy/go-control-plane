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
	"encoding/hex"
	"errors"
	"fmt"
	"hash/fnv"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/log"
)

// cachedResource is used to track resources added by the user in the cache.
// It contains the resource itself and its associated version (currently in two different modes).
type cachedResource struct {
	types.Resource

	// cacheVersion is the version of the cache at the time of last update, used in sotw.
	cacheVersion string
	// stableVersion is the version of the resource itself (a hash of its content after deterministic marshaling).
	// It is lazy initialized and should be accessed through getStableVersion.
	stableVersion string
}

func newCachedResource(res types.Resource, cacheVersion string) *cachedResource {
	return &cachedResource{
		Resource:     res,
		cacheVersion: cacheVersion,
	}
}

func (c *cachedResource) getStableVersion() (string, error) {
	if c.stableVersion != "" {
		return c.stableVersion, nil
	}

	// TODO(valerian-roche): store serialized resource as part of the cachedResource
	// to reuse it when marshaling the responses instead of remarshaling and recomputing the version then.
	marshaledResource, err := MarshalResource(c.Resource)
	if err != nil {
		return "", err
	}
	c.stableVersion = HashResource(marshaledResource)
	return c.stableVersion, nil
}

func (c *cachedResource) getVersion(useStableVersion bool) (string, error) {
	if !useStableVersion {
		return c.cacheVersion, nil
	}

	return c.getStableVersion()
}

type watches struct {
	// sotw keeps track of current sotw watches, indexed per watch id.
	sotw map[uint64]ResponseWatch
	// delta keeps track of current delta watches, indexed per watch id.
	delta map[uint64]DeltaResponseWatch
}

func newWatches() watches {
	return watches{
		sotw:  make(map[uint64]ResponseWatch),
		delta: make(map[uint64]DeltaResponseWatch),
	}
}

func (w *watches) empty() bool {
	return len(w.sotw)+len(w.delta) == 0
}

// LinearCache supports collections of opaque resources. This cache has a
// single collection indexed by resource names and manages resource versions
// internally. It implements the cache interface for a single type URL and
// should be combined with other caches via type URL muxing. It can be used to
// supply EDS entries, for example, uniformly across a fleet of proxies.
type LinearCache struct {
	// typeURL provides the type of resources managed by the cache.
	// This information is used to reject requests watching another type, as well as to make
	// decisions based on resource type (e.g. whether sotw must return full-state).
	typeURL string

	// resources contains all resources currently set in the cache and associated versions.
	resources map[string]*cachedResource

	// resourceWatches keeps track of watches currently opened specifically tracking a resource.
	// It does not contain wildcard watches.
	// It can contain resources not present in resources.
	resourceWatches map[string]watches
	// wildcardWatches keeps track of all wildcard watches currently opened.
	wildcardWatches watches
	// currentWatchID is used to index new watches.
	currentWatchID uint64

	// version is the current version of the cache. It is incremented each time resources are updated.
	version uint64
	// versionPrefix is used to modify the version returned to clients, and can be used to uniquely identify
	// cache instances and avoid issues of version reuse.
	versionPrefix string

	// useStableVersionsInSotw switches to a new version model for sotw watches.
	// When activated, versions are stored in subscriptions using stable versions, and the response version
	// is an hash of the returned versions to allow watch resumptions when reconnecting to the cache with a
	// new subscription.
	useStableVersionsInSotw bool

	log log.Logger

	mu sync.RWMutex
}

var _ Cache = &LinearCache{}

// Options for modifying the behavior of the linear cache.
type LinearCacheOption func(*LinearCache)

// WithVersionPrefix sets a version prefix of the form "prefixN" in the version info.
// Version prefix can be used to distinguish replicated instances of the cache, in case
// a client re-connects to another instance.
// Deprecated: use WithSotwStableVersions instead to avoid issues when reconnecting to other instances
// while avoiding resending resources if unchanged.
func WithVersionPrefix(prefix string) LinearCacheOption {
	return func(cache *LinearCache) {
		cache.versionPrefix = prefix
	}
}

// WithInitialResources initializes the initial set of resources.
func WithInitialResources(resources map[string]types.Resource) LinearCacheOption {
	return func(cache *LinearCache) {
		for name, resource := range resources {
			cache.resources[name] = &cachedResource{
				Resource: resource,
			}
		}
	}
}

func WithLogger(log log.Logger) LinearCacheOption {
	return func(cache *LinearCache) {
		cache.log = log
	}
}

// WithSotwStableVersions changes the versions returned in sotw to encode the list of resources known
// in the subscription.
// The use of stable versions for sotw also deduplicates updates to clients if the cache updates are
// not changing the content of the resource.
// When used, the use of WithVersionPrefix is no longer needed to manage reconnection to other instances
// and should not be used.
func WithSotwStableVersions() LinearCacheOption {
	return func(cache *LinearCache) {
		cache.useStableVersionsInSotw = true
	}
}

// NewLinearCache creates a new cache. See the comments on the struct definition.
func NewLinearCache(typeURL string, opts ...LinearCacheOption) *LinearCache {
	out := &LinearCache{
		typeURL:         typeURL,
		resources:       make(map[string]*cachedResource),
		resourceWatches: make(map[string]watches),
		wildcardWatches: newWatches(),
		version:         0,
		currentWatchID:  0,
		log:             log.NewDefaultLogger(),
	}
	for _, opt := range opts {
		opt(out)
	}
	for name, resource := range out.resources {
		resource.cacheVersion = out.getVersion()
		out.resources[name] = resource
	}
	return out
}

// computeResourceChange compares the subscription known resources and the cache current state to compute the list of resources
// which have changed and should be notified to the user.
//
// The alwaysConsiderAllResources argument removes the consideration of the subscription known resources (e.g. if the version did not match),
// and return all known subscribed resources.
//
// The useStableVersion argument defines what version type to use for resources:
//   - if set to false versions are based on when resources were updated in the cache.
//   - if set to true versions are a stable property of the resource, with no regard to when it was added to the cache.
func (cache *LinearCache) computeResourceChange(sub Subscription, alwaysConsiderAllResources, useStableVersion bool) (updated, removed []string, err error) {
	var changedResources []string
	var removedResources []string

	knownVersions := sub.ReturnedResources()
	if alwaysConsiderAllResources {
		// The response will include all resources, with no regards of resources potentially already returned.
		knownVersions = make(map[string]string)
	}

	if sub.IsWildcard() {
		for resourceName, resource := range cache.resources {
			knownVersion, ok := knownVersions[resourceName]
			if !ok {
				// This resource is not yet known by the client (new resource added in the cache or newly subscribed).
				changedResources = append(changedResources, resourceName)
			} else {
				resourceVersion, err := resource.getVersion(useStableVersion)
				if err != nil {
					return nil, nil, fmt.Errorf("failed to compute version of %s: %w", resourceName, err)
				}
				if knownVersion != resourceVersion {
					// The client knows an outdated version.
					changedResources = append(changedResources, resourceName)
				}
			}
		}

		// Negative check to identify resources that have been removed in the cache.
		// Sotw does not support returning "deletions", but in the case of full state resources
		// a response must then be returned.
		for resourceName := range knownVersions {
			if _, ok := cache.resources[resourceName]; !ok {
				removedResources = append(removedResources, resourceName)
			}
		}
	} else {
		for resourceName := range sub.SubscribedResources() {
			res, exists := cache.resources[resourceName]
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
			} else {
				resourceVersion, err := res.getVersion(useStableVersion)
				if err != nil {
					return nil, nil, fmt.Errorf("failed to compute version of %s: %w", resourceName, err)
				}
				if knownVersion != resourceVersion {
					// The client knows an outdated version.
					changedResources = append(changedResources, resourceName)
				}
			}
		}

		for resourceName := range knownVersions {
			// If the subscription no longer watches a resource,
			// we mark it as unknown on the client side to ensure it will be resent to the client if subscribing again later on.
			if _, ok := sub.SubscribedResources()[resourceName]; !ok {
				removedResources = append(removedResources, resourceName)
			}
		}
	}

	return changedResources, removedResources, nil
}

func computeSotwStableVersion(versionMap map[string]string) string {
	// To enforce a stable hash we need to have an ordered vision of the map.
	keys := make([]string, 0, len(versionMap))
	for key := range versionMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	mapHasher := fnv.New64()

	buffer := make([]byte, 0, 8)
	itemHasher := fnv.New64()
	for _, key := range keys {
		buffer = buffer[:0]
		itemHasher.Reset()
		itemHasher.Write([]byte(key))
		mapHasher.Write(itemHasher.Sum(buffer))
		buffer = buffer[:0]
		itemHasher.Reset()
		itemHasher.Write([]byte(versionMap[key]))
		mapHasher.Write(itemHasher.Sum(buffer))
	}
	buffer = buffer[:0]
	return hex.EncodeToString(mapHasher.Sum(buffer))
}

func (cache *LinearCache) computeSotwResponse(watch ResponseWatch, alwaysConsiderAllResources bool) (*RawResponse, error) {
	changedResources, removedResources, err := cache.computeResourceChange(watch.subscription, alwaysConsiderAllResources, cache.useStableVersionsInSotw)
	if err != nil {
		return nil, err
	}
	if len(changedResources) == 0 && len(removedResources) == 0 && !alwaysConsiderAllResources {
		// Nothing changed.
		return nil, nil
	}

	// In sotw the list of resources to actually return depends on:
	//  - whether the type requires full-state in each reply (lds and cds).
	//  - whether the request is wildcard.
	// resourcesToReturn will include all the resource names to reply based on the changes detected.
	var resourcesToReturn []string

	switch {
	// For lds and cds, answers will always include all existing subscribed resources, with no regard to which resource was changed or removed.
	// For other types, the response only includes updated resources (sotw cannot notify for deletion).
	case !ResourceRequiresFullStateInSotw(cache.typeURL):
		if !alwaysConsiderAllResources && len(changedResources) == 0 {
			// If the request is not the initial one, and the type does not require full updates,
			// do not return if nothing is to be set.
			// For full-state resources an empty response does have a semantic meaning.
			return nil, nil
		}

		// changedResources is already filtered based on the subscription.
		resourcesToReturn = changedResources
	case watch.subscription.IsWildcard():
		// Include all resources for the type.
		resourcesToReturn = make([]string, 0, len(cache.resources))
		for resourceName := range cache.resources {
			resourcesToReturn = append(resourcesToReturn, resourceName)
		}
	default:
		// Include all resources matching the subscription, with no concern on whether it has been updated or not.
		requestedResources := watch.subscription.SubscribedResources()
		// The linear cache could be very large (e.g. containing all potential CLAs)
		// Therefore drives on the subscription requested resources.
		resourcesToReturn = make([]string, 0, len(requestedResources))
		for resourceName := range requestedResources {
			if _, ok := cache.resources[resourceName]; ok {
				resourcesToReturn = append(resourcesToReturn, resourceName)
			}
		}
	}

	// returnedVersions includes all resources currently known to the subscription and their version.
	returnedVersions := make(map[string]string, len(watch.subscription.ReturnedResources()))
	// Clone the current returned versions. The cache should not alter the subscription.
	for resourceName, version := range watch.subscription.ReturnedResources() {
		returnedVersions[resourceName] = version
	}

	resources := make([]types.ResourceWithTTL, 0, len(resourcesToReturn))
	for _, resourceName := range resourcesToReturn {
		cachedResource := cache.resources[resourceName]
		resources = append(resources, types.ResourceWithTTL{Resource: cachedResource.Resource})
		version, err := cachedResource.getVersion(cache.useStableVersionsInSotw)
		if err != nil {
			return nil, fmt.Errorf("failed to compute version of %s: %w", resourceName, err)
		}
		returnedVersions[resourceName] = version
	}
	// Cleanup resources no longer existing in the cache or no longer subscribed.
	// In sotw we cannot return those if not full state,
	// but this ensures we detect unsubscription then resubscription.
	for _, resourceName := range removedResources {
		delete(returnedVersions, resourceName)
	}

	responseVersion := cache.getVersion()
	if cache.useStableVersionsInSotw {
		responseVersion = cache.versionPrefix + computeSotwStableVersion(returnedVersions)
	}

	return &RawResponse{
		Request:           watch.Request,
		Resources:         resources,
		ReturnedResources: returnedVersions,
		Version:           responseVersion,
		Ctx:               context.Background(),
	}, nil
}

func (cache *LinearCache) computeDeltaResponse(watch DeltaResponseWatch) (*RawDeltaResponse, error) {
	changedResources, removedResources, err := cache.computeResourceChange(watch.subscription, false, true)
	if err != nil {
		return nil, err
	}

	// On first request on a wildcard subscription, envoy does expect a response to come in to
	// conclude initialization.
	isFirstWildcardRequest := watch.subscription.IsWildcard() && watch.Request.GetResponseNonce() == ""
	if len(changedResources) == 0 && len(removedResources) == 0 && !isFirstWildcardRequest {
		// Nothing changed.
		return nil, nil
	}

	returnedVersions := make(map[string]string, len(watch.subscription.ReturnedResources()))
	// Clone the current returned versions. The cache should not alter the subscription
	for resourceName, version := range watch.subscription.ReturnedResources() {
		returnedVersions[resourceName] = version
	}

	cacheVersion := cache.getVersion()
	resources := make([]types.Resource, 0, len(changedResources))
	for _, resourceName := range changedResources {
		resource := cache.resources[resourceName]
		resources = append(resources, resource.Resource)
		version, err := resource.getStableVersion()
		if err != nil {
			return nil, fmt.Errorf("failed to compute stable version of %s: %w", resourceName, err)
		}
		returnedVersions[resourceName] = version
	}
	// Cleanup resources no longer existing in the cache or no longer subscribed.
	for _, resourceName := range removedResources {
		delete(returnedVersions, resourceName)
	}

	return &RawDeltaResponse{
		DeltaRequest:      watch.Request,
		Resources:         resources,
		RemovedResources:  removedResources,
		NextVersionMap:    returnedVersions,
		SystemVersionInfo: cacheVersion,
		Ctx:               context.Background(),
	}, nil
}

func (cache *LinearCache) notifyAll(modified []string) error {
	// Gather the list of watches impacted by the modified resources.
	sotwWatches := make(map[uint64]ResponseWatch)
	deltaWatches := make(map[uint64]DeltaResponseWatch)
	for _, name := range modified {
		for watchID, watch := range cache.resourceWatches[name].sotw {
			sotwWatches[watchID] = watch
		}
		for watchID, watch := range cache.resourceWatches[name].delta {
			deltaWatches[watchID] = watch
		}
	}

	// sotw watches
	for watchID, watch := range sotwWatches {
		response, err := cache.computeSotwResponse(watch, false)
		if err != nil {
			return err
		}

		if response != nil {
			watch.Response <- response
			cache.removeWatch(watchID, watch.subscription)
		} else {
			cache.log.Infof("[Linear cache] Watch %d detected as triggered but no change was found", watchID)
		}
	}

	for watchID, watch := range cache.wildcardWatches.sotw {
		response, err := cache.computeSotwResponse(watch, false)
		if err != nil {
			return err
		}

		if response != nil {
			watch.Response <- response
			delete(cache.wildcardWatches.sotw, watchID)
		} else {
			cache.log.Infof("[Linear cache] Wildcard watch %d detected as triggered but no change was found", watchID)
		}
	}

	// delta watches
	for watchID, watch := range deltaWatches {
		response, err := cache.computeDeltaResponse(watch)
		if err != nil {
			return err
		}

		if response != nil {
			watch.Response <- response
			cache.removeDeltaWatch(watchID, watch.subscription)
		} else {
			cache.log.Infof("[Linear cache] Delta watch %d detected as triggered but no change was found", watchID)
		}
	}

	for watchID, watch := range cache.wildcardWatches.delta {
		response, err := cache.computeDeltaResponse(watch)
		if err != nil {
			return err
		}

		if response != nil {
			watch.Response <- response
			delete(cache.wildcardWatches.delta, watchID)
		} else {
			cache.log.Infof("[Linear cache] Wildcard delta watch %d detected as triggered but no change was found", watchID)
		}
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
	cache.resources[name] = newCachedResource(res, cache.getVersion())

	return cache.notifyAll([]string{name})
}

// DeleteResource removes a resource in the collection.
func (cache *LinearCache) DeleteResource(name string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.version++
	delete(cache.resources, name)

	return cache.notifyAll([]string{name})
}

// UpdateResources updates/deletes a list of resources in the cache.
// Calling UpdateResources instead of iterating on UpdateResource and DeleteResource
// is significantly more efficient when using delta or wildcard watches.
func (cache *LinearCache) UpdateResources(toUpdate map[string]types.Resource, toDelete []string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.version++
	version := cache.getVersion()
	modified := make([]string, 0, len(toUpdate)+len(toDelete))
	for name, resource := range toUpdate {
		cache.resources[name] = newCachedResource(resource, version)
		modified = append(modified, name)
	}
	for _, name := range toDelete {
		delete(cache.resources, name)
		modified = append(modified, name)
	}

	return cache.notifyAll(modified)
}

// SetResources replaces current resources with a new set of resources.
// If only some resources are to be updated, UpdateResources is more efficient.
func (cache *LinearCache) SetResources(resources map[string]types.Resource) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.version++
	version := cache.getVersion()

	modified := make([]string, 0, len(resources))
	// Collect deleted resource names.
	for name := range cache.resources {
		if _, found := resources[name]; !found {
			delete(cache.resources, name)
			modified = append(modified, name)
		}
	}

	// We assume all resources passed to SetResources are changed.
	// In delta and if stable versions are used for sotw, identical resources will not trigger watches.
	// In sotw without stable versions used, all those resources will trigger watches, even if identical.
	for name, resource := range resources {
		cache.resources[name] = newCachedResource(resource, version)
		modified = append(modified, name)
	}

	if err := cache.notifyAll(modified); err != nil {
		cache.log.Errorf("Failed to notify watches: %s", err.Error())
	}
}

// GetResources returns current resources stored in the cache
func (cache *LinearCache) GetResources() map[string]types.Resource {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	// create a copy of our internal storage to avoid data races
	// involving mutations of our backing map
	resources := make(map[string]types.Resource, len(cache.resources))
	for k, v := range cache.resources {
		resources[k] = v.Resource
	}
	return resources
}

// The implementations of sotw and delta watches handling is nearly identical. The main distinctions are:
//   - handling of version in sotw when the request is the first of a subscription. Delta has a proper handling based on the request providing known versions.
//   - building the initial resource versions in delta if they've not been computed yet.
//   - computeSotwResponse and computeDeltaResponse has slightly different implementations due to sotw requirements to return full state for certain resources only.
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

	response, err := cache.computeSotwResponse(watch, ignoreCurrentSubscriptionResources)
	if err != nil {
		return nil, fmt.Errorf("failed to compute the watch respnse: %w", err)
	}
	shouldReply := false
	if response != nil {
		// If the request
		//  - is the first
		//  - provides a non-empty version, matching the version prefix
		// and the cache uses stable versions, if the generated versions are the same as the previous one, we do not return the response.
		// This avoids resending all data if the new subscription is just a resumption of the previous one.
		if cache.useStableVersionsInSotw && request.GetResponseNonce() == "" && !ignoreCurrentSubscriptionResources {
			shouldReply = request.GetVersionInfo() != response.Version

			// We confirmed the content of the known resources, store them in the watch we create.
			subscription := newWatchSubscription(sub)
			subscription.returnedResources = response.ReturnedResources
			watch.subscription = subscription
			sub = subscription
		} else {
			shouldReply = true
		}
	}

	if shouldReply {
		cache.log.Debugf("[linear cache] replying to the watch with resources %v (subscription values %v, known %v)", response.GetReturnedResources(), sub.SubscribedResources(), sub.ReturnedResources())
		watch.Response <- response
		return func() {}, nil
	}

	watchID := cache.nextWatchID()
	// Create open watches since versions are up to date.
	if sub.IsWildcard() {
		cache.log.Infof("[linear cache] open watch %d for %s all resources, known versions %v, system version %q", watchID, cache.typeURL, sub.ReturnedResources(), cache.getVersion())
		cache.wildcardWatches.sotw[watchID] = watch
		return func() {
			cache.mu.Lock()
			defer cache.mu.Unlock()
			delete(cache.wildcardWatches.sotw, watchID)
		}, nil
	}

	cache.log.Infof("[linear cache] open watch %d for %s resources %v, known versions %v, system version %q", watchID, cache.typeURL, sub.SubscribedResources(), sub.ReturnedResources(), cache.getVersion())
	for name := range sub.SubscribedResources() {
		watches, exists := cache.resourceWatches[name]
		if !exists {
			watches = newWatches()
			cache.resourceWatches[name] = watches
		}
		watches.sotw[watchID] = watch
	}
	return func() {
		cache.mu.Lock()
		defer cache.mu.Unlock()
		cache.removeWatch(watchID, watch.subscription)
	}, nil
}

// Must be called under lock
func (cache *LinearCache) removeWatch(watchID uint64, sub Subscription) {
	// Make sure we clean the watch for ALL resources it might be associated with,
	// as the channel will no longer be listened to
	for resource := range sub.SubscribedResources() {
		resourceWatches := cache.resourceWatches[resource]
		delete(resourceWatches.sotw, watchID)
		if resourceWatches.empty() {
			delete(cache.resourceWatches, resource)
		}
	}
}

func (cache *LinearCache) CreateDeltaWatch(request *DeltaRequest, sub Subscription, value chan DeltaResponse) (func(), error) {
	if request.GetTypeUrl() != cache.typeURL {
		return nil, fmt.Errorf("request type %s does not match cache type %s", request.GetTypeUrl(), cache.typeURL)
	}

	watch := DeltaResponseWatch{Request: request, Response: value, subscription: sub}

	cache.mu.Lock()
	defer cache.mu.Unlock()

	response, err := cache.computeDeltaResponse(watch)
	if err != nil {
		return nil, fmt.Errorf("failed to compute the watch respnse: %w", err)
	}

	if response != nil {
		cache.log.Debugf("[linear cache] replying to the delta watch (subscription values %v, known %v)", sub.SubscribedResources(), sub.ReturnedResources())
		watch.Response <- response
		return nil, nil
	}

	watchID := cache.nextWatchID()
	// Create open watches since versions are up to date.
	if sub.IsWildcard() {
		cache.log.Infof("[linear cache] open delta watch %d for all %s resources, system version %q", watchID, cache.typeURL, cache.getVersion())
		cache.wildcardWatches.delta[watchID] = watch
		return func() {
			cache.mu.Lock()
			defer cache.mu.Unlock()
			delete(cache.wildcardWatches.delta, watchID)
		}, nil
	}

	cache.log.Infof("[linear cache] open delta watch %d for %s resources %v, system version %q", watchID, cache.typeURL, sub.SubscribedResources(), cache.getVersion())
	for name := range sub.SubscribedResources() {
		watches, exists := cache.resourceWatches[name]
		if !exists {
			watches = newWatches()
			cache.resourceWatches[name] = watches
		}
		watches.delta[watchID] = watch
	}
	return func() {
		cache.mu.Lock()
		defer cache.mu.Unlock()
		cache.removeDeltaWatch(watchID, watch.subscription)
	}, nil
}

func (cache *LinearCache) getVersion() string {
	return cache.versionPrefix + strconv.FormatUint(cache.version, 10)
}

// cancellation function for cleaning stale watches
func (cache *LinearCache) removeDeltaWatch(watchID uint64, sub Subscription) {
	// Make sure we clean the watch for ALL resources it might be associated with,
	// as the channel will no longer be listened to
	for resource := range sub.SubscribedResources() {
		resourceWatches := cache.resourceWatches[resource]
		delete(resourceWatches.delta, watchID)
		if resourceWatches.empty() {
			delete(cache.resourceWatches, resource)
		}
	}
}

func (cache *LinearCache) nextWatchID() uint64 {
	cache.currentWatchID++
	if cache.currentWatchID == 0 {
		panic("watch id count overflow")
	}
	return cache.currentWatchID
}

func (cache *LinearCache) Fetch(context.Context, *Request) (Response, error) {
	return nil, errors.New("not implemented")
}

// NumResources returns the number of resources currently in the cache.
// As GetResources is building a clone it is expensive to get metrics otherwise.
func (cache *LinearCache) NumResources() int {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	return len(cache.resources)
}

// NumWatches returns the number of active sotw watches for a resource name.
func (cache *LinearCache) NumWatches(name string) int {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	return len(cache.resourceWatches[name].sotw) + len(cache.wildcardWatches.sotw)
}

// NumDeltaWatchesForResource returns the number of active delta watches for a resource name.
func (cache *LinearCache) NumDeltaWatchesForResource(name string) int {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	return len(cache.resourceWatches[name].delta) + len(cache.wildcardWatches.delta)
}

// NumDeltaWatches returns the total number of active delta watches.
// Warning: it is quite inefficient, and NumDeltaWatchesForResource should be preferred.
func (cache *LinearCache) NumDeltaWatches() int {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	uniqueWatches := map[uint64]struct{}{}
	for _, watches := range cache.resourceWatches {
		for id := range watches.delta {
			uniqueWatches[id] = struct{}{}
		}
	}
	return len(uniqueWatches) + len(cache.wildcardWatches.delta)
}
