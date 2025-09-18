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
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource/v3"
)

type group struct{}

const (
	key = "node"
)

func (group) ID(node *core.Node) string {
	if node != nil {
		return node.GetId()
	}
	return key
}

var (
	ttl                = 2 * time.Second
	snapshotWithTTL, _ = cache.NewSnapshotWithTTLs(fixture.version, map[rsrc.Type][]types.ResourceWithTTL{
		rsrc.EndpointType:        {{Resource: testEndpoint, TTL: &ttl}},
		rsrc.ClusterType:         {{Resource: testCluster}},
		rsrc.RouteType:           {{Resource: testRoute}, {Resource: testEmbeddedRoute}},
		rsrc.ScopedRouteType:     {{Resource: testScopedRoute}},
		rsrc.VirtualHostType:     {{Resource: testVirtualHost}},
		rsrc.ListenerType:        {{Resource: testScopedListener}, {Resource: testListener}},
		rsrc.RuntimeType:         {{Resource: testRuntime}},
		rsrc.SecretType:          {{Resource: testSecret[0]}},
		rsrc.ExtensionConfigType: {{Resource: testExtensionConfig}},
	})

	names = map[string][]string{
		rsrc.EndpointType:    {clusterName},
		rsrc.ClusterType:     nil,
		rsrc.RouteType:       {routeName, embeddedRouteName},
		rsrc.ScopedRouteType: nil,
		rsrc.VirtualHostType: nil,
		rsrc.ListenerType:    nil,
		rsrc.RuntimeType:     nil,
	}

	testTypes = []string{
		rsrc.EndpointType,
		rsrc.ClusterType,
		rsrc.RouteType,
		rsrc.ScopedRouteType,
		rsrc.VirtualHostType,
		rsrc.ListenerType,
		rsrc.RuntimeType,
	}
)

type logger struct {
	t *testing.T
}

func (log logger) Debugf(format string, args ...interface{}) {
	log.t.Helper()
	log.t.Logf(format, args...)
}

func (log logger) Infof(format string, args ...interface{}) {
	log.t.Helper()
	log.t.Logf(format, args...)
}

func (log logger) Warnf(format string, args ...interface{}) {
	log.t.Helper()
	log.t.Logf(format, args...)
}

func (log logger) Errorf(format string, args ...interface{}) {
	log.t.Helper()
	log.t.Logf(format, args...)
}

func TestSnapshotCacheWithTTL(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := cache.NewSnapshotCacheWithHeartbeating(ctx, true, group{}, logger{t: t}, time.Second)

	_, err := c.GetSnapshot(key)
	require.Errorf(t, err, "unexpected snapshot found for key %q", key)

	require.NoError(t, c.SetSnapshot(context.Background(), key, snapshotWithTTL))

	snap, err := c.GetSnapshot(key)
	require.NoError(t, err)
	assert.Truef(t, reflect.DeepEqual(snap, snapshotWithTTL), "expect snapshot: %v, got: %v", snapshotWithTTL, snap)

	wg := sync.WaitGroup{}
	// All the resources should respond immediately when version is not up to date.
	streamState := stream.NewStreamState(false, map[string]string{})
	for _, typ := range testTypes {
		wg.Add(1)
		t.Run(typ, func(t *testing.T) {
			defer wg.Done()
			value := make(chan cache.Response, 1)
			c.CreateWatch(&discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]}, streamState, value)
			select {
			case out := <-value:
				gotVersion, _ := out.GetVersion()
				assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
				assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).Resources), snapshotWithTTL.GetResourcesAndTTL(typ)), "get resources %v, want %v", out.(*cache.RawResponse).Resources, snapshotWithTTL.GetResourcesAndTTL(typ))
				// Update streamState
				streamState.SetKnownResourceNamesAsList(typ, out.GetRequest().GetResourceNames())
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
				value := make(chan cache.Response, 1)
				cancel := c.CreateWatch(&discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ], VersionInfo: fixture.version},
					streamState, value)

				select {
				case out := <-value:
					gotVersion, _ := out.GetVersion()
					assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
					assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).Resources), snapshotWithTTL.GetResourcesAndTTL(typ)), "get resources %v, want %v", out.(*cache.RawResponse).Resources, snapshotWithTTL.GetResources(typ))

					assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).Resources), snapshotWithTTL.GetResourcesAndTTL(typ)), "get resources %v, want %v", out.(*cache.RawResponse).Resources, snapshotWithTTL.GetResources(typ))

					updatesByType[typ]++

					streamState.SetKnownResourceNamesAsList(typ, out.GetRequest().GetResourceNames())
				case <-end:
					cancel()
					return
				}
			}
		}(typ)
	}

	wg.Wait()

	assert.Lenf(t, updatesByType, 1, "expected to only receive updates for TTL'd type, got %v", updatesByType)
	// Avoid an exact match on number of triggers to avoid this being flaky.
	assert.GreaterOrEqualf(t, updatesByType[rsrc.EndpointType], 2, "expected at least two TTL updates for endpoints, got %d", updatesByType[rsrc.EndpointType])
}

