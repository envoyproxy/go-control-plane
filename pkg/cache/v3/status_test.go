// Copyright 2018 Envoyproxy Authors
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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
)

func TestIDHash(t *testing.T) {
	node := &core.Node{Id: "test"}
	got := (IDHash{}).ID(node)
	assert.Equalf(t, "test", got, "IDHash.ID(%v) => got %s, want %s", node, got, node.GetId())
	got = (IDHash{}).ID(nil)
	assert.Emptyf(t, got, "IDHash.ID(nil) => got %s, want empty", got)
}

func TestNewStatusInfo(t *testing.T) {
	node := &core.Node{Id: "test"}
	info := newStatusInfo(node)

	gotNode := info.GetNode()
	assert.Truef(t, reflect.DeepEqual(gotNode, node), "GetNode() => got %#v, want %#v", gotNode, node)

	gotNumWatches := info.GetNumWatches()
	assert.Equalf(t, 0, gotNumWatches, "GetNumWatches() => got %d, want 0", gotNumWatches)

	gotLastWatchRequestTime := info.GetLastWatchRequestTime()
	assert.Truef(t, gotLastWatchRequestTime.IsZero(), "GetLastWatchRequestTime() => got %v, want zero time", gotLastWatchRequestTime)

	gotNumDeltaWatches := info.GetNumDeltaWatches()
	assert.Equalf(t, 0, gotNumDeltaWatches, "GetNumDeltaWatches() => got %d, want 0", gotNumDeltaWatches)

	gotLastDeltaWatchRequestTime := info.GetLastDeltaWatchRequestTime()
	assert.Truef(t, gotLastDeltaWatchRequestTime.IsZero(), "GetLastDeltaWatchRequestTime() => got %v, want zero time", gotLastDeltaWatchRequestTime)
}
