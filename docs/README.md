# Knowledge Base

Below lies a list of resources that may be helpful to those looking to understand the go-control-plane API. The aim of these artifacts is to provide enough knowledge and understanding to newcomers and users who wish to use this API within their own codebases to implement an xDS compliant control-plane.

## Snapshot Cache
The following guides may be helpful on how to use go-control-plane's Snapshot Cache:
- [Snapshot.md](cache/Snapshot.md)
- [Server.md](cache/Server.md)

## Getting Started
Below is an example of a simple xDS ready server utilizing the provided Snapshot Cache and gRPC server logic.

```go
import (
    "context"
    "net"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/keepalive"

    api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
    discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
    "github.com/envoyproxy/go-control-plane/pkg/cache/v2"
    xds "github.com/envoyproxy/go-control-plane/pkg/server/v2"
)

const (
    grpcKeepaliveTime        = 30 * time.Second
    grpcKeepaliveTimeout     = 5 * time.Second
    grpcKeepaliveMinTime     = 30 * time.Second
    grpcMaxConcurrentStreams = 1000000
)

func main() {
    snapshotCache := cache.NewSnapshotCache(false, cache.IDHash{}, nil)
    server := xds.NewServer(context.Background(), snapshotCache, nil)
    var grpcOptions []grpc.ServerOption
    grpcOptions = append(grpcOptions,
        grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams),
        grpc.KeepaliveParams(keepalive.ServerParameters{
            Time:    grpcKeepaliveTime,
            Timeout: grpcKeepaliveTimeout,
        }),
        grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
            MinTime:             grpcKeepaliveMinTime,
            PermitWithoutStream: true,
        }),
    )
    grpcServer := grpc.NewServer(grpcOptions...)
    lis, _ := net.Listen("tcp", ":8080")

    discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
    api.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
    api.RegisterClusterDiscoveryServiceServer(grpcServer, server)
    api.RegisterRouteDiscoveryServiceServer(grpcServer, server)
    api.RegisterListenerDiscoveryServiceServer(grpcServer, server)
    go func() {
        if err := grpcServer.Serve(lis); err != nil {
            // error handling
        }
    }()
}
```

As mentioned in the README's [Scope](https://github.com/envoyproxy/go-control-plane/blob/master/README.md#scope), you need to cache Envoy configurations.
Generate the key for the corresponding snapshot based on the node information provided from an Envoy node, then cache the configurations.

```go
import (
    "github.com/envoyproxy/go-control-plane/pkg/cache/v2"
    "github.com/envoyproxy/go-control-plane/pkg/cache/types"
)

var clusters, endpoints, routes, listeners, runtimes []types.Resource

snapshotCache := cache.NewSnapshotCache(false, cache.IDHash{}, nil)
snapshot := cache.NewSnapshot("1.0", endpoints, clusters, routes, listeners, runtimes)
_ = snapshotCache.SetSnapshot("node1", snapshot)
```