func TestSnapshotCache(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})

	_, err := c.GetSnapshot(key)
	require.Errorf(t, err, "unexpected snapshot found for key %q", key)

	require.NoError(t, c.SetSnapshot(context.Background(), key, fixture.snapshot()))

	snap, err := c.GetSnapshot(key)
	require.NoError(t, err)
	assert.Truef(t, reflect.DeepEqual(snap, fixture.snapshot()), "expect snapshot: %v, got: %v", fixture.snapshot(), snap)

	// try to get endpoints with incorrect list of names
	// should not receive response
	value := make(chan cache.Response, 1)
	streamState := stream.NewStreamState(false, map[string]string{})
	c.CreateWatch(&discovery.DiscoveryRequest{TypeUrl: rsrc.EndpointType, ResourceNames: []string{"none"}},
		streamState, value)
	select {
	case out := <-value:
		t.Errorf("watch for endpoints and mismatched names => got %v, want none", out)
	case <-time.After(time.Second / 4):
	}

	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			value := make(chan cache.Response, 1)
			streamState := stream.NewStreamState(false, map[string]string{})
			c.CreateWatch(&discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]},
				streamState, value)
			select {
			case out := <-value:
				snapshot := fixture.snapshot()
				gotVersion, _ := out.GetVersion()
				assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
				assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).Resources), snapshot.GetResourcesAndTTL(typ)), "get resources %v, want %v", out.(*cache.RawResponse).Resources, snapshot.GetResourcesAndTTL(typ))
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}
}

func TestSnapshotCacheFetch(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	require.NoError(t, c.SetSnapshot(context.Background(), key, fixture.snapshot()))

	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			resp, err := c.Fetch(context.Background(), &discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]})
			require.NoErrorf(t, err, "unexpected error")
			require.NotNilf(t, resp, "null response")
			gotVersion, _ := resp.GetVersion()
			assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
		})
	}

	// no response for missing snapshot
	resp, err := c.Fetch(context.Background(),
		&discovery.DiscoveryRequest{TypeUrl: rsrc.ClusterType, Node: &core.Node{Id: "oof"}})
	require.Errorf(t, err, "missing snapshot: response is not nil %v", resp)
	assert.Nilf(t, resp, "missing snapshot: response is not nil %v", resp)

	// no response for latest version
	resp, err = c.Fetch(context.Background(),
		&discovery.DiscoveryRequest{TypeUrl: rsrc.ClusterType, VersionInfo: fixture.version})
	require.Errorf(t, err, "latest version: response is not nil %v", resp)
	assert.Nilf(t, resp, "latest version: response is not nil %v", resp)
}

