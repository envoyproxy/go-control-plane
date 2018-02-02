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

package resource

import (
	"time"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	bootstrap "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
)

// MakeBootstrap creates a bootstrap envoy configuration
func MakeBootstrap(ads bool, controlPort, adminPort uint32) *bootstrap.Bootstrap {
	source := &core.ApiConfigSource{
		ApiType:      core.ApiConfigSource_GRPC,
		ClusterNames: []string{XdsCluster},
	}

	var dynamic *bootstrap.Bootstrap_DynamicResources
	if ads {
		dynamic = &bootstrap.Bootstrap_DynamicResources{
			LdsConfig: &core.ConfigSource{
				ConfigSourceSpecifier: &core.ConfigSource_Ads{Ads: &core.AggregatedConfigSource{}},
			},
			CdsConfig: &core.ConfigSource{
				ConfigSourceSpecifier: &core.ConfigSource_Ads{Ads: &core.AggregatedConfigSource{}},
			},
			AdsConfig: source,
		}
	} else {
		dynamic = &bootstrap.Bootstrap_DynamicResources{
			LdsConfig: &core.ConfigSource{
				ConfigSourceSpecifier: &core.ConfigSource_ApiConfigSource{ApiConfigSource: source},
			},
			CdsConfig: &core.ConfigSource{
				ConfigSourceSpecifier: &core.ConfigSource_ApiConfigSource{ApiConfigSource: source},
			},
		}
	}

	return &bootstrap.Bootstrap{
		Node: &core.Node{
			Id:      "test-id",
			Cluster: "test-cluster",
		},
		Admin: bootstrap.Admin{
			AccessLogPath: "/dev/null",
			Address: core.Address{
				Address: &core.Address_SocketAddress{
					SocketAddress: &core.SocketAddress{
						Address: "127.0.0.1",
						PortSpecifier: &core.SocketAddress_PortValue{
							PortValue: adminPort,
						},
					},
				},
			},
		},
		StaticResources: &bootstrap.Bootstrap_StaticResources{
			Clusters: []v2.Cluster{{
				Name:           XdsCluster,
				ConnectTimeout: 5 * time.Second,
				Type:           v2.Cluster_STATIC,
				Hosts: []*core.Address{{
					Address: &core.Address_SocketAddress{
						SocketAddress: &core.SocketAddress{
							Address: "127.0.0.1",
							PortSpecifier: &core.SocketAddress_PortValue{
								PortValue: controlPort,
							},
						},
					},
				}},
				LbPolicy:             v2.Cluster_ROUND_ROBIN,
				Http2ProtocolOptions: &core.Http2ProtocolOptions{},
			}},
		},
		DynamicResources: dynamic,
	}
}
