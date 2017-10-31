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

// Package main contains the test driver for testing xDS manually.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/envoyproxy/go-control-plane/api"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	xds "github.com/envoyproxy/go-control-plane/pkg/grpc"
	"github.com/envoyproxy/go-control-plane/pkg/test"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

var (
	upstreamPort uint
	listenPort   uint
	xdsPort      uint
)

func init() {
	flag.UintVar(&upstreamPort, "upstream", 18080, "Upstream HTTP/1.1 port")
	flag.UintVar(&listenPort, "listen", 9000, "Listener port")
	flag.UintVar(&xdsPort, "xds", 18000, "xDS server port")
}

func main() {
	flag.Parse()
	// start upstream
	go runHTTP()

	// create a cache
	config := cache.NewSimpleCache(hasher{}, nil)

	// populate the cache
	endpoint := test.MakeEndpoint("cluster", uint32(upstreamPort))
	cluster := test.MakeCluster("cluster")
	route := test.MakeRoute("route", "cluster")
	listener := test.MakeListener("listener", uint32(listenPort), "route")

	config.SetResource("", cache.EndpointResponse, []proto.Message{endpoint})
	config.SetResource("", cache.ClusterResponse, []proto.Message{cluster})
	config.SetResource("", cache.RouteResponse, []proto.Message{route})
	config.SetResource("", cache.ListenerResponse, []proto.Message{listener})

	// start the xDS server
	server := xds.NewServer(config)
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", xdsPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server.Register(grpcServer)
	log.Printf("xDS server listening on %d", xdsPort)
	if err = grpcServer.Serve(lis); err != nil {
		log.Println(err.Error())
	}
}

type hasher struct {
}

func (hash hasher) Hash(*api.Node) cache.Key {
	return ""
}

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %q...", r.RemoteAddr)
	body := bytes.Buffer{}
	w.Header().Set("Content-Type", "application/text")
	if _, err := w.Write(body.Bytes()); err != nil {
		log.Println(err.Error())
	}
}

func runHTTP() {
	log.Printf("Upstream listening HTTP1.1 on %d", upstreamPort)
	h := handler{}
	if err := http.ListenAndServe(fmt.Sprintf(":%d", upstreamPort), h); err != nil {
		log.Println(err.Error())
	}
}
