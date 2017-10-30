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

import "github.com/golang/protobuf/proto"

// SimpleCache is a snapshot-based cache that maintains a single version per
// xDS response, node group tuple.
type SimpleCache struct {
	responses map[CacheKey][ResponseType]proto.Message
	versions  map[CacheKey][ResponseType]int
}

func (cache *SimpleCache) SetResource(key CacheKey, typ ResponseType) {
	// update the existing entry

	// trigger watches
}

type SimpleWatcher struct {
}

func (watcher *SimpleWatcher) WatchEndpoints(*api.Node, string, []string) {
	return watcher.Watch(EndpointType, node, version, 
}

