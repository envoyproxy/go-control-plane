package cache_test

import (
	"testing"

	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v2"
	ttl_helper "github.com/envoyproxy/go-control-plane/pkg/ttl/v2"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
)

const (
	resourceName = "route1"
)

func TestResponseGetDiscoveryResponse(t *testing.T) {
	routes := []types.ResourceWithTtl{{Resource: &route.RouteConfiguration{Name: resourceName}}}
	resp := cache.RawResponse{
		Request:   &discovery.DiscoveryRequest{TypeUrl: resource.RouteType},
		Version:   "v",
		Resources: routes,
	}

	discoveryResponse, err := resp.GetDiscoveryResponse()
	assert.Nil(t, err)
	assert.Equal(t, discoveryResponse.VersionInfo, resp.Version)
	assert.Equal(t, len(discoveryResponse.Resources), 1)

	cachedResponse, err := resp.GetDiscoveryResponse()
	assert.Nil(t, err)
	assert.Same(t, discoveryResponse, cachedResponse)

	r := &route.RouteConfiguration{}
	err = ptypes.UnmarshalAny(discoveryResponse.Resources[0], r)
	assert.Nil(t, err)
	assert.Equal(t, r.Name, resourceName)
}

func TestPassthroughResponseGetDiscoveryResponse(t *testing.T) {
	routes := []types.Resource{&route.RouteConfiguration{Name: resourceName}}
	rsrc, err := ptypes.MarshalAny(routes[0])
	assert.Nil(t, err)
	dr := &discovery.DiscoveryResponse{
		TypeUrl:     resource.RouteType,
		Resources:   []*any.Any{rsrc},
		VersionInfo: "v",
	}
	resp := cache.PassthroughResponse{
		Request:           &discovery.DiscoveryRequest{TypeUrl: resource.RouteType},
		DiscoveryResponse: dr,
	}

	discoveryResponse, err := resp.GetDiscoveryResponse()
	assert.Nil(t, err)
	assert.Equal(t, discoveryResponse.VersionInfo, resp.DiscoveryResponse.VersionInfo)
	assert.Equal(t, len(discoveryResponse.Resources), 1)

	r := &route.RouteConfiguration{}
	err = ptypes.UnmarshalAny(discoveryResponse.Resources[0], r)
	assert.Nil(t, err)
	assert.Equal(t, r.Name, resourceName)
	assert.Equal(t, discoveryResponse, dr)
}

func TestResponseGetDeltaDiscoveryResponse(t *testing.T) {
	r := []types.Resource{&route.RouteConfiguration{Name: resourceName}}

	resp := cache.RawDeltaResponse{
		DeltaRequest: &discovery.DeltaDiscoveryRequest{TypeUrl: resource.RouteType},
		Resources:    r,
	}

	discoveryResponse, err := resp.GetDeltaDiscoveryResponse()
	assert.Nil(t, err)
	assert.Equal(t, discoveryResponse.SystemVersionInfo, resp.SystemVersionInfo)
	assert.Equal(t, len(discoveryResponse.Resources), 1)

	cachedResponse, err := resp.GetDeltaDiscoveryResponse()
	assert.Nil(t, err)
	assert.Same(t, discoveryResponse, cachedResponse)

	route := &route.RouteConfiguration{}
	err = ptypes.UnmarshalAny(discoveryResponse.Resources[0].GetResource(), route)
	assert.Nil(t, err)
	assert.Equal(t, route.Name, resourceName)
}

func TestDeltaPassthroughResponseGetDiscoveryResponse(t *testing.T) {
	routes := []types.Resource{&route.RouteConfiguration{Name: resourceName}}
	rsrc, err := ptypes.MarshalAny(routes[0])
	assert.Nil(t, err)
	dr := &discovery.DeltaDiscoveryResponse{
		TypeUrl: resource.RouteType,
		Resources: []*discovery.Resource{{
			Name:     resourceName,
			Aliases:  []string{resourceName},
			Resource: rsrc,
		}},
		SystemVersionInfo: "v0",
	}

	resp := cache.DeltaPassthroughResponse{
		DeltaRequest:           &discovery.DeltaDiscoveryRequest{TypeUrl: resource.RouteType},
		DeltaDiscoveryResponse: dr,
	}

	discoveryResponse, err := resp.GetDeltaDiscoveryResponse()
	assert.Nil(t, err)
	assert.Equal(t, discoveryResponse.SystemVersionInfo, resp.DeltaDiscoveryResponse.SystemVersionInfo)
	assert.Equal(t, len(discoveryResponse.Resources), 1)

	r := &route.RouteConfiguration{}
	err = ptypes.UnmarshalAny(discoveryResponse.Resources[0].GetResource(), r)
	assert.Nil(t, err)
	assert.Equal(t, r.Name, resourceName)
	assert.Equal(t, discoveryResponse, dr)
}
func TestHeartbeatResponseGetDiscoveryResponse(t *testing.T) {
	routes := []types.ResourceWithTtl{{Resource: &route.RouteConfiguration{Name: resourceName}}}
	resp := cache.RawResponse{
		Request:   &discovery.DiscoveryRequest{TypeUrl: resource.RouteType},
		Version:   "v",
		Resources: routes,
		Heartbeat: true,
	}

	discoveryResponse, err := resp.GetDiscoveryResponse()
	assert.Nil(t, err)
	assert.Equal(t, discoveryResponse.VersionInfo, resp.Version)
	assert.Equal(t, len(discoveryResponse.Resources), 1)
	assert.True(t, ttl_helper.IsTTLResource(discoveryResponse.Resources[0]))

	cachedResponse, err := resp.GetDiscoveryResponse()
	assert.Nil(t, err)
	assert.Same(t, discoveryResponse, cachedResponse)

	r := &route.RouteConfiguration{}
	err = ptypes.UnmarshalAny(discoveryResponse.Resources[0], r)
	assert.Nil(t, err)
	assert.Equal(t, r.Name, resourceName)
}
