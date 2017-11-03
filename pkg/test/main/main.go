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
	"time"

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
	interval     time.Duration
	node         string
)

func init() {
	flag.UintVar(&upstreamPort, "upstream", 18080, "Upstream HTTP/1.1 port")
	flag.UintVar(&listenPort, "listen", 9000, "Listener port")
	flag.UintVar(&xdsPort, "xds", 18000, "xDS server port")
	flag.DurationVar(&interval, "interval", 30*time.Second, "Interval between cache refresh")
	flag.StringVar(&node, "node", "node", "Node group name")
}

func main() {
	flag.Parse()
	// start upstream
	go runHTTP()

	// create a cache
	config := cache.NewSimpleCache(hasher{}, nil)

	// update the cache at a regular interval
	go func() {
		i := 0
		for {
			version := fmt.Sprintf("version%d", i)
			clusterName := fmt.Sprintf("cluster%d", i)
			routeName := fmt.Sprintf("route%d", i)
			// listener name must be same since ports are shared and previous listener is drained
			listenerName := "listener"

			endpoint := test.MakeEndpoint(clusterName, uint32(upstreamPort))
			cluster := test.MakeCluster(clusterName)
			route := test.MakeRoute(routeName, clusterName)
			listener := test.MakeListener(listenerName, uint32(listenPort), routeName)

			log.Printf("updating cache with %d-labelled responses", i)
			snapshot := cache.NewSnapshot(version,
				[]proto.Message{endpoint},
				[]proto.Message{cluster},
				[]proto.Message{route},
				[]proto.Message{listener})
			config.SetSnapshot(cache.Key(node), snapshot)

			time.Sleep(interval)
			i++
		}
	}()

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

func (hash hasher) Hash(*api.Node) (cache.Key, error) {
	return cache.Key(node), nil
}

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request from %q...", r.RemoteAddr)
	body := bytes.Buffer{}
	w.Header().Set("Content-Type", "application/text")
	if _, err := w.Write(body.Bytes()); err != nil {
		log.Println(err.Error())
	}
}

func runHTTP() {
	log.Printf("upstream listening HTTP1.1 on %d", upstreamPort)
	h := handler{}
	if err := http.ListenAndServe(fmt.Sprintf(":%d", upstreamPort), h); err != nil {
		log.Println(err.Error())
	}
}
