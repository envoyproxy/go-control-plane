package cache

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// wcStream drives one stream against a cache the way the servers do: on every response it hands the
// returned versions back to the subscription and re-arms a watch with a non-initial request.
type wcStream struct {
	delta bool
	sub   stream.Subscription
	req   *Request
	dreq  *DeltaRequest
	sotwC chan Response
	delC  chan DeltaResponse
}

func newWCStream(c *LinearCache, delta bool, subscribed []string, initial map[string]string, version string) *wcStream {
	s := &wcStream{delta: delta}
	if delta {
		s.sub = stream.NewDeltaSubscription(subscribed, nil, initial, true)
		s.dreq = &DeltaRequest{TypeUrl: c.typeURL}
		if initial != nil || len(subscribed) > 0 {
			s.dreq.ResponseNonce = "reconnect"
		}
		s.delC = make(chan DeltaResponse, 1)
		_, err := c.CreateDeltaWatch(s.dreq, s.sub, s.delC)
		if err != nil {
			panic(err)
		}
		return s
	}
	s.sub = stream.NewSotwSubscription(subscribed, true)
	s.req = &Request{TypeUrl: c.typeURL, ResourceNames: subscribed, VersionInfo: version}
	s.sotwC = make(chan Response, 1)
	_, err := c.CreateWatch(s.req, s.sub, s.sotwC)
	if err != nil {
		panic(err)
	}
	return s
}

// pump delivers all the responses currently available and returns a comparable description of each.
func (s *wcStream) pump(c *LinearCache) ([]string, error) {
	var out []string
	for {
		if s.delta {
			select {
			case resp := <-s.delC:
				dr, err := resp.GetDeltaDiscoveryResponse()
				if err != nil {
					return nil, err
				}
				var names []string
				for _, r := range dr.GetResources() {
					names = append(names, r.GetName()+"@"+r.GetVersion())
				}
				removed := append([]string(nil), dr.GetRemovedResources()...)
				sort.Strings(names)
				sort.Strings(removed)
				out = append(out, fmt.Sprint(names, removed, dr.GetSystemVersionInfo(), resp.GetReturnedResources()))
				s.sub.SetReturnedResources(resp.GetReturnedResources())
				s.dreq = &DeltaRequest{TypeUrl: c.typeURL, ResponseNonce: "n"}
				s.delC = make(chan DeltaResponse, 1)
				if _, err := c.CreateDeltaWatch(s.dreq, s.sub, s.delC); err != nil {
					return nil, err
				}
			default:
				return out, nil
			}
			continue
		}
		select {
		case resp := <-s.sotwC:
			dr, err := resp.GetDiscoveryResponse()
			if err != nil {
				return nil, err
			}
			var values []string
			for _, r := range dr.GetResources() {
				values = append(values, string(r.GetValue()))
			}
			sort.Strings(values)
			out = append(out, fmt.Sprint(values, dr.GetVersionInfo(), resp.GetReturnedResources()))
			updateFromSotwResponse(resp, &s.sub, s.req)
			s.req = &Request{TypeUrl: c.typeURL, ResourceNames: s.req.GetResourceNames(), VersionInfo: s.req.GetVersionInfo()}
			s.sotwC = make(chan Response, 1)
			if _, err := c.CreateWatch(s.req, s.sub, s.sotwC); err != nil {
				return nil, err
			}
		default:
			return out, nil
		}
	}
}

// clearSynced makes the cache treat every wildcard watch as not in sync, so the next update compares
// the whole cache for it. This is the behavior the incremental path has to match.
func clearSynced(c *LinearCache) {
	for id := range c.syncedWildcardWatches {
		delete(c.syncedWildcardWatches, id)
	}
}

type wcKind struct {
	typeURL string
	mk      func(name string, rev int) types.Resource
}

func wcKinds() map[string]wcKind {
	return map[string]wcKind{
		// sotw replies only with what changed for this type
		"endpoint": {resource.EndpointType, func(name string, rev int) types.Resource {
			return &endpoint.ClusterLoadAssignment{ClusterName: name, Endpoints: []*endpoint.LocalityLbEndpoints{{Priority: uint32(rev)}}}
		}},
		// sotw has to reply with every resource for this type
		"cluster": {resource.ClusterType, func(name string, rev int) types.Resource {
			return &cluster.Cluster{Name: name, ConnectTimeout: nil, AltStatName: strconv.Itoa(rev)}
		}},
	}
}

