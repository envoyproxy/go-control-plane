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
	"github.com/envoyproxy/go-control-plane/pkg/log"
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

func subFromRequest(req *cache.Request) stream.Subscription {
	return stream.NewSotwSubscription(req.GetResourceNames(), true)
}

// This method represents the expected behavior of client and servers regarding the request and the subscription.
// For edge cases it should ignore those
func updateFromSotwResponse(resp cache.Response, sub *stream.Subscription, req *cache.Request) {
	sub.SetReturnedResources(resp.GetReturnedResources())
	// Never returns an error when not using passthrough responses
	version, _ := resp.GetVersion()
	req.VersionInfo = version
}

func srfrt(res types.ResourceWithTTL) types.SnapshotResource {
	return types.SnapshotResource{
		Name:     cache.GetResourceName(res.Resource),
		Resource: res.Resource,
		TTL:      res.TTL,
	}
}

var (
	ttl           = 2 * time.Second
	testResources = map[rsrc.Type][]types.ResourceWithTTL{
		rsrc.EndpointType:        {{Resource: testEndpoint, TTL: &ttl}},
		rsrc.ClusterType:         {{Resource: testCluster}},
		rsrc.RouteType:           {{Resource: testRoute}, {Resource: testEmbeddedRoute}},
		rsrc.ScopedRouteType:     {{Resource: testScopedRoute}},
		rsrc.VirtualHostType:     {{Resource: testVirtualHost}},
		rsrc.ListenerType:        {{Resource: testScopedListener}, {Resource: testListener}},
		rsrc.RuntimeType:         {{Resource: testRuntime}},
		rsrc.SecretType:          {{Resource: testSecret[0]}},
		rsrc.ExtensionConfigType: {{Resource: testExtensionConfig}},
	}

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

func TestSnapshotCacheWithTTL(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := cache.NewSnapshotCacheWithHeartbeating(ctx, true, group{}, log.NewTestLogger(t), time.Second)

	_, err := c.GetSnapshot(key)
	require.Errorf(t, err, "unexpected snapshot found for key %q", key)

	snapshots := []types.TypeSnapshot{}
	for typeURL, resources := range testResources {
		snapRes := []types.SnapshotResource{}
		for _, res := range resources {
			snapRes = append(snapRes, srfrt(res))
		}
		snap, err := types.NewTypeSnapshot(typeURL, fixture.version, snapRes)
		require.NoError(t, err)
		snapshots = append(snapshots, snap)
	}
	snapshotWithTTL, err := types.NewSnapshotFromTypeSnapshots(fixture.version, snapshots)
	require.NoError(t, err)
	require.NoError(t, c.SetSnapshot(context.Background(), key, snapshotWithTTL))

	snap, err := c.GetSnapshot(key)
	require.NoError(t, err)
	assert.Truef(t, reflect.DeepEqual(snap, snapshotWithTTL), "expect snapshot: %v, got: %v", snapshotWithTTL, snap)

	wg := sync.WaitGroup{}
	// All the resources should respond immediately when version is not up to date.
	subs := map[string]stream.Subscription{}
	mu := sync.Mutex{}
	for _, typ := range testTypes {
		wg.Add(1)
		t.Run(typ, func(t *testing.T) {
			defer wg.Done()
			req := &discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]}
			sub := subFromRequest(req)
			value := make(chan cache.Response, 1)
			_, err = c.CreateWatch(req, sub, value)
			require.NoError(t, err)
			select {
			case out := <-value:
				gotVersion, _ := out.GetVersion()
				assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
				assert.ElementsMatch(t, out.(*cache.RawResponse).GetRawResources(), testResources[typ])

				updateFromSotwResponse(out, &sub, req)
				mu.Lock()
				subs[typ] = sub
				mu.Unlock()
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
				mu.Lock()
				sub := subs[typ]
				mu.Unlock()
				value := make(chan cache.Response, 1)
				cancel, err := c.CreateWatch(&discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ], VersionInfo: fixture.version},
					sub, value)
				require.NoError(t, err)

				select {
				case out := <-value:
					t.Logf("received reply for %s, expected %s", out.GetRequest().TypeUrl, typ)
					gotVersion, _ := out.GetVersion()
					assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
					assert.ElementsMatch(t, out.(*cache.RawResponse).GetRawResources(), testResources[typ])

					mu.Lock()
					updatesByType[typ]++
					mu.Unlock()

					returnedResources := make(map[string]string)
					// Update sub to track what was returned
					for _, resource := range out.GetRequest().GetResourceNames() {
						returnedResources[resource] = fixture.version
					}
					sub.SetReturnedResources(returnedResources)
					mu.Lock()
					subs[typ] = sub
					mu.Unlock()
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
	c := cache.NewSnapshotCache(true, group{}, log.NewTestLogger(t))

	_, err := c.GetSnapshot(key)
	require.Errorf(t, err, "unexpected snapshot found for key %q", key)

	require.NoError(t, c.SetSnapshot(context.Background(), key, fixture.snapshot()))

	snap, err := c.GetSnapshot(key)
	require.NoError(t, err)
	assert.Truef(t, reflect.DeepEqual(snap, fixture.snapshot()), "expect snapshot: %v, got: %v", fixture.snapshot(), snap)

	// try to get endpoints with incorrect list of names
	// should not receive response
	value := make(chan cache.Response, 1)
	req := &discovery.DiscoveryRequest{TypeUrl: rsrc.EndpointType, ResourceNames: []string{"none"}}
	sub := subFromRequest(req)
	_, err = c.CreateWatch(req, sub, value)
	require.NoError(t, err)
	select {
	case out := <-value:
		t.Errorf("watch for endpoints and mismatched names => got %v, want none", out)
	case <-time.After(time.Second / 4):
	}

	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			value := make(chan cache.Response, 1)
			req := &discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]}
			sub := subFromRequest(req)
			_, err = c.CreateWatch(req, sub, value)
			require.NoError(t, err)
			select {
			case out := <-value:
				snapshot := fixture.snapshot()
				gotVersion, _ := out.GetVersion()
				assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
				assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).GetRawResources()), snapshot.GetResourcesAndTTL(typ)), "get resources %v, want %v", out.(*cache.RawResponse).GetRawResources(), snapshot.GetResourcesAndTTL(typ))
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}
}

