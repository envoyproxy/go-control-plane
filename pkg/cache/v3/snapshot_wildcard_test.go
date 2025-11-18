// Copyright 2025 Envoyproxy Authors
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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource/v3"
)

func TestNewSnapshotWithExplicitWildcard(t *testing.T) {
	cluster1 := resource.MakeCluster(resource.Ads, "cluster1")
	cluster2 := resource.MakeCluster(resource.Ads, "cluster2")
	cluster3 := resource.MakeCluster(resource.Ads, "cluster3")

	snapshot, err := cache.NewSnapshotWithExplicitWildcard("v1", map[rsrc.Type][]cache.SnapshotResource{
		rsrc.ClusterType: {
			{Resource: types.ResourceWithTTL{Resource: cluster1}, Wildcard: true},
			{Resource: types.ResourceWithTTL{Resource: cluster2}, Wildcard: false},
			{Resource: types.ResourceWithTTL{Resource: cluster3}, Wildcard: true},
		},
	})

	require.NoError(t, err)
	require.NotNil(t, snapshot)

	allResources := snapshot.GetResources(rsrc.ClusterType)
	assert.Len(t, allResources, 3)
	assert.Contains(t, allResources, "cluster1")
	assert.Contains(t, allResources, "cluster2")
	assert.Contains(t, allResources, "cluster3")

	wildcardResources := snapshot.GetWildcardResources(rsrc.ClusterType)
	assert.Len(t, wildcardResources, 2)
	assert.Contains(t, wildcardResources, "cluster1")
	assert.Contains(t, wildcardResources, "cluster3")
	assert.NotContains(t, wildcardResources, "cluster2")

	assert.True(t, snapshot.IsResourceWildcard(rsrc.ClusterType, "cluster1"))
	assert.False(t, snapshot.IsResourceWildcard(rsrc.ClusterType, "cluster2"))
	assert.True(t, snapshot.IsResourceWildcard(rsrc.ClusterType, "cluster3"))
}

func TestSnapshotWithExplicitWildcard_AllNonWildcard(t *testing.T) {
	cluster1 := resource.MakeCluster(resource.Ads, "cluster1")
	cluster2 := resource.MakeCluster(resource.Ads, "cluster2")

	snapshot, err := cache.NewSnapshotWithExplicitWildcard("v1", map[rsrc.Type][]cache.SnapshotResource{
		rsrc.ClusterType: {
			{Resource: types.ResourceWithTTL{Resource: cluster1}, Wildcard: false},
			{Resource: types.ResourceWithTTL{Resource: cluster2}, Wildcard: false},
		},
	})

	require.NoError(t, err)
	require.NotNil(t, snapshot)

	allResources := snapshot.GetResources(rsrc.ClusterType)
	assert.Len(t, allResources, 2)

	wildcardResources := snapshot.GetWildcardResources(rsrc.ClusterType)
	assert.Empty(t, wildcardResources)

	assert.False(t, snapshot.IsResourceWildcard(rsrc.ClusterType, "cluster1"))
	assert.False(t, snapshot.IsResourceWildcard(rsrc.ClusterType, "cluster2"))
}

func TestSnapshotWithExplicitWildcard_AllWildcard(t *testing.T) {
	cluster1 := resource.MakeCluster(resource.Ads, "cluster1")
	cluster2 := resource.MakeCluster(resource.Ads, "cluster2")

	snapshot, err := cache.NewSnapshotWithExplicitWildcard("v1", map[rsrc.Type][]cache.SnapshotResource{
		rsrc.ClusterType: {
			{Resource: types.ResourceWithTTL{Resource: cluster1}, Wildcard: true},
			{Resource: types.ResourceWithTTL{Resource: cluster2}, Wildcard: true},
		},
	})

	require.NoError(t, err)
	require.NotNil(t, snapshot)

	allResources := snapshot.GetResources(rsrc.ClusterType)
	assert.Len(t, allResources, 2)

	wildcardResources := snapshot.GetWildcardResources(rsrc.ClusterType)
	assert.Len(t, wildcardResources, 2)
	assert.Contains(t, wildcardResources, "cluster1")
	assert.Contains(t, wildcardResources, "cluster2")

	assert.True(t, snapshot.IsResourceWildcard(rsrc.ClusterType, "cluster1"))
	assert.True(t, snapshot.IsResourceWildcard(rsrc.ClusterType, "cluster2"))
}

func TestBackwardCompatibility_RegularSnapshot(t *testing.T) {
	cluster1 := resource.MakeCluster(resource.Ads, "cluster1")
	cluster2 := resource.MakeCluster(resource.Ads, "cluster2")

	snapshot, err := cache.NewSnapshot("v1", map[rsrc.Type][]types.Resource{
		rsrc.ClusterType: {cluster1, cluster2},
	})

	require.NoError(t, err)
	require.NotNil(t, snapshot)

	allResources := snapshot.GetResources(rsrc.ClusterType)
	assert.Len(t, allResources, 2)

	wildcardResources := snapshot.GetWildcardResources(rsrc.ClusterType)
	assert.Len(t, wildcardResources, 2)
	assert.Contains(t, wildcardResources, "cluster1")
	assert.Contains(t, wildcardResources, "cluster2")

	assert.True(t, snapshot.IsResourceWildcard(rsrc.ClusterType, "cluster1"))
	assert.True(t, snapshot.IsResourceWildcard(rsrc.ClusterType, "cluster2"))
}

