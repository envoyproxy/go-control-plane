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
