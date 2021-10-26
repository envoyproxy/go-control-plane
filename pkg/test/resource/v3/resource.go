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

// Package resource creates test xDS resources
package resource

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"

	alf "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3"
	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	als "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/grpc/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	tcp "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"
	auth "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	runtime "github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
)

const (
	localhost = "127.0.0.1"

	// XdsCluster is the cluster name for the control server (used by non-ADS set-up).
	XdsCluster = "xds_cluster"

	// AlsCluster is the clustername for gRPC access log service (ALS)
	AlsCluster = "als_cluster"

	// Ads mode for resources: one aggregated xDS service
	Ads = "ads"

	// Xds mode for resources: individual xDS services.
	Xds = "xds"

	// Rest mode for resources: polling using Fetch.
	Rest = "rest"

	// Delta mode for resources: individual delta xDS services.
	Delta = "delta"

	// Delta Ads mode for resource: one aggregated delta xDS service.
	DeltaAds = "delta-ads"
)

var (
	// RefreshDelay for the polling config source.
	RefreshDelay = 500 * time.Millisecond
)

// MakeEndpoint creates a localhost endpoint on a given port.
func MakeEndpoint(clusterName string, port uint32) *endpoint.ClusterLoadAssignment {
	return &endpoint.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{{
			LbEndpoints: []*endpoint.LbEndpoint{{
				HostIdentifier: &endpoint.LbEndpoint_Endpoint{
					Endpoint: &endpoint.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Protocol: core.SocketAddress_TCP,
									Address:  localhost,
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: port,
									},
								},
							},
						},
					},
				},
			}},
		}},
	}
}

// MakeCluster creates a cluster using either ADS or EDS.
func MakeCluster(mode string, clusterName string) *cluster.Cluster {
	edsSource := configSource(mode)

	connectTimeout := 5 * time.Second
	return &cluster.Cluster{
		Name:                 clusterName,
		ConnectTimeout:       durationpb.New(connectTimeout),
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_EDS},
		EdsClusterConfig: &cluster.Cluster_EdsClusterConfig{
			EdsConfig: edsSource,
		},
	}
}

// MakeRoute creates an HTTP route that routes to a given cluster.
func MakeRoute(routeName, clusterName string) *route.RouteConfiguration {
	return &route.RouteConfiguration{
		Name: routeName,
		VirtualHosts: []*route.VirtualHost{{
			Name:    routeName,
			Domains: []string{"*"},
			Routes: []*route.Route{{
				Match: &route.RouteMatch{
					PathSpecifier: &route.RouteMatch_Prefix{
						Prefix: "/",
					},
				},
				Action: &route.Route_Route{
					Route: &route.RouteAction{
						ClusterSpecifier: &route.RouteAction_Cluster{
							Cluster: clusterName,
						},
					},
				},
			}},
		}},
	}
}

// MakeScopedRoute creates an HTTP scoped route that routes to a given cluster.
func MakeScopedRoute(scopedRouteName string, routeConfigurationName string, keyFragments []string) *route.ScopedRouteConfiguration {
	k := &route.ScopedRouteConfiguration_Key{}

	for _, key := range keyFragments {
		fragment := &route.ScopedRouteConfiguration_Key_Fragment{
			Type: &route.ScopedRouteConfiguration_Key_Fragment_StringKey{
				StringKey: key,
			},
		}
		k.Fragments = append(k.Fragments, fragment)
	}

	return &route.ScopedRouteConfiguration{
		OnDemand:               false,
		Name:                   scopedRouteName,
		RouteConfigurationName: routeConfigurationName,
		Key:                    k,
	}
}

