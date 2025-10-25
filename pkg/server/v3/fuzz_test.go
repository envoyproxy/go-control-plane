// Copyright 2025 Envoyproxy Authors
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

package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	deltav3 "github.com/envoyproxy/go-control-plane/pkg/server/delta/v3"
	restv3 "github.com/envoyproxy/go-control-plane/pkg/server/rest/v3"
	streamv3 "github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/anypb"
)

// fakeFetcher lets the fuzzer control REST Fetch() responses.
var _ cachev3.ConfigFetcher = (*fakeFetcher)(nil)

type fakeFetcher struct {
	makeResp func(*cachev3.Request) (cachev3.Response, error)
}

func (ff *fakeFetcher) Fetch(ctx context.Context, r *cachev3.Request) (cachev3.Response, error) {
	if ff.makeResp == nil {
		return nil, fmt.Errorf("no makeResp configured")
	}
	return ff.makeResp(r)
}

// fuzzWatcher lets the fuzzer control Delta/SotW watch responses
var _ cachev3.ConfigWatcher = (*fuzzWatcher)(nil)

type fuzzWatcher struct {
	makeDelta func(*cachev3.DeltaRequest) (cachev3.DeltaResponse, error)
	makeSotW  func(*cachev3.Request) (cachev3.Response, error)
}

// CreateWatch implements cache.ConfigWatcher for SotW.
// It immediately delivers a single fuzzer-controlled Response (if makeSotW is set)
// and returns a no-op cancel function.
func (fw *fuzzWatcher) CreateWatch(
	r *cachev3.Request,
	sub cachev3.Subscription,
	ch chan cachev3.Response,
) (func(), error) {
	if fw.makeSotW != nil && ch != nil {
		if resp, err := fw.makeSotW(r); err == nil && resp != nil {
			select {
			case ch <- resp:
			case <-time.After(10 * time.Millisecond):
			}
		}
	}
	return func() {}, nil
}

// CreateDeltaWatch implements cache.ConfigWatcher for Delta xDS.
// It immediately delivers a single fuzzer-controlled DeltaResponse (if makeDelta is set)
// and returns a no-op cancel function.
func (fw *fuzzWatcher) CreateDeltaWatch(
	r *cachev3.DeltaRequest,
	sub cachev3.Subscription,
	ch chan cachev3.DeltaResponse,
) (func(), error) {
	if fw.makeDelta != nil && ch != nil {
		if resp, err := fw.makeDelta(r); err == nil && resp != nil {
			select {
			case ch <- resp:
			case <-time.After(10 * time.Millisecond):
			}
		}
	}
	return func() {}, nil
}

// stubSOTW implements the minimal sotw.Server interface required by NewServerAdvanced.
type stubSOTW struct{}

func (s *stubSOTW) StreamHandler(_ streamv3.Stream, _ string) error { return nil }

// mockDeltaStream is an in-process fake Delta stream for fuzzing.
// Implements streamv3.DeltaStream and grpc.ServerStream.
type mockDeltaStream struct {
	ctx  context.Context
	req  *discovery.DeltaDiscoveryRequest
	sent []*discovery.DeltaDiscoveryResponse
	done bool
}

// grpc.ServerStream methods
func (m *mockDeltaStream) SetHeader(md metadata.MD) error  { return nil }
func (m *mockDeltaStream) SendHeader(md metadata.MD) error { return nil }
func (m *mockDeltaStream) SetTrailer(md metadata.MD)       {}
func (m *mockDeltaStream) Context() context.Context        { return m.ctx }
func (m *mockDeltaStream) SendMsg(interface{}) error       { return nil }
func (m *mockDeltaStream) RecvMsg(interface{}) error {
	if m.done {
		return io.EOF
	}
	m.done = true
	return nil
}

// Delta stream methods
func (m *mockDeltaStream) Recv() (*discovery.DeltaDiscoveryRequest, error) {
	if m.done {
		return nil, io.EOF
	}
	m.done = true
	return m.req, nil
}

func (m *mockDeltaStream) Send(resp *discovery.DeltaDiscoveryResponse) error {
	m.sent = append(m.sent, resp)
	return nil
}

