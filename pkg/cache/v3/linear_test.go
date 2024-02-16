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

type validationContext struct {
	expectedType string
}

func newValidationContext(opts []validateOption) validationContext {
	context := validationContext{
		expectedType: testType,
	}
	for _, opt := range opts {
		opt(&context)
	}
	return context
}

type validateOption = func(*validationContext)

func responseType(t string) validateOption {
	return func(vo *validationContext) {
		vo.expectedType = t
	}
}

func testResource(s string) types.Resource {
	return wrapperspb.String(s)
}

func verifyResponseContent(t *testing.T, ch <-chan Response, expectedType, expectedVersion string) (Response, *discovery.DiscoveryResponse) {
	t.Helper()
	var r Response
	select {
	case r = <-ch:
	case <-time.After(1 * time.Second):
		t.Fatal("failed to receive response after 1 second")
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
	return r, out
}

func verifyResponse(t *testing.T, ch <-chan Response, expectedVersion string, expectedResourcesNb int) Response {
	t.Helper()
	response, r := verifyResponseContent(t, ch, testType, expectedVersion)
	if r == nil {
		return nil
	}
	if n := len(r.GetResources()); n != expectedResourcesNb {
		t.Errorf("unexpected number of responses: got %d, want %d", n, expectedResourcesNb)
	}
	return response
}

func verifyResponseResources(t *testing.T, ch <-chan Response, expectedType, expectedVersion string, expectedResources ...string) Response {
	t.Helper()
	r, _ := verifyResponseContent(t, ch, expectedType, expectedVersion)
	if r == nil {
		return nil
	}
	out := r.(*RawResponse)
	resourceNames := []string{}
	for _, res := range out.Resources {
		resourceNames = append(resourceNames, GetResourceName(res.Resource))
	}
	assert.ElementsMatch(t, resourceNames, expectedResources)
	return r
}

type resourceInfo struct {
	name    string
	version string
}

func validateDeltaResponse(t *testing.T, resp DeltaResponse, resources []resourceInfo, deleted []string, options ...validateOption) {
	t.Helper()

	validationCtx := newValidationContext(options)

	if resp.GetDeltaRequest().GetTypeUrl() != validationCtx.expectedType {
		t.Errorf("unexpected request type URL: received %s and expected %s", resp.GetDeltaRequest().GetTypeUrl(), validationCtx.expectedType)
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
	if out.GetTypeUrl() != validationCtx.expectedType {
		t.Errorf("unexpected type URL: received %s and expected %s", out.GetTypeUrl(), validationCtx.expectedType)
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

func verifyDeltaResponse(t *testing.T, ch <-chan DeltaResponse, resources []resourceInfo, deleted []string, options ...validateOption) DeltaResponse {
	t.Helper()
	var r DeltaResponse
	select {
	case r = <-ch:
	case <-time.After(1 * time.Second):
		t.Error("timeout waiting for delta response")
		return nil
	}
	validateDeltaResponse(t, r, resources, deleted, options...)
	return r
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

func checkStableVersionsAreNotComputed(t *testing.T, c *LinearCache, resources ...string) {
	t.Helper()
	for _, res := range resources {
		assert.Empty(t, c.resources[res].stableVersion, "stable version not set on resource %s", res)
	}
}

func checkStableVersionsAreComputed(t *testing.T, c *LinearCache, resources ...string) {
	t.Helper()
	for _, res := range resources {
		assert.NotEmpty(t, c.resources[res].stableVersion, "stable version not set on resource %s", res)
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
	return HashResource(marshaledResource)
}

func createWildcardDeltaWatch(t *testing.T, initialReq bool, c *LinearCache, w chan DeltaResponse) {
	t.Helper()
	sub := stream.NewDeltaSubscription(nil, nil, nil)
	req := &DeltaRequest{TypeUrl: testType}
	if !initialReq {
		req.ResponseNonce = "1"
	}
	_, err := c.CreateDeltaWatch(req, sub, w)
	require.NoError(t, err)
	resp := <-w
	sub.SetReturnedResources(resp.GetNextVersionMap())
	_, err = c.CreateDeltaWatch(req, sub, w) // Ensure the watch is set properly with cache values
	require.NoError(t, err)
}

func subFromRequest(req *Request) stream.Subscription {
	return stream.NewSotwSubscription(req.GetResourceNames())
}

// This method represents the expected behavior of client and servers regarding the request and the subscription.
// For edge cases it should ignore those
func updateFromSotwResponse(resp Response, sub *stream.Subscription, req *Request) {
	sub.SetReturnedResources(resp.GetReturnedResources())
	// Never returns an error when not using passthrough responses
	version, _ := resp.GetVersion()
	req.VersionInfo = version
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

	resp1 := verifyResponse(t, w1, "1", 1)
	updateFromSotwResponse(resp1, &sub1, req1)
	resp2 := verifyResponse(t, w2, "1", 1)
	updateFromSotwResponse(resp2, &sub2, req2)

	// Request again, should get same response
	w1 = make(chan Response, 1)
	req1 = &Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}
	sub1 = subFromRequest(req1)
	_, err = c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	checkWatchCount(t, c, "a", 0)
	resp1 = verifyResponse(t, w1, "1", 1)
	updateFromSotwResponse(resp1, &sub1, req1)

	w2 = make(chan Response, 1)
	req2 = &Request{TypeUrl: testType, VersionInfo: "0"}
	sub2 = subFromRequest(req2)
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	checkWatchCount(t, c, "a", 0)
	resp2 = verifyResponse(t, w2, "1", 1)
	updateFromSotwResponse(resp2, &sub2, req2)

	// Add another element and update the first, response should be different
	require.NoError(t, c.UpdateResource("b", testResource("b")))
	require.NoError(t, c.UpdateResource("a", testResource("aa")))
	_, err = c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	verifyResponse(t, w1, "3", 1)
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	verifyResponse(t, w2, "3", 2)
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
	resp1 := verifyResponse(t, w1, "1", 1)
	updateFromSotwResponse(resp1, &sub1, req1)
	resp2 := verifyResponse(t, w2, "1", 2) // the version was only incremented once for all resources
	updateFromSotwResponse(resp2, &sub2, req2)

	// Add another element and update the first, response should be different
	_, err = c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	mustBlock(t, w1)
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	mustBlock(t, w2)

	c.SetResources(map[string]types.Resource{
		"a": testResource("aa"),
		"b": testResource("b"),
		"c": testResource("c"),
	})
	resp1 = verifyResponse(t, w1, "2", 1)
	updateFromSotwResponse(resp1, &sub1, req1)
	resp2 = verifyResponse(t, w2, "2", 3)
	updateFromSotwResponse(resp2, &sub2, req2)

	// Delete resource
	_, err = c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	mustBlock(t, w1)
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	mustBlock(t, w2)

	c.SetResources(map[string]types.Resource{
		"b": testResource("b"),
		"c": testResource("c"),
	})
	mustBlock(t, w1) // removing a resource from the set does not trigger the watch for non full state resources
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
	resp := verifyResponse(t, w, "instance1-0", 0)
	updateFromSotwResponse(resp, &sub, req)

	require.NoError(t, c.UpdateResource("a", testResource("a")))

	_, err = c.CreateWatch(req, sub, w)
	require.NoError(t, err)
	resp = verifyResponse(t, w, "instance1-1", 1)

	req.VersionInfo = "instance1-1"
	sub.SetReturnedResources(resp.GetReturnedResources())
	w = make(chan Response, 1)
	_, err = c.CreateWatch(req, sub, w)
	require.NoError(t, err)
	mustBlock(t, w)
	checkWatchCount(t, c, "a", 1)
}

func TestLinearDeletion(t *testing.T) {
	t.Run("non full-state resource", func(t *testing.T) {
		c := NewLinearCache(testType, WithInitialResources(map[string]types.Resource{"a": testResource("a"), "b": testResource("b")}), WithLogger(log.NewTestLogger(t)))
		w := make(chan Response, 1)
		req := &Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}
		sub := subFromRequest(req)
		sub.SetReturnedResources(map[string]string{"a": "0"})
		cancel, err := c.CreateWatch(req, sub, w)
		require.NoError(t, err)
		mustBlock(t, w)
		checkWatchCount(t, c, "a", 1)

		require.NoError(t, c.DeleteResource("a"))
		// For non full-state type, we don't return on deletion
		mustBlock(t, w)

		cancel()
		checkWatchCount(t, c, "a", 0)

		// Create a wildcard watch
		req = &Request{TypeUrl: testType, VersionInfo: "0"}
		sub = subFromRequest(req)
		_, err = c.CreateWatch(req, sub, w)
		require.NoError(t, err)
		resp := verifyResponse(t, w, "1", 1)
		updateFromSotwResponse(resp, &sub, req)
		checkWatchCount(t, c, "b", 0)
		require.NoError(t, c.DeleteResource("b"))

		req.VersionInfo = "1"
		_, err = c.CreateWatch(req, sub, w)
		require.NoError(t, err)
		// b is watched by wildcard, but for non-full-state resources we cannot report deletions
		mustBlock(t, w)
		assert.Len(t, c.wildcardWatches.sotw, 1)
	})

	t.Run("full-state resource", func(t *testing.T) {
		c := NewLinearCache(resource.ClusterType, WithInitialResources(map[string]types.Resource{"a": &cluster.Cluster{Name: "a"}, "b": &cluster.Cluster{Name: "b"}}))
		w := make(chan Response, 1)
		req := &Request{ResourceNames: []string{"a"}, TypeUrl: resource.ClusterType, VersionInfo: "0"}
		sub := subFromRequest(req)
		sub.SetReturnedResources(map[string]string{"a": "0"})
		_, err := c.CreateWatch(req, sub, w)
		require.NoError(t, err)
		mustBlock(t, w)
		checkWatchCount(t, c, "a", 1)

		require.NoError(t, c.DeleteResource("a"))
		// We get a response with no resource as full update and only a was requested
		resp := verifyResponseResources(t, w, resource.ClusterType, "1")
		updateFromSotwResponse(resp, &sub, req)
		checkWatchCount(t, c, "a", 0)

		// New wildcard request
		req = &Request{TypeUrl: resource.ClusterType, VersionInfo: "0"}
		sub = subFromRequest(req)
		_, err = c.CreateWatch(req, sub, w)
		require.NoError(t, err)
		// b still exists in the cache
		resp = verifyResponseResources(t, w, resource.ClusterType, "1", "b")
		updateFromSotwResponse(resp, &sub, req)
		checkWatchCount(t, c, "b", 0)
		require.NoError(t, c.DeleteResource("b"))

		req.VersionInfo = "1"
		_, err = c.CreateWatch(req, sub, w)
		require.NoError(t, err)
		// The cache no longer contains any resource, and as full-state is requested a response is provided
		_ = verifyResponseResources(t, w, resource.ClusterType, "2")
		checkWatchCount(t, c, "b", 0)
	})
}

func TestLinearWatchTwo(t *testing.T) {
	c := NewLinearCache(testType, WithInitialResources(map[string]types.Resource{"a": testResource("a"), "b": testResource("b")}))

	w1 := make(chan Response, 1)
	req1 := &Request{ResourceNames: []string{"a", "b"}, TypeUrl: testType, VersionInfo: "0"}
	sub1 := subFromRequest(req1)
	sub1.SetReturnedResources(map[string]string{"a": "0", "b": "0"})
	_, err := c.CreateWatch(req1, sub1, w1)
	require.NoError(t, err)
	mustBlock(t, w1)

	w2 := make(chan Response, 1)
	req2 := &Request{TypeUrl: testType, VersionInfo: "0"}
	sub2 := subFromRequest(req2)
	sub2.SetReturnedResources(map[string]string{"a": "0", "b": "0"})
	_, err = c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	mustBlock(t, w2)

	require.NoError(t, c.UpdateResource("a", testResource("aa")))
	// should only get the modified resource
	verifyResponse(t, w1, "1", 1)
	verifyResponse(t, w2, "1", 1)
}

func TestLinearCancel(t *testing.T) {
	c := NewLinearCache(testType)
	require.NoError(t, c.UpdateResource("a", testResource("a")))

	// cancel watch-all
	w1 := make(chan Response, 1)
	req1 := &Request{TypeUrl: testType, VersionInfo: "1"}
	sub1 := subFromRequest(req1)
	sub1.SetReturnedResources(map[string]string{"a": "1"})
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
	sub2 := subFromRequest(req2)
	cancel2, err := c.CreateWatch(req2, sub2, w2)
	require.NoError(t, err)
	req3 := &Request{TypeUrl: testType, VersionInfo: "1"}
	sub3 := subFromRequest(req3)
	sub3.SetReturnedResources(map[string]string{"a": "1"})
	cancel3, err := c.CreateWatch(req3, sub3, w3)
	require.NoError(t, err)
	req4 := &Request{TypeUrl: testType, VersionInfo: "1"}
	sub4 := subFromRequest(req4)
	sub4.SetReturnedResources(map[string]string{"a": "1"})
	cancel4, err := c.CreateWatch(req4, sub4, w4)
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
	sub1 := subFromDeltaRequest(req1)
	_, err := c.CreateDeltaWatch(req1, sub1, w1)
	require.NoError(t, err)
	resp1 := verifyDeltaResponse(t, w1, nil, nil)
	sub1.SetReturnedResources(resp1.GetNextVersionMap())
	req1.ResponseNonce = "1"

	req2 := &DeltaRequest{TypeUrl: testType}
	w2 := make(chan DeltaResponse, 1)
	sub2 := subFromDeltaRequest(req2)
	_, err = c.CreateDeltaWatch(req2, sub2, w2)
	require.NoError(t, err)
	resp2 := verifyDeltaResponse(t, w2, nil, nil)
	sub2.SetReturnedResources(resp2.GetNextVersionMap())
	req2.ResponseNonce = "1"

	_, err = c.CreateDeltaWatch(req1, sub1, w1)
	require.NoError(t, err)
	mustBlockDelta(t, w1)

	_, err = c.CreateDeltaWatch(req2, sub2, w2)
	require.NoError(t, err)
	mustBlockDelta(t, w2)

	a := &endpoint.ClusterLoadAssignment{ClusterName: "a"}
	hash := hashResource(t, a)
	err = c.UpdateResource("a", a)
	require.NoError(t, err)
	checkDeltaWatchCount(t, c, 0)
	verifyDeltaResponse(t, w1, []resourceInfo{{"a", hash}}, nil)
	verifyDeltaResponse(t, w2, []resourceInfo{{"a", hash}}, nil)
}

func TestLinearDeltaExistingResources(t *testing.T) {
	c := NewLinearCache(testType, WithLogger(log.NewTestLogger(t)))
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
	c := NewLinearCache(testType, WithLogger(log.NewTestLogger(t)))
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
	checkStableVersionsAreNotComputed(t, c, "a", "b")

	req := &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"a", "b"}}
	w := make(chan DeltaResponse, 1)
	_, err = c.CreateDeltaWatch(req, subFromDeltaRequest(req), w)
	require.NoError(t, err)
	checkDeltaWatchCount(t, c, 0)
	verifyDeltaResponse(t, w, []resourceInfo{{"b", hashB}, {"a", hashA}}, nil)
	checkStableVersionsAreComputed(t, c, "a", "b")

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
	checkStableVersionsAreComputed(t, c, "a")
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
	assert.Equal(t, 0, c.NumResources())

	// Initial update
	req := &DeltaRequest{TypeUrl: testType, ResourceNamesSubscribe: []string{"a", "b"}}
	sub := subFromDeltaRequest(req)
	_, err := c.CreateDeltaWatch(req, subFromDeltaRequest(req), w)
	require.NoError(t, err)
	mustBlockDelta(t, w)
	checkDeltaWatchCount(t, c, 1)
	a := &endpoint.ClusterLoadAssignment{ClusterName: "a"}
	hashA := hashResource(t, a)
	b := &endpoint.ClusterLoadAssignment{ClusterName: "b"}
	hashB := hashResource(t, b)
	err = c.UpdateResources(map[string]types.Resource{"a": a, "b": b}, nil)
	require.NoError(t, err)
	resp := <-w
	validateDeltaResponse(t, resp, []resourceInfo{{"a", hashA}, {"b", hashB}}, nil)
	checkStableVersionsAreComputed(t, c, "a", "b")
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
	checkStableVersionsAreComputed(t, c, "a", "b")
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
	checkStableVersionsAreComputed(t, c, "a")
	// d is not watched currently
	checkStableVersionsAreNotComputed(t, c, "d")
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
	checkStableVersionsAreComputed(t, c, "b")
	assert.Equal(t, 2, c.NumResources())
	sub.SetReturnedResources(resp.GetNextVersionMap())

	// Wildcard create/update
	createWildcardDeltaWatch(t, false, c, w)
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
	// d is now watched and should be returned
	checkStableVersionsAreComputed(t, c, "b", "d")
	assert.Equal(t, 3, c.NumResources())

	// Wildcard update/delete
	createWildcardDeltaWatch(t, false, c, w)
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
	assert.Equal(t, 2, c.NumResources())
}

func TestLinearMixedWatches(t *testing.T) {
	c := NewLinearCache(resource.EndpointType, WithLogger(log.NewTestLogger(t)))
	a := &endpoint.ClusterLoadAssignment{ClusterName: "a"}
	err := c.UpdateResource("a", a)
	require.NoError(t, err)
	b := &endpoint.ClusterLoadAssignment{ClusterName: "b"}
	hashB := hashResource(t, b)
	err = c.UpdateResource("b", b)
	require.NoError(t, err)
	assert.Equal(t, 2, c.NumResources())

	w := make(chan Response, 1)
	sotwReq := &Request{ResourceNames: []string{"a", "b"}, TypeUrl: resource.EndpointType, VersionInfo: c.getVersion()}
	sotwSub := subFromRequest(sotwReq)
	sotwSub.SetReturnedResources(map[string]string{"a": "1", "b": "2"})
	_, err = c.CreateWatch(sotwReq, sotwSub, w)
	require.NoError(t, err)
	mustBlock(t, w)
	// Only sotw watches, should not have triggered stable resource computation
	checkStableVersionsAreNotComputed(t, c, "a", "b")

	a = &endpoint.ClusterLoadAssignment{ClusterName: "a", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
		{Priority: 25},
	}}
	hashA := hashResource(t, a)
	err = c.UpdateResources(map[string]types.Resource{"a": a}, nil)
	require.NoError(t, err)
	// This behavior is currently invalid for cds and lds, but due to a current limitation of linear cache sotw implementation
	resp := verifyResponseResources(t, w, resource.EndpointType, c.getVersion(), "a")
	updateFromSotwResponse(resp, &sotwSub, sotwReq)
	checkStableVersionsAreNotComputed(t, c, "a", "b")

	sotwReq.VersionInfo = c.getVersion()
	_, err = c.CreateWatch(sotwReq, sotwSub, w)
	require.NoError(t, err)
	mustBlock(t, w)
	checkStableVersionsAreNotComputed(t, c, "a", "b")

	deltaReq := &DeltaRequest{TypeUrl: resource.EndpointType, ResourceNamesSubscribe: []string{"a", "b"}, InitialResourceVersions: map[string]string{"a": hashA, "b": hashB}}
	wd := make(chan DeltaResponse, 1)

	// Initial update
	_, err = c.CreateDeltaWatch(deltaReq, subFromDeltaRequest(deltaReq), wd)
	require.NoError(t, err)
	mustBlockDelta(t, wd)
	checkDeltaWatchCount(t, c, 1)
	checkStableVersionsAreComputed(t, c, "a", "b")

	err = c.UpdateResources(nil, []string{"b"})
	require.NoError(t, err)
	mustBlock(t, w) // For sotw with non full-state resources, we don't report deletions
	verifyDeltaResponse(t, wd, nil, []string{"b"}, responseType(resource.EndpointType))
}

func TestLinearSotwWatches(t *testing.T) {
	t.Run("watches are properly removed from all objects", func(t *testing.T) {
		cache := NewLinearCache(testType, WithLogger(log.NewTestLogger(t)))
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
		sotwSub.SetReturnedResources(map[string]string{"a": "1", "b": "2"})
		_, err = cache.CreateWatch(sotwReq, sotwSub, w)
		require.NoError(t, err)
		mustBlock(t, w)

		assert.Len(t, cache.resourceWatches["a"].sotw, 1)
		assert.Len(t, cache.resourceWatches["b"].sotw, 1)
		assert.Len(t, cache.resourceWatches["c"].sotw, 1)

		// Update a and c without touching b
		a = &endpoint.ClusterLoadAssignment{ClusterName: "a", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
			{Priority: 25},
		}}
		err = cache.UpdateResources(map[string]types.Resource{"a": a}, nil)
		require.NoError(t, err)
		resp := verifyResponseResources(t, w, testType, cache.getVersion(), "a")
		updateFromSotwResponse(resp, &sotwSub, sotwReq)

		assert.Empty(t, cache.resourceWatches["a"].sotw)
		assert.Empty(t, cache.resourceWatches["b"].sotw)
		assert.Empty(t, cache.resourceWatches["c"].sotw)

		// c no longer watched
		w = make(chan Response, 1)
		sotwReq.ResourceNames = []string{"a", "b"}
		sotwSub.SetResourceSubscription(sotwReq.ResourceNames)
		_, err = cache.CreateWatch(sotwReq, sotwSub, w)
		require.NoError(t, err)
		mustBlock(t, w)

		b = &endpoint.ClusterLoadAssignment{ClusterName: "b", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
			{Priority: 15},
		}}
		err = cache.UpdateResources(map[string]types.Resource{"b": b}, nil)

		assert.Empty(t, cache.resourceWatches["a"].sotw)
		assert.Empty(t, cache.resourceWatches["b"].sotw)
		assert.Empty(t, cache.resourceWatches["c"].sotw)

		require.NoError(t, err)
		resp = verifyResponseResources(t, w, testType, cache.getVersion(), "b")
		updateFromSotwResponse(resp, &sotwSub, sotwReq)

		w = make(chan Response, 1)
		sotwReq.ResourceNames = []string{"c"}
		sotwSub.SetResourceSubscription(sotwReq.ResourceNames)
		_, err = cache.CreateWatch(sotwReq, sotwSub, w)
		require.NoError(t, err)
		mustBlock(t, w)

		c := &endpoint.ClusterLoadAssignment{ClusterName: "c", Endpoints: []*endpoint.LocalityLbEndpoints{ // resource update
			{Priority: 15},
		}}
		err = cache.UpdateResources(map[string]types.Resource{"c": c}, nil)
		require.NoError(t, err)
		verifyResponseResources(t, w, testType, cache.getVersion(), "c")

		assert.Empty(t, cache.resourceWatches["a"].sotw)
		assert.Empty(t, cache.resourceWatches["b"].sotw)
		assert.Empty(t, cache.resourceWatches["c"].sotw)
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
		nonWildcardSub.SetReturnedResources(map[string]string{"a": cache.getVersion(), "b": cache.getVersion()})
		w1 := make(chan Response, 1)
		_, err := cache.CreateWatch(nonWildcardReq, nonWildcardSub, w1)
		require.NoError(t, err)
		mustBlock(t, w1)

		// wildcard request
		wildcardReq := &Request{ResourceNames: nil, TypeUrl: resource.ClusterType, VersionInfo: cache.getVersion()}
		wildcardSub := subFromRequest(wildcardReq)
		wildcardSub.SetReturnedResources(map[string]string{"a": cache.getVersion(), "b": cache.getVersion(), "c": cache.getVersion()})
		w2 := make(chan Response, 1)
		_, err = cache.CreateWatch(wildcardReq, wildcardSub, w2)
		require.NoError(t, err)
		mustBlock(t, w2)

		// request not requesting b
		otherReq := &Request{ResourceNames: []string{"a", "c", "d"}, TypeUrl: resource.ClusterType, VersionInfo: cache.getVersion()}
		otherSub := subFromRequest(otherReq)
		otherSub.SetReturnedResources(map[string]string{"a": cache.getVersion(), "c": cache.getVersion()})
		w3 := make(chan Response, 1)
		_, err = cache.CreateWatch(otherReq, otherSub, w3)
		require.NoError(t, err)
		mustBlock(t, w3)

		b.AltStatName = "othername"
		err = cache.UpdateResources(map[string]types.Resource{"b": b}, nil)
		require.NoError(t, err)

		// Other watch has not triggered
		mustBlock(t, w3)

		resp1 := verifyResponseResources(t, w1, resource.ClusterType, cache.getVersion(), "a", "b") // a is also returned as cluster requires full state
		updateFromSotwResponse(resp1, &nonWildcardSub, nonWildcardReq)
		resp2 := verifyResponseResources(t, w2, resource.ClusterType, cache.getVersion(), "a", "b", "c") // a and c are also returned wildcard
		updateFromSotwResponse(resp2, &wildcardSub, wildcardReq)

		// Recreate the watches
		w1 = make(chan Response, 1)
		_, err = cache.CreateWatch(nonWildcardReq, nonWildcardSub, w1)
		require.NoError(t, err)
		mustBlock(t, w1)

		w2 = make(chan Response, 1)
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

func TestLinearSotwNonWildcard(t *testing.T) {
	var cache *LinearCache
	resourceType := resource.EndpointType

	reqs := make([]*discovery.DiscoveryRequest, 4)
	subs := make([]stream.Subscription, 4)
	watches := make([]chan Response, 4)

	buildRequest := func(res []string, version string) *discovery.DiscoveryRequest {
		return &discovery.DiscoveryRequest{
			ResourceNames: res,
			TypeUrl:       resourceType,
			VersionInfo:   version,
		}
	}
	updateReqResources := func(index int, res []string) {
		t.Helper()
		reqs[index-1].ResourceNames = res
		subs[index-1].SetResourceSubscription(reqs[index-1].ResourceNames)
	}

	createWatchWithCancel := func(index int) func() {
		t.Helper()
		w := make(chan Response, 1)
		cancel, err := cache.CreateWatch(reqs[index-1], subs[index-1], w)
		require.NoError(t, err)
		watches[index-1] = w
		return cancel
	}
	createWatch := func(index int) {
		t.Helper()
		_ = createWatchWithCancel(index)
	}
	validateResponse := func(index int, res []string) {
		t.Helper()
		resp := verifyResponseResources(t, watches[index-1], resourceType, cache.getVersion(), res...)
		updateFromSotwResponse(resp, &subs[index-1], reqs[index-1])
	}
	checkPendingWatch := func(index int) {
		t.Helper()
		mustBlock(t, watches[index-1])
	}

	// We run twice the same sequence of events,
	// First run is for endpoints, currently not a full-state resource
	// Second run is for clusters, currently a full-state resource
	t.Run("type not returning full state", func(t *testing.T) {
		resourceType := resource.EndpointType
		buildEndpoint := func(name string) *endpoint.ClusterLoadAssignment {
			return &endpoint.ClusterLoadAssignment{ClusterName: name}
		}

		cache = NewLinearCache(resourceType, WithLogger(log.NewTestLogger(t)), WithInitialResources(
			map[string]types.Resource{
				"a": buildEndpoint("a"),
				"b": buildEndpoint("b"),
				"c": buildEndpoint("c"),
			},
		))

		// Create watches
		// Watch 1, wildcard, starting with the current cache version
		reqs[0] = buildRequest(nil, cache.getVersion())
		subs[0] = subFromRequest(reqs[0])
		// Watch 2, wildcard, starting with no version (https://github.com/envoyproxy/go-control-plane/issues/855)
		reqs[1] = buildRequest([]string{"*"}, "")
		subs[1] = subFromRequest(reqs[1])
		// Watch 3, non-wildcard, starting with a different cache prefix
		reqs[2] = buildRequest([]string{"a", "b"}, "prefix-"+cache.getVersion())
		subs[2] = subFromRequest(reqs[2])
		// Watch 4, non-wildcard, starting with no version
		reqs[3] = buildRequest([]string{"d"}, "")
		subs[3] = subFromRequest(reqs[3])

		// Create watches
		// Version is ignored as we cannot guarantee the state, so everything is returned
		createWatch(1)
		validateResponse(1, []string{"a", "b", "c"})
		// Standard first wilcard request
		createWatch(2)
		validateResponse(2, []string{"a", "b", "c"})
		// Version has a different prefix and we send everything requested
		createWatch(3)
		validateResponse(3, []string{"a", "b"})
		// No requested version is available, so we return an empty response on first request
		createWatch(4)
		validateResponse(4, []string{})

		// Recreate watches
		createWatch(1)
		checkPendingWatch(1)
		createWatch(2)
		checkPendingWatch(2)
		createWatch(3)
		checkPendingWatch(3)
		createWatch(4)
		checkPendingWatch(4)

		// Update the cache
		_ = cache.UpdateResources(map[string]types.Resource{
			"b": buildEndpoint("b"),
			"d": buildEndpoint("d"),
		}, []string{"a"})

		validateResponse(1, []string{"b", "d"})
		validateResponse(2, []string{"b", "d"})
		validateResponse(3, []string{"b"})
		validateResponse(4, []string{"d"})

		createWatch(1)
		checkPendingWatch(1)
		// Make watch 2 no longer wildcard
		updateReqResources(2, []string{"a", "c", "d"})
		c2 := createWatchWithCancel(2)
		checkPendingWatch(2)
		c3 := createWatchWithCancel(3)
		checkPendingWatch(3)
		// Add a resource to watch 4 (https://github.com/envoyproxy/go-control-plane/issues/608)
		updateReqResources(4, []string{"c", "d"})
		createWatch(4)
		validateResponse(4, []string{"c"}) // c is newly requested, and should be returned
		createWatch(4)
		checkPendingWatch(4)

		// Add a new resource not requested in all subscriptions
		_ = cache.UpdateResource("e", buildEndpoint("e"))
		validateResponse(1, []string{"e"})
		createWatch(1)
		checkPendingWatch(1)
		checkPendingWatch(2) // No longer wildcard
		checkPendingWatch(3)
		checkPendingWatch(4)

		// Cancel two watches to change resources
		assert.Len(t, cache.resourceWatches["c"].sotw, 2)
		c2()
		assert.Len(t, cache.resourceWatches["c"].sotw, 1)
		assert.Len(t, cache.resourceWatches["b"].sotw, 1)
		c3()
		assert.Empty(t, cache.resourceWatches["b"].sotw)

		// Remove a resource from 2 (was a, c, d)
		updateReqResources(2, []string{"a", "d"})
		createWatch(2)
		checkPendingWatch(2)

		// 3 is now wildcard (was a, b). The version still matches the previous one
		updateReqResources(3, []string{"*"})
		createWatch(3)
		validateResponse(3, []string{"c", "d", "e"})
		createWatch(3)
		checkPendingWatch(3)

		// Do an update removing a resource only
		// This type is not full update, and therefore does not return
		_ = cache.UpdateResources(nil, []string{"c"})
		checkPendingWatch(1)
		checkPendingWatch(2)
		checkPendingWatch(3)
		checkPendingWatch(4)

		// Do an update in the cache to confirm all is well
		_ = cache.UpdateResources(map[string]types.Resource{
			"a": buildEndpoint("a"),
			"b": buildEndpoint("b"),
		}, nil)
		validateResponse(1, []string{"a", "b"})
		validateResponse(2, []string{"a"})
		validateResponse(3, []string{"a", "b"})
		checkPendingWatch(4)
	})

	t.Run("type returning full state", func(t *testing.T) {
		resourceType = resource.ClusterType
		buildCluster := func(name string) *cluster.Cluster {
			return &cluster.Cluster{Name: name}
		}

		cache = NewLinearCache(resource.ClusterType, WithLogger(log.NewTestLogger(t)), WithInitialResources(
			map[string]types.Resource{
				"a": buildCluster("a"),
				"b": buildCluster("b"),
				"c": buildCluster("c"),
			},
		))

		// Create watches
		// Watch 1, wildcard, starting with the current cache version
		reqs[0] = buildRequest(nil, cache.getVersion())
		subs[0] = subFromRequest(reqs[0])
		// Watch 2, wildcard, starting with no version (https://github.com/envoyproxy/go-control-plane/issues/855)
		reqs[1] = buildRequest([]string{"*"}, "")
		subs[1] = subFromRequest(reqs[1])
		// Watch 3, non-wildcard, starting with a different cache prefix
		reqs[2] = buildRequest([]string{"a", "b"}, "prefix-"+cache.getVersion())
		subs[2] = subFromRequest(reqs[2])
		// Watch 4, non-wildcard, starting with no version
		reqs[3] = buildRequest([]string{"d"}, "")
		subs[3] = subFromRequest(reqs[3])

		// Create watches
		// Version is ignored as we cannot guarantee the state, so everything is returned
		createWatch(1)
		validateResponse(1, []string{"a", "b", "c"})
		// Standard first wilcard request
		createWatch(2)
		validateResponse(2, []string{"a", "b", "c"})
		// Version has a different prefix and we send everything requested
		createWatch(3)
		validateResponse(3, []string{"a", "b"})
		// No requested version is available, so we return an empty response on first request
		createWatch(4)
		validateResponse(4, []string{})

		// Recreate watches
		createWatch(1)
		checkPendingWatch(1)
		createWatch(2)
		checkPendingWatch(2)
		createWatch(3)
		checkPendingWatch(3)
		createWatch(4)
		checkPendingWatch(4)

		// Update the cache
		_ = cache.UpdateResources(map[string]types.Resource{
			"b": buildCluster("b"),
			"d": buildCluster("d"),
		}, []string{"a"})

		validateResponse(1, []string{"b", "c", "d"})
		validateResponse(2, []string{"b", "c", "d"})
		validateResponse(3, []string{"b"})
		validateResponse(4, []string{"d"})

		createWatch(1)
		checkPendingWatch(1)
		// Make watch 2 no longer wildcard
		updateReqResources(2, []string{"a", "c", "d"})
		c2 := createWatchWithCancel(2)
		checkPendingWatch(2)
		c3 := createWatchWithCancel(3)
		checkPendingWatch(3)
		// Add a resource to req4 (https://github.com/envoyproxy/go-control-plane/issues/608)
		updateReqResources(4, []string{"c", "d"})
		createWatch(4)
		validateResponse(4, []string{"c", "d"}) // c is newly requested, and should be returned
		createWatch(4)
		checkPendingWatch(4)

		// Add a new resource not request in all subscriptions
		_ = cache.UpdateResource("e", buildCluster("e"))
		validateResponse(1, []string{"b", "c", "d", "e"})
		createWatch(1)
		checkPendingWatch(1)
		checkPendingWatch(2) // No longer wildcard
		checkPendingWatch(3)
		checkPendingWatch(4)

		// Cancel two watches to change resources
		assert.Len(t, cache.resourceWatches["c"].sotw, 2)
		c2()
		assert.Len(t, cache.resourceWatches["c"].sotw, 1)
		assert.Len(t, cache.resourceWatches["b"].sotw, 1)
		c3()
		assert.Empty(t, cache.resourceWatches["b"].sotw)

		// Remove a resource from 2 (was a, c, d)
		updateReqResources(2, []string{"a", "d"})
		createWatch(2)
		checkPendingWatch(2)

		// 3 is now wildcard (was a, b). The version still matches the previous one
		updateReqResources(3, []string{"*"})
		createWatch(3)
		validateResponse(3, []string{"b", "c", "d", "e"})
		createWatch(3)
		checkPendingWatch(3)

		// Do an update removing a resource only
		// This type is not full update, and therefore does not return
		_ = cache.UpdateResources(nil, []string{"c"})
		validateResponse(1, []string{"b", "d", "e"})
		checkPendingWatch(2)
		validateResponse(3, []string{"b", "d", "e"})
		validateResponse(4, []string{"d"})

		createWatch(1)
		checkPendingWatch(1)
		createWatch(3)
		checkPendingWatch(3)
		createWatch(4)
		checkPendingWatch(4)

		// Do an update in the cache to confirm all is well
		_ = cache.UpdateResources(map[string]types.Resource{
			"a": buildCluster("a"),
			"b": buildCluster("b"),
		}, nil)
		validateResponse(1, []string{"a", "b", "d", "e"})
		validateResponse(2, []string{"a", "d"})
		validateResponse(3, []string{"a", "b", "d", "e"})
		checkPendingWatch(4)
	})
}