func TestSnapshotCacheWatch(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	watches := make(map[string]chan cache.Response)
	streamState := stream.NewStreamState(false, map[string]string{})
	for _, typ := range testTypes {
		watches[typ] = make(chan cache.Response, 1)
		c.CreateWatch(&discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]}, streamState, watches[typ])
	}

	require.NoError(t, c.SetSnapshot(context.Background(), key, fixture.snapshot()))
	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			select {
			case out := <-watches[typ]:
				gotVersion, _ := out.GetVersion()
				assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
				snapshot := fixture.snapshot()
				assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).Resources), snapshot.GetResourcesAndTTL(typ)), "get resources %v, want %v", out.(*cache.RawResponse).Resources, snapshot.GetResourcesAndTTL(typ))
				streamState.SetKnownResourceNamesAsList(typ, out.GetRequest().GetResourceNames())
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}

	// open new watches with the latest version
	for _, typ := range testTypes {
		watches[typ] = make(chan cache.Response, 1)
		c.CreateWatch(&discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ], VersionInfo: fixture.version},
			streamState, watches[typ])
	}
	count := c.GetStatusInfo(key).GetNumWatches()
	assert.Lenf(t, testTypes, count, "watches should be created for the latest version: %d", count)

	// set partially-versioned snapshot
	snapshot2 := fixture.snapshot()
	snapshot2.Resources[types.Endpoint] = cache.NewResources(fixture.version2, []types.Resource{resource.MakeEndpoint(clusterName, 9090)})
	require.NoError(t, c.SetSnapshot(context.Background(), key, snapshot2))
	count = c.GetStatusInfo(key).GetNumWatches()
	assert.Equalf(t, count, len(testTypes)-1, "watches should be preserved for all but one: %d", count)

	// validate response for endpoints
	select {
	case out := <-watches[rsrc.EndpointType]:
		gotVersion, _ := out.GetVersion()
		assert.Equalf(t, gotVersion, fixture.version2, "got version %q, want %q", gotVersion, fixture.version2)
		assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).Resources), snapshot2.Resources[types.Endpoint].Items), "got resources %v, want %v", out.(*cache.RawResponse).Resources, snapshot2.Resources[types.Endpoint].Items)
	case <-time.After(time.Second):
		t.Fatal("failed to receive snapshot response")
	}
}

func TestConcurrentSetWatch(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, logger{t: t})
	for i := 0; i < 50; i++ {
		i := i
		t.Run(fmt.Sprintf("worker%d", i), func(t *testing.T) {
			t.Parallel()
			id := t.Name()
			value := make(chan cache.Response, 1)
			if i < 25 {
				snap := cache.Snapshot{}
				snap.Resources[types.Endpoint] = cache.NewResources(fmt.Sprintf("v%d", i), []types.Resource{resource.MakeEndpoint(clusterName, uint32(i))})
				err := c.SetSnapshot(context.Background(), id, &snap)
				require.NoErrorf(t, err, "failed to set snapshot %q", id)
			} else {
				streamState := stream.NewStreamState(false, map[string]string{})
				cancel := c.CreateWatch(&discovery.DiscoveryRequest{
					Node:    &core.Node{Id: id},
					TypeUrl: rsrc.EndpointType,
				}, streamState, value)
				defer cancel()
			}
		})
	}
}

func TestSnapshotCacheWatchCancel(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	streamState := stream.NewStreamState(false, map[string]string{})
	for _, typ := range testTypes {
		value := make(chan cache.Response, 1)
		cancel := c.CreateWatch(&discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]}, streamState, value)
		cancel()
	}
	// should be status info for the node
	assert.NotEmptyf(t, c.GetStatusKeys(), "got 0, want status info for the node")

	for _, typ := range testTypes {
		count := c.GetStatusInfo(key).GetNumWatches()
		assert.LessOrEqualf(t, count, 0, "watches should be released for %s", typ)
	}

	empty := c.GetStatusInfo("missing")
	assert.Nilf(t, empty, "should not return a status for unknown key: got %#v", empty)
}

