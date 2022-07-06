package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSotwSubscription(t *testing.T) {
	t.Run("initial stream with no resource", func(t *testing.T) {
		state := NewStreamState(false, nil)
		assert.False(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())

		state = NewStreamState(true, nil)
		assert.True(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())
	})

	t.Run("no resources passed maintains wildcard status", func(t *testing.T) {
		state := NewStreamState(false, nil)
		assert.False(t, state.IsWildcard())
		state.RegisterSubscribedResources([]string{})
		assert.False(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())

		state = NewStreamState(true, nil)
		assert.True(t, state.IsWildcard())
		state.RegisterSubscribedResources([]string{})
		assert.True(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())
	})

	t.Run("resources passed without wildcard does remove wildcard state", func(t *testing.T) {
		state := NewStreamState(false, nil)
		assert.False(t, state.IsWildcard())
		state.RegisterSubscribedResources([]string{"a", "b"})
		assert.False(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"a": {}, "b": {}}, state.GetSubscribedResourceNames())

		state = NewStreamState(true, nil)
		assert.True(t, state.IsWildcard())
		state.RegisterSubscribedResources([]string{"a", "b"})
		assert.False(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"a": {}, "b": {}}, state.GetSubscribedResourceNames())
	})

	t.Run("new wildcard setup is supported", func(t *testing.T) {
		state := NewStreamState(false, nil)
		assert.False(t, state.IsWildcard())
		state.RegisterSubscribedResources([]string{})
		assert.False(t, state.IsWildcard())

		// We become wildcard as expected
		state.RegisterSubscribedResources([]string{"*"})
		assert.True(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())

		// This is not matching xDS behavior but is expected given the current implementation
		// Passing no resource after the initial call does reset the resource list but does not affect wildcard status
		state.RegisterSubscribedResources([]string{})
		assert.True(t, state.IsWildcard())

		// Passing any resource does reset wildcard status
		state.RegisterSubscribedResources([]string{"a"})
		assert.False(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"a": {}}, state.GetSubscribedResourceNames())

		state.RegisterSubscribedResources([]string{})
		assert.False(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())
	})

	t.Run("resources are properly handled when mixed with wildcard", func(t *testing.T) {
		state := NewStreamState(false, nil)
		assert.False(t, state.IsWildcard())

		state.RegisterSubscribedResources([]string{"*", "a"})
		assert.True(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"a": {}}, state.GetSubscribedResourceNames())

		state.RegisterSubscribedResources([]string{"*", "b"})
		assert.True(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"b": {}}, state.GetSubscribedResourceNames())

		state.RegisterSubscribedResources([]string{"b", "c"})
		assert.False(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"b": {}, "c": {}}, state.GetSubscribedResourceNames())
	})
}

func TestDeltaSubscription(t *testing.T) {
	t.Run("initial stream with no resource", func(t *testing.T) {
		state := NewStreamState(false, nil)
		assert.False(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())

		state = NewStreamState(true, nil)
		assert.True(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())
	})

	t.Run("no resources added or removed maintains status", func(t *testing.T) {
		state := NewStreamState(false, nil)
		assert.False(t, state.IsWildcard())
		state.SubscribeToResources([]string{})
		state.UnsubscribeFromResources([]string{})
		assert.False(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())

		state = NewStreamState(true, nil)
		assert.True(t, state.IsWildcard())
		state.SubscribeToResources([]string{})
		state.UnsubscribeFromResources([]string{})
		assert.True(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())

		state = NewStreamState(true, nil)
		state.SubscribeToResources([]string{"a"})
		assert.True(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"a": {}}, state.GetSubscribedResourceNames())
		state.SubscribeToResources([]string{})
		state.UnsubscribeFromResources([]string{})
		assert.True(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"a": {}}, state.GetSubscribedResourceNames())
	})

	t.Run("wildcard status is properly handled when starting as wildcard", func(t *testing.T) {
		state := NewStreamState(true, nil)
		assert.True(t, state.IsWildcard())
		state.SubscribeToResources([]string{"a"})
		assert.True(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"a": {}}, state.GetSubscribedResourceNames())

		state.SubscribeToResources([]string{"b", "c"})
		state.UnsubscribeFromResources([]string{"a"})
		assert.True(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"b": {}, "c": {}}, state.GetSubscribedResourceNames())

		state.UnsubscribeFromResources([]string{"*", "b"})
		assert.False(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"c": {}}, state.GetSubscribedResourceNames())

		state.SubscribeToResources([]string{"*"})
		assert.True(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"c": {}}, state.GetSubscribedResourceNames())

		state.UnsubscribeFromResources([]string{"c"})
		assert.True(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())

		state.UnsubscribeFromResources([]string{"*"})
		assert.False(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())
	})

	t.Run("wildcard status is properly handled when starting as non-wildcard", func(t *testing.T) {
		state := NewStreamState(false, nil)
		assert.False(t, state.IsWildcard())
		state.SubscribeToResources([]string{"a"})
		assert.False(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"a": {}}, state.GetSubscribedResourceNames())

		state.SubscribeToResources([]string{"b", "c"})
		state.UnsubscribeFromResources([]string{"a"})
		assert.False(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"b": {}, "c": {}}, state.GetSubscribedResourceNames())

		state.SubscribeToResources([]string{"*"})
		state.UnsubscribeFromResources([]string{"b"})
		assert.True(t, state.IsWildcard())
		assert.Equal(t, map[string]struct{}{"c": {}}, state.GetSubscribedResourceNames())

		state.UnsubscribeFromResources([]string{"c"})
		assert.True(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())

		state.UnsubscribeFromResources([]string{"*"})
		assert.False(t, state.IsWildcard())
		assert.Empty(t, state.GetSubscribedResourceNames())
	})
}
