package server

import (
	"errors"
	"log"
	"strconv"
	"sync/atomic"

	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type stream interface {
	grpc.ServerStream

	Send(*discovery.DiscoveryResponse) error
	Recv() (*discovery.DiscoveryRequest, error)
}

func createResponse(resp cache.Response, typeURL string) (*discovery.DiscoveryResponse, error) {
	if resp == nil {
		return nil, errors.New("missing response")
	}

	marshalledResponse, err := resp.GetDiscoveryResponse()
	if err != nil {
		return nil, err
	}

	return marshalledResponse, nil
}

// process handles a bi-di stream request
func (s *server) process(stream stream, reqCh <-chan *discovery.DiscoveryRequest, defaultTypeURL string) error {
	// increment stream count
	streamID := atomic.AddInt64(&s.streamCount, 1)

	// unique nonce generator for req-resp pairs per xDS stream; the server
	// ignores stale nonces. nonce is only modified within send() function.
	var streamNonce int64

	// a collection of watches per request type
	var values watches
	defer func() {
		values.Cancel()
		if s.callbacks != nil {
			s.callbacks.OnStreamClosed(streamID)
		}
	}()

	// sends a response by serializing to protobuf Any
	send := func(resp cache.Response, typeURL string) (string, error) {
		out, err := createResponse(resp, typeURL)
		if err != nil {
			return "", err
		}

		// increment nonce
		streamNonce = streamNonce + 1
		out.Nonce = strconv.FormatInt(streamNonce, 10)
		if s.callbacks != nil {
			s.callbacks.OnStreamResponse(streamID, resp.GetRequest(), out)
		}
		return out.Nonce, stream.Send(out)
	}

	if s.callbacks != nil {
		if err := s.callbacks.OnStreamOpen(stream.Context(), streamID, defaultTypeURL); err != nil {
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
		case resp, more := <-values.endpoints:
			if !more {
				return status.Errorf(codes.Unavailable, "endpoints watch failed")
			}
			nonce, err := send(resp, resource.EndpointType)
			if err != nil {
				return err
			}
			values.endpointNonce = nonce

		case resp, more := <-values.clusters:
			if !more {
				return status.Errorf(codes.Unavailable, "clusters watch failed")
			}
			nonce, err := send(resp, resource.ClusterType)
			if err != nil {
				return err
			}
			values.clusterNonce = nonce

		case resp, more := <-values.routes:
			if !more {
				return status.Errorf(codes.Unavailable, "routes watch failed")
			}
			nonce, err := send(resp, resource.RouteType)
			if err != nil {
				return err
			}
			values.routeNonce = nonce

		case resp, more := <-values.listeners:
			if !more {
				return status.Errorf(codes.Unavailable, "listeners watch failed")
			}
			nonce, err := send(resp, resource.ListenerType)
			if err != nil {
				return err
			}
			values.listenerNonce = nonce

		case resp, more := <-values.secrets:
			if !more {
				return status.Errorf(codes.Unavailable, "secrets watch failed")
			}
			nonce, err := send(resp, resource.SecretType)
			if err != nil {
				return err
			}
			values.secretNonce = nonce

		case resp, more := <-values.runtimes:
			if !more {
				return status.Errorf(codes.Unavailable, "runtimes watch failed")
			}
			nonce, err := send(resp, resource.RuntimeType)
			if err != nil {
				return err
			}
			values.runtimeNonce = nonce

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
				if err := s.callbacks.OnStreamRequest(streamID, req); err != nil {
					return err
				}
			}

			// cancel existing watches to (re-)request a newer version
			switch {
			case req.TypeUrl == resource.EndpointType && (values.endpointNonce == "" || values.endpointNonce == nonce):
				if values.endpointCancel != nil {
					values.endpointCancel()
				}
				values.endpoints, values.endpointCancel = s.cache.CreateWatch(*req)
			case req.TypeUrl == resource.ClusterType && (values.clusterNonce == "" || values.clusterNonce == nonce):
				if values.clusterCancel != nil {
					values.clusterCancel()
				}
				log.Println(nonce)
				values.clusters, values.clusterCancel = s.cache.CreateWatch(*req)
			case req.TypeUrl == resource.RouteType && (values.routeNonce == "" || values.routeNonce == nonce):
				if values.routeCancel != nil {
					values.routeCancel()
				}
				values.routes, values.routeCancel = s.cache.CreateWatch(*req)
			case req.TypeUrl == resource.ListenerType && (values.listenerNonce == "" || values.listenerNonce == nonce):
				if values.listenerCancel != nil {
					values.listenerCancel()
				}
				values.listeners, values.listenerCancel = s.cache.CreateWatch(*req)
			case req.TypeUrl == resource.SecretType && (values.secretNonce == "" || values.secretNonce == nonce):
				if values.secretCancel != nil {
					values.secretCancel()
				}
				values.secrets, values.secretCancel = s.cache.CreateWatch(*req)
			case req.TypeUrl == resource.RuntimeType && (values.runtimeNonce == "" || values.runtimeNonce == nonce):
				if values.runtimeCancel != nil {
					values.runtimeCancel()
				}
				values.runtimes, values.runtimeCancel = s.cache.CreateWatch(*req)
			}
		}
	}
}

// handler converts a blocking read call to channels and initiates stream processing
func (s *server) handler(stream stream, typeURL string) error {
	// a channel for receiving incoming requests
	reqCh := make(chan *discovery.DiscoveryRequest)
	reqStop := int32(0)
	go func() {
		for {
			req, err := stream.Recv()
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

	err := s.process(stream, reqCh, typeURL)

	// prevents writing to a closed channel if send failed on blocked recv
	// TODO(kuat) figure out how to unblock recv through gRPC API
	atomic.StoreInt32(&reqStop, 1)

	return err
}
