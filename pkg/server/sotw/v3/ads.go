package sotw

import (
	"fmt"
	"reflect"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// process handles a bi-di stream request
func (s *server) processADS(sw *streamWrapper, reqCh <-chan *discovery.DiscoveryRequest, defaultTypeURL string) error {
	// process := func(resp cache.Response) error {
	// 	nonce, err := send(resp)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	typeUrl := resp.GetRequest().TypeUrl
	// 	watches.nonces[typeUrl] = nonce
	// 	return nil
	// }

	// processAllExcept := func(typeUrl string) error {
	// 	for {
	// 		select {
	// 		case resp := <-values.responses:
	// 			if resp.GetRequest().TypeUrl != typeUrl {
	// 				if err := process(resp); err != nil {
	// 					return err
	// 				}
	// 			}
	// 		default:
	// 			return nil
	// 		}
	// 	}
	// }

	// node may only be set on the first discovery request
	var node = &core.Node{}

	// This control loop strictly orders resources when running in ADS mode.
	// It should be treated as a child process of the original process() loop
	// and should return on close of stream or error. This will cause the cleanup routines
	// in the parent process() loop to execute.
	for {
		fmt.Printf("%+v\n", sw.watches.cases)
		index, value, ok := reflect.Select(sw.watches.cases)
		fmt.Println(index)
		fmt.Println(value)
		fmt.Println(ok)

		switch index {
		case 0:
			fmt.Println("recieved a Done() on the ADS code path.")
			// ctx.Done() -> no further computation is needed
			return nil
		case 1:
			fmt.Println("This is the ordered ADS code path request pipeline")
			if !ok {
				return nil
			}

			req := value.Interface().(*discovery.DiscoveryRequest)
			if req == nil {
				return status.Errorf(codes.Unavailable, "empty request")
			}

			// Only first request is guaranteed to hold node info so if it's missing, reassign.
			if req.Node != nil {
				node = req.Node
			} else {
				req.Node = node
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

			// if lastResponse, ok := lastDiscoveryResponses[req.TypeUrl]; ok {
			// 	if lastResponse.nonce == "" || lastResponse.nonce == nonce {
			// 		// Let's record Resource names that a client has received.
			// 		streamState.SetKnownResourceNames(req.TypeUrl, lastResponse.resources)
			// 	}
			// }

			typeURL := req.GetTypeUrl()
			responder := make(chan cache.Response, 1)
			if w, ok := sw.watches.responders[typeURL]; ok {
				// We've found a pre-existing watch, lets check and update if needed.
				// If these requirements aren't satisfied, leave an open watch.
				if w.nonce == "" || w.nonce == nonce {
					w.close()

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
			fmt.Println("This is the ordered ADS code path default pipeline")

			// Channel n -> these are the dynamic list of responders that correspond to the stream request typeURL
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

// typeUrl := req.TypeUrl
// responseNonce, seen := values.nonces[typeUrl]
// if !seen || responseNonce == nonce {
// 	if cancel, seen := values.cancellations[typeUrl]; seen {
// 		if cancel != nil {
// 			cancel()
// 		}
// 		if err := processAllExcept(typeUrl); err != nil {
// 			return err
// 		}
// 	}
// 	values.cancellations[typeUrl] = s.cache.CreateWatch(req, values.responses)
// }
