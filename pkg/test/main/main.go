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
	"context"
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/envoyproxy/go-control-plane/pkg/test"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/glog"
)

var (
	upstreamPort  uint
	listenPort    uint
	xdsPort       uint
	interval      time.Duration
	ads           bool
	bootstrapFile string
	envoyBinary   string
)

func init() {
	flag.UintVar(&upstreamPort, "upstream", 18080, "Upstream HTTP/1.1 port")
	flag.UintVar(&listenPort, "listen", 9000, "Listener port")
	flag.UintVar(&xdsPort, "xds", 18000, "xDS server port")
	flag.DurationVar(&interval, "interval", 10*time.Second, "Interval between cache refresh")
	flag.BoolVar(&ads, "ads", true, "Use ADS instead of separate xDS services")
	flag.StringVar(&bootstrapFile, "bootstrap", "bootstrap.json", "Bootstrap file name")
	flag.StringVar(&envoyBinary, "envoy", "envoy", "Envoy binary file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// start upstream
	go test.RunHTTP(ctx, upstreamPort)

	// create a cache
	config := cache.NewSnapshotCache(true, test.Hasher{}, nil)

	// update the cache at a regular interval
	go test.RunCacheUpdate(ctx, config, ads, interval, upstreamPort, listenPort)

	// start the xDS server
	go test.RunXDS(ctx, config, xdsPort)

	// write bootstrap file
	bootstrap := resource.MakeBootstrap(ads, uint32(xdsPort), 19000)
	buf := &bytes.Buffer{}
	if err := (&jsonpb.Marshaler{OrigName: true}).Marshal(buf, bootstrap); err != nil {
		glog.Fatal(err)
	}
	if err := ioutil.WriteFile(bootstrapFile, buf.Bytes(), 0644); err != nil {
		glog.Fatal(err)
	}

	// start envoy
	envoy := exec.Command("envoy", "-c", bootstrapFile, "--drain-time-s", "1")
	envoy.Stdout = os.Stdout
	envoy.Stderr = os.Stderr
	envoy.Start()

	for {
		if err := test.CheckResponse(listenPort); err != nil {
			glog.Errorf("ERROR %v", err)
		} else {
			glog.Info("OK")
		}

		time.Sleep(1 * time.Second)
	}
}
