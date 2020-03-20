package utils

import (
	"github.com/golang/protobuf/ptypes"

	listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	hcm "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/envoyproxy/go-control-plane/pkg/conversion"
)

// GetHTTPConnectionManager creates a HttpConnectionManager from filter
func GetHTTPConnectionManager(filter *listener.Filter) *hcm.HttpConnectionManager {
	config := &hcm.HttpConnectionManager{}

	// use typed config if available
	if typedConfig := filter.GetTypedConfig(); typedConfig != nil {
		ptypes.UnmarshalAny(typedConfig, config)
	} else {
		conversion.StructToMessage(filter.GetConfig(), config)
	}
	return config
}
