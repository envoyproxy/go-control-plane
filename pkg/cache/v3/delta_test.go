package cache_test

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource/v3"
)

func assertResourceMapEqual(t *testing.T, want, got map[string]types.Resource) {
	t.Helper()

	if !cmp.Equal(want, got, protocmp.Transform()) {
		t.Errorf("got resources %v, want %v", got, want)
	}
}

func TestSnapshotCacheDeltaWatch(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, logger{t: t})
	watches := make(map[string]chan cache.DeltaResponse)

	// Make our initial request as a wildcard to get all resources and make sure the wildcard requesting works as intended
	for _, typ := range testTypes {
		watches[typ] = make(chan cache.DeltaResponse, 1)
		c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: "node",
			},
			TypeUrl:                typ,
			ResourceNamesSubscribe: names[typ],
		}, stream.NewStreamState(true, nil), watches[typ])
	}

	require.NoError(t, c.SetSnapshot(context.Background(), key, fixture.snapshot()))

	versionMap := make(map[string]map[string]string)
	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			select {
			case out := <-watches[typ]:
				snapshot := fixture.snapshot()
				assertResourceMapEqual(t, cache.IndexRawResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshot.GetResources(typ))
				vMap := out.GetNextVersionMap()
				versionMap[typ] = vMap
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}

	// On re-request we want to use non-wildcard so we can verify the logic path of not requesting
	// all resources as well as individual resource removals
	for _, typ := range testTypes {
		watches[typ] = make(chan cache.DeltaResponse, 1)
		state := stream.NewStreamState(false, versionMap[typ])
		for resource := range versionMap[typ] {
			state.GetSubscribedResourceNames()[resource] = struct{}{}
		}
		c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: "node",
			},
			TypeUrl:                typ,
			ResourceNamesSubscribe: names[typ],
		}, state, watches[typ])
	}

	count := c.GetStatusInfo(key).GetNumDeltaWatches()
	assert.Lenf(t, testTypes, count, "watches should be created for the latest version, saw %d watches expected %d", count, len(testTypes))

	// set partially-versioned snapshot
	snapshot2 := fixture.snapshot()
	snapshot2.Resources[types.Endpoint] = cache.NewResources(fixture.version2, []types.Resource{resource.MakeEndpoint(clusterName, 9090)})
	require.NoError(t, c.SetSnapshot(context.Background(), key, snapshot2))
	count = c.GetStatusInfo(key).GetNumDeltaWatches()
	assert.Equalf(t, count, len(testTypes)-1, "watches should be preserved for all but one, got: %d open watches instead of the expected %d open watches", count, len(testTypes)-1)

	// validate response for endpoints
	select {
	case out := <-watches[testTypes[0]]:
		snapshot2 := fixture.snapshot()
		snapshot2.Resources[types.Endpoint] = cache.NewResources(fixture.version2, []types.Resource{resource.MakeEndpoint(clusterName, 9090)})
		assertResourceMapEqual(t, cache.IndexRawResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshot2.GetResources(rsrc.EndpointType))
		vMap := out.GetNextVersionMap()
		versionMap[testTypes[0]] = vMap
	case <-time.After(time.Second):
		t.Fatal("failed to receive snapshot response")
	}
}

