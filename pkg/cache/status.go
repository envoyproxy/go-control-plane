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

import "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"

// NodeHash computes string identifiers for Envoy nodes.
type NodeHash interface {
	// ID function defines a unique string identifier for the remote Envoy node.
	ID(node *core.Node) string
}

// StatusInfo tracks the server state for the remote Envoy node.
// Not all fields are used by all cache implementations.
type StatusInfo struct {
	// Node is the constant Envoy node metadata.
	Node *core.Node

	// ResponseWatches are indexed channels for the response watches and the original requests.
	ResponseWatches map[int64]ResponseWatch
}

// ResponseWatch is a watch record keeping both the request and an open channel for the response.
type ResponseWatch struct {
	// Request is the original request for the watch.
	Request Request

	// Response is the channel to push response to.
	Response chan Response
}

// NewStatusInfo initializes a status info data structure.
func NewStatusInfo(node *core.Node) StatusInfo {
	out := StatusInfo{
		Node:            node,
		ResponseWatches: make(map[int64]ResponseWatch),
	}
	return out
}