func TestSnapshotCacheWatchTimeout(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})

	// Create a non-buffered channel that will block sends.
	watchCh := make(chan cache.Response)
	streamState := stream.NewStreamState(false, map[string]string{})
	c.CreateWatch(&discovery.DiscoveryRequest{TypeUrl: rsrc.EndpointType, ResourceNames: names[rsrc.EndpointType]},
		streamState, watchCh)

	// The first time we set the snapshot without consuming from the blocking channel, so this should time out.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	err := c.SetSnapshot(ctx, key, fixture.snapshot())
	require.EqualError(t, err, context.Canceled.Error())

	// Now reset the snapshot with a consuming channel. This verifies that if setting the snapshot fails,
	// we can retry by setting the same snapshot. In other words, we keep the watch open even if we failed
	// to respond to it within the deadline.
	watchTriggeredCh := make(chan cache.Response)
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

func TestSnapshotCreateWatchWithResourcePreviouslyNotRequested(t *testing.T) {
	clusterName2 := "clusterName2"
	routeName2 := "routeName2"
	listenerName2 := "listenerName2"
	c := cache.NewSnapshotCache(false, group{}, logger{t: t})

	snapshot2, _ := cache.NewSnapshot(fixture.version, map[rsrc.Type][]types.Resource{
		rsrc.EndpointType:        {testEndpoint, resource.MakeEndpoint(clusterName2, 8080)},
		rsrc.ClusterType:         {testCluster, resource.MakeCluster(resource.Ads, clusterName2)},
		rsrc.RouteType:           {testRoute, resource.MakeRouteConfig(routeName2, clusterName2)},
		rsrc.ListenerType:        {testScopedListener, resource.MakeRouteHTTPListener(resource.Ads, listenerName2, 80, routeName2)},
		rsrc.RuntimeType:         {},
		rsrc.SecretType:          {},
		rsrc.ExtensionConfigType: {},
	})
	require.NoError(t, c.SetSnapshot(context.Background(), key, snapshot2))
	watch := make(chan cache.Response)

	// Request resource with name=ClusterName
	go func() {
		c.CreateWatch(&discovery.DiscoveryRequest{TypeUrl: rsrc.EndpointType, ResourceNames: []string{clusterName}},
			stream.NewStreamState(false, map[string]string{}), watch)
	}()

	select {
	case out := <-watch:
		gotVersion, _ := out.GetVersion()
		assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
		want := map[string]types.ResourceWithTTL{clusterName: snapshot2.Resources[types.Endpoint].Items[clusterName]}
		assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).Resources), want), "got resources %v, want %v", out.(*cache.RawResponse).Resources, want)
	case <-time.After(time.Second):
		t.Fatal("failed to receive snapshot response")
	}

	// Request additional resource with name=clusterName2 for same version
	go func() {
		state := stream.NewStreamState(false, map[string]string{})
		state.SetKnownResourceNames(rsrc.EndpointType, map[string]struct{}{clusterName: {}})
		c.CreateWatch(&discovery.DiscoveryRequest{
			TypeUrl: rsrc.EndpointType, VersionInfo: fixture.version,
			ResourceNames: []string{clusterName, clusterName2},
		}, state, watch)
	}()

	select {
	case out := <-watch:
		gotVersion, _ := out.GetVersion()
		assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
		assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).Resources), snapshot2.Resources[types.Endpoint].Items), "got resources %v, want %v", out.(*cache.RawResponse).Resources, snapshot2.Resources[types.Endpoint].Items)
	case <-time.After(time.Second):
		t.Fatal("failed to receive snapshot response")
	}

	// Repeat request for with same version and make sure a watch is created
	state := stream.NewStreamState(false, map[string]string{})
	state.SetKnownResourceNames(rsrc.EndpointType, map[string]struct{}{clusterName: {}, clusterName2: {}})
	if cancel := c.CreateWatch(&discovery.DiscoveryRequest{
		TypeUrl: rsrc.EndpointType, VersionInfo: fixture.version,
		ResourceNames: []string{clusterName, clusterName2},
	}, state, watch); cancel == nil {
		t.Fatal("Should create a watch")
	} else {
		cancel()
	}
}

