package sotw

import (
	"reflect"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// process handles a bi-di stream request
func (s *server) processADS(sw *streamWrapper, reqCh chan *discovery.DiscoveryRequest, defaultTypeURL string) error {
	process := func(resp cache.Response) error {
		nonce, err := sw.send(resp)
		if err != nil {
			return err
		}

		typeURL := resp.GetRequest().TypeUrl
		sw.watches.responders[typeURL].nonce = nonce
		return nil
	}

	processAllExcept := func(typeURL string) error {
		for {
			index, value, ok := reflect.Select(sw.watches.cases)
			// index is checked because if we receive a value
			// from the Done() or request channel here
			// we can ignore. The main control loop will handle that accordingly.
			// This is strictly for handling incoming resources off the dynamic channels list.
			if !ok || index < 2 {
				// We just exit and assume the ordered parent resource isn't ready if !ok
				return nil
			}

			res := value.Interface().(cache.Response)
			if res.GetRequest().TypeUrl != typeURL {
				if err := process(res); err != nil {
					return err
				}
			}
		}
	}

	// This control loop strictly orders resources when running in ADS mode.
	// It should be treated as a child process of the original process() loop
	// and should return on close of stream or error.
	// This will cause the cleanup routinesin the parent process() loop to execute.
	for {
		index, value, ok := reflect.Select(sw.watches.cases)
		switch index {
		case 0:
			// ctx.Done() -> no further computation is needed
			return nil
		case 1:
			if !ok {
				return nil
			}

			req := value.Interface().(*discovery.DiscoveryRequest)
			if req == nil {
				return status.Errorf(codes.Unavailable, "empty request")
			}

			// Only first request is guaranteed to hold node info so if it's missing, reassign.
			if req.Node != nil {
				sw.node = req.Node
			} else {
				req.Node = sw.node
			}

			// nonces can be reused across streams; we verify nonce only if nonce is not initialized
			nonce := req.GetResponseNonce()

			// type URL is required for ADS but is implicit for xDS
			if defaultTypeURL == resource.AnyType {
				if req.TypeUrl == "" {
					return status.Errorf(codes.InvalidArgument, "type URL is required for ADS")
				}
			} else if req.TypeUrl == "" {
				req.TypeUrl = defaultTypeURL
			}

			if s.callbacks != nil {
				if err := s.callbacks.OnStreamRequest(sw.ID, req); err != nil {
					return err
				}
			}

			if lastResponse, ok := sw.lastDiscoveryResponses[req.TypeUrl]; ok {
				if lastResponse.nonce == "" || lastResponse.nonce == nonce {
					// Let's record Resource names that a client has received.
					sw.streamState.SetKnownResourceNames(req.TypeUrl, lastResponse.resources)
				}
			}

			typeURL := req.GetTypeUrl()
			responder := make(chan cache.Response, 1)
			if w, ok := sw.watches.responders[typeURL]; ok {
				// We've found a pre-existing watch, lets check and update if needed.
				// If these requirements aren't satisfied, leave an open watch.
				if w.nonce == "" || w.nonce == nonce {
					w.close()

					// Only process if we have an existing watch otherwise go ahead and create.
					if err := processAllExcept(typeURL); err != nil {
						return err
					}

					sw.watches.addWatch(typeURL, &watch{
						cancel:   s.cache.CreateWatch(req, stream.StreamState{}, responder),
						response: responder,
					})
				}
			} else {
				// No pre-existing watch exists, let's create one.
				// We need to precompute the watches first then open a watch in the cache.
				sw.watches.addWatch(typeURL, &watch{
					cancel:   s.cache.CreateWatch(req, stream.StreamState{}, responder),
					response: responder,
				})
			}

			// Recompute the dynamic select cases for this stream.
			sw.watches.recompute(s.ctx, reqCh)
		default:
			// Channel n -> these are the dynamic list of responders
			// that correspond to the stream request typeURL
			// No special processing here for ordering. We just send on the channels.
			if !ok {
				// Receiver channel was closed. TODO(jpeach): probably cancel the watch or something?
				return status.Errorf(codes.Unavailable, "resource watch %d -> failed", index)
			}

			res := value.Interface().(cache.Response)
			nonce, err := sw.send(res)
			if err != nil {
				return err
			}

			sw.watches.responders[res.GetRequest().TypeUrl].nonce = nonce
		}
	}
}
