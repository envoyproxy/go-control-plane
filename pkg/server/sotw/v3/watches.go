package sotw

import (
	"sync"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
)

// watches for all xDS resource types
type watches struct {
	mu         sync.RWMutex
	responders map[string]*watch
}

// newWatches creates and initializes watches.
func newWatches() watches {
	return watches{
		responders: make(map[string]*watch, int(types.UnknownType)),
	}
}

// addWatch creates a new watch entry in the watches map.
// Watches are sorted by typeURL.
func (w *watches) addWatch(typeURL string, watch *watch) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.responders[typeURL] = watch
}

func (w *watches) getWatch(typeURL string) (watch *watch) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.responders[typeURL]
}

// close all open watches
func (w *watches) close() {
	for _, watch := range w.responders {
		watch.close()
	}
}

// watch contains the necessary modifiables for receiving resource responses
type watch struct {
	mu     sync.RWMutex
	cancel func()
	nonce  string
}

func (w *watch) getNonce() (n string) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.nonce
}

func (w *watch) setNonce(n string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.nonce = n
}

// close cancels an open watch
func (w *watch) close() {
	if w.cancel != nil {
		w.cancel()
	}
}
