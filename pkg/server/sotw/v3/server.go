// Copyright 2020 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Package sotw provides an implementation of GRPC SoTW (State of The World) part of XDS server
package sotw

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

type Server interface {
	StreamHandler(stream stream.Stream, typeURL string) error
}

type Callbacks interface {
	// OnStreamOpen is called once an xDS stream is open with a stream ID and the type URL (or "" for ADS).
	// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
	OnStreamOpen(context.Context, int64, string) error
	// OnStreamClosed is called immediately prior to closing an xDS stream with a stream ID.
	OnStreamClosed(int64)
	// OnStreamRequest is called once a request is received on a stream.
	// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
	OnStreamRequest(int64, *discovery.DiscoveryRequest) error
	// OnStreamResponse is called immediately prior to sending a response on a stream.
	OnStreamResponse(context.Context, int64, *discovery.DiscoveryRequest, *discovery.DiscoveryResponse)
}

// NewServer creates handlers from a config watcher and callbacks.
func NewServer(ctx context.Context, config cache.ConfigWatcher, callbacks Callbacks) Server {
	return &server{cache: config, callbacks: callbacks, ctx: ctx}
}

type server struct {
	cache     cache.ConfigWatcher
	callbacks Callbacks
	ctx       context.Context

	// streamCount for counting bi-di streams
	streamCount int64
}

// Token response value used to signal a watch failure in muxed watches.
var errorResponse = &cache.RawResponse{}