// data source configuration
func configSource(mode string) *core.ConfigSource {
	source := &core.ConfigSource{}
	source.ResourceApiVersion = resource.DefaultAPIVersion
	switch mode {
	case Ads:
		source.ConfigSourceSpecifier = &core.ConfigSource_Ads{
			Ads: &core.AggregatedConfigSource{},
		}
	case DeltaAds:
		source.ConfigSourceSpecifier = &core.ConfigSource_Ads{
			Ads: &core.AggregatedConfigSource{},
		}
	case Xds:
		source.ConfigSourceSpecifier = &core.ConfigSource_ApiConfigSource{
			ApiConfigSource: &core.ApiConfigSource{
				TransportApiVersion:       resource.DefaultAPIVersion,
				ApiType:                   core.ApiConfigSource_GRPC,
				SetNodeOnFirstMessageOnly: true,
				GrpcServices: []*core.GrpcService{{
					TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
						EnvoyGrpc: &core.GrpcService_EnvoyGrpc{ClusterName: XdsCluster},
					},
				}},
			},
		}
	case Rest:
		source.ConfigSourceSpecifier = &core.ConfigSource_ApiConfigSource{
			ApiConfigSource: &core.ApiConfigSource{
				ApiType:             core.ApiConfigSource_REST,
				TransportApiVersion: resource.DefaultAPIVersion,
				ClusterNames:        []string{XdsCluster},
				RefreshDelay:        durationpb.New(RefreshDelay),
			},
		}
	case Delta:
		source.ConfigSourceSpecifier = &core.ConfigSource_ApiConfigSource{
			ApiConfigSource: &core.ApiConfigSource{
				TransportApiVersion:       resource.DefaultAPIVersion,
				ApiType:                   core.ApiConfigSource_DELTA_GRPC,
				SetNodeOnFirstMessageOnly: true,
				GrpcServices: []*core.GrpcService{{
					TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
						EnvoyGrpc: &core.GrpcService_EnvoyGrpc{ClusterName: XdsCluster},
					},
				}},
			},
		}
	}
	return source
}

func buildHTTPConnectionManager() *hcm.HttpConnectionManager {
	// access log service configuration.
	alsConfig := &als.HttpGrpcAccessLogConfig{
		CommonConfig: &als.CommonGrpcAccessLogConfig{
			LogName:             "echo",
			TransportApiVersion: resource.DefaultAPIVersion,
			GrpcService: &core.GrpcService{
				TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &core.GrpcService_EnvoyGrpc{
						ClusterName: AlsCluster,
					},
				},
			},
		},
	}
	alsConfigPbst, err := anypb.New(alsConfig)
	if err != nil {
		panic(err)
	}

	// HTTP filter configuration.
	manager := &hcm.HttpConnectionManager{
		CodecType:  hcm.HttpConnectionManager_AUTO,
		StatPrefix: "http",
		HttpFilters: []*hcm.HttpFilter{{
			Name: wellknown.Router,
		}},
		AccessLog: []*alf.AccessLog{{
			Name: wellknown.HTTPGRPCAccessLog,
			ConfigType: &alf.AccessLog_TypedConfig{
				TypedConfig: alsConfigPbst,
			},
		}},
	}

	return manager
}

func makeListener(listenerName string, port uint32, filterChains []*listener.FilterChain) *listener.Listener {
	return &listener.Listener{
		Name: listenerName,
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.SocketAddress_TCP,
					Address:  localhost,
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: port,
					},
				},
			},
		},
		FilterChains: filterChains,
	}
}

func MakeRouteHTTPListener(mode string, listenerName string, port uint32, route string) *listener.Listener {
	rdsSource := configSource(mode)
	routeSpecifier := &hcm.HttpConnectionManager_Rds{
		Rds: &hcm.Rds{
			ConfigSource:    rdsSource,
			RouteConfigName: route,
		},
	}

	manager := buildHTTPConnectionManager()
	manager.RouteSpecifier = routeSpecifier

	pbst, err := anypb.New(manager)
	if err != nil {
		panic(err)
	}

	filterChains := []*listener.FilterChain{
		{
			Filters: []*listener.Filter{
				{
					Name: wellknown.HTTPConnectionManager,
					ConfigType: &listener.Filter_TypedConfig{
						TypedConfig: pbst,
					},
				},
			},
		},
	}

	return makeListener(listenerName, port, filterChains)
}

// Creates a HTTP listener using Scoped Routes, which extracts the "Host" header field as the key.
func MakeScopedRouteHTTPListener(mode string, listenerName string, port uint32, scopedRouteConfigName string) *listener.Listener {
	source := configSource(mode)
	routeSpecifier := &hcm.HttpConnectionManager_ScopedRoutes{
		ScopedRoutes: &hcm.ScopedRoutes{
			Name: scopedRouteConfigName,
			ScopeKeyBuilder: &hcm.ScopedRoutes_ScopeKeyBuilder{
				Fragments: []*hcm.ScopedRoutes_ScopeKeyBuilder_FragmentBuilder{
					{
						Type: &hcm.ScopedRoutes_ScopeKeyBuilder_FragmentBuilder_HeaderValueExtractor_{
							HeaderValueExtractor: &hcm.ScopedRoutes_ScopeKeyBuilder_FragmentBuilder_HeaderValueExtractor{
								Name: "Host",
								ExtractType: &hcm.ScopedRoutes_ScopeKeyBuilder_FragmentBuilder_HeaderValueExtractor_Index{
									Index: 0,
								},
							},
						},
					},
				},
			},
			RdsConfigSource: source,
			ConfigSpecifier: &hcm.ScopedRoutes_ScopedRds{
				ScopedRds: &hcm.ScopedRds{
					ScopedRdsConfigSource: source,
				},
			},
		},
	}

	manager := buildHTTPConnectionManager()
	manager.RouteSpecifier = routeSpecifier

	pbst, err := anypb.New(manager)
	if err != nil {
		panic(err)
	}

	filterChains := []*listener.FilterChain{
		{
			Filters: []*listener.Filter{
				{
					Name: wellknown.HTTPConnectionManager,
					ConfigType: &listener.Filter_TypedConfig{
						TypedConfig: pbst,
					},
				},
			},
		},
	}

	return makeListener(listenerName, port, filterChains)
}

