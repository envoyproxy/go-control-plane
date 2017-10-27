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

// ConfigCache ...
type ConfigCache interface {
	// Listen requests a channel for receiving configuration resources for
	// a given node, config resources, and last applied version identifier
	Listen(*api.Node, ResourceSelector, string) chan api.DiscoveryResponse
}

// ResourceSelector for selecting monitored configuration resources
type ResourceSelector struct {
	// Types of configuration resources to monitor (or all if empty)
	Types []string

	// Names of configuration resources (or all if empty)
	Names []string
}

// Resource types in xDS v2
const (
	typePrefix   = "type.googleapis.com/envoy.api.v2."
	EndpointType = typePrefix + "LbEndpoint"
	ClusterType  = typePrefix + "Cluster"
	RouteType    = typePrefix + "Route"
	ListenerType = typePrefix + "Listener"
)
