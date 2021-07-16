package sotw

import (
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

type orderedAdsState struct {
	// Opaque resources share a muxed channel. Nonces and watch cancellations are indexed by type URL.
	responses     chan cache.Response
	cancellations map[string]func()
	nonces        map[string]string
}

func newOrderedAdsState() orderedAdsState {
	return orderedAdsState{
		responses:     make(chan cache.Response, 8),
		cancellations: make(map[string]func()),
		nonces:        make(map[string]string),
	}
}

func (v *orderedAdsState) Cancel() {
	for _, cancel := range v.cancellations {
		if cancel != nil {
			cancel()
		}
	}
}

// process handles a bi-di stream request
func (s *server) processOrderedAds(stream Stream, reqCh <-chan *discovery.DiscoveryRequest, defaultTypeURL string) error {
	sw := streamWrapper{
		stream:    stream,
		id:        atomic.AddInt64(&s.streamCount, 1), // increment stream count
		callbacks: s.callbacks,
	}

	// a collection of stack allocated watches per request type
	values := newOrderedAdsState()
	defer values.Cancel()

	process := func(resp cache.Response) error {
		nonce, err := sw.send(resp)
		if err != nil {
			return err
		}
		typeUrl := resp.GetRequest().TypeUrl
		values.nonces[typeUrl] = nonce
		delete(values.cancellations, typeUrl)
		return nil
	}

	processAllExcept := func(typeUrl string) error {
		for {
			select {
			case resp := <-values.responses:
				if resp.GetRequest().TypeUrl != typeUrl {
					if err := process(resp); err != nil {
						return err
					}
				}
			default:
				return nil
			}
		}
	}

	defer func() {
		if s.callbacks != nil {
			s.callbacks.OnStreamClosed(sw.id)
		}
	}()

	if s.callbacks != nil {
		if err := s.callbacks.OnStreamOpen(stream.Context(), sw.id, defaultTypeURL); err != nil {
			return err
		}
	}

	// node may only be set on the first discovery request
	var node = &core.Node{}

	for {
		select {
		case <-s.ctx.Done():
			return nil
		case resp := <-values.responses:
			req := resp.GetRequest()
			nonce, err := sw.send(resp)
			if err != nil {
				return err
			}
			values.nonces[req.TypeUrl] = nonce
		case req, recvOk := <-reqCh:
			if !recvOk {
				return nil
			}
			if req == nil {
				return status.Errorf(codes.Unavailable, "empty request")
			}

			// node field in discovery request is delta-compressed
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
				if err := s.callbacks.OnStreamRequest(sw.id, req); err != nil {
					return err
				}
			}

			typeUrl := req.TypeUrl
			responseNonce, seen := values.nonces[typeUrl]
			if !seen || responseNonce == nonce {
				if cancel, seen := values.cancellations[typeUrl]; seen {
					if cancel != nil {
						cancel()
					}
					if err := processAllExcept(typeUrl); err != nil {
						return err
					}
				}
				values.cancellations[typeUrl] = s.cache.CreateWatch(req, values.responses)
			}
		}
	}
}
