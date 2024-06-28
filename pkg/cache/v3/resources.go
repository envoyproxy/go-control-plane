package cache

import (
	"fmt"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
)

// Resources is a versioned group of resources.
type Resources struct {
	// Version information.
	Version string

	// Items in the group indexed by name.
	Items map[string]types.ResourceWithTTL
}

// IndexResourcesByName creates a map from the resource name to the resource.
func IndexResourcesByName(items []types.ResourceWithTTL) (map[string]types.ResourceWithTTL, error) {
	indexed := make(map[string]types.ResourceWithTTL, len(items))
	for _, item := range items {
		name := GetResourceName(item.Resource)
		if _, seen := indexed[name]; seen {
			return nil, fmt.Errorf("duplicate name: %q", name)
		}
		indexed[name] = item
	}
	return indexed, nil
}

// IndexRawResourcesByName creates a map from the resource name to the resource.
func IndexRawResourcesByName(items []types.Resource) (map[string]types.Resource, error) {
	indexed := make(map[string]types.Resource, len(items))
	for _, item := range items {
		name := GetResourceName(item)
		if _, seen := indexed[name]; seen {
			return nil, fmt.Errorf("duplicate name: %q", name)
		}

		indexed[name] = item
	}
	return indexed, nil
}

// NewResources creates a new resource group.
func NewResources(version string, items []types.Resource) (Resources, error) {
	itemsWithTTL := make([]types.ResourceWithTTL, 0, len(items))
	for _, item := range items {
		itemsWithTTL = append(itemsWithTTL, types.ResourceWithTTL{Resource: item})
	}
	return NewResourcesWithTTL(version, itemsWithTTL)
}

// NewResourcesWithTTL creates a new resource group.
func NewResourcesWithTTL(version string, items []types.ResourceWithTTL) (Resources, error) {
	nameIndexedItems, err := IndexResourcesByName(items)
	if err != nil {
		return Resources{}, fmt.Errorf("indexing resource by name: %w", err)
	}
	return Resources{
		Version: version,
		Items:   nameIndexedItems,
	}, nil
}
