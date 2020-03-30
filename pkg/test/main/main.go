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

// Package main contains the test driver for testing xDS manually.
package main

import (
	"context"
	cryptotls "crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	cachev2 "github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	serverv2 "github.com/envoyproxy/go-control-plane/pkg/server/v2"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	testv2 "github.com/envoyproxy/go-control-plane/pkg/test/v2"
	testv3 "github.com/envoyproxy/go-control-plane/pkg/test/v3"
	"google.golang.org/grpc"

	gcplogger "github.com/envoyproxy/go-control-plane/pkg/log"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource/v2"
)

const (
	// Hello is the echo message
	hello = "Hi, there!\n"

	grpcMaxConcurrentStreams = 1000000
)

type echo struct{}

var (
	debug bool

	port         uint
	gatewayPort  uint
	upstreamPort uint
	basePort     uint
	alsPort      uint

	delay    time.Duration
	requests int
	updates  int

	mode          string
	clusters      int
	httpListeners int
	tcpListeners  int
	runtimes      int
	tls           bool

	nodeID string
)

func init() {
	flag.BoolVar(&debug, "debug", false, "Use debug logging")
	flag.UintVar(&port, "port", 18000, "Management server port")
	flag.UintVar(&gatewayPort, "gateway", 18001, "Management server port for HTTP gateway")
	flag.UintVar(&upstreamPort, "upstream", 18080, "Upstream HTTP/1.1 port")
	flag.UintVar(&basePort, "base", 9000, "Listener port")
	flag.UintVar(&alsPort, "als", 18090, "Accesslog server port")
	flag.DurationVar(&delay, "delay", 500*time.Millisecond, "Interval between request batch retries")
	flag.IntVar(&requests, "r", 5, "Number of requests between snapshot updates")
	flag.IntVar(&updates, "u", 3, "Number of snapshot updates")
	flag.StringVar(&mode, "xds", resource.Ads, "Management server type (ads, xds, rest)")
	flag.IntVar(&clusters, "clusters", 4, "Number of clusters")
	flag.IntVar(&httpListeners, "http", 2, "Number of HTTP listeners (and RDS configs)")
	flag.IntVar(&tcpListeners, "tcp", 2, "Number of TCP pass-through listeners")
	flag.IntVar(&runtimes, "runtimes", 1, "Number of RTDS layers")
	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")
	flag.BoolVar(&tls, "tls", false, "Enable TLS on all listeners and use SDS for secret delivery")
}

// main returns code 1 if any of the batches failed to pass all requests
func main() {
	flag.Parse()
	ctx := context.Background()

	// start upstream
	go runHTTP(ctx, upstreamPort)

	// create a cache
	signal := make(chan struct{})
	cbv2 := &testv2.Callbacks{Signal: signal, Debug: debug}
	cbv3 := &testv3.Callbacks{Signal: signal, Debug: debug}

	configv2 := cachev2.NewSnapshotCache(mode == resource.Ads, cachev2.IDHash{}, logger{})
	configv3 := cachev3.NewSnapshotCache(mode == resource.Ads, cachev3.IDHash{}, logger{})
	srv2 := serverv2.NewServer(context.Background(), configv2, cbv2)
	srv3 := serverv3.NewServer(context.Background(), configv3, cbv3)
	alsv2 := &testv2.AccessLogService{}
	alsv3 := &testv3.AccessLogService{}

	// create a test snapshot
	snapshots := resource.TestSnapshot{
		Xds:              mode,
		UpstreamPort:     uint32(upstreamPort),
		BasePort:         uint32(basePort),
		NumClusters:      clusters,
		NumHTTPListeners: httpListeners,
		NumTCPListeners:  tcpListeners,
		TLS:              tls,
		NumRuntimes:      runtimes,
	}

	// start the xDS server
	go runAccessLogServer(ctx, alsv2, alsv3)
	go runManagementServer(ctx, srv2, srv3)
	go runManagementGateway(ctx, srv2, srv3, gatewayPort, logger{})

	log.Println("waiting for the first request...")
	select {
	case <-signal:
		break
	case <-time.After(1 * time.Minute):
		log.Println("timeout waiting for the first request")
		os.Exit(1)
	}
	log.Printf("initial snapshot %+v\n", snapshots)
	log.Printf("executing sequence updates=%d request=%d\n", updates, requests)

	for i := 0; i < updates; i++ {
		snapshots.Version = fmt.Sprintf("v%d", i)
		log.Printf("update snapshot %v\n", snapshots.Version)

		snapshot := snapshots.Generate()
		if err := snapshot.Consistent(); err != nil {
			log.Printf("snapshot inconsistency: %+v\n", snapshot)
		}

		err := configv2.SetSnapshot(nodeID, snapshot)
		if err != nil {
			log.Printf("snapshot error %q for %+v\n", err, snapshot)
			os.Exit(1)
		}

		// pass is true if all requests succeed at least once in a run
		pass := false
		for j := 0; j < requests; j++ {
			ok, failed := callEcho()
			if failed == 0 && !pass {
				pass = true
			}
			log.Printf("request batch %d, ok %v, failed %v, pass %v\n", j, ok, failed, pass)
			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return
			}
		}

		alsv2.Dump(func(s string) {
			if debug {
				log.Println(s)
			}
		})
		cbv2.Report()

		if !pass {
			log.Printf("failed all requests in a run %d\n", i)
			os.Exit(1)
		}
	}

	log.Printf("Test for %s passed!\n", mode)
}

