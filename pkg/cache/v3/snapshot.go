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
	"maps"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

// Snapshot is an internally consistent snapshot of xDS resources.
// Consistency is important for the convergence as different resource types
// from the snapshot may be delivered to the proxy in arbitrary order.
// Deprecated: use types.Snapshot instead.
type Snapshot struct {
	Resources [types.UnknownType]Resources
}

var _ ResourceSnapshot = &Snapshot{}

// NewSnapshot creates a snapshot from response types and a version.
// The resources map is keyed off the type URL of a resource, followed by the slice of resource objects.
// Deprecated: use types.NewSnapshot or types.NewSnapshotFromTypeSnapshots instead.
func NewSnapshot(version string, resources map[resource.Type][]types.Resource) (*Snapshot, error) {
	out := Snapshot{}

	for typ, resource := range resources {
		index := GetResponseType(typ)
		if index == types.UnknownType {
			return nil, errors.New("unknown resource type: " + typ)
		}

		out.Resources[index] = NewResources(version, resource)
	}

	return &out, nil
}

// NewSnapshotWithTTLs creates a snapshot of ResourceWithTTLs.
// The resources map is keyed off the type URL of a resource, followed by the slice of resource objects.
// Deprecated: use types.NewSnapshot or types.NewSnapshotFromTypeSnapshots instead.
func NewSnapshotWithTTLs(version string, resources map[resource.Type][]types.ResourceWithTTL) (*Snapshot, error) {
	out := Snapshot{}

	for typ, resource := range resources {
		index := GetResponseType(typ)
		if index == types.UnknownType {
			return nil, errors.New("unknown resource type: " + typ)
		}

		out.Resources[index] = NewResourcesWithTTL(version, resource)
	}

	return &out, nil
}

// Consistent check verifies that the dependent resources are exactly listed in the
// snapshot:
// - all EDS resources are listed by name in CDS resources
// - all SRDS/RDS resources are listed by name in LDS resources
// - all RDS resources are listed by name in SRDS resources
//
// Note that clusters and listeners are requested without name references, so
// Envoy will accept the snapshot list of clusters as-is even if it does not match
// all references found in xDS.
func (s *Snapshot) Consistent() error {
	if s == nil {
		return errors.New("nil snapshot")
	}

	referencedResources := GetAllResourceReferences(s.Resources)

	// Loop through each referenced resource.
	referencedResponseTypes := map[types.ResponseType]struct{}{
		types.Endpoint: {},
		types.Route:    {},
	}

	for idx, items := range s.Resources {
		// We only want to check resource types that are expected to be referenced by another resource type.
		// Basically, if the consistency relationship is modeled as a DAG, we only want
		// to check nodes that are expected to have edges pointing to it.
		responseType := types.ResponseType(idx)
		if _, ok := referencedResponseTypes[responseType]; ok {
			typeURL, err := GetResponseTypeURL(responseType)
			if err != nil {
				return err
			}
			referenceSet := referencedResources[typeURL]

			if len(referenceSet) != len(items.Items) {
				return fmt.Errorf("mismatched %q reference and resource lengths: len(%v) != %d",
					typeURL, referenceSet, len(items.Items))
			}

			// Check superset.
			if missing := difference(maps.Keys(items.Items), referenceSet); len(missing) > 0 {
				return fmt.Errorf("inconsistent %q reference: missing resources %v", typeURL, missing)
			}
		}
	}

	return nil
}

// GetResources selects snapshot resources by type, returning the map of resources.
func (s *Snapshot) GetResources(typeURL resource.Type) map[string]types.Resource {
	resources := s.GetResourcesAndTTL(typeURL)
	if resources == nil {
		return nil
	}

	withoutTTL := make(map[string]types.Resource, len(resources))

	for k, v := range resources {
		withoutTTL[k] = v.Resource
	}

	return withoutTTL
}

// GetResourcesAndTTL selects snapshot resources by type, returning the map of resources and the associated TTL.
func (s *Snapshot) GetResourcesAndTTL(typeURL resource.Type) map[string]types.ResourceWithTTL {
	if s == nil {
		return nil
	}
	typ := GetResponseType(typeURL)
	if typ == types.UnknownType {
		return nil
	}
	return s.Resources[typ].Items
}

// GetVersion returns the version for a resource type.
func (s *Snapshot) GetVersion(typeURL resource.Type) string {
	if s == nil {
		return ""
	}
	typ := GetResponseType(typeURL)
	if typ == types.UnknownType {
		return ""
	}
	return s.Resources[typ].Version
}

func (s *Snapshot) GetTypeSnapshot(typeURL string) types.TypeSnapshot {
	if s == nil {
		return types.TypeSnapshot{}
	}
	typ := GetResponseType(typeURL)
	if typ == types.UnknownType {
		return types.TypeSnapshot{}
	}

	items := s.Resources[typ].Items
	resources := make([]types.SnapshotResource, 0, len(items))
	for name, res := range items {
		resources = append(resources, types.SnapshotResource{
			Name:     name,
			Resource: res.Resource,
			TTL:      res.TTL,
		})
	}
	return types.NewTypeSnapshot(typeURL, s.GetVersion(typeURL), resources)
}