// Tests the scenario where an Envoy client unsubscribes from a resource by
// omitting it from the ResourceNames in a DiscoveryRequest, then later
// resubscribes to it by including it again in ResourceNames.
func TestSnapshotCreateWatch_UnsubscribeFollowedByResubscribe(t *testing.T) {
	const (
		listenerName1 = "listenerName1"
		listenerName2 = "listenerName2"
		routeName1    = "routeName1"
		routeName2    = "routeName2"
		version       = "1"
	)
	node := &core.Node{Id: t.Name()}

	// Create a snapshot with two listener resources, each with its own route
	// configuration.
	snapshot, err := cache.NewSnapshot(version, map[rsrc.Type][]types.Resource{
		rsrc.ListenerType: {
			resource.MakeRouteHTTPListener(resource.Ads, listenerName1, 80, routeName1),
			resource.MakeRouteHTTPListener(resource.Ads, listenerName2, 80, routeName2),
		},
	})
	if err != nil {
		t.Fatalf("Failed to create snapshot: %v", err)
	}

	// Create a cache and set the snapshot.
	c := cache.NewSnapshotCache(false, cache.IDHash{}, logger{t: t})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := c.SetSnapshot(ctx, node.GetId(), snapshot); err != nil {
		t.Fatalf("Failed to set snapshot: %v", err)
	}

	// Create a watchResponseCh for both listeners.
	watchResponseCh := make(chan cache.Response, 1)
	streamState := stream.NewStreamState(false, map[string]string{})
	req := &discovery.DiscoveryRequest{
		TypeUrl:       rsrc.ListenerType,
		ResourceNames: []string{listenerName1, listenerName2},
		Node:          node,
	}
	c.CreateWatch(req, streamState, watchResponseCh)

	// Wait for the initial response with both listeners.
	var gotWatchResponse cache.Response
	var gotDiscoveryResponse *discovery.DiscoveryResponse
	select {
	case gotWatchResponse = <-watchResponseCh:
		var err error
		gotDiscoveryResponse, err = gotWatchResponse.GetDiscoveryResponse()
		if err != nil {
			t.Fatalf("Failed to get discovery response: %v", err)
		}
	case <-ctx.Done():
		t.Fatal("Test timed out when waiting for a response from the watch")
	}

	// Validate the initial response.
	gotVersion := gotDiscoveryResponse.GetVersionInfo()
	if gotVersion != version {
		t.Fatalf("Got version %q, want %q", gotVersion, version)
	}
	wantResources := map[string]types.ResourceWithTTL{
		listenerName1: snapshot.Resources[types.Listener].Items[listenerName1],
		listenerName2: snapshot.Resources[types.Listener].Items[listenerName2],
	}
	gotResources := cache.IndexResourcesByName(gotWatchResponse.(*cache.RawResponse).Resources)
	if diff := cmp.Diff(wantResources, gotResources, protocmp.Transform()); diff != "" {
		t.Fatalf("Response resources (-want +got):\n%s", diff)
	}

	// Simulate an ACK from Envoy by creating a watch for the same set of
	// resources, but with the updated version and nonce from the initial
	// response. The set of known resource names in the stream state is actually
	// updated by the server implementation when it sees the ACK, before it
	// recreates the watch.
	streamState.SetKnownResourceNamesAsList(rsrc.ListenerType, []string{listenerName1, listenerName2})
	req = &discovery.DiscoveryRequest{
		TypeUrl:       rsrc.ListenerType,
		ResourceNames: []string{listenerName1, listenerName2},
		VersionInfo:   gotVersion,
		ResponseNonce: gotDiscoveryResponse.GetNonce(),
		Node:          node,
	}
	c.CreateWatch(req, streamState, watchResponseCh)

	// Wait for the ACK to be processed. Rely on the stream state being updated
	// with the known resource names.
	wantKnownResourceNames := map[string]struct{}{
		listenerName1: {},
		listenerName2: {},
	}
	var gotKnownResourceNames map[string]struct{}
	for ; ctx.Err() == nil; <-time.After(10 * time.Millisecond) {
		gotKnownResourceNames = streamState.GetKnownResourceNames(rsrc.ListenerType)
		if diff := cmp.Diff(wantKnownResourceNames, gotKnownResourceNames); diff == "" {
			break
		}
	}
	if ctx.Err() != nil {
		diff := cmp.Diff(wantKnownResourceNames, gotKnownResourceNames)
		t.Fatalf("Test timed out when waiting for the ACK to be processed\nKnown resource names (-want +got):\n%s", diff)
	}

	// Now simulate unsubscribing from listenerName2 by creating a watch for
	// just listenerName1, with the current version and nonce.
	req = &discovery.DiscoveryRequest{
		TypeUrl:       rsrc.ListenerType,
		ResourceNames: []string{listenerName1},
		VersionInfo:   gotVersion,
		ResponseNonce: gotDiscoveryResponse.GetNonce(),
		Node:          node,
	}
	c.CreateWatch(req, streamState, watchResponseCh)

	// Ensure there is no response from the cache after a short wait, since
	// listenerName1 is already known at the current version.
	sCtx, sCancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer sCancel()
	select {
	case <-watchResponseCh:
		t.Fatalf("Received unexpected response when unsubscribing from a resource")
	case <-sCtx.Done():
	}

	// Wait for the stream state to be updated with just listenerName1 as the
	// known resource name, indicating that the unsubscribe has been processed.
	wantKnownResourceNames = map[string]struct{}{listenerName1: {}}
	for ; ctx.Err() == nil; <-time.After(10 * time.Millisecond) {
		gotKnownResourceNames = streamState.GetKnownResourceNames(rsrc.ListenerType)
		if diff := cmp.Diff(wantKnownResourceNames, gotKnownResourceNames); diff == "" {
			break
		}
	}
	if ctx.Err() != nil {
		diff := cmp.Diff(wantKnownResourceNames, gotKnownResourceNames)
		t.Fatalf("Test timed out when waiting for the unsubscribe to be processed\nKnown resource names (-want +got):\n%s", diff)
	}

	// Now simulate resubscribing to listenerName2 by creating a watch for
	// both listenerName1 and listenerName2, with the current version and nonce.
	req = &discovery.DiscoveryRequest{
		TypeUrl:       rsrc.ListenerType,
		ResourceNames: []string{listenerName1, listenerName2},
		VersionInfo:   gotVersion,
		ResponseNonce: gotDiscoveryResponse.GetNonce(),
		Node:          node,
	}
	c.CreateWatch(req, streamState, watchResponseCh)

	// Wait for the response with both listeners.
	select {
	case gotWatchResponse = <-watchResponseCh:
	case <-ctx.Done():
		t.Fatal("Test timed out when waiting for a response from the watch")
	}

	// Validate the initial response.
	if gotVersion = gotDiscoveryResponse.GetVersionInfo(); gotVersion != version {
		t.Fatalf("Got version %q, want %q", gotVersion, version)
	}
	wantResources = map[string]types.ResourceWithTTL{
		listenerName1: snapshot.Resources[types.Listener].Items[listenerName1],
		listenerName2: snapshot.Resources[types.Listener].Items[listenerName2],
	}
	gotResources = cache.IndexResourcesByName(gotWatchResponse.(*cache.RawResponse).Resources)
	if diff := cmp.Diff(wantResources, gotResources, protocmp.Transform()); diff != "" {
		t.Fatalf("Response resources (-want +got):\n%s", diff)
	}
}

