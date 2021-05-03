package delta

import (
	"context"
	"errors"
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

	// streamNonce holds a unique nonce for req-resp pairs per xDS stream. The server
	// ignores stale nonces and nonce is only modified within send() function.
	var streamNonce int64

	// a collection of stack allocated watches per request type
	watches := newWatches()

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

	if s.callbacks != nil {
		if err := s.callbacks.OnDeltaStreamOpen(str.Context(), streamID, defaultTypeURL); err != nil {
			return err
		}
	}

	var node = &core.Node{}
	isWildcard := map[string]bool{}

	for {
		select {
		case <-s.ctx.Done():
			return nil
		case resp, more := <-watches.deltaMuxedResponses:
			if more {
				typ := resp.GetDeltaRequest().GetTypeUrl()

				if isErr := resp == deltaErrorResponse; isErr {
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
		case req, more := <-reqCh:
			// input stream ended or errored out
			if !more {
				return nil
			}
			if req == nil {
				return status.Errorf(codes.Unavailable, "empty request")
			}

			if s.callbacks != nil {
				if err := s.callbacks.OnStreamDeltaRequest(streamID, req); err != nil {
					return err
				}
			}

			// we verify nonce only if nonce is not initialized
			var nonce string
			// the node information may only be set on the first incoming delta discovery request
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

			typeURL := req.GetTypeUrl()
			var state stream.StreamState
			// Initialize the state map if we haven't already.
			// This means that we're handling the first request for this type on the stream.
			if s, ok := watches.deltaStreamStates[typeURL]; !ok {
				state.ResourceVersions = make(map[string]string)
			} else {
				state = s
			}

			// We are in the wildcard mode if the first request of a particular type has an empty subscription list
			var found bool
			if state.IsWildcard, found = isWildcard[typeURL]; !found {
				state.IsWildcard = len(req.GetResourceNamesSubscribe()) == 0
				isWildcard[typeURL] = state.IsWildcard
			}

			s.subscribe(req.GetResourceNamesSubscribe(), state.ResourceVersions)
			// We assume this is the first request on the stream
			if streamNonce == 0 {
				for r, v := range req.InitialResourceVersions {
					state.ResourceVersions[r] = v
				}
			} else if streamNonce > 0 && len(req.InitialResourceVersions) > 0 {
				return status.Errorf(codes.InvalidArgument, "InitialResourceVersions can only be set on initial stream request")
			}
			s.unsubscribe(req.GetResourceNamesUnsubscribe(), state.ResourceVersions)

			// cancel existing watch to (re-)request a newer version
			watch, seen := watches.deltaResponses[typeURL]
			if !seen || watch.nonce == nonce {
				// We must signal goroutine termination to prevent a race between the cancel closing the watch
				// and the producer closing the watch.`	`
				if terminate, exists := watches.deltaTerminations[typeURL]; exists {
					close(terminate)
				}
				if watch.cancel != nil {
					watch.cancel()
				}

				watch.responses, watch.cancel = s.cache.CreateDeltaWatch(req, &state)

				// a go-routine. Golang does not allow selecting over a dynamic set of channels.
				terminate := make(chan struct{})
				watches.deltaTerminations[typeURL] = terminate
				go func() {
					select {
					case resp, more := <-watch.responses:
						if more {
							watches.deltaMuxedResponses <- resp
						} else {
							// Check again if the watch is cancelled.
							select {
							case <-terminate: // do nothing
							// TODO: it might be worth logging something here to provide notice of a termination
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

func (s *server) DeltaStreamHandler(str stream.DeltaStream, typeURL string) error {
	// a channel for receiving incoming delta requests
	reqCh := make(chan *discovery.DeltaDiscoveryRequest)

	go func() {
		for {
			select {
			case <-str.Context().Done():
				close(reqCh)
				return
			default:
				req, err := str.Recv()
				if err != nil {
					close(reqCh)
					return
				}

				reqCh <- req
			}
		}
	}()

	return s.processDelta(str, reqCh, typeURL)
}

// When we subscribe, we just want to make the cache know we are subscribing to a resource.
// Providing a name with an empty version is enough to make that happen.
func (s *server) subscribe(resources []string, sv map[string]string) {
	for _, resource := range resources {
		sv[resource] = ""
	}
}

// Unsubscriptions remove resources from the stream state to
// indicate to the cache that we don't care about the resource anymore
func (s *server) unsubscribe(resources []string, sv map[string]string) {
	for _, resource := range resources {
		delete(sv, resource)
	}
}
