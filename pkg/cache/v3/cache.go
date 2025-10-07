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
	"errors"
	"sync/atomic"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
)

// Request is an alias for the discovery request type.
type Request = discovery.DiscoveryRequest

// DeltaRequest is an alias for the delta discovery request type.
type DeltaRequest = discovery.DeltaDiscoveryRequest

// Subscription stores the server view of the client state for a given resource type.
// This allows proper implementation of stateful aspects of the protocol (e.g. returning only some updated resources).
// Though the methods may return mutable parts of the state for performance reasons,
// the cache is expected to consider this state as immutable and thread safe between a watch creation and its cancellation.
type Subscription interface {
	// ReturnedResources returns a list of resources that were sent to the client and their associated versions.
	// The versions are:
	//  - delta protocol: version of the specific resource set in the response.
	//  - sotw protocol: version of the global response when the resource was last sent.
	ReturnedResources() map[string]string

	// SubscribedResources returns the list of resources currently subscribed to by the client for the type.
	// For delta it keeps track of subscription updates across requests
	// For sotw it is a normalized view of the last request resources
	SubscribedResources() map[string]struct{}

	// IsWildcard returns whether the client has a wildcard watch.
	// This considers subtleties related to the current migration of wildcard definitions within the protocol.
	// More details on the behavior of wildcard are present at https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#how-the-client-specifies-what-resources-to-return
	IsWildcard() bool

	// WatchesResources returns whether at least one of the resources provided is currently being watched by the subscription.
	// It is currently only applicable to delta-xds.
	// If the request is wildcard, it will always return true,
	// otherwise it will compare the provided resources to the list of resources currently subscribed
	WatchesResources(resourceNames map[string]struct{}) bool
}

// ConfigWatcher requests watches for configuration resources by a node, last
// applied version identifier, and resource names hint. The watch should send
// the responses when they are ready. The watch can be canceled by the
// consumer, in effect terminating the watch for the request.
// ConfigWatcher implementation must be thread-safe.
type ConfigWatcher interface {
	// CreateWatch returns a new open watch from a non-empty request.
	// This is the entrypoint to propagate configuration changes the
	// provided Response channel. State from the gRPC server is utilized
	// to make sure consuming cache implementations can see what the server has sent to clients.
	//
	// An individual consumer normally issues a single open watch by each type URL.
	//
	// The provided channel produces requested resources as responses, once they are available.
	//
	// Cancel is an optional function to release resources in the producer. If
	// provided, the consumer may call this function multiple times.
	CreateWatch(*Request, Subscription, chan Response) (cancel func(), err error)

	// CreateDeltaWatch returns a new open incremental xDS watch.
	// This is the entrypoint to propagate configuration changes the
	// provided DeltaResponse channel. State from the gRPC server is utilized
	// to make sure consuming cache implementations can see what the server has sent to clients.
	//
	// The provided channel produces requested resources as responses, or spontaneous updates in accordance
	// with the incremental xDS specification.
	//
	// Cancel is an optional function to release resources in the producer. If
	// provided, the consumer may call this function multiple times.
	CreateDeltaWatch(*DeltaRequest, Subscription, chan DeltaResponse) (cancel func(), err error)
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

	// GetReturnedResources returns the map of resources and their versions returned in the subscription.
	// It may include more resources than directly set in the response to consider the full state of the client.
	// The caller is expected to provide this unchanged to the next call to CreateWatch as part of the subscription.
	GetReturnedResources() map[string]string

	// Get the context provided during response creation.
	GetContext() context.Context
}

// DeltaResponse is a wrapper around Envoy's DeltaDiscoveryResponse
type DeltaResponse interface {
	// Get the constructed DeltaDiscoveryResponse
	GetDeltaDiscoveryResponse() (*discovery.DeltaDiscoveryResponse, error)

	// Get the request that created the watch that we're now responding to. This is provided to allow the caller to correlate the
	// response with a request. Generally this will be the latest request seen on the stream for the specific type.
	GetDeltaRequest() *discovery.DeltaDiscoveryRequest

	// Get the version in the DeltaResponse. This field is generally used for debugging purposes as noted by the Envoy documentation.
	GetSystemVersion() (string, error)

	// Get the version map of the internal cache.
	// The version map consists of updated version mappings after this response is applied
	GetNextVersionMap() map[string]string

	// Get the context provided during response creation
	GetContext() context.Context
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
	Resources []types.ResourceWithTTL

	// ReturnedResources tracks the resources returned for the subscription and the version when it was last returned,
	// including previously returned ones when using non-full state resources.
	// It allows the cache to know what the client knows. The server will transparently forward this
	// across requests, and the cache is responsible for its interpretation.
	ReturnedResources map[string]string

	// Whether this is a heartbeat response. For xDS versions that support TTL, this
	// will be converted into a response that doesn't contain the actual resource protobuf.
	// This allows for more lightweight updates that server only to update the TTL timer.
	Heartbeat bool

	// Context provided at the time of response creation. This allows associating additional
	// information with a generated response.
	Ctx context.Context

	// marshaledResponse holds an atomic reference to the serialized discovery response.
	marshaledResponse atomic.Value
}

