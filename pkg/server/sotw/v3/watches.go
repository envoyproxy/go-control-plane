package sotw

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

// watches for all xDS resource types
type watches struct {
	watches map[string]watch

	// Opaque resources share a muxed channel
	muxedResponses chan cache.Response
}

// newWatches creates and initializes watches.
func newWatches() watches {
	// deltaMuxedResponses needs a buffer to release go-routines populating it
	return watches{
		watches:        make(map[string]watch, int(types.UnknownType)),
		muxedResponses: make(chan cache.Response, int(types.UnknownType)),
	}
}

// Cancel all watches
func (w *watches) Cancel() {
	for _, watch := range w.watches {
		watch.Cancel()
	}
}

// watch contains the necessary modifiables for receiving resource responses
type watch struct {
	responses chan cache.Response
	cancel    func()
	nonce     string

	termination chan struct{}
}

// Cancel calls terminate and cancel
func (w *watch) Cancel() {
	if w.cancel != nil {
		w.cancel()
	}

	w.terminate()
}

func (w *watch) terminate() {
	if w.termination != nil {
		close(w.termination)
	}
}