// TestLinearWildcardIncrementalMatchesFullScan applies the same random operations to a cache using the
// incremental wildcard path and to one comparing the whole cache for every watch, with the same
// streams on both, and requires identical responses.
func TestLinearWildcardIncrementalMatchesFullScan(t *testing.T) {
	for name, kind := range wcKinds() {
		t.Run(name, func(t *testing.T) {
			for seed := int64(1); seed <= 10; seed++ {
				runWildcardDifferential(t, kind, seed)
			}
		})
	}
}

func runWildcardDifferential(t *testing.T, kind wcKind, seed int64) {
	t.Helper()
	rnd := rand.New(rand.NewSource(seed))
	const pool = 100
	fast, ref := NewLinearCache(kind.typeURL), NewLinearCache(kind.typeURL)
	names := make([]string, pool)
	for i := range names {
		names[i] = fmt.Sprintf("r%03d", i)
	}
	pick := func(n int) []string {
		out := make([]string, n)
		for i := range out {
			out[i] = names[rnd.Intn(pool)]
		}
		return out
	}

	var fs, rs []*wcStream
	connect := func(delta bool, subscribed []string, stale bool) {
		var initial map[string]string
		var version string
		if stale {
			// a reconnecting client with partly outdated state, unknown resources, and an old version
			initial = map[string]string{"ghost": "x"}
			for _, n := range pick(30) {
				initial[n] = "old"
			}
			version = "old"
		}
		clone := func(m map[string]string) map[string]string {
			if m == nil {
				return nil
			}
			c := make(map[string]string, len(m))
			for k, v := range m {
				c[k] = v
			}
			return c
		}
		fs = append(fs, newWCStream(fast, delta, subscribed, clone(initial), version))
		rs = append(rs, newWCStream(ref, delta, subscribed, clone(initial), version))
	}
	compare := func(step string) {
		for i := range fs {
			got, err := fs[i].pump(fast)
			require.NoError(t, err)
			want, err := rs[i].pump(ref)
			require.NoError(t, err)
			require.Equal(t, want, got, "seed %d %s: stream %d responses differ", seed, step, i)
		}
	}

	seedRes := map[string]types.Resource{}
	for _, n := range pick(30) {
		seedRes[n] = kind.mk(n, 0)
	}
	for _, c := range []*LinearCache{fast, ref} {
		require.NoError(t, c.UpdateResources(seedRes, nil))
	}
	clearSynced(ref)
	connect(true, nil, false)
	connect(false, nil, false)
	connect(true, pick(8), false)
	connect(false, pick(8), false)
	compare("connect")

	for step := 0; step < 300; step++ {
		// The reference never has synced watches, so it compares the whole cache on every update.
		clearSynced(ref)
		switch op := rnd.Intn(10); {
		case op <= 3:
			upd := map[string]types.Resource{}
			for _, n := range pick(1 + rnd.Intn(5)) {
				upd[n] = kind.mk(n, rnd.Intn(3))
			}
			// deleting a name that is also updated lists it twice in the modified names
			del := append(pick(rnd.Intn(3)), pick(1)[0])
			for _, c := range []*LinearCache{fast, ref} {
				require.NoError(t, c.UpdateResources(upd, del))
			}
		case op == 4:
			n := names[rnd.Intn(pool)]
			rev := rnd.Intn(3)
			for _, c := range []*LinearCache{fast, ref} {
				require.NoError(t, c.UpdateResource(n, kind.mk(n, rev)))
			}
		case op == 5:
			n := names[rnd.Intn(pool)]
			for _, c := range []*LinearCache{fast, ref} {
				require.NoError(t, c.DeleteResource(n))
			}
		case op == 6:
			set := map[string]types.Resource{}
			for _, n := range pick(5 + rnd.Intn(30)) {
				set[n] = kind.mk(n, rnd.Intn(3))
			}
			for _, c := range []*LinearCache{fast, ref} {
				cp := make(map[string]types.Resource, len(set))
				for k, v := range set {
					cp[k] = v
				}
				require.NoError(t, c.SetResources(cp))
			}
		case op == 7:
			connect(rnd.Intn(2) == 0, nil, rnd.Intn(2) == 0)
		case op == 8:
			connect(rnd.Intn(2) == 0, pick(6), rnd.Intn(2) == 0)
		default:
			connect(true, nil, false)
		}
		compare(fmt.Sprintf("step %d", step))
	}

	// the shortcut is only valid while a synced watch really holds the cache's versions
	for id := range fast.syncedWildcardWatches {
		w, ok := fast.wildcardWatches[id]
		require.True(t, ok, "seed %d: synced watch %d is not registered", seed, id)
		sub := w.getSubscription()
		known := sub.ReturnedResources()
		assert.Len(t, known, len(fast.resources), "seed %d: synced watch %d diverged", seed, id)
		for name, res := range fast.resources {
			v, err := res.getVersion(w.useResourceVersion())
			require.NoError(t, err)
			assert.Equal(t, v, known[name], "seed %d: synced watch %d diverged on %s", seed, id, name)
		}
	}
	assert.NotEmpty(t, fast.syncedWildcardWatches, "seed %d: the incremental path was not exercised", seed)
}

