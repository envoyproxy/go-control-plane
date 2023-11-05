package cache_test

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

var fixture = &fixtureGenerator{
	version:  "x",
	version2: "y",
}

type fixtureGenerator struct {
	version  string
	version2 string
}

func (f *fixtureGenerator) snapshot() *cache.Snapshot {
	snapshot, err := cache.NewSnapshot(
		f.version,
		map[rsrc.Type][]types.Resource{
			rsrc.EndpointType:        {testEndpoint},
			rsrc.ClusterType:         {testCluster},
			rsrc.RouteType:           {testRoute, testEmbeddedRoute},
			rsrc.ScopedRouteType:     {testScopedRoute},
			rsrc.VirtualHostType:     {testVirtualHost},
			rsrc.ListenerType:        {testScopedListener, testListener},
			rsrc.RuntimeType:         {testRuntime},
			rsrc.SecretType:          {testSecret[0]},
			rsrc.ExtensionConfigType: {testExtensionConfig},
		},
	)
	if err != nil {
		panic(err.Error())
	}

	return snapshot
}
