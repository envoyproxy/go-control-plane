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
	"errors"
	"strconv"

	"github.com/envoyproxy/go-control-plane/api"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

// ManagementServer is a collection of handlers for streaming discovery requests
type ManagementServer interface {
	// Register the handlers in a gRPC server
	Register(*grpc.Server)
}

type server struct {
	cache cache.ConfigCache
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

func (s *server) handler(stream stream, selector cache.ResourceSelector) error {
	// a channel for receiving incoming requests
	reqCh := make(chan *api.DiscoveryRequest, 0)
	go func() {
		for {
			req, err := stream.Recv()
			if err != nil {
				close(reqCh)
				return
			}
			reqCh <- req
		}
	}()

	// the node issuing the last request (assumed to be constant for the session)
	var node *api.Node

	// unique nonce for req-resp pairs; the server ignores stale nonces
	var nonce int

	// last successfully applied version as set in the requests
	var version string

	for {
		// make a cache request
		var configCh chan api.DiscoveryResponse
		if node != nil {
			configCh = s.cache.Listen(node, selector, version)
		}

		select {
		case resp := <-configCh:
			resp.Nonce = strconv.Itoa(nonce)
			nonce = nonce + 1
			if err := stream.Send(&resp); err != nil {
				return err
			}

		case req, closed := <-reqCh:
			// input stream ended or errored out
			if closed {
				return nil
			}

			// ignore stale non-empty nonces
			if req.GetResponseNonce() != "" && req.GetResponseNonce() != strconv.Itoa(nonce) {
				continue
			}

			// update cache request
			node = req.GetNode()
			version = req.GetVersionInfo()
			selector.Names = req.GetResourceNames()
			selector.Types = []string{req.TypeUrl}
		}
	}
}

func (s *server) StreamAggregatedResources(stream api.AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	return s.handler(stream, cache.ResourceSelector{})
}

func (s *server) StreamEndpoints(stream api.EndpointDiscoveryService_StreamEndpointsServer) error {
	return s.handler(stream, cache.ResourceSelector{
		Types: []string{cache.EndpointType},
	})
}

func (s *server) StreamLoadStats(stream api.EndpointDiscoveryService_StreamLoadStatsServer) error {
	return errors.New("not implemented")
}

func (s *server) StreamClusters(stream api.ClusterDiscoveryService_StreamClustersServer) error {
	return s.handler(stream, cache.ResourceSelector{
		Types: []string{cache.ClusterType},
	})
}

func (s *server) StreamRoutes(stream api.RouteDiscoveryService_StreamRoutesServer) error {
	return s.handler(stream, cache.ResourceSelector{
		Types: []string{cache.RouteType},
	})
}

func (s *server) StreamListeners(stream api.ListenerDiscoveryService_StreamListenersServer) error {
	return s.handler(stream, cache.ResourceSelector{
		Types: []string{cache.ListenerType},
	})
}

func (s *server) FetchEndpoints(ctx context.Context, req *api.DiscoveryRequest) (*api.DiscoveryResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *server) FetchClusters(ctx context.Context, req *api.DiscoveryRequest) (*api.DiscoveryResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *server) FetchRoutes(ctx context.Context, req *api.DiscoveryRequest) (*api.DiscoveryResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *server) FetchListeners(ctx context.Context, req *api.DiscoveryRequest) (*api.DiscoveryResponse, error) {
	return nil, errors.New("not implemented")
}
