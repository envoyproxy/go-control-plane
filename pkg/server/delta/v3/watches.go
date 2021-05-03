package delta

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// watches for all delta xDS resource types
type watches struct {
	deltaResponses    map[string]watch
	deltaTerminations map[string]chan struct{}

	// Organize stream state by resource type
	deltaStreamStates map[string]stream.StreamState

	// Opaque resources share a muxed channel. Nonces and watch cancellations are indexed by type URL.
	deltaMuxedResponses chan cache.DeltaResponse
}

// watch contains the necessary modifiables for receiving resource responses
type watch struct {
	responses chan cache.DeltaResponse
	cancel    func()
	nonce     string
}

// newWatches creates and initializes watches.
func newWatches() watches {
	// deltaMuxedResponses needs a buffer to release go-routines populating it
	return watches{
		deltaResponses:      make(map[string]watch, int(types.UnknownType)),
		deltaMuxedResponses: make(chan cache.DeltaResponse, int(types.UnknownType)),
		deltaTerminations:   make(map[string]chan struct{}),
		deltaStreamStates:   make(map[string]stream.StreamState, int(types.UnknownType)),
	}
}

// Cancel all watches
func (w *watches) Cancel() {
	for _, watch := range w.deltaResponses {
		if watch.cancel != nil {
			watch.cancel()
		}
	}

	for _, terminate := range w.deltaTerminations {
		close(terminate)
	}
}
