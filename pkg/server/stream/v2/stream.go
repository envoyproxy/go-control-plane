package stream

import (
	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	"google.golang.org/grpc"
)

// Generic RPC stream.
type Stream interface {
	grpc.ServerStream

	Send(*discovery.DiscoveryResponse) error
	Recv() (*discovery.DiscoveryRequest, error)
}

type DeltaStream interface {
	grpc.ServerStream

	Send(*discovery.DeltaDiscoveryResponse) error
	Recv() (*discovery.DeltaDiscoveryRequest, error)
}

// StreamState will keep track of resource state on a stream
type StreamState struct {
	// Keep track of a version on the stream
	SystemVersion string `json:"system_version_info"`

	// Nonce for the delta stream
	Nonce string `json:"nonce"`

	// ResourceVersions contain an alias and a hash of the resource as a version map
	ResourceVersions map[string]cache.DeltaVersionInfo `json:"resource_versions"`
}

// GetVersionMap will return the currently applied version map within a streams state
func (s StreamState) GetVersionMap() map[string]cache.DeltaVersionInfo {
	return s.ResourceVersions
}

// GetStreamNonce will return the currently applied stream nonce with a streams state
func (s StreamState) GetStreamNonce() string {
	return s.Nonce
}

// GetSystemVersion will return the currently applied system version within a stream (i.e. what envoy set)
func (s StreamState) GetSystemVersion() string {
	return s.SystemVersion
}