// RawDeltaResponse is a pre-serialized xDS response that utilizes the delta discovery request/response objects.
type RawDeltaResponse struct {
	// Request is the latest delta request on the stream.
	DeltaRequest *discovery.DeltaDiscoveryRequest

	// SystemVersionInfo holds the currently applied response system version and should be used for debugging purposes only.
	SystemVersionInfo string

	// Resources to be included in the response.
	Resources []types.Resource

	// RemovedResources is a list of resource aliases which should be dropped by the consuming client.
	RemovedResources []string

	// NextVersionMap consists of updated version mappings after this response is applied
	NextVersionMap map[string]string

	// Context provided at the time of response creation. This allows associating additional
	// information with a generated response.
	Ctx context.Context

	// Marshaled Resources to be included in the response.
	marshaledResponse atomic.Value
}

var (
	_ Response      = &RawResponse{}
	_ DeltaResponse = &RawDeltaResponse{}
)

// PassthroughResponse is a pre constructed xDS response that need not go through marshaling transformations.
type PassthroughResponse struct {
	// Request is the original request.
	Request *discovery.DiscoveryRequest

	// The discovery response that needs to be sent as is, without any marshaling transformations.
	DiscoveryResponse *discovery.DiscoveryResponse

	ctx context.Context

	// ReturnedResources tracks the resources returned for the subscription and the version when it was last returned,
	// including previously returned ones when using non-full state resources.
	// It allows the cache to know what the client knows. The server will transparently forward this
	// across requests, and the cache is responsible for its interpretation.
	ReturnedResources map[string]string
}

// DeltaPassthroughResponse is a pre constructed xDS response that need not go through marshaling transformations.
type DeltaPassthroughResponse struct {
	// Request is the latest delta request on the stream
	DeltaRequest *discovery.DeltaDiscoveryRequest

	// NextVersionMap consists of updated version mappings after this response is applied
	NextVersionMap map[string]string

	// This discovery response that needs to be sent as is, without any marshaling transformations
	DeltaDiscoveryResponse *discovery.DeltaDiscoveryResponse

	ctx context.Context
}

var (
	_ Response      = &PassthroughResponse{}
	_ DeltaResponse = &DeltaPassthroughResponse{}
)

// GetDiscoveryResponse performs the marshaling the first time its called and uses the cached response subsequently.
// This is necessary because the marshaled response does not change across the calls.
// This caching behavior is important in high throughput scenarios because grpc marshaling has a cost and it drives the cpu utilization under load.
func (r *RawResponse) GetDiscoveryResponse() (*discovery.DiscoveryResponse, error) {
	marshaledResponse := r.marshaledResponse.Load()

	if marshaledResponse == nil {
		marshaledResources := make([]*anypb.Any, len(r.Resources))

		for i, resource := range r.Resources {
			maybeTtldResource, resourceType, err := r.maybeCreateTTLResource(resource)
			if err != nil {
				return nil, err
			}
			marshaledResource, err := MarshalResource(maybeTtldResource)
			if err != nil {
				return nil, err
			}
			marshaledResources[i] = &anypb.Any{
				TypeUrl: resourceType,
				Value:   marshaledResource,
			}
		}

		marshaledResponse = &discovery.DiscoveryResponse{
			VersionInfo: r.Version,
			Resources:   marshaledResources,
			TypeUrl:     r.GetRequest().GetTypeUrl(),
		}

		r.marshaledResponse.Store(marshaledResponse)
	}

	return marshaledResponse.(*discovery.DiscoveryResponse), nil
}

func (r *RawResponse) GetReturnedResources() map[string]string {
	return r.ReturnedResources
}

