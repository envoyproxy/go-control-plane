package delta

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// watches for all delta xDS resource types
type watches struct {
	deltaWatches map[string]watch

	// Opaque resources share a muxed channel
	deltaMuxedResponses chan cache.DeltaResponse
}

// newWatches creates and initializes watches.
func newWatches() watches {
	// deltaMuxedResponses needs a buffer to release go-routines populating it
	return watches{
		deltaWatches:        make(map[string]watch, int(types.UnknownType)),
		deltaMuxedResponses: make(chan cache.DeltaResponse, int(types.UnknownType)),
	}
}

// Cancel all watches
func (w *watches) Cancel() {
	for _, watch := range w.deltaWatches {
		if watch.cancel != nil {
			watch.cancel()
		}

		watch.terminate()
	}
}

// watch contains the necessary modifiables for receiving resource responses
type watch struct {
	responses chan cache.DeltaResponse
	cancel    func()
	nonce     string

	termination chan struct{}

	state stream.StreamState
}

// GetState returns the current state of the watch
func (w *watch) GetState() stream.StreamState {
	return w.state
}

// Cancel calls terminate and cancel
func (w *watch) Cancel() {
	if w.cancel != nil {
		w.cancel()
	}

	if w.termination != nil {
		close(w.termination)
	}
}

func (w *watch) terminate() {
	if w.termination != nil {
		close(w.termination)
	}
}
