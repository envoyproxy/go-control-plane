package internal

import (
	"errors"
	"testing"
	"time"

	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/anypb"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSotwResource(t *testing.T) {
	ttl := 52 * time.Second
	t.Run("successful", func(t *testing.T) {
		c := &cluster.Cluster{
			Name: "test",
		}

		t.Run("without ttl", func(t *testing.T) {
			res := NewCachedResource(c.Name, c, WithCacheVersion("42"))

			serialized, err := res.GetSotwResource(false)
			require.NoError(t, err)

			// Returned value is simply anypb of the resource.
			ret := &cluster.Cluster{}
			require.NoError(t, serialized.UnmarshalTo(ret))
			assert.Empty(t, cmp.Diff(c, ret, protocmp.Transform()))
		})

		t.Run("with ttl", func(t *testing.T) {
			res := NewCachedResource(c.Name, c, WithCacheVersion("42"), WithResourceTTL(&ttl))

			serialized, err := res.GetSotwResource(false)
			require.NoError(t, err)

			// Returned resource is a discovery resource.
			wrapper := &discovery.Resource{}
			require.NoError(t, serialized.UnmarshalTo(wrapper))
			assert.Equal(t, c.Name, wrapper.Name)
			require.NotNil(t, wrapper.Ttl)
			assert.Equal(t, ttl, wrapper.Ttl.AsDuration())
			require.NotNil(t, wrapper.Resource)

			ret := &cluster.Cluster{}
			require.NoError(t, wrapper.Resource.UnmarshalTo(ret))
			assert.Empty(t, cmp.Diff(c, ret, protocmp.Transform()))
		})

		t.Run("without resource", func(t *testing.T) {
			anybytes, err := anypb.New(c)
			require.NoError(t, err)
			// Do not provide any resource, only an opaque serialized one.
			res := NewCachedResource(c.Name, nil, WithCacheVersion("42"), WithMarshaledResource(anybytes))

			serialized, err := res.GetSotwResource(false)
			require.NoError(t, err)

			// Returned value is simply anypb of the resource.
			ret := &cluster.Cluster{}
			require.NoError(t, serialized.UnmarshalTo(ret))
			assert.Empty(t, cmp.Diff(c, ret, protocmp.Transform()))
		})
	})

	t.Run("error on marshaling", func(t *testing.T) {
		returnedErr := errors.New("failed to serialize")
		c := &cluster.Cluster{}
		res := NewCachedResource("test", c, WithCacheVersion("42"))
		res.marshalFunc = func() (*anypb.Any, error) { return nil, returnedErr }

		_, err := res.GetSotwResource(false)
		assert.ErrorIs(t, err, returnedErr)
	})

	t.Run("error on hashing", func(t *testing.T) {
		returnedErr := errors.New("failed to compute version")
		c := &cluster.Cluster{}
		res := NewCachedResource("test", c, WithCacheVersion("42"))
		res.computeResourceVersionFunc = func() (string, error) { return "", returnedErr }

		_, err := res.GetSotwResource(false)
		// Resource version is not computed in sotw
		assert.NoError(t, err)
	})
}

func TestSotwHeartBeatResource(t *testing.T) {
	ttl := 52 * time.Second
	t.Run("successful", func(t *testing.T) {
		c := &cluster.Cluster{
			Name: "test",
		}

		t.Run("without ttl", func(t *testing.T) {
			res := NewCachedResource(c.Name, c, WithCacheVersion("42"))

			_, err := res.GetSotwResource(true)
			assert.Error(t, err)
		})

		t.Run("with ttl", func(t *testing.T) {
			res := NewCachedResource(c.Name, c, WithCacheVersion("42"), WithResourceTTL(&ttl))

			serialized, err := res.GetSotwResource(true)
			require.NoError(t, err)

			// Returned resource is a discovery resource without the actual content.
			wrapper := &discovery.Resource{}
			require.NoError(t, serialized.UnmarshalTo(wrapper))
			assert.Equal(t, c.Name, wrapper.Name)
			require.NotNil(t, wrapper.Ttl)
			assert.Equal(t, ttl, wrapper.Ttl.AsDuration())
			assert.Nil(t, wrapper.Resource)
		})
	})

	t.Run("error on marshaling", func(t *testing.T) {
		returnedErr := errors.New("failed to serialize")
		c := &cluster.Cluster{}
		res := NewCachedResource("test", c, WithCacheVersion("42"), WithResourceTTL(&ttl))
		res.marshalFunc = func() (*anypb.Any, error) { return nil, returnedErr }

		_, err := res.GetSotwResource(true)
		// Resource content is not computed in sotw
		assert.NoError(t, err)
	})

	t.Run("error on hashing", func(t *testing.T) {
		returnedErr := errors.New("failed to compute version")
		c := &cluster.Cluster{}
		res := NewCachedResource("test", c, WithCacheVersion("42"), WithResourceTTL(&ttl))
		res.computeResourceVersionFunc = func() (string, error) { return "", returnedErr }

		_, err := res.GetSotwResource(true)
		// Resource version is not computed in sotw
		assert.NoError(t, err)
	})
}

