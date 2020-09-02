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
	"errors"
	"fmt"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
)

// Resources is a versioned group of resources.
type Resources struct {
	// Version information.
	Version string

	// Items in the group indexed by name.
	Items map[string]types.Resource
}

// DeltaResources is a versioned group of resources which also contains individual resource versions per the incremental xDS protocol
type DeltaResources struct {
	// Version information
	SystemVersion string

	// Items in the group indexed by name
	Items resourceItems
}

type resourceItems struct {
	Version string
	Items   map[string]types.Resource
}

// IndexResourcesByName creates a map from the resource name to the resource.
func IndexResourcesByName(items []types.Resource) map[string]types.Resource {
	indexed := make(map[string]types.Resource, len(items))
	for _, item := range items {
		indexed[GetResourceName(item)] = item
	}
	return indexed
}

// NewResources creates a new resource group.
func NewResources(version string, items []types.Resource) Resources {
	return Resources{
		Version: version,
		Items:   IndexResourcesByName(items),
	}
}

// Snapshot is an internally consistent snapshot of xDS resources.
// Consistency is important for the convergence as different resource types
// from the snapshot may be delivered to the proxy in arbitrary order.
type Snapshot struct {
	Resources [types.UnknownType]Resources
}

// NewSnapshot creates a snapshot from response types and a version.
func NewSnapshot(version string,
	endpoints []types.Resource,
	clusters []types.Resource,
	routes []types.Resource,
	listeners []types.Resource,
	runtimes []types.Resource,
	secrets []types.Resource) Snapshot {
	out := Snapshot{}
	out.Resources[types.Endpoint] = NewResources(version, endpoints)
	out.Resources[types.Cluster] = NewResources(version, clusters)
	out.Resources[types.Route] = NewResources(version, routes)
	out.Resources[types.Listener] = NewResources(version, listeners)
	out.Resources[types.Runtime] = NewResources(version, runtimes)
	out.Resources[types.Secret] = NewResources(version, secrets)
	return out
}

// Consistent check verifies that the dependent resources are exactly listed in the
// snapshot:
// - all EDS resources are listed by name in CDS resources
// - all RDS resources are listed by name in LDS resources
//
// Note that clusters and listeners are requested without name references, so
// Envoy will accept the snapshot list of clusters as-is even if it does not match
// all references found in xDS.
func (s *Snapshot) Consistent() error {
	if s == nil {
		return errors.New("nil snapshot")
	}
	endpoints := GetResourceReferences(s.Resources[types.Cluster].Items)
	if len(endpoints) != len(s.Resources[types.Endpoint].Items) {
		return fmt.Errorf("mismatched endpoint reference and resource lengths: %v != %d", endpoints, len(s.Resources[types.Endpoint].Items))
	}
	if err := superset(endpoints, s.Resources[types.Endpoint].Items); err != nil {
		return err
	}

	routes := GetResourceReferences(s.Resources[types.Listener].Items)
	if len(routes) != len(s.Resources[types.Route].Items) {
		return fmt.Errorf("mismatched route reference and resource lengths: %v != %d", routes, len(s.Resources[types.Route].Items))
	}
	return superset(routes, s.Resources[types.Route].Items)
}

// GetResources selects snapshot resources by type.
func (s *Snapshot) GetResources(typeURL string) map[string]types.Resource {
	if s == nil {
		return nil
	}
	typ := GetResponseType(typeURL)
	if typ == types.UnknownType {
		return nil
	}
	return s.Resources[typ].Items
}

// GetSubscribedResources selects requested snapshot resources by type and alias.
// This function is used for Incremental/Delta xDS and follows the subscribed resource model.
func (s *Snapshot) GetSubscribedResources(aliases []string, typeURL string) map[string]types.Resource {
	if s == nil {
		return nil
	}

	t := GetResponseType(typeURL)
	if t == types.UnknownType {
		return nil
	}

	subscribed := make(map[string]types.Resource, len(aliases))
	r := s.Resources[t].Items

	// TODO:
	// This right now is O(n^2) which is not performant. Will need to revisit
	for _, item := range r {
		for _, alias := range aliases {
			if GetResourceName(item) == alias {
				subscribed[alias] = item
			}
		}
	}

	return subscribed
}

// GetResource will return a single resource by alias within a given resource group
func (s *Snapshot) GetResource(alias string, typeURL string) types.Resource {
	if s == nil {
		return nil
	}

	t := GetResponseType(typeURL)
	if t == types.UnknownType {
		return nil
	}

	for key, item := range s.Resources[t].Items {
		if alias == key {
			return item
		}
	}

	return nil
}

// GetVersion returns the version for a resource type.
func (s *Snapshot) GetVersion(typeURL string) string {
	if s == nil {
		return ""
	}
	typ := GetResponseType(typeURL)
	if typ == types.UnknownType {
		return ""
	}
	return s.Resources[typ].Version
}
