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
	// Indicates whether the delta stream currently has a wildcard watch
	wildcard bool

	// Provides the list of resources explicitly requested by the client
	// This list might be non-empty even when set as wildcard
	subscribedResourceNames map[string]struct{}

	// ResourceVersions contains a hash of the resource as the value and the resource name as the key.
	// This field stores the last state sent to the client.
	resourceVersions map[string]string
}

// GetSubscribedResourceNames returns the list of resources currently explicitly subscribed to
// If the request is set to wildcard it may be empty
// Currently populated only when using delta-xds
func (s *StreamState) GetSubscribedResources() map[string]struct{} {
	return s.subscribedResourceNames
}

// SetSubscribedResourceNames is setting the list of resources currently explicitly subscribed to
// It is decorrelated from the wildcard state of the stream
// Currently used only when using delta-xds
func (s *StreamState) SetSubscribedResources(subscribedResourceNames map[string]struct{}) {
	s.subscribedResourceNames = subscribedResourceNames
}

func (s *StreamState) GetKnownResources() map[string]string {
	return s.resourceVersions
}

func (s *StreamState) SetResourceVersions(resourceVersions map[string]string) {
	s.resourceVersions = resourceVersions
}

func (s *StreamState) SetWildcard(wildcard bool) {
	s.wildcard = wildcard
}

func (s *StreamState) IsWildcard() bool {
	return s.wildcard
}

// NewStreamState initializes a stream state.
func NewStreamState(wildcard bool, initialResourceVersions map[string]string) StreamState {
	state := StreamState{
		wildcard:                wildcard,
		subscribedResourceNames: map[string]struct{}{},
		resourceVersions:        initialResourceVersions,
	}

	if initialResourceVersions == nil {
		state.resourceVersions = make(map[string]string)
	}

	return state
}