func TestSnapshotCacheFetch(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, log.NewTestLogger(t))
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
	c := cache.NewSnapshotCache(true, group{}, log.NewTestLogger(t))
	watches := make(map[string]chan cache.Response)
	subs := map[string]stream.Subscription{}
	for _, typ := range testTypes {
		req := &discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]}
		sub := subFromRequest(req)
		subs[typ] = sub
		watches[typ] = make(chan cache.Response, 1)
		_, err := c.CreateWatch(req, sub, watches[typ])
		require.NoError(t, err)
	}

	require.NoError(t, c.SetSnapshot(context.Background(), key, fixture.snapshot()))
	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			select {
			case out := <-watches[typ]:
				gotVersion, _ := out.GetVersion()
				assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
				snapshot := fixture.snapshot()
				assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).GetRawResources()), snapshot.GetResourcesAndTTL(typ)), "get resources %v, want %v", out.(*cache.RawResponse).GetRawResources(), snapshot.GetResourcesAndTTL(typ))
				returnedResources := make(map[string]string)
				// Update sub to track what was returned
				for _, resource := range out.GetRequest().GetResourceNames() {
					returnedResources[resource] = fixture.version
				}
				sub := subs[typ]
				updateFromSotwResponse(out, &sub, out.GetRequest())
				subs[typ] = sub
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}

	// open new watches with the latest version
	for _, typ := range testTypes {
		watches[typ] = make(chan cache.Response, 1)
		_, err := c.CreateWatch(&discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ], VersionInfo: fixture.version},
			subs[typ], watches[typ])
		require.NoError(t, err)
	}
	count := c.GetStatusInfo(key).GetNumWatches()
	assert.Lenf(t, testTypes, count, "watches should be created for the latest version: %d", count)

	// set partially-versioned snapshot
	snapshot2 := fixture.snapshot()
	snapshot2.Resources[types.Cluster] = cache.NewResources(fixture.version2, []types.Resource{resource.MakeCluster(resource.Ads, clusterName)})
	require.NoError(t, c.SetSnapshot(context.Background(), key, snapshot2))
	count = c.GetStatusInfo(key).GetNumWatches()
	assert.Equalf(t, count, len(testTypes)-1, "watches should be preserved for all but one: %d", count)

	// validate response for endpoints
	select {
	case out := <-watches[rsrc.ClusterType]:
		gotVersion, _ := out.GetVersion()
		assert.Equalf(t, gotVersion, fixture.version2, "got version %q, want %q", gotVersion, fixture.version2)
		assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).GetRawResources()), snapshot2.Resources[types.Cluster].Items), "got resources %v, want %v", out.(*cache.RawResponse).GetRawResources(), snapshot2.Resources[types.Endpoint].Items)
	case <-time.After(time.Second):
		t.Fatal("failed to receive snapshot response")
	}
}

