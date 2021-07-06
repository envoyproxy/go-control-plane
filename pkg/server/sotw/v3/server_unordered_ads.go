package sotw

import (
	"reflect"
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

type unorderedAdsState struct {
	// Opaque resources share a muxed channel. Nonces and watch cancellations are indexed by type URL.
	selectCasesToTypeUrls map[int]string
	typeUrlsToSelectCases map[string]int
	selectCases           []reflect.SelectCase
	responses             map[string]chan cache.Response
	cancellations         map[string]func()
	nonces                map[string]string
}

func newUnorderedAdsState() unorderedAdsState {
	return unorderedAdsState{
		selectCasesToTypeUrls: make(map[int]string),
		typeUrlsToSelectCases: make(map[string]int),
		responses:             make(map[string]chan cache.Response),
		cancellations:         make(map[string]func()),
		nonces:                make(map[string]string),
	}
}

func (v *unorderedAdsState) Cancel() {
	for _, cancel := range v.cancellations {
		if cancel != nil {
			cancel()
		}
	}
}

// process handles a bi-di stream request
func (s *server) processUnorderedAds(stream Stream, reqCh <-chan *discovery.DiscoveryRequest, defaultTypeURL string) error {
	sw := streamWrapper{
		stream:    stream,
		id:        atomic.AddInt64(&s.streamCount, 1), // increment stream count
		callbacks: s.callbacks,
	}

	// a collection of stack allocated watches per request type
	values := newUnorderedAdsState()
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

	values.selectCases = append(values.selectCases, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(s.ctx.Done()),
	})
	values.selectCases = append(values.selectCases, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(reqCh),
	})

	for {
		chosen, recv, recvOk := reflect.Select(values.selectCases)
		if chosen == 0 { // done
			return nil
		} else if chosen == 1 { // req
			if !recvOk {
				return nil
			}
			req := recv.Interface().(*discovery.DiscoveryRequest)

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
				if cancel, seen := values.cancellations[typeUrl]; seen && cancel != nil {
					cancel()
				}
				values.responses[typeUrl] = make(chan cache.Response, 1)
				if !seen {
					values.selectCases = append(values.selectCases, reflect.SelectCase{
						Dir:  reflect.SelectRecv,
						Chan: reflect.ValueOf(values.responses[typeUrl]),
					})
					values.selectCasesToTypeUrls[len(values.selectCases)-1] = typeUrl
					values.typeUrlsToSelectCases[typeUrl] = len(values.selectCases) - 1
				} else {
					values.selectCases[values.typeUrlsToSelectCases[typeUrl]].Chan = reflect.ValueOf(values.responses[typeUrl])
				}
				values.cancellations[typeUrl] = s.cache.CreateWatch(req, values.responses[typeUrl])
			}
		} else {
			resp := recv.Interface().(cache.Response)
			req := resp.GetRequest()
			nonce, err := sw.send(resp)
			if err != nil {
				return err
			}
			values.nonces[req.TypeUrl] = nonce
		}
	}
}
