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

// Package cache defines a configuration cache for the server.
package cache

import (
	"context"
	"fmt"
	"sync/atomic"

	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/golang/protobuf/ptypes/any"
)

// Request is an alias for the discovery request type.
type Request = discovery.DiscoveryRequest

// DeltaRequest is an alias for the delta discovery request type.
type DeltaRequest = *discovery.DeltaDiscoveryRequest

// ConfigWatcher requests watches for configuration resources by a node, last
// applied version identifier, and resource names hint. The watch should send
// the responses when they are ready. The watch can be cancelled by the
// consumer, in effect terminating the watch for the request.
// ConfigWatcher implementation must be thread-safe.
type ConfigWatcher interface {
	// CreateWatch returns a new open watch from a non-empty request.
	// An individual consumer normally issues a single open watch by each type URL.
	//
	// Value channel produces requested resources, once they are available.  If
	// the channel is closed prior to cancellation of the watch, an unrecoverable
	// error has occurred in the producer, and the consumer should close the
	// corresponding stream.
	//
	// Cancel is an optional function to release resources in the producer. If
	// provided, the consumer may call this function multiple times.
	CreateWatch(Request) (value chan Response, cancel func())

	// CreateDeltaWatch returns a new open incremental xDS watch.
	//
	// Value channel produces requested resources, or spontaneous updates in accordance
	// with the incremental xDS specification. If the channel is closed
	// prior to cancellation of the watch, an unrecoverable error has occurred in the producer,
	// and the consumer should close the corresponding stream.
	//
	// Cancel is an optional function to release resources in the producer. If
	// provided, the consumer may call this function multiple times.
	CreateDeltaWatch(DeltaRequest, string) (value chan DeltaResponse, cancel func())
}

// ConfigFetcher fetches configuration resources from cache
type ConfigFetcher interface {
	// Fetch implements the polling method of the config cache using a non-empty request.
	Fetch(context.Context, *Request) (Response, error)
}

// Cache is a generic config cache with a watcher.
type Cache interface {
	ConfigWatcher
	ConfigFetcher
}

// Response is a wrapper around Envoy's DiscoveryResponse.
type Response interface {
	// Get the Constructed DiscoveryResponse
	GetDiscoveryResponse() (*discovery.DiscoveryResponse, error)

	// Get the original Request for the Response.
	GetRequest() *discovery.DiscoveryRequest

	// Get the version in the Response.
	GetVersion() (string, error)
}

// DeltaResponse is a wrapper around Envoy's DeltaDiscoveryResponse
type DeltaResponse interface {
	// Get the constructed DeltaDiscoveryResponse
	GetDeltaDiscoveryResponse() (*discovery.DeltaDiscoveryResponse, error)

	// Get te original Request for the Response.
	GetDeltaRequest() *discovery.DeltaDiscoveryRequest

	// Get the version in the Response.
	GetSystemVersion() (string, error)
}

// RawResponse is a pre-serialized xDS response containing the raw resources to
// be included in the final Discovery Response.
type RawResponse struct {
	// Request is the original request.
	Request *discovery.DiscoveryRequest

	// Version of the resources as tracked by the cache for the given type.
	// Proxy responds with this version as an acknowledgement.
	Version string

	// Resources to be included in the response.
	Resources []types.Resource

	// marshaledResponse holds an atomic reference to the serialized discovery response.
	marshaledResponse atomic.Value
}

// RawDeltaResponse is a pre-serialized xDS response that utilizes the delta discovery request/response objects.
type RawDeltaResponse struct {
	DeltaResponse

	// Request is the original request.
	DeltaRequest *discovery.DeltaDiscoveryRequest

	// System Version Info
	SystemVersion string

	// Resources to be included in the response.
	Resources []types.Resource

	// RemovedResources is a list of unsubscribed aliases to be included in the response
	RemovedResources []string

	// isResourceMarshaled indicates whether the resources have been marshaled.
	// This is internally maintained by go-control-plane to prevent future
	// duplication in marshaling efforts.
	isResourceMarshaled bool

	// Marshaled Resources to be included in the response.
	marshaledResponse *discovery.DeltaDiscoveryResponse
}

// PassthroughResponse is a pre constructed xDS response that need not go through marshalling transformations.
type PassthroughResponse struct {
	Response

	// Request is the original request.
	Request *discovery.DiscoveryRequest

	// The discovery response that needs to be sent as is, without any marshalling transformations.
	DiscoveryResponse *discovery.DiscoveryResponse
}

// DeltaPassthroughResponse is a pre constructed xDS response that need not go through marshalling transformations.
type DeltaPassthroughResponse struct {
	DeltaResponse

	// Request is the original request
	DeltaRequest discovery.DeltaDiscoveryRequest

	// This discovery response that needs to be sent as is, without any marshalling transformations
	DeltaDiscoveryResponse *discovery.DeltaDiscoveryResponse
}

