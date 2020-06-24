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
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/golang/protobuf/proto"

	cluster "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	listener "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	hcm "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	runtime "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v2"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
)

// GetResponseType returns the enumeration for a valid xDS type URL
func GetResponseType(typeURL string) types.ResponseType {
	switch typeURL {
	case resource.EndpointType:
		return types.Endpoint
	case resource.ClusterType:
		return types.Cluster
	case resource.RouteType:
		return types.Route
	case resource.ListenerType:
		return types.Listener
	case resource.SecretType:
		return types.Secret
	case resource.RuntimeType:
		return types.Runtime
	}
	return types.UnknownType
}

// GetResponseTypeURL returns the type url for a valid enum
func GetResponseTypeURL(responseType types.ResponseType) string {
	switch responseType {
	case types.Endpoint:
		return resource.EndpointType
	case types.Cluster:
		return resource.ClusterType
	case types.Route:
		return resource.RouteType
	case types.Listener:
		return resource.ListenerType
	case types.Secret:
		return resource.SecretType
	case types.Runtime:
		return resource.RuntimeType
	}
	return resource.AnyType
}

// GetResourceName returns the resource name for a valid xDS response type.
func GetResourceName(res types.Resource) string {
	switch v := res.(type) {
	case *endpoint.ClusterLoadAssignment:
		return v.GetClusterName()
	case *cluster.Cluster:
		return v.GetName()
	case *route.RouteConfiguration:
		return v.GetName()
	case *listener.Listener:
		return v.GetName()
	case *auth.Secret:
		return v.GetName()
	case *runtime.Runtime:
		return v.GetName()
	default:
		return ""
	}
}

// MarshalResource converts the Resource to MarshaledResource
func MarshalResource(resource types.Resource) (types.MarshaledResource, error) {
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
func GetResourceReferences(resources map[string]types.Resource) map[string]bool {
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

					config := resource.GetHTTPConnectionManager(filter)

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

// HashResources will take a list of resources and create a SHA256 hash sum out of the marshaled bytes
func HashResources(resources []types.Resource) (string, error) {
	var source []byte
	hasher := sha256.New()

	// Add our resources to the byte source for the hash
	for _, resource := range resources {
		b := proto.NewBuffer(nil)
		b.SetDeterministic(true)
		err := b.Marshal(resource)
		if err != nil {
			return "", err
		}

		source = append(source, b.Bytes()...)
	}

	// Use the bytes as our hashing source
	hasher.Write(source)
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// RandomHash will take a list of resources and create a SHA256 hash sum out of the marshaled bytes
func RandomHash() string {
	// Use current time in nanoseconds
	source := rand.NewSource(time.Now().UnixNano())

	b := make([]byte, 32)
	binary.LittleEndian.PutUint64(b[0:], rand.New(source).Uint64())

	hasher := sha256.New()
	hasher.Write(b)

	return hex.EncodeToString(hasher.Sum(nil))
}
