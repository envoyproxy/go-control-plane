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
	"sync/atomic"
	"time"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
)

func (cache *snapshotCache) SetSnapshotDelta(node string, snapshot Snapshot) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// update the existing entry
	cache.snapshots[node] = snapshot

	// trigger existing watches for which version changed
	// TODO:
	// Note the discarding of an old watch in regular SOTW snapshot set when a new version is found
	// In delta there's no concept of versioning against the request since the version_info field is no longer applicable
	// IDEA:
	// My current belief would be that this is where the diff needs to happen, so if the state of the node does not match up to what we have a watch for
	// then we need to discard the watch and it's resource version
	// I'm not sure how this will get done or if it's correct but that idea just popped into my head
	if info, ok := cache.status[node]; ok {
		info.mu.Lock()
		for id, watch := range info.deltaWatches {
			// Get the version of the current resources per type url
			t := watch.Request.GetTypeUrl()
			subscribed := snapshot.GetSubscribedResources(watch.Request.GetResourceNamesSubscribe(), watch.Request.GetTypeUrl())
			version := snapshot.GetVersion(t)

			// Handle the case of an initial delta request and having no previous state
			// We can assume we just want to set the state as the initially requested resources
			if info.deltaState[t].Version == "" && len(info.deltaState[t].Items) == 0 {
				info.deltaState[t] = Resources{
					Version: version,
					Items:   subscribed,
				}

				// handle wildcard on the initial request
				// if this case is met, just subscribe to all clusters and listeners in the snapshot
				if len(subscribed) == 0 {
					cache.log.Debugf("setting wildcard")
					// Maybe set the resources for all the types here???
					for i := 0; i < int(types.UnknownType); i++ {
						tURL := GetResponseTypeURL(types.ResponseType(i))
						info.deltaState[tURL] = Resources{
							Version: version,
							Items:   snapshot.GetResources(tURL),
						}
					}

					// info.deltaState[t] = Resources{
					// 	Version: version,
					// 	Items:   snapshot.GetResources(t),
					// }
				}

				if cache.log != nil {
					if subscribed := watch.Request.GetResourceNamesSubscribe(); len(subscribed) != 0 {
						cache.log.Debugf("subscribing to resources: %+v", subscribed)
					}
					cache.log.Infof("initial snapshot set - respond to open watch ID:%d Resources:%+v", id, info.deltaState[t])
				}

				// Send out the response right away since we have nothing else to do
				cache.respondDelta(
					watch.Request,
					watch.Response,
					info.deltaState[t].Items,
					version,
				)

				// discard the old watch
				delete(info.deltaWatches, id)
			} else if version != info.deltaState[t].Version {
				// Assume we've received a new resource and we want to send new resources and cancel old watches
				diff := cache.checkState(subscribed, info.deltaState[t].Items)
				if len(diff) > 0 {
					if cache.log != nil {
						cache.log.Debugf("node: %s, found new items to subscribe too: %v ", watch.Request.GetNode().GetId(), diff)
					}

					// Add our new subscription items to our state to watch that we've found
					r := Resources{
						Version: version,
					}
					for key, value := range diff {
						for rKey := range info.deltaState[t].Items {
							// Handle the case when a new item could be added to the state and also if we need to overwrite a previous resource
							if key == rKey {
								info.deltaState[t].Items[key] = value
							} else if _, found := info.deltaState[t].Items[key]; !found {
								info.deltaState[t].Items[key] = value
							}
						}
					}
					r.Items = info.deltaState[t].Items

					info.deltaState[t] = r
				}

				if cache.log != nil {
					// We only want to show the specific resources we're sending back from the diff
					cache.log.Debugf("delta respond open watch ID:%d Resources:%+v with new version %q", id, diff, version)
				}

				// Respond to our delta stream with the subcribed resources
				cache.respondDelta(
					watch.Request,
					watch.Response,
					diff, // We want to only send the updated resources here
					version,
				)

				// discard the old watch
				delete(info.deltaWatches, id)
			}
		}
		info.mu.Unlock()
	}

	return nil
}

