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

type xdsState struct {
	responses chan cache.Response
	nonce     string
	cancel    func()
}

func newXdsState() xdsState {
	return xdsState{}
}

func (v *xdsState) Cancel() {
	if v.cancel != nil {
		v.cancel()
	}
}

// process handles a bi-di stream request
func (s *server) processXds(stream Stream, reqCh <-chan *discovery.DiscoveryRequest, defaultTypeURL string) error {
	sw := streamWrapper{
		stream:    stream,
		id:        atomic.AddInt64(&s.streamCount, 1), // increment stream count
		callbacks: s.callbacks,
	}

	// a collection of stack allocated watches per request type
	values := newXdsState()
	defer values.Cancel()
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
		// config watcher can send the requested resources types in any order
		case resp, more := <-values.responses:
			if !more {
				return status.Errorf(codes.Unavailable, "resource watch failed")
			}
			nonce, err := sw.send(resp)
			if err != nil {
				return err
			}
			values.nonce = nonce

		case req, more := <-reqCh:
			// input stream ended or errored out
			if !more {
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

			// cancel existing watches to (re-)request a newer version
			if values.nonce == "" || values.nonce == nonce {
				values.Cancel()
				values.responses = make(chan cache.Response, 1)
				values.cancel = s.cache.CreateWatch(req, values.responses)
			}
		}
	}
}
