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

	"golang.org/x/sync/errgroup"
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
	OnStreamClosed(int64, *core.Node)
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

	streamState := stream.NewSTOWStreamState()

	// a collection of stack allocated watches per request type
	watches := newWatches()

	// node may only be set on the first discovery request
	var node = &core.Node{}

	defer func() {
		if s.callbacks != nil {
			s.callbacks.OnStreamClosed(streamID, node)
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

		lastResponse := stream.NewLastDiscoveryResponse(out.Nonce, resp.GetRequest().ResourceNames)
		streamState.Set(resp.GetRequest().TypeUrl, lastResponse)

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

	var resCh = make(chan cache.Response, 1)

	ctx, cancel := context.WithCancel(s.ctx)
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		defer func() {
			watches.close() // this should remove all watches from the cache
			close(resCh)    // close resCh and let the second eg.Go drain it
		}()

		for {
			select {
			case <-ctx.Done():
				return nil
			case req, more := <-reqCh:
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

				if lastResponse, ok := streamState.Get(req.TypeUrl); ok {
					if lastResponse.Nonce == "" || lastResponse.Nonce == nonce {
						// Let's record Resource names that a client has received.
						streamState.SetKnownResourceNames(req.TypeUrl, lastResponse.Resources)
					}
				}

				typeURL := req.GetTypeUrl()
				if w := watches.getWatch(typeURL); w != nil {
					// We've found a pre-existing watch, lets check and update if needed.
					// If these requirements aren't satisfied, leave an open watch.
					if n := w.getNonce(); n == "" || n == nonce {
						w.close()

						watches.addWatch(typeURL, &watch{
							cancel: s.cache.CreateWatch(req, streamState.StreamState, resCh),
						})
					}
				} else {
					// No pre-existing watch exists, let's create one.
					// We need to precompute the watches first then open a watch in the cache.
					watches.addWatch(typeURL, &watch{
						cancel: s.cache.CreateWatch(req, streamState.StreamState, resCh),
					})
				}
			}
		}
	})

	eg.Go(func() (err error) {
		var nonce string
		for res := range resCh {
			if res == nil || err != nil {
				continue // this loop should not exit until resCh closed
			}
			if nonce, err = send(res); err == nil {
				if w := watches.getWatch(res.GetRequest().TypeUrl); w != nil {
					w.setNonce(nonce)
				}
			} else {
				cancel()
			}
		}
		return err
	})

	return eg.Wait()
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
