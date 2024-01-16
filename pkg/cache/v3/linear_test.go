// Copyright 2020 Envoyproxy Authors
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

package cache

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/log"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

const (
	testType = "google.protobuf.StringValue"
)

func testResource(s string) types.Resource {
	return wrapperspb.String(s)
}

func verifyResponseContent(t *testing.T, ch <-chan Response, expectedType, expectedVersion string) (Response, *discovery.DiscoveryResponse) {
	t.Helper()
	var r Response
	select {
	case r = <-ch:
	case <-time.After(1 * time.Second):
		t.Error("failed to receive response after 1 second")
		return nil, nil
	}

	if r.GetRequest().GetTypeUrl() != expectedType {
		t.Errorf("unexpected request type URL: %q", r.GetRequest().GetTypeUrl())
	}
	if r.GetContext() == nil {
		t.Errorf("unexpected empty response context")
	}
	out, err := r.GetDiscoveryResponse()
	if err != nil {
		t.Fatal(err)
	}
	if out.GetVersionInfo() == "" {
		t.Error("unexpected response empty version")
	}
	if expectedVersion != "" && out.GetVersionInfo() != expectedVersion {
		t.Errorf("unexpected version: got %q, want %q", out.GetVersionInfo(), expectedVersion)
	}
	if out.GetTypeUrl() != expectedType {
		t.Errorf("unexpected type URL: %q", out.GetTypeUrl())
	}
	if len(r.GetRequest().GetResourceNames()) != 0 && len(r.GetRequest().GetResourceNames()) < len(out.Resources) {
		t.Errorf("received more resources (%d) than requested (%d)", len(r.GetRequest().GetResourceNames()), len(out.Resources))
	}
	return r, out
}

func verifyResponse(t *testing.T, ch <-chan Response, expectedVersion string, expectedResourcesNb int) {
	t.Helper()
	_, r := verifyResponseContent(t, ch, testType, expectedVersion)
	if r == nil {
		return
	}
	if n := len(r.GetResources()); n != expectedResourcesNb {
		t.Errorf("unexpected number of responses: got %d, want %d", n, expectedResourcesNb)
	}
}

func verifyResponseResources(t *testing.T, ch <-chan Response, expectedType, expectedVersion string, expectedResources ...string) {
	t.Helper()
	r, _ := verifyResponseContent(t, ch, expectedType, expectedVersion)
	if r == nil {
		return
	}
	out := r.(*RawResponse)
	resourceNames := []string{}
	for _, res := range out.Resources {
		resourceNames = append(resourceNames, GetResourceName(res.Resource))
	}
	assert.ElementsMatch(t, resourceNames, expectedResources)
}

type resourceInfo struct {
	name    string
	version string
}

