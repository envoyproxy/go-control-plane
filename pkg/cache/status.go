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
	"strings"
	"sync"

	"github.com/envoyproxy/go-control-plane/api"
)

// NodeID provides the function to obtain a primary key from the node.
type NodeID func(*api.Node) string

// DefaultID is the default ID function combining the node cluster and the node ID.
func DefaultID(node *api.Node) string {
	return fmt.Sprintf("%s/%s", node.Cluster, node.Id)
}

// VersionCompare compares version identifiers and returns of three values, +1 if
// if first argument is greater than the second, 0 if equal, and -1 if less.
type VersionCompare func(string, string) int

// NodeInfo is node status.
type NodeInfo struct {
	id string
	mu sync.RWMutex

	// latest acknowledged version
	acked map[ResponseType]string

	// latest responded version
	responded map[ResponseType]string
}

// StatusInfo is a status tracker for configuration version acknowledgement by
// individual proxy nodes.
type StatusInfo struct {
	id      NodeID
	compare VersionCompare

	mu    sync.RWMutex
	nodes map[string]*NodeInfo
}

// NewStatusInfo creates a new status tracker.
func NewStatusInfo(id NodeID, compare VersionCompare) *StatusInfo {
	return &StatusInfo{
		id:      id,
		compare: compare,
		nodes:   make(map[string]*NodeInfo),
	}
}

// GetInfo retrieves node info or creates a record for the node.
func (info *StatusInfo) GetInfo(node *api.Node) *NodeInfo {
	info.mu.Lock()
	defer info.mu.Unlock()

	id := info.id(node)
	ni, exists := info.nodes[id]
	if !exists {
		ni = &NodeInfo{
			id:        id,
			acked:     make(map[ResponseType]string),
			responded: make(map[ResponseType]string),
		}
		info.nodes[id] = ni
	}
	return ni
}

// Ack records an acknowledgement for the latest version from a remote node.
func (info *StatusInfo) Ack(typ ResponseType, node *api.Node, version string) {
	ni := info.GetInfo(node)

	ni.mu.Lock()
	defer ni.mu.Unlock()
	if old, exists := ni.acked[typ]; !exists || info.compare(version, old) > 0 {
		ni.acked[typ] = version
	}
}

// AllowResponse checks whether a response should be allowed..
func (info *StatusInfo) AllowResponse(typ ResponseType, node *api.Node, version string) bool {
	ni := info.GetInfo(node)

	ni.mu.Lock()
	defer ni.mu.Unlock()
	if ni.responded[typ] == version {
		// reject if the version is responded once already
		// there are two main causes of rejection by the proxy:
		// - internal validity is violated:
		//   the response fails to validate and repeated responses do not lead to progress
		// - referential validity is violated (e.g. validate_clusters option):
		//   retries may lead to progress but this requires Envoy to enforce referential validity (a non-default behavior)
		//   it is important to stagger updates and not immediately reply in this case
		return false
	}

	ni.responded[typ] = version
	return true
}

// Dump serializes the status info into a string.
func (info *StatusInfo) Dump() string {
	info.mu.RLock()
	nodes := make([]*NodeInfo, 0, len(info.nodes))
	for _, ni := range info.nodes {
		nodes = append(nodes, ni)
	}
	info.mu.RUnlock()

	stats := make([]string, 0, 2*len(nodes)*len(ResponseTypes))
	for _, ni := range nodes {
		ni.mu.RLock()
		for _, typ := range ResponseTypes {
			stats = append(stats, ni.id+".acked."+typ.String()+"="+ni.acked[typ])
			stats = append(stats, ni.id+".responded."+typ.String()+"="+ni.responded[typ])
		}
		ni.mu.RUnlock()
	}

	return strings.Join(stats, "\n")
}
