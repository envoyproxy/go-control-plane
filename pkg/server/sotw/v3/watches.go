package sotw

import (
	"context"
	"reflect"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// subscription is the state of the stream for a given resource type (cannot be AnyType)
type subscription struct {
	nonce string
	watch watch
	state stream.SubscriptionState
	// ToDo: move tracking of previous responses in state
	lastResponseResources map[string]string
}

func newSubscription(wildcard bool) *subscription {
	return &subscription{
		state: stream.NewSubscriptionState(wildcard, nil),
	}
}

// subscriptions for all xDS resource types
type subscriptions struct {
	responders map[string]*subscription

	// cases is a dynamic select case for the watched channels.
	cases []reflect.SelectCase
}

// newSubscriptions creates and initializes subscriptions.
func newSubscriptions() subscriptions {
	return subscriptions{
		responders: make(map[string]*subscription, int(types.UnknownType)),
		cases:      make([]reflect.SelectCase, 0),
	}
}

// addSubscription creates a new subscription entry in the subscriptions map.
// subscriptions are sorted by typeURL.
func (w *subscriptions) addSubscription(typeURL string, subscription *subscription) {
	w.responders[typeURL] = subscription
}

// close all open watches
func (w *subscriptions) close() {
	for _, subscription := range w.responders {
		subscription.watch.close()
	}
}

// recomputeWatches rebuilds the known list of dynamic channels if needed
func (w *subscriptions) recomputeWatches(ctx context.Context, req <-chan *discovery.DiscoveryRequest) {
	w.cases = w.cases[:0] // Clear the existing cases while retaining capacity.

	w.cases = append(w.cases,
		reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ctx.Done()),
		}, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(req),
		},
	)

	for _, subscription := range w.responders {
		w.cases = append(w.cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(subscription.watch.response),
		})
	}
}

// watch contains the necessary modifiable data for receiving resource responses
type watch struct {
	cancel   func()
	response chan cache.Response
}

// close cancels an open watch
func (w watch) close() {
	if w.cancel != nil {
		w.cancel()
	}
}
