package ttl

import (
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func MaybeCreateTtlResource(resource types.ResourceWithTtl, name string, heartbeat bool) (types.Resource, error) {
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
