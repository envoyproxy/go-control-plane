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

package cache_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource"
)

type group struct{}

const (
	key = "node"
)

func (group) ID(node *core.Node) string {
	if node != nil {
		return node.Id
	}
	return key
}

var (
	version  = "x"
	version2 = "y"

	snapshot = cache.NewSnapshot(version,
		[]cache.Resource{endpoint},
		[]cache.Resource{cluster},
		[]cache.Resource{route},
		[]cache.Resource{listener})

	names = map[string][]string{
		cache.EndpointType: []string{clusterName},
		cache.ClusterType:  nil,
		cache.RouteType:    []string{routeName},
		cache.ListenerType: nil,
	}
)

type logger struct {
	t *testing.T
}

func (log logger) Infof(format string, args ...interface{}) {
	log.t.Logf(format, args...)
}

func TestSnapshotCache(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	if err := c.SetSnapshot(key, snapshot); err != nil {
		t.Fatal(err)
	}

	// try to get endpoints with incorrect list of names
	// should not receive response
	value, _ := c.CreateWatch(v2.DiscoveryRequest{TypeUrl: cache.EndpointType, ResourceNames: []string{"none"}})
	select {
	case out := <-value:
		t.Errorf("watch for endpoints and mismatched names => got %v, want none", out)
	case <-time.After(time.Second / 4):
	}

	for _, typ := range cache.ResponseTypes {
		t.Run(typ, func(t *testing.T) {
			value, _ := c.CreateWatch(v2.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]})
			select {
			case out := <-value:
				if out.Version != version {
					t.Errorf("got version %q, want %q", out.Version, version)
				}
				if !reflect.DeepEqual(out.Resources, snapshot.GetResources(typ)) {
					t.Errorf("get resources %v, want %v", out.Resources, snapshot.GetResources(typ))
				}
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}
}

func TestSnapshotCacheFetch(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	if err := c.SetSnapshot(key, snapshot); err != nil {
		t.Fatal(err)
	}

	for _, typ := range cache.ResponseTypes {
		t.Run(typ, func(t *testing.T) {
			resp, err := c.Fetch(context.Background(), v2.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]})
			if err != nil || resp == nil {
				t.Fatal("unexpected error or null response")
			}
			if resp.Version != version {
				t.Errorf("got version %q, want %q", resp.Version, version)
			}
		})
	}

	// no response for missing snapshot
	if resp, err := c.Fetch(context.Background(),
		v2.DiscoveryRequest{TypeUrl: cache.ClusterType, Node: &core.Node{Id: "oof"}}); resp != nil || err == nil {
		t.Errorf("missing snapshot: response is not nil %q", resp)
	}

	// no response for latest version
	if resp, err := c.Fetch(context.Background(),
		v2.DiscoveryRequest{TypeUrl: cache.ClusterType, VersionInfo: version}); resp != nil || err == nil {
		t.Errorf("latest version: response is not nil %q", resp)
	}

}

func TestSnapshotCacheWatch(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	watches := make(map[string]chan cache.Response)
	for _, typ := range cache.ResponseTypes {
		watches[typ], _ = c.CreateWatch(v2.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]})
	}
	if err := c.SetSnapshot(key, snapshot); err != nil {
		t.Fatal(err)
	}
	for _, typ := range cache.ResponseTypes {
		t.Run(typ, func(t *testing.T) {
			select {
			case out := <-watches[typ]:
				if out.Version != version {
					t.Errorf("got version %q, want %q", out.Version, version)
				}
				if !reflect.DeepEqual(out.Resources, snapshot.GetResources(typ)) {
					t.Errorf("get resources %v, want %v", out.Resources, snapshot.GetResources(typ))
				}
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}

	// open new watches with the latest version
	for _, typ := range cache.ResponseTypes {
		watches[typ], _ = c.CreateWatch(v2.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ], VersionInfo: version})
	}
	if count := len(c.GetStatusInfo(key).ResponseWatches); count != len(cache.ResponseTypes) {
		t.Errorf("watches should be created for the latest version: %d", count)
	}

	// set partially-versioned snapshot
	snapshot2 := snapshot
	snapshot2.Endpoints = cache.Resources{
		Version: version2,
		Items:   []cache.Resource{resource.MakeEndpoint(clusterName, 9090)},
	}
	if err := c.SetSnapshot(key, snapshot2); err != nil {
		t.Fatal(err)
	}
	if count := len(c.GetStatusInfo(key).ResponseWatches); count != len(cache.ResponseTypes)-1 {
		t.Errorf("watches should be preserved for all but one: %d", count)
	}

	// validate response for endpoints
	select {
	case out := <-watches[cache.EndpointType]:
		if out.Version != version2 {
			t.Errorf("got version %q, want %q", out.Version, version2)
		}
		if !reflect.DeepEqual(out.Resources, snapshot2.Endpoints.Items) {
			t.Errorf("get resources %v, want %v", out.Resources, snapshot2.Endpoints.Items)
		}
	case <-time.After(time.Second):
		t.Fatal("failed to receive snapshot response")
	}
}

func TestSnapshotCacheWatchCancel(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	for _, typ := range cache.ResponseTypes {
		_, cancel := c.CreateWatch(v2.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]})
		cancel()
	}
	for _, typ := range cache.ResponseTypes {
		if count := len(c.GetStatusInfo(key).ResponseWatches); count > 0 {
			t.Errorf("watches should be released for %s", typ)
		}
	}

	if empty := c.GetStatusInfo("missing"); empty != nil {
		t.Errorf("missing node key")
	}
}
