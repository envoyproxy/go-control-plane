package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Resource is the base interface for the xDS payload.
type Resource interface {
	proto.Message
}

// ResourceWithTTL is a Resource with an optional TTL.
type ResourceWithTTL struct {
	Resource Resource
	TTL      *time.Duration
}

// CachedResource is used to track resources added by the user in the cache.
// It contains the resource itself and its associated version (currently in two different modes).
// It should not be altered once created, to allow concurrent access.
type CachedResource struct {
	Name    string
	typeURL string

	resource Resource
	ttl      *time.Duration

	// cacheVersion is the version of the cache at the time of last update, used in sotw.
	cacheVersion string

	marshalFunc                func() ([]byte, error)
	computeResourceVersionFunc func() (string, error)
}

type CachedResourceOption = func(*CachedResource)

// WithCacheVersion allows specifying the cacheVersion when the resource is set.
func WithCacheVersion(version string) CachedResourceOption {
	return func(r *CachedResource) { r.cacheVersion = version }
}

// WithMarshaledResource enables the user to provide the already marshaled bytes if they have them.
// Those bytes should strive at being consistent if the object has not changed (beware protobuf non-deterministic marshaling)
// or alternatively the resource version should also then be set.
// By default it is computed by performing a deterministic protobuf marshaling.
func WithMarshaledResource(bytes []byte) CachedResourceOption {
	if len(bytes) == 0 {
		return func(*CachedResource) {}
	}
	return func(r *CachedResource) { r.marshalFunc = func() ([]byte, error) { return bytes, nil } }
}

// WithResourceVersion enables the user to provide the resource version to be used.
// This version should be constant if the object has not changed to avoid needlessly sending resources to clients.
// By default it is computed by hashing the serialized version of the resource.
func WithResourceVersion(version string) CachedResourceOption {
	if version == "" {
		return func(*CachedResource) {}
	}
	return func(r *CachedResource) { r.computeResourceVersionFunc = func() (string, error) { return version, nil } }
}

// WithResourceTTL sets a TTL on the resource, that will be sent to the client with the payload.
func WithResourceTTL(ttl *time.Duration) CachedResourceOption {
	return func(r *CachedResource) { r.ttl = ttl }
}

func NewCachedResource(name, typeURL string, res Resource, opts ...CachedResourceOption) *CachedResource {
	cachedRes := &CachedResource{
		Name:     name,
		typeURL:  typeURL,
		resource: res,
	}
	for _, opt := range opts {
		opt(cachedRes)
	}
	if cachedRes.marshalFunc == nil {
		cachedRes.marshalFunc = sync.OnceValues(func() ([]byte, error) {
			return marshalResource(res)
		})
	}
	if cachedRes.computeResourceVersionFunc == nil {
		cachedRes.computeResourceVersionFunc = sync.OnceValues(func() (string, error) {
			marshaled, err := cachedRes.marshalFunc()
			if err != nil {
				return "", fmt.Errorf("marshaling resource: %w", err)
			}
			return hashResource(marshaled), nil
		})
	}
	return cachedRes
}

// SetCacheVersion updates the cache version. This violates the assumption that all fields can be safely read concurrently.
// It's required today for the linear cache constructor, where we are guaranteed resources are not being used concurently,
// and should not be used elsewhere.
func (c *CachedResource) SetCacheVersion(version string) {
	c.cacheVersion = version
}

// HasTTL returns whether the resource has a TTL set.
func (c *CachedResource) HasTTL() bool {
	return c.ttl != nil
}

// getMarshaledResource lazily marshals the resource and returns the bytes.
func (c *CachedResource) getMarshaledResource() ([]byte, error) {
	return c.marshalFunc()
}

// GetResourceVersion returns a stable version reflecting the resource content.
// By default it is built by hashing the serialized version of the object, using deterministic serializing.
func (c *CachedResource) GetResourceVersion() (string, error) {
	return c.computeResourceVersionFunc()
}

// GetVersion returns the version for the resource.
// By default it returns the cache version when the resource was added, but if requested it will return a
// version specific to the resource content.
func (c *CachedResource) GetVersion(useResourceVersion bool) (string, error) {
	if !useResourceVersion {
		return c.cacheVersion, nil
	}

	return c.GetResourceVersion()
}

// GetRawResource returns the underlying resource for use in legacy accessors.
func (c *CachedResource) GetRawResource() ResourceWithTTL {
	return ResourceWithTTL{
		Resource: c.resource,
		TTL:      c.ttl,
	}
}

var deltaResourceTypeURL = "type.googleapis.com/" + string(proto.MessageName(&discovery.Resource{}))

// getResourceVersion lazily hashes the resource and returns the stable hash used to track version changes.
func (c *CachedResource) GetSotwResource(isHeartbeat bool) (*anypb.Any, error) {
	buildResource := func() (*anypb.Any, error) {
		marshaled, err := c.getMarshaledResource()
		if err != nil {
			return nil, fmt.Errorf("marshaling: %w", err)
		}
		return &anypb.Any{
			TypeUrl: c.typeURL,
			Value:   marshaled,
		}, nil
	}

	if c.ttl == nil {
		return buildResource()
	}

	wrappedResource := &discovery.Resource{
		Name: c.Name,
		Ttl:  durationpb.New(*c.ttl),
	}

	if !isHeartbeat {
		rsrc, err := buildResource()
		if err != nil {
			return nil, err
		}
		wrappedResource.Resource = rsrc
	}

	marshaled, err := marshalResource(wrappedResource)
	if err != nil {
		return nil, fmt.Errorf("marshaling discovery resource: %w", err)
	}

	return &anypb.Any{
		TypeUrl: deltaResourceTypeURL,
		Value:   marshaled,
	}, nil
}

// getResourceVersion lazily hashes the resource and returns the stable hash used to track version changes.
func (c *CachedResource) GetDeltaResource() (*discovery.Resource, error) {
	marshaled, err := c.getMarshaledResource()
	if err != nil {
		return nil, fmt.Errorf("marshaling: %w", err)
	}
	version, err := c.GetResourceVersion()
	if err != nil {
		return nil, fmt.Errorf("computing version: %w", err)
	}
	return &discovery.Resource{
		Name: c.Name,
		Resource: &anypb.Any{
			TypeUrl: c.typeURL,
			Value:   marshaled,
		},
		Version: version,
	}, nil
}

// hashResource will take a resource and create a SHA256 hash sum out of the marshaled bytes
func hashResource(resource []byte) string {
	hasher := sha256.New()
	hasher.Write(resource)
	return hex.EncodeToString(hasher.Sum(nil))
}

// marshalResource converts the Resource to MarshaledResource.
func marshalResource(resource Resource) ([]byte, error) {
	return proto.MarshalOptions{Deterministic: true}.Marshal(resource)
}
