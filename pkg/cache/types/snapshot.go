package types //nolint:revive // var-naming: avoid meaningless package names

import (
	"fmt"
	"time"

	"github.com/envoyproxy/go-control-plane/pkg/cache/internal"

	"google.golang.org/protobuf/types/known/anypb"
)

// SnapshotResource represents a resource to be provided to caches.
type SnapshotResource struct {
	// Mandatory
	Name string
	// Optional
	TTL *time.Duration

	// Raw protobuf message. If unset, calls to `GetRawResource` methods on snapshot are undefined behavior.
	// One of Resource or Serialized must be set.
	Resource Resource
	// Serialized protobuf message. Allows the user to provide a totally opaque content.
	// For control-plane and data-plane efficiency, it is critical for the serialization to be consisten
	// (e.g. beware the anypb default marshaling), or the version below should be provided.
	// One of Resource or Serialized must be set. If both are set only Serialized will be sent to the data-plane.
	Serialized *anypb.Any

	// Resource-intrisic version. This version MUST change if the underlying resource changes to be sent to the data-plane.
	// Optional, if not set a version will be computed from the marshaled representation of the resource.
	Version string

	// Optional - marks resource as on-demand only (e.g. for OdCDS).
	// Only supported for the snapshot cache.
	// When false (default), this resource is sent to all clients with wildcard subscriptions.
	// When true, this resource is only sent when explicitly requested by name.
	OnDemandOnly bool
}

func (r *SnapshotResource) AsCachedResource(cacheVersion string) *internal.CachedResource {
	return internal.NewCachedResource(r.Name, r.Resource,
		internal.WithCacheVersion(cacheVersion),
		internal.WithResourceTTL(r.TTL),
		internal.WithResourceVersion(r.Version),
		internal.WithMarshaledResource(r.Serialized),
		internal.OnDemandOnly(r.OnDemandOnly),
	)
}

// From anypb code.
const urlPrefix = "type.googleapis.com/"

// TypeURL returns the type of the included resources
func (r *SnapshotResource) TypeURL() string {
	if r.Serialized != nil {
		return r.Serialized.GetTypeUrl()
	}
	if r.Resource == nil {
		return ""
	}
	return urlPrefix + string(r.Resource.ProtoReflect().Descriptor().FullName())
}

// TypeSnapshot represents the resources for a given type, associated with an opaque version.
// The snapshot (and associated resources) must not be modified once provided to the cache.
// A given TypeSnapshot instance can be provided in multiple snapshots, including for multiple node ids, as long as it remains immutable.
type TypeSnapshot struct {
	typeURL   string
	version   string
	resources map[string]*internal.CachedResource
}

func NewTypeSnapshot(typeURL, version string, resources []SnapshotResource) (TypeSnapshot, error) {
	s := TypeSnapshot{
		typeURL:   typeURL,
		version:   version,
		resources: make(map[string]*internal.CachedResource, len(resources)),
	}
	for _, res := range resources {
		if res.TypeURL() != typeURL {
			return TypeSnapshot{}, fmt.Errorf("resource %s has wrong type: expected %s and received %s", res.Name, typeURL, res.Serialized.GetTypeUrl())
		}
		s.resources[res.Name] = res.AsCachedResource(version)
	}
	return s, nil
}

// GetVersion returns the version of the snapshot.
// Multiple TypeSnapshots in a given Snapshot can have different versions.
func (s TypeSnapshot) GetVersion() string {
	return s.version
}

// GetResources returns the resources in the snapshot.
// The map must not be modified by the caller.
func (s TypeSnapshot) GetResources() map[string]*internal.CachedResource {
	return s.resources
}

// Snapshot represents a consistent set of resources for multiple types.
// Once provided to a cache it should not be altered in any way.
type Snapshot struct {
	// defaultVersion is the negative version returned if there is no snapshot set for the provided type.
	defaultVersion string
	resources      map[string]TypeSnapshot
}

// NewSnapshot creates a snapshot with a single version for all resource types.
func NewSnapshot(version string, resources map[string][]SnapshotResource) (*Snapshot, error) {
	s := &Snapshot{
		defaultVersion: version,
		resources:      make(map[string]TypeSnapshot, len(resources)),
	}
	for typeURL, res := range resources {
		snap, err := NewTypeSnapshot(typeURL, version, res)
		if err != nil {
			return nil, fmt.Errorf("building snapshot for type %s: %w", typeURL, err)
		}
		s.resources[typeURL] = snap
	}
	return s, nil
}

// NewSnapshotFromTypeSnapshots creates a snapshot from per-type snapshots.
// TypeSnapshot instances can be shared across snapshots, but must not be altered in any way once provided to at least one snapshot.
func NewSnapshotFromTypeSnapshots(version string, snapshots []TypeSnapshot) (*Snapshot, error) {
	s := &Snapshot{
		defaultVersion: version,
		resources:      make(map[string]TypeSnapshot, len(snapshots)),
	}
	for _, snap := range snapshots {
		s.resources[snap.typeURL] = snap
	}
	return s, nil
}

// GetVersion returns the current version of the resource indicated by typeURL.
// The version string that is returned is opaque and should only be compared for equality.
func (s *Snapshot) GetVersion(typeURL string) string {
	typeSnapshot, ok := s.resources[typeURL]
	if !ok {
		return s.defaultVersion
	}
	return typeSnapshot.version
}

func (s *Snapshot) GetTypeSnapshot(typeURL string) TypeSnapshot {
	return s.resources[typeURL]
}
