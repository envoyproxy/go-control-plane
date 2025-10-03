package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSotwSubscriptions(t *testing.T) {
	t.Run("legacy mode properly handled", func(t *testing.T) {
		sub := NewSotwSubscription([]string{})
		assert.True(t, sub.IsWildcard())

		// Requests always set empty in legacy mode
		sub.SetResourceSubscription([]string{})
		assert.True(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())

		// Requests always set empty in legacy mode
		sub.SetResourceSubscription(nil)
		assert.True(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())

		// Set any resource, no longer wildcard
		sub.SetResourceSubscription([]string{"resource"})
		assert.False(t, sub.IsWildcard())
		assert.Equal(t, map[string]struct{}{"resource": {}}, sub.SubscribedResources())

		// No longer watch any resource, should not come back to wildcard as no longer in legacy mode
		// We end up with a watch to nothing
		sub.SetResourceSubscription(nil)
		assert.False(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())
	})

	t.Run("new wildcard mode from start", func(t *testing.T) {
		// A resource is provided so the subscription was created in wildcard
		sub := NewSotwSubscription([]string{"*"})
		assert.True(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())

		// Keep wildcard, no change
		sub.SetResourceSubscription([]string{"*"})
		assert.True(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())

		// Add resource to wildcard
		sub.SetResourceSubscription([]string{"*", "resource"})
		assert.True(t, sub.IsWildcard())
		assert.Equal(t, map[string]struct{}{"resource": {}}, sub.SubscribedResources())

		// Add/Remove resource to wildcard
		sub.SetResourceSubscription([]string{"*", "otherresource"})
		assert.True(t, sub.IsWildcard())
		assert.Equal(t, map[string]struct{}{"otherresource": {}}, sub.SubscribedResources())

		// Remove wildcard
		sub.SetResourceSubscription([]string{"otherresource"})
		assert.False(t, sub.IsWildcard())
		assert.Equal(t, map[string]struct{}{"otherresource": {}}, sub.SubscribedResources())

		// Remove last resource
		sub.SetResourceSubscription([]string{})
		assert.False(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())

		// Re-subscribe to wildcard
		sub.SetResourceSubscription([]string{"*"})
		assert.True(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())
	})
}

func TestDeltaSubscriptions(t *testing.T) {
	t.Run("legacy mode properly handled", func(t *testing.T) {
		sub := NewDeltaSubscription([]string{}, []string{}, map[string]string{"resource": "version"})
		assert.True(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())
		assert.Equal(t, map[string]string{"resource": "version"}, sub.ReturnedResources())

		// New request with no additional subscription
		sub.UpdateResourceSubscriptions(nil, nil)
		assert.True(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())
		assert.Equal(t, map[string]string{"resource": "version"}, sub.ReturnedResources())

		// New request adding a resource
		sub.UpdateResourceSubscriptions([]string{"resource"}, nil)
		assert.True(t, sub.IsWildcard()) // Wildcard not unsubscribed
		assert.Equal(t, map[string]struct{}{"resource": {}}, sub.SubscribedResources())
		assert.Equal(t, map[string]string{"resource": "version"}, sub.ReturnedResources())

		// Unsubscribe from "resource", still wildcard
		sub.UpdateResourceSubscriptions(nil, []string{"resource"})
		assert.True(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())
		// Version is set to "" to trigger an update or have the resource in the "removed" field
		// when explicitly unsubscribing from wildcard, to align with
		// https://www.envoyproxy.io/docs/envoy/v1.29.0/api-docs/xds_protocol#xds-protocol-unsubscribe
		assert.Equal(t, map[string]string{"resource": ""}, sub.ReturnedResources())
	})

	t.Run("new wildcard mode", func(t *testing.T) {
		// A resource is provided so the subscription was created in wildcard
		sub := NewDeltaSubscription([]string{"*"}, []string{}, map[string]string{"resource": "version"})
		assert.True(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())

		// New request with no additional subscription
		sub.UpdateResourceSubscriptions(nil, nil)
		assert.True(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())
		assert.Equal(t, map[string]string{"resource": "version"}, sub.ReturnedResources())

		// Add resource to wildcard
		sub.UpdateResourceSubscriptions([]string{"resource"}, []string{})
		assert.True(t, sub.IsWildcard())
		assert.Equal(t, map[string]struct{}{"resource": {}}, sub.SubscribedResources())
		assert.Equal(t, map[string]string{"resource": "version"}, sub.ReturnedResources())

		// Unsubscribe from resource while wildcard
		sub.UpdateResourceSubscriptions([]string{"otherresource"}, []string{"resource"})
		assert.True(t, sub.IsWildcard())
		assert.Equal(t, map[string]struct{}{"otherresource": {}}, sub.SubscribedResources())
		// Version is set to "" to trigger an update or have the resource in the "removed" field
		// when explicitly unsubscribing from wildcard, to align with
		// https://www.envoyproxy.io/docs/envoy/v1.29.0/api-docs/xds_protocol#xds-protocol-unsubscribe
		assert.Equal(t, map[string]string{"resource": ""}, sub.ReturnedResources())

		sub.SetReturnedResources(nil)

		// Remove subscription to wildcard
		sub.UpdateResourceSubscriptions([]string{"resource"}, []string{"*"})
		assert.False(t, sub.IsWildcard())
		assert.Equal(t, map[string]struct{}{"resource": {}, "otherresource": {}}, sub.SubscribedResources())
		assert.Empty(t, sub.ReturnedResources())

		// Remove all subscriptions
		// Does not come back to wildcard
		sub.UpdateResourceSubscriptions([]string{}, []string{"resource", "otherresource"})
		assert.False(t, sub.IsWildcard())
		assert.Empty(t, sub.SubscribedResources())

		// Attempt to remove wildcard when not subscribed
		sub.UpdateResourceSubscriptions([]string{"resource"}, []string{"*"})
		assert.False(t, sub.IsWildcard())
		assert.Equal(t, map[string]struct{}{"resource": {}}, sub.SubscribedResources())

		// Resubscribe to wildcard
		sub.UpdateResourceSubscriptions([]string{"*"}, nil)
		assert.True(t, sub.IsWildcard())
		assert.Equal(t, map[string]struct{}{"resource": {}}, sub.SubscribedResources())

		// Attempt to remove not-subscribed resource. Should just be ignored
		sub.UpdateResourceSubscriptions([]string{}, []string{"otherresource"})
		assert.True(t, sub.IsWildcard())
		assert.Equal(t, map[string]struct{}{"resource": {}}, sub.SubscribedResources())
	})
}
