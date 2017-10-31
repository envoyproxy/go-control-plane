// Copyright 2017 Envoyproxy Authors
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
	"fmt"
	"log"
	"sync"

	"github.com/envoyproxy/go-control-plane/api"
	"github.com/golang/protobuf/proto"
)

// SimpleCache is a snapshot-based cache that maintains a single version per
// xDS response, node group tuple, with no canary updates.
type SimpleCache struct {
	// responses are cached resources
	responses map[Key]map[ResponseType][]proto.Message

	// versions must have the same key set as responses
	versions map[Key]map[ResponseType]int

	// watches keeps track of open watches
	watches map[Key]map[ResponseType]map[int64]Watch

	// callback requests missing responses
	callback func(Key, ResponseType)

	// watchCount is the ID generator for watches
	watchCount int64

	// groups is the hashing function for proxy nodes
	groups NodeGroup

	mu sync.Mutex
}

// NewSimpleCache initializes a simple cache.
// callback function is called on every new cache key and response type if there is no response available.
// callback is executed in a go-routine.
func NewSimpleCache(groups NodeGroup, callback func(Key, ResponseType)) Cache {
	return &SimpleCache{
		responses:  make(map[Key]map[ResponseType][]proto.Message),
		versions:   make(map[Key]map[ResponseType]int),
		watches:    make(map[Key]map[ResponseType]map[int64]Watch),
		callback:   callback,
		watchCount: 0,
		groups:     groups,
	}
}

func (cache *SimpleCache) SetResource(group Key, typ ResponseType, resources []proto.Message) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// update the existing entry
	if _, ok := cache.responses[group]; !ok {
		cache.responses[group] = make(map[ResponseType][]proto.Message)
		cache.versions[group] = make(map[ResponseType]int)
	}

	cache.responses[group][typ] = resources

	// increment version
	cache.versions[group][typ]++

	// trigger watches
	if _, ok := cache.watches[group]; ok {
		if _, ok := cache.watches[group][typ]; ok {
			// signal the update (channels have buffer size 1)
			for _, watch := range cache.watches[group][typ] {
				watch.Value <- Response{
					Version:   fmt.Sprintf("%d", cache.versions[group][typ]),
					Resources: resources,
				}
				// remove clean-up as the watch is discarded immediately
				watch.stop = nil
			}
			// discard all watches; the client must request a new watch to receive updates after ACK/NACK
			cache.watches[group][typ] = make(map[int64]Watch)
		}
	}
}

func (cache *SimpleCache) Watch(typ ResponseType, node *api.Node, version string, names []string) Watch {
	// do nothing case
	if node == nil {
		return Watch{}
	}

	group := cache.groups.Hash(node)

	cache.mu.Lock()
	defer cache.mu.Unlock()

	// if the requested version is up-to-date or missing a response, leave an open watch
	versions, exists := cache.versions[group]
	if !exists || version == fmt.Sprintf("%d", versions[typ]) {
		// invoke callback in a go-routine
		if !exists && cache.callback != nil {
			log.Printf("requesting resources for %v from %q at %q", typ, group, version)
			go cache.callback(group, typ)
		}

		// make sure watches map exists
		if _, ok := cache.watches[group]; !ok {
			cache.watches[group] = make(map[ResponseType]map[int64]Watch)
		}
		if _, ok := cache.watches[group][typ]; !ok {
			cache.watches[group][typ] = make(map[int64]Watch)
		}

		log.Printf("open watch for %s%v from key %q at version %q", typ.String(), names, group, version)
		value := make(chan Response, 1)
		cache.watchCount++
		id := cache.watchCount
		out := Watch{
			Value: value,
			stop: func() {
				if _, ok := cache.watches[group]; ok {
					if _, ok := cache.watches[group][typ]; ok {
						delete(cache.watches[group][typ], id)
					}
				}
			},
		}
		cache.watches[group][typ][id] = out
		return out
	}

	// otherwise, respond with the latest version
	// TODO(kuat) this responds immediately and can cause the remote node to spin if it consistently fails to ACK the update
	log.Printf("respond for %v from %q at %q", typ, group, version)
	value := make(chan Response, 1)
	value <- Response{
		Version:   fmt.Sprintf("%d", cache.versions[group][typ]),
		Resources: cache.responses[group][typ],
	}
	return Watch{Value: value}
}

func (cache *SimpleCache) WatchEndpoints(node *api.Node, version string, names []string) Watch {
	return cache.Watch(EndpointResponse, node, version, names)
}

func (cache *SimpleCache) WatchClusters(node *api.Node, version string, names []string) Watch {
	return cache.Watch(ClusterResponse, node, version, names)
}

func (cache *SimpleCache) WatchRoutes(node *api.Node, version string, names []string) Watch {
	return cache.Watch(RouteResponse, node, version, names)
}

func (cache *SimpleCache) WatchListeners(node *api.Node, version string, names []string) Watch {
	return cache.Watch(ListenerResponse, node, version, names)
}
