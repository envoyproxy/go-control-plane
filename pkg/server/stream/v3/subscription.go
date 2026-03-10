package stream

import "strings"

const (
	explicitWildcard = "*"
	globSuffix       = "/*"
)

// Subscription stores the server view of a given type subscription in a stream.
type Subscription struct {
	// wildcard indicates if the subscription currently has a wildcard watch.
	wildcard bool

	// allowLegacyWildcard indicates that the stream never provided any resource
	// and is de facto wildcard.
	// As soon as a resource or an explicit subscription to wildcard is provided,
	// this flag will be set to false
	allowLegacyWildcard bool

	// subscribedResourceNames provides the resources explicitly requested by the client.
	// Prefix glob subscriptions (e.g. "collection/*") are stored separately in subscribedPrefixes.
	// This list might be non-empty even when set as wildcard.
	subscribedResourceNames map[string]struct{}

	// subscribedPrefixes stores the extracted prefixes from glob collection subscriptions
	// (e.g. subscribing to "collection/*" stores "collection/" here).
	subscribedPrefixes map[string]struct{}

	// returnedResources contains the resources acknowledged by the client and the acknowledged versions.
	returnedResources map[string]string
}

// newSubscription initializes a subscription state.
func newSubscription(emptyRequest, allowLegacyWildcard bool, initialResourceVersions map[string]string) Subscription {
	// By default we set the subscription as a wildcard only if the request was empty
	// and in legacy mode. Later on, outside of this constructor, when we actually
	// process the request, if the request was non-empty, it may have an
	// explicit wildcard subscription, in which case
	// we will set the wildcard field on the subscription accordingly.
	wildcard := emptyRequest && allowLegacyWildcard

	state := Subscription{
		wildcard:                wildcard,
		allowLegacyWildcard:     allowLegacyWildcard,
		subscribedResourceNames: map[string]struct{}{},
		subscribedPrefixes:      map[string]struct{}{},
		returnedResources:       initialResourceVersions,
	}

	if initialResourceVersions == nil {
		state.returnedResources = make(map[string]string)
	}

	return state
}

func NewSotwSubscription(subscribed []string, allowLegacyWildcard bool) Subscription {
	sub := newSubscription(len(subscribed) == 0, allowLegacyWildcard, nil)
	sub.SetResourceSubscription(subscribed)
	return sub
}

// SetResourceSubscription updates the subscribed resources (including the wildcard state)
// based on the full state of subscribed resources provided in the request
// Used in sotw subscriptions
// Behavior is based on
// https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#how-the-client-specifies-what-resources-to-return
func (s *Subscription) SetResourceSubscription(subscribed []string) {
	if s.allowLegacyWildcard {
		if len(subscribed) == 0 {
			// We were wildcard based on legacy behavior and still don't request any resource
			// The watch remains wildcard
			return
		}

		// A resource was provided (might be an explicit wildcard)
		// Documentation states that we should no longer allow to fallback to the previous case
		// and no longer setting wildcard would no longer subscribe to anything
		s.allowLegacyWildcard = false
	}

	subscribedResources := make(map[string]struct{}, len(subscribed))
	subscribedPrefixes := make(map[string]struct{})
	explicitWildcardSet := false
	for _, resource := range subscribed {
		switch {
		case resource == explicitWildcard:
			explicitWildcardSet = true
		case strings.HasSuffix(resource, globSuffix):
			subscribedPrefixes[strings.TrimSuffix(resource, "*")] = struct{}{}
		default:
			subscribedResources[resource] = struct{}{}
		}
	}

	if !explicitWildcardSet {
		// Cleanup resources no longer subscribed to.
		// This ensures later subscriptions will trigger responses,
		// even if the version has not changed
		for resource := range s.returnedResources {
			if _, ok := subscribedResources[resource]; ok {
				continue
			}
			if matchesAnyPrefix(subscribedPrefixes, resource) {
				continue
			}
			delete(s.returnedResources, resource)
		}
	}

	// Explicit subscription to wildcard as we are not in legacy wildcard behavior
	s.wildcard = explicitWildcardSet
	s.subscribedResourceNames = subscribedResources
	s.subscribedPrefixes = subscribedPrefixes
}

