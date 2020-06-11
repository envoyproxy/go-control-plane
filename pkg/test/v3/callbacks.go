package test

import (
	"context"
	"log"
	"sync"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
)

// Callbacks ...
type Callbacks struct {
	Signal   chan struct{}
	Debug    bool
	Fetches  int
	Requests int
	mu       sync.Mutex
}

// Report ...
func (cb *Callbacks) Report() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	log.Printf("server callbacks fetches=%d requests=%d\n", cb.Fetches, cb.Requests)
}

// OnStreamOpen ...
func (cb *Callbacks) OnStreamOpen(_ context.Context, id int64, typ string) error {
	if cb.Debug {
		log.Printf("stream %d open for %s\n", id, typ)
	}
	return nil
}

// OnStreamClosed ...
func (cb *Callbacks) OnStreamClosed(id int64) {
	if cb.Debug {
		log.Printf("stream %d closed\n", id)
	}
}

// OnStreamRequest ...
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

// OnStreamResponse ...
func (cb *Callbacks) OnStreamResponse(int64, *discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {
}

// OnStreamDeltaResponse ...
func (cb *Callbacks) OnStreamDeltaResponse(id int64, req *discovery.DeltaDiscoveryRequest, res *discovery.DeltaDiscoveryResponse) {
}

// OnStreamDeltaRequest ...
func (cb *Callbacks) OnStreamDeltaRequest(int64, *discovery.DeltaDiscoveryRequest) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.Requests++
	if cb.Signal != nil {
		close(cb.Signal)
		cb.Signal = nil
	}
	return nil
}

// OnFetchRequest ...
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

// OnFetchResponse ...
func (cb *Callbacks) OnFetchResponse(*discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {}
