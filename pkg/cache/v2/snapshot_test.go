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

	common "github.com/envoyproxy/go-control-plane/pkg/cache/common"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource/v2"
	utils "github.com/envoyproxy/go-control-plane/pkg/utils/v2"
)

func TestSnapshotConsistent(t *testing.T) {
	if err := snapshot.Consistent(); err != nil {
		t.Errorf("got inconsistent snapshot for %#v", snapshot)
	}
	if snap := cache.NewSnapshot(version, []common.Resource{testEndpoint}, nil, nil, nil, nil); snap.Consistent() == nil {
		t.Errorf("got consistent snapshot %#v", snap)
	}
	if snap := cache.NewSnapshot(version, []common.Resource{resource.MakeEndpoint("missing", 8080)},
		[]common.Resource{testCluster}, nil, nil, nil); snap.Consistent() == nil {
		t.Errorf("got consistent snapshot %#v", snap)
	}
	if snap := cache.NewSnapshot(version, nil, nil, nil, []common.Resource{testListener}, nil); snap.Consistent() == nil {
		t.Errorf("got consistent snapshot %#v", snap)
	}
	if snap := cache.NewSnapshot(version, nil, nil,
		[]common.Resource{resource.MakeRoute("test", clusterName)}, []common.Resource{testListener}, nil); snap.Consistent() == nil {
		t.Errorf("got consistent snapshot %#v", snap)
	}
}

func TestSnapshotGetters(t *testing.T) {
	var nilsnap *cache.Snapshot
	if out := nilsnap.GetResources(utils.EndpointType); out != nil {
		t.Errorf("got non-empty resources for nil snapshot: %#v", out)
	}
	if out := nilsnap.Consistent(); out == nil {
		t.Errorf("nil snapshot should be inconsistent")
	}
	if out := nilsnap.GetVersion(utils.EndpointType); out != "" {
		t.Errorf("got non-empty version for nil snapshot: %#v", out)
	}
	if out := snapshot.GetResources("not a type"); out != nil {
		t.Errorf("got non-empty resources for unknown type: %#v", out)
	}
	if out := snapshot.GetVersion("not a type"); out != "" {
		t.Errorf("got non-empty version for unknown type: %#v", out)
	}
}