// FuzzHTTPGatewayServeHTTP
//   - Fuzzes REST gateway across all REST endpoints.
//   - Mirrors request TypeUrl in responses so each fetch kind is exercised.
//   - Also builds a real delta server - in-process - with a fuzzWatcher and calls
//     DeltaStreamHandler directly via a mock stream.
func FuzzHTTPGatewayServeHTTP(f *testing.F) {
	f.Fuzz(func(t *testing.T,
		methodIndex uint8,
		body []byte,
		headerBytes1 []byte,
		headerBytes2 []byte,
		headerName1 string,
		headerName2 string,
		resourcePayload []byte,
		version string,
		nonce string,
	) {
		methods := []string{"GET", "POST", "PATCH", "DELETE"}
		paths := []string{
			"/v3/discovery:endpoints",
			"/v3/discovery:clusters",
			"/v3/discovery:routes",
			"/v3/discovery:scoped-routes",
			"/v3/discovery:listeners",
			"/v3/discovery:secrets",
			"/v3/discovery:runtime",
			"/v3/discovery:extensionconfigs",
		}

		// Choose a method
		method := methods[int(methodIndex)%len(methods)]
		
		// Choose a path and content type
		path := paths[int(methodIndex)%len(paths)]
		contentType := "application/json"

		ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
		defer cancel()

		// REST: fuzzer-controlled response that mirrors request TypeUrl.
		ff := &fakeFetcher{makeResp: func(r *cachev3.Request) (cachev3.Response, error) {
			any := &anypb.Any{TypeUrl: r.TypeUrl, Value: append([]byte{}, resourcePayload...)}
			dr := &discovery.DiscoveryResponse{
				TypeUrl:     r.TypeUrl,
				VersionInfo: version,
				Resources:   []*anypb.Any{any},
				Nonce:       nonce,
			}
			return &cachev3.PassthroughResponse{DiscoveryResponse: dr}, nil
		}}
		restSrv := restv3.NewServer(ff, nil)

		// DELTA: in-process watcher producing a single Delta response.
		fw := &fuzzWatcher{
			makeDelta: func(dr *cachev3.DeltaRequest) (cachev3.DeltaResponse, error) {
				ddr := &discovery.DeltaDiscoveryResponse{
					TypeUrl:           dr.TypeUrl,
					SystemVersionInfo: version,
				}
				return &cachev3.DeltaPassthroughResponse{DeltaDiscoveryResponse: ddr}, nil
			},
			makeSotW: func(r *cachev3.Request) (cachev3.Response, error) {
				any := &anypb.Any{TypeUrl: r.TypeUrl, Value: append([]byte{}, resourcePayload...)}
				dr := &discovery.DiscoveryResponse{
					TypeUrl:     r.TypeUrl,
					VersionInfo: version,
					Resources:   []*anypb.Any{any},
					Nonce:       nonce,
				}
				return &cachev3.PassthroughResponse{DiscoveryResponse: dr}, nil
			},
		}

		// Create the Delta server
		deltaSrv := deltav3.NewServer(ctx, fw, nil)

		// Create a full server with REST + SOTW(stub) + Delta(real).
		srv := NewServerAdvanced(restSrv, &stubSOTW{}, deltaSrv)

		req, err := http.NewRequest(method, path, bytes.NewReader(body))
		if err != nil {
			t.Skip(fmt.Sprintf("http.NewRequest error: %v", err))
			return
		}
		req = req.WithContext(ctx)
		req.Header.Set("Content-Type", contentType)
		if len(headerBytes1) > 0 {
			req.Header.Set(headerName1, fmt.Sprintf("%x", headerBytes1))
		}
		if len(headerBytes2) > 0 {
			req.Header.Set(headerName2, fmt.Sprintf("%x", headerBytes2))
		}

		gw := &HTTPGateway{Server: srv}
		_, _, _ = gw.ServeHTTP(req)

		deltaTypes := []string{
			"type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment",
			"type.googleapis.com/envoy.config.cluster.v3.Cluster",
			"type.googleapis.com/envoy.config.route.v3.RouteConfiguration",
			"type.googleapis.com/envoy.config.route.v3.ScopedRouteConfiguration",
			"type.googleapis.com/envoy.config.listener.v3.Listener",
			"type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.Secret",
			"type.googleapis.com/envoy.service.runtime.v3.Runtime",
			"type.googleapis.com/envoy.config.core.v3.TypedExtensionConfig",
		}

		// Chose a Delta type
		du := deltaTypes[int(methodIndex)%len(deltaTypes)]
		mds := &mockDeltaStream{ctx: ctx, req: &discovery.DeltaDiscoveryRequest{TypeUrl: du, ResponseNonce: nonce}}
		_ = deltaSrv.DeltaStreamHandler(mds, du)
	})
}
