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

	"github.com/stretchr/testify/require"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
)

const (
	testType = "google.protobuf.StringValue"
)

func testResource(s string) types.Resource {
	return &wrappers.StringValue{Value: s}
}

func verifyResponse(t *testing.T, ch <-chan Response, version string, num int) {
	t.Helper()
	r := <-ch
	if r.GetRequest().TypeUrl != testType {
		t.Errorf("unexpected empty request type URL: %q", r.GetRequest().TypeUrl)
	}
	out, err := r.GetDiscoveryResponse()
	if err != nil {
		t.Fatal(err)
	}
	if out.VersionInfo == "" {
		t.Error("unexpected response empty version")
	}
	if n := len(out.Resources); n != num {
		t.Errorf("unexpected number of responses: got %d, want %d", n, num)
	}
	if version != "" && out.VersionInfo != version {
		t.Errorf("unexpected version: got %q, want %q", out.VersionInfo, version)
	}
	if out.TypeUrl != testType {
		t.Errorf("unexpected type URL: %q", out.TypeUrl)
	}
}

func checkWatchCount(t *testing.T, c *LinearCache, name string, count int) {
	t.Helper()
	if i := c.NumWatches(name); i != count {
		t.Errorf("unexpected number of watches for %q: got %d, want %d", name, i, count)
	}
}

func mustBlock(t *testing.T, w <-chan Response) {
	select {
	case <-w:
		t.Error("watch must block")
	default:
	}
}

func TestLinearInitialResources(t *testing.T) {
	c := NewLinearCache(testType, WithInitialResources(map[string]types.Resource{"a": testResource("a"), "b": testResource("b")}))
	w := make(chan Response, 1)
	c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType}, w)
	verifyResponse(t, w, "0", 1)
	c.CreateWatch(&Request{TypeUrl: testType}, w)
	verifyResponse(t, w, "0", 2)
}

func TestLinearCornerCases(t *testing.T) {
	c := NewLinearCache(testType)
	err := c.UpdateResource("a", nil)
	if err == nil {
		t.Error("expected error on nil resource")
	}
	// create an incorrect type URL request
	w := make(chan Response, 1)
	c.CreateWatch(&Request{TypeUrl: "test"}, w)
	select {
	case r := <-w:
		if r != nil {
			t.Error("response should be nil")
		}
	default:
		t.Error("should receive nil response")
	}
}

func TestLinearBasic(t *testing.T) {
	c := NewLinearCache(testType)

	// Create watches before a resource is ready
	w1 := make(chan Response, 1)
	c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}, w1)
	mustBlock(t, w1)

	w := make(chan Response, 1)
	c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "0"}, w)
	mustBlock(t, w)
	checkWatchCount(t, c, "a", 2)
	checkWatchCount(t, c, "b", 1)
	require.NoError(t, c.UpdateResource("a", testResource("a")))
	checkWatchCount(t, c, "a", 0)
	checkWatchCount(t, c, "b", 0)
	verifyResponse(t, w1, "1", 1)
	verifyResponse(t, w, "1", 1)

	// Request again, should get same response
	c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}, w)
	checkWatchCount(t, c, "a", 0)
	verifyResponse(t, w, "1", 1)
	c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "0"}, w)
	checkWatchCount(t, c, "a", 0)
	verifyResponse(t, w, "1", 1)

	// Add another element and update the first, response should be different
	require.NoError(t, c.UpdateResource("b", testResource("b")))
	require.NoError(t, c.UpdateResource("a", testResource("aa")))
	c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}, w)
	verifyResponse(t, w, "3", 1)
	c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "0"}, w)
	verifyResponse(t, w, "3", 2)
}