func validateDeltaResponse(t *testing.T, resp DeltaResponse, resources []resourceInfo, deleted []string) {
	t.Helper()

	if resp.GetDeltaRequest().GetTypeUrl() != testType {
		t.Errorf("unexpected empty request type URL: %q", resp.GetDeltaRequest().GetTypeUrl())
	}
	out, err := resp.GetDeltaDiscoveryResponse()
	if err != nil {
		t.Fatal(err)
	}
	if len(out.GetResources()) != len(resources) {
		t.Errorf("unexpected number of responses: got %d, want %d", len(out.GetResources()), len(resources))
	}
	for _, r := range resources {
		found := false
		for _, r1 := range out.GetResources() {
			if r1.GetName() == r.name && r1.GetVersion() == r.version {
				found = true
				break
			} else if r1.GetName() == r.name {
				t.Errorf("unexpected version for resource %q: got %q, want %q", r.name, r1.GetVersion(), r.version)
				found = true
				break
			}
		}
		if !found {
			t.Errorf("resource with name %q not found in response", r.name)
		}
	}
	if out.GetTypeUrl() != testType {
		t.Errorf("unexpected type URL: %q", out.GetTypeUrl())
	}
	if len(out.GetRemovedResources()) != len(deleted) {
		t.Errorf("unexpected number of removed resurces: got %d, want %d", len(out.GetRemovedResources()), len(deleted))
	}
	for _, r := range deleted {
		found := false
		for _, rr := range out.GetRemovedResources() {
			if r == rr {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected resource %s to be deleted", r)
		}
	}
}

func verifyDeltaResponse(t *testing.T, ch <-chan DeltaResponse, resources []resourceInfo, deleted []string) {
	t.Helper()
	var r DeltaResponse
	select {
	case r = <-ch:
	case <-time.After(5 * time.Second):
		t.Error("timeout waiting for delta response")
		return
	}
	validateDeltaResponse(t, r, resources, deleted)
}

func checkWatchCount(t *testing.T, c *LinearCache, name string, count int) {
	t.Helper()
	if i := c.NumWatches(name); i != count {
		t.Errorf("unexpected number of watches for %q: got %d, want %d", name, i, count)
	}
}

func checkDeltaWatchCount(t *testing.T, c *LinearCache, count int) {
	t.Helper()
	if i := c.NumDeltaWatches(); i != count {
		t.Errorf("unexpected number of delta watches: got %d, want %d", i, count)
	}
}

func checkVersionMapNotSet(t *testing.T, c *LinearCache) {
	t.Helper()
	if c.versionMap != nil {
		t.Errorf("version map is set on the cache with %d elements", len(c.versionMap))
	}
}

func checkVersionMapSet(t *testing.T, c *LinearCache) {
	t.Helper()
	if c.versionMap == nil {
		t.Errorf("version map is not set on the cache")
	} else if len(c.versionMap) != len(c.resources) {
		t.Errorf("version map has the wrong number of elements: %d instead of %d expected", len(c.versionMap), len(c.resources))
	}
}

func mustBlock(t *testing.T, w <-chan Response) {
	t.Helper()
	select {
	case <-w:
		t.Error("watch must block")
	default:
	}
}

func mustBlockDelta(t *testing.T, w <-chan DeltaResponse) {
	t.Helper()
	select {
	case <-w:
		t.Error("watch must block")
	default:
	}
}

func hashResource(t *testing.T, resource types.Resource) string {
	t.Helper()
	marshaledResource, err := MarshalResource(resource)
	if err != nil {
		t.Fatal(err)
	}
	v := HashResource(marshaledResource)
	if v == "" {
		t.Fatal(errors.New("failed to build resource version"))
	}
	return v
}

func createWildcardDeltaWatch(t *testing.T, c *LinearCache, w chan DeltaResponse) {
	t.Helper()
	sub := stream.NewDeltaSubscription(nil, nil, nil)
	_, err := c.CreateDeltaWatch(&DeltaRequest{TypeUrl: testType}, sub, w)
	require.NoError(t, err)
	resp := <-w
	sub.SetReturnedResources(resp.GetNextVersionMap())
	_, err = c.CreateDeltaWatch(&DeltaRequest{TypeUrl: testType}, sub, w) // Ensure the watch is set properly with cache values
	require.NoError(t, err)
}

func subFromRequest(req *Request) stream.Subscription {
	return stream.NewSotwSubscription(req.GetResourceNames())
}

func subFromDeltaRequest(req *DeltaRequest) stream.Subscription {
	return stream.NewDeltaSubscription(req.GetResourceNamesSubscribe(), req.GetResourceNamesUnsubscribe(), req.GetInitialResourceVersions())
}

func TestLinearInitialResources(t *testing.T) {
	c := NewLinearCache(testType, WithInitialResources(map[string]types.Resource{"a": testResource("a"), "b": testResource("b")}))
	w := make(chan Response, 1)
	req := &Request{ResourceNames: []string{"a"}, TypeUrl: testType}
	sub := subFromRequest(req)
	_, err := c.CreateWatch(req, sub, w)
	require.NoError(t, err)
	verifyResponse(t, w, "0", 1)

	req = &Request{TypeUrl: testType}
	sub = subFromRequest(req)
	_, err = c.CreateWatch(req, sub, w)
	require.NoError(t, err)
	verifyResponse(t, w, "0", 2)
	checkVersionMapNotSet(t, c)
}

func TestLinearCornerCases(t *testing.T) {
	c := NewLinearCache(testType)
	err := c.UpdateResource("a", nil)
	if err == nil {
		t.Error("expected error on nil resource")
	}
	// create an incorrect type URL request
	req := &Request{TypeUrl: "test"}
	w := make(chan Response, 1)
	sub := subFromRequest(req)
	_, err = c.CreateWatch(req, sub, w)
	require.Error(t, err, "cache should have rejected the watch")
}

func TestLinearBasic(t *testing.T) {
	c := NewLinearCache(testType)

	// Create watches before a resource is ready
	w1 := make(chan Response, 1)
	req1 := &Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}
	sub1 := subFromRequest(req1)
	_, err := c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	mustBlock(t, w1)
	checkVersionMapNotSet(t, c)

	w2 := make(chan Response, 1)
	req2 := &Request{TypeUrl: testType, VersionInfo: "0"}
	sub2 := subFromRequest(req2)
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	mustBlock(t, w2)
	checkWatchCount(t, c, "a", 2)
	checkWatchCount(t, c, "b", 1)
	require.NoError(t, c.UpdateResource("a", testResource("a")))
	checkWatchCount(t, c, "a", 0)
	checkWatchCount(t, c, "b", 0)
	verifyResponse(t, w1, "1", 1)
	verifyResponse(t, w2, "1", 1)

	// Request again, should get same response
	_, err = c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	checkWatchCount(t, c, "a", 0)
	verifyResponse(t, w1, "1", 1)
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	checkWatchCount(t, c, "a", 0)
	verifyResponse(t, w2, "1", 1)

	// Add another element and update the first, response should be different
	require.NoError(t, c.UpdateResource("b", testResource("b")))
	require.NoError(t, c.UpdateResource("a", testResource("aa")))
	_, err = c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	verifyResponse(t, w1, "3", 1)
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	verifyResponse(t, w2, "3", 2)
	// Ensure the version map was not created as we only ever used stow watches
	checkVersionMapNotSet(t, c)
}

