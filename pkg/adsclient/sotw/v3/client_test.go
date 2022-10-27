package sotw_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	client "github.com/envoyproxy/go-control-plane/pkg/adsclient/sotw/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	snapCache := cache.NewSnapshotCache(true, cache.IDHash{}, nil)
	go func() {
		err := startAdsServer(t, ctx, snapCache)
		assert.NoError(t, err)
	}()

	conn, err := grpc.Dial(":18001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	assert.NoError(t, err)
	defer conn.Close()

	c := client.NewAdsClient(ctx, "node_1", resource.ClusterType)
	c.InitConnect(conn)

	t.Run("Test initial fetch", testInitialFetch(t, ctx, snapCache, c))
	t.Run("Test next fetch", testNextFetch(t, ctx, snapCache, c))
}

func testInitialFetch(t *testing.T, ctx context.Context, snapCache cache.SnapshotCache, c client.AdsClient) func(t *testing.T) {
	return func(t *testing.T) {
		go func() {
			// watch for configs
			resp, err := c.Fetch()
			assert.NoError(t, err)
			assert.Equal(t, 3, len(resp.Resources))
			for i, r := range resp.Resources {
				cluster := &clusterv3.Cluster{}
				anypb.UnmarshalTo(r, cluster, proto.UnmarshalOptions{})
				assert.Equal(t, fmt.Sprint("cluster_", i), cluster.Name)
			}

			err = c.Ack()
			assert.NoError(t, err)
		}()

		snapshot, err := cache.NewSnapshot("1", map[resource.Type][]types.Resource{
			resource.ClusterType: {
				&clusterv3.Cluster{Name: "cluster_1"},
				&clusterv3.Cluster{Name: "cluster_2"},
				&clusterv3.Cluster{Name: "cluster_3"},
			},
		})
		assert.NoError(t, err)

		err = snapshot.Consistent()
		assert.NoError(t, err)
		snapCache.SetSnapshot(ctx, "node_1", snapshot)
	}
}

func testNextFetch(t *testing.T, ctx context.Context, snapCache cache.SnapshotCache, c client.AdsClient) func(t *testing.T) {
	return func(t *testing.T) {
		go func() {
			// watch for configs
			resp, err := c.Fetch()
			assert.NoError(t, err)
			assert.Equal(t, 2, len(resp.Resources))
			for i, r := range resp.Resources {
				cluster := &clusterv3.Cluster{}
				anypb.UnmarshalTo(r, cluster, proto.UnmarshalOptions{})
				assert.Equal(t, fmt.Sprint("cluster_", i), cluster.Name)
			}

			err = c.Ack()
			assert.NoError(t, err)
		}()

		snapshot, err := cache.NewSnapshot("2", map[resource.Type][]types.Resource{
			resource.ClusterType: {
				&clusterv3.Cluster{Name: "cluster_1"},
				&clusterv3.Cluster{Name: "cluster_2"},
			},
		})
		assert.NoError(t, err)

		err = snapshot.Consistent()
		assert.NoError(t, err)
		snapCache.SetSnapshot(ctx, "node_1", snapshot)
	}
}

func startAdsServer(t *testing.T, ctx context.Context, snapCache cache.SnapshotCache) error {
	lis, err := net.Listen("tcp", ":18001")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	s := server.NewServer(ctx, snapCache, nil)
	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, s)

	if e := grpcServer.Serve(lis); e != nil {
		err = e
	}

	return err
}
