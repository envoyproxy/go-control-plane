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
	if info, ok := cache.status[node]; ok {
		info.mu.Lock()
		for id, watch := range info.deltaWatches {
			// Get the version of the current resources per type url
			t := watch.Request.GetTypeUrl()
			subscribed := snapshot.GetSubscribedResources(watch.Request.GetResourceNamesSubscribe(), watch.Request.GetTypeUrl())
			// This can be nil, so we don't need to worry about passing nil to our deltaResponse creators
			unsubscribed := watch.Request.GetResourceNamesUnsubscribe()
			version := snapshot.GetVersion(t)

			if version != info.deltaState.state[t].Version {
				// We want to perform some pre-processing before we create the diff
				if len(subscribed) == 0 && len(unsubscribed) == 0 {
					cache.log.Debugf("wildcard request")

					// Set the new state and since this is wildcard
					// we can just populat the state with all since that is what the stream has requested
					info.deltaState.state[t] = Resources{
						Version: version,
						Items:   snapshot.GetResources(t),
					}

					// Respond immediately
					cache.respondDelta(
						watch.Request,
						watch.Response,
						snapshot.GetVersionMap(),
						snapshot.GetResources(t),
						unsubscribed,
					)

					// Clean up and since we've responded we can continue going through the rest of the watches
					delete(info.deltaWatches, id)
					continue
				}

				if len(unsubscribed) > 0 {
					// we need to remove the previously subscribed resources from the state so we no longer send updates
					if cache.log != nil {
						cache.log.Debugf("node: %s, recieved items to unsubscribe from: %v", node, unsubscribed)
					}

					// Mutates deltaState and will remove the items from the map
					cache.unsubscribe(unsubscribed, info.deltaState.state)
				}

				// Now calculate the diff and see what has changed in the state
				diff := cache.checkState(subscribed, info.deltaState.state[t].Items)
				if len(diff) > 0 {
					if cache.log != nil {
						cache.log.Debugf("node: %s, found new items to subscribe too: %v ", node, diff)
					}

					// Add our new subscription items to our state to watch that we've found
					r := Resources{
						Version: version,
					}
					for key, value := range diff {
						for rKey := range info.deltaState.state[t].Items {
							// Handle the case when a new item could be added to the state and also if we need to overwrite a previous resource
							if key == rKey {
								info.deltaState.state[t].Items[key] = value
							} else if _, found := info.deltaState.state[t].Items[key]; !found {
								info.deltaState.state[t].Items[key] = value
							}
						}
					}
					r.Items = info.deltaState.state[t].Items

					info.deltaState.state[t] = r
				}

				if cache.log != nil {
					// We only want to show the specific resources we're sending back from the diff
					cache.log.Debugf("delta respond open watch ID:%d Resources:%+v with new version %q", id, diff, version)
				}

				// Respond to our delta stream with the new diff after processing has completed
				cache.respondDelta(
					watch.Request,
					watch.Response,
					snapshot.GetVersionMap(),
					diff, // We want to only send the updated resources here
					unsubscribed,
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
	diff := make(map[string]types.Resource, len(deltaState))

	// Check our diff map to see what has changed
	// Even is an underlying resource has changed we need to update the diff
	for key, value := range resources {
		if resource, found := deltaState[key]; !found || resource != value {
			if cache.log != nil {
				cache.log.Debugf("Detected change in deltaState: %s -> %s\n", key, GetResourceName(resource))
			}
			diff[key] = value
		}
	}

	return diff
}

// Unscrubscribe will remove resources from the tracked list in the management server when received from a client
func (cache *snapshotCache) unsubscribe(resources []string, deltaState map[string]Resources) {
	// here we need to search and remove from the current subscribed list in the snapshot
	for _, items := range deltaState {
		for i := 0; i < len(resources); i++ {
			if cache.log != nil {
				cache.log.Debugf("unsubscribing from resource: %s", resources[i])
			}
			delete(items.Items, resources[i])
		}
	}
}

// CreateDeltaWatch returns a watch for a delta xDS request.
// Requester now sets version info when creating new delta watches
<<<<<<< HEAD
func (cache *snapshotCache) CreateDeltaWatch(request DeltaRequest, requestVersion string) (chan DeltaResponse, func()) {
=======
func (cache *snapshotCache) CreateDeltaWatch(request *DeltaRequest) (chan DeltaResponse, func()) {
>>>>>>> progress save... restructuring version logic in cache
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

	// update last watch request times
	info.mu.Lock()
	info.lastDeltaWatchRequestTime = time.Now()
	info.mu.Unlock()

	// allocate capacity 1 to allow one-time non-blocking use
	value := make(chan DeltaResponse, 1)

	// find the current cache snapshot for the provided node
	snapshot, exists := cache.snapshots[nodeID]

	// I don't think we use the snapshot version here it is pretty much irrelevant for delta xDS
	snapshotVersion := snapshot.GetVersion(t)
	vMap := snapshot.GetVersionMap()

	// This logic is going to need to change to compare individual resource versions,
	// if we detect a change in resource version from the previous snapshot then we should create a new watch accordingly
	// if the requested version is up-to-date or missing a response, leave an open watch
	if !exists || snapshotVersion != info.GetDeltaStateSystemVersion() {
		watchID := cache.nextDeltaWatchID()
		if cache.log != nil {
			cache.log.Infof("open delta watch ID:%d for %s Resources:%v from nodeID: %q, system version %q", watchID,
				t, aliases, nodeID, snapshotVersion)
		}

		info.mu.Lock()
		info.deltaWatches[watchID] = DeltaResponseWatch{Request: request, Response: value}
		// Set our initial state when a watch is created
		if snapshotVersion != "" {
			// Create our delta state versions if the snapshot exists
			for _, resource := range snapshot.GetSubscribedResources(aliases, t) {
				v, err := HashResource(resource)
				if err != nil {
					panic(err)
				}

				cache.log.Debugf("resource version created: %s", v)
			}
			info.deltaState.state[t] = Resources{}
		} else {
			info.deltaState.state[t] = Resources{
				Version: "",
				Items:   snapshot.GetSubscribedResources(aliases, t),
			}
		}
		info.mu.Unlock()

		return value, cache.cancelDeltaWatch(nodeID, watchID)
	}

	// otherwise, the watch may be responded to immediately with the subscribed resources
	// we don't want to ask for all the resources by type here
	// we do want to respond with the full resource version map though
	info.mu.RLock()
	cache.respondDelta(request, value, vMap, info.deltaState.state[t].Items, nil)
	info.mu.RUnlock()

	return value, nil
}

func (cache *snapshotCache) nextDeltaWatchID() int64 {
	return atomic.AddInt64(&cache.deltaWatchCount, 1)
}

// cancellation function for cleaning stale watches
func (cache *snapshotCache) cancelDeltaWatch(nodeID string, watchID int64) func() {
	return func() {
		// uses the cache mutex
		cache.mu.Lock()
		defer cache.mu.Unlock()
		if info, ok := cache.status[nodeID]; ok {
			info.mu.Lock()
			delete(info.deltaWatches, watchID)
			info.mu.Unlock()
		}
	}
}

func (cache *snapshotCache) respondDelta(request *DeltaRequest, value chan DeltaResponse, versionMap map[string][]DeltaVersionInfo, resources map[string]types.Resource, unsubscribed []string) {
	if cache.log != nil {
		cache.log.Debugf("node: %s sending delta response %s with resource versions: %q",
			request.GetNode().GetId(), request.TypeUrl)
	}

	value <- createDeltaResponse(request, versionMap, resources, unsubscribed)
}

func createDeltaResponse(request *DeltaRequest, versionMap map[string][]DeltaVersionInfo, resources map[string]types.Resource, unsubscribed []string) DeltaResponse {
	filtered := make([]types.Resource, 0, len(resources))

	// Reply only with the requested resources. Envoy may ask each resource
	// individually in a separate stream. It is ok to reply with the same version
	// on separate streams since requests do not share their response versions.

	// This logic is probably broken so we'll revisit
	// if len(request.ResourceNamesSubscribe) != 0 {
	// 	set := nameSet(request.ResourceNamesSubscribe)
	// 	for name, resource := range resources {
	// 		if set[name] {
	// 			filtered = append(filtered, resource)
	// 		}
	// 	}
	// } else {
	for _, resource := range resources {
		filtered = append(filtered, resource)
	}
	// }

	// send through our version map
	return &RawDeltaResponse{
		DeltaRequest:     request,
		Resources:        filtered,
		VersionMap:       versionMap,
		RemovedResources: unsubscribed,
	}
}
