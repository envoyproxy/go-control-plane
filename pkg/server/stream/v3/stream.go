package stream

import (
	"sync"

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

// StreamState will keep track of resource state on a stream
type StreamState struct {
	mu sync.RWMutex

	// Indicates whether the original DeltaRequest was a wildcard LDS/RDS request.
	IsWildcard map[string]bool

	// Indicates whether this is the first request on the stream for the corresponding resource type
	IsFirst map[string]bool

	// ResourceVersions contains a hash of the resource as the value and the resource name as the key.
	// This field stores the last state sent to the client.
	ResourceVersions map[string]string
}

// NewStreamState initializes a stream state.
func NewStreamState() *StreamState {
	return &StreamState{
		IsWildcard:       make(map[string]bool),
		IsFirst:          make(map[string]bool),
		ResourceVersions: make(map[string]string),
	}
}

func (s *StreamState) GetResourceVersions() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.ResourceVersions
}
