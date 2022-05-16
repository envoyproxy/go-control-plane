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
	"reflect"
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

// process handles a bi-di stream request
func (s *server) process(str stream.Stream, reqCh <-chan *discovery.DiscoveryRequest, defaultTypeURL string) error {
	// increment stream count
	streamID := atomic.AddInt64(&s.streamCount, 1)

	// unique nonce generator for req-resp pairs per xDS stream; the server
	// ignores stale nonces. nonce is only modified within send() function.
	var streamNonce int64

	// a collection of stack allocated watches per request type
	watches := newWatches()

	defer func() {
		watches.close()
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

		if s.callbacks != nil {
			s.callbacks.OnStreamResponse(resp.GetContext(), streamID, resp.GetRequest(), out)
		}
		return out.Nonce, str.Send(out)
	}

	if s.callbacks != nil {
		if err := s.callbacks.OnStreamOpen(str.Context(), streamID, defaultTypeURL); err != nil {
			return err
		}
	}

	// node may only be set on the first discovery request
	var node = &core.Node{}

	// recompute dynamic channels for this stream
	watches.recompute(s.ctx, reqCh)

	for {
		// The list of select cases looks like this:
		// 0: <- ctx.Done
		// 1: <- reqCh
		// 2...: per type watches
		index, value, ok := reflect.Select(watches.cases)
		switch index {
		// ctx.Done() -> if we receive a value here we return as no further computation is needed
		case 0:
			return nil
		// Case 1 handles any request inbound on the stream and handles all initialization as needed
		case 1:
			// input stream ended or errored out
			if !ok {
				return nil
			}

			req := value.Interface().(*discovery.DiscoveryRequest)
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

			typeURL := req.GetTypeUrl()
			responder := make(chan cache.Response, 1)

			var state stream.StreamState
			if w, ok := watches.responders[typeURL]; ok {
				if w.nonce != "" && w.nonce != nonce {
					// The request received does not match the current state of the type, we ignore it and keep our current watch
					continue
				}
				w.close()
				state = w.state
			} else {
				// Initialize the state of the stream.
				// Since there was no previous state, we know we're handling the first request of this type.
				// We also set the stream as wildcard based on its legacy meaning (no resource name sent in resource_names_subscribe).
				// If the state starts with this legacy mode, adding new resources will not unsubscribe from wildcard.
				// It can still be done by explicitly unsubscribing from "*"
				state = stream.NewStreamState(len(req.ResourceNames) == 0, nil)
			}

			// This update of registered resources must occur after the previous watch is closed if existing
			// It would otherwise potentially leak watches in caches
			s.registerResourceNames(req.ResourceNames, &state)
			watches.addWatch(typeURL, &watch{
				cancel:   s.cache.CreateWatch(req, state, responder),
				response: responder,
				state:    state,
			})

			// Recompute the dynamic select cases for this stream.
			watches.recompute(s.ctx, reqCh)
		default:
			// Channel n -> these are the dynamic list of responders that correspond to the stream request typeURL
			if !ok {
				// Receiver channel was closed. TODO(jpeach): probably cancel the watch or something?
				return status.Errorf(codes.Unavailable, "resource watch %d -> failed", index)
			}

			res := value.Interface().(cache.Response)
			nonce, err := send(res)
			if err != nil {
				return err
			}
			watches.responders[res.GetRequest().TypeUrl].nonce = nonce

			rawResponse := value.Interface().(*cache.RawResponse)
			watches.responders[res.GetRequest().TypeUrl].state.SetResourceVersions(rawResponse.NextVersionMap)
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

func (s *server) registerResourceNames(resources []string, streamState *stream.StreamState) {
	if len(resources) == 0 && streamState.IsWildcard() {
		// The xDS protocol states that if there has never been any resource set, the request should be considered wildcard
		// This would theoretically require keeping track on whether we ever became non-empty.
		// As it is also technically allowed to return resources which have not been subscribed to, it is a best effort here.
		return
	}

	// When resources are provided, they may still include the wildcard symbol '*', as well as potentially other resources
	// This allows the client to subscribe/unsubscribe to wildcard during the stream lifespan.
	wantsWildcard := false
	wantedResources := make(map[string]struct{}, len(resources))
	for _, resourceName := range resources {
		// We do not track '*' as a resource name to avoid confusion in further processing and rely on the IsWildcard method instead
		if resourceName == "*" {
			wantsWildcard = true
			continue
		}
		wantedResources[resourceName] = struct{}{}
	}
	streamState.SetWildcard(wantsWildcard)
	streamState.SetSubscribedResourceNames(wantedResources)
}
