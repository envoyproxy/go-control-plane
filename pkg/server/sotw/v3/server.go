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

	"google.golang.org/grpc"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

type Server interface {
	StreamHandler(stream Stream, typeURL string) error
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
	OnStreamResponse(int64, *discovery.DiscoveryRequest, *discovery.DiscoveryResponse)
}

type ProcessMethod int

const (
	ProcessDefault      ProcessMethod = iota
	ProcessXds          ProcessMethod = iota
	ProcessOrderedAds   ProcessMethod = iota
	ProcessUnorderedAds ProcessMethod = iota
)

// NewServer creates handlers from a config watcher and callbacks.
func NewServer(ctx context.Context, config cache.ConfigWatcher, callbacks Callbacks, opts ...Opt) Server {
	s := &server{cache: config, callbacks: callbacks, ctx: ctx}
	s.xdsProcessor = s.process
	s.adsProcessor = s.process
	for _, opt := range opts {
		opt(s)
	}
	return s
}

type Opt func(*server)

func WithCustomProcessMethods(xds, ads ProcessMethod) Opt {
	if ads == ProcessXds {
		panic("PROCESS_XDS cannot be supported for ADS mode")
	}
	return func(s *server) {
		switch xds {
		case ProcessDefault:
			s.xdsProcessor = s.process
		case ProcessXds:
			s.xdsProcessor = s.processXds
		case ProcessOrderedAds:
			s.xdsProcessor = s.processOrderedAds
		case ProcessUnorderedAds:
			s.xdsProcessor = s.processUnorderedAds
		}

		switch ads {
		case ProcessDefault:
			s.adsProcessor = s.process
		case ProcessOrderedAds:
			s.adsProcessor = s.processOrderedAds
		case ProcessUnorderedAds:
			s.adsProcessor = s.processUnorderedAds
		}
	}
}

type handler func(stream Stream, reqCh <-chan *discovery.DiscoveryRequest, defaultTypeURL string) error

type server struct {
	cache     cache.ConfigWatcher
	callbacks Callbacks
	ctx       context.Context

	// streamCount for counting bi-di streams
	streamCount int64

	xdsProcessor handler
	adsProcessor handler
}

// Generic RPC stream.
type Stream interface {
	grpc.ServerStream

	Send(*discovery.DiscoveryResponse) error
	Recv() (*discovery.DiscoveryRequest, error)
}

type streamWrapper struct {
	stream    Stream
	id        int64
	callbacks Callbacks

	// unique nonce generator for req-resp pairs per xDS stream; the server
	// ignores stale nonces. nonce is only modified within send() function.
	nonce int64
}

func (sw *streamWrapper) nextNonce() string {
	// increment nonce
	sw.nonce += 1
	return strconv.FormatInt(sw.nonce, 10)
}

// sends a response by serializing to protobuf Any
func (sw *streamWrapper) send(resp cache.Response) (string, error) {
	if resp == nil {
		return "", errors.New("missing response")
	}

	out, err := resp.GetDiscoveryResponse()
	if err != nil {
		return "", err
	}

	out.Nonce = sw.nextNonce()
	if sw.callbacks != nil {
		sw.callbacks.OnStreamResponse(sw.id, resp.GetRequest(), out)
	}
	return out.Nonce, sw.stream.Send(out)
}

// Token response value used to signal a watch failure in muxed watches.
var errorResponse = &cache.RawResponse{}

// StreamHandler converts a blocking read call to channels and initiates stream processing
func (s *server) StreamHandler(stream Stream, typeURL string) error {
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

	if typeURL == resource.AnyType {
		return s.adsProcessor(stream, reqCh, typeURL)
	} else {
		return s.xdsProcessor(stream, reqCh, typeURL)
	}
}
