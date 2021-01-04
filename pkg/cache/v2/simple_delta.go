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
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v2"
)

// CreateDeltaWatch returns a watch for a delta xDS request.
func (cache *snapshotCache) CreateDeltaWatch(request *DeltaRequest, st *stream.StreamState) (chan DeltaResponse, func()) {
	nodeID := cache.hash.ID(request.Node)
	t := request.GetTypeUrl()

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

	value := make(chan DeltaResponse, 1)

	// find the current cache snapshot for the provided node
	snapshot, exists := cache.snapshots[nodeID]

	// if respondDelta returns nil this means that there is no change in any resource version from the previous snapshot
	// create a new watch accordingly
	if !exists || cache.respondDelta(request, value, st, snapshot.GetResources(t)) == nil {
		watchID := cache.nextDeltaWatchID()
		if cache.log != nil {
			cache.log.Infof("open delta watch ID:%d for %s Resources:%v from nodeID: %q, system version %q", watchID,
				t, st.ResourceVersions, nodeID, snapshot.GetVersion(t))
		}

		info.mu.Lock()
		info.deltaWatches[watchID] = DeltaResponseWatch{Request: request, Response: value, StreamState: st}
		info.mu.Unlock()

		return value, cache.cancelDeltaWatch(nodeID, watchID)
	}

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

func (cache *snapshotCache) respondDelta(request *DeltaRequest, value chan DeltaResponse, st *stream.StreamState, resources map[string]types.Resource) *RawDeltaResponse {
	resp, err := createDeltaResponse(request, st, resources)
	if err != nil {
		if cache.log != nil {
			cache.log.Errorf("Error creating delta response: %v", err)
		}
		return nil
	}
	// One send response if there were some actual updates
	if len(resp.Resources) > 0 || len(resp.RemovedResources) > 0 {
		if cache.log != nil {
			cache.log.Debugf("node: %s, sending delta response:\n---> old Version Map: %v\n---> new resources: %v\n---> new Version Map: %v\n---> removed resources %v\n---> is wildcard: %t",
				request.GetNode().GetId(), st.ResourceVersions, resp.Resources, resp.VersionMap, resp.RemovedResources, st.IsWildcard)
		}
		value <- resp
		return resp
	}
	return nil
}

func createDeltaResponse(request *DeltaRequest, st *stream.StreamState, resources map[string]types.Resource) (*RawDeltaResponse, error) {
	newVersionMap := make(map[string]string)
	filtered := make([]types.Resource, 0)
	toRemove := make([]string, 0)
	if st.IsWildcard {
		for resourceName, resource := range resources {
			newVersion, err := HashResource(resource)
			if err != nil {
				return nil, err
			}
			newVersionMap[resourceName] = newVersion
			oldVersion, found := st.ResourceVersions[resourceName]

			if !found || oldVersion != newVersion {
				filtered = append(filtered, resource)
			}
		}
	} else {
		// Reply only with the requested resources. Envoy may ask each resource
		// individually in a separate stream. It is ok to reply with the same version
		// on separate streams since requests do not share their response states.
		for resourceName, oldVersion := range st.ResourceVersions {
			if r, ok := resources[resourceName]; ok {
				newVersion, err := HashResource(r)
				if err != nil {
					return nil, err
				}
				if oldVersion != newVersion {
					filtered = append(filtered, r)
				}
				newVersionMap[resourceName] = newVersion
			} else {
				// if oldVersion == "" this means that the resourse was already removed or desn't yet exist on the client
				// no need to remove it once again
				if oldVersion != "" {
					toRemove = append(toRemove, resourceName)
				}
				// the resource has gone but we keep watching for it so we detect an update if the resource is back
				newVersionMap[resourceName] = ""
			}
		}
	}

	// send through our version map
	return &RawDeltaResponse{
		DeltaRequest:     request,
		Resources:        filtered,
		RemovedResources: toRemove,
		VersionMap:       newVersionMap,
	}, nil
}