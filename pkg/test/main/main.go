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
	"flag"
	"time"

	"context"
	cryptotls "crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	cache "github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	server "github.com/envoyproxy/go-control-plane/pkg/server/v2"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	testv2 "github.com/envoyproxy/go-control-plane/pkg/test/v2"
	testv3 "github.com/envoyproxy/go-control-plane/pkg/test/v3"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	resource "github.com/envoyproxy/go-control-plane/pkg/test/resource/v2"
	resourcev3 "github.com/envoyproxy/go-control-plane/pkg/test/resource/v3"
)

// Flags type
type Flags struct {
	Debug bool

	Port         uint
	GatewayPort  uint
	UpstreamPort uint
	BasePort     uint
	AlsPort      uint

	Delay    time.Duration
	Requests int
	Updates  int

	Mode          string
	Clusters      int
	HTTPListeners int
	TCPListeners  int
	Runtimes      int
	TLS           bool

	NodeID string

	V2 bool
	V3 bool
}

var flags Flags

func init() {
	flags := Flags{}
	flag.BoolVar(&flags.Debug, "debug", false, "Use debug logging")
	flag.UintVar(&flags.Port, "port", 18000, "Management server port")
	flag.UintVar(&flags.GatewayPort, "gateway", 18001, "Management server port for HTTP gateway")
	flag.UintVar(&flags.UpstreamPort, "upstream", 18080, "Upstream HTTP/1.1 port")
	flag.UintVar(&flags.BasePort, "base", 9000, "Listener port")
	flag.UintVar(&flags.AlsPort, "als", 18090, "Accesslog server port")
	flag.DurationVar(&flags.Delay, "delay", 500*time.Millisecond, "Interval between request batch retries")
	flag.IntVar(&flags.Requests, "r", 5, "Number of requests between snapshot updates")
	flag.IntVar(&flags.Updates, "u", 3, "Number of snapshot updates")
	flag.StringVar(&flags.Mode, "xds", resource.Ads, "Management server type (ads, xds, rest)")
	flag.IntVar(&flags.Clusters, "clusters", 4, "Number of clusters")
	flag.IntVar(&flags.HTTPListeners, "http", 2, "Number of HTTP listeners (and RDS configs)")
	flag.IntVar(&flags.TCPListeners, "tcp", 2, "Number of TCP pass-through listeners")
	flag.IntVar(&flags.Runtimes, "runtimes", 1, "Number of RTDS layers")
	flag.StringVar(&flags.NodeID, "nodeID", "test-id", "Node ID")
	flag.BoolVar(&flags.V2, "v2", true, "Run v2 api tests")
	flag.BoolVar(&flags.V3, "v3", false, "Run v3 api tests")
}

// main returns code 1 if any of the batches failed to pass all requests
func main() {
	flag.Parse()
	if flags.V2 {
		testdriverv2 := TestDriverV2{Flags: flags}
		testdriverv2.Run()
	}

	if flags.V3 {
		testdriverv2 := TestDriverV3{Flags: flags}
		testdriverv2.Run()
	}
}

// TestDriver is defined here
type TestDriver interface {
	Run()
}

// TestDriverV2 is defined here
type TestDriverV2 struct {
	TestDriver
	Flags Flags
}

// TestDriverV3 is defined here
type TestDriverV3 struct {
	TestDriver
	Flags Flags
}

// Run is a method
func (d *TestDriverV2) Run() {
	ctx := context.Background()

	// start upstream
	go testv2.RunHTTP(ctx, d.Flags.UpstreamPort)

	// create a cache
	signal := make(chan struct{})
	cb := &callbacksv2{signal: signal, Debug: d.Flags.Debug}
	config := cache.NewSnapshotCache(d.Flags.Mode == resource.Ads, cache.IDHash{}, logger{Debug: d.Flags.Debug})
	srv := server.NewServer(context.Background(), config, cb)
	als := &testv2.AccessLogService{}

	// create a test snapshot
	snapshots := resource.TestSnapshot{
		Xds:              d.Flags.Mode,
		UpstreamPort:     uint32(d.Flags.UpstreamPort),
		BasePort:         uint32(d.Flags.BasePort),
		NumClusters:      d.Flags.Clusters,
		NumHTTPListeners: d.Flags.HTTPListeners,
		NumTCPListeners:  d.Flags.TCPListeners,
		TLS:              d.Flags.TLS,
		NumRuntimes:      d.Flags.Runtimes,
	}

	// start the xDS server
	go testv2.RunAccessLogServer(ctx, als, d.Flags.AlsPort)
	go testv2.RunManagementServer(ctx, srv, d.Flags.Port)
	go testv2.RunManagementGateway(ctx, srv, d.Flags.GatewayPort, logger{Debug: d.Flags.Debug})

	log.Println("waiting for the first request...")
	select {
	case <-signal:
		break
	case <-time.After(1 * time.Minute):
		log.Println("timeout waiting for the first request")
		os.Exit(1)
	}
	log.Printf("initial snapshot %+v\n", snapshots)
	log.Printf("executing sequence updates=%d request=%d\n", d.Flags.Updates, d.Flags.Requests)

	for i := 0; i < d.Flags.Updates; i++ {
		snapshots.Version = fmt.Sprintf("v%d", i)
		log.Printf("update snapshot %v\n", snapshots.Version)

		snapshot := snapshots.Generate()
		if err := snapshot.Consistent(); err != nil {
			log.Printf("snapshot inconsistency: %+v\n", snapshot)
		}

		err := config.SetSnapshot(d.Flags.NodeID, snapshot)
		if err != nil {
			log.Printf("snapshot error %q for %+v\n", err, snapshot)
			os.Exit(1)
		}

		// pass is true if all requests succeed at least once in a run
		pass := false
		for j := 0; j < d.Flags.Requests; j++ {
			ok, failed := d.callEcho()
			if failed == 0 && !pass {
				pass = true
			}
			log.Printf("request batch %d, ok %v, failed %v, pass %v\n", j, ok, failed, pass)
			select {
			case <-time.After(d.Flags.Delay):
			case <-ctx.Done():
				return
			}
		}

		als.Dump(func(s string) {
			if d.Flags.Debug {
				log.Println(s)
			}
		})
		cb.Report()

		if !pass {
			log.Printf("failed all requests in a run %d\n", i)
			os.Exit(1)
		}
	}

	log.Printf("Test for %s passed!\n", d.Flags.Mode)
}

