package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
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
	Name string

	resource Resource
	ttl      *time.Duration

	// cacheVersion is the version of the cache at the time of last update, used in sotw.
	cacheVersion string

	marshalFunc func() (*anypb.Any, error)

	// onDemandOnly indicates if this resource is only sent when explicitly requested.
	// When false (default), resource is sent to wildcard subscriptions.
	// When true, resource is only sent when explicitly requested by name.
	onDemandOnly bool

	computeResourceVersionFunc func() (string, error)
}

type CachedResourceOption = func(*CachedResource)

// WithCacheVersion allows specifying the cacheVersion when the resource is set.
func WithCacheVersion(version string) CachedResourceOption {
	return func(r *CachedResource) { r.cacheVersion = version }
}

// WithMarshaledResource enables the user to provide the already marshaled resource if they have them.
// This serialization should strive at being consistent if the object has not changed (beware protobuf non-deterministic marshaling through anypb.New)
// or alternatively the resource version should also then be set.
// By default it is computed by performing a deterministic protobuf marshaling.
func WithMarshaledResource(res *anypb.Any) CachedResourceOption {
	if res == nil {
		return func(r *CachedResource) {
			r.marshalFunc = nil
		}
	}
	return func(r *CachedResource) {
		r.marshalFunc = func() (*anypb.Any, error) { return res, nil }
	}
}

// WithResourceVersion enables the user to provide the resource version to be used.
// This version should be constant if the object has not changed to avoid needlessly sending resources to clients.
// By default it is computed by hashing the serialized version of the resource.
func WithResourceVersion(version string) CachedResourceOption {
	if version == "" {
		return func(r *CachedResource) {
			r.computeResourceVersionFunc = nil
		}
	}
	return func(r *CachedResource) { r.computeResourceVersionFunc = func() (string, error) { return version, nil } }
}

// WithResourceTTL sets a TTL on the resource, that will be sent to the client with the payload.
func WithResourceTTL(ttl *time.Duration) CachedResourceOption {
	return func(r *CachedResource) { r.ttl = ttl }
}

// OnDemandOnly marks the resource as on-demand only.
func OnDemandOnly(onDemandOnly bool) CachedResourceOption {
	return func(r *CachedResource) { r.onDemandOnly = onDemandOnly }
}

func NewCachedResource(name string, res Resource, opts ...CachedResourceOption) *CachedResource {
	cachedRes := &CachedResource{
		Name:         name,
		resource:     res,
		onDemandOnly: false, // Default to wildcard-eligible
	}
	for _, opt := range opts {
		opt(cachedRes)
	}

	if cachedRes.marshalFunc == nil {
		cachedRes.marshalFunc = sync.OnceValues(func() (*anypb.Any, error) {
			return marshalResource(res)
		})
	}
	if cachedRes.computeResourceVersionFunc == nil {
		cachedRes.computeResourceVersionFunc = sync.OnceValues(func() (string, error) {
			marshaled, err := cachedRes.marshalFunc()
			if err != nil {
				return "", fmt.Errorf("marshaling resource: %w", err)
			}
			return hashResource(marshaled.Value), nil
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

// OnDemandOnly returns whether the resource should be ignored by wildcard watches if not explicitly requested.
func (c *CachedResource) OnDemandOnly() bool {
	return c.onDemandOnly
}

// getMarshaledResource lazily marshals the resource and returns the bytes.
func (c *CachedResource) getMarshaledResource() (*anypb.Any, error) {
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

// GetSotwResource returns the resource as is to be wrapped in a sotw response (i.e. DiscoveryResponse).
// If the response is a heartbeat the underlying resource is not included.
func (c *CachedResource) GetSotwResource(isHeartbeat bool) (*anypb.Any, error) {
	if isHeartbeat && c.ttl == nil {
		return nil, errors.New("heartbeat requested without ttl set")
	}

	if c.ttl == nil {
		// No TTL set, directly return the anypb format of the resource.
		return c.getMarshaledResource()
	}

	// A TTL is set, wrapped the resource into a discovery resource.
	wrappedResource := &discovery.Resource{
		Name: c.Name,
		Ttl:  durationpb.New(*c.ttl),
	}

	if !isHeartbeat {
		rsrc, err := c.getMarshaledResource()
		if err != nil {
			return nil, fmt.Errorf("marshaling resource: %w", err)
		}
		wrappedResource.Resource = rsrc
	}

	return marshalResource(wrappedResource)
}

// GetDeltaResource returns the resource as is to be wrapped in a delta response (i.e. DeltaDiscoveryResponse).
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
		Name:     c.Name,
		Resource: marshaled,
		Version:  version,
	}, nil
}

// hashResource will take a resource and create a SHA256 hash sum out of the marshaled bytes
func hashResource(resource []byte) string {
	hasher := sha256.New()
	hasher.Write(resource)
	return hex.EncodeToString(hasher.Sum(nil))
}

// marshalResource performs the same operation as anypb.New but using deterministic marshaling.
func marshalResource(resource Resource) (*anypb.Any, error) {
	ret := new(anypb.Any)
	err := anypb.MarshalFrom(ret, resource, proto.MarshalOptions{Deterministic: true})
	return ret, err
}
