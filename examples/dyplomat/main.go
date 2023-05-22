package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"

	"google.golang.org/grpc"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	clusterservice "github.com/envoyproxy/go-control-plane/envoy/service/cluster/v3"
	discoverygrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	endpointservice "github.com/envoyproxy/go-control-plane/envoy/service/endpoint/v3"
	listenerservice "github.com/envoyproxy/go-control-plane/envoy/service/listener/v3"
	routeservice "github.com/envoyproxy/go-control-plane/envoy/service/route/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	xds "github.com/envoyproxy/go-control-plane/pkg/server/v3"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	k8scache "k8s.io/client-go/tools/cache"
)

type EnvoyCluster struct {
	name      string
	port      uint32
	endpoints []string
}

var (
	endpoints         []types.Resource
	version           int
	snapshotCache     cache.SnapshotCache
	endpointInformers []k8scache.SharedIndexInformer
)

func main() {
	version = 0
	snapshotCache = cache.NewSnapshotCache(false, cache.IDHash{}, nil)
	server := xds.NewServer(context.Background(), snapshotCache, nil)
	grpcServer := grpc.NewServer()
	lis, _ := net.Listen("tcp", ":8080")

	discoverygrpc.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	endpointservice.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
	clusterservice.RegisterClusterDiscoveryServiceServer(grpcServer, server)
	routeservice.RegisterRouteDiscoveryServiceServer(grpcServer, server)
	listenerservice.RegisterListenerDiscoveryServiceServer(grpcServer, server)

	clusters, _ := CreateBootstrapClients()

	for _, cluster := range clusters {
		stop := make(chan struct{})
		defer close(stop)

		factory := informers.NewSharedInformerFactoryWithOptions(cluster, time.Second*10, informers.WithNamespace("demo"))
		informer := factory.Core().V1().Endpoints().Informer()
		endpointInformers = append(endpointInformers, informer)

		informer.AddEventHandler(k8scache.ResourceEventHandlerFuncs{
			UpdateFunc: HandleEndpointsUpdate,
		})

		go func() {
			informer.Run(stop)
		}()
	}

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("%v", err)
	}
}

func HandleEndpointsUpdate(oldObj, newObj interface{}) {

	edsServiceData := map[string]*EnvoyCluster{}

	for _, inform := range endpointInformers {
		for _, ep := range inform.GetStore().List() {

			endpoints := ep.(*corev1.Endpoints)
			if _, ok := endpoints.Labels["xds"]; !ok {
				continue
			}

			if _, ok := edsServiceData[endpoints.Name]; !ok {
				edsServiceData[endpoints.Name] = &EnvoyCluster{
					name: endpoints.Name,
				}
			}

			for _, subset := range endpoints.Subsets {
				for i, addr := range subset.Addresses {
					edsServiceData[endpoints.Name].port = uint32(subset.Ports[i].Port)
					edsServiceData[endpoints.Name].endpoints = append(edsServiceData[endpoints.Name].endpoints, addr.IP)
				}
			}
		}
	}

	// for each service create endpoints
	edsEndpoints := make([]types.Resource, len(edsServiceData))
	for _, envoyCluster := range edsServiceData {
		edsEndpoints = append(edsEndpoints, MakeEndpointsForCluster(envoyCluster))
	}

	snapshot, err := cache.NewSnapshot(fmt.Sprintf("%v.0", version), map[resource.Type][]types.Resource{
		resource.EndpointType: edsEndpoints,
	})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	IDs := snapshotCache.GetStatusKeys()
	for _, id := range IDs {
		err = snapshotCache.SetSnapshot(context.Background(), id, snapshot)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}

	version++
}

func MakeEndpointsForCluster(service *EnvoyCluster) *endpointv3.ClusterLoadAssignment {
	fmt.Printf("Updating endpoints for cluster %s: %v\n", service.name, service.endpoints)
	cla := &endpointv3.ClusterLoadAssignment{
		ClusterName: service.name,
		Endpoints:   []*endpointv3.LocalityLbEndpoints{},
	}

	for _, endpoint := range service.endpoints {
		cla.Endpoints = append(cla.Endpoints,
			&endpointv3.LocalityLbEndpoints{
				LbEndpoints: []*endpointv3.LbEndpoint{{
					HostIdentifier: &endpointv3.LbEndpoint_Endpoint{
						Endpoint: &endpointv3.Endpoint{
							Address: &core.Address{
								Address: &core.Address_SocketAddress{
									SocketAddress: &core.SocketAddress{
										Protocol: core.SocketAddress_TCP,
										Address:  endpoint,
										PortSpecifier: &core.SocketAddress_PortValue{
											PortValue: service.port,
										},
									},
								},
							},
						},
					},
				}},
			},
		)
	}
	return cla
}
