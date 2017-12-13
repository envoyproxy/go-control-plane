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

// Package test contains test utilities
package test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/envoyproxy/go-control-plane/api"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	xds "github.com/envoyproxy/go-control-plane/pkg/server"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

const (
	node  = "service-node"
	hello = "Hi, there!\n"
)

// Hasher is a single cache key hash.
type Hasher struct {
}

// Hash function that always returns same value.
func (h Hasher) Hash(*api.Node) (cache.Key, error) {
	return cache.Key(node), nil
}

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	if _, err := w.Write([]byte(hello)); err != nil {
		glog.Error(err)
	}
}

// RunHTTP opens a simple listener on the port.
func RunHTTP(ctx context.Context, upstreamPort uint) {
	glog.Infof("upstream listening HTTP1.1 on %d", upstreamPort)
	h := handler{}
	server := &http.Server{Addr: fmt.Sprintf(":%d", upstreamPort), Handler: h}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			glog.Error(err)
		}
	}()
	if err := server.Shutdown(ctx); err != nil {
		glog.Error(err)
	}
}

// RunXDS starts an xDS server at the given port.
func RunXDS(ctx context.Context, config cache.Cache, port uint) {
	server := xds.NewServer(config)
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	api.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	api.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
	api.RegisterClusterDiscoveryServiceServer(grpcServer, server)
	api.RegisterRouteDiscoveryServiceServer(grpcServer, server)
	api.RegisterListenerDiscoveryServiceServer(grpcServer, server)
	glog.Infof("xDS server listening on %d", port)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			glog.Error(err)
		}
	}()
	<-ctx.Done()
	grpcServer.GracefulStop()
}

// RunCacheUpdate executes a config update sequence every second.
func RunCacheUpdate(ctx context.Context,
	config cache.Cache,
	ads bool,
	interval time.Duration,
	upstreamPort, listenPort uint) {
	i := 0
	for {
		version := fmt.Sprintf("version%d", i)
		clusterName := fmt.Sprintf("cluster%d", i)
		routeName := fmt.Sprintf("route%d", i)
		// listener name must be same since ports are shared and previous listener is drained
		listenerName := "listener"

		endpoint := resource.MakeEndpoint(clusterName, uint32(upstreamPort))
		cluster := resource.MakeCluster(ads, clusterName)
		route := resource.MakeRoute(routeName, clusterName)
		listener := resource.MakeListener(ads, listenerName, uint32(listenPort), routeName)

		glog.Infof("updating cache with %d-labelled responses", i)
		snapshot := cache.NewSnapshot(version,
			[]proto.Message{endpoint},
			[]proto.Message{cluster},
			[]proto.Message{route},
			[]proto.Message{listener})
		config.SetSnapshot(cache.Key(node), snapshot)

		select {
		case <-time.After(interval):
		case <-ctx.Done():
			return
		}
		i++
	}
}

// CheckResponse makes a request to localhost at the given port and checks that the response body matches.
func CheckResponse(port uint) error {
	glog.Infof("making a request to :%d", port)
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	req, err := client.Get(fmt.Sprintf("http://localhost:%d", port))
	if err != nil {
		return err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	if string(body) != hello {
		return fmt.Errorf("unexpected return %q", string(body))
	}
	return nil
}
