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

	streamStates := map[string]stream.StreamState{}

	// a collection of stack allocated watches per request type
	watches := newWatches()

	// node may only be set on the first discovery request
	var node = &core.Node{}

	defer func() {
		watches.close()
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
			// State cannot be modified until any potential watch is closed
			state, ok := streamStates[typeURL]
			if !ok {
				// We don't have a current state for this type, create one
				state = stream.NewStreamState(len(req.ResourceNames) == 0, nil)
			} else if nonce != state.LastResponseNonce() {
				// The request does not match the last response we sent.
				// The protocol is a bit unclear in this case, but currently we discard such request and wait
				// for the client to acknowledge the previous response.
				// This is unclear how this handles cases where a response would be missed as any subsequent request will be discarded.

				// We can continue here as the watch list hasn't changed so we don't need to recompute the watches select
				continue
			}

			responder := make(chan cache.Response, 1)
			if w, ok := watches.responders[typeURL]; ok {
				// If we had an open watch, close it to make sure we don't end up sending a cache response while we update the state of the request
				w.close()

				// Check if the new request ACKs the previous response
				// This is defined as nonce and versions are matching, as well as no error present
				if state.LastResponseNonce() == nonce && state.LastResponseVersion() == req.VersionInfo && req.ErrorDetail == nil {
					// The nonce and versions are matching, this is an ACK from the client
					state.CommitPendingResources()
				}
			}

			// Remove resources no longer subscribed from the stream state
			// This ensures we will send a resource if it is unsubscribed then subscribed again
			// without a cache version change
			knownResources := state.GetKnownResources()
			unsubscribedResources := getUnsubscribedResources(req.ResourceNames, knownResources)
			for _, resourceName := range unsubscribedResources {
				delete(knownResources, resourceName)
			}
			// Remove from pending resources to ensure we won't lose this state when commiting
			state.RemovePendingResources(unsubscribedResources)

			watches.addWatch(typeURL, &watch{
				cancel:   s.cache.CreateWatch(req, &state, responder),
				response: responder,
			})

			streamStates[typeURL] = state

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

			// Track the resources returned in the response
			// Those are staged pending the client ACK
			// The protocol clearly states that if we send another response prior to an ACK
			// the previous one is to be considered as discarded
			version, err := res.GetVersion()
			if err != nil {
				return err
			}

			state := streamStates[res.GetRequest().TypeUrl]
			resources := make(map[string]string, len(res.GetResourceNames()))
			for _, name := range res.GetResourceNames() {
				resources[name] = version
			}
			// Pending resources can be modified in the server while a watch is opened
			// It will only be visible to caches once those resources are commited
			state.SetPendingResources(nonce, version, resources)
			streamStates[res.GetRequest().TypeUrl] = state
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

func getUnsubscribedResources(newResources []string, knownResources map[string]string) (removedResources []string) {
	newResourcesMap := make(map[string]struct{}, len(newResources))
	for _, resourceName := range newResources {
		newResourcesMap[resourceName] = struct{}{}
	}
	for resourceName := range knownResources {
		if _, ok := newResourcesMap[resourceName]; !ok {
			removedResources = append(removedResources, resourceName)
		}
	}
	return
}
