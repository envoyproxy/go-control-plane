package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource/v3"
)

func TestSnapshotWithWildcardFilter(t *testing.T) {
	// Test that OnDemandOnly flag is correctly stored and retrieved
	wildcardCluster := resource.MakeCluster(resource.Ads, "wildcard-cluster")
	onDemandCluster := resource.MakeCluster(resource.Ads, "on-demand-cluster")

	snapshot, err := types.NewSnapshot("v1", map[string][]types.SnapshotResource{
		rsrc.ClusterType: {
			{Name: "wildcard-cluster", Resource: wildcardCluster, OnDemandOnly: false},
			{Name: "on-demand-cluster", Resource: onDemandCluster, OnDemandOnly: true},
		},
	})
	require.NoError(t, err)

	typeSnapshot := snapshot.GetTypeSnapshot(rsrc.ClusterType)
	resources := typeSnapshot.GetResources()

	require.Len(t, resources, 2)
	assert.False(t, resources["wildcard-cluster"].IsOnDemandOnly())
	assert.True(t, resources["on-demand-cluster"].IsOnDemandOnly())
}

func TestSnapshotResourceDefaultWildcard(t *testing.T) {
	// Test that SnapshotResource defaults to wildcard-eligible (OnDemandOnly=false)
	cluster := resource.MakeCluster(resource.Ads, "test-cluster")

	snapshot, err := types.NewSnapshot("v1", map[string][]types.SnapshotResource{
		rsrc.ClusterType: {
			{Name: "test-cluster", Resource: cluster},
		},
	})
	require.NoError(t, err)

	typeSnapshot := snapshot.GetTypeSnapshot(rsrc.ClusterType)
	resources := typeSnapshot.GetResources()

	require.Len(t, resources, 1)
	assert.False(t, resources["test-cluster"].IsOnDemandOnly())
}
