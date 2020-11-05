package ttl

import "github.com/envoyproxy/go-control-plane/pkg/cache/types"

func MaybeCreateTtlResource(resource types.ResourceWithTtl, name string, heartbeat bool) (types.Resource, error) {
	return resource.Resource, nil
}
