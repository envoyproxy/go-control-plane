package types

import (
	"time"

	"github.com/envoyproxy/go-control-plane/pkg/cache/internal/resources"

	"google.golang.org/protobuf/types/known/anypb"
)

// SnapshotResource represents a resource to be provided to caches.
type SnapshotResource struct {
	// Mandatory
	Name string
	// Mandatory
	Resource Resource
	// Optional
	TTL *time.Duration

	// Optional
	Serialized *anypb.Any
	// Optional
	Version string
}

func (r *SnapshotResource) asCachedResource(typeURL, cacheVersion string) *resources.CachedResource {
	var serialized []byte
	if r.Serialized != nil {
		serialized = r.Serialized.Value
	}
	return resources.NewCachedResource(r.Name, typeURL, r.Resource,
		resources.WithCacheVersion(cacheVersion),
		resources.WithTTL(r.TTL),
		resources.WithVersion(r.Version),
		resources.WithMarshaled(serialized))
}

// TypeSnapshot represents the resources for a given type, associated with an opaque version.
// The snapshot (and associated resources) must not be modified once provided to the cache.
// A given TypeSnapshot instance can be provided in multiple snapshots, including for multiple node ids, as long as it remains immutable.
type TypeSnapshot struct {
	typeURL   string
	version   string
	resources map[string]*resources.CachedResource
}

func NewTypeSnapshot(typeURL, version string, res []SnapshotResource) TypeSnapshot {
	s := TypeSnapshot{
		typeURL:   typeURL,
		version:   version,
		resources: make(map[string]*resources.CachedResource, len(res)),
	}
	for _, r := range res {
		s.resources[r.Name] = r.asCachedResource(typeURL, version)
	}
	return s
}

// GetVersion returns the version of the snapshot.
// Multiple TypeSnapshots in a given Snapshot can have different versions.
func (s TypeSnapshot) GetVersion() string {
	return s.version
}

// GetResources returns the resources in the snapshot.
// The map must not be modified by the caller.
func (s TypeSnapshot) GetResources() map[string]*resources.CachedResource {
	return s.resources
}

// NodeSnapshot represents a consistent set of resources for multiple types.
// Once provided to a cache it should not be altered in any way.
type NodeSnapshot struct {
	// defaultVersion is the negative version returned if there is no snapshot set for the provided type.
	defaultVersion string
	resources      map[string]TypeSnapshot
}

// NewNodeSnapshot creates a snapshot with a single version for all resource types.
func NewNodeSnapshot(version string, resources map[string][]SnapshotResource) (*NodeSnapshot, error) {
	s := &NodeSnapshot{
		defaultVersion: version,
		resources:      make(map[string]TypeSnapshot, len(resources)),
	}
	for typeURL, res := range resources {
		s.resources[typeURL] = NewTypeSnapshot(typeURL, version, res)
	}
	return s, nil
}

// NewNodeSnapshotFromTypeSnapshots creates a snapshot from per-type snapshots.
// TypeSnapshot instances can be shared across snapshots, but must not be altered in any way once provided to at least one snapshot.
func NewNodeSnapshotFromTypeSnapshots(version string, snapshots []TypeSnapshot) (*NodeSnapshot, error) {
	s := &NodeSnapshot{
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
func (s *NodeSnapshot) GetVersion(typeURL string) string {
	typeSnapshot, ok := s.resources[typeURL]
	if !ok {
		return s.defaultVersion
	}
	return typeSnapshot.version
}

func (s *NodeSnapshot) GetTypeSnapshot(typeURL string) TypeSnapshot {
	return s.resources[typeURL]
}