func TestLinearSetResources(t *testing.T) {
	c := NewLinearCache(testType)

	// Create new resources
	w1 := make(chan Response, 1)
	c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}, w1)
	mustBlock(t, w1)
	w2 := make(chan Response, 1)
	c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "0"}, w2)
	mustBlock(t, w2)
	c.SetResources(map[string]types.Resource{
		"a": testResource("a"),
		"b": testResource("b"),
	})
	verifyResponse(t, w1, "1", 1)
	verifyResponse(t, w2, "1", 2) // the version was only incremented once for all resources

	// Add another element and update the first, response should be different
	c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "1"}, w1)
	mustBlock(t, w1)
	c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "1"}, w2)
	mustBlock(t, w2)
	c.SetResources(map[string]types.Resource{
		"a": testResource("aa"),
		"b": testResource("b"),
		"c": testResource("c"),
	})
	verifyResponse(t, w1, "2", 1)
	verifyResponse(t, w2, "2", 3)

	// Delete resource
	c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "2"}, w1)
	mustBlock(t, w1)
	c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "2"}, w2)
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
	c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}, w)
	verifyResponse(t, w, "instance1-0", 0)

	require.NoError(t, c.UpdateResource("a", testResource("a")))
	c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}, w)
	verifyResponse(t, w, "instance1-1", 1)

	c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "instance1-1"}, w)
	mustBlock(t, w)
	checkWatchCount(t, c, "a", 1)
}

func TestLinearDeletion(t *testing.T) {
	c := NewLinearCache(testType, WithInitialResources(map[string]types.Resource{"a": testResource("a"), "b": testResource("b")}))
	w := make(chan Response, 1)
	c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "0"}, w)
	mustBlock(t, w)
	checkWatchCount(t, c, "a", 1)
	require.NoError(t, c.DeleteResource("a"))
	verifyResponse(t, w, "1", 0)
	checkWatchCount(t, c, "a", 0)
	c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "0"}, w)
	verifyResponse(t, w, "1", 1)
	checkWatchCount(t, c, "b", 0)
	require.NoError(t, c.DeleteResource("b"))
	c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "1"}, w)
	verifyResponse(t, w, "2", 0)
	checkWatchCount(t, c, "b", 0)
}

func TestLinearWatchTwo(t *testing.T) {
	c := NewLinearCache(testType, WithInitialResources(map[string]types.Resource{"a": testResource("a"), "b": testResource("b")}))
	w := make(chan Response, 1)
	c.CreateWatch(&Request{ResourceNames: []string{"a", "b"}, TypeUrl: testType, VersionInfo: "0"}, w)
	mustBlock(t, w)
	w1 := make(chan Response, 1)
	c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "0"}, w1)
	mustBlock(t, w1)
	require.NoError(t, c.UpdateResource("a", testResource("aa")))
	// should only get the modified resource
	verifyResponse(t, w, "1", 1)
	verifyResponse(t, w1, "1", 2)
}

func TestLinearCancel(t *testing.T) {
	c := NewLinearCache(testType)
	require.NoError(t, c.UpdateResource("a", testResource("a")))

	// cancel watch-all
	w := make(chan Response, 1)
	cancel := c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "1"}, w)
	mustBlock(t, w)
	checkWatchCount(t, c, "a", 1)
	cancel()
	checkWatchCount(t, c, "a", 0)

	// cancel watch for "a"
	cancel = c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "1"}, w)
	mustBlock(t, w)
	checkWatchCount(t, c, "a", 1)
	cancel()
	checkWatchCount(t, c, "a", 0)

	// open four watches for "a" and "b" and two for all, cancel one of each, make sure the second one is unaffected
	w2 := make(chan Response, 1)
	w3 := make(chan Response, 1)
	w4 := make(chan Response, 1)
	cancel = c.CreateWatch(&Request{ResourceNames: []string{"a"}, TypeUrl: testType, VersionInfo: "1"}, w)
	cancel2 := c.CreateWatch(&Request{ResourceNames: []string{"b"}, TypeUrl: testType, VersionInfo: "1"}, w2)
	cancel3 := c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "1"}, w3)
	cancel4 := c.CreateWatch(&Request{TypeUrl: testType, VersionInfo: "1"}, w4)
	mustBlock(t, w)
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
					c.CreateWatch(&Request{
						// Only expect one to become stale
						ResourceNames: []string{id, id2},
						VersionInfo:   "0",
						TypeUrl:       testType,
					}, value)
					// wait until all updates apply
					verifyResponse(t, value, "", 1)
				}
			})
		}(i)
	}
}
