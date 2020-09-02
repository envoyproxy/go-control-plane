package cache_test

import (
	"fmt"
	"testing"
	"time"

	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v2"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource/v2"
)

func TestSnapshotCacheDelta(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, logger{t: t})
	watches := make(map[string]chan cache.DeltaResponse)

	for _, typ := range testTypes {
		watches[typ], _ = c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: "node",
			},
			TypeUrl:                typ,
			ResourceNamesSubscribe: names[typ],
		})
	}

	if err := c.SetSnapshot(key, snapshot); err != nil {
		t.Fatal(err)
	}

	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			select {
			case out := <-watches[typ]:
				// if !reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshot.GetResources(typ)) {
				// 	t.Errorf("get resources %v, want %v", out.(*cache.RawDeltaResponse).Resources, snapshot.GetResources(typ))
				// }
				res, err := out.GetDeltaDiscoveryResponse()
				if err != nil {
					t.Fatal(err)
				}
				t.Log(res)
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
		})
	}
}

// func TestSnapshotCacheDeltaWatch(t *testing.T) {
// 	c := cache.NewSnapshotCache(false, group{}, logger{t: t})
// 	watches := make(map[string]chan cache.DeltaResponse)

// 	for _, typ := range testTypes {
// 		watches[typ], _ = c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
// 			Node: &core.Node{
// 				Id: "node",
// 			},
// 			TypeUrl:                typ,
// 			ResourceNamesSubscribe: names[typ],
// 		})
// 	}

// 	if err := c.SetSnapshot(key, snapshot); err != nil {
// 		t.Fatal(err)
// 	}

// 	for _, typ := range testTypes {
// 		t.Run(typ, func(t *testing.T) {
// 			select {
// 			case out := <-watches[typ]:
// 				// if v, _ := out.GetSystemVersion(); v != version {
// 				// 	t.Errorf("got version %q, want %q", v, version)
// 				// }
// 				// if !reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshot.GetResources(typ)) {
// 				// 	t.Errorf("get resources %v, want %v", out.(*cache.RawDeltaResponse).Resources, snapshot.GetResources(typ))
// 				// }
// 			case <-time.After(time.Second):
// 				t.Fatal("failed to receive snapshot response")
// 			}
// 		})
// 	}

// 	// open new watches with the latest version
// 	for _, typ := range testTypes {
// 		watches[typ], _ = c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
// 			Node: &core.Node{
// 				Id: "node",
// 			},
// 			TypeUrl:                typ,
// 			ResourceNamesSubscribe: names[typ],
// 		})
// 	}

// 	if count := c.GetStatusInfo(key).GetNumDeltaWatches(); count != len(testTypes) {
// 		t.Errorf("watches should be created for the latest version: %d", count)
// 	}

// 	// set partially-versioned snapshot
// 	snapshot2 := snapshot
// 	snapshot2.Resources[types.Endpoint] = cache.NewResources(version2, []types.Resource{resource.MakeEndpoint(clusterName, 9090)})
// 	if err := c.SetSnapshot(key, snapshot2); err != nil {
// 		t.Fatal(err)
// 	}

// 	if count := c.GetStatusInfo(key).GetNumDeltaWatches(); count != len(testTypes)-1 {
// 		t.Errorf("watches should be preserved for all but one: %d", count)
// 	}

// 	// validate response for endpoints
// 	select {
// 	case out := <-watches[rsrc.EndpointType]:
// 		// if v, _ := out.GetSystemVersion(); v != version2 {
// 		// 	t.Errorf("got version %q, want %q", v, version2)
// 		// }
// 		// res, err := out.GetDeltaDiscoveryResponse()
// 		// if err != nil {
// 		// 	t.Fatal("failed to retrieve DeltaDiscoveryResponse")
// 		// }

// 		// fmt.Println(res)
// 		// if !reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshot2.Resources[types.Endpoint].Items) {
// 		// 	t.Fatalf("got resources %v, want %v", out.(*cache.RawDeltaResponse).Resources, snapshot2.Resources[types.Endpoint].Items)
// 		// }
// 	case <-time.After(time.Second * 5):
// 		t.Fatal("failed to receive snapshot response")
// 	}

// 	// test an unsubscribe scenario
// 	// Assume we got a request from the grpc server to unsubscribe from a resource so we can initiate a request
// 	watches[testTypes[0]], _ = c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
// 		Node: &core.Node{
// 			Id: "node",
// 		},
// 		TypeUrl:                  testTypes[0],
// 		ResourceNamesUnsubscribe: []string{clusterName},
// 	})

// 	if count := c.GetStatusInfo(key).GetNumDeltaWatches(); count != len(testTypes) {
// 		t.Errorf("watches should be preserved for all but one %d", count)
// 	}
// }

func TestCheckState(t *testing.T) {
	deltaState := map[string][]string{
		rsrc.EndpointType: {clusterName},
		rsrc.ClusterType:  {clusterName},
		rsrc.RouteType:    {routeName},
		rsrc.ListenerType: {listenerName},
		rsrc.RuntimeType:  nil,
	}
	subscribed := map[string][]string{
		rsrc.EndpointType: {clusterName, "clusterDelta2"},
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
					})
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
		})

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
