package cache_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

func TestIndexResourcesByName(t *testing.T) {
	tests := []struct {
		name      string
		resources []types.ResourceWithTTL
		want      map[string]types.ResourceWithTTL
	}{
		{
			name:      "empty",
			resources: nil,
			want:      map[string]types.ResourceWithTTL{},
		},
		{
			name: "more than one",
			resources: []types.ResourceWithTTL{
				{Resource: testEndpoint, TTL: &ttl},
				{Resource: testRoute, TTL: &ttl},
			},
			want: map[string]types.ResourceWithTTL{
				"cluster0": {Resource: testEndpoint, TTL: &ttl},
				"route0":   {Resource: testRoute, TTL: &ttl},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cache.IndexResourcesByName(tt.resources)
			require.NoError(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIndexResourcesByName_DuplicateName(t *testing.T) {
	resources := []types.ResourceWithTTL{
		{
			Resource: testEndpoint,
			TTL:      &ttl,
		},
		{
			Resource: testEndpoint,
			TTL:      &ttl,
		},
	}

	_, err := cache.IndexResourcesByName(resources)
	require.Error(t, err, "Expected to get an error if we duplicated resource names")
}

func TestIndexRawResourceByName(t *testing.T) {
	tests := []struct {
		name      string
		resources []types.Resource
		want      map[string]types.Resource
	}{
		{
			name:      "empty",
			resources: nil,
			want:      map[string]types.Resource{},
		},
		{
			name: "more than one",
			resources: []types.Resource{
				testEndpoint,
				testRoute,
			},
			want: map[string]types.Resource{
				"cluster0": testEndpoint,
				"route0":   testRoute,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cache.IndexRawResourcesByName(tt.resources)
			require.NoError(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewResources(t *testing.T) {
	resources, err := cache.NewResources("x", []types.Resource{
		testEndpoint,
		testRoute,
	})
	require.NoError(t, err)

	assert.NotNil(t, resources.Items)
	assert.Equal(t, "x", resources.Version)
}
