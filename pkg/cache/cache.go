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

import "github.com/envoyproxy/go-control-plane/api"

// ConfigWatcher requests promises for configuration resources by a node,
// last applied version identifier and resource names hint.
// The promise should send the response once it is ready or be cancelled
// by the servier, in effect terminating the computation for the request.
type ConfigWatcher interface {
	// WatchEndpoints resources
	WatchEndpoints(*api.Node, string, []string) Promise

	// WatchClusters resources
	WatchClusters(*api.Node, string, []string) Promise

	// WatchRoutes resources
	WatchRoutes(*api.Node, string, []string) Promise

	// WatchListeners resources
	WatchListeners(*api.Node, string, []string) Promise
}

// Promise is used once to complete a response asynchronously
type Promise struct {
	// Value is to be used once
	Value <-chan api.DiscoveryResponse

	// Stop terminates the promise computation
	Stop func()
}

// Cancel the promise computation
func (promise Promise) Cancel() {
	if promise.Stop != nil {
		promise.Stop()
	}
}