func TestLinearSetResources(t *testing.T) {
	c := NewLinearCache(testType)

	// Create new resources
	w1 := make(chan Response, 1)
	req1 := &Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}
	sub1 := subFromRequest(req1)
	_, err := c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	mustBlock(t, w1)

	w2 := make(chan Response, 1)
	req2 := &Request{TypeUrl: testType, VersionInfo: "0"}
	sub2 := subFromRequest(req2)
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	mustBlock(t, w2)

	c.SetResources(map[string]types.Resource{
		"a": testResource("a"),
		"b": testResource("b"),
	})
	verifyResponse(t, w1, "1", 1)
	verifyResponse(t, w2, "1", 2) // the version was only incremented once for all resources

	// Add another element and update the first, response should be different
	req1.VersionInfo = "1"
	_, err = c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	mustBlock(t, w1)
	req2.VersionInfo = "1"
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	mustBlock(t, w2)

	c.SetResources(map[string]types.Resource{
		"a": testResource("aa"),
		"b": testResource("b"),
		"c": testResource("c"),
	})
	verifyResponse(t, w1, "2", 1)
	verifyResponse(t, w2, "2", 3)

	// Delete resource
	req1.VersionInfo = "2"
	_, err = c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	mustBlock(t, w1)
	req2.VersionInfo = "2"
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	mustBlock(t, w2)

	c.SetResources(map[string]types.Resource{
		"b": testResource("b"),
		"c": testResource("c"),
	})
	verifyResponse(t, w1, "", 0) // removing a resource from the set triggers existing watches for deleted resources
	verifyResponse(t, w2, "3", 2)
}

func TestLinearGetResources(t *testing.T) {
	c := NewLinearCache(testType)

	expectedResources := map[string]types.Resource{
		"a": testResource("a"),
		"b": testResource("b"),
	}

	c.SetResources(expectedResources)

	resources := c.GetResources()

	if !reflect.DeepEqual(expectedResources, resources) {
		t.Errorf("resources are not equal. got: %v want: %v", resources, expectedResources)
	}
}

func TestLinearVersionPrefix(t *testing.T) {
	c := NewLinearCache(testType, WithVersionPrefix("instance1-"))

	w := make(chan Response, 1)
	req := &Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}
	sub := subFromRequest(req)
	_, err := c.CreateWatch(req, sub, w)
	require.NoError(t, err)
	verifyResponse(t, w, "instance1-0", 0)

	require.NoError(t, c.UpdateResource("a", testResource("a")))

	req.VersionInfo = "instance1-0"
	_, err = c.CreateWatch(req, sub, w)
	require.NoError(t, err)
	verifyResponse(t, w, "instance1-1", 1)

	req.VersionInfo = "instance1-1"
	_, err = c.CreateWatch(req, sub, w)
	require.NoError(t, err)
	mustBlock(t, w)
	checkWatchCount(t, c, "a", 1)
}

func TestLinearDeletion(t *testing.T) {
	c := NewLinearCache(testType, WithInitialResources(map[string]types.Resource{"a": testResource("a"), "b": testResource("b")}))
	w := make(chan Response, 1)
	req := &Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}
	sub := subFromRequest(req)
	_, err := c.CreateWatch(req, sub, w)
	require.NoError(t, err)
	mustBlock(t, w)
	checkWatchCount(t, c, "a", 1)

	require.NoError(t, c.DeleteResource("a"))
	verifyResponse(t, w, "1", 0)
	checkWatchCount(t, c, "a", 0)

	req = &Request{TypeUrl: testType, VersionInfo: "0"}
	sub = subFromRequest(req)
	_, err = c.CreateWatch(req, sub, w)
	require.NoError(t, err)
	verifyResponse(t, w, "1", 1)
	checkWatchCount(t, c, "b", 0)
	require.NoError(t, c.DeleteResource("b"))

	req.VersionInfo = "1"
	_, err = c.CreateWatch(req, sub, w)
	require.NoError(t, err)
	verifyResponse(t, w, "2", 0)
	checkWatchCount(t, c, "b", 0)
}

func TestLinearWatchTwo(t *testing.T) {
	c := NewLinearCache(testType, WithInitialResources(map[string]types.Resource{"a": testResource("a"), "b": testResource("b")}))

	w1 := make(chan Response, 1)
	req1 := &Request{ResourceNames: []string{"a", "b"}, TypeUrl: testType, VersionInfo: "0"}
	sub1 := subFromRequest(req1)
	_, err := c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	mustBlock(t, w1)

	w2 := make(chan Response, 1)
	req2 := &Request{TypeUrl: testType, VersionInfo: "0"}
	sub2 := subFromRequest(req2)
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	mustBlock(t, w2)

	require.NoError(t, c.UpdateResource("a", testResource("aa")))
	// should only get the modified resource
	verifyResponse(t, w1, "1", 1)
	verifyResponse(t, w2, "1", 2)
}

