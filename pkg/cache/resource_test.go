// Copyright 2017 Envoyproxy Authors
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
	"testing"

	"github.com/envoyproxy/go-control-plane/pkg/test/resource"
)

const (
	clusterName  = "cluster0"
	routeName    = "route0"
	listenerName = "listener0"
)

var (
	endpoint = resource.MakeEndpoint(clusterName, 8080)
	cluster  = resource.MakeCluster(true, clusterName)
	route    = resource.MakeRoute(routeName, clusterName)
	listener = resource.MakeListener(true, listenerName, 80, routeName)
)

func TestGetResourceName(t *testing.T) {
	if name := GetResourceName(endpoint); name != clusterName {
		t.Errorf("GetResourceName(%v) => got %q, want %q", endpoint, name, clusterName)
	}
	if name := GetResourceName(cluster); name != clusterName {
		t.Errorf("GetResourceName(%v) => got %q, want %q", cluster, name, clusterName)
	}
	if name := GetResourceName(route); name != routeName {
		t.Errorf("GetResourceName(%v) => got %q, want %q", route, name, routeName)
	}
	if name := GetResourceName(listener); name != listenerName {
		t.Errorf("GetResourceName(%v) => got %q, want %q", listener, name, listenerName)
	}
	if name := GetResourceName(nil); name != "" {
		t.Errorf("GetResourceName(nil) => got %q, want none", name)
	}
}
