package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

const (
	resourceName = "route1"
)

func TestResponseGetDiscoveryResponse(t *testing.T) {
	routes := []resourceWithTTLAndName{ttlResource{resource: &route.RouteConfiguration{Name: resourceName}, name: resourceName}}
	resp := rawResponse{
		request:   &discovery.DiscoveryRequest{TypeUrl: resource.RouteType},
		version:   "v",
		resources: routes,
	}

	discoveryResponse, err := resp.GetDiscoveryResponse()
	assert.Nil(t, err)
	assert.Equal(t, discoveryResponse.VersionInfo, resp.version)
	assert.Equal(t, len(discoveryResponse.Resources), 1)

	cachedResponse, err := resp.GetDiscoveryResponse()
	assert.Nil(t, err)
	assert.Same(t, discoveryResponse, cachedResponse)

	r := &route.RouteConfiguration{}
	err = anypb.UnmarshalTo(discoveryResponse.Resources[0], r, proto.UnmarshalOptions{})
	assert.Nil(t, err)
	assert.Equal(t, r.Name, resourceName)
}

func TestPassthroughResponseGetDiscoveryResponse(t *testing.T) {
	routes := []types.Resource{&route.RouteConfiguration{Name: resourceName}}
	rsrc, err := anypb.New(routes[0])
	assert.Nil(t, err)
	dr := &discovery.DiscoveryResponse{
		TypeUrl:     resource.RouteType,
		Resources:   []*anypb.Any{rsrc},
		VersionInfo: "v",
	}
	resp := passthroughResponse{
		discoveryRequest:  &discovery.DiscoveryRequest{TypeUrl: resource.RouteType},
		discoveryResponse: dr,
	}

	discoveryResponse, err := resp.GetDiscoveryResponse()
	assert.Nil(t, err)
	assert.Equal(t, discoveryResponse.VersionInfo, resp.discoveryResponse.VersionInfo)
	assert.Equal(t, len(discoveryResponse.Resources), 1)

	r := &route.RouteConfiguration{}
	err = anypb.UnmarshalTo(discoveryResponse.Resources[0], r, proto.UnmarshalOptions{})
	assert.Nil(t, err)
	assert.Equal(t, r.Name, resourceName)
	assert.Equal(t, discoveryResponse, dr)
}

func TestHeartbeatResponseGetDiscoveryResponse(t *testing.T) {
	routes := []resourceWithTTLAndName{ttlResource{resource: &route.RouteConfiguration{Name: resourceName}, name: resourceName}}
	resp := rawResponse{
		request:   &discovery.DiscoveryRequest{TypeUrl: resource.RouteType},
		version:   "v",
		resources: routes,
		heartbeat: true,
	}

	discoveryResponse, err := resp.GetDiscoveryResponse()
	assert.Nil(t, err)
	assert.Equal(t, discoveryResponse.VersionInfo, resp.version)
	assert.Equal(t, len(discoveryResponse.Resources), 1)
	assert.False(t, isTTLResource(discoveryResponse.Resources[0]))

	cachedResponse, err := resp.GetDiscoveryResponse()
	assert.Nil(t, err)
	assert.Same(t, discoveryResponse, cachedResponse)

	r := &route.RouteConfiguration{}
	err = anypb.UnmarshalTo(discoveryResponse.Resources[0], r, proto.UnmarshalOptions{})
	assert.Nil(t, err)
	assert.Equal(t, r.Name, resourceName)
}

func isTTLResource(resource *anypb.Any) bool {
	wrappedResource := &discovery.Resource{}
	err := protojson.Unmarshal(resource.Value, wrappedResource)
	if err != nil {
		return false
	}

	return wrappedResource.Resource == nil
}