func TestLinearCancel(t *testing.T) {
	c := NewLinearCache(testType)
	require.NoError(t, c.UpdateResource("a", testResource("a")))

	// cancel watch-all
	w1 := make(chan Response, 1)
	req1 := &Request{TypeUrl: testType, VersionInfo: "1"}
	sub1 := subFromRequest(req1)
	cancel, err := c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	mustBlock(t, w1)
	checkWatchCount(t, c, "a", 1)
	cancel()
	checkWatchCount(t, c, "a", 0)

	// cancel watch for "a"
	sub1.SetResourceSubscription([]string{"a"})
	req1.ResourceNames = []string{"a"}
	req1.VersionInfo = "1"
	cancel, err = c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	mustBlock(t, w1)
	checkWatchCount(t, c, "a", 1)
	cancel()
	checkWatchCount(t, c, "a", 0)

	// open four watches for "a" and "b" and two for all, cancel one of each, make sure the second one is unaffected
	w2 := make(chan Response, 1)
	w3 := make(chan Response, 1)
	w4 := make(chan Response, 1)
	cancel, err = c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)

	req2 := &Request{ResourceNames: []string{"b"}, TypeUrl: testType, VersionInfo: "1"}
	cancel2, err := c.CreateWatch(req2, subFromRequest(req2), w2)
	require.NoError(t, err)
	req3 := &Request{TypeUrl: testType, VersionInfo: "1"}
	cancel3, err := c.CreateWatch(req3, subFromRequest(req3), w3)
	require.NoError(t, err)
	req4 := &Request{TypeUrl: testType, VersionInfo: "1"}
	cancel4, err := c.CreateWatch(req4, subFromRequest(req4), w4)
	require.NoError(t, err)
	mustBlock(t, w1)
	mustBlock(t, w2)
	mustBlock(t, w3)
	mustBlock(t, w4)
	checkWatchCount(t, c, "a", 3)
	checkWatchCount(t, c, "b", 3)
	cancel()
	checkWatchCount(t, c, "a", 2)
	checkWatchCount(t, c, "b", 3)
	cancel3()
	checkWatchCount(t, c, "a", 1)
	checkWatchCount(t, c, "b", 2)
	cancel2()
	cancel4()
	checkWatchCount(t, c, "a", 0)
	checkWatchCount(t, c, "b", 0)
}

// TODO(mattklein123): This test requires GOMAXPROCS or -parallel >= 100. This should be
// rewritten to not require that. This is not the case in the GH actions environment.
func TestLinearConcurrentSetWatch(t *testing.T) {
	c := NewLinearCache(testType)
	n := 50
	for i := 0; i < 2*n; i++ {
		func(i int) {
			t.Run(fmt.Sprintf("worker%d", i), func(t *testing.T) {
				t.Parallel()
				id := fmt.Sprintf("%d", i)
				if i%2 == 0 {
					t.Logf("update resource %q", id)
					require.NoError(t, c.UpdateResource(id, testResource(id)))
				} else {
					id2 := fmt.Sprintf("%d", i-1)
					t.Logf("request resources %q and %q", id, id2)
					value := make(chan Response, 1)
					req := &Request{
						// Only expect one to become stale
						ResourceNames: []string{id, id2},
						VersionInfo:   "0",
						TypeUrl:       testType,
					}
					_, err := c.CreateWatch(req, subFromRequest(req), value)
					require.NoError(t, err)
					// wait until all updates apply
					verifyResponse(t, value, "", 1)
				}
			})
		}(i)
	}
}

func TestLinearDeltaWildcard(t *testing.T) {
	c := NewLinearCache(testType)
	req1 := &DeltaRequest{TypeUrl: testType}
	w1 := make(chan DeltaResponse, 1)
	_, err := c.CreateDeltaWatch(req1, subFromDeltaRequest(req1), w1)
	require.NoError(t, err)
	mustBlockDelta(t, w1)

	req2 := &DeltaRequest{TypeUrl: testType}
	w2 := make(chan DeltaResponse, 1)
	_, err = c.CreateDeltaWatch(req2, subFromDeltaRequest(req2), w2)
	require.NoError(t, err)
	mustBlockDelta(t, w1)
	checkDeltaWatchCount(t, c, 2)

	a := &endpoint.ClusterLoadAssignment{ClusterName: "a"}
	hash := hashResource(t, a)
	err = c.UpdateResource("a", a)
	require.NoError(t, err)
	checkDeltaWatchCount(t, c, 0)
	verifyDeltaResponse(t, w1, []resourceInfo{{"a", hash}}, nil)
	verifyDeltaResponse(t, w2, []resourceInfo{{"a", hash}}, nil)
}

func TestLinearDeltaExistingResources(t *testing.T) {
	c := NewLinearCache(testType)
	a := &endpoint.ClusterLoadAssignment{ClusterName: "a"}
	hashA := hashResource(t, a)
	err := c.UpdateResource("a", a)
	require.NoError(t, err)
	b := &endpoint.ClusterLoadAssignment{ClusterName: "b"}
	hashB := hashResource(t, b)
	err = c.UpdateResource("b", b)
	require.NoError(t, err)

	// watching b and c - not interested in a
	req := &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"b", "c"}}
	w := make(chan DeltaResponse, 1)
	_, err = c.CreateDeltaWatch(req, subFromDeltaRequest(req), w)
	require.NoError(t, err)
	checkDeltaWatchCount(t, c, 0)
	verifyDeltaResponse(t, w, []resourceInfo{{"b", hashB}}, []string{})

	req = &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"a", "b"}}
	w = make(chan DeltaResponse, 1)
	_, err = c.CreateDeltaWatch(req, subFromDeltaRequest(req), w)
	require.NoError(t, err)
	checkDeltaWatchCount(t, c, 0)
	verifyDeltaResponse(t, w, []resourceInfo{{"b", hashB}, {"a", hashA}}, nil)
}

