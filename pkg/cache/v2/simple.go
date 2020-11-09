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
)

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
	SetSnapshot(node string, snapshot Snapshot) error

	// GetSnapshots gets the snapshot for a node.
	GetSnapshot(node string) (Snapshot, error)

	// ClearSnapshot removes all status and snapshot information associated with a node.
	ClearSnapshot(node string)

	// GetStatusInfo retrieves status information for a node ID.
	GetStatusInfo(string) StatusInfo

	// GetStatusKeys retrieves node IDs for all statuses.
	GetStatusKeys() []string
}

type heartbeatHandle struct {
	cancel func()
}

type snapshotCache struct {
	log log.Logger

	// ads flag to hold responses until all resources are named
	ads bool

	// snapshots are cached resources indexed by node IDs
	snapshots map[string]Snapshot

	// status information for all nodes indexed by node IDs
	status map[string]*statusInfo

	// hash is the hashing function for Envoy nodes
	hash NodeHash

	// watchCount is an atomic counter incremented for each watch
	watchCount int64

	heartbeatRoutines map[string]heartbeatHandle

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
	return &snapshotCache{
		log:               logger,
		ads:               ads,
		snapshots:         make(map[string]Snapshot),
		status:            make(map[string]*statusInfo),
		hash:              hash,
		heartbeatRoutines: make(map[string]heartbeatHandle),
	}
}

func (cache *snapshotCache) sendHeartbeats(ctx context.Context, wg *sync.WaitGroup, node string) {
	snapshot := cache.snapshots[node]
	if info, ok := cache.status[node]; ok {
		info.mu.Lock()
		for id, watch := range info.watches {
			// Respond with the current version regardless of whether the version has changed.
			version := snapshot.GetVersion(watch.Request.TypeUrl)
			resources := snapshot.GetResourcesAndTtl(watch.Request.TypeUrl)

			// TODO(snowp): Construct this once per type instead of once per watch.
			resourcesWithTtl := map[string]types.ResourceWithTtl{}
			for k, v := range resources {
				if v.Ttl != nil {
					resourcesWithTtl[k] = v
				}
			}

			if len(resourcesWithTtl) == 0 {
				continue
			}
			if cache.log != nil {
				cache.log.Debugf("respond open watch %d%v with heartbeat for version %q", id, watch.Request.ResourceNames, version)
			}

			wg.Add(1)
			// Spawn a goroutine here so that we can give up the lock before issuing responses. This prevents lock contention if the
			// response channels aren't ready to receive the response.
			go func(request *Request, response chan<- Response) {
				defer wg.Done()
				select {
				case <-ctx.Done():
					return
				default:
				}
				cache.respond(request, response, resourcesWithTtl, version, true)
			}(watch.Request, watch.Response)

			// The watch must be deleted and we must rely on the client to ack this response to create a new watch.
			delete(info.watches, id)
		}
		info.mu.Unlock()
	}
}

// SetSnapshotCache updates a snapshot for a node.
func (cache *snapshotCache) SetSnapshot(node string, snapshot Snapshot) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// update the existing entry
	cache.snapshots[node] = snapshot

	handle, ok := cache.heartbeatRoutines[node]
	// No heartbeat configured and no existing heartbeat, nothing to do.
	if ok {
		handle.cancel()
	}

	if snapshot.HeartbeatInterval != nil {
		heartBeatCtx, cancel := context.WithCancel(context.Background())
		cache.heartbeatRoutines[node] = heartbeatHandle{cancel: cancel}

		go func(heartbeatInterval time.Duration) {
			ticker := time.NewTicker(heartbeatInterval)
			for {
				select {
				case <-heartBeatCtx.Done():
					return
				case <-ticker.C:
					wg := &sync.WaitGroup{}
					cache.mu.Lock()

					// It's possible for this goroutine to be canceled while we're waiting to grab the lock. In this case
					// we want to avoid sending a heartbeat here and just end. If heartbeats should continue the another
					// goroutine will have been created.
					select {
					case <-heartBeatCtx.Done():
					default:
						cache.sendHeartbeats(heartBeatCtx, wg, node)
					}
					cache.mu.Unlock()

					// Block until we've finished sending all the heartbeats.
					wg.Wait()
				}
			}
		}(*snapshot.HeartbeatInterval)
	}

	// trigger existing watches for which version changed
	if info, ok := cache.status[node]; ok {
		info.mu.Lock()
		for id, watch := range info.watches {
			version := snapshot.GetVersion(watch.Request.TypeUrl)
			if version != watch.Request.VersionInfo {
				if cache.log != nil {
					cache.log.Debugf("respond open watch %d%v with new version %q", id, watch.Request.ResourceNames, version)
				}
				resources := snapshot.GetResourcesAndTtl(watch.Request.TypeUrl)
				cache.respond(watch.Request, watch.Response, resources, version, false)

				// discard the watch
				delete(info.watches, id)
			}
		}
		info.mu.Unlock()
	}

	return nil
}