// callEcho calls upstream echo service on all listener ports and returns an error
// if any of the listeners returned an error.
func callEcho() (int, int) {
	total := httpListeners + tcpListeners
	ok, failed := 0, 0
	ch := make(chan error, total)

	// spawn requests
	for i := 0; i < total; i++ {
		go func(i int) {
			client := http.Client{
				Timeout: 100 * time.Millisecond,
				Transport: &http.Transport{
					TLSClientConfig: &cryptotls.Config{InsecureSkipVerify: true},
				},
			}
			proto := "http"
			if tls {
				proto = "https"
			}
			req, err := client.Get(fmt.Sprintf("%s://127.0.0.1:%d", proto, basePort+uint(i)))
			if err != nil {
				ch <- err
				return
			}
			defer req.Body.Close()
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				ch <- err
				return
			}
			if string(body) != hello {
				ch <- fmt.Errorf("unexpected return %q", string(body))
				return
			}
			ch <- nil
		}(i)
	}

	for {
		out := <-ch
		if out == nil {
			ok++
		} else {
			failed++
		}
		if ok+failed == total {
			return ok, failed
		}
	}
}

type logger struct{}

func (logger logger) Debugf(format string, args ...interface{}) {
	if debug {
		log.Printf(format+"\n", args...)
	}
}

func (logger logger) Infof(format string, args ...interface{}) {
	if debug {
		log.Printf(format+"\n", args...)
	}
}

func (logger logger) Warnf(format string, args ...interface{}) {
	log.Printf(format+"\n", args...)
}

func (logger logger) Errorf(format string, args ...interface{}) {
	log.Printf(format+"\n", args...)
}

func runAccessLogServer(ctx context.Context, alsv2 *testv2.AccessLogService, alsv3 *testv3.AccessLogService) {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", alsPort))
	if err != nil {
		log.Fatal(err)
	}

	testv2.RegisterAccessLogServer(grpcServer, alsv2)
	log.Printf("access log server listening on %d\n", alsPort)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Println(err)
		}
	}()
	<-ctx.Done()

	grpcServer.GracefulStop()
}

// RunManagementServer starts an xDS server at the given port.
func runManagementServer(ctx context.Context, srv2 serverv2.Server, srv3 serverv3.Server) {
	// gRPC golang library sets a very small upper bound for the number gRPC/h2
	// streams over a single TCP connection. If a proxy multiplexes requests over
	// a single connection to the management server, then it might lead to
	// availability problems.
	var grpcOptions []grpc.ServerOption
	grpcOptions = append(grpcOptions, grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams))
	grpcServer := grpc.NewServer(grpcOptions...)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	testv2.RegisterServer(grpcServer, srv2)
	testv3.RegisterServer(grpcServer, srv3)

	log.Printf("management server listening on %d\n", port)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Println(err)
		}
	}()
	<-ctx.Done()

	grpcServer.GracefulStop()
}

// RunManagementGateway starts an HTTP gateway to an xDS server.
func runManagementGateway(ctx context.Context, srv2 serverv2.Server, srv3 serverv3.Server, port uint, lg gcplogger.Logger) {
	log.Printf("gateway listening HTTP/1.1 on %d\n", port)
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: &HTTPGateway{
			GatewayV2: serverv2.HTTPGateway{Log: lg, Server: srv2},
			GatewayV3: serverv3.HTTPGateway{Log: lg, Server: srv3},
			Log:       lg,
		},
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}

// HTTPGateway is a custom implementation of [gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)
// specialized to Envoy xDS API.
type HTTPGateway struct {
	// Log is an optional log for errors in response write
	Log gcplogger.Logger

	GatewayV2 serverv2.HTTPGateway

	GatewayV3 serverv3.HTTPGateway
}

func (h *HTTPGateway) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	err := h.GatewayV2.ServeHTTP(resp, req)
	if err != nil {
		h.GatewayV3.ServeHTTP(resp, req)
	}
}

func (h echo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	if _, err := w.Write([]byte(hello)); err != nil {
		log.Println(err)
	}
}

// RunHTTP opens a simple listener on the port.
func runHTTP(ctx context.Context, upstreamPort uint) {
	log.Printf("upstream listening HTTP/1.1 on %d\n", upstreamPort)
	server := &http.Server{Addr: fmt.Sprintf(":%d", upstreamPort), Handler: echo{}}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}
