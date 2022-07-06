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
	// Indicates whether the stream currently has a wildcard watch
	wildcard bool

	// Provides the list of resources explicitly requested by the client
	// This list might be non-empty even when set as wildcard
	subscribedResourceNames map[string]struct{}

	// ResourceVersions contains a hash of the resource as the value and the resource name as the key.
	// This field stores the last state sent to the client.
	resourceVersions map[string]string

	// indicates whether the stream is handling its first request/response or not
	first bool
}

// GetSubscribedResourceNames returns the list of resources currently explicitly subscribed to
// If the request is set to wildcard it may be empty.
// User should not alter the map through this accessor and update subscriptions through the methods below.
func (s *StreamState) GetSubscribedResourceNames() map[string]struct{} {
	return s.subscribedResourceNames
}

// delta-xds subscriptions

// In delta mode, this method registers new resources requested considering the specific wildcard case.
// The list of subscribed resources remains populated in wildcard mode to properly handle unsubscription from wildcard.
func (s *StreamState) SubscribeToResources(resources []string) {
	for _, resource := range resources {
		if resource == "*" {
			s.wildcard = true
			continue
		}
		s.subscribedResourceNames[resource] = struct{}{}
	}
}

// In delta mode, this method removes resources no longer requested considering the specific wildcard case.
// If a client explicitly unsubscribes from wildcard, the stream wildcard status is updated and now tracks only previously subscribed resources.
func (s *StreamState) UnsubscribeFromResources(resources []string) {
	for _, resource := range resources {
		if resource == "*" {
			s.wildcard = false
			continue
		}
		if _, ok := s.subscribedResourceNames[resource]; ok && s.IsWildcard() {
			// The XDS protocol states that:
			// * if a watch is currently wildcard
			// * a resource is explicitly unsubscribed by name
			// Then the control-plane must return in the response whether the resource is removed (if no longer present for this node)
			// or still existing. In the latter case the entire resource must be returned, same as if it had been created or updated
			// To achieve that, we mark the resource as having been returned with an empty version. While creating the response, the cache will either:
			// * detect the version change, and return the resource (as an update)
			// * detect the resource deletion, and set it as removed in the response
			s.GetResourceVersions()[resource] = ""
		}
		delete(s.subscribedResourceNames, resource)
	}
}

// sotw-xds subscriptions

// RegisterSubscribedResources is used in sotw mode to update the streamstate when new requests are received
// It sets both the wildcard flag and the list of resources currently being subscribed to.
func (s *StreamState) RegisterSubscribedResources(resources []string) {
	// sotw fully wipes out the previously registered resources.
	s.subscribedResourceNames = make(map[string]struct{}, len(resources))
	if len(resources) == 0 {
		// The xDS protocol states that if there has never been any resource set, the request should be considered wildcard (legacy mode),
		// but also indicates that if it ever included any resource, it should be considered empty and non-wildcard.
		// This would require keeping track on whether we ever received any resource (including wildcard itself).
		// As the protocol also allows to return resources which have not been subscribed to, it is an implementation choice to potentially
		// not get out of wildcard mode instead of empty mode in a case when no resources are received while the stream is currently wildcard.
		return
	}

	// When resources are provided, they may still include the wildcard symbol '*', as well as potentially other resources.
	wantsWildcard := false
	for _, resourceName := range resources {
		// We do not track '*' as a resource name to avoid confusion in further processing and rely on the IsWildcard method instead.
		if resourceName == "*" {
			wantsWildcard = true
			continue
		}
		s.subscribedResourceNames[resourceName] = struct{}{}
	}
	s.wildcard = wantsWildcard
}

// WatchesResources returns whether at least one of the resource provided is currently watch by the stream
// It is currently only applicable to delta-xds
// If the request is wildcard, it will always return true
// Otherwise it will compare the provided resources to the list of resources currently subscribed
func (s *StreamState) WatchesResources(resourceNames map[string]struct{}) bool {
	if s.IsWildcard() {
		return true
	}
	for resourceName := range resourceNames {
		if _, ok := s.subscribedResourceNames[resourceName]; ok {
			return true
		}
	}
	return false
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
// The user is currently required to provide the initial wildcard status as compatibility
// with the legacy wildcard mode cannot be trivially implemented for delta mode.
func NewStreamState(wildcard bool, initialResourceVersions map[string]string) StreamState {
	state := StreamState{
		wildcard:                wildcard,
		subscribedResourceNames: map[string]struct{}{},
		resourceVersions:        initialResourceVersions,
		first:                   true,
	}

	if initialResourceVersions == nil {
		state.resourceVersions = make(map[string]string)
	}

	return state
}