func TestLinearDeltaInitialResourcesVersionSet(t *testing.T) {
	c := NewLinearCache(testType)
	a := &endpoint.ClusterLoadAssignment{ClusterName: "a"}
	hashA := hashResource(t, a)
	err := c.UpdateResource("a", a)
	require.NoError(t, err)
	b := &endpoint.ClusterLoadAssignment{ClusterName: "b"}
	hashB := hashResource(t, b)
	err = c.UpdateResource("b", b)
	require.NoError(t, err)

	req := &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"a", "b"}, InitialResourceVersions: map[string]string{"b": hashB}}
	w := make(chan DeltaResponse, 1)
	_, err = c.CreateDeltaWatch(req, subFromDeltaRequest(req), w)
	require.NoError(t, err)
	checkDeltaWatchCount(t, c, 0)
	verifyDeltaResponse(t, w, []resourceInfo{{"a", hashA}}, nil) // b is up to date and shouldn't be returned

	req = &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"a", "b"}, InitialResourceVersions: map[string]string{"a": hashA, "b": hashB}}
	w = make(chan DeltaResponse, 1)
	_, err = c.CreateDeltaWatch(req, subFromDeltaRequest(req), w)
	require.NoError(t, err)
	mustBlockDelta(t, w)
	checkDeltaWatchCount(t, c, 1)
	b = &endpoint.ClusterLoadAssignment{ClusterName: "b", Endpoints: []*endpoint.LocalityLbEndpoints{{Priority: 10}}} // new version of b
	hashB = hashResource(t, b)
	err = c.UpdateResource("b", b)
	require.NoError(t, err)
	checkDeltaWatchCount(t, c, 0)
	verifyDeltaResponse(t, w, []resourceInfo{{"b", hashB}}, nil)
}

func TestLinearDeltaResourceUpdate(t *testing.T) {
	c := NewLinearCache(testType)
	a := &endpoint.ClusterLoadAssignment{ClusterName: "a"}
	hashA := hashResource(t, a)
	err := c.UpdateResource("a", a)
	require.NoError(t, err)
	b := &endpoint.ClusterLoadAssignment{ClusterName: "b"}
	hashB := hashResource(t, b)
	err = c.UpdateResource("b", b)
	require.NoError(t, err)
	// There is currently no delta watch
	checkVersionMapNotSet(t, c)

	req := &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"a", "b"}}
	w := make(chan DeltaResponse, 1)
	_, err = c.CreateDeltaWatch(req, subFromDeltaRequest(req), w)
	require.NoError(t, err)
	checkDeltaWatchCount(t, c, 0)
	verifyDeltaResponse(t, w, []resourceInfo{{"b", hashB}, {"a", hashA}}, nil)
	checkVersionMapSet(t, c)

	req = &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"a", "b"}, InitialResourceVersions: map[string]string{"a": hashA, "b": hashB}}
	w = make(chan DeltaResponse, 1)
	_, err = c.CreateDeltaWatch(req, subFromDeltaRequest(req), w)
	require.NoError(t, err)
	mustBlockDelta(t, w)
	checkDeltaWatchCount(t, c, 1)

	a = &endpoint.ClusterLoadAssignment{ClusterName: "a", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
		{Priority: 10},
	}}
	hashA = hashResource(t, a)
	err = c.UpdateResource("a", a)
	require.NoError(t, err)
	verifyDeltaResponse(t, w, []resourceInfo{{"a", hashA}}, nil)
	checkVersionMapSet(t, c)
}

func TestLinearDeltaResourceDelete(t *testing.T) {
	c := NewLinearCache(testType)
	a := &endpoint.ClusterLoadAssignment{ClusterName: "a"}
	hashA := hashResource(t, a)
	err := c.UpdateResource("a", a)
	require.NoError(t, err)
	b := &endpoint.ClusterLoadAssignment{ClusterName: "b"}
	hashB := hashResource(t, b)
	err = c.UpdateResource("b", b)
	require.NoError(t, err)

	req := &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"a", "b"}}
	w := make(chan DeltaResponse, 1)
	_, err = c.CreateDeltaWatch(req, subFromDeltaRequest(req), w)
	require.NoError(t, err)
	checkDeltaWatchCount(t, c, 0)
	verifyDeltaResponse(t, w, []resourceInfo{{"b", hashB}, {"a", hashA}}, nil)

	req = &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"a", "b"}, InitialResourceVersions: map[string]string{"a": hashA, "b": hashB}}
	w = make(chan DeltaResponse, 1)
	_, err = c.CreateDeltaWatch(req, subFromDeltaRequest(req), w)
	require.NoError(t, err)
	mustBlockDelta(t, w)
	checkDeltaWatchCount(t, c, 1)

	a = &endpoint.ClusterLoadAssignment{ClusterName: "a", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
		{Priority: 10},
	}}
	hashA = hashResource(t, a)
	c.SetResources(map[string]types.Resource{"a": a})
	verifyDeltaResponse(t, w, []resourceInfo{{"a", hashA}}, []string{"b"})
}

