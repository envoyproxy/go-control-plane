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

package cache

import (
	"github.com/gogo/protobuf/proto"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	hcm "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/envoyproxy/go-control-plane/pkg/util"
)

// Resource is the base interface for the xDS payload.
type Resource interface {
	proto.Message
	Equal(interface{}) bool
}

// Common names for Envoy filters.
const (
	// Buffer HTTP filter
	Buffer = "envoy.buffer"
	// CORS HTTP filter
	CORS = "envoy.cors"
	// Dynamo HTTP filter
	Dynamo = "envoy.http_dynamo_filter"
	// Fault HTTP filter
	Fault = "envoy.fault"
	// GRPCHTTP1Bridge HTTP filter
	GRPCHTTP1Bridge = "envoy.grpc_http1_bridge"
	// GRPCJSONTranscoder HTTP filter
	GRPCJSONTranscoder = "envoy.grpc_json_transcoder"
	// GRPCWeb HTTP filter
	GRPCWeb = "envoy.grpc_web"
	// Gzip HTTP filter
	Gzip = "envoy.gzip"
	// IPTagging HTTP filter
	IPTagging = "envoy.ip_tagging"
	// HTTPRateLimit filter
	HTTPRateLimit = "envoy.rate_limit"
	// Router HTTP filter
	Router = "envoy.router"
	// Health checking HTTP filter
	HealthCheck = "envoy.health_check"
	// Lua HTTP filter
	Lua = "envoy.lua"
	// Squash HTTP filter
	Squash = "envoy.squash"
	// HTTPExternalAuthorization HTTP filter
	HTTPExternalAuthorization = "envoy.ext_authz"

	// Echo network filter
	Echo = "envoy.echo"
	// HTTPConnectionManager network filter
	HTTPConnectionManager = "envoy.http_connection_manager"
	// TCPProxy network filter
	TCPProxy = "envoy.tcp_proxy"
	// RateLimit network filter
	RateLimit = "envoy.ratelimit"
	// MongoProxy network filter
	MongoProxy = "envoy.mongo_proxy"
	// RedisProxy network filter
	RedisProxy = "envoy.redis_proxy"
	// ExternalAuthorization network filter
	ExternalAuthorization = "envoy.ext_authz"
)

// Resource types in xDS v2.
const (
	typePrefix   = "type.googleapis.com/envoy.api.v2."
	EndpointType = typePrefix + "ClusterLoadAssignment"
	ClusterType  = typePrefix + "Cluster"
	RouteType    = typePrefix + "RouteConfiguration"
	ListenerType = typePrefix + "Listener"

	// AnyType is used only by ADS
	AnyType = ""
)

var (
	// ResponseTypes are supported response types.
	ResponseTypes = []string{
		EndpointType,
		ClusterType,
		RouteType,
		ListenerType,
	}
)

// GetResourceName returns the resource name for a valid xDS response type.
func GetResourceName(res Resource) string {
	switch v := res.(type) {
	case *v2.ClusterLoadAssignment:
		return v.GetClusterName()
	case *v2.Cluster:
		return v.GetName()
	case *v2.RouteConfiguration:
		return v.GetName()
	case *v2.Listener:
		return v.GetName()
	default:
		return ""
	}
}

// GetResourceReferences returns the names for dependent resources (EDS cluster
// names for CDS, RDS routes names for LDS).
func GetResourceReferences(resources map[string]Resource) map[string]bool {
	out := make(map[string]bool)
	for _, res := range resources {
		if res == nil {
			continue
		}
		switch v := res.(type) {
		case *v2.ClusterLoadAssignment:
			// no dependencies
		case *v2.Cluster:
			// for EDS type, use cluster name or ServiceName override
			if v.Type == v2.Cluster_EDS {
				if v.EdsClusterConfig != nil && v.EdsClusterConfig.ServiceName != "" {
					out[v.EdsClusterConfig.ServiceName] = true
				} else {
					out[v.Name] = true
				}
			}
		case *v2.RouteConfiguration:
			// References to clusters in both routes (and listeners) are not included
			// in the result, because the clusters are retrieved in bulk currently,
			// and not by name.
		case *v2.Listener:
			// extract route configuration names from HTTP connection manager
			for _, chain := range v.FilterChains {
				for _, filter := range chain.Filters {
					if filter.Name != HTTPConnectionManager {
						continue
					}

					config := &hcm.HttpConnectionManager{}
					if util.StructToMessage(filter.Config, config) == nil && config != nil {
						if rds, ok := config.RouteSpecifier.(*hcm.HttpConnectionManager_Rds); ok && rds != nil && rds.Rds != nil {
							out[rds.Rds.RouteConfigName] = true
						}
					}
				}
			}
		}
	}
	return out
}
