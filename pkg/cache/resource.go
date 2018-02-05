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

package cache

import (
	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/gogo/protobuf/proto"
)

// ResponseType is an enumeration of cache response types.
type ResponseType int

const (
	// EndpointResponse for EDS.
	EndpointResponse ResponseType = iota

	// ClusterResponse for CDS.
	ClusterResponse

	// RouteResponse for RDS.
	RouteResponse

	// ListenerResponse for LDS.
	ListenerResponse
)

var (
	// ResponseTypes are supported response types.
	ResponseTypes = []ResponseType{
		EndpointResponse,
		ClusterResponse,
		RouteResponse,
		ListenerResponse,
	}
)

func (typ ResponseType) String() string {
	switch typ {
	case EndpointResponse:
		return "endpoints"
	case ClusterResponse:
		return "clusters"
	case RouteResponse:
		return "routes"
	case ListenerResponse:
		return "listeners"
	default:
		return "unknown"
	}
}

// GetResourceName returns the resource name for a valid xDS response type.
func GetResourceName(xds proto.Message) string {
	switch v := xds.(type) {
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
