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
	"errors"
	"sync/atomic"
	"time"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
)

// CreateDeltaWatch returns a watch for a delta xDS request.
func (cache *snapshotCache) CreateDeltaWatch(request *DeltaRequest, sv StreamVersion) (chan DeltaResponse, func()) {
	nodeID := cache.hash.ID(request.Node)
	t := request.GetTypeUrl()
	aliases := request.GetResourceNamesSubscribe()
	if sv == nil {
		panic(errors.New("StreamVersion cannot be nil when creating a delta watch"))
	}

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
	snapshotVersion := snapshot.GetVersion(t)
	vMap := snapshot.GetVersionMap()

	// Compare our requested versions with the existing snapshot
	var versionChange bool
	for alias, version := range vMap[t] {
		for a, v := range sv.GetVersionMap() {
			if a == alias && v == version {
				versionChange = true
			}
		}
	}

	// if we detect a change in resource version from the previous snapshot then we should create a new watch accordingly
	// if the requested version is up-to-date or missing a response, leave an open watch
	if !exists || versionChange {
		watchID := cache.nextDeltaWatchID()
		if cache.log != nil {
			cache.log.Infof("open delta watch ID:%d for %s Resources:%v from nodeID: %q, system version %q", watchID,
				t, aliases, nodeID, snapshotVersion)
		}

		info.mu.Lock()
		info.deltaWatches[watchID] = DeltaResponseWatch{Request: request, Response: value, VersionMap: sv.GetVersionMap()}
		info.mu.Unlock()

		return value, cache.cancelDeltaWatch(nodeID, watchID)
	}

	// otherwise, the watch may be responded to immediately with the subscribed resources
	// we don't want to ask for all the resources by type here
	// we do want to respond with the full resource version map though
	cache.respondDelta(request, value, vMap[t], snapshot.GetResources(t), nil)
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

func (cache *snapshotCache) respondDelta(request *DeltaRequest, value chan DeltaResponse, versionMap map[string]DeltaVersionInfo, resources map[string]types.Resource, unsubscribed []string) {
	if cache.ads && len(request.ResourceNamesSubscribe) != 0 {
		if err := superset(nameSet(request.ResourceNamesSubscribe), resources); err != nil {
			if cache.log != nil {
				cache.log.Debugf("Delta ADS mode: not responding to request %v", err)
			}
			return
		}
	}

	if cache.log != nil {
		cache.log.Debugf("node: %s sending delta response %s with resource versions: %v",
			request.GetNode().GetId(), request.TypeUrl, versionMap)
	}

	value <- createDeltaResponse(request, versionMap, resources, unsubscribed)
}

func createDeltaResponse(request *DeltaRequest, versionMap map[string]DeltaVersionInfo, resources map[string]types.Resource, unsubscribed []string) DeltaResponse {
	filtered := make([]types.Resource, 0, len(resources))

	// Reply only with the requested resources. Envoy may ask each resource
	// individually in a separate stream. It is ok to reply with the same version
	// on separate streams since requests do not share their response versions.
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

	// send through our version map
	return &RawDeltaResponse{
		DeltaRequest:     request,
		Resources:        filtered,
		VersionMap:       versionMap,
		RemovedResources: unsubscribed,
	}
}