func TestDeltaRemoveResources(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, logger{t: t})
	watches := make(map[string]chan cache.DeltaResponse)
	streams := make(map[string]*stream.StreamState)

	for _, typ := range testTypes {
		watches[typ] = make(chan cache.DeltaResponse, 1)
		state := stream.NewStreamState(true, make(map[string]string))
		streams[typ] = &state
		// We don't specify any resource name subscriptions here because we want to make sure we test wildcard
		// functionality. This means we should receive all resources back without requesting a subscription by name.
		c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: "node",
			},
			TypeUrl: typ,
		}, *streams[typ], watches[typ])
	}

	require.NoError(t, c.SetSnapshot(context.Background(), key, fixture.snapshot()))

	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			select {
			case out := <-watches[typ]:
				snapshot := fixture.snapshot()
				assertResourceMapEqual(t, cache.IndexRawResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshot.GetResources(typ))
				nextVersionMap := out.GetNextVersionMap()
				streams[typ].SetResourceVersions(nextVersionMap)
			case <-time.After(time.Second):
				t.Fatal("failed to receive a snapshot response")
			}
		})
	}

	// We want to continue to do wildcard requests here so we can later
	// test the removal of certain resources from a partial snapshot
	for _, typ := range testTypes {
		watches[typ] = make(chan cache.DeltaResponse, 1)
		c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: "node",
			},
			TypeUrl: typ,
		}, *streams[typ], watches[typ])
	}

	count := c.GetStatusInfo(key).GetNumDeltaWatches()
	assert.Lenf(t, testTypes, count, "watches should be created for the latest version, saw %d watches expected %d", count, len(testTypes))

	// set a partially versioned snapshot with no endpoints
	snapshot2 := fixture.snapshot()
	snapshot2.Resources[types.Endpoint] = cache.NewResources(fixture.version2, []types.Resource{})
	require.NoError(t, c.SetSnapshot(context.Background(), key, snapshot2))

	// validate response for endpoints
	select {
	case out := <-watches[testTypes[0]]:
		snapshot2 := fixture.snapshot()
		snapshot2.Resources[types.Endpoint] = cache.NewResources(fixture.version2, []types.Resource{})
		assertResourceMapEqual(t, cache.IndexRawResourcesByName(out.(*cache.RawDeltaResponse).Resources), snapshot2.GetResources(rsrc.EndpointType))
		nextVersionMap := out.GetNextVersionMap()

		// make sure the version maps are different since we no longer are tracking any endpoint resources
		if reflect.DeepEqual(streams[testTypes[0]].GetResourceVersions(), nextVersionMap) {
			t.Fatalf("versionMap for the endpoint resource type did not change, received: %v, instead of an empty map", nextVersionMap)
		}
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
				id := strconv.Itoa(i % 2)
				responses := make(chan cache.DeltaResponse, 1)
				if i < 25 {
					snap, err := cache.NewSnapshot("", map[rsrc.Type][]types.Resource{})
					require.NoError(t, err)
					snap.Resources[types.Endpoint] = cache.NewResources(version, []types.Resource{resource.MakeEndpoint(clusterName, uint32(i))})
					err = c.SetSnapshot(context.Background(), key, snap)
					require.NoErrorf(t, err, "snapshot failed")
				} else {
					cancel := c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
						Node: &core.Node{
							Id: id,
						},
						TypeUrl:                rsrc.EndpointType,
						ResourceNamesSubscribe: []string{clusterName},
					}, stream.NewStreamState(false, make(map[string]string)), responses)

					defer cancel()
				}
			})
		}(i)
	}
}

type testKey struct{}

func TestSnapshotDeltaCacheWatchTimeout(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})

	// Create a non-buffered channel that will block sends.
	watchCh := make(chan cache.DeltaResponse)
	state := stream.NewStreamState(false, nil)
	state.SetSubscribedResourceNames(map[string]struct{}{names[rsrc.EndpointType][0]: {}})
	c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
		Node: &core.Node{
			Id: key,
		},
		TypeUrl:                rsrc.EndpointType,
		ResourceNamesSubscribe: names[rsrc.EndpointType],
	}, state, watchCh)

	// The first time we set the snapshot without consuming from the blocking channel, so this should time out.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	err := c.SetSnapshot(ctx, key, fixture.snapshot())
	require.EqualError(t, err, context.Canceled.Error())

	// Now reset the snapshot with a consuming channel. This verifies that if setting the snapshot fails,
	// we can retry by setting the same snapshot. In other words, we keep the watch open even if we failed
	// to respond to it within the deadline.
	watchTriggeredCh := make(chan cache.DeltaResponse)
	go func() {
		response := <-watchCh
		watchTriggeredCh <- response
		close(watchTriggeredCh)
	}()

	err = c.SetSnapshot(context.WithValue(context.Background(), testKey{}, "bar"), key, fixture.snapshot())
	require.NoError(t, err)

	// The channel should get closed due to the watch trigger.
	select {
	case response := <-watchTriggeredCh:
		// Verify that we pass the context through.
		assert.Equal(t, "bar", response.GetContext().Value(testKey{}))
	case <-time.After(time.Second):
		t.Fatalf("timed out")
	}
}

func TestSnapshotCacheDeltaWatchCancel(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	for _, typ := range testTypes {
		responses := make(chan cache.DeltaResponse, 1)
		cancel := c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
			Node: &core.Node{
				Id: key,
			},
			TypeUrl:                typ,
			ResourceNamesSubscribe: names[typ],
		}, stream.NewStreamState(false, make(map[string]string)), responses)

		// Cancel the watch
		cancel()
	}
	// c.GetStatusKeys() should return at least 1 because we register a node ID with the above watch creations
	keys := c.GetStatusKeys()
	assert.NotEmptyf(t, keys, "expected to see a status info registered for watch, saw %d entries", len(keys))

	for _, typ := range testTypes {
		count := c.GetStatusInfo(key).GetNumDeltaWatches()
		assert.LessOrEqualf(t, count, 0, "watches should be released for %s", typ)
	}

	s := c.GetStatusInfo("missing")
	assert.Nilf(t, s, "should not return a status for unknown key: got %#v", s)
}
