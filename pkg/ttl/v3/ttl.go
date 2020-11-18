package ttl

import (
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

// Helper functions for interacting with TTL resources for xDS V3. A resource will be wrapped in a discovery.Resource in order
// to allow specifying a TTL. If the resource is meant to be a heartbeat response, only the resource name and TTL will be set
// to avoid having to send the entire resource down.

func MaybeCreateTtlResourceIfSupported(resource types.ResourceWithTtl, name string, heartbeat bool) (types.Resource, error) {
	if resource.Ttl != nil {
		wrappedResource := &discovery.Resource{
			Name: name,
			Ttl:  ptypes.DurationProto(*resource.Ttl),
		}

		if !heartbeat {
			any, err := ptypes.MarshalAny(resource.Resource)
			if err != nil {
				return nil, err
			}
			wrappedResource.Resource = any
		}

		return wrappedResource, nil
	}

	return resource.Resource, nil
}

func IsTTLResource(resource *any.Any) bool {
	wrappedResource := &discovery.Resource{}
	err := ptypes.UnmarshalAny(resource, wrappedResource)
	if err != nil {
		return false
	}

	return wrappedResource.Resource == nil
}
