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

type ttlHandle struct {
	ttl    time.Duration
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

	ttls map[string]map[int]ttlHandle

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
		log:       logger,
		ads:       ads,
		snapshots: make(map[string]Snapshot),
		status:    make(map[string]*statusInfo),
		hash:      hash,
		ttls:      make(map[string]map[int]ttlHandle),
	}
}

// SetSnapshotCache updates a snapshot for a node.
func (cache *snapshotCache) SetSnapshot(node string, snapshot Snapshot) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// update the existing entry
	cache.snapshots[node] = snapshot

	if _, ok := cache.ttls[node]; !ok {
		cache.ttls[node] = make(map[int]ttlHandle)
	}

	for typeUrl, resource := range snapshot.Resources {
		ttl, ok := cache.ttls[node][typeUrl]
		// No TTL configured and no existing TTL, nothing to do.
		if !ok && resource.Ttl == nil {
			continue
		}
		// Exisiting TTL matches new TTL, nothing to do.
		if ok && resource.Ttl != nil && *resource.Ttl == ttl.ttl {
			continue
		}

		// No TTL, nothing to do.
		if resource.Ttl == nil {
			continue
		}

		// Clean up the old goroutine, then spin up a new one with the new TTL.
		if ok {
			ttl.cancel()
		}

		ctx, cancel := context.WithCancel(context.Background())
		cache.ttls[node][typeUrl] = ttlHandle{ttl: *resource.Ttl, cancel: cancel}
		go func(ttl time.Duration, typeUrl types.ResponseType) {
			// By default we heartbeat TTL'd resources at 1/3 of the TTL duration. This is somewhat arbitrary but is meant to reduce the likelihood for a race between the
			// heartbeat being sent and the client removing the resource due to hitting the TTL.
			// TODO(snowp): We probably would want to add some safety guards here: a 100ms TTL on a lot of resources would cause a lot of stress on the server, both through
			// lock contention on the cache and gRPC serialization related costs.
			ticker := time.NewTicker(ttl / 3)
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					cache.mu.Lock()

					// Trigger heartbeat updates for all watches for this particular resource.
					if info, ok := cache.status[node]; ok {
						info.mu.Lock()
						for id, watch := range info.watches {
							// This goroutine is only responsible for triggering updates for a specific type.
							if GetResponseType(watch.Request.TypeUrl) != typeUrl {
								continue
							}

							snapshot := cache.snapshots[node]
							// Respond with the current version regardless of whether the version has changed.
							version := snapshot.GetVersion(watch.Request.TypeUrl)
							resources, ttl := snapshot.GetResourcesAndTtl(watch.Request.TypeUrl)
							if cache.log != nil {
								cache.log.Debugf("respond open watch %d%v with heartbeat for version %q", id, watch.Request.ResourceNames, version)
							}
							cache.respond(watch.Request, watch.Response, resources, version, ttl)

							// The watch must be deleted and we must rely on the client to ack this response to create a new watch.
							delete(info.watches, id)
						}
						info.mu.Unlock()
					}
					cache.mu.Unlock()
				}
			}
		}(*resource.Ttl, types.ResponseType(typeUrl))
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
				resources, ttl := snapshot.GetResourcesAndTtl(watch.Request.TypeUrl)
				cache.respond(watch.Request, watch.Response, resources, version, ttl)

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
func superset(names map[string]bool, resources map[string]types.Resource) error {
	for resourceName := range resources {
		if _, exists := names[resourceName]; !exists {
			return fmt.Errorf("%q not listed", resourceName)
		}
	}
	return nil
}

// CreateWatch returns a watch for an xDS request.
func (cache *snapshotCache) CreateWatch(request *Request) (chan Response, func()) {
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

	// allocate capacity 1 to allow one-time non-blocking use
	value := make(chan Response, 1)

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
		return value, cache.cancelWatch(nodeID, watchID)
	}

	// otherwise, the watch may be responded immediately
	resources, ttl := snapshot.GetResourcesAndTtl(request.TypeUrl)
	cache.respond(request, value, resources, version, ttl)

	return value, nil
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
func (cache *snapshotCache) respond(request *Request, value chan Response, resources map[string]types.Resource, version string, ttl *time.Duration) {
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

	value <- createResponse(request, resources, version, ttl)
}

func createResponse(request *Request, resources map[string]types.Resource, version string, ttl *time.Duration) Response {
	filtered := make([]types.Resource, 0, len(resources))

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
		Ttl:       ttl,
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

		resources, ttl := snapshot.GetResourcesAndTtl(request.TypeUrl)
		out := createResponse(request, resources, version, ttl)
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
