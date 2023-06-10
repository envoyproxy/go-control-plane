// Copyright 2020 Envoyproxy Authors
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
	"sync/atomic"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type resourceWithNameAndVersion interface {
	getResource() types.Resource
	getName() string
	getVersion() string
}

func getResourceNames(resources []resourceWithNameAndVersion) []string {
	names := make([]string, 0, len(resources))
	for _, res := range resources {
		names = append(names, res.getName())
	}
	return names
}

var _ DeltaResponse = &rawDeltaResponse{}

// RawDeltaResponse is a pre-serialized xDS response that utilizes the delta discovery request/response objects.
type rawDeltaResponse struct {
	// Request is the original request.
	deltaRequest DeltaRequest

	// SystemVersionInfo holds the currently applied response system version and should be used for debugging purposes only.
	systemVersionInfo string

	// Resources to be included in the response.
	resources []resourceWithNameAndVersion

	// RemovedResources is a list of resource aliases which should be dropped by the consuming client.
	removedResources []string

	// Context provided at the time of response creation. This allows associating additional
	// information with a generated response.
	ctx context.Context

	// Marshaled Resources to be included in the response.
	marshaledResponse atomic.Value
}

// ToDo: merge with versionedResource
type TestNamedResource struct {
	Name     string
	Version  string
	Resource types.Resource
}

func (r TestNamedResource) getResource() types.Resource { return r.Resource }
func (r TestNamedResource) getName() string             { return r.Name }
func (r TestNamedResource) getVersion() string          { return r.Version }

func NewTestRawDeltaResponse(request DeltaRequest, version string, resources []TestNamedResource, removedResources []string) *rawDeltaResponse {
	resourceInterfaces := make([]resourceWithNameAndVersion, 0, len(resources))
	for _, resource := range resources {
		resourceInterfaces = append(resourceInterfaces, resource)
	}
	return &rawDeltaResponse{
		deltaRequest:      request,
		systemVersionInfo: version,
		resources:         resourceInterfaces,
		removedResources:  removedResources,
	}
}

// GetDeltaDiscoveryResponse performs the marshaling the first time its called and uses the cached response subsequently.
// We can do this because the marshaled response does not change across the calls.
// This caching behavior is important in high throughput scenarios because grpc marshaling has a cost and it drives the cpu utilization under load.
func (r *rawDeltaResponse) GetDeltaDiscoveryResponse() (*discovery.DeltaDiscoveryResponse, error) {
	marshaledResponse := r.marshaledResponse.Load()

	if marshaledResponse == nil {
		marshaledResources := make([]*discovery.Resource, len(r.resources))

		for i, resource := range r.resources {
			marshaledResource, err := MarshalResource(resource.getResource())
			if err != nil {
				return nil, err
			}
			marshaledResources[i] = &discovery.Resource{
				Name: resource.getName(),
				Resource: &anypb.Any{
					TypeUrl: r.deltaRequest.GetTypeUrl(),
					Value:   marshaledResource,
				},
				Version: resource.getVersion(),
			}
		}

		marshaledResponse = &discovery.DeltaDiscoveryResponse{
			Resources:         marshaledResources,
			RemovedResources:  r.removedResources,
			TypeUrl:           r.deltaRequest.GetTypeUrl(),
			SystemVersionInfo: r.systemVersionInfo,
		}
		r.marshaledResponse.Store(marshaledResponse)
	}

	return marshaledResponse.(*discovery.DeltaDiscoveryResponse), nil
}

// Get the list of resources included in the response.
// If the resource was set in Resources, its version is provided
// If the version was set in RemovedResources, its version is ""
func (r *rawDeltaResponse) GetReturnedResources() map[string]string {
	resources := make(map[string]string, len(r.resources)+len(r.removedResources))
	for _, resource := range r.resources {
		resources[resource.getName()] = resource.getVersion()
	}
	for _, resourceName := range r.removedResources {
		resources[resourceName] = ""
	}
	return resources
}

func (r *rawDeltaResponse) GetDeltaRequest() DeltaRequest {
	return r.deltaRequest
}

// GetDeltaDiscoveryRequest returns the original Discovery DeltaRequest if available
func (r *rawDeltaResponse) GetDeltaDiscoveryRequest() (*discovery.DeltaDiscoveryRequest, error) {
	if r.deltaRequest == nil {
		return nil, errors.New("no request was provided")
	} else if req, ok := r.deltaRequest.(*discovery.DeltaDiscoveryRequest); ok {
		return req, nil
	} else {
		return nil, errors.New("provided request is not a DeltaDiscoveryRequest")
	}
}

// GetSystemVersion returns the raw SystemVersion
func (r *rawDeltaResponse) GetSystemVersion() (string, error) {
	return r.systemVersionInfo, nil
}

func (r *rawDeltaResponse) GetContext() context.Context {
	return r.ctx
}

