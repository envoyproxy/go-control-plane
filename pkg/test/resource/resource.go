// Copyright 2017 Envoyproxy Authors
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

// Package resource creates test xDS resources
package resource

import (
	"github.com/envoyproxy/go-control-plane/api"
	"github.com/envoyproxy/go-control-plane/api/filter/http"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/golang/protobuf/ptypes/duration"
	structpb "github.com/golang/protobuf/ptypes/struct"
)

const (
	localhost  = "127.0.0.1"
	router     = "envoy.router"
	httpFilter = "envoy.http_connection_manager"
	xdsCluster = "xds_cluster"
)

// MakeEndpoint creates a localhost endpoint.
func MakeEndpoint(cluster string, port uint32) *api.ClusterLoadAssignment {
	return &api.ClusterLoadAssignment{
		ClusterName: cluster,
		Endpoints: []*api.LocalityLbEndpoints{{
			LbEndpoints: []*api.LbEndpoint{{
				Endpoint: &api.Endpoint{
					Address: &api.Address{
						Address: &api.Address_SocketAddress{
							SocketAddress: &api.SocketAddress{
								Protocol: api.SocketAddress_TCP,
								Address:  localhost,
								PortSpecifier: &api.SocketAddress_PortValue{
									PortValue: port,
								},
							},
						},
					},
				},
			}},
		}},
	}
}

// MakeCluster creates a cluster.
func MakeCluster(ads bool, cluster string) *api.Cluster {
	var edsSource *api.ConfigSource
	if ads {
		edsSource = &api.ConfigSource{
			ConfigSourceSpecifier: &api.ConfigSource_Ads{
				Ads: &api.AggregatedConfigSource{},
			},
		}
	} else {
		edsSource = &api.ConfigSource{
			ConfigSourceSpecifier: &api.ConfigSource_ApiConfigSource{
				ApiConfigSource: &api.ApiConfigSource{
					ApiType:     api.ApiConfigSource_GRPC,
					ClusterName: []string{xdsCluster},
				},
			},
		}
	}

	return &api.Cluster{
		Name:           cluster,
		ConnectTimeout: &duration.Duration{Seconds: 5},
		Type:           api.Cluster_EDS,
		EdsClusterConfig: &api.Cluster_EdsClusterConfig{
			EdsConfig:   edsSource,
			ServiceName: cluster,
		},
	}
}

// MakeRoute creates an HTTP route.
func MakeRoute(route, cluster string) *api.RouteConfiguration {
	return &api.RouteConfiguration{
		Name: route,
		VirtualHosts: []*api.VirtualHost{{
			Name:    "all",
			Domains: []string{"*"},
			Routes: []*api.Route{{
				Match: &api.RouteMatch{
					PathSpecifier: &api.RouteMatch_Prefix{
						Prefix: "/",
					},
				},
				Action: &api.Route_Route{
					Route: &api.RouteAction{
						ClusterSpecifier: &api.RouteAction_Cluster{
							Cluster: cluster,
						},
					},
				},
			}},
		}},
	}
}

// MakeListener creates a listener.
func MakeListener(ads bool, listener string, port uint32, route string) *api.Listener {
	var rdsSource *api.ConfigSource
	if ads {
		rdsSource = &api.ConfigSource{
			ConfigSourceSpecifier: &api.ConfigSource_Ads{
				Ads: &api.AggregatedConfigSource{},
			},
		}
	} else {
		rdsSource = &api.ConfigSource{
			ConfigSourceSpecifier: &api.ConfigSource_ApiConfigSource{
				ApiConfigSource: &api.ApiConfigSource{
					ApiType:     api.ApiConfigSource_GRPC,
					ClusterName: []string{xdsCluster},
				},
			},
		}
	}
	manager := &http.HttpConnectionManager{
		CodecType: http.HttpConnectionManager_AUTO,
		RouteSpecifier: &http.HttpConnectionManager_Rds{
			Rds: &http.Rds{
				ConfigSource:    rdsSource,
				RouteConfigName: route,
			},
		},
		HttpFilters: []*http.HttpFilter{{
			Name: router,
			// TODO(kuat) surprising requirement to allow LDS retrieve this route
			Config: &structpb.Struct{
				Fields: map[string]*structpb.Value{
					"deprecated_v1": {Kind: &structpb.Value_BoolValue{BoolValue: true}},
				},
			},
		}},
	}
	pbst, err := cache.MessageToStruct(manager)
	if err != nil {
		panic(err)
	}

	return &api.Listener{
		Name: listener,
		Address: &api.Address{
			Address: &api.Address_SocketAddress{
				SocketAddress: &api.SocketAddress{
					Protocol: api.SocketAddress_TCP,
					Address:  localhost,
					PortSpecifier: &api.SocketAddress_PortValue{
						PortValue: port,
					},
				},
			},
		},
		FilterChains: []*api.FilterChain{{
			Filters: []*api.Filter{{
				Name:   httpFilter,
				Config: pbst,
			}},
		}},
	}
}