// difference returns the elements in `a` that aren't in `b`.
// TODO: SLOW this will need to be revisited
func (cache *snapshotCache) checkState(resources, deltaState map[string]types.Resource) map[string]types.Resource {
	mb := make(map[string]types.Resource, len(deltaState))
	diff := make(map[string]types.Resource, 0)

	for key, value := range deltaState {
		mb[key] = value
	}

	// Check our diff map to see what has changed
	// Even is an underlying resource has changed we need to update the diff
	for key, value := range resources {
		if _, found := mb[key]; !found {
			cache.log.Debugf("found new key: %s", key)
			diff[key] = value
		} else if resource, found := mb[key]; found && (resource != value) {
			cache.log.Debugf("found updated resource from existing key %s, modifying resource map", key)
			diff[key] = value
		}
	}

	return diff
}

// CreateDeltaWatch returns a watch for a delta xDS request.
// Requester now sets version info when creating new delta watches
func (cache *snapshotCache) CreateDeltaWatch(request DeltaRequest, requestVersion string) (chan DeltaResponse, func()) {
	nodeID := cache.hash.ID(request.Node)
	t := request.GetTypeUrl()
	aliases := request.GetResourceNamesSubscribe()

	cache.mu.Lock()
	defer cache.mu.Unlock()

	info, ok := cache.status[nodeID]
	if !ok {
		info = newStatusInfo(request.Node)
		cache.status[nodeID] = info
	}

	// update last watch request time
	info.mu.Lock()
	info.lastDeltaWatchRequestTime = time.Now()
	info.mu.Unlock()

	// allocate capacity 1 to allow one-time non-blocking use
	value := make(chan DeltaResponse, 1)

	// find the current cache snapshot for the provided node
	snapshot, exists := cache.snapshots[nodeID]
	version := snapshot.GetVersion(t)

	// if the requested version is up-to-date or missing a response, leave an open watch
	if !exists || version == requestVersion {
		watchID := cache.nextDeltaWatchID()
		if cache.log != nil {
			cache.log.Infof("open delta watch ID:%d for %s Resources:%v from nodeID: %q, version %q", watchID,
				t, aliases, nodeID, requestVersion)
		}

		info.mu.Lock()
		info.deltaWatches[watchID] = DeltaResponseWatch{Request: request, Response: value}
		// Set our initial state when a watch is created
		info.deltaState[t] = Resources{
			Version: requestVersion,
			Items:   snapshot.GetSubscribedResources(aliases, t),
		}
		info.mu.Unlock()

		return value, cache.cancelWatch(nodeID, watchID)
	}

	// otherwise, the watch may be responded to immediately with the subscribed resources
	// we don't want to ask for all the resources by type here
	cache.respondDelta(
		request,
		value,
		info.deltaState[t].Items,
		info.deltaState[t].Version,
	)

	return value, nil
}

func (cache *snapshotCache) nextDeltaWatchID() int64 {
	return atomic.AddInt64(&cache.deltaWatchCount, 1)
}

func (cache *snapshotCache) respondDelta(request DeltaRequest, value chan DeltaResponse, resources map[string]types.Resource, version string) {
	if cache.log != nil {
		cache.log.Debugf("node: %s sending delta response %s with version %q",
			request.GetNode().GetId(), request.TypeUrl, version)
	}

	value <- createDeltaResponse(request, resources, version)
}

func createDeltaResponse(request DeltaRequest, resources map[string]types.Resource, version string) DeltaResponse {
	filtered := make([]types.Resource, 0, len(resources))

	// Reply only with the requested resources. Envoy may ask each resource
	// individually in a separate stream. It is ok to reply with the same version
	// on separate streams since requests do not share their response versions.

	// This logic is probably broken so we'll revisit
	if len(request.ResourceNamesSubscribe) != 0 {
		set := nameSet(request.ResourceNamesSubscribe)
		for name, resource := range resources {
			if set[name] {
				filtered = append(filtered, resource)
			}
		}
	} else {
		for _, resource := range resources {
			filtered = append(filtered, resource)
		}
	}

	return DeltaResponse{
		DeltaRequest:  request,
		Resources:     filtered,
		SystemVersion: version,
	}
}
