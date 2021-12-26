# SnapshotCache

[SnapshotCache](https://github.com/envoyproxy/go-control-plane/blob/main/pkg/cache/v3/simple.go#L40) is a snapshot-based caching mechanism that maintains a versioned config snapshot per node. See the original README scope section for more detailed explanations on the individual cache systems.

> *NOTE*: SnapshotCache can operate as a REST or bi-di streaming xDS server

## Create a Snapshot Cache

To create a snapshot cache, we simply call the provided constructor:
```go
// Create a cache
cache := cache.NewSnapshotCache(false, cache.IDHash{}, nil)
```
This `cache` object holds a fully compliant [SnapshotCache](https://github.com/envoyproxy/go-control-plane/blob/main/pkg/cache/v3/simple.go#L40) with the necessary methods to manage a configuration cache lifecycle.

> *NOTE*: The cache needs a high access level inside the management server as it's the source of truth for xDS resources.

## Snapshots

Snapshots are groupings of resources at a given point in time for a node cluster. In other words Envoy's and consuming xDS clients registering as node `abc` all share a snapshot of config. This snapshot is the singular source of truth in the cache that represents config for any of those consumers.

> *NOTE*: Snapshots can be partial, e.g., only including RDS or EDS resources. 

```go
snap, err := cache.NewSnapshot("v0", map[resource.Type][]types.Resource{
    resource.EndpointType:        endpoints,
    resource.ClusterType:         clusters,
    resource.RouteType:           routes,
    resource.ScopedRouteType:     scopedRoutes,
    resource.ListenerType:        listeners,
    resource.RuntimeType:         runtimes,
    resource.SecretType:          secrets,
    resource.ExtensionConfigType: extensions,
})
```

For a more in-depth example of how to genereate a snapshot, explore our example found [here](https://github.com/envoyproxy/go-control-plane/blob/main/internal/example/resource.go#L168).

We recommend verifying that your new `snapshot` is consistent within itself meaning that the dependent resources are exactly listed in the snapshot:

```go
if err := snapshot.Consistent(); err != nil {
   l.Errorf("snapshot inconsistency: %+v\n%+v", snapshot, err)
   os.Exit(1)
}
```

- all EDS resources are listed by name in CDS resources
- all SRDS/RDS resources are listed by name in LDS resources
- all RDS resources are listed by name in SRDS resources

> *NOTE*: Clusters and Listeners are requested without name references, so Envoy will accept the snapshot list of clusters as-is even if it does not match all references found in xDS.

Setting a snapshot is as simple as:
```go
// Add the snapshot to the cache
if err := cache.SetSnapshot("envoy-node-id", snapshot); err != nil {
    l.Errorf("snapshot error %q for %+v", err, snapshot)
    os.Exit(1)
}
```

This will trigger all open watches internal to the caching [config watchers](https://github.com/envoyproxy/go-control-plane/blob/main/pkg/cache/v3/cache.go#L45) and anything listening for changes will received updates and responses from the new snapshot.

*Note*: that a node ID must be provided along with the snapshot object. Internally a mapping of the two is kept so each node can receive the latest version of its configuration.