// callEcho calls upstream echo service on all listener ports and returns an error
// if any of the listeners returned an error.
func (d *TestDriverV2) callEcho() (int, int) {
	total := d.Flags.HTTPListeners + d.Flags.TCPListeners
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
			if d.Flags.TLS {
				proto = "https"
			}
			req, err := client.Get(fmt.Sprintf("%s://127.0.0.1:%d", proto, d.Flags.BasePort+uint(i)))
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
			if string(body) != testv2.Hello {
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

// Run is a method
func (d *TestDriverV3) Run() {
	ctx := context.Background()

	// start upstream
	go testv3.RunHTTP(ctx, d.Flags.UpstreamPort)

	// create a cache
	signal := make(chan struct{})
	cb := &callbacksv3{signal: signal, Debug: d.Flags.Debug}
	config := cachev3.NewSnapshotCache(d.Flags.Mode == resource.Ads, cachev3.IDHash{}, logger{Debug: d.Flags.Debug})
	srv := serverv3.NewServer(context.Background(), config, cb)
	als := &testv3.AccessLogService{}

	// create a test snapshot
	snapshots := resourcev3.TestSnapshot{
		Xds:              d.Flags.Mode,
		UpstreamPort:     uint32(d.Flags.UpstreamPort),
		BasePort:         uint32(d.Flags.BasePort),
		NumClusters:      d.Flags.Clusters,
		NumHTTPListeners: d.Flags.HTTPListeners,
		NumTCPListeners:  d.Flags.TCPListeners,
		TLS:              d.Flags.TLS,
		NumRuntimes:      d.Flags.Runtimes,
	}

	// start the xDS server
	go testv3.RunAccessLogServer(ctx, als, d.Flags.AlsPort)
	go testv3.RunManagementServer(ctx, srv, d.Flags.Port)
	go testv3.RunManagementGateway(ctx, srv, d.Flags.GatewayPort, logger{Debug: d.Flags.Debug})

	log.Println("waiting for the first request...")
	select {
	case <-signal:
		break
	case <-time.After(1 * time.Minute):
		log.Println("timeout waiting for the first request")
		os.Exit(1)
	}
	log.Printf("initial snapshot %+v\n", snapshots)
	log.Printf("executing sequence updates=%d request=%d\n", d.Flags.Updates, d.Flags.Requests)

	for i := 0; i < d.Flags.Updates; i++ {
		snapshots.Version = fmt.Sprintf("v%d", i)
		log.Printf("update snapshot %v\n", snapshots.Version)

		snapshot := snapshots.Generate()
		if err := snapshot.Consistent(); err != nil {
			log.Printf("snapshot inconsistency: %+v\n", snapshot)
		}

		err := config.SetSnapshot(d.Flags.NodeID, snapshot)
		if err != nil {
			log.Printf("snapshot error %q for %+v\n", err, snapshot)
			os.Exit(1)
		}

		// pass is true if all requests succeed at least once in a run
		pass := false
		for j := 0; j < d.Flags.Requests; j++ {
			ok, failed := d.callEcho()
			if failed == 0 && !pass {
				pass = true
			}
			log.Printf("request batch %d, ok %v, failed %v, pass %v\n", j, ok, failed, pass)
			select {
			case <-time.After(d.Flags.Delay):
			case <-ctx.Done():
				return
			}
		}

		als.Dump(func(s string) {
			if d.Flags.Debug {
				log.Println(s)
			}
		})
		cb.Report()

		if !pass {
			log.Printf("failed all requests in a run %d\n", i)
			os.Exit(1)
		}
	}

	log.Printf("Test for %s passed!\n", d.Flags.Mode)
}

// callEcho calls upstream echo service on all listener ports and returns an error
// if any of the listeners returned an error.
func (d *TestDriverV3) callEcho() (int, int) {
	total := d.Flags.HTTPListeners + d.Flags.TCPListeners
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
			if d.Flags.TLS {
				proto = "https"
			}
			req, err := client.Get(fmt.Sprintf("%s://127.0.0.1:%d", proto, d.Flags.BasePort+uint(i)))
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
			if string(body) != testv3.Hello {
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

type logger struct {
	Debug bool
}

func (logger logger) Debugf(format string, args ...interface{}) {
	if logger.Debug {
		log.Printf(format+"\n", args...)
	}
}

func (logger logger) Infof(format string, args ...interface{}) {
	if logger.Debug {
		log.Printf(format+"\n", args...)
	}
}

func (logger logger) Warnf(format string, args ...interface{}) {
	log.Printf(format+"\n", args...)
}

func (logger logger) Errorf(format string, args ...interface{}) {
	log.Printf(format+"\n", args...)
}

type callbacksv2 struct {
	signal   chan struct{}
	fetches  int
	requests int
	mu       sync.Mutex
	Debug    bool
}

type callbacksv3 struct {
	signal   chan struct{}
	fetches  int
	requests int
	mu       sync.Mutex
	Debug    bool
}

func (cb *callbacksv2) Report() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	log.Printf("server callbacks fetches=%d requests=%d\n", cb.fetches, cb.requests)
}
func (cb *callbacksv2) OnStreamOpen(_ context.Context, id int64, typ string) error {
	if cb.Debug {
		log.Printf("stream %d open for %s\n", id, typ)
	}
	return nil
}
func (cb *callbacksv2) OnStreamClosed(id int64) {
	if cb.Debug {
		log.Printf("stream %d closed\n", id)
	}
}

func (cb *callbacksv2) OnStreamRequest(int64, *v2.DiscoveryRequest) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.requests++
	if cb.signal != nil {
		close(cb.signal)
		cb.signal = nil
	}
	return nil
}

func (cb *callbacksv2) OnStreamResponse(int64, *v2.DiscoveryRequest, *v2.DiscoveryResponse) {}

func (cb *callbacksv2) OnFetchRequest(_ context.Context, req *v2.DiscoveryRequest) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.fetches++
	if cb.signal != nil {
		close(cb.signal)
		cb.signal = nil
	}
	return nil
}

func (cb *callbacksv2) OnFetchResponse(*v2.DiscoveryRequest, *v2.DiscoveryResponse) {}

func (cb *callbacksv3) Report() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	log.Printf("server callbacks fetches=%d requests=%d\n", cb.fetches, cb.requests)
}
func (cb *callbacksv3) OnStreamOpen(_ context.Context, id int64, typ string) error {
	if cb.Debug {
		log.Printf("stream %d open for %s\n", id, typ)
	}
	return nil
}
func (cb *callbacksv3) OnStreamClosed(id int64) {
	if cb.Debug {
		log.Printf("stream %d closed\n", id)
	}
}

func (cb *callbacksv3) OnStreamRequest(int64, *discovery.DiscoveryRequest) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.requests++
	if cb.signal != nil {
		close(cb.signal)
		cb.signal = nil
	}
	return nil
}

func (cb *callbacksv3) OnStreamResponse(int64, *discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {
}

func (cb *callbacksv3) OnFetchRequest(_ context.Context, req *discovery.DiscoveryRequest) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.fetches++
	if cb.signal != nil {
		close(cb.signal)
		cb.signal = nil
	}
	return nil
}

func (cb *callbacksv3) OnFetchResponse(*discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {}