func TestConcurrentSetWatch(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, log.NewTestLogger(t))
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
				req := &discovery.DiscoveryRequest{
					Node:    &core.Node{Id: id},
					TypeUrl: rsrc.EndpointType,
				}
				cancel, err := c.CreateWatch(req, subFromRequest(req), value)
				require.NoError(t, err)
				defer cancel()
			}
		})
	}
}

func TestSnapshotCacheWatchCancel(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, log.NewTestLogger(t))
	for _, typ := range testTypes {
		req := &discovery.DiscoveryRequest{TypeUrl: typ, ResourceNames: names[typ]}
		value := make(chan cache.Response, 1)
		cancel, err := c.CreateWatch(req, subFromRequest(req), value)
		require.NoError(t, err)
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
	c := cache.NewSnapshotCache(true, group{}, log.NewTestLogger(t))

	// Create a non-buffered channel that will block sends.
	watchCh := make(chan cache.Response)
	req := &discovery.DiscoveryRequest{TypeUrl: rsrc.EndpointType, ResourceNames: names[rsrc.EndpointType]}
	sub := subFromRequest(req)
	_, err := c.CreateWatch(req, sub, watchCh)
	require.NoError(t, err)

	// The first time we set the snapshot without consuming from the blocking channel, so this should time out.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	err = c.SetSnapshot(ctx, key, fixture.snapshot())
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
	c := cache.NewSnapshotCache(false, group{}, log.NewTestLogger(t))

	testResources := map[rsrc.Type][]types.ResourceWithTTL{
		rsrc.EndpointType:        {{Resource: testEndpoint}, {Resource: resource.MakeEndpoint(clusterName2, 8080)}},
		rsrc.ClusterType:         {{Resource: testCluster}, {Resource: resource.MakeCluster(resource.Ads, clusterName2)}},
		rsrc.RouteType:           {{Resource: testRoute}, {Resource: resource.MakeRouteConfig(routeName2, clusterName2)}},
		rsrc.ListenerType:        {{Resource: testScopedListener}, {Resource: resource.MakeRouteHTTPListener(resource.Ads, listenerName2, 80, routeName2)}},
		rsrc.RuntimeType:         {},
		rsrc.SecretType:          {},
		rsrc.ExtensionConfigType: {},
	}

	snapshots := []types.TypeSnapshot{}
	for typeURL, resources := range testResources {
		snapRes := []types.SnapshotResource{}
		for _, res := range resources {
			snapRes = append(snapRes, srfrt(res))
		}
		snap, err := types.NewTypeSnapshot(typeURL, fixture.version, snapRes)
		require.NoError(t, err)
		snapshots = append(snapshots, snap)
	}
	snapshot2, err := types.NewSnapshotFromTypeSnapshots(fixture.version, snapshots)
	require.NoError(t, err)
	require.NoError(t, c.SetSnapshot(context.Background(), key, snapshot2))
	watch := make(chan cache.Response)

	// Request resource with name=ClusterName
	go func() {
		req := &discovery.DiscoveryRequest{TypeUrl: rsrc.ClusterType, ResourceNames: []string{clusterName}}
		_, err := c.CreateWatch(req, subFromRequest(req), watch)
		require.NoError(t, err)
	}()

	select {
	case out := <-watch:
		gotVersion, _ := out.GetVersion()
		assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
		want := map[string]types.ResourceWithTTL{clusterName: {Resource: testCluster}}
		assert.Truef(t, reflect.DeepEqual(cache.IndexResourcesByName(out.(*cache.RawResponse).GetRawResources()), want), "got resources %v, want %v", out.(*cache.RawResponse).GetRawResources(), want)
	case <-time.After(time.Second):
		t.Fatal("failed to receive snapshot response")
	}

	// Request additional resource with name=clusterName2 for same version
	go func() {
		req := &discovery.DiscoveryRequest{
			TypeUrl: rsrc.ClusterType, VersionInfo: fixture.version,
			ResourceNames: []string{clusterName, clusterName2},
			ResponseNonce: "1",
		}
		sub := subFromRequest(req)
		sub.SetReturnedResources(map[string]string{clusterName: fixture.version})
		_, err := c.CreateWatch(req, sub, watch)
		require.NoError(t, err)
	}()

	select {
	case out := <-watch:
		gotVersion, _ := out.GetVersion()
		assert.Equalf(t, gotVersion, fixture.version, "got version %q, want %q", gotVersion, fixture.version)
		assert.ElementsMatch(t, testResources[rsrc.ClusterType], out.(*cache.RawResponse).GetRawResources())
	case <-time.After(time.Second):
		t.Fatal("failed to receive snapshot response")
	}

	// Repeat request for with same version and make sure a watch is created
	req := &discovery.DiscoveryRequest{
		TypeUrl: rsrc.ClusterType, VersionInfo: fixture.version,
		ResourceNames: []string{clusterName, clusterName2},
		ResponseNonce: "2",
	}
	sub := subFromRequest(req)
	sub.SetReturnedResources(map[string]string{clusterName: fixture.version, clusterName2: fixture.version})
	cancel, err := c.CreateWatch(req, sub, watch)
	require.NoError(t, err)
	if cancel == nil {
		t.Fatal("Should create a watch")
	} else {
		cancel()
	}
}

