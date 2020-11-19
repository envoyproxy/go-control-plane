package cache_test

import (
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

type mockStreamState struct {
	SystemVersion    string
	Nonce            string
	ResourceVersions map[string]cache.DeltaVersionInfo
}

func (s mockStreamState) GetVersionMap() map[string]cache.DeltaVersionInfo {
	return s.ResourceVersions
}
func (s mockStreamState) GetStreamNonce() string {
	return s.Nonce
}
func (s mockStreamState) GetSystemVersion() string {
	return s.SystemVersion
}

func TestSnapshotCacheDeltaWatch(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, logger{t: t})
	watches := make(map[string]chan cache.DeltaResponse)

	for _, typ := range testTypes {
		watches[typ] = make(chan cache.DeltaResponse, 1)
		c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: "node",
			},
			TypeUrl:                typ,
			ResourceNamesSubscribe: names[typ],
		}, watches[typ], mockStreamState{ResourceVersions: nil, SystemVersion: ""})
	}

	if err := c.SetSnapshot(key, snapshot); err != nil {
		t.Fatal(err)
	}

	vm := make(map[string]map[string]cache.DeltaVersionInfo)
	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			select {
			case out := <-watches[typ]:
				if !reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshot.GetResources(typ)) {
					t.Errorf("got resources %v, want %v", out.(*cache.RawDeltaResponse).Resources, snapshot.GetResources(typ))
				}
				vMap, err := out.GetDeltaVersionMap()
				if err != nil {
					t.Fatal(err)
				}
				vm[typ] = vMap
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}

	for _, typ := range testTypes {
		watches[typ] = make(chan cache.DeltaResponse, 1)
		c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: "node",
			},
			TypeUrl:                typ,
			ResourceNamesSubscribe: names[typ],
		}, watches[typ], mockStreamState{ResourceVersions: vm[typ], SystemVersion: "x"})
	}

	if count := c.GetStatusInfo(key).GetNumDeltaWatches(); count != len(testTypes) {
		t.Errorf("watches should be created for the latest version %d", count)
	}

	// set partially-versioned snapshot
	snapshot2 := snapshot
	snapshot2.Resources[types.Endpoint] = cache.NewResources(version2, []types.Resource{resource.MakeEndpoint(clusterName, 9090)})
	if err := c.SetSnapshot(key, snapshot2); err != nil {
		t.Fatal(err)
	}
	if count := c.GetStatusInfo(key).GetNumDeltaWatches(); count != len(testTypes)-1 {
		t.Errorf("watches should be preserved for all but one: %d", count)
	}

	// validate response for endpoints
	select {
	case out := <-watches[testTypes[0]]:
		if !reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshot2.Resources[types.Endpoint].Items) {
			t.Fatalf("got resources %v, want %v", out.(*cache.RawDeltaResponse).Resources, snapshot2.Resources[types.Endpoint].Items)
		}
		vMap, err := out.GetDeltaVersionMap()
		if err != nil {
			t.Fatal(err)
		}
		vm[testTypes[0]] = vMap
	case <-time.After(time.Second):
		t.Fatal("failed to receive snapshot response")
	}
}

func TestConcurrentSetDeltaWatch(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, logger{t: t})
	for i := 0; i < 50; i++ {
		version := fmt.Sprintf("v%d", i)
		func(i int) {
			t.Run(fmt.Sprintf("worker%d", i), func(t *testing.T) {
				t.Parallel()
				id := fmt.Sprintf("%d", i%2)
				var cancel func()
				value := make(chan cache.DeltaResponse, 1)
				if i < 25 {
					snap := cache.Snapshot{}
					snap.Resources[types.Endpoint] = cache.NewResources(version, []types.Resource{resource.MakeEndpoint(clusterName, uint32(i))})
					c.SetSnapshot(id, snap)
				} else {
					if cancel != nil {
						cancel()
					}
					c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
						Node: &core.Node{
							Id: id,
						},
						TypeUrl:                rsrc.EndpointType,
						ResourceNamesSubscribe: []string{clusterName},
					}, value, mockStreamState{ResourceVersions: nil, SystemVersion: "x"})
				}
			})
		}(i)
	}
}

func TestSnapshotCacheDeltaWatchCancel(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	for _, typ := range testTypes {
		value := make(chan cache.DeltaResponse, 1)
		cancel := c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: key,
			},
			TypeUrl:                typ,
			ResourceNamesSubscribe: names[typ],
		}, value, mockStreamState{ResourceVersions: nil, SystemVersion: "x"})

		// Cancel the watch
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
