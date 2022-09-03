# xDS Server Implementation

go-control-plane ships with a full [streaming implementation](https://github.com/envoyproxy/go-control-plane/blob/main/pkg/server/v3/server.go#L175) of the xDS protocol. Current support for the servers lists as follows:
- REST HTTP/1.1 *(This will soon be deprecated)*
- gRPC Bi-Di
	- State of the World
	- Incremental

## Getting Started

For a fully functional gRPC server, check out the provided example for what that looks like:
- https://github.com/envoyproxy/go-control-plane/blob/main/internal/example/server.go

### Callbacks

All go-control-plane xDS server implementations require `Callback` methods. Callbacks are executed at certain steps of the management server lifecycle. The interface to be implemented can be found [here](https://godoc.org/github.com/envoyproxy/go-control-plane/pkg/server/v2#Callbacks).

An example implemention of the Callback interface can be found below:
```go
import (
	"context"
	"log"
	"sync"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
)

type Callbacks struct {
	Signal         chan struct{}
	Debug          bool
	Fetches        int
	Requests       int
	DeltaRequests  int
	DeltaResponses int
	mu             sync.Mutex
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
func (cb *Callbacks) OnStreamClosed(id int64, node *core.Node) {
	if cb.Debug {
		log.Printf("stream %d of node %s closed\n", id, node.Id)
	}
}
func (cb *Callbacks) OnDeltaStreamOpen(_ context.Context, id int64, typ string) error {
	if cb.Debug {
		log.Printf("delta stream %d open for %s\n", id, typ)
	}
	return nil
}
func (cb *Callbacks) OnDeltaStreamClosed(id int64, node *core.Node) {
	if cb.Debug {
		log.Printf("delta stream %d of node %s closed\n", id, node.Id)
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
func (cb *Callbacks) OnStreamResponse(context.Context, int64, *discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {
}
func (cb *Callbacks) OnStreamDeltaResponse(id int64, req *discovery.DeltaDiscoveryRequest, res *discovery.DeltaDiscoveryResponse) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.DeltaResponses++
}
func (cb *Callbacks) OnStreamDeltaRequest(id int64, req *discovery.DeltaDiscoveryRequest) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.DeltaRequests++
	if cb.Signal != nil {
		close(cb.Signal)
		cb.Signal = nil
	}

	return nil
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

The internal go-control-plane gRPC server implementations take care of managing watches with the [Config Watcher](https://github.com/envoyproxy/go-control-plane/blob/main/pkg/cache/v3/cache.go#L45) when new xDS clients register themselves.

> *NOTE*: The server supports REST/JSON as well as gRPC bi-di streaming