var _ DeltaResponse = &deltaPassthroughResponse{}

// DeltaPassthroughResponse is a pre constructed xDS response that need not go through marshaling transformations.
type deltaPassthroughResponse struct {
	// DeltaDiscoveryRequest is the original request.
	deltaDiscoveryRequest *discovery.DeltaDiscoveryRequest

	// This discovery response that needs to be sent as is, without any marshaling transformations
	deltaDiscoveryResponse *discovery.DeltaDiscoveryResponse

	ctx context.Context
}

var deltaResourceTypeURL = "type.googleapis.com/" + string(proto.MessageName(&discovery.Resource{}))

func (r *deltaPassthroughResponse) GetDeltaRequest() DeltaRequest {
	return r.deltaDiscoveryRequest
}

// GetDeltaDiscoveryResponse returns the final passthrough Delta Discovery Response.
func (r *deltaPassthroughResponse) GetDeltaDiscoveryResponse() (*discovery.DeltaDiscoveryResponse, error) {
	return r.deltaDiscoveryResponse, nil
}

// GetDeltaRequest returns the original Request
func (r *deltaPassthroughResponse) GetDeltaDiscoveryRequest() (*discovery.DeltaDiscoveryRequest, error) {
	return r.deltaDiscoveryRequest, nil
}

// GetSystemVersion returns the response version.
func (r *deltaPassthroughResponse) GetSystemVersion() (string, error) {
	if r.deltaDiscoveryResponse != nil {
		return r.deltaDiscoveryResponse.SystemVersionInfo, nil
	}
	return "", fmt.Errorf("DeltaDiscoveryResponse is nil")
}

func (r *deltaPassthroughResponse) GetContext() context.Context {
	return r.ctx
}

// Get the list of resources included in the response.
// If the resource was set in Resources, its version is provided
// If the version was set in RemovedResources, its version is ""
func (r *deltaPassthroughResponse) GetReturnedResources() map[string]string {
	resources := make(map[string]string, len(r.deltaDiscoveryResponse.Resources)+len(r.deltaDiscoveryResponse.RemovedResources))
	for _, resource := range r.deltaDiscoveryResponse.Resources {
		resources[resource.Name] = resource.Version
	}
	for _, resourceName := range r.deltaDiscoveryResponse.RemovedResources {
		resources[resourceName] = ""
	}
	return resources
}

// groups together resource-related arguments for the createDeltaResponse function
type resourceContainer struct {
	resourceMap   map[string]types.Resource
	versionMap    map[string]string
	systemVersion string
}

type versionedResource struct {
	resource types.Resource
	name     string
	version  string
}

func (r versionedResource) getResource() types.Resource {
	return r.resource
}

func (r versionedResource) getName() string {
	return r.name
}

func (r versionedResource) getVersion() string {
	return r.version
}

func createDeltaResponse(ctx context.Context, req DeltaRequest, state SubscriptionState, resources resourceContainer) *rawDeltaResponse {
	// variables to build our response with
	var filtered []resourceWithNameAndVersion
	var toRemove []string

	// If we are handling a wildcard request, we want to respond with all resources
	switch {
	case state.IsWildcard():
		if len(state.GetKnownResources()) == 0 {
			filtered = make([]resourceWithNameAndVersion, 0, len(resources.resourceMap))
		}
		for name, r := range resources.resourceMap {
			// Since we've already precomputed the version hashes of the new snapshot,
			// we can just set it here to be used for comparison later
			version := resources.versionMap[name]
			prevVersion, found := state.GetKnownResources()[name]
			if !found || (prevVersion != version) {
				filtered = append(filtered, versionedResource{
					name:     name,
					version:  version,
					resource: r,
				})
			}
		}

		// Compute resources for removal
		for name, version := range state.GetKnownResources() {
			if version == "" {
				// The resource version can be set to "" to indicate that we already notified the deletion
				continue
			}
			if _, ok := resources.resourceMap[name]; !ok {
				toRemove = append(toRemove, name)
			}
		}
	default:
		// state.GetResourceVersions() may include resources no longer subscribed
		// In the current code this gets silently cleaned when updating the version map
		for name := range state.GetSubscribedResources() {
			prevVersion, found := state.GetKnownResources()[name]
			if r, ok := resources.resourceMap[name]; ok {
				nextVersion := resources.versionMap[name]
				if prevVersion != nextVersion {
					filtered = append(filtered, versionedResource{
						name:     name,
						version:  nextVersion,
						resource: r,
					})
				}
			} else if found && prevVersion != "" {
				toRemove = append(toRemove, name)
			}
		}
	}

	return &rawDeltaResponse{
		deltaRequest:      req,
		resources:         filtered,
		removedResources:  toRemove,
		systemVersionInfo: resources.systemVersion,
		ctx:               ctx,
	}
}