func TestLinearWildcardIncrementalUnchangedAndStale(t *testing.T) {
	c := NewLinearCache(resource.EndpointType)
	mk := func(name string, rev int) types.Resource {
		return &endpoint.ClusterLoadAssignment{ClusterName: name, Endpoints: []*endpoint.LocalityLbEndpoints{{Priority: uint32(rev)}}}
	}
	require.NoError(t, c.UpdateResources(map[string]types.Resource{"a": mk("a", 0), "b": mk("b", 0)}, nil))

	// a delta wildcard stream is answered first, then registered as synced
	s := newWCStream(c, true, nil, nil, "")
	got, err := s.pump(c)
	require.NoError(t, err)
	require.Len(t, got, 1)
	assert.Len(t, c.syncedWildcardWatches, 1)

	// rewriting identical content does not answer a delta stream, and it stays synced
	require.NoError(t, c.UpdateResources(map[string]types.Resource{"a": mk("a", 0)}, nil))
	got, err = s.pump(c)
	require.NoError(t, err)
	assert.Empty(t, got)
	assert.Len(t, c.syncedWildcardWatches, 1)

	// only the changed resource is returned, and removal is reported
	require.NoError(t, c.UpdateResources(map[string]types.Resource{"b": mk("b", 1)}, []string{"a"}))
	got, err = s.pump(c)
	require.NoError(t, err)
	require.Len(t, got, 1)
	assert.Contains(t, got[0], "b@")
	assert.Contains(t, got[0], "[a]")

	// a stream reconnecting with stale versions is not synced until it has been answered
	stale := newWCStream(c, true, nil, map[string]string{"b": "old", "gone": "x"}, "")
	assert.Len(t, c.syncedWildcardWatches, 1, "a stale stream must not be tracked as synced")
	got, err = stale.pump(c)
	require.NoError(t, err)
	require.Len(t, got, 1)
	assert.Len(t, c.syncedWildcardWatches, 2)
}

func TestLinearWildcardIncrementalCancelUntracks(t *testing.T) {
	c := NewLinearCache(resource.EndpointType)
	sub := stream.NewDeltaSubscription(nil, nil, map[string]string{}, true)
	w := make(chan DeltaResponse, 1)
	cancel, err := c.CreateDeltaWatch(&DeltaRequest{TypeUrl: resource.EndpointType, ResponseNonce: "n"}, sub, w)
	require.NoError(t, err)
	assert.Len(t, c.syncedWildcardWatches, 1)
	cancel()
	assert.Empty(t, c.syncedWildcardWatches)
	assert.Zero(t, c.NumWildcardWatches())
}

// BenchmarkLinearWildcardUpdate changes one resource in a large cache that has several wildcard delta
// watches open. write-ns/op is the UpdateResources call alone, ns/op includes the watches being re-armed.
func BenchmarkLinearWildcardUpdate(b *testing.B) {
	const resources, watches = 100000, 8
	for _, tc := range []struct {
		name string
		full bool
	}{{"full-scan", true}, {"incremental", false}} {
		b.Run(tc.name, func(b *testing.B) {
			c := NewLinearCache(resource.EndpointType)
			mk := func(i, rev int) types.Resource {
				return &endpoint.ClusterLoadAssignment{ClusterName: strconv.Itoa(i), Endpoints: []*endpoint.LocalityLbEndpoints{{Priority: uint32(rev)}}}
			}
			initial := make(map[string]types.Resource, resources)
			for i := 0; i < resources; i++ {
				initial[strconv.Itoa(i)] = mk(i, 0)
			}
			require.NoError(b, c.UpdateResources(initial, nil))
			ss := make([]*wcStream, watches)
			for i := range ss {
				ss[i] = newWCStream(c, true, nil, nil, "")
				_, err := ss[i].pump(c)
				require.NoError(b, err)
			}
			b.ResetTimer()
			var write time.Duration
			for i := 0; i < b.N; i++ {
				if tc.full {
					clearSynced(c)
				}
				start := time.Now()
				require.NoError(b, c.UpdateResources(map[string]types.Resource{strconv.Itoa(i % resources): mk(i%resources, i+1)}, nil))
				write += time.Since(start)
				for _, s := range ss {
					_, err := s.pump(c)
					require.NoError(b, err)
				}
			}
			b.ReportMetric(float64(write.Nanoseconds())/float64(b.N), "write-ns/op")
		})
	}
}