func TestLinearDeltaMultiResourceUpdates(t *testing.T) {
	c := NewLinearCache(testType)

	w := make(chan DeltaResponse, 1)
	checkVersionMapNotSet(t, c)
	assert.Equal(t, 0, c.NumResources())

	// Initial update
	req := &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"a", "b"}}
	sub := subFromDeltaRequest(req)
	_, err := c.CreateDeltaWatch(req, subFromDeltaRequest(req), w)
	require.NoError(t, err)
	mustBlockDelta(t, w)
	checkDeltaWatchCount(t, c, 1)
	// The version map should now be created, even if empty
	checkVersionMapSet(t, c)
	a := &endpoint.ClusterLoadAssignment{ClusterName: "a"}
	hashA := hashResource(t, a)
	b := &endpoint.ClusterLoadAssignment{ClusterName: "b"}
	hashB := hashResource(t, b)
	err = c.UpdateResources(map[string]types.Resource{"a": a, "b": b}, nil)
	require.NoError(t, err)
	resp := <-w
	validateDeltaResponse(t, resp, []resourceInfo{{"a", hashA}, {"b", hashB}}, nil)
	checkVersionMapSet(t, c)
	assert.Equal(t, 2, c.NumResources())

	sub.SetReturnedResources(resp.GetNextVersionMap())

	// Multiple updates
	req.ResourceNamesSubscribe = nil // No change in subscription
	_, err = c.CreateDeltaWatch(req, sub, w)
	require.NoError(t, err)
	mustBlockDelta(t, w)
	checkDeltaWatchCount(t, c, 1)
	a = &endpoint.ClusterLoadAssignment{ClusterName: "a", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
		{Priority: 10},
	}}
	b = &endpoint.ClusterLoadAssignment{ClusterName: "b", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
		{Priority: 15},
	}}
	hashA = hashResource(t, a)
	hashB = hashResource(t, b)
	err = c.UpdateResources(map[string]types.Resource{"a": a, "b": b}, nil)
	require.NoError(t, err)
	resp = <-w
	validateDeltaResponse(t, resp, []resourceInfo{{"a", hashA}, {"b", hashB}}, nil)
	checkVersionMapSet(t, c)
	assert.Equal(t, 2, c.NumResources())
	sub.SetReturnedResources(resp.GetNextVersionMap())

	// Update/add/delete
	_, err = c.CreateDeltaWatch(req, sub, w)
	require.NoError(t, err)
	mustBlockDelta(t, w)
	checkDeltaWatchCount(t, c, 1)
	a = &endpoint.ClusterLoadAssignment{ClusterName: "a", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
		{Priority: 15},
	}}
	d := &endpoint.ClusterLoadAssignment{ClusterName: "d", Endpoints: []*endpoint.LocalityLbEndpoints{}} // resource created, but not watched
	hashA = hashResource(t, a)
	err = c.UpdateResources(map[string]types.Resource{"a": a, "d": d}, []string{"b"})
	require.NoError(t, err)
	assert.Contains(t, c.resources, "d", "resource with name d not found in cache")
	assert.NotContains(t, c.resources, "b", "resource with name b was found in cache")
	resp = <-w
	validateDeltaResponse(t, resp, []resourceInfo{{"a", hashA}}, []string{"b"})
	checkVersionMapSet(t, c)
	assert.Equal(t, 2, c.NumResources())
	sub.SetReturnedResources(resp.GetNextVersionMap())

	// Re-add previously deleted watched resource
	_, err = c.CreateDeltaWatch(req, sub, w)
	require.NoError(t, err)
	mustBlockDelta(t, w)
	checkDeltaWatchCount(t, c, 1)
	b = &endpoint.ClusterLoadAssignment{ClusterName: "b", Endpoints: []*endpoint.LocalityLbEndpoints{}} // recreate watched resource
	hashB = hashResource(t, b)
	err = c.UpdateResources(map[string]types.Resource{"b": b}, []string{"d"})
	require.NoError(t, err)
	assert.Contains(t, c.resources, "b", "resource with name b not found in cache")
	assert.NotContains(t, c.resources, "d", "resource with name d was found in cache")
	resp = <-w
	validateDeltaResponse(t, resp, []resourceInfo{{"b", hashB}}, nil) // d is not watched and should not be returned
	checkVersionMapSet(t, c)
	assert.Equal(t, 2, c.NumResources())
	sub.SetReturnedResources(resp.GetNextVersionMap())

	// Wildcard create/update
	createWildcardDeltaWatch(t, c, w)
	mustBlockDelta(t, w)
	checkDeltaWatchCount(t, c, 1)
	b = &endpoint.ClusterLoadAssignment{ClusterName: "b", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
		{Priority: 15},
	}}
	d = &endpoint.ClusterLoadAssignment{ClusterName: "d", Endpoints: []*endpoint.LocalityLbEndpoints{}} // resource create
	hashB = hashResource(t, b)
	hashD := hashResource(t, d)
	err = c.UpdateResources(map[string]types.Resource{"b": b, "d": d}, nil)
	require.NoError(t, err)
	verifyDeltaResponse(t, w, []resourceInfo{{"b", hashB}, {"d", hashD}}, nil)
	checkVersionMapSet(t, c)
	assert.Equal(t, 3, c.NumResources())

	// Wildcard update/delete
	createWildcardDeltaWatch(t, c, w)
	mustBlockDelta(t, w)
	checkDeltaWatchCount(t, c, 1)
	a = &endpoint.ClusterLoadAssignment{ClusterName: "a", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
		{Priority: 25},
	}}
	hashA = hashResource(t, a)
	err = c.UpdateResources(map[string]types.Resource{"a": a}, []string{"d"})
	require.NoError(t, err)
	assert.NotContains(t, c.resources, "d", "resource with name d was found in cache")
	verifyDeltaResponse(t, w, []resourceInfo{{"a", hashA}}, []string{"d"})

	checkDeltaWatchCount(t, c, 0)
	// Confirm that the map is still set even though there is currently no watch
	checkVersionMapSet(t, c)
	assert.Equal(t, 2, c.NumResources())
}

