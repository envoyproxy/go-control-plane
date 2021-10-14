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
	"sync"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

// Snapshot is an internally consistent snapshot of xDS resources.
type Snapshot interface {
	// GetVersion returns the current version of the resources indicated
	// by typeURL. The version string that is returned is opaque and should
	// only be compared for equality. This should strictly be used for SOTW and
	// is not a valid version identifier for Delta xDS (see ConstructVersionMap).
	GetVersion(typeURL string) string

	// GetResourcesAndTTL returns all resources of the type indicted by
	// typeURL, together with their TTL.
	GetResourcesAndTTL(typeURL string) map[string]types.ResourceWithTTL

	// GetResources returns all resources of the type indicted by
	// typeURL. This is identical to GetResourcesAndTTL, except that
	// the TTL is omitted.
	GetResources(typeURL string) map[string]types.Resource

	// SetResourcesByType sets a grouping of cache resources at the specificed
	// resource type in a Snapshot.
	SetResourcesByType(typeURL string, resources Resources) error

	// ConstructVersionMap constructs an internal opaque version string for each resource in a collection.
	// If this is not done correctly the cache system won't be able to properly diff resources when using Delta xDSs.
	ConstructVersionMap() error

	// GetVersionMap returns a map of resource name to resource version for
	// all the resources of type indicated by typeURL.
	GetVersionMap(typeURL string) map[string]string

	// Consistent checks that the snapshot is internally consistent.and all references resources are listed.
	Consistent() error
}

// Snapshot is an internally consistent snapshot of xDS resources.
// Consistency is important for the convergence as different resource types
// from the snapshot may be delivered to the proxy in arbitrary order.
type snapshot struct {
	Resources [types.UnknownType]Resources

	// VersionMap holds the current hash map of all resources in the snapshot.
	// This field should remain nil until it is used, at which point should be
	// instantiated by calling ConstructVersionMap().
	// VersionMap is only to be used with delta xDS.
	VersionMap map[string]map[string]string

	// Lock for multi-thread processing
	mu sync.RWMutex
}

// NewSnapshot creates a snapshot from response types and a version.
// The resources map is keyed off the type URL of a resource, followed by the slice of resource objects.
func NewSnapshot(version string, resources map[resource.Type][]types.Resource) (Snapshot, error) {
	out := &snapshot{}

	// Return a nil snapshot if there are no provided resources.
	if resources == nil {
		return out, nil
	}

	for typ, resource := range resources {
		index := GetResponseType(typ)
		if index == types.UnknownType {
			return out, errors.New("unknown resource type: " + typ)
		}

		out.Resources[index] = NewResources(version, resource)
	}

	return out, nil
}

// NewSnapshotWithTTLs creates a snapshot of ResourceWithTTLs.
// The resources map is keyed off the type URL of a resource, followed by the slice of resource objects.
func NewSnapshotWithTTLs(version string, resources map[resource.Type][]types.ResourceWithTTL) (Snapshot, error) {
	out := &snapshot{}

	// Return a nil snapshot if there are no provided resources.
	if resources == nil {
		return out, nil
	}

	for typ, resource := range resources {
		index := GetResponseType(typ)
		if index == types.UnknownType {
			return out, errors.New("unknown resource type: " + typ)
		}

		out.Resources[index] = NewResourcesWithTTL(version, resource)
	}

	return out, nil
}

// Consistent check verifies that the dependent resources are exactly listed in the
// snapshot:
// - all EDS resources are listed by name in CDS resources
// - all SRDS/RDS resources are listed by name in LDS resources
// - all RDS resources are listed by name in SRDS resources
// - empty resource lists are considered consistent since there are no references
//
// Note that clusters and listeners are requested without name references, so
// Envoy will accept the snapshot list of clusters as-is even if it does not match
// all references found in xDS.
func (s *snapshot) Consistent() error {
	if s == nil {
		return errors.New("nil snapshot")
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

	referencedResources := GetAllResourceReferences(s.Resources)

	// Loop through each referenced resource.
	referencedResponseTypes := map[types.ResponseType]struct{}{
		types.Endpoint:    {},
		types.ScopedRoute: {},
		types.Route:       {},
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
				return fmt.Errorf("mismatched reference and resource lengths: len(%v) != %d", referenceSet, len(items.Items))
			}

			// Check superset.
			if err := superset(referenceSet, items.Items); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetResources selects snapshot resources by type, returning the map of resources.
func (s *snapshot) GetResources(typeURL resource.Type) map[string]types.Resource {
	if s == nil {
		return nil
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

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
func (s *snapshot) GetResourcesAndTTL(typeURL resource.Type) map[string]types.ResourceWithTTL {
	if s == nil {
		return nil
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

	typ := GetResponseType(typeURL)
	if typ == types.UnknownType {
		return nil
	}
	return s.Resources[typ].Items
}

// SetResourcesByType sets a grouping of cache resources at the specificed typeURL in a snapshot.
func (s *snapshot) SetResourcesByType(typeURL resource.Type, resources Resources) error {
	if s == nil {
		return errors.New("snapshot can not be nil")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Resources[GetResponseType(typeURL)] = resources
	return nil
}

// GetVersion returns the version for a resource type.
func (s *snapshot) GetVersion(typeURL resource.Type) string {
	if s == nil {
		return ""
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

	typ := GetResponseType(typeURL)
	if typ == types.UnknownType {
		return ""
	}
	return s.Resources[typ].Version
}

// GetVersionMap will return the internal version map of the currently applied snapshot.
func (s *snapshot) GetVersionMap(typeURL string) map[string]string {
	if s == nil {
		return nil
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.VersionMap[typeURL]
}

// ConstructVersionMap will construct a version map based on the current state of a snapshot
func (s *snapshot) ConstructVersionMap() error {
	if s == nil {
		return fmt.Errorf("missing snapshot")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	// TODO: why was this ever in here? It doesn't make sense
	// The snapshot resources never change, so no need to ever rebuild.
	// if s.VersionMap != nil {
	// 	return nil
	// }

	s.VersionMap = make(map[string]map[string]string)

	for i, resources := range s.Resources {
		typeURL, err := GetResponseTypeURL(types.ResponseType(i))
		if err != nil {
			return err
		}
		if _, ok := s.VersionMap[typeURL]; !ok {
			s.VersionMap[typeURL] = make(map[string]string)
		}

		for _, r := range resources.Items {
			// Hash our version in here and build the version map.
			marshaledResource, err := MarshalResource(r.Resource)
			if err != nil {
				return err
			}
			v := HashResource(marshaledResource)
			if v == "" {
				return fmt.Errorf("failed to build resource version: %v", err)
			}

			s.VersionMap[typeURL][GetResourceName(r.Resource)] = v
		}
	}

	return nil
}
