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
type StreamState struct { // nolint:golint,revive
	// Indicates whether the original DeltaRequest was a wildcard LDS/RDS request.
	wildcard bool

	// ResourceVersions contains a hash of the resource as the value and the resource name as the key.
	// This field stores the last state sent to the client.
	resourceVersions map[string]string

	// indicates whether the object has beed modified since its creation
	first bool
}

func (s *StreamState) GetResourceVersions() map[string]string {
	return s.resourceVersions
}

func (s *StreamState) SetResourceVersions(resourceVersions map[string]string) {
	s.first = false
	s.resourceVersions = resourceVersions
}

func (s *StreamState) IsFirst() bool {
	return s.first
}

func (s *StreamState) IsWildcard() bool {
	return s.wildcard
}

// NewStreamState initializes a stream state.
func NewStreamState(wildcard bool, initialResourceVersions map[string]string) StreamState {
	state := StreamState{
		wildcard:         wildcard,
		resourceVersions: initialResourceVersions,
		first:            true,
	}

	if initialResourceVersions == nil {
		state.resourceVersions = make(map[string]string)
	}

	return state
}