func TestWildcardResourcesWithTTL(t *testing.T) {
	cluster1 := resource.MakeCluster(resource.Ads, "cluster1")
	cluster2 := resource.MakeCluster(resource.Ads, "cluster2")

	snapshot, err := cache.NewSnapshotWithExplicitWildcard("v1", map[rsrc.Type][]cache.SnapshotResource{
		rsrc.ClusterType: {
			{Resource: types.ResourceWithTTL{Resource: cluster1}, Wildcard: true},
			{Resource: types.ResourceWithTTL{Resource: cluster2}, Wildcard: false},
		},
	})

	require.NoError(t, err)
	require.NotNil(t, snapshot)

	wildcardResourcesWithTTL := snapshot.GetWildcardResourcesAndTTL(rsrc.ClusterType)
	assert.Len(t, wildcardResourcesWithTTL, 1)
	assert.Contains(t, wildcardResourcesWithTTL, "cluster1")
	assert.NotContains(t, wildcardResourcesWithTTL, "cluster2")

	res, ok := wildcardResourcesWithTTL["cluster1"]
	assert.True(t, ok)
	assert.NotNil(t, res.Resource)
}

func TestMixedResourceTypes(t *testing.T) {
	cluster1 := resource.MakeCluster(resource.Ads, "cluster1")
	cluster2 := resource.MakeCluster(resource.Ads, "cluster2")
	endpoint1 := resource.MakeEndpoint("cluster1", 8080)
	endpoint2 := resource.MakeEndpoint("cluster2", 8080)

	snapshot, err := cache.NewSnapshotWithExplicitWildcard("v1", map[rsrc.Type][]cache.SnapshotResource{
		rsrc.ClusterType: {
			{Resource: types.ResourceWithTTL{Resource: cluster1}, Wildcard: true},
			{Resource: types.ResourceWithTTL{Resource: cluster2}, Wildcard: false},
		},
		rsrc.EndpointType: {
			{Resource: types.ResourceWithTTL{Resource: endpoint1}, Wildcard: false},
			{Resource: types.ResourceWithTTL{Resource: endpoint2}, Wildcard: true},
		},
	})

	require.NoError(t, err)
	require.NotNil(t, snapshot)

	wildcardClusters := snapshot.GetWildcardResources(rsrc.ClusterType)
	assert.Len(t, wildcardClusters, 1)
	assert.Contains(t, wildcardClusters, "cluster1")

	wildcardEndpoints := snapshot.GetWildcardResources(rsrc.EndpointType)
	assert.Len(t, wildcardEndpoints, 1)
	assert.Contains(t, wildcardEndpoints, "cluster2")
}

func TestEmptySnapshot(t *testing.T) {
	snapshot, err := cache.NewSnapshotWithExplicitWildcard("v1", map[rsrc.Type][]cache.SnapshotResource{})

	require.NoError(t, err)
	require.NotNil(t, snapshot)

	// Empty resources should return empty maps
	allResources := snapshot.GetResources(rsrc.ClusterType)
	assert.Nil(t, allResources)

	wildcardResources := snapshot.GetWildcardResources(rsrc.ClusterType)
	assert.Empty(t, wildcardResources)

	// Non-existent resources should be non-wildcard
	assert.False(t, snapshot.IsResourceWildcard(rsrc.ClusterType, "nonexistent"))
}

func TestUnknownResourceType(t *testing.T) {
	cluster1 := &cluster.Cluster{Name: "cluster1"}

	snapshot, err := cache.NewSnapshotWithExplicitWildcard("v1", map[rsrc.Type][]cache.SnapshotResource{
		"unknown.type": {
			{Resource: types.ResourceWithTTL{Resource: cluster1}, Wildcard: true},
		},
	})

	assert.Error(t, err)
	assert.Nil(t, snapshot)
	assert.Contains(t, err.Error(), "unknown resource type")
}

func TestNilSnapshot(t *testing.T) {
	var snapshot *cache.Snapshot

	assert.False(t, snapshot.IsResourceWildcard(rsrc.ClusterType, "cluster1"))
	assert.Nil(t, snapshot.GetWildcardResources(rsrc.ClusterType))
	assert.Nil(t, snapshot.GetWildcardResourcesAndTTL(rsrc.ClusterType))
}

func TestResourceNotInWildcardMap(t *testing.T) {
	cluster1 := resource.MakeCluster(resource.Ads, "cluster1")

	snapshot, err := cache.NewSnapshotWithExplicitWildcard("v1", map[rsrc.Type][]cache.SnapshotResource{
		rsrc.ClusterType: {
			{Resource: types.ResourceWithTTL{Resource: cluster1}, Wildcard: true},
		},
	})

	require.NoError(t, err)
	require.NotNil(t, snapshot)

	// Querying endpoint resources should return empty (not panic)
	wildcardEndpoints := snapshot.GetWildcardResources(rsrc.EndpointType)
	assert.Empty(t, wildcardEndpoints)

	assert.False(t, snapshot.IsResourceWildcard(rsrc.EndpointType, "some-endpoint"))
}
