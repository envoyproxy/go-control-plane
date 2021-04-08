package cache_test

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource"
)

func TestSnapshotCacheWithDeltaTtl(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := cache.NewSnapshotCacheWithHeartbeating(ctx, true, group{}, logger{t: t}, time.Second)

	if _, err := c.GetSnapshot(key); err == nil {
		t.Errorf("unexpected snapshot found for key %q", key)
	}

	if err := c.SetSnapshot(key, snapshotWithTtl); err != nil {
		t.Fatal(err)
	}

	snap, err := c.GetSnapshot(key)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(snap, snapshotWithTtl) {
		t.Errorf("expect snapshot: %v, got: %v", snapshotWithTtl, snap)
	}

	wg := sync.WaitGroup{}
	vm := make(map[string]map[string]string)
	// All the resources should respond immediately when version is not up to date.
	for _, typ := range testTypes {
		wg.Add(1)
		t.Run(typ, func(t *testing.T) {
			defer wg.Done()
			value, _ := c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
				Node: &core.Node{
					Id: "node",
				},
				TypeUrl:                typ,
				ResourceNamesSubscribe: names[typ],
			}, &stream.StreamState{IsWildcard: true, ResourceVersions: nil})

			select {
			case out := <-value:
				if !reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshotWithTtl.GetResourcesAndTtl(typ)) {
					t.Errorf("get resources %v, want %v", out.(*cache.RawDeltaResponse).Resources, snapshotWithTtl.GetResourcesAndTtl(typ))
				}
				vMap := out.GetNextVersionMap()
				vm[typ] = vMap
			case <-time.After(2 * time.Second):
				t.Errorf("failed to receive snapshot response")
			}
		})
	}
	wg.Wait()

	// Once everything is up to date, only the TTL'd resource should send out updates.
	wg = sync.WaitGroup{}
	updatesByType := map[string]int{}
	for _, typ := range testTypes {
		wg.Add(1)
		go func(typ string) {
			defer wg.Done()

			end := time.After(5 * time.Second)
			for {
				value, cancel := c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
					Node: &core.Node{
						Id: "node",
					},
					TypeUrl:                typ,
					ResourceNamesSubscribe: names[typ],
				}, &stream.StreamState{IsWildcard: true, ResourceVersions: vm[typ]})

				select {
				case out := <-value:
					if !reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshotWithTtl.GetResourcesAndTtl(typ)) {
						t.Errorf("get resources %v, want %v", out.(*cache.RawDeltaResponse).Resources, snapshotWithTtl.GetResources(typ))
					}

					if !reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshotWithTtl.GetResourcesAndTtl(typ)) {
						t.Errorf("get resources %v, want %v", out.(*cache.RawDeltaResponse).Resources, snapshotWithTtl.GetResources(typ))
					}
					vMap := out.GetNextVersionMap()
					vm[typ] = vMap

					updatesByType[typ]++
				case <-end:
					cancel()
					return
				}
			}
		}(typ)
	}

	wg.Wait()

	if len(updatesByType) != 1 {
		t.Errorf("expected to only receive updates for TTL'd type, got %v", updatesByType)
	}
	// Avoid an exact match on number of triggers to avoid this being flaky.
	if updatesByType[rsrc.EndpointType] < 2 {
		t.Errorf("expected at least two TTL updates for endpoints, got %d", updatesByType[rsrc.EndpointType])
	}
}

func TestSnapshotCacheDeltaWatch(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, logger{t: t})
	watches := make(map[string]chan cache.DeltaResponse)

	for _, typ := range testTypes {
		watches[typ], _ = c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: "node",
			},
			TypeUrl:                typ,
			ResourceNamesSubscribe: names[typ],
		}, &stream.StreamState{IsWildcard: true, ResourceVersions: nil})
	}

	if err := c.SetSnapshot(key, snapshot); err != nil {
		t.Fatal(err)
	}

	vm := make(map[string]map[string]string)
	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			select {
			case out := <-watches[typ]:
				if !reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshot.GetResourcesAndTtl(typ)) {
					t.Errorf("got resources %v, want %v", out.(*cache.RawDeltaResponse).Resources, snapshot.GetResourcesAndTtl(typ))
				}
				vMap := out.GetNextVersionMap()
				vm[typ] = vMap
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}

	for _, typ := range testTypes {
		watches[typ], _ = c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: "node",
			},
			TypeUrl:                typ,
			ResourceNamesSubscribe: names[typ],
		}, &stream.StreamState{IsWildcard: false, ResourceVersions: vm[typ]})
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
		if !reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshot2.GetResourcesAndTtl(rsrc.EndpointType)) {
			t.Fatalf("got resources %v, want %v", out.(*cache.RawDeltaResponse).Resources, snapshot2.GetResourcesAndTtl(rsrc.EndpointType))
		}
		vMap := out.GetNextVersionMap()
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
				if i < 25 {
					snap := cache.Snapshot{}
					snap.Resources[types.Endpoint] = cache.NewResources(version, []types.Resource{resource.MakeEndpoint(clusterName, uint32(i))})
					c.SetSnapshot(id, snap)
				} else {
					if cancel != nil {
						cancel()
					}

					_, cancel = c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
						Node: &core.Node{
							Id: id,
						},
						TypeUrl:                rsrc.EndpointType,
						ResourceNamesSubscribe: []string{clusterName},
					}, &stream.StreamState{IsWildcard: true, ResourceVersions: nil})
				}
			})
		}(i)
	}
}

func TestSnapshotCacheDeltaWatchCancel(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	for _, typ := range testTypes {
		_, cancel := c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: key,
			},
			TypeUrl:                typ,
			ResourceNamesSubscribe: names[typ],
		}, &stream.StreamState{IsWildcard: true, ResourceVersions: nil})

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
