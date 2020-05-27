// Code generated by create_version. DO NOT EDIT.

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
	"fmt"
	"reflect"
	"testing"
	"time"

	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v2"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource/v2"
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
		[]types.Resource{testEndpoint, testEndpoint1, testEndpoint2},
		[]types.Resource{testCluster},
		[]types.Resource{testRoute},
		[]types.Resource{testListener},
		[]types.Resource{testRuntime})

	names = map[string][]string{
		rsrc.EndpointType: {clusterName, "cluster1", "clusterDelta"},
		rsrc.ClusterType:  {clusterName},
		rsrc.RouteType:    {routeName},
		rsrc.ListenerType: {listenerName},
		rsrc.RuntimeType:  nil,
	}

	testTypes = []string{
		rsrc.EndpointType,
		rsrc.ClusterType,
		rsrc.RouteType,
		rsrc.ListenerType,
		rsrc.RuntimeType,
	}
)

type logger struct {
	t *testing.T
}

func (log logger) Debugf(format string, args ...interface{}) { log.t.Logf(format, args...) }
func (log logger) Infof(format string, args ...interface{})  { log.t.Logf(format, args...) }
func (log logger) Warnf(format string, args ...interface{})  { log.t.Logf(format, args...) }
func (log logger) Errorf(format string, args ...interface{}) { log.t.Logf(format, args...) }

func TestSnapshotCache(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})

	if _, err := c.GetSnapshot(key); err == nil {
		t.Errorf("unexpected snapshot found for key %q", key)
	}

	if err := c.SetSnapshot(key, snapshot); err != nil {
		t.Fatal(err)
	}

	snap, err := c.GetSnapshot(key)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(snap, snapshot) {
		t.Errorf("expect snapshot: %v, got: %v", snapshot, snap)
	}

	// try to get endpoints with incorrect list of names
	// should not receive response
	value, _ := c.CreateWatch(discovery.DiscoveryRequest{TypeUrl: rsrc.EndpointType, ResourceNames: []string{"none"}})
	select {
	case out := <-value:
		t.Errorf("watch for endpoints and mismatched names => got %v, want none", out)
	case <-time.After(time.Second / 4):
	}

	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			value, _ := c.CreateWatch(discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]})
			select {
			case out := <-value:
				if out.Version != version {
					t.Errorf("got version %q, want %q", out.Version, version)
				}
				if !reflect.DeepEqual(cache.IndexResourcesByName(out.Resources), snapshot.GetResources(typ)) {
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

	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			resp, err := c.Fetch(context.Background(), discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]})
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
		discovery.DiscoveryRequest{TypeUrl: rsrc.ClusterType, Node: &core.Node{Id: "oof"}}); resp != nil || err == nil {
		t.Errorf("missing snapshot: response is not nil %v", resp)
	}

	// no response for latest version
	if resp, err := c.Fetch(context.Background(),
		discovery.DiscoveryRequest{TypeUrl: rsrc.ClusterType, VersionInfo: version}); resp != nil || err == nil {
		t.Errorf("latest version: response is not nil %v", resp)
	}
}

func TestSnapshotCacheWatch(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	watches := make(map[string]chan cache.Response)
	for _, typ := range testTypes {
		watches[typ], _ = c.CreateWatch(discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]})
	}
	if err := c.SetSnapshot(key, snapshot); err != nil {
		t.Fatal(err)
	}
	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			select {
			case out := <-watches[typ]:
				if out.Version != version {
					t.Errorf("got version %q, want %q", out.Version, version)
				}
				if !reflect.DeepEqual(cache.IndexResourcesByName(out.Resources), snapshot.GetResources(typ)) {
					t.Errorf("get resources %v, want %v", out.Resources, snapshot.GetResources(typ))
				}
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}

	// open new watches with the latest version
	for _, typ := range testTypes {
		watches[typ], _ = c.CreateWatch(discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ], VersionInfo: version})
	}

	t.Logf("%+v\n", c.GetStatusInfo(key))

	if count := c.GetStatusInfo(key).GetNumWatches(); count != len(testTypes) {
		t.Errorf("watches should be created for the latest version: %d", count)
	}

	// set partially-versioned snapshot
	snapshot2 := snapshot
	snapshot2.Resources[types.Endpoint] = cache.NewResources(version2, []types.Resource{resource.MakeEndpoint(clusterName, 9090)})
	if err := c.SetSnapshot(key, snapshot2); err != nil {
		t.Fatal(err)
	}
	if count := c.GetStatusInfo(key).GetNumWatches(); count != len(testTypes)-1 {
		t.Errorf("watches should be preserved for all but one: %d", count)
	}

	// validate response for endpoints
	select {
	case out := <-watches[rsrc.EndpointType]:
		if out.Version != version2 {
			t.Errorf("got version %q, want %q", out.Version, version2)
		}
		if !reflect.DeepEqual(cache.IndexResourcesByName(out.Resources), snapshot2.Resources[types.Endpoint].Items) {
			t.Errorf("get resources %v, want %v", out.Resources, snapshot2.Resources[types.Endpoint].Items)
		}
	case <-time.After(time.Second):
		t.Fatal("failed to receive snapshot response")
	}
}

