package ttl

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/golang/protobuf/ptypes/any"
)

func MaybeCreateTtlResource(resource types.ResourceWithTtl, name string, heartbeat bool) (types.Resource, error) {
	return resource.Resource, nil
}

func IsTTLResource(resource *any.Any) bool {
	return true
}