func TestLinearMixedWatches(t *testing.T) {
	c := NewLinearCache(testType)
	a := &endpoint.ClusterLoadAssignment{ClusterName: "a"}
	err := c.UpdateResource("a", a)
	require.NoError(t, err)
	b := &endpoint.ClusterLoadAssignment{ClusterName: "b"}
	hashB := hashResource(t, b)
	err = c.UpdateResource("b", b)
	require.NoError(t, err)
	assert.Equal(t, 2, c.NumResources())

	w := make(chan Response, 1)
	sotwReq := &Request{ResourceNames: []string{"a", "b"}, TypeUrl: testType, VersionInfo: c.getVersion()}
	sotwSub := subFromRequest(sotwReq)
	_, err = c.CreateWatch(sotwReq, subFromRequest(sotwReq), w)
	require.NoError(t, err)
	mustBlock(t, w)
	checkVersionMapNotSet(t, c)

	a = &endpoint.ClusterLoadAssignment{ClusterName: "a", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
		{Priority: 25},
	}}
	hashA := hashResource(t, a)
	err = c.UpdateResources(map[string]types.Resource{"a": a}, nil)
	require.NoError(t, err)
	// This behavior is currently invalid for cds and lds, but due to a current limitation of linear cache sotw implementation
	verifyResponse(t, w, c.getVersion(), 1)
	checkVersionMapNotSet(t, c)

	sotwReq.VersionInfo = c.getVersion()
	_, err = c.CreateWatch(sotwReq, sotwSub, w)
	require.NoError(t, err)
	mustBlock(t, w)
	checkVersionMapNotSet(t, c)

	deltaReq := &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"a", "b"}, InitialResourceVersions: map[string]string{"a": hashA, "b": hashB}}
	wd := make(chan DeltaResponse, 1)

	// Initial update
	_, err = c.CreateDeltaWatch(deltaReq, subFromDeltaRequest(deltaReq), wd)
	require.NoError(t, err)
	mustBlockDelta(t, wd)
	checkDeltaWatchCount(t, c, 1)
	checkVersionMapSet(t, c)

	err = c.UpdateResources(nil, []string{"b"})
	require.NoError(t, err)
	checkVersionMapSet(t, c)

	verifyResponse(t, w, c.getVersion(), 0)
	verifyDeltaResponse(t, wd, nil, []string{"b"})
}

