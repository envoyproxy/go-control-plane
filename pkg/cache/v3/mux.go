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
	"errors"

	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

// MuxCache multiplexes across several caches using a classification function.
// If there is no matching cache for a classification result, the cache
// responds with an empty closed channel, which effectively terminates the
// stream on the server. It might be preferred to respond with a "nil" channel
// instead which will leave the stream open in case the stream is aggregated by
// making sure there is always a matching cache.
type MuxCache struct {
	// Classification functions.
	Classify func(Request) string
	// Muxed caches.
	Caches map[string]Cache
}

var _ Cache = &MuxCache{}

func (mux *MuxCache) CreateWatch(request *Request, value chan Response) func() {
	// Passing a Request by value to Classify triggers a govet copylocks
	// error. This is because protobuf messages specifically embed
	// a Mutex in order to trigger this check (i.e. not for actually
	// locking fields).
	//
	// So in this specific case, it happens to be safe to copy the
	// lock. Fixing this would require adding a new API to classify
	// requests, and deprecating Classify.
	//
	// nolint:govet
	key := mux.Classify(*request)
	cache, exists := mux.Caches[key]
	if !exists {
		value <- nil
		return nil
	}
	return cache.CreateWatch(request, value)
}

func (mux *MuxCache) CreateDeltaWatch(request *DeltaRequest, state stream.StreamState, value chan DeltaResponse) func() {
	return nil
}

func (mux *MuxCache) Fetch(ctx context.Context, request *Request) (Response, error) {
	return nil, errors.New("not implemented")
}
