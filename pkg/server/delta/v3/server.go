package delta

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// Server is a wrapper interface which is meant to hold the proper stream handler for each xDS protocol.
type Server interface {
	DeltaStreamHandler(stream stream.DeltaStream, typeURL string) error
}

type Callbacks interface {
	// OnDeltaStreamOpen is called once an incremental xDS stream is open with a stream ID and the type URL (or "" for ADS).
	// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
	OnDeltaStreamOpen(context.Context, int64, string) error
	// OnDeltaStreamClosed is called immediately prior to closing an xDS stream with a stream ID.
	OnDeltaStreamClosed(int64)
	// OnStreamDeltaRequest is called once a request is received on a stream.
	// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
	OnStreamDeltaRequest(int64, *discovery.DeltaDiscoveryRequest) error
	// OnStreamDelatResponse is called immediately prior to sending a response on a stream.
	OnStreamDeltaResponse(int64, *discovery.DeltaDiscoveryRequest, *discovery.DeltaDiscoveryResponse)
}

var deltaErrorResponse = &cache.RawDeltaResponse{}

type server struct {
	cache     cache.ConfigWatcher
	callbacks Callbacks

	// total stream count for counting bi-di streams
	streamCount int64
	ctx         context.Context
}

// NewServer creates a delta xDS specific server which utilizes a ConfigWatcher and delta Callbacks.
func NewServer(ctx context.Context, config cache.ConfigWatcher, callbacks Callbacks) Server {
	return &server{
		cache:     config,
		callbacks: callbacks,
		ctx:       ctx,
	}
}