// Creates a TCP listener HTTP manager.
func MakeTCPListener(listenerName string, port uint32, clusterName string) *listener.Listener {
	// TCP filter configuration
	config := &tcp.TcpProxy{
		StatPrefix: "tcp",
		ClusterSpecifier: &tcp.TcpProxy_Cluster{
			Cluster: clusterName,
		},
	}
	pbst, err := anypb.New(config)
	if err != nil {
		panic(err)
	}

	filterChains := []*listener.FilterChain{
		{
			Filters: []*listener.Filter{
				{
					Name: wellknown.TCPProxy,
					ConfigType: &listener.Filter_TypedConfig{
						TypedConfig: pbst,
					},
				},
			},
		},
	}

	return makeListener(listenerName, port, filterChains)
}

// MakeRuntime creates an RTDS layer with some fields.
func MakeRuntime(runtimeName string) *runtime.Runtime {
	return &runtime.Runtime{
		Name: runtimeName,
		Layer: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"field-0": {
					Kind: &structpb.Value_NumberValue{NumberValue: 100},
				},
				"field-1": {
					Kind: &structpb.Value_StringValue{StringValue: "foobar"},
				},
			},
		},
	}
}

// MakeExtensionConfig creates a extension config for a cluster.
func MakeExtensionConfig(mode string, extensionConfigName string, route string) *core.TypedExtensionConfig {
	rdsSource := configSource(mode)

	// HTTP filter configuration
	manager := &hcm.HttpConnectionManager{
		CodecType:  hcm.HttpConnectionManager_AUTO,
		StatPrefix: "http",
		RouteSpecifier: &hcm.HttpConnectionManager_Rds{
			Rds: &hcm.Rds{
				ConfigSource:    rdsSource,
				RouteConfigName: route,
			},
		},
		HttpFilters: []*hcm.HttpFilter{{
			Name: wellknown.Router,
		}},
	}
	pbst, err := anypb.New(manager)
	if err != nil {
		panic(err)
	}

	return &core.TypedExtensionConfig{
		Name:        extensionConfigName,
		TypedConfig: pbst,
	}
}

// TestSnapshot holds parameters for a synthetic snapshot.
type TestSnapshot struct {
	// Xds indicates snapshot mode: ads, xds, rest, or delta
	Xds string
	// Version for the snapshot.
	Version string
	// UpstreamPort for the single endpoint on the localhost.
	UpstreamPort uint32
	// BasePort is the initial port for the listeners.
	BasePort uint32
	// NumClusters is the total number of clusters to generate.
	NumClusters int
	// NumHTTPListeners is the total number of HTTP listeners to generate.
	NumHTTPListeners int
	// NumScopedHTTPListeners is the total number of scoped route HTTP listeners to generate.
	NumScopedHTTPListeners int
	// NumTCPListeners is the total number of TCP listeners to generate.
	// Listeners are assigned clusters in a round-robin fashion.
	NumTCPListeners int
	// NumRuntimes is the total number of RTDS layers to generate.
	NumRuntimes int
	// TLS enables SDS-enabled TLS mode on all listeners
	TLS bool
	// NumExtension is the total number of Extension Config
	NumExtension int
}