// process handles a bi-di stream request
func (s *server) process(stream stream.Stream, reqCh <-chan *discovery.DiscoveryRequest, defaultTypeURL string) error {
	// increment stream count
	streamID := atomic.AddInt64(&s.streamCount, 1)

	// unique nonce generator for req-resp pairs per xDS stream; the server
	// ignores stale nonces. nonce is only modified within send() function.
	var streamNonce int64

	streamState := streamv3.NewStreamState(false, map[string]string{})
	lastDiscoveryResponses := map[string]lastDiscoveryResponse{}

	// a collection of stack allocated watches per request type
	watches := newWatches()

	defer func() {
		watches.Cancel()
		if s.callbacks != nil {
			s.callbacks.OnStreamClosed(streamID)
		}
	}()

	// sends a response by serializing to protobuf Any
	send := func(resp cache.Response) (string, error) {
		if resp == nil {
			return "", errors.New("missing response")
		}

		out, err := resp.GetDiscoveryResponse()
		if err != nil {
			return "", err
		}

		// increment nonce
		streamNonce = streamNonce + 1
		out.Nonce = strconv.FormatInt(streamNonce, 10)

		lastResponse := lastDiscoveryResponse{
			nonce:     out.Nonce,
			resources: make(map[string]struct{}),
		}
		for _, r := range resp.GetRequest().ResourceNames {
			lastResponse.resources[r] = struct{}{}
		}
		lastDiscoveryResponses[resp.GetRequest().TypeUrl] = lastResponse

		if s.callbacks != nil {
			s.callbacks.OnStreamResponse(resp.GetContext(), streamID, resp.GetRequest(), out)
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
		case resp, more := <-watches.muxedResponses:
			if !more {
				break
			}

			typ := resp.GetRequest().GetTypeUrl()
			if resp == errorResponse {
				return status.Errorf(codes.Unavailable, typ+" watch failed")
			}

<<<<<<< HEAD
<<<<<<< HEAD
		case resp, more := <-values.runtimes:
			if !more {
				return status.Errorf(codes.Unavailable, "runtimes watch failed")
			}
			nonce, err := send(resp)
			if err != nil {
				return err
			}
			values.runtimeNonce = nonce

		case resp, more := <-values.extensionConfigs:
			if !more {
				return status.Errorf(codes.Unavailable, "extensionConfigs watch failed")
			}
			nonce, err := send(resp)
			if err != nil {
				return err
			}
			values.extensionConfigNonce = nonce

		case resp, more := <-values.responses:
			if more {
				if resp == errorResponse {
					return status.Errorf(codes.Unavailable, "resource watch failed")
				}
				typeURL := resp.GetRequest().TypeUrl
				nonce, err := send(resp)
				if err != nil {
					return err
				}
				values.nonces[typeURL] = nonce
			}
=======
			nonce, err := send(resp, typ)
=======
			nonce, err := send(resp)
>>>>>>> a31cdfc0 (rebase and fix errors)
			if err != nil {
				return err
			}
>>>>>>> c635a6b7 (dedupe and linearization)

			watch := watches.watches[typ]
			watch.nonce = nonce
			watches.watches[typ] = watch
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

<<<<<<< HEAD
			if lastResponse, ok := lastDiscoveryResponses[req.TypeUrl]; ok {
				if lastResponse.nonce == "" || lastResponse.nonce == nonce {
					// Let's record Resource names that a client has received.
					streamState.SetKnownResourceNames(req.TypeUrl, lastResponse.resources)
				}
			}

			// cancel existing watches to (re-)request a newer version
			switch {
			case req.TypeUrl == resource.EndpointType:
				if values.endpointNonce == "" || values.endpointNonce == nonce {
					if values.endpointCancel != nil {
						values.endpointCancel()
					}
					values.endpoints = make(chan cache.Response, 1)
					values.endpointCancel = s.cache.CreateWatch(req, streamState, values.endpoints)
				}
			case req.TypeUrl == resource.ClusterType:
				if values.clusterNonce == "" || values.clusterNonce == nonce {
					if values.clusterCancel != nil {
						values.clusterCancel()
					}
					values.clusters = make(chan cache.Response, 1)
					values.clusterCancel = s.cache.CreateWatch(req, streamState, values.clusters)
				}
			case req.TypeUrl == resource.RouteType:
				if values.routeNonce == "" || values.routeNonce == nonce {
					if values.routeCancel != nil {
						values.routeCancel()
					}
					values.routes = make(chan cache.Response, 1)
					values.routeCancel = s.cache.CreateWatch(req, streamState, values.routes)
				}
			case req.TypeUrl == resource.ScopedRouteType:
				if values.scopedRouteNonce == "" || values.scopedRouteNonce == nonce {
					if values.scopedRouteCancel != nil {
						values.scopedRouteCancel()
					}
					values.scopedRoutes = make(chan cache.Response, 1)
					values.scopedRouteCancel = s.cache.CreateWatch(req, streamState, values.scopedRoutes)
				}
			case req.TypeUrl == resource.ListenerType:
				if values.listenerNonce == "" || values.listenerNonce == nonce {
					if values.listenerCancel != nil {
						values.listenerCancel()
					}
					values.listeners = make(chan cache.Response, 1)
					values.listenerCancel = s.cache.CreateWatch(req, streamState, values.listeners)
				}
			case req.TypeUrl == resource.SecretType:
				if values.secretNonce == "" || values.secretNonce == nonce {
					if values.secretCancel != nil {
						values.secretCancel()
					}
					values.secrets = make(chan cache.Response, 1)
					values.secretCancel = s.cache.CreateWatch(req, streamState, values.secrets)
				}
			case req.TypeUrl == resource.RuntimeType:
				if values.runtimeNonce == "" || values.runtimeNonce == nonce {
					if values.runtimeCancel != nil {
						values.runtimeCancel()
					}
					values.runtimes = make(chan cache.Response, 1)
					values.runtimeCancel = s.cache.CreateWatch(req, streamState, values.runtimes)
				}
			case req.TypeUrl == resource.ExtensionConfigType:
				if values.extensionConfigNonce == "" || values.extensionConfigNonce == nonce {
					if values.extensionConfigCancel != nil {
						values.extensionConfigCancel()
					}
					values.extensionConfigs = make(chan cache.Response, 1)
					values.extensionConfigCancel = s.cache.CreateWatch(req, streamState, values.extensionConfigs)
				}
			default:
				typeURL := req.TypeUrl
				responseNonce, seen := values.nonces[typeURL]
				if !seen || responseNonce == nonce {
					if cancel, seen := values.cancellations[typeURL]; seen && cancel != nil {
						cancel()
					}
					values.cancellations[typeURL] = s.cache.CreateWatch(req, streamState, values.responses)
				}
=======
			typeURL := req.GetTypeUrl()

			// cancel existing watches to (re-)request a newer version
			watch, ok := watches.watches[typeURL]
			// nonces can be reused across streams; we verify nonce only if nonce is not initialized
			if !ok || watch.nonce == req.GetResponseNonce() {
				watch.Cancel()

				watch.responses = make(chan cache.Response, 1)
				watch.cancel = s.cache.CreateWatch(req, watch.responses)

				watches.watches[typeURL] = watch
				go func() {
					resp, more := <-watch.responses
					if !more {
						watches.muxedResponses <- errorResponse
						return
					}

					watches.muxedResponses <- resp
				}()
>>>>>>> c635a6b7 (dedupe and linearization)
			}
		}
	}
}

// StreamHandler converts a blocking read call to channels and initiates stream processing
func (s *server) StreamHandler(stream stream.Stream, typeURL string) error {
	// a channel for receiving incoming requests
	reqCh := make(chan *discovery.DiscoveryRequest)
	go func() {
		defer close(reqCh)
		for {
			req, err := stream.Recv()
			if err != nil {
				return
			}
			select {
			case reqCh <- req:
			case <-stream.Context().Done():
				return
			case <-s.ctx.Done():
				return
			}
		}
	}()

	return s.process(stream, reqCh, typeURL)
}