func TestSnapshotCacheDeltaWatch(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, logger{t: t})
	watches := make(map[string]chan cache.DeltaResponse)

	for _, typ := range testTypes {
		watches[typ], _ = c.CreateDeltaWatch(discovery.DeltaDiscoveryRequest{
			TypeUrl:                typ,
			ResourceNamesSubscribe: names[typ],
		}, "")
	}

	if err := c.SetSnapshotDelta(key, snapshot); err != nil {
		t.Fatal(err)
	}

	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			select {
			case out := <-watches[typ]:
				if !reflect.DeepEqual(cache.IndexResourcesByName(out.Resources), snapshot.GetSubscribedResources(names[typ], typ)) {
					t.Errorf("get resources %v, want %v", out.Resources, snapshot.GetSubscribedResources(names[typ], typ))
				}
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}

	// open new watches with the latest version
	watches[rsrc.EndpointType], _ = c.CreateDeltaWatch(discovery.DeltaDiscoveryRequest{
		TypeUrl:                rsrc.EndpointType,
		ResourceNamesSubscribe: []string{"cluster0"},
	}, version)

	if count := c.GetStatusInfo(key).GetNumDeltaWatches(); count != len(testTypes) {
		t.Errorf("watches should be created for the latest version: %d", count)
	}

	// set partially-versioned snapshot
	snapshot2 := snapshot
	snapshot2.Resources[types.Endpoint] = cache.NewResources(version2, []types.Resource{resource.MakeEndpoint(clusterName, 9090)})
	if err := c.SetSnapshotDelta(key, snapshot2); err != nil {
		t.Fatal(err)
	}
	if count := c.GetStatusInfo(key).GetNumDeltaWatches(); count != len(testTypes)-1 {
		t.Errorf("watches should be preserved for all but one: %d", count)
	}

	// validate response for endpoints
	select {
	case out := <-watches[rsrc.EndpointType]:
		if !reflect.DeepEqual(cache.IndexResourcesByName(out.Resources), snapshot2.Resources[types.Endpoint].Items) {
			t.Fatalf("got resources %v, want %v", out.Resources, snapshot2.Resources[types.Endpoint].Items)
		}
	case <-time.After(time.Second):
		t.Fatal("failed to receive snapshot response")
	}
}

func TestCheckState(t *testing.T) {
	deltaState := map[string][]string{
		rsrc.EndpointType: {clusterName, "cluster1", "clusterDelta"},
		rsrc.ClusterType:  {clusterName},
		rsrc.RouteType:    {routeName},
		rsrc.ListenerType: {listenerName},
		rsrc.RuntimeType:  nil,
	}
	subscribed := map[string][]string{
		rsrc.EndpointType: {clusterName, "cluster1", "clusterDelta", "clusterDelta2"},
		rsrc.ClusterType:  {clusterName},
		rsrc.RouteType:    {routeName},
		rsrc.ListenerType: {listenerName},
		rsrc.RuntimeType:  nil,
	}

	mb := make(map[string][]string, len(deltaState))
	diff := make(map[string][]string, 0)

	for key, value := range deltaState {
		mb[key] = value
	}

	// Check our diff map to see what has changed
	// Even is an underlying resource has changed we need to update the diff
	for key, value := range subscribed {
		t.Log(key)
		if _, found := mb[key]; !found {
			t.Log("Found a new key, adding to diff")
			diff[key] = value
		} else if resources, found := mb[key]; found && (len(resources) != len(value)) {
			t.Log("Found a new resource to an existing key, modifying resource map")
			diff[key] = value
		} else {
			t.Log("Found no changes")
		}
	}

	if len(diff) == 0 {
		t.Fatalf("Expected diff greater than 0")
	}
}

func TestConcurrentSetWatch(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, logger{t: t})
	for i := 0; i < 50; i++ {
		func(i int) {
			t.Run(fmt.Sprintf("worker%d", i), func(t *testing.T) {
				t.Parallel()
				id := fmt.Sprintf("%d", i%2)
				var cancel func()
				if i < 25 {
					snap := cache.Snapshot{}
					snap.Resources[types.Endpoint] = cache.NewResources(fmt.Sprintf("v%d", i), []types.Resource{resource.MakeEndpoint(clusterName, uint32(i))})
					c.SetSnapshot(id, snap)
				} else {
					if cancel != nil {
						cancel()
					}
					_, cancel = c.CreateWatch(discovery.DiscoveryRequest{
						Node:    &core.Node{Id: id},
						TypeUrl: rsrc.EndpointType,
					})
				}
			})
		}(i)
	}
}

func TestSnapshotCacheWatchCancel(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	for _, typ := range testTypes {
		_, cancel := c.CreateWatch(discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]})
		cancel()
	}
	// should be status info for the node
	if keys := c.GetStatusKeys(); len(keys) == 0 {
		t.Error("got 0, want status info for the node")
	}

	for _, typ := range testTypes {
		if count := c.GetStatusInfo(key).GetNumWatches(); count > 0 {
			t.Errorf("watches should be released for %s", typ)
		}
	}

	if empty := c.GetStatusInfo("missing"); empty != nil {
		t.Errorf("should not return a status for unknown key: got %#v", empty)
	}
}

func TestSnapshotClear(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	if err := c.SetSnapshot(key, snapshot); err != nil {
		t.Fatal(err)
	}
	c.ClearSnapshot(key)
	if empty := c.GetStatusInfo(key); empty != nil {
		t.Errorf("cache should be cleared")
	}
	if keys := c.GetStatusKeys(); len(keys) != 0 {
		t.Errorf("keys should be empty")
	}
}
