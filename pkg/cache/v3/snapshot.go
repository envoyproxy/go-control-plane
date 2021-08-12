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
	"context"
	"errors"
	"fmt"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/log"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// Resources is a versioned group of resources.
type Resources struct {
	// Version information.
	Version string

	// Items in the group indexed by name.
	Items map[string]types.ResourceWithTTL
}

// DeltaResources is a versioned group of resources which also contains individual resource versions per the incremental xDS protocol
type DeltaResources struct {
	// Version information
	SystemVersion string

	// Items in the group indexed by name
	Items resourceItems
}

// resourceItems contain the lower level versioned resource map
type resourceItems struct {
	Version string
	Items   map[string]types.ResourceWithTTL
}

// IndexResourcesByName creates a map from the resource name to the resource.
func IndexResourcesByName(items []types.ResourceWithTTL) map[string]types.ResourceWithTTL {
	indexed := make(map[string]types.ResourceWithTTL)
	for _, item := range items {
		indexed[GetResourceName(item.Resource)] = item
	}
	return indexed
}

// IndexRawResourcesByName creates a map from the resource name to the resource.
func IndexRawResourcesByName(items []types.Resource) map[string]types.Resource {
	indexed := make(map[string]types.Resource)
	for _, item := range items {
		indexed[GetResourceName(item)] = item
	}
	return indexed
}

// NewResources creates a new resource group.
func NewResources(version string, items []types.Resource) Resources {
	itemsWithTTL := []types.ResourceWithTTL{}
	for _, item := range items {
		itemsWithTTL = append(itemsWithTTL, types.ResourceWithTTL{Resource: item})
	}
	return NewResourcesWithTtl(version, itemsWithTTL)
}

// NewResources creates a new resource group.
func NewResourcesWithTtl(version string, items []types.ResourceWithTTL) Resources { // nolint:golint,revive
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

	// VersionMap holds the current hash map of all resources in the snapshot.
	// This field should remain nil until it is used, at which point should be
	// instantiated by calling ConstructVersionMap().
	// VersionMap is only to be used with delta xDS.
	VersionMap map[string]map[string]string
}

// NewSnapshot creates a snapshot from response types and a version.
func NewSnapshot(version string,
	endpoints []types.Resource,
	clusters []types.Resource,
	routes []types.Resource,
	listeners []types.Resource,
	runtimes []types.Resource,
	secrets []types.Resource,
	extensionConfigs []types.Resource) Snapshot {
	return NewSnapshotWithResources(version, SnapshotResources{
		Endpoints:        endpoints,
		Clusters:         clusters,
		Routes:           routes,
		Listeners:        listeners,
		Runtimes:         runtimes,
		Secrets:          secrets,
		ExtensionConfigs: extensionConfigs,
	})
}

// SnapshotResources contains the resources to construct a snapshot from.
type SnapshotResources struct {
	Endpoints        []types.Resource
	Clusters         []types.Resource
	Routes           []types.Resource
	Listeners        []types.Resource
	Runtimes         []types.Resource
	Secrets          []types.Resource
	ExtensionConfigs []types.Resource
}

// NewSnapshotWithResources creates a snapshot from response types and a version.
func NewSnapshotWithResources(version string, resources SnapshotResources) Snapshot {
	out := Snapshot{}
	out.Resources[types.Endpoint] = NewResources(version, resources.Endpoints)
	out.Resources[types.Cluster] = NewResources(version, resources.Clusters)
	out.Resources[types.Route] = NewResources(version, resources.Routes)
	out.Resources[types.Listener] = NewResources(version, resources.Listeners)
	out.Resources[types.Runtime] = NewResources(version, resources.Runtimes)
	out.Resources[types.Secret] = NewResources(version, resources.Secrets)
	out.Resources[types.ExtensionConfig] = NewResources(version, resources.ExtensionConfigs)

	return out
}

func NewSnapshotWithTtls(version string,
	endpoints []types.ResourceWithTTL,
	clusters []types.ResourceWithTTL,
	routes []types.ResourceWithTTL,
	listeners []types.ResourceWithTTL,
	runtimes []types.ResourceWithTTL,
	secrets []types.ResourceWithTTL) Snapshot {
	out := Snapshot{}
	out.Resources[types.Endpoint] = NewResourcesWithTtl(version, endpoints)
	out.Resources[types.Cluster] = NewResourcesWithTtl(version, clusters)
	out.Resources[types.Route] = NewResourcesWithTtl(version, routes)
	out.Resources[types.Listener] = NewResourcesWithTtl(version, listeners)
	out.Resources[types.Runtime] = NewResourcesWithTtl(version, runtimes)
	out.Resources[types.Secret] = NewResourcesWithTtl(version, secrets)
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

// GetResources selects snapshot resources by type, returning the map of resources.
func (s *Snapshot) GetResources(typeURL string) map[string]types.Resource {
	resources := s.GetResourcesAndTtl(typeURL)
	if resources == nil {
		return nil
	}

	withoutTTL := make(map[string]types.Resource, len(resources))

	for k, v := range resources {
		withoutTTL[k] = v.Resource
	}

	return withoutTTL
}

// GetResourcesAndTtl selects snapshot resources by type, returning the map of resources and the associated TTL.
func (s *Snapshot) GetResourcesAndTtl(typeURL string) map[string]types.ResourceWithTTL { // nolint:golint,revive
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

// GetVersionMap will return the internal version map of the currently applied snapshot.
func (s *Snapshot) GetVersionMap(typeUrl string) map[string]string {
	return s.VersionMap[typeUrl]
}

// ConstructVersionMap will construct a version map based on the current state of a snapshot
func (s *Snapshot) ConstructVersionMap() error {
	if s == nil {
		return fmt.Errorf("missing snapshot")
	}

	// The snapshot resources never change, so no need to ever rebuild.
	if s.VersionMap != nil {
		return nil
	}

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

// Respond to a delta watch with the provided snapshot value. If the response is nil, there has been no state change.
func (s *Snapshot) respondDelta(ctx context.Context, request *DeltaRequest, value chan DeltaResponse, state stream.StreamState, log log.Logger) (*RawDeltaResponse, error) {
	resources := &resourceContainer{
		resourceMap:   s.GetResources(request.TypeUrl),
		versionMap:    s.GetVersionMap(request.TypeUrl),
		systemVersion: s.GetVersion(request.TypeUrl),
	}
	resp := createDeltaResponse(ctx, request, state, resources, log)

	// Only send a response if there were changes
	// We want to respond immediately for the first wildcard request in a stream, even if the response is empty
	// otherwise, envoy won't complete initialization
	if len(resp.Resources) > 0 || len(resp.RemovedResources) > 0 || (state.IsWildcard() && state.IsFirst()) {
		if log != nil {
			log.Debugf("node: %s, sending delta response with resources: %v removed resources %v wildcard: %t",
				request.GetNode().GetId(), resp.Resources, resp.RemovedResources, state.IsWildcard())
		}
		select {
		case value <- resp:
			return resp, nil
		case <-ctx.Done():
			return resp, context.Canceled
		}
	}
	return nil, nil
}
