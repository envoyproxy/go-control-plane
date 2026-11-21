package cache

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/log"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

func newClusterSnapshot(t *testing.T, version string, clusters ...types.Resource) *Snapshot {
	t.Helper()
	snap, err := NewSnapshot(version, map[resource.Type][]types.Resource{resource.ClusterType: clusters})
	require.NoError(t, err)
	return snap
}

func TestSetSnapshotInheritsVersionsOfSameObjects(t *testing.T) {
	c := NewSnapshotCache(false, IDHash{}, log.NewTestLogger(t), WithReuseUnchangedResources())
	const node = "node"

	kept := &cluster.Cluster{Name: "kept"}
	replaced := &cluster.Cluster{Name: "replaced"}
	changed := &cluster.Cluster{Name: "changed"}

	first := newClusterSnapshot(t, "1", kept, replaced, changed)
	require.NoError(t, c.SetSnapshot(context.Background(), node, first))
	require.NoError(t, first.ConstructVersionMap())
	firstVersions := first.GetVersionMap(resource.ClusterType)

	// Modifying kept in place breaks the option's contract; it is done here only to
	// prove the next snapshot inherits kept's version instead of hashing it again.
	kept.AltStatName = "mutated"
	replacedCopy := &cluster.Cluster{Name: "replaced"}
	changedCopy := &cluster.Cluster{Name: "changed", AltStatName: "new"}

	second := newClusterSnapshot(t, "2", kept, replacedCopy, changedCopy)
	require.NoError(t, c.SetSnapshot(context.Background(), node, second))
	require.NoError(t, second.ConstructVersionMap())
	secondVersions := second.GetVersionMap(resource.ClusterType)

	assert.Equal(t, firstVersions["kept"], secondVersions["kept"])
	assert.Nil(t, second.getMarshaledResource(resource.ClusterType, "kept"))

	// A new object is hashed again, so equal content keeps an equal version and new
	// content gets a new one.
	assert.Equal(t, firstVersions["replaced"], secondVersions["replaced"])
	assert.NotNil(t, second.getMarshaledResource(resource.ClusterType, "replaced"))
	assert.NotEqual(t, firstVersions["changed"], secondVersions["changed"])
	assert.Nil(t, second.inheritedVersions)
}

func TestSetSnapshotWithoutReuseHashesEveryResource(t *testing.T) {
	c := NewSnapshotCache(false, IDHash{}, log.NewTestLogger(t))
	const node = "node"

	res := &cluster.Cluster{Name: "a"}
	first := newClusterSnapshot(t, "1", res)
	require.NoError(t, c.SetSnapshot(context.Background(), node, first))
	require.NoError(t, first.ConstructVersionMap())
	assert.Nil(t, first.getMarshaledResource(resource.ClusterType, "a"))

	res.AltStatName = "mutated"
	second := newClusterSnapshot(t, "2", res)
	require.NoError(t, c.SetSnapshot(context.Background(), node, second))
	assert.Nil(t, second.inheritedVersions)
	require.NoError(t, second.ConstructVersionMap())
	assert.NotEqual(t, first.GetVersionMap(resource.ClusterType)["a"], second.GetVersionMap(resource.ClusterType)["a"])
}

func TestSetSnapshotWithoutVersionMapInheritsNothing(t *testing.T) {
	c := NewSnapshotCache(false, IDHash{}, log.NewTestLogger(t), WithReuseUnchangedResources())
	const node = "node"

	res := &cluster.Cluster{Name: "a"}
	require.NoError(t, c.SetSnapshot(context.Background(), node, newClusterSnapshot(t, "1", res)))

	second := newClusterSnapshot(t, "2", res)
	require.NoError(t, c.SetSnapshot(context.Background(), node, second))
	assert.Nil(t, second.inheritedVersions)
}

func TestDeltaResponseWithReuseSendsHashedBytes(t *testing.T) {
	c := NewSnapshotCache(false, IDHash{}, log.NewTestLogger(t), WithReuseUnchangedResources())
	const node = "node"

	kept := &cluster.Cluster{Name: "kept"}
	first := newClusterSnapshot(t, "1", kept)
	require.NoError(t, c.SetSnapshot(context.Background(), node, first))
	require.NoError(t, first.ConstructVersionMap())

	added := &cluster.Cluster{Name: "added"}
	second := newClusterSnapshot(t, "2", kept, added)
	require.NoError(t, c.SetSnapshot(context.Background(), node, second))

	// A wildcard watch with no known resources gets both: kept through lazy encoding,
	// added through the bytes kept while hashing.
	watch := make(chan DeltaResponse, 1)
	_, err := c.CreateDeltaWatch(&discovery.DeltaDiscoveryRequest{
		Node:    &core.Node{Id: node},
		TypeUrl: resource.ClusterType,
	}, stream.NewDeltaSubscription(nil, nil, nil, true), watch)
	require.NoError(t, err)

	var out DeltaResponse
	select {
	case out = <-watch:
	case <-time.After(time.Second):
		t.Fatal("failed to receive delta response")
	}
	require.NotNil(t, second.getMarshaledResource(resource.ClusterType, "added"))
	require.Nil(t, second.getMarshaledResource(resource.ClusterType, "kept"))

	resp, err := out.GetDeltaDiscoveryResponse()
	require.NoError(t, err)
	require.Len(t, resp.GetResources(), 2)
	versions := second.GetVersionMap(resource.ClusterType)
	for _, r := range resp.GetResources() {
		want, err := MarshalResource(map[string]types.Resource{"kept": kept, "added": added}[r.GetName()])
		require.NoError(t, err)
		assert.Equal(t, want, r.GetResource().GetValue(), r.GetName())
		assert.Equal(t, versions[r.GetName()], r.GetVersion(), r.GetName())
	}
}
