// Copyright 2020 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package cache

import (
	"context"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/log"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

type ResourceContainer interface {
	GetResources(typeURL string) map[string]types.Resource

	GetVersionMap() map[string]map[string]string

	GetVersion(typeURL string) string
}

// This struct exist only because we want to reuse createDeltaResponse function from linear cache code
// so we create a wrapper, which adopts LinearCche to satisfy ResourceContainer interface
type cacheWrapper struct {
	cache *LinearCache
}

func (w *cacheWrapper) GetResources(typeURL string) map[string]types.Resource {
	return w.cache.resources
}

func (w *cacheWrapper) GetVersionMap() map[string]map[string]string {
	return map[string]map[string]string{w.cache.typeURL: w.cache.versionMap}
}

func (w *cacheWrapper) GetVersion(typeURL string) string {
	return w.cache.getVersion()
}

// Respond to a delta watch with the provided snapshot value. If the response is nil, there has been no state change.
func respondDeltaSnapshot(ctx context.Context, request *DeltaRequest, value chan DeltaResponse, state stream.StreamState, snapshot *Snapshot, log log.Logger) (*RawDeltaResponse, error) {
	resp := createDeltaResponse(ctx, request, state, snapshot, log)

	// Only send a response if there were changes
	// We want to respond immediatly for the first wildcard request in a stream, even if the response is empty
	// otherwise, envoy won't complete initialization
	if len(resp.Resources) > 0 || len(resp.RemovedResources) > 0 || (state.IsWildcard() && state.IsFirst()) {
		if log != nil {
			log.Debugf("node: %s, sending delta response with resources: %v removed resources %v wildcard: %t",
				request.GetNode().GetId(), resp.Resources, resp.RemovedResources, state.IsWildcard())
		}
		select {
		case value <- resp:
			return resp, nil
		case <-ctx.Done():
			return resp, context.Canceled
		}
	}
	return nil, nil
}

func respondDeltaLinear(request *DeltaRequest, value chan DeltaResponse, state stream.StreamState, cache *LinearCache, log log.Logger) *RawDeltaResponse {
	resp := createDeltaResponse(context.Background(), request, state, &cacheWrapper{cache}, log)

	// Only send a response if there were changes
	if len(resp.Resources) > 0 || len(resp.RemovedResources) > 0 {
		if log != nil {
			log.Debugf("node: %s, sending delta response with resources: %v removed resources %v wildcard: %t",
				request.GetNode().GetId(), resp.Resources, resp.RemovedResources, state.IsWildcard())
		}
		value <- resp
		return resp
	}
	return nil
}

func createDeltaResponse(ctx context.Context, req *DeltaRequest, state stream.StreamState, snapshot ResourceContainer, log log.Logger) *RawDeltaResponse {
	resources := snapshot.GetResources((req.TypeUrl))

	// variables to build our response with
	nextVersionMap := make(map[string]string)
	filtered := make([]types.Resource, 0, len(resources))
	toRemove := make([]string, 0)

	// If we are handling a wildcard request, we want to respond with all resources
	switch {
	case state.IsWildcard():
		for name, r := range resources {
			// Since we've already precomputed the version hashes of the new snapshot,
			// we can just set it here to be used for comparison later
			version := snapshot.GetVersionMap()[req.TypeUrl][name]
			nextVersionMap[name] = version
			prevVersion, found := state.GetResourceVersions()[name]
			if !found || (prevVersion != nextVersionMap[name]) {
				filtered = append(filtered, r)
			}
		}
	default:
		// Reply only with the requested resources
		for name, prevVersion := range state.GetResourceVersions() {
			if r, ok := resources[name]; ok {
				nextVersion := snapshot.GetVersionMap()[req.TypeUrl][name]
				if prevVersion != nextVersion {
					filtered = append(filtered, r)
				}
				nextVersionMap[name] = nextVersion
			}
		}
	}

	// Compute resources for removal regardless of the request type
	for name := range state.GetResourceVersions() {
		if _, ok := resources[name]; !ok {
			toRemove = append(toRemove, name)
		}
	}

	return &RawDeltaResponse{
		DeltaRequest:      req,
		Resources:         filtered,
		RemovedResources:  toRemove,
		NextVersionMap:    nextVersionMap,
		SystemVersionInfo: snapshot.GetVersion(req.TypeUrl),
		Ctx:               ctx,
	}
}