func (s *server) processDelta(str stream.DeltaStream, reqCh <-chan *discovery.DeltaDiscoveryRequest, defaultTypeURL string) error {
	streamID := atomic.AddInt64(&s.streamCount, 1)

	// unique nonce generator for req-resp pairs per xDS stream; the server
	// ignores stale nonces. nonce is only modified within send() function.
	var streamNonce int64

	// a collection of stack allocated watches per request type
	watches := NewWatches()

	defer func() {
		watches.Cancel()
		if s.callbacks != nil {
			s.callbacks.OnDeltaStreamClosed(streamID)
		}
	}()

	// Sends a response, returns the new stream nonce
	send := func(resp cache.DeltaResponse) (string, error) {
		if resp == nil {
			return "", errors.New("missing response")
		}

		out, err := resp.GetDeltaDiscoveryResponse()
		if err != nil {
			return "", err
		}

		streamNonce = streamNonce + 1
		out.Nonce = strconv.FormatInt(streamNonce, 10)
		if s.callbacks != nil {
			s.callbacks.OnStreamDeltaResponse(streamID, resp.GetDeltaRequest(), out)
		}

		return out.Nonce, str.Send(out)
	}

	// Processes a response and updates the current server state
	process := func(resp cache.DeltaResponse, more bool) error {
		if resp == nil {
			return errors.New("missing response")
		}

		typ := resp.GetDeltaRequest().GetTypeUrl()

		if more {
			watches.mu.Lock()
			defer watches.mu.Unlock()

			if resp == deltaErrorResponse {
				return status.Errorf(codes.Unavailable, typ+" watch failed")
			}

			nonce, err := send(resp)
			if err != nil {
				return err
			}

			watch := watches.deltaResponses[typ]
			watch.nonce = nonce
			watches.deltaResponses[typ] = watch

			state := watches.deltaStreamStates[typ]
			if state.ResourceVersions == nil {
				state.ResourceVersions = make(map[string]string)
			}

			state.ResourceVersions = resp.GetNextVersionMap()
			watches.deltaStreamStates[typ] = state
		}

		return nil
	}

	rerequest := func(typ, nonce string, req *discovery.DeltaDiscoveryRequest, state *stream.StreamState) {
		watches.mu.Lock()
		defer watches.mu.Unlock()

		watch := watches.deltaResponses[typ]
		if watch.nonce == "" || watch.nonce == nonce {
			if watch.cancel != nil {
				if cancel := watch.cancel; cancel != nil {
					cancel()
				}
			}
			watch.responses, watch.cancel = s.cache.CreateDeltaWatch(req, state)
			watches.deltaResponses[typ] = watch
		}
	}

	if s.callbacks != nil {
		if err := s.callbacks.OnDeltaStreamOpen(str.Context(), streamID, defaultTypeURL); err != nil {
			return err
		}
	}

	// node may only be set on the first discovery request
	var node = &core.Node{}
	isWildcard := map[string]bool{}

	for {
		select {
		case <-s.ctx.Done():
			return nil
		case resp, more := <-watches.deltaResponses[resource.EndpointType].responses:
			err := process(resp, more)
			if err != nil {
				return err
			}
		case resp, more := <-watches.deltaResponses[resource.ClusterType].responses:
			err := process(resp, more)
			if err != nil {
				return err
			}
		case resp, more := <-watches.deltaResponses[resource.RouteType].responses:
			err := process(resp, more)
			if err != nil {
				return err
			}
		case resp, more := <-watches.deltaResponses[resource.ListenerType].responses:
			err := process(resp, more)
			if err != nil {
				return err
			}
		case resp, more := <-watches.deltaResponses[resource.SecretType].responses:
			err := process(resp, more)
			if err != nil {
				return err
			}
		case resp, more := <-watches.deltaResponses[resource.RuntimeType].responses:
			err := process(resp, more)
			if err != nil {
				return err
			}
		case resp, more := <-watches.deltaMuxedResponses:
			err := process(resp, more)
			if err != nil {
				return err
			}
		case req, more := <-reqCh:
			// input stream ended or errored out
			if !more {
				return nil
			}
			if req == nil {
				return status.Errorf(codes.Unavailable, "empty request")
			}

			// Log out our error from envoy if we get one, don't do anything crazy here yet
			// TODO: embed a logger in the server so we can be more verbose when needed
			if req.ErrorDetail != nil {
				fmt.Printf("received error from xDS client: %s", req.ErrorDetail.GetMessage())
			}

			// node field in discovery request is delta-compressed
			// nonces can be reused across streams; we verify nonce only if nonce is not initialized
			var nonce string
			if req.Node != nil {
				node = req.Node
				nonce = req.GetResponseNonce()
			} else {
				req.Node = node
				// If we have no nonce, i.e. this is the first request on a delta stream, set one
				nonce = strconv.FormatInt(streamNonce, 10)
			}

			// type URL is required for ADS but is implicit for xDS
			if defaultTypeURL == resource.AnyType {
				if req.TypeUrl == "" {
					return status.Errorf(codes.InvalidArgument, "type URL is required for ADS")
				}
			} else if req.TypeUrl == "" {
				req.TypeUrl = defaultTypeURL
			}

			state := watches.deltaStreamStates[req.GetTypeUrl()]
			// If this is empty we can assume this is the first time state
			// is being set on this stream for this resource type
			if state.ResourceVersions == nil {
				state.ResourceVersions = make(map[string]string)
			}

			// We are in the wildcard mode if the first request of a particular type has an empty subscription list
			var found bool
			if state.IsWildcard, found = isWildcard[req.TypeUrl]; !found {
				state.IsWildcard = len(req.GetResourceNamesSubscribe()) == 0
				isWildcard[req.TypeUrl] = state.IsWildcard
			}

			if sub := req.GetResourceNamesSubscribe(); len(sub) > 0 {
				s.subscribe(sub, state.ResourceVersions)
			}
			for r, v := range req.InitialResourceVersions {
				state.ResourceVersions[r] = v
			}
			if unsub := req.GetResourceNamesUnsubscribe(); len(unsub) > 0 {
				s.unsubscribe(unsub, state.ResourceVersions)
			}

			if s.callbacks != nil {
				if err := s.callbacks.OnStreamDeltaRequest(streamID, req); err != nil {
					return err
				}
			}

			// cancel existing watches to (re-)request a newer version
			for typ := range watches.deltaResponses {
				// If we've found our type, we go ahead and initiate the createWatch cycle
				if typ == req.TypeUrl {
					rerequest(typ, nonce, req, &state)
					continue
				}

				typeURL := req.TypeUrl
				responseNonce, seen := watches.deltaNonces[typeURL]
				if !seen || responseNonce == nonce {
					// We must signal goroutine termination to prevent a race between the cancel closing the watch
					// and the producer closing the watch.
					if terminate, exists := watches.deltaTerminations[typeURL]; exists {
						close(terminate)
					}
					if cancel, seen := watches.deltaCancellations[typeURL]; seen && cancel != nil {
						cancel()
					}

					var watch chan cache.DeltaResponse
					watch, watches.deltaCancellations[typeURL] = s.cache.CreateDeltaWatch(req, &state)

					// a go-routine. Golang does not allow selecting over a dynamic set of channels.
					terminate := make(chan struct{})
					watches.deltaTerminations[typeURL] = terminate
					go func() {
						select {
						case resp, more := <-watch:
							if more {
								watches.deltaMuxedResponses <- resp
							} else {
								// Check again if the watch is cancelled.
								select {
								case <-terminate: // do nothing
								default:
									// We cannot close the responses channel since it can be closed twice.
									// Instead we send a fake error response.
									watches.deltaMuxedResponses <- deltaErrorResponse
								}
							}
							break
						case <-terminate:
							break
						}
					}()
				}
			}
		}
	}
}

func (s *server) DeltaStreamHandler(str stream.DeltaStream, typeURL string) error {
	// a channel for receiving incoming delta requests
	reqCh := make(chan *discovery.DeltaDiscoveryRequest)
	reqStop := int32(0)

	go func() {
		for {
			req, err := str.Recv()
			if atomic.LoadInt32(&reqStop) != 0 {
				return
			}
			if err != nil {
				close(reqCh)
				return
			}

			reqCh <- req
		}
	}()

	err := s.processDelta(str, reqCh, typeURL)
	atomic.StoreInt32(&reqStop, 1)

	return err
}

// When we subscribe, we just want to make the cache know we are subscribing to a resource.
// Providing a name with an empty version is enough to make that happen.
func (s *server) subscribe(resources []string, sv map[string]string) {
	for _, resource := range resources {
		sv[resource] = ""
	}
}

// When we unsubscribe, we need to search and remove from the current subscribed list in the servers state
// so when we send that down to the cache, it knows to no longer track that resource
func (s *server) unsubscribe(resources []string, sv map[string]string) {
	for _, resource := range resources {
		delete(sv, resource)
	}
}
