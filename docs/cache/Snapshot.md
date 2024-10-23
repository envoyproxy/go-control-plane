# SnapshotCache

SnapshotCache is a snapshot-based cache that maintains a single versioned snapshot of responses per node. SnapshotCache consistently replies with the latest snapshot. For the protocol to work correctly in ADS mode, EDS/RDS requests are responded to only when all resources in the snapshot xDS response are named as part of the request. It is expected that the CDS response names all EDS clusters, and the LDS response names all RDS routes in a snapshot, to ensure that Envoy makes the request for all EDS clusters or RDS routes eventually.

SnapshotCache can operate as a REST or regular xDS backend. The snapshot can be partial, e.g., only including RDS or EDS resources.

## Create a Snapshot Cache

To create a snapshot cache, we simply need to execute the provided constructor which will return an implemented `cache.SnapshotCache` interface.

```go
// Create a cache
cache := cachev3.NewSnapshotCache(false, cachev3.IDHash{}, logger)
```

The `IDHash{}` struct is requested by the cache to use when mapping snapshots to nodes. If no hashing function is present, it will simply use the Envoy node ID as the key within the cache map.

We have defined a `logger` interface which enables us to log at different levels within the source. Please refer to the [godoc](https://godoc.org/github.com/envoyproxy/go-control-plane/pkg/log) on what 
functions to implement when creating the logger.

> *NOTE*: This cache will need high level access within your management server as it is the source of truth for all configuration.

## Setting a Snapshot

To create a new snapshot with resources that must be mapped to an Envoy node, we provide the following helper method:

```go
snapshot := cache.NewSnapshot(
    "v1",
    endpoints,
    clusters,
    routes,
    listeners,
    runtimes,
    secrets,
)
```

For a more in-depth example on how to genereate a snapshot programmatically, explore how we create a generated snapshot in the integration test found [here](https://github.com/envoyproxy/go-control-plane/blob/master/pkg/test/resource/v2/resource.go#L317).

> *NOTE*: Not all resources are required to generate a configuration snapshot as it can be a partial.

We recommend verifying that your new `snapshot` is consistent within itself meaning that the dependent resources are exactly listed in the snapshot:

- all EDS resources are listed by name in CDS resources
- all RDS resources are listed by name in LDS resources.

> *NOTE*: Clusters and listeners are requested without name references, so Envoy will accept the snapshot list of clusters as-is even if it does not match all references found in xDS.

```go
if err := snapshot.Consistent(); err != nil {
   l.Errorf("snapshot inconsistency: %+v\n%+v", snapshot, err)
   os.Exit(1)
}
```

When ready to set a snapshot to an Envoy node, execute the below logic:

```go
// Add the snapshot to the cache
if err := cache.SetSnapshot("envoy-node-id", snapshot); err != nil {
    l.Errorf("snapshot error %q for %+v", err, snapshot)
    os.Exit(1)
}
```

Note that an Envoy node ID must be provided along with the snapshot object. Internally we map the two together so each Envoy node can receive the latest version of its configuration snapshot. When this is executed, go-control-plane will kick off new watches on the new resources and watch for changes throughout the server.
