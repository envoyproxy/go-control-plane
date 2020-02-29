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
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	secret "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	runtime "github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
)

// Resource is the base interface for the xDS payload.
type Resource interface {
	proto.Message
}

// Resource types in xDS v3.
const (
	apiTypePrefix = "type.googleapis.com/"
	EndpointType  = apiTypePrefix + "envoy.config.endpoint.v3.ClusterLoadAssignment"
	ClusterType   = apiTypePrefix + "envoy.config.cluster.v3.Cluster"
	RouteType     = apiTypePrefix + "envoy.config.route.v3.RouteConfiguration"
	ListenerType  = apiTypePrefix + "envoy.config.listener.v3.Listener"
	SecretType    = apiTypePrefix + "envoy.extensions.transport_sockets.tls.v3.Secret"
	RuntimeType   = apiTypePrefix + "envoy.service.runtime.v3.Runtime"

	// AnyType is used only by ADS
	AnyType = ""
)

// ResponseType enumeration of supported response types
type ResponseType int

const (
	Endpoint ResponseType = iota
	Cluster
	Route
	Listener
	Secret
	Runtime
	UnknownType // token to count the total number of supported types
)

// GetResponseType returns the enumeration for a valid xDS type URL
func GetResponseType(typeURL string) ResponseType {
	switch typeURL {
	case EndpointType:
		return Endpoint
	case ClusterType:
		return Cluster
	case RouteType:
		return Route
	case ListenerType:
		return Listener
	case SecretType:
		return Secret
	case RuntimeType:
		return Runtime
	}
	return UnknownType
}

// GetResourceName returns the resource name for a valid xDS response type.
func GetResourceName(res Resource) string {
	switch v := res.(type) {
	case *endpoint.ClusterLoadAssignment:
		return v.GetClusterName()
	case *cluster.Cluster:
		return v.GetName()
	case *route.RouteConfiguration:
		return v.GetName()
	case *listener.Listener:
		return v.GetName()
	case *secret.Secret:
		return v.GetName()
	case *runtime.Runtime:
		return v.GetName()
	default:
		return ""
	}
}

// MarshalResource converts the Resource to MarshaledResource
func MarshalResource(resource Resource) (MarshaledResource, error) {
	b := proto.NewBuffer(nil)
	b.SetDeterministic(true)
	err := b.Marshal(resource)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
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
		case *endpoint.ClusterLoadAssignment:
			// no dependencies
		case *cluster.Cluster:
			// for EDS type, use cluster name or ServiceName override
			switch typ := v.ClusterDiscoveryType.(type) {
			case *cluster.Cluster_Type:
				if typ.Type == cluster.Cluster_EDS {
					if v.EdsClusterConfig != nil && v.EdsClusterConfig.ServiceName != "" {
						out[v.EdsClusterConfig.ServiceName] = true
					} else {
						out[v.Name] = true
					}
				}
			}
		case *route.RouteConfiguration:
			// References to clusters in both routes (and listeners) are not included
			// in the result, because the clusters are retrieved in bulk currently,
			// and not by name.
		case *listener.Listener:
			// extract route configuration names from HTTP connection manager
			for _, chain := range v.FilterChains {
				for _, filter := range chain.Filters {
					if filter.Name != wellknown.HTTPConnectionManager {
						continue
					}

					config := &hcm.HttpConnectionManager{}

					ptypes.UnmarshalAny(filter.GetTypedConfig(), config)

					if config == nil {
						continue
					}

					if rds, ok := config.RouteSpecifier.(*hcm.HttpConnectionManager_Rds); ok && rds != nil && rds.Rds != nil {
						out[rds.Rds.RouteConfigName] = true
					}
				}
			}
		case *runtime.Runtime:
			// no dependencies
		}
	}
	return out
}
