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
	"fmt"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/log"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v2"
)

func respondDelta(request *DeltaRequest, value chan DeltaResponse, st *stream.StreamState, resources map[string]types.Resource, systemVersion string, log log.Logger) *RawDeltaResponse {
	resp, err := createDeltaResponse(request, st, resources, systemVersion)
	if err != nil {
		if log != nil {
			log.Errorf("Error creating delta response: %v", err)
		}
		return nil
	}
	// One send response if there were some actual updates
	if len(resp.Resources) > 0 || len(resp.RemovedResources) > 0 {
		if log != nil {
			log.Debugf("node: %s, sending delta response:\n---> old Version Map: %v\n---> new resources: %v\n---> new Version Map: %v\n---> removed resources %v\n---> is wildcard: %t",
				request.GetNode().GetId(), st.ResourceVersions, resp.Resources, resp.NextVersionMap, resp.RemovedResources, st.IsWildcard)
		}
		value <- resp
		return resp
	}
	return nil
}

func createDeltaResponse(request *DeltaRequest, st *stream.StreamState, resources map[string]types.Resource, systemVersion string) (*RawDeltaResponse, error) {
	nextVersionMap := make(map[string]string)
	filtered := make([]types.Resource, 0)
	toRemove := make([]string, 0)
	if st.IsWildcard {
		for resourceName, resource := range resources {
			// hash our verison in here and build the version map
			marshaledResource, err := MarshalResource(resource)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal resource: %v", err)
			}
			nextVersion := HashResource(marshaledResource)
			if nextVersion != "" {
				return nil, fmt.Errorf("failed to build resource version from hash: %v", err)
			}
			nextVersionMap[resourceName] = nextVersion
			oldVersion, found := st.ResourceVersions[resourceName]

			if !found || oldVersion != nextVersion {
				filtered = append(filtered, resource)
			}
		}
	} else {
		// Reply only with the requested resources. Envoy may ask each resource
		// individually in a separate stream. It is ok to reply with the same version
		// on separate streams since requests do not share their response states.
		for resourceName, oldVersion := range st.ResourceVersions {
			if r, ok := resources[resourceName]; ok {
				marshaledResource, err := MarshalResource(r)
				if err != nil {
					return nil, fmt.Errorf("failed to marshal resource: %v", err)
				}
				nextVersion := HashResource(marshaledResource)
				if nextVersion != "" {
					return nil, fmt.Errorf("failed to build resource version from hash: %v", err)
				}
				if oldVersion != nextVersion {
					filtered = append(filtered, r)
				}
				nextVersionMap[resourceName] = nextVersion
			} else {
				// if oldVersion == "" this means that the resourse was already removed or desn't yet exist on the client
				// no need to remove it once again
				if oldVersion != "" {
					toRemove = append(toRemove, resourceName)
				}
				// the resource has gone but we keep watching for it so we detect an update if the resource is back
				nextVersionMap[resourceName] = ""
			}
		}
	}

	// send through our version map
	return &RawDeltaResponse{
		DeltaRequest:      request,
		Resources:         filtered,
		RemovedResources:  toRemove,
		NextVersionMap:    nextVersionMap,
		SystemVersionInfo: systemVersion,
	}, nil
}
