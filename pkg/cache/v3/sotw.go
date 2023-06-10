package cache

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
)

type resourceWithName interface {
	// getResource is returning the resource which can be directly serialized for the response
	getResource() types.Resource
	getName() string
}

type resourceWithTTLAndName interface {
	resourceWithName
	getTTL() *time.Duration
}

type ttlResource struct {
	resource types.Resource
	name     string
	ttl      *time.Duration
}

func (r ttlResource) getResource() types.Resource {
	return r.resource
}

func (r ttlResource) getName() string {
	return r.name
}

func (r ttlResource) getTTL() *time.Duration {
	return r.ttl
}

var _ Response = &rawResponse{}

// RawResponse is a pre-serialized xDS response containing the raw resources to
// be included in the final Discovery Response.
type rawResponse struct {
	// Request is the original request.
	request Request

	// Version of the resources as tracked by the cache for the given type.
	// Proxy responds with this version as an acknowledgement.
	version string

	// Resources to be included in the response.
	resources []resourceWithTTLAndName

	// Whether this is a heartbeat response. For xDS versions that support TTL, this
	// will be converted into a response that doesn't contain the actual resource protobuf.
	// This allows for more lightweight updates that server only to update the TTL timer.
	heartbeat bool

	// Context provided at the time of response creation. This allows associating additional
	// information with a generated response.
	ctx context.Context

	// marshaledResponse holds an atomic reference to the serialized discovery response.
	marshaledResponse atomic.Value
}

func NewTestRawResponse(request Request, version string, resources []types.ResourceWithTTL) *rawResponse {
	var namedResources []resourceWithTTLAndName
	for _, res := range resources {
		namedResources = append(namedResources, ttlResource{
			name:     GetResourceName(res.Resource),
			resource: res.Resource,
			ttl:      res.TTL,
		})
	}
	return &rawResponse{
		request:   request,
		version:   version,
		resources: namedResources,
	}
}

// GetDiscoveryResponse performs the marshaling the first time its called and uses the cached response subsequently.
// This is necessary because the marshaled response does not change across the calls.
// This caching behavior is important in high throughput scenarios because grpc marshaling has a cost and it drives the cpu utilization under load.
func (r *rawResponse) GetDiscoveryResponse() (*discovery.DiscoveryResponse, error) {
	marshaledResponse := r.marshaledResponse.Load()

	if marshaledResponse == nil {
		marshaledResources := make([]*anypb.Any, len(r.resources))

		for i, resource := range r.resources {
			maybeTtldResource, resourceType, err := r.maybeCreateTTLResource(resource)
			if err != nil {
				return nil, err
			}
			marshaledResource, err := MarshalResource(maybeTtldResource.getResource())
			if err != nil {
				return nil, err
			}
			marshaledResources[i] = &anypb.Any{
				TypeUrl: resourceType,
				Value:   marshaledResource,
			}
		}

		marshaledResponse = &discovery.DiscoveryResponse{
			VersionInfo: r.version,
			Resources:   marshaledResources,
			TypeUrl:     r.request.GetTypeUrl(),
		}

		r.marshaledResponse.Store(marshaledResponse)
	}

	return marshaledResponse.(*discovery.DiscoveryResponse), nil
}

func (r *rawResponse) GetReturnedResources() []string {
	names := make([]string, 0, len(r.resources))
	for _, res := range r.resources {
		names = append(names, res.getName())
	}
	return names
}

func (r *rawResponse) GetRequest() Request {
	return r.request
}

// GetDiscoveryRequest returns the original Discovery Request if available
func (r *rawResponse) GetDiscoveryRequest() (*discovery.DiscoveryRequest, error) {
	if r.request == nil {
		return nil, errors.New("no request was provided")
	} else if req, ok := r.request.(*discovery.DiscoveryRequest); ok {
		return req, nil
	} else {
		return nil, errors.New("provided request is not a DiscoveryRequest")
	}
}

func (r *rawResponse) GetContext() context.Context {
	return r.ctx
}

// GetVersion returns the response version.
func (r *rawResponse) GetVersion() (string, error) {
	return r.version, nil
}

type namedResource struct {
	types.Resource
	name string
}

func (r namedResource) getResource() types.Resource {
	return r.Resource
}

func (r namedResource) getName() string {
	return r.name
}

func (r *rawResponse) maybeCreateTTLResource(resource resourceWithTTLAndName) (resourceWithName, string, error) {
	if resource.getTTL() != nil {
		wrappedResource := &discovery.Resource{
			Name: resource.getName(),
			Ttl:  durationpb.New(*resource.getTTL()),
		}

		if !r.heartbeat {
			rsrc, err := anypb.New(resource.getResource())
			if err != nil {
				return namedResource{}, "", err
			}
			rsrc.TypeUrl = r.request.GetTypeUrl()
			wrappedResource.Resource = rsrc
		}

		return namedResource{
			Resource: wrappedResource,
			name:     resource.getName(),
		}, deltaResourceTypeURL, nil
	}

	return resource, r.request.GetTypeUrl(), nil
}

var _ Response = &passthroughResponse{}

// passthroughResponse is a pre constructed xDS response that need not go through marshaling transformations.
type passthroughResponse struct {
	// discoveryRequest is the original request.
	discoveryRequest *discovery.DiscoveryRequest

	// The discovery response that needs to be sent as is, without any marshaling transformations.
	discoveryResponse *discovery.DiscoveryResponse

	ctx context.Context
}

func (r *passthroughResponse) GetDiscoveryResponse() (*discovery.DiscoveryResponse, error) {
	return r.discoveryResponse, nil
}

func (r *passthroughResponse) GetReturnedResources() []string {
	if r.discoveryResponse == nil {
		return nil
	}
	names := make([]string, 0, len(r.discoveryResponse.Resources))
	for _, res := range r.discoveryResponse.Resources {
		names = append(names, GetResourceName(res))
	}
	return names
}

func (r *passthroughResponse) GetRequest() Request {
	return r.discoveryRequest
}

func (r *passthroughResponse) GetDiscoveryRequest() (*discovery.DiscoveryRequest, error) {
	return r.discoveryRequest, nil
}

func (r *passthroughResponse) GetContext() context.Context {
	return r.ctx
}

// GetVersion returns the response version.
func (r *passthroughResponse) GetVersion() (string, error) {
	if r.discoveryResponse != nil {
		return r.discoveryResponse.VersionInfo, nil
	}
	return "", fmt.Errorf("DiscoveryResponse is nil")
}
