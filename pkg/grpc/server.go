// Copyright 2017 Envoyproxy Authors
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

// Package server provides an implementation of a streaming xDS server
package server

import (
	"strconv"
	"sync/atomic"

	"github.com/envoyproxy/go-control-plane/api"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Resource types in xDS v2
const (
	typePrefix   = "type.googleapis.com/envoy.api.v2."
	EndpointType = typePrefix + "LbEndpoint"
	ClusterType  = typePrefix + "Cluster"
	RouteType    = typePrefix + "Route"
	ListenerType = typePrefix + "Listener"
)

// Server is a collection of handlers for streaming discovery requests
type Server interface {
	// Register the handlers in a gRPC server
	Register(*grpc.Server)
}

// NewServer creates handlers from a config watcher
func NewServer(config cache.ConfigWatcher) Server {
	return &server{config: config}
}

type server struct {
	config cache.ConfigWatcher
}

func (s *server) Register(srv *grpc.Server) {
	api.RegisterAggregatedDiscoveryServiceServer(srv, s)
	api.RegisterEndpointDiscoveryServiceServer(srv, s)
	api.RegisterClusterDiscoveryServiceServer(srv, s)
	api.RegisterRouteDiscoveryServiceServer(srv, s)
	api.RegisterListenerDiscoveryServiceServer(srv, s)
}

type stream interface {
	Send(*api.DiscoveryResponse) error
	Recv() (*api.DiscoveryRequest, error)
}

// promises for all xDS resource types
type promises struct {
	endpoints cache.Promise
	clusters  cache.Promise
	routes    cache.Promise
	listeners cache.Promise
}

func (values promises) Cancel() {
	values.endpoints.Cancel()
	values.clusters.Cancel()
	values.routes.Cancel()
	values.listeners.Cancel()
}

func (s *server) process(stream stream, reqCh <-chan *api.DiscoveryRequest, implicitTypeURL string) error {
	// unique nonce for req-resp pairs; the server ignores stale nonces
	var nonce int64
	send := func(resp *api.DiscoveryResponse) error {
		resp.Nonce = strconv.FormatInt(nonce, 10)
		nonce = nonce + 1
		return stream.Send(resp)
	}

	var values promises
	defer func() {
		values.Cancel()
	}()

	for {
		select {
		case resp := <-values.endpoints.Value:
			if err := send(&resp); err != nil {
				return err
			}
		case resp := <-values.clusters.Value:
			if err := send(&resp); err != nil {
				return err
			}
		case resp := <-values.routes.Value:
			if err := send(&resp); err != nil {
				return err
			}
		case resp := <-values.listeners.Value:
			if err := send(&resp); err != nil {
				return err
			}
		case req, closed := <-reqCh:
			switch {
			case closed:
				// input stream ended or errored out
				return nil
			case req.GetResponseNonce() != "" && req.GetResponseNonce() != strconv.FormatInt(nonce, 10):
				// ignore stale non-empty nonces
			default:
				// proxy requests a new resource
				// proxy can NACKs by sending an older version for the same resource type
				typeURL := implicitTypeURL
				if req.TypeUrl != "" {
					typeURL = req.TypeUrl
				}
				// cancel existing promises to (re-)request a newer version
				switch typeURL {
				case EndpointType:
					values.endpoints.Cancel()
					values.endpoints = s.config.WatchEndpoints(req.GetNode(), req.GetVersionInfo(), req.GetResourceNames())
				case ClusterType:
					values.clusters.Cancel()
					values.clusters = s.config.WatchClusters(req.GetNode(), req.GetVersionInfo(), req.GetResourceNames())
				case RouteType:
					values.routes.Cancel()
					values.routes = s.config.WatchRoutes(req.GetNode(), req.GetVersionInfo(), req.GetResourceNames())
				case ListenerType:
					values.listeners.Cancel()
					values.listeners = s.config.WatchListeners(req.GetNode(), req.GetVersionInfo(), req.GetResourceNames())
				}
			}
		}
	}
}

func (s *server) handler(stream stream, implicitTypeURL string) error {
	// a channel for receiving incoming requests
	reqCh := make(chan *api.DiscoveryRequest, 0)
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

	err := s.process(stream, reqCh, implicitTypeURL)

	// prevents writing to a closed channel if send failed on blocked recv
	// TODO(kuat) figure out how to unblock recv through gRPC API
	atomic.StoreInt32(&reqStop, 1)

	return err
}

func (s *server) StreamAggregatedResources(stream api.AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	return s.handler(stream, "")
}

func (s *server) StreamEndpoints(stream api.EndpointDiscoveryService_StreamEndpointsServer) error {
	return s.handler(stream, EndpointType)
}

func (s *server) StreamLoadStats(stream api.EndpointDiscoveryService_StreamLoadStatsServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *server) StreamClusters(stream api.ClusterDiscoveryService_StreamClustersServer) error {
	return s.handler(stream, ClusterType)
}

func (s *server) StreamRoutes(stream api.RouteDiscoveryService_StreamRoutesServer) error {
	return s.handler(stream, RouteType)
}

func (s *server) StreamListeners(stream api.ListenerDiscoveryService_StreamListenersServer) error {
	return s.handler(stream, ListenerType)
}

func (s *server) FetchEndpoints(ctx context.Context, req *api.DiscoveryRequest) (*api.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *server) FetchClusters(ctx context.Context, req *api.DiscoveryRequest) (*api.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *server) FetchRoutes(ctx context.Context, req *api.DiscoveryRequest) (*api.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *server) FetchListeners(ctx context.Context, req *api.DiscoveryRequest) (*api.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}