func TestSnapshotClear(t *testing.T) {
	c := cache.NewSnapshotCache(true, group{}, log.NewTestLogger(t))
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

func (s *singleResourceSnapshot) GetTypeSnapshot(typeURL string) types.TypeSnapshot {
	if typeURL != s.typeurl {
		return types.TypeSnapshot{}
	}

	ttl := 5 * time.Second
	res := types.SnapshotResource{
		Name:     s.name,
		Resource: s.resource,
		TTL:      &ttl,
	}
	snapshot, _ := types.NewTypeSnapshot(typeURL, s.version, []types.SnapshotResource{res})
	return snapshot
}

// TestSnapshotSingleResourceFetch is a basic test to verify that simple
// cache functions work with a type that is not `Snapshot`.
func TestSnapshotSingleResourceFetch(t *testing.T) {
	durationTypeURL := "type.googleapis.com/" + string(proto.MessageName(&durationpb.Duration{}))

	anyDuration := func(d time.Duration) *anypb.Any {
		bytes, err := proto.MarshalOptions{Deterministic: true}.Marshal(durationpb.New(d))
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

	c := cache.NewSnapshotCache(true, group{}, log.NewTestLogger(t))
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
	ss := stream.NewSotwSubscription([]string{"rtds"}, true)
	ss.SetReturnedResources(map[string]string{"cluster": "abcdef"})
	responder := make(chan cache.Response)
	_, err := c.CreateWatch(req, ss, responder)
	require.NoError(t, err)

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

func TestSotwSnapshotCacheWithODCDS(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, log.NewTestLogger(t))

	// Create snapshot with mixed wildcard eligibility (ODCDS)
	wildcardCluster := resource.MakeCluster(resource.Ads, "wildcard-cluster")
	onDemandCluster := resource.MakeCluster(resource.Ads, "on-demand-cluster")
	onDemandCluster2 := resource.MakeCluster(resource.Ads, "on-demand-cluster2")

	snapshot, err := types.NewSnapshot("v1", map[string][]types.SnapshotResource{
		rsrc.ClusterType: {
			{Name: "wildcard-cluster", Resource: wildcardCluster, OnDemandOnly: false},
			{Name: "on-demand-cluster", Resource: onDemandCluster, OnDemandOnly: true},
			{Name: "on-demand-cluster2", Resource: onDemandCluster2, OnDemandOnly: true},
		},
	})
	require.NoError(t, err)
	require.NoError(t, c.SetSnapshot(context.Background(), key, snapshot))

	t.Run("wildcard subscription only returns wildcard resources", func(t *testing.T) {
		value := make(chan cache.Response, 1)
		req := &discovery.DiscoveryRequest{TypeUrl: rsrc.ClusterType, ResourceNames: []string{"*"}}
		sub := stream.NewSotwSubscription([]string{"*"}, false)
		_, err := c.CreateWatch(req, sub, value)
		require.NoError(t, err)

		select {
		case out := <-value:
			resources := out.(*cache.RawResponse).GetRawResources()
			require.Len(t, resources, 1)
			assert.Equal(t, "wildcard-cluster", cache.GetResourceName(resources[0].Resource))
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for response")
		}
	})

	t.Run("ODCDS: wildcard + explicit returns union", func(t *testing.T) {
		value := make(chan cache.Response, 1)
		req := &discovery.DiscoveryRequest{TypeUrl: rsrc.ClusterType, ResourceNames: []string{"*", "on-demand-cluster"}}
		sub := stream.NewSotwSubscription([]string{"*", "on-demand-cluster"}, false)
		_, err := c.CreateWatch(req, sub, value)
		require.NoError(t, err)

		select {
		case out := <-value:
			resources := out.(*cache.RawResponse).GetRawResources()
			require.Len(t, resources, 2)
			names := []string{cache.GetResourceName(resources[0].Resource), cache.GetResourceName(resources[1].Resource)}
			assert.ElementsMatch(t, []string{"wildcard-cluster", "on-demand-cluster"}, names)
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for response")
		}
	})

	t.Run("explicit only returns requested resource", func(t *testing.T) {
		value := make(chan cache.Response, 1)
		req := &discovery.DiscoveryRequest{TypeUrl: rsrc.ClusterType, ResourceNames: []string{"on-demand-cluster2"}}
		sub := stream.NewSotwSubscription([]string{"on-demand-cluster2"}, false)
		_, err := c.CreateWatch(req, sub, value)
		require.NoError(t, err)

		select {
		case out := <-value:
			resources := out.(*cache.RawResponse).GetRawResources()
			require.Len(t, resources, 1)
			assert.Equal(t, "on-demand-cluster2", cache.GetResourceName(resources[0].Resource))
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for response")
		}
	})
}

func TestSnapshotCacheDeltaWithODCDS(t *testing.T) {
	c := cache.NewSnapshotCache(false, group{}, log.NewTestLogger(t))

	// Create snapshot with mixed wildcard eligibility (ODCDS)
	wildcardCluster := resource.MakeCluster(resource.Ads, "wildcard-cluster")
	onDemandCluster := resource.MakeCluster(resource.Ads, "on-demand-cluster")
	onDemandCluster2 := resource.MakeCluster(resource.Ads, "on-demand-cluster2")

	snapshot, err := types.NewSnapshot("v1", map[string][]types.SnapshotResource{
		rsrc.ClusterType: {
			{Name: "wildcard-cluster", Resource: wildcardCluster, OnDemandOnly: false},
			{Name: "on-demand-cluster", Resource: onDemandCluster, OnDemandOnly: true},
			{Name: "on-demand-cluster2", Resource: onDemandCluster2, OnDemandOnly: true},
		},
	})
	require.NoError(t, err)
	require.NoError(t, c.SetSnapshot(context.Background(), key, snapshot))

	t.Run("delta wildcard subscription only returns wildcard resources", func(t *testing.T) {
		value := make(chan cache.DeltaResponse, 1)
		req := &discovery.DeltaDiscoveryRequest{
			TypeUrl:                rsrc.ClusterType,
			ResourceNamesSubscribe: []string{"*"},
			Node:                   &core.Node{Id: key},
		}
		sub := stream.NewDeltaSubscription([]string{"*"}, nil, nil, false)
		_, err := c.CreateDeltaWatch(req, sub, value)
		require.NoError(t, err)

		select {
		case out := <-value:
			resources := out.(*cache.RawDeltaResponse).GetRawResources()
			require.Len(t, resources, 1)
			assert.Equal(t, "wildcard-cluster", cache.GetResourceName(resources[0].Resource))
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for response")
		}
	})

	t.Run("delta ODCDS: wildcard + explicit returns union", func(t *testing.T) {
		value := make(chan cache.DeltaResponse, 1)
		req := &discovery.DeltaDiscoveryRequest{
			TypeUrl:                rsrc.ClusterType,
			ResourceNamesSubscribe: []string{"*", "on-demand-cluster"},
			Node:                   &core.Node{Id: key},
		}
		sub := stream.NewDeltaSubscription([]string{"*", "on-demand-cluster"}, nil, nil, false)
		_, err := c.CreateDeltaWatch(req, sub, value)
		require.NoError(t, err)

		select {
		case out := <-value:
			resources := out.(*cache.RawDeltaResponse).GetRawResources()
			require.Len(t, resources, 2)
			names := []string{cache.GetResourceName(resources[0].Resource), cache.GetResourceName(resources[1].Resource)}
			assert.ElementsMatch(t, []string{"wildcard-cluster", "on-demand-cluster"}, names)
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for response")
		}
	})
}