// GetDeltaDiscoveryResponse performs the marshaling the first time its called and uses the cached response subsequently.
// We can do this because the marshaled response does not change across the calls.
// This caching behavior is important in high throughput scenarios because grpc marshaling has a cost and it drives the cpu utilization under load.
func (r *RawDeltaResponse) GetDeltaDiscoveryResponse() (*discovery.DeltaDiscoveryResponse, error) {
	marshaledResponse := r.marshaledResponse.Load()

	if marshaledResponse == nil {
		marshaledResources := make([]*discovery.Resource, len(r.Resources))

		for i, resource := range r.Resources {
			name := GetResourceName(resource)
			marshaledResource, err := MarshalResource(resource)
			if err != nil {
				return nil, err
			}
			version := HashResource(marshaledResource)
			if version == "" {
				return nil, errors.New("failed to create a resource hash")
			}
			marshaledResources[i] = &discovery.Resource{
				Name: name,
				Resource: &anypb.Any{
					TypeUrl: r.GetDeltaRequest().GetTypeUrl(),
					Value:   marshaledResource,
				},
				Version: version,
			}
		}

		marshaledResponse = &discovery.DeltaDiscoveryResponse{
			Resources:         marshaledResources,
			RemovedResources:  r.RemovedResources,
			TypeUrl:           r.GetDeltaRequest().GetTypeUrl(),
			SystemVersionInfo: r.SystemVersionInfo,
		}
		r.marshaledResponse.Store(marshaledResponse)
	}

	return marshaledResponse.(*discovery.DeltaDiscoveryResponse), nil
}

// GetRequest returns the original Discovery Request.
func (r *RawResponse) GetRequest() *discovery.DiscoveryRequest {
	return r.Request
}

func (r *RawResponse) GetContext() context.Context {
	return r.Ctx
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
	return r.SystemVersionInfo, nil
}

// NextVersionMap returns the version map which consists of updated version mappings after this response is applied
func (r *RawDeltaResponse) GetNextVersionMap() map[string]string {
	return r.NextVersionMap
}

func (r *RawDeltaResponse) GetContext() context.Context {
	return r.Ctx
}

var deltaResourceTypeURL = "type.googleapis.com/" + string(proto.MessageName(&discovery.Resource{}))

func (r *RawResponse) maybeCreateTTLResource(resource types.ResourceWithTTL) (types.Resource, string, error) {
	if resource.TTL != nil {
		wrappedResource := &discovery.Resource{
			Name: GetResourceName(resource.Resource),
			Ttl:  durationpb.New(*resource.TTL),
		}

		if !r.Heartbeat {
			rsrc, err := anypb.New(resource.Resource)
			if err != nil {
				return nil, "", err
			}
			rsrc.TypeUrl = r.GetRequest().GetTypeUrl()
			wrappedResource.Resource = rsrc
		}

		return wrappedResource, deltaResourceTypeURL, nil
	}

	return resource.Resource, r.GetRequest().GetTypeUrl(), nil
}

// GetDiscoveryResponse returns the final passthrough Discovery Response.
func (r *PassthroughResponse) GetDiscoveryResponse() (*discovery.DiscoveryResponse, error) {
	return r.DiscoveryResponse, nil
}

func (r *PassthroughResponse) GetReturnedResources() map[string]string {
	return r.ReturnedResources
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
	return r.DeltaRequest
}

// GetVersion returns the response version.
func (r *PassthroughResponse) GetVersion() (string, error) {
	discoveryResponse, _ := r.GetDiscoveryResponse()
	if discoveryResponse != nil {
		return discoveryResponse.GetVersionInfo(), nil
	}
	return "", errors.New("DiscoveryResponse is nil")
}

func (r *PassthroughResponse) GetContext() context.Context {
	return r.ctx
}

// GetSystemVersion returns the response version.
func (r *DeltaPassthroughResponse) GetSystemVersion() (string, error) {
	deltaDiscoveryResponse, _ := r.GetDeltaDiscoveryResponse()
	if deltaDiscoveryResponse != nil {
		return deltaDiscoveryResponse.GetSystemVersionInfo(), nil
	}
	return "", errors.New("DeltaDiscoveryResponse is nil")
}

// NextVersionMap returns the version map from a DeltaPassthroughResponse
func (r *DeltaPassthroughResponse) GetNextVersionMap() map[string]string {
	return r.NextVersionMap
}

func (r *DeltaPassthroughResponse) GetContext() context.Context {
	return r.ctx
}
