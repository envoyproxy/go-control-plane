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
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

// Snapshot is an internally consistent snapshot of xDS resources.
// Consistency is important for the convergence as different resource types
// from the snapshot may be delivered to the proxy in arbitrary order.
type Snapshot struct {
	Resources [types.UnknownType]Resources

	// VersionMap holds the current hash map of all resources in the snapshot.
	// This field should remain nil until it is used, at which point should be
	// instantiated by calling ConstructVersionMap().
	// VersionMap is only to be used with delta xDS.
	VersionMap map[string]map[string]string

	// inheritedVersions holds versions taken from a previous snapshot for resources
	// that are the very same object under the same name. ConstructVersionMap uses
	// them instead of re-hashing and then drops them.
	inheritedVersions map[string]map[string]string

	// marshaled holds the bytes encoded while hashing in ConstructVersionMap when
	// keepMarshaled is set, so a delta response doesn't encode the same resource again.
	marshaled     map[string]map[string][]byte
	keepMarshaled bool
}

var _ ResourceSnapshot = &Snapshot{}

// NewSnapshot creates a snapshot from response types and a version.
// The resources map is keyed off the type URL of a resource, followed by the slice of resource objects.
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

			// Check difference.
			if missing := difference(items.Items, referenceSet); len(missing) > 0 {
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

// GetVersionMap will return the internal version map of the currently applied snapshot.
func (s *Snapshot) GetVersionMap(typeURL string) map[string]string {
	return s.VersionMap[typeURL]
}

// ConstructVersionMap will construct a version map based on the current state of a snapshot.
func (s *Snapshot) ConstructVersionMap() error {
	if s == nil {
		return errors.New("missing snapshot")
	}

	// The snapshot resources never change, so no need to ever rebuild.
	if s.VersionMap != nil {
		return nil
	}

	versionMap := make(map[string]map[string]string)
	marshaled := make(map[string]map[string][]byte)

	for i, resources := range s.Resources {
		typeURL, err := GetResponseTypeURL(types.ResponseType(i))
		if err != nil {
			return err
		}
		versions := make(map[string]string, len(resources.Items))
		versionMap[typeURL] = versions
		inherited := s.inheritedVersions[typeURL]

		for name, r := range resources.Items {
			if v, ok := inherited[name]; ok {
				versions[name] = v
				continue
			}

			marshaledResource, err := MarshalResource(r.Resource)
			if err != nil {
				return err
			}
			v := HashResource(marshaledResource)
			if v == "" {
				return errors.New("failed to build resource version")
			}

			versions[name] = v
			if !s.keepMarshaled {
				continue
			}
			if marshaled[typeURL] == nil {
				marshaled[typeURL] = make(map[string][]byte)
			}
			marshaled[typeURL][name] = marshaledResource
		}
	}

	s.VersionMap = versionMap
	s.marshaled = marshaled
	s.inheritedVersions = nil

	return nil
}

// inheritVersions records, for each resource that is the same object under the same
// name in prev, the version prev already computed, so ConstructVersionMap can skip
// encoding and hashing it. It relies on resources never being modified once they are
// part of a snapshot.
func (s *Snapshot) inheritVersions(prev *Snapshot) {
	if s == nil || prev == nil || s == prev || s.VersionMap != nil || prev.VersionMap == nil {
		return
	}

	for i, resources := range s.Resources {
		// Every index of Resources is a known type, so the lookup can't fail.
		typeURL, _ := GetResponseTypeURL(types.ResponseType(i))
		prevItems := prev.Resources[i].Items
		prevVersions := prev.VersionMap[typeURL]
		if len(prevItems) == 0 || len(prevVersions) == 0 {
			continue
		}

		for name, r := range resources.Items {
			prevRes, sameName := prevItems[name]
			v, hasVersion := prevVersions[name]
			if !sameName || !hasVersion || prevRes.Resource != r.Resource {
				continue
			}
			if s.inheritedVersions == nil {
				s.inheritedVersions = make(map[string]map[string]string)
			}
			if s.inheritedVersions[typeURL] == nil {
				s.inheritedVersions[typeURL] = make(map[string]string)
			}
			s.inheritedVersions[typeURL][name] = v
		}
	}
}

// getMarshaledResource returns the bytes encoded for name while building the version
// map, if any.
func (s *Snapshot) getMarshaledResource(typeURL, name string) []byte {
	return s.marshaled[typeURL][name]
}
