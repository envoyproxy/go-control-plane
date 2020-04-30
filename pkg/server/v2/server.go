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

// Package server provides an implementation of a streaming xDS server.
package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	clusterservice "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	endpointservice "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	listenerservice "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	routeservice "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	discoverygrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	runtimeservice "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v2"
)

// Server is a collection of handlers for streaming discovery requests.
type Server interface {
	endpointservice.EndpointDiscoveryServiceServer
	clusterservice.ClusterDiscoveryServiceServer
	routeservice.RouteDiscoveryServiceServer
	listenerservice.ListenerDiscoveryServiceServer
	discoverygrpc.AggregatedDiscoveryServiceServer
	secretservice.SecretDiscoveryServiceServer
	runtimeservice.RuntimeDiscoveryServiceServer

	// Fetch is the universal fetch method.
	Fetch(context.Context, *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error)
}

// Callbacks is a collection of callbacks inserted into the server operation.
// The callbacks are invoked synchronously.
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

	OnStreamDeltaRequest(int64, *v2.DeltaDiscoveryRequest) error
	OnStreamDeltaResponse(int64, *v2.DeltaDiscoveryRequest, *v2.DeltaDiscoveryResponse)

	// OnFetchRequest is called for each Fetch request. Returning an error will end processing of the
	// request and respond with an error.
	OnFetchRequest(context.Context, *discovery.DiscoveryRequest) error
	// OnFetchResponse is called immediately prior to sending a response.
	OnFetchResponse(*discovery.DiscoveryRequest, *discovery.DiscoveryResponse)
}

// NewServer creates handlers from a config watcher and callbacks.
func NewServer(ctx context.Context, config cache.Cache, callbacks Callbacks) Server {
	return &server{cache: config, callbacks: callbacks, ctx: ctx}
}

type server struct {
	cache     cache.Cache
	callbacks Callbacks

	// streamCount for counting bi-di streams
	streamCount int64
	ctx         context.Context
}

// watches for all xDS resource types
type watches struct {
	endpoints chan cache.Response
	clusters  chan cache.Response
	routes    chan cache.Response
	listeners chan cache.Response
	secrets   chan cache.Response
	runtimes  chan cache.Response

	deltaEndpoints chan cache.DeltaResponse
	deltaClusters  chan cache.DeltaResponse
	deltaRoutes    chan cache.DeltaResponse
	deltaListeners chan cache.DeltaResponse
	deltaSecrets   chan cache.DeltaResponse
	deltaRuntimes  chan cache.DeltaResponse

	endpointCancel func()
	clusterCancel  func()
	routeCancel    func()
	listenerCancel func()
	secretCancel   func()
	runtimeCancel  func()

	endpointNonce string
	clusterNonce  string
	routeNonce    string
	listenerNonce string
	secretNonce   string
	runtimeNonce  string

	deltaClusterCancel func()

	deltaClusterNonce string
}

// Cancel all watches
func (values watches) Cancel() {
	if values.endpointCancel != nil {
		values.endpointCancel()
	}
	if values.clusterCancel != nil {
		values.clusterCancel()
	}
	if values.routeCancel != nil {
		values.routeCancel()
	}
	if values.listenerCancel != nil {
		values.listenerCancel()
	}
	if values.secretCancel != nil {
		values.secretCancel()
	}
	if values.runtimeCancel != nil {
		values.runtimeCancel()
	}
}

func (s *server) StreamAggregatedResources(stream discoverygrpc.AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	return s.handler(stream, resource.AnyType)
}

func (s *server) StreamEndpoints(stream endpointservice.EndpointDiscoveryService_StreamEndpointsServer) error {
	return s.handler(stream, resource.EndpointType)
}

func (s *server) StreamClusters(stream clusterservice.ClusterDiscoveryService_StreamClustersServer) error {
	return s.handler(stream, resource.ClusterType)
}

func (s *server) StreamRoutes(stream routeservice.RouteDiscoveryService_StreamRoutesServer) error {
	return s.handler(stream, resource.RouteType)
}

func (s *server) StreamListeners(stream listenerservice.ListenerDiscoveryService_StreamListenersServer) error {
	return s.handler(stream, resource.ListenerType)
}

func (s *server) StreamSecrets(stream secretservice.SecretDiscoveryService_StreamSecretsServer) error {
	return s.handler(stream, resource.SecretType)
}

func (s *server) StreamRuntime(stream runtimeservice.RuntimeDiscoveryService_StreamRuntimeServer) error {
	return s.handler(stream, resource.RuntimeType)
}

// Fetch is the universal fetch method.
func (s *server) Fetch(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if s.callbacks != nil {
		if err := s.callbacks.OnFetchRequest(ctx, req); err != nil {
			return nil, err
		}
	}
	resp, err := s.cache.Fetch(ctx, *req)
	if err != nil {
		return nil, err
	}
	out, err := createResponse(resp, req.TypeUrl)
	if s.callbacks != nil {
		s.callbacks.OnFetchResponse(req, out)
	}
	return out, err
}

func (s *server) FetchEndpoints(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = resource.EndpointType
	return s.Fetch(ctx, req)
}

func (s *server) FetchClusters(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = resource.ClusterType
	return s.Fetch(ctx, req)
}

func (s *server) FetchRoutes(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = resource.RouteType
	return s.Fetch(ctx, req)
}

func (s *server) FetchListeners(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = resource.ListenerType
	return s.Fetch(ctx, req)
}

func (s *server) FetchSecrets(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = resource.SecretType
	return s.Fetch(ctx, req)
}

func (s *server) FetchRuntime(ctx context.Context, req *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.Unavailable, "empty request")
	}
	req.TypeUrl = resource.RuntimeType
	return s.Fetch(ctx, req)
}

func (s *server) DeltaAggregatedResources(stream discoverygrpc.AggregatedDiscoveryService_DeltaAggregatedResourcesServer) error {
	return s.deltaHandler(stream, resource.AnyType)
}

func (s *server) DeltaEndpoints(stream endpointservice.EndpointDiscoveryService_DeltaEndpointsServer) error {
	return s.deltaHandler(stream, resource.EndpointType)
}

func (s *server) DeltaClusters(stream clusterservice.ClusterDiscoveryService_DeltaClustersServer) error {
	return s.deltaHandler(stream, resource.ClusterType)
}

func (s *server) DeltaRoutes(stream routeservice.RouteDiscoveryService_DeltaRoutesServer) error {
	return s.deltaHandler(stream, resource.RouteType)
}

func (s *server) DeltaListeners(stream listenerservice.ListenerDiscoveryService_DeltaListenersServer) error {
	return s.deltaHandler(stream, resource.ListenerType)
}

func (s *server) DeltaSecrets(stream discoverygrpc.SecretDiscoveryService_DeltaSecretsServer) error {
	return s.deltaHandler(stream, resource.SecretType)
}

func (s *server) DeltaRuntime(stream discoverygrpc.RuntimeDiscoveryService_DeltaRuntimeServer) error {
	return s.deltaHandler(stream, resource.RuntimeType)
}