func TestSnapshotClear(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, logger{t: t})

	require.NoError(t, c.SetSnapshot(context.Background(), key, fixture.snapshot()))
	c.ClearSnapshot(key)
	assert.Nilf(t, c.GetStatusInfo(key), "cache should be cleared")
	assert.Emptyf(t, c.GetStatusKeys(), "keys should be empty")
}

type singleResourceSnapshot struct {
	version  string
	typeurl  string
	name     string
	resource types.Resource
}

func (s *singleResourceSnapshot) GetVersion(typeURL string) string {
	if typeURL != s.typeurl {
		return ""
	}

	return s.version
}

func (s *singleResourceSnapshot) GetResourcesAndTTL(typeURL string) map[string]types.ResourceWithTTL {
	if typeURL != s.typeurl {
		return nil
	}

	ttl := time.Second
	return map[string]types.ResourceWithTTL{
		s.name: {Resource: s.resource, TTL: &ttl},
	}
}

func (s *singleResourceSnapshot) GetResources(typeURL string) map[string]types.Resource {
	if typeURL != s.typeurl {
		return nil
	}

	return map[string]types.Resource{
		s.name: s.resource,
	}
}

func (s *singleResourceSnapshot) ConstructVersionMap() error {
	return nil
}

func (s *singleResourceSnapshot) GetVersionMap(typeURL string) map[string]string {
	if typeURL != s.typeurl {
		return nil
	}
	return map[string]string{
		s.name: s.version,
	}
}