func TestLinearSotwWatches(t *testing.T) {
	t.Run("watches are properly removed from all objects", func(t *testing.T) {
		cache := NewLinearCache(testType)
		a := &endpoint.ClusterLoadAssignment{ClusterName: "a"}
		err := cache.UpdateResource("a", a)
		require.NoError(t, err)
		b := &endpoint.ClusterLoadAssignment{ClusterName: "b"}
		err = cache.UpdateResource("b", b)
		require.NoError(t, err)
		assert.Equal(t, 2, cache.NumResources())

		// A watch tracks three different objects.
		// An update is done for the three objects in a row
		// If the watches are no properly purged, all three updates will send responses in the channel, but only the first one is tracked
		// The buffer will therefore saturate and the third request will deadlock the entire cache as occurring under the mutex
		w := make(chan Response, 1)
		sotwReq := &Request{ResourceNames: []string{"a", "b", "c"}, TypeUrl: testType, VersionInfo: cache.getVersion()}
		sotwSub := subFromRequest(sotwReq)
		_, err = cache.CreateWatch(sotwReq, sotwSub, w)
		require.NoError(t, err)
		mustBlock(t, w)
		checkVersionMapNotSet(t, cache)

		assert.Len(t, cache.watches["a"], 1)
		assert.Len(t, cache.watches["b"], 1)
		assert.Len(t, cache.watches["c"], 1)

		// Update a and c without touching b
		a = &endpoint.ClusterLoadAssignment{ClusterName: "a", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
			{Priority: 25},
		}}
		err = cache.UpdateResources(map[string]types.Resource{"a": a}, nil)
		require.NoError(t, err)
		verifyResponseResources(t, w, testType, cache.getVersion(), "a")
		checkVersionMapNotSet(t, cache)

		assert.Empty(t, cache.watches["a"])
		assert.Empty(t, cache.watches["b"])
		assert.Empty(t, cache.watches["c"])

		// c no longer watched
		w = make(chan Response, 1)
		sotwReq.ResourceNames = []string{"a", "b"}
		sotwReq.VersionInfo = cache.getVersion()
		sotwSub.SetResourceSubscription([]string{"a", "b"})
		_, err = cache.CreateWatch(sotwReq, sotwSub, w)
		require.NoError(t, err)
		mustBlock(t, w)
		checkVersionMapNotSet(t, cache)

		b = &endpoint.ClusterLoadAssignment{ClusterName: "b", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
			{Priority: 15},
		}}
		err = cache.UpdateResources(map[string]types.Resource{"b": b}, nil)

		assert.Empty(t, cache.watches["a"])
		assert.Empty(t, cache.watches["b"])
		assert.Empty(t, cache.watches["c"])

		require.NoError(t, err)
		verifyResponseResources(t, w, testType, cache.getVersion(), "b")
		checkVersionMapNotSet(t, cache)

		w = make(chan Response, 1)
		sotwReq.ResourceNames = []string{"c"}
		sotwReq.VersionInfo = cache.getVersion()
		sotwSub.SetResourceSubscription([]string{"c"})
		_, err = cache.CreateWatch(sotwReq, sotwSub, w)
		require.NoError(t, err)
		mustBlock(t, w)
		checkVersionMapNotSet(t, cache)

		c := &endpoint.ClusterLoadAssignment{ClusterName: "c", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
			{Priority: 15},
		}}
		err = cache.UpdateResources(map[string]types.Resource{"c": c}, nil)
		require.NoError(t, err)
		verifyResponseResources(t, w, testType, cache.getVersion(), "c")
		checkVersionMapNotSet(t, cache)

		assert.Empty(t, cache.watches["a"])
		assert.Empty(t, cache.watches["b"])
		assert.Empty(t, cache.watches["c"])
	})

	t.Run("watches return full state for types requesting it", func(t *testing.T) {
		a := &cluster.Cluster{Name: "a"}
		b := &cluster.Cluster{Name: "b"}
		c := &cluster.Cluster{Name: "c"}
		// ClusterType requires all resources to always be returned
		cache := NewLinearCache(resource.ClusterType, WithInitialResources(map[string]types.Resource{
			"a": a,
			"b": b,
			"c": c,
		}), WithLogger(log.NewTestLogger(t)))
		assert.Equal(t, 3, cache.NumResources())

		// Non-wildcard request
		nonWildcardReq := &Request{ResourceNames: []string{"a", "b", "d"}, TypeUrl: resource.ClusterType, VersionInfo: cache.getVersion()}
		nonWildcardSub := subFromRequest(nonWildcardReq)
		w1 := make(chan Response, 1)
		_, err := cache.CreateWatch(nonWildcardReq, nonWildcardSub, w1)
		require.NoError(t, err)
		mustBlock(t, w1)
		checkVersionMapNotSet(t, cache)

		// wildcard request
		wildcardReq := &Request{ResourceNames: nil, TypeUrl: resource.ClusterType, VersionInfo: cache.getVersion()}
		wildcardSub := subFromRequest(wildcardReq)
		w2 := make(chan Response, 1)
		_, err = cache.CreateWatch(wildcardReq, wildcardSub, w2)
		require.NoError(t, err)
		mustBlock(t, w2)
		checkVersionMapNotSet(t, cache)

		// request not requesting b
		otherReq := &Request{ResourceNames: []string{"a", "c", "d"}, TypeUrl: resource.ClusterType, VersionInfo: cache.getVersion()}
		otherSub := subFromRequest(otherReq)
		w3 := make(chan Response, 1)
		_, err = cache.CreateWatch(otherReq, otherSub, w3)
		require.NoError(t, err)
		mustBlock(t, w3)
		checkVersionMapNotSet(t, cache)

		b.AltStatName = "othername"
		err = cache.UpdateResources(map[string]types.Resource{"b": b}, nil)
		require.NoError(t, err)

		// Other watch has not triggered
		mustBlock(t, w3)

		verifyResponseResources(t, w1, resource.ClusterType, cache.getVersion(), "a", "b")      // a is also returned as cluster requires full state
		verifyResponseResources(t, w2, resource.ClusterType, cache.getVersion(), "a", "b", "c") // a and c are also returned wildcard

		// Recreate the watches
		w1 = make(chan Response, 1)
		nonWildcardReq.VersionInfo = cache.getVersion()
		_, err = cache.CreateWatch(nonWildcardReq, nonWildcardSub, w1)
		require.NoError(t, err)
		mustBlock(t, w1)
		w2 = make(chan Response, 1)
		wildcardReq.VersionInfo = cache.getVersion()
		_, err = cache.CreateWatch(wildcardReq, wildcardSub, w2)
		require.NoError(t, err)
		mustBlock(t, w2)

		// Update d, new resource in the cache
		d := &cluster.Cluster{Name: "d"}
		err = cache.UpdateResource("d", d)
		require.NoError(t, err)

		verifyResponseResources(t, w1, resource.ClusterType, cache.getVersion(), "a", "b", "d")
		verifyResponseResources(t, w2, resource.ClusterType, cache.getVersion(), "a", "b", "c", "d")
		verifyResponseResources(t, w3, resource.ClusterType, cache.getVersion(), "a", "c", "d")
	})
}