func TestDeltaResource(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		c := &cluster.Cluster{
			Name: "test",
		}
		anySerialized, err := anypb.New(c)
		require.NoError(t, err)
		cacheVersion := "42"
		defaultResourceVersion := hashResource(anySerialized.Value)

		validate := func(t testing.TB, serialized *discovery.Resource, expectedVersion string) {
			assert.Equal(t, c.Name, serialized.Name)
			if expectedVersion != "" {
				assert.Equal(t, expectedVersion, serialized.Version)
			} else {
				assert.Equal(t, defaultResourceVersion, serialized.Version)
			}
			assert.Nil(t, serialized.Ttl)
			require.NotEmpty(t, serialized.Resource.Value)

			ret := &cluster.Cluster{}
			require.NoError(t, serialized.Resource.UnmarshalTo(ret))
			assert.Empty(t, cmp.Diff(c, ret, protocmp.Transform()))
		}

		t.Run("without ttl", func(t *testing.T) {
			res := NewCachedResource(c.Name, c, WithCacheVersion(cacheVersion))

			serialized, err := res.GetDeltaResource()
			require.NoError(t, err)

			validate(t, serialized, "")
		})

		t.Run("with ttl", func(t *testing.T) {
			// TTL is currently ignored in delta
			ttl := 52 * time.Second
			res := NewCachedResource(c.Name, c, WithCacheVersion(cacheVersion), WithResourceTTL(&ttl))

			serialized, err := res.GetDeltaResource()
			require.NoError(t, err)

			validate(t, serialized, "")
		})

		t.Run("without resource", func(t *testing.T) {
			anybytes, err := anypb.New(c)
			require.NoError(t, err)
			// Do not provide any resource, only an opaque serialized one.
			res := NewCachedResource(c.Name, nil, WithCacheVersion(cacheVersion), WithMarshaledResource(anybytes))

			serialized, err := res.GetDeltaResource()
			require.NoError(t, err)

			validate(t, serialized, "")
		})

		t.Run("without resource and with version", func(t *testing.T) {
			anybytes, err := anypb.New(c)
			require.NoError(t, err)
			// Do not provide any resource, only an opaque serialized one.
			res := NewCachedResource(c.Name, nil, WithCacheVersion(cacheVersion), WithMarshaledResource(anybytes), WithResourceVersion("testversion"))

			serialized, err := res.GetDeltaResource()
			require.NoError(t, err)

			validate(t, serialized, "testversion")
		})
	})

	t.Run("error on marshaling", func(t *testing.T) {
		returnedErr := errors.New("failed to serialize")
		c := &cluster.Cluster{}
		res := NewCachedResource("test", c, WithCacheVersion("42"))
		res.marshalFunc = func() (*anypb.Any, error) { return nil, returnedErr }

		_, err := res.GetDeltaResource()
		require.ErrorIs(t, err, returnedErr)
	})

	t.Run("error on hashing", func(t *testing.T) {
		returnedErr := errors.New("failed to compute version")
		c := &cluster.Cluster{}
		res := NewCachedResource("test", c, WithCacheVersion("42"))
		res.computeResourceVersionFunc = func() (string, error) { return "", returnedErr }

		_, err := res.GetDeltaResource()
		require.ErrorIs(t, err, returnedErr)
	})
}
