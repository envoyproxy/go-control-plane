package stream

// SubscriptionState stores the server view of a given type subscription in a stream.
type SubscriptionState struct {
	// wildcard indicates if the subscription currently has a wildcard watch.
	wildcard bool

	// subscribedResourceNames provides the resources explicitly requested by the client
	// This list might be non-empty even when set as wildcard.
	subscribedResourceNames map[string]struct{}

	// ackedResources contains the resources acknowledged by the client and the acknowledged versions.
	ackedResources map[string]string
}

// NewSubscriptionState initializes a stream state.
func NewSubscriptionState(wildcard bool, initialResourceVersions map[string]string) SubscriptionState {
	state := SubscriptionState{
		wildcard:                wildcard,
		subscribedResourceNames: map[string]struct{}{},
		ackedResources:          initialResourceVersions,
	}

	if initialResourceVersions == nil {
		state.ackedResources = make(map[string]string)
	}

	return state
}

// GetSubscribedResources returns the list of resources currently explicitly subscribed to
// If the request is set to wildcard it may be empty
// Currently populated only when using delta-xds
func (s SubscriptionState) GetSubscribedResources() map[string]struct{} {
	return s.subscribedResourceNames
}

// SetSubscribedResources is setting the list of resources currently explicitly subscribed to
// It is decorrelated from the wildcard state of the stream
// Currently used only when using delta-xds
func (s *SubscriptionState) SetSubscribedResources(subscribedResourceNames map[string]struct{}) {
	s.subscribedResourceNames = subscribedResourceNames
}

// GetACKedResources returns the list of resources acknowledged by the client
// and their acknowledged version
func (s SubscriptionState) GetACKedResources() map[string]string {
	return s.ackedResources
}

// SetACKedResources sets a list of resource versions currently known by the client
// The cache can use this state to compute resources added/updated/deleted
func (s *SubscriptionState) SetACKedResources(resourceVersions map[string]string) {
	s.ackedResources = resourceVersions
}

// SetWildcard will set the subscription to return all known resources
func (s *SubscriptionState) SetWildcard(wildcard bool) {
	s.wildcard = wildcard
}

// IsWildcard returns whether or not the subscription currently has a wildcard watch
func (s SubscriptionState) IsWildcard() bool {
	return s.wildcard
}

// WatchesResources returns whether at least one of the resources provided is currently being watched by the subscription.
// If the request is wildcard, it will always return true,
// otherwise it will compare the provided resources to the list of resources currently subscribed
func (s SubscriptionState) WatchesResources(resourceNames map[string]struct{}) bool {
	if s.wildcard {
		return true
	}
	for resourceName := range resourceNames {
		if _, ok := s.subscribedResourceNames[resourceName]; ok {
			return true
		}
	}
	return false
}