func NewDeltaSubscription(subscribed, unsubscribed []string, initialResourceVersions map[string]string, allowLegacyWildcard bool) Subscription {
	sub := newSubscription(len(subscribed) == 0, allowLegacyWildcard, initialResourceVersions)
	sub.UpdateResourceSubscriptions(subscribed, unsubscribed)
	return sub
}

// UpdateResourceSubscriptions updates the subscribed resources (including the wildcard state)
// based on newly subscribed or unsubscribed resources
// Used in delta subscriptions.
func (s *Subscription) UpdateResourceSubscriptions(subscribed, unsubscribed []string) {
	// Handles legacy wildcard behavior first to exit if we are still in this behavior
	if s.allowLegacyWildcard {
		// The protocol (as of v1.36.0) only references subscribed as triggering
		// exiting legacy wildcard behavior, so we currently not check unsubscribed
		if len(subscribed) == 0 {
			// We were wildcard based on legacy behavior and still don't request any resource
			// The watch remains wildcard
			return
		}

		// A resource was provided (might be an explicit wildcard)
		// Documentation states that we should no longer allow to fallback to the previous case
		// and no longer setting wildcard would no longer subscribe to anything
		// The watch does remain wildcard if not explicitly unsubscribed (from the example in
		// https://www.envoyproxy.io/docs/envoy/v1.29.0/api-docs/xds_protocol#how-the-client-specifies-what-resources-to-return)
		s.allowLegacyWildcard = false
	}

	// Handle subscriptions first
	for _, resource := range subscribed {
		if resource == explicitWildcard {
			s.wildcard = true
			continue
		}
		if strings.HasSuffix(resource, globSuffix) {
			s.subscribedPrefixes[strings.TrimSuffix(resource, "*")] = struct{}{}
			continue
		}
		s.subscribedResourceNames[resource] = struct{}{}
	}

	// Then unsubscriptions
	for _, resource := range unsubscribed {
		if resource == explicitWildcard {
			s.wildcard = false
			continue
		}
		if strings.HasSuffix(resource, globSuffix) {
			prefix := strings.TrimSuffix(resource, "*")
			delete(s.subscribedPrefixes, prefix)
			continue
		}
		if _, ok := s.subscribedResourceNames[resource]; ok && s.wildcard {
			// The XDS protocol states that:
			// * if a watch is currently wildcard
			// * a resource is explicitly unsubscribed by name
			// Then the control-plane must return in the response whether the resource is removed (if no longer present for this node)
			// or still existing. In the latter case the entire resource must be returned, same as if it had been created or updated
			// To achieve that, we mark the resource as having been returned with an empty version. While creating the response, the cache will either:
			// * detect the version change, and return the resource (as an update)
			// * detect the resource deletion, and set it as removed in the response
			s.returnedResources[resource] = ""
		} else {
			// Cleanup unsubscribed resources. This avoids returning a response
			// if the versions have not changed
			delete(s.returnedResources, resource)
		}
		delete(s.subscribedResourceNames, resource)
	}
}

// SubscribedResources returns the list of resources currently explicitly subscribed to.
// Prefix glob subscriptions are not included; see SubscribedPrefixes.
// If the request is set to wildcard it may be empty.
func (s Subscription) SubscribedResources() map[string]struct{} {
	return s.subscribedResourceNames
}

// SubscribedPrefixes returns the set of prefix glob subscriptions, keyed by prefix
// (e.g. subscribing to "collection/*" yields the key "collection/").
func (s Subscription) SubscribedPrefixes() map[string]struct{} {
	return s.subscribedPrefixes
}

func matchesAnyPrefix(prefixes map[string]struct{}, name string) bool {
	for prefix := range prefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}

// IsWildcard returns whether or not the subscription currently has a wildcard watch.
func (s Subscription) IsWildcard() bool {
	return s.wildcard
}

// ReturnedResources returns the list of resources returned to the client
// and their version.
func (s Subscription) ReturnedResources() map[string]string {
	return s.returnedResources
}

// SetReturnedResources sets a list of resource versions currently known by the client
// The cache can use this state to compute resources added/updated/deleted.
func (s *Subscription) SetReturnedResources(resourceVersions map[string]string) {
	s.returnedResources = resourceVersions
}