// GetDiscoveryResponse performs the marshalling the first time its called and uses the cached response subsequently.
// This is necessary because the marshalled response does not change across the calls.
// This caching behavior is important in high throughput scenarios because grpc marshalling has a cost and it drives the cpu utilization under load.
func (r *RawResponse) GetDiscoveryResponse() (*discovery.DiscoveryResponse, error) {

	marshaledResponse := r.marshaledResponse.Load()

	if marshaledResponse == nil {

		marshaledResources := make([]*any.Any, len(r.Resources))

		for i, resource := range r.Resources {
			marshaledResource, err := MarshalResource(resource)
			if err != nil {
				return nil, err
			}
			marshaledResources[i] = &any.Any{
				TypeUrl: r.Request.TypeUrl,
				Value:   marshaledResource,
			}
		}

		marshaledResponse = &discovery.DiscoveryResponse{
			VersionInfo: r.Version,
			Resources:   marshaledResources,
			TypeUrl:     r.Request.TypeUrl,
		}

		r.marshaledResponse.Store(marshaledResponse)
	}

	return marshaledResponse.(*discovery.DiscoveryResponse), nil
}

// GetDeltaDiscoveryResponse performs the marshalling the first time its called and uses the cached response subsequently.
// This is necessary because the marshalled response does not change across the calls.
// This caching behavior is important in high throughput scenarios because grpc marshalling has a cost and it drives the cpu utilization under load.
func (r *RawDeltaResponse) GetDeltaDiscoveryResponse() (*discovery.DeltaDiscoveryResponse, error) {
	if r.isResourceMarshaled {
		return r.marshaledResponse, nil
	}

	marshaledResources := make([]*discovery.Resource, len(r.Resources))

	for i, resource := range r.Resources {
		marshaledResource, err := MarshalResource(resource)
		if err != nil {
			return nil, err
		}

		name := GetResourceName(resource)
		marshaledResources[i] = &discovery.Resource{
			Name:    name,
			Aliases: []string{name},
			Resource: &any.Any{
				TypeUrl: r.DeltaRequest.TypeUrl,
				Value:   marshaledResource,
			},
		}
	}

	r.marshaledResponse = &discovery.DeltaDiscoveryResponse{
		SystemVersionInfo: r.SystemVersion,
		Resources:         marshaledResources,
		RemovedResources:  r.RemovedResources,
		TypeUrl:           r.DeltaRequest.TypeUrl,
	}
	r.isResourceMarshaled = true

	return r.marshaledResponse, nil
}

// GetRequest returns the original Discovery Request.
func (r *RawResponse) GetRequest() *discovery.DiscoveryRequest {
	return r.Request
}

// GetDeltaRequest returns the original DeltaRequest
func (r *RawDeltaResponse) GetDeltaRequest() *discovery.DeltaDiscoveryRequest {
	return r.DeltaRequest
}

// GetVersion returns the response version.
func (r *RawResponse) GetVersion() (string, error) {
	return r.Version, nil
}

// GetSystemVersion returns the raw SystemVersion
func (r *RawDeltaResponse) GetSystemVersion() (string, error) {
	return r.SystemVersion, nil
}

// GetDiscoveryResponse returns the final passthrough Discovery Response.
func (r *PassthroughResponse) GetDiscoveryResponse() (*discovery.DiscoveryResponse, error) {
	return r.DiscoveryResponse, nil
}

// GetDeltaDiscoveryResponse returns the final passthrough Delta Discovery Response.
func (r *DeltaPassthroughResponse) GetDeltaDiscoveryResponse() (*discovery.DeltaDiscoveryResponse, error) {
	return r.DeltaDiscoveryResponse, nil
}

// GetRequest returns the original Discovery Request
func (r *PassthroughResponse) GetRequest() *discovery.DiscoveryRequest {
	return r.Request
}

// GetDeltaRequest returns the original Delta Discovery Request
func (r *DeltaPassthroughResponse) GetDeltaRequest() *discovery.DeltaDiscoveryRequest {
	return &r.DeltaRequest
}

// GetVersion returns the response version.
func (r *PassthroughResponse) GetVersion() (string, error) {
	if r.DiscoveryResponse != nil {
		return r.DiscoveryResponse.VersionInfo, nil
	}
	return "", fmt.Errorf("DiscoveryResponse is nil")
}

// GetSystemVersion returns the response version.
func (r *DeltaPassthroughResponse) GetSystemVersion() (string, error) {
	if r.DeltaDiscoveryResponse != nil {
		return r.DeltaDiscoveryResponse.SystemVersionInfo, nil
	}
	return "", fmt.Errorf("DeltaDiscoveryResponse is nil")
}
