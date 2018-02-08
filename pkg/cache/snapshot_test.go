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

package cache_test

import (
	"testing"

	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource"
)

func TestSnapshotConsistent(t *testing.T) {
	if err := snapshot.Consistent(); err != nil {
		t.Errorf("got inconsistent snapshot for %#v", snapshot)
	}
	if snap := cache.NewSnapshot(version, []cache.Resource{endpoint}, nil, nil, nil); snap.Consistent() == nil {
		t.Errorf("got consistent snapshot %#v", snap)
	}
	if snap := cache.NewSnapshot(version, []cache.Resource{resource.MakeEndpoint("missing", 8080)},
		[]cache.Resource{cluster}, nil, nil); snap.Consistent() == nil {
		t.Errorf("got consistent snapshot %#v", snap)
	}
	if snap := cache.NewSnapshot(version, nil, nil, nil, []cache.Resource{listener}); snap.Consistent() == nil {
		t.Errorf("got consistent snapshot %#v", snap)
	}
	if snap := cache.NewSnapshot(version, nil, nil,
		[]cache.Resource{resource.MakeRoute("test", clusterName)}, []cache.Resource{listener}); snap.Consistent() == nil {
		t.Errorf("got consistent snapshot %#v", snap)
	}
}