// Generate produces a snapshot from the parameters.
func (ts TestSnapshot) Generate() cache.Snapshot {
	clusters := make([]types.Resource, ts.NumClusters)
	endpoints := make([]types.Resource, ts.NumClusters)
	for i := 0; i < ts.NumClusters; i++ {
		name := fmt.Sprintf("cluster-%s-%d", ts.Version, i)
		clusters[i] = MakeCluster(ts.Xds, name)
		endpoints[i] = MakeEndpoint(name, ts.UpstreamPort)
	}

	totalHTTPListeners := ts.NumHTTPListeners + ts.NumScopedHTTPListeners
	routes := make([]types.Resource, totalHTTPListeners)
	scopedRoutes := make([]types.Resource, ts.NumScopedHTTPListeners)

	for i := 0; i < totalHTTPListeners; i++ {
		suffix := fmt.Sprintf("%s-%d", ts.Version, i)
		routeName := fmt.Sprintf("route-%s", suffix)
		routes[i] = MakeRoute(routeName, cache.GetResourceName(clusters[i%ts.NumClusters]))

		// Scoped Routes.
		if i >= ts.NumHTTPListeners {
			scopedRouteName := fmt.Sprintf("scopedroute-%s", suffix)
			port := ts.BasePort + uint32(i)
			scopedRoutes[i-ts.NumHTTPListeners] = MakeScopedRoute(scopedRouteName, routeName, []string{fmt.Sprintf("127.0.0.1:%d", port)})
		}
	}

	numHTTPListeners := ts.NumHTTPListeners
	numScopedHTTPListeners := ts.NumScopedHTTPListeners
	numTCPListeners := ts.NumTCPListeners
	total := numHTTPListeners + numScopedHTTPListeners + numTCPListeners

	listeners := make([]types.Resource, total)
	for i := 0; i < total; i++ {
		port := ts.BasePort + uint32(i)
		// listener name must be same since ports are shared and previous listener is drained
		name := fmt.Sprintf("listener-%d", port)
		var listener *listener.Listener

		if numHTTPListeners > 0 {
			listener = MakeRouteHTTPListener(ts.Xds, name, port, cache.GetResourceName(routes[i]))
			numHTTPListeners--
		} else if numScopedHTTPListeners > 0 {
			listener = MakeScopedRouteHTTPListener(ts.Xds, name, port, cache.GetResourceName(scopedRoutes[numScopedHTTPListeners-1]))
			numScopedHTTPListeners--
		} else if numTCPListeners > 0 {
			listener = MakeTCPListener(name, port, cache.GetResourceName(clusters[i%ts.NumClusters]))
			numTCPListeners--
		}

		if ts.TLS {
			for i, chain := range listener.FilterChains {
				tlsc := &auth.DownstreamTlsContext{
					CommonTlsContext: &auth.CommonTlsContext{
						TlsCertificateSdsSecretConfigs: []*auth.SdsSecretConfig{{
							Name:      tlsName,
							SdsConfig: configSource(ts.Xds),
						}},
						ValidationContextType: &auth.CommonTlsContext_ValidationContextSdsSecretConfig{
							ValidationContextSdsSecretConfig: &auth.SdsSecretConfig{
								Name:      rootName,
								SdsConfig: configSource(ts.Xds),
							},
						},
					},
				}
				mt, _ := anypb.New(tlsc)
				chain.TransportSocket = &core.TransportSocket{
					Name: "envoy.transport_sockets.tls",
					ConfigType: &core.TransportSocket_TypedConfig{
						TypedConfig: mt,
					},
				}
				listener.FilterChains[i] = chain
			}
		}

		listeners[i] = listener
	}

	runtimes := make([]types.Resource, ts.NumRuntimes)
	for i := 0; i < ts.NumRuntimes; i++ {
		name := fmt.Sprintf("runtime-%d", i)
		runtimes[i] = MakeRuntime(name)
	}

	var secrets []types.Resource
	if ts.TLS {
		for _, s := range MakeSecrets(tlsName, rootName) {
			secrets = append(secrets, s)
		}
	}

	extensions := make([]types.Resource, ts.NumExtension)
	for i := 0; i < ts.NumExtension; i++ {
		routeName := fmt.Sprintf("route-%s-%d", ts.Version, i)
		extensionConfigName := fmt.Sprintf("extensionConfig-%d", i)
		extensions[i] = MakeExtensionConfig(Ads, extensionConfigName, routeName)
	}

	out, _ := cache.NewSnapshot(ts.Version, map[resource.Type][]types.Resource{
		resource.EndpointType:        endpoints,
		resource.ClusterType:         clusters,
		resource.RouteType:           routes,
		resource.ScopedRouteType:     scopedRoutes,
		resource.ListenerType:        listeners,
		resource.RuntimeType:         runtimes,
		resource.SecretType:          secrets,
		resource.ExtensionConfigType: extensions,
	})

	return out
}
