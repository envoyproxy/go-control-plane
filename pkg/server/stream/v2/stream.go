package stream

import (
	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
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
	// Indicates whether
	IsWildcard bool

	// ResourceVersions contain a hash of the resource as a version map
	// this field stores the last state sent to the client
	ResourceVersions map[string]string
}
