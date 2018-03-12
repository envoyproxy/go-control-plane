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
	"sync"
	"time"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
)

// NodeHash computes string identifiers for Envoy nodes.
type NodeHash interface {
	// ID function defines a unique string identifier for the remote Envoy node.
	ID(node *core.Node) string
}

// StatusInfo tracks the server state for the remote Envoy node.
// Not all fields are used by all cache implementations.
type StatusInfo interface {
	// GetNode returns the node metadata.
	GetNode() *core.Node

	// GetNumWatches returns the number of open watches.
	GetNumWatches() int

	// GetLastWatchRequestTime returns the timestamp of the last discovery watch request.
	GetLastWatchRequestTime() time.Time

	// GetLastWatchVersion returns the version info from the last watch request
	// per type URL. Note that this is mostly useful in the ADS mode, since the
	// server only responds once for the entire set of resource names.
	GetLastWatchVersion(string) string
}

type statusInfo struct {
	// node is the constant Envoy node metadata.
	node *core.Node

	// watches are indexed channels for the response watches and the original requests.
	watches map[int64]ResponseWatch

	// the timestamp of the last watch request
	lastWatchRequestTime time.Time

	// the version from the last watch request per discovery type
	lastWatchEndpointsVersion string
	lastWatchClustersVersion  string
	lastWatchRoutesVersion    string
	lastWatchListenersVersion string

	// mutex to protect the status fields.
	// should not acquire mutex of the parent cache after acquiring this mutex.
	mu sync.RWMutex
}

// ResponseWatch is a watch record keeping both the request and an open channel for the response.
type ResponseWatch struct {
	// Request is the original request for the watch.
	Request Request

	// Response is the channel to push response to.
	Response chan Response
}

// newStatusInfo initializes a status info data structure.
func newStatusInfo(node *core.Node) *statusInfo {
	out := statusInfo{
		node:    node,
		watches: make(map[int64]ResponseWatch),
	}
	return &out
}

// setWatchVersion updates the last known accepted version from the watch request.
func (info *statusInfo) setWatchVersion(typeURL string, version string) {
	switch typeURL {
	case EndpointType:
		info.lastWatchEndpointsVersion = version
	case ClusterType:
		info.lastWatchClustersVersion = version
	case RouteType:
		info.lastWatchRoutesVersion = version
	case ListenerType:
		info.lastWatchListenersVersion = version
	}
}

func (info *statusInfo) GetNode() *core.Node {
	info.mu.RLock()
	defer info.mu.RUnlock()
	return info.node
}

func (info *statusInfo) GetNumWatches() int {
	info.mu.RLock()
	defer info.mu.RUnlock()
	return len(info.watches)
}

func (info *statusInfo) GetLastWatchRequestTime() time.Time {
	info.mu.RLock()
	defer info.mu.RUnlock()
	return info.lastWatchRequestTime
}

func (info *statusInfo) GetLastWatchVersion(typeURL string) string {
	switch typeURL {
	case EndpointType:
		return info.lastWatchEndpointsVersion
	case ClusterType:
		return info.lastWatchClustersVersion
	case RouteType:
		return info.lastWatchRoutesVersion
	case ListenerType:
		return info.lastWatchListenersVersion
	}
	return ""
}
