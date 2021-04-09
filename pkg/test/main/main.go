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
	"net/http"
	"os"
	"time"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/envoyproxy/go-control-plane/pkg/test"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource/v3"
	testv3 "github.com/envoyproxy/go-control-plane/pkg/test/v3"
)

var (
	debug bool

	port            uint
	gatewayPort     uint
	upstreamPort    uint
	upstreamMessage string
	basePort        uint
	alsPort         uint

	delay    time.Duration
	requests int
	updates  int

	mode          string
	clusters      int
	httpListeners int
	tcpListeners  int
	runtimes      int
	tls           bool
	mux           bool

	nodeID string
)

func init() {
	flag.BoolVar(&debug, "debug", false, "Use debug logging")

	//
	// These parameters control the ports that the integration test
	// components use to talk to one another
	//

	// The port that the Envoy xDS client uses to talk to the control
	// plane xDS server (part of this program)
	flag.UintVar(&port, "port", 18000, "xDS management server port")

	// The port that the Envoy REST client uses to talk to the control
	// plane gateway (which translates from REST to xDS)
	flag.UintVar(&gatewayPort, "gateway", 18001, "Management HTTP gateway (from HTTP to xDS) server port")

	// The port that Envoy uses to talk to the upstream http "echo"
	// server
	flag.UintVar(&upstreamPort, "upstream", 18080, "Upstream HTTP/1.1 port")

	// The port that the tests below use to talk to Envoy's proxy of the
	// upstream server
	flag.UintVar(&basePort, "base", 9000, "Envoy Proxy listener port")

	// The control plane accesslog server port (currently unused)
	flag.UintVar(&alsPort, "als", 18090, "Control plane accesslog server port")

	//
	// These parameters control Envoy configuration
	//

	// Tell Envoy to request configurations from the control plane using
	// this protocol
	flag.StringVar(&mode, "xds", resource.Ads, "Management protocol to test (ADS, xDS, REST)")

	// Tell Envoy to use this Node ID
	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")

	// Tell Envoy to use TLS to talk to the control plane
	flag.BoolVar(&tls, "tls", false, "Enable TLS on all listeners and use SDS for secret delivery")

	// Tell Envoy to configure this many clusters for each snapshot
	flag.IntVar(&clusters, "clusters", 4, "Number of clusters")

	// Tell Envoy to configure this many Runtime Discovery Service
	// layers for each snapshot
	flag.IntVar(&runtimes, "runtimes", 1, "Number of RTDS layers")

	//
	// These parameters control the test harness
	//

	// The message that the tests expect to receive from the upstream
	// server
	flag.StringVar(&upstreamMessage, "message", "Default message", "Upstream HTTP server response message")

	// Time to wait between test request batches
	flag.DurationVar(&delay, "delay", 500*time.Millisecond, "Interval between request batch retries")

	// Each test loads a configuration snapshot into the control plane
	// which is then picked up by Envoy.  This parameter specifies how
	// many snapshots to test
	flag.IntVar(&updates, "u", 3, "Number of snapshot updates")

	// Each snapshot test sends this many requests to the upstream
	// server for each snapshot for each listener port
	flag.IntVar(&requests, "r", 5, "Number of requests between snapshot updates")

	// Test this many HTTP listeners per snapshot
	flag.IntVar(&httpListeners, "http", 2, "Number of HTTP listeners (and RDS configs)")
	// Test this many TCP listeners per snapshot
	flag.IntVar(&tcpListeners, "tcp", 2, "Number of TCP pass-through listeners")

	// Enable a muxed cache with partial snapshots
	flag.BoolVar(&mux, "mux", false, "Enable muxed linear cache for EDS")
}

// main returns code 1 if any of the batches failed to pass all requests
func main() {
	flag.Parse()
	ctx := context.Background()

	// create a cache
	signal := make(chan struct{})
	cb := &testv3.Callbacks{Signal: signal, Debug: debug}

	// mux integration
	config := cache.NewSnapshotCache(mode == resource.Ads, cache.IDHash{}, logger{})
	var configCache cache.Cache = config
	typeURL := "type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment"
	eds := cache.NewLinearCache(typeURL)
	if mux {
		configCache = &cache.MuxCache{
			Classify: func(req cache.Request) string {
				if req.TypeUrl == typeURL {
					return "eds"
				}
				return "default"
			},
			Caches: map[string]cache.Cache{
				"default": config,
				"eds":     eds,
			},
		}
	}
	srv := server.NewServer(context.Background(), configCache, cb)
	als := &testv3.AccessLogService{}

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
	go test.RunAccessLogServer(ctx, als, alsPort)
	go test.RunManagementServer(ctx, srv, port)
	go test.RunManagementGateway(ctx, srv, gatewayPort, logger{})

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

		err := config.SetSnapshot(nodeID, snapshot)
		if err != nil {
			log.Printf("snapshot error %q for %+v\n", err, snapshot)
			os.Exit(1)
		}

		if mux {
			for name, res := range snapshot.GetResources(typeURL) {
				eds.UpdateResource(name, res)
			}
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

		als.Dump(func(s string) {
			if debug {
				log.Println(s)
			}
		})
		cb.Report()

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
			if string(body) != upstreamMessage {
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
