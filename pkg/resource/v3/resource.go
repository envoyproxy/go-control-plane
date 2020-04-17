package resource

import (
	"github.com/golang/protobuf/ptypes"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
)

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

// Fetch urls in xDS v3.
const (
	FetchEndpoints = "/v3/discovery:endpoints"
	FetchClusters  = "/v3/discovery:clusters"
	FetchListeners = "/v3/discovery:listeners"
	FetchRoutes    = "/v3/discovery:routes"
	FetchSecrets   = "/v3/discovery:secrets"
	FetchRuntimes  = "/v3/discovery:runtime"
)

// DefaultAPIVersion is the api version
const DefaultAPIVersion = core.ApiVersion_V3

// GetHTTPConnectionManager creates a HttpConnectionManager from filter
func GetHTTPConnectionManager(filter *listener.Filter) *hcm.HttpConnectionManager {
	config := &hcm.HttpConnectionManager{}

	// use typed config if available
	if typedConfig := filter.GetTypedConfig(); typedConfig != nil {
		ptypes.UnmarshalAny(typedConfig, config)
	}
	return config
}

// ConvertTypeURL converts typeurl from v2->v3 and vice versa
func ConvertTypeURL(typeURL string) string {
	var convertedURL string
	switch typeURL {
	case EndpointType:
		convertedURL = "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment"
	case ClusterType:
		convertedURL = "type.googleapis.com/envoy.api.v2.Cluster"
	case RouteType:
		convertedURL = "type.googleapis.com/envoy.api.v2.RouteConfiguration"
	case ListenerType:
		convertedURL = "type.googleapis.com/envoy.api.v2.Listener"
	case SecretType:
		convertedURL = "type.googleapis.com/envoy.api.v2.auth.Secret"
	case RuntimeType:
		convertedURL = "type.googleapis.com/envoy.service.discovery.v2.Runtime"
	}
	return convertedURL
}