// TestSnapshotSingleResourceFetch is a basic test to verify that simple
// cache functions work with a type that is not `Snapshot`.
func TestSnapshotSingleResourceFetch(t *testing.T) {
	durationTypeURL := "type.googleapis.com/" + string(proto.MessageName(&durationpb.Duration{}))

	anyDuration := func(d time.Duration) *anypb.Any {
		bytes, err := cache.MarshalResource(durationpb.New(d))
		require.NoError(t, err)
		return &anypb.Any{
			TypeUrl: durationTypeURL,
			Value:   bytes,
		}
	}

	unwrapResource := func(src *anypb.Any) *discovery.Resource {
		dst := &discovery.Resource{}
		require.NoError(t, anypb.UnmarshalTo(src, dst, proto.UnmarshalOptions{}))
		return dst
	}

	c := cache.NewSnapshotCache(true, group{}, logger{t: t})
	require.NoError(t, c.SetSnapshot(context.Background(), key, &singleResourceSnapshot{
		version:  "version-one",
		typeurl:  durationTypeURL,
		name:     "one-second",
		resource: durationpb.New(time.Second),
	}))

	resp, err := c.Fetch(context.Background(), &discovery.DiscoveryRequest{
		TypeUrl:       durationTypeURL,
		ResourceNames: []string{"one-second"},
	})
	require.NoError(t, err)

	vers, err := resp.GetVersion()
	require.NoError(t, err)
	assert.Equal(t, "version-one", vers)

	discoveryResponse, err := resp.GetDiscoveryResponse()
	require.NoError(t, err)
	assert.Equal(t, durationTypeURL, discoveryResponse.GetTypeUrl())
	require.Len(t, discoveryResponse.GetResources(), 1)
	assert.Empty(t, cmp.Diff(
		unwrapResource(discoveryResponse.GetResources()[0]).GetResource(),
		anyDuration(time.Second),
		protocmp.Transform()),
	)
}

func TestAvertPanicForWatchOnNonExistentSnapshot(t *testing.T) {
	ctx := context.Background()
	c := cache.NewSnapshotCacheWithHeartbeating(ctx, false, cache.IDHash{}, nil, time.Millisecond)

	// Create watch.
	req := &cache.Request{
		Node:          &core.Node{Id: "test"},
		ResourceNames: []string{"rtds"},
		TypeUrl:       rsrc.RuntimeType,
	}
	ss := stream.NewStreamState(false, map[string]string{"cluster": "abcdef"})
	responder := make(chan cache.Response)
	c.CreateWatch(req, ss, responder)

	go func() {
		// Wait for at least one heartbeat to occur, then set snapshot.
		time.Sleep(time.Millisecond * 5)
		srs := &singleResourceSnapshot{
			version:  "version-one",
			typeurl:  rsrc.RuntimeType,
			name:     "one-second",
			resource: durationpb.New(time.Second),
		}
		err := c.SetSnapshot(ctx, "test", srs)
		assert.NoErrorf(t, err, "unexpected error setting snapshot %v", err)
	}()

	<-responder
}
