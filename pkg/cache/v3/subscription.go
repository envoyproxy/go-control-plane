package cache

// watchSubscription is used to provide the minimum feature set we need from subscription.
// As only a few APIs are used, and we don't return the subscription, this allows the cache
// to store an altered version of the subscription without modifying the immutable one provided.
type watchSubscription struct {
	returnedResources   map[string]string
	subscribedResources map[string]struct{}
	isWildcard          bool
}

func (s watchSubscription) ReturnedResources() map[string]string {
	return s.returnedResources
}

func (s watchSubscription) SubscribedResources() map[string]struct{} {
	return s.subscribedResources
}

func (s watchSubscription) IsWildcard() bool {
	return s.isWildcard
}

func newWatchSubscription(s Subscription) watchSubscription {
	clone := watchSubscription{
		isWildcard:          s.IsWildcard(),
		returnedResources:   make(map[string]string, len(s.ReturnedResources())),
		subscribedResources: make(map[string]struct{}, len(s.SubscribedResources())),
	}
	for name, version := range s.ReturnedResources() {
		clone.returnedResources[name] = version
	}
	for name := range s.SubscribedResources() {
		clone.subscribedResources[name] = struct{}{}
	}
	return clone
}
