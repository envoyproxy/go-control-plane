package stream

import (
	"google.golang.org/grpc"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
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

// StreamState will keep track of resource state per type on a stream.
type StreamState struct {
	// Indicates whether the original DeltaRequest was a wildcard LDS/RDS request.
	Wildcard bool

	// ResourceVersions contains a hash of the resource as the value and the resource name as the key.
	// This field stores the last state sent to the client.
	ResourceVersions map[string]string
}

// NewStreamState initializes a stream state.
func NewStreamState(wildcard bool, initialResourceVersions map[string]string) StreamState {
	state := StreamState{
		Wildcard:         wildcard,
		ResourceVersions: initialResourceVersions,
	}

	if initialResourceVersions == nil {
		state.ResourceVersions = make(map[string]string)
	}

	return state
}
