# xDS Server Implementation

go-control-plane ships with a full streaming implementation of the xDS protocol. We provide an interface as well as helper functions to instantiate a fully xDS compliant gRPC server which utilizes the [SnapshotCache](Snapshot.md) logic behind the scenes.

## Getting Started

Below is an example of a functional gRPC server utilizing the go-control-plane interfaces and helper functions.

```go
// Create a cache
cache := cachev3.NewSnapshotCache(false, cachev3.IDHash{}, example.Logger{})

// Create the snapshot that we'll serve to Envoy
snapshot := example.GenerateSnapshot()
if err := snapshot.Consistent(); err != nil {
    l.Errorf("snapshot inconsistency: %+v\n%+v", snapshot, err)
    os.Exit(1)
}
l.Debugf("will serve snapshot %+v", snapshot)

// Add the snapshot to the cache
if err := cache.SetSnapshot("envoy-node-id", snapshot); err != nil {
    l.Errorf("snapshot error %q for %+v", err, snapshot)
    os.Exit(1)
}

// Run the xDS server
ctx := context.Background()
cb := &testv3.Callbacks{}
srv := serverv3.NewServer(ctx, cache, cb)
example.RunServer(ctx, srv, port)
```

For a full in-depth example of this code, please view the integration test located in `pkg/test/main` and the example code in `internal/example`.

### Callbacks

All go-control-plane xDS servers require `Callback` methods to run properly. These callbacks are executed throughout certain steps of the xDS lifecycle. The Callback interface which must be implemented can be found [here](https://godoc.org/github.com/envoyproxy/go-control-plane/pkg/server/v2#Callbacks).

An example implemention of the Callback interface can be found below:
```go
type Callbacks struct {
	Signal   chan struct{}
	Debug    bool
	Fetches  int
	Requests int
	mu       sync.Mutex
}

func (cb *Callbacks) Report() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	log.Printf("server callbacks fetches=%d requests=%d\n", cb.Fetches, cb.Requests)
}

func (cb *Callbacks) OnStreamOpen(_ context.Context, id int64, typ string) error {
	if cb.Debug {
		log.Printf("stream %d open for %s\n", id, typ)
	}
	return nil
}

func (cb *Callbacks) OnStreamClosed(id int64) {
	if cb.Debug {
		log.Printf("stream %d closed\n", id)
	}
}

func (cb *Callbacks) OnStreamRequest(int64, *discovery.DiscoveryRequest) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.Requests++
	if cb.Signal != nil {
		close(cb.Signal)
		cb.Signal = nil
	}
	return nil
}

func (cb *Callbacks) OnStreamResponse(int64, *discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {
}

func (cb *Callbacks) OnFetchRequest(_ context.Context, req *discovery.DiscoveryRequest) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.Fetches++
	if cb.Signal != nil {
		close(cb.Signal)
		cb.Signal = nil
	}
	return nil
}
func (cb *Callbacks) OnFetchResponse(*discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {}
```

## Info

The provided gRPC server will take care of creating new watches when new Envoy nodes register with the management server.

> *NOTE*: The server supports REST/JSON as well as gRPC streaming
