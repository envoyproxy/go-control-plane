package delta

import (
	"sync"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// watches for all delta xDS resource types
type watches struct {
	mu *sync.RWMutex

	deltaResponses     map[string]watch
	deltaCancellations map[string]func()
	deltaNonces        map[string]string
	deltaTerminations  map[string]chan struct{}

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

// NewWatches creates and initializes watches.
func NewWatches() watches {
	dr := make(map[string]watch, 7)
	for i := 0; i < 6; i++ {
		typ, err := cache.GetResponseTypeURL(types.ResponseType(i))
		if err != nil {
			panic(err)
		}

		dr[typ] = watch{
			responses: make(chan cache.DeltaResponse),
		}
	}

	// deltaMuxedResponses needs a buffer to release go-routines populating it
	return watches{
		deltaMuxedResponses: make(chan cache.DeltaResponse, 6),
		deltaResponses:      dr,
		deltaNonces:         make(map[string]string),
		deltaTerminations:   make(map[string]chan struct{}),
		deltaCancellations:  make(map[string]func()),
		deltaStreamStates:   make(map[string]stream.StreamState, int(types.UnknownType)),
		mu:                  &sync.RWMutex{},
	}
}

// Cancel all watches
func (w *watches) Cancel() {
	for _, watch := range w.deltaResponses {
		if watch.cancel != nil {
			watch.cancel()
		}
	}

	for _, cancel := range w.deltaCancellations {
		if cancel != nil {
			cancel()
		}
	}

	for _, terminate := range w.deltaTerminations {
		close(terminate)
	}
}
