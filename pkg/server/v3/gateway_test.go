// Copyright 2018 Envoyproxy Authors
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

package server_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
)

func TestGateway(t *testing.T) {
	config := makeMockConfigWatcher()
	config.responses = map[string][]cache.Response{
		resource.ClusterType: {
			cache.NewTestRawResponse(
				&discovery.DiscoveryRequest{TypeUrl: resource.ClusterType},
				"2",
				[]types.ResourceWithTTL{{Resource: cluster}},
			),
		},
		resource.RouteType: {
			cache.NewTestRawResponse(
				&discovery.DiscoveryRequest{TypeUrl: resource.RouteType},
				"3",
				[]types.ResourceWithTTL{{Resource: route}},
			),
		},
		resource.ListenerType: {
			cache.NewTestRawResponse(
				&discovery.DiscoveryRequest{TypeUrl: resource.ListenerType},
				"4",
				[]types.ResourceWithTTL{{Resource: httpListener}, {Resource: httpScopedListener}},
			),
		},
	}
	gtw := server.HTTPGateway{Server: server.NewServer(context.Background(), config, nil)}

	failCases := []struct {
		path   string
		body   io.Reader
		expect int
	}{
		{
			path:   "/hello/",
			expect: http.StatusNotFound,
		},
		{
			path:   resource.FetchEndpoints,
			expect: http.StatusBadRequest,
		},
		{
			path:   resource.FetchEndpoints,
			body:   iotest.TimeoutReader(strings.NewReader("hello")),
			expect: http.StatusBadRequest,
		},
		{
			path:   resource.FetchEndpoints,
			body:   strings.NewReader("hello"),
			expect: http.StatusBadRequest,
		},
		{
			// missing response
			path:   resource.FetchEndpoints,
			body:   strings.NewReader("{\"node\": {\"id\": \"test\"}}"),
			expect: http.StatusInternalServerError,
		},
	}
	for _, cs := range failCases {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, cs.path, cs.body)
		require.NoError(t, err)
		resp, code, err := gtw.ServeHTTP(req)
		require.Errorf(t, err, "ServeHTTP succeeded, but should have failed")
		assert.Nilf(t, resp, "handler returned wrong response")
		assert.Equalf(t, code, cs.expect, "handler returned wrong status: %d, want %d", code, cs.expect)
	}

	for _, path := range []string{resource.FetchClusters, resource.FetchRoutes, resource.FetchListeners} {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, path, strings.NewReader("{\"node\": {\"id\": \"test\"}}"))
		require.NoError(t, err)
		resp, code, err := gtw.ServeHTTP(req)
		require.NoError(t, err)
		assert.NotNilf(t, resp, "handler returned wrong response")
		assert.Equalf(t, 200, code, "handler returned wrong status: %d, want %d", code, 200)
	}
}
