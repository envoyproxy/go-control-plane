package utils

import (
	"github.com/golang/protobuf/ptypes"

	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
)

// GetHTTPConnectionManager creates a HttpConnectionManager from filter
func GetHTTPConnectionManager(filter listener.Filter) *hcm.HttpConnectionManager {
	config := &hcm.HttpConnectionManager{}

	// use typed config if available
	if typedConfig := filter.GetTypedConfig(); typedConfig != nil {
		ptypes.UnmarshalAny(typedConfig, config)
	}
	return config
}
