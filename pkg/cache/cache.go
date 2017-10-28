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

// Package cache defines a configuration cache for the server
package cache

import (
	"github.com/envoyproxy/go-control-plane/api"
	"github.com/golang/protobuf/proto"
)

// ConfigWatcher requests watches for configuration resources by a node, last
// applied version identifier, and resource names hint.  The watch should send
// the responses when they are ready.  The watch can be cancelled by the
// consumer, in effect terminating the watch for the request.
type ConfigWatcher interface {
	// WatchEndpoints resources
	WatchEndpoints(*api.Node, string, []string) Watch

	// WatchClusters resources
	WatchClusters(*api.Node, string, []string) Watch

	// WatchRoutes resources
	WatchRoutes(*api.Node, string, []string) Watch

	// WatchListeners resources
	WatchListeners(*api.Node, string, []string) Watch
}

// Watch is a dedicated stream of configuration resources
type Watch struct {
	// Value channel should send at least one response.  The producer can eagerly
	// send multiple values ahead of the consumer requesting a newer watch.  The
	// producer can close the channel if an unrecoverable error occurs, in which
	// case the consumer stream is terminated.
	Value <-chan Response

	stop func()
}

// Response is a pre-serialized response for an assumed type
type Response struct {
	// Version of the resources as tracked by the cache for the given type.
	// Proxy responds with this version if the resources get applied remotely.
	Version string

	// Resources to be included in the response
	Resources []proto.Message

	// Canary bit to control Envoy config transition
	Canary bool
}

// Cancel the watch.
func (watch Watch) Cancel() {
	if watch.stop != nil {
		watch.stop()
	}
}

// NodeGroup aggregates configuration resources by a hash of the node.
type NodeGroup interface {
	Hash(*api.Node) string
}