// GetSnapshots gets the snapshot for a node, and returns an error if not found.
func (cache *snapshotCache) GetSnapshot(node string) (Snapshot, error) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	snap, ok := cache.snapshots[node]
	if !ok {
		return Snapshot{}, fmt.Errorf("no snapshot found for node %s", node)
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

// nameSet creates a map from a string slice to value true.
func nameSet(names []string) map[string]bool {
	set := make(map[string]bool)
	for _, name := range names {
		set[name] = true
	}
	return set
}

// superset checks that all resources are listed in the names set.
func superset(names map[string]bool, resources map[string]types.ResourceWithTtl) error {
	for resourceName := range resources {
		if _, exists := names[resourceName]; !exists {
			return fmt.Errorf("%q not listed", resourceName)
		}
	}
	return nil
}

// CreateWatch returns a watch for an xDS request.
func (cache *snapshotCache) CreateWatch(request *Request, value chan<- Response) func() {
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

	snapshot, exists := cache.snapshots[nodeID]
	version := snapshot.GetVersion(request.TypeUrl)

	// if the requested version is up-to-date or missing a response, leave an open watch
	if !exists || request.VersionInfo == version {
		watchID := cache.nextWatchID()
		if cache.log != nil {
			cache.log.Debugf("open watch %d for %s%v from nodeID %q, version %q", watchID,
				request.TypeUrl, request.ResourceNames, nodeID, request.VersionInfo)
		}
		info.mu.Lock()
		info.watches[watchID] = ResponseWatch{Request: request, Response: value}
		info.mu.Unlock()
		return cache.cancelWatch(nodeID, watchID)
	}

	// otherwise, the watch may be responded immediately
	resources := snapshot.GetResourcesAndTtl(request.TypeUrl)
	cache.respond(request, value, resources, version, false)

	return nil
}

func (cache *snapshotCache) nextWatchID() int64 {
	return atomic.AddInt64(&cache.watchCount, 1)
}

// cancellation function for cleaning stale watches
func (cache *snapshotCache) cancelWatch(nodeID string, watchID int64) func() {
	return func() {
		// uses the cache mutex
		cache.mu.Lock()
		defer cache.mu.Unlock()
		if info, ok := cache.status[nodeID]; ok {
			info.mu.Lock()
			delete(info.watches, watchID)
			info.mu.Unlock()
		}
	}
}

// Respond to a watch with the snapshot value. The value channel should have capacity not to block.
// TODO(kuat) do not respond always, see issue https://github.com/envoyproxy/go-control-plane/issues/46
func (cache *snapshotCache) respond(request *Request, value chan<- Response, resources map[string]types.ResourceWithTtl, version string, heartbeat bool) {
	// for ADS, the request names must match the snapshot names
	// if they do not, then the watch is never responded, and it is expected that envoy makes another request
	if len(request.ResourceNames) != 0 && cache.ads {
		if err := superset(nameSet(request.ResourceNames), resources); err != nil {
			if cache.log != nil {
				cache.log.Debugf("ADS mode: not responding to request: %v", err)
			}
			return
		}
	}
	if cache.log != nil {
		cache.log.Debugf("respond %s%v version %q with version %q",
			request.TypeUrl, request.ResourceNames, request.VersionInfo, version)
	}

	value <- createResponse(request, resources, version, heartbeat)
}

func createResponse(request *Request, resources map[string]types.ResourceWithTtl, version string, heartbeat bool) Response {
	filtered := make([]types.ResourceWithTtl, 0, len(resources))

	// Reply only with the requested resources. Envoy may ask each resource
	// individually in a separate stream. It is ok to reply with the same version
	// on separate streams since requests do not share their response versions.
	if len(request.ResourceNames) != 0 {
		set := nameSet(request.ResourceNames)
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

	return &RawResponse{
		Request:   request,
		Version:   version,
		Resources: filtered,
		Heartbeat: heartbeat,
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
			if cache.log != nil {
				cache.log.Warnf("skip fetch: version up to date")
			}
			return nil, &types.SkipFetchError{}
		}

		resources := snapshot.GetResourcesAndTtl(request.TypeUrl)
		out := createResponse(request, resources, version, false)
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
		if cache.log != nil {
			cache.log.Warnf("node does not exist")
		}
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
