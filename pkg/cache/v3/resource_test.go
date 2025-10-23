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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource/v3"
)

const (
	clusterName         = "cluster0"
	routeName           = "route0"
	embeddedRouteName   = "embeddedRoute0"
	scopedRouteName     = "scopedRoute0"
	listenerName        = "listener0"
	scopedListenerName  = "scopedListener0"
	virtualHostName     = "virtualHost0"
	runtimeName         = "runtime0"
	tlsName             = "secret0"
	rootName            = "root0"
	extensionConfigName = "extensionConfig0"
)

var (
	testEndpoint        = resource.MakeEndpoint(clusterName, 8080)
	testCluster         = resource.MakeCluster(resource.Ads, clusterName)
	testRoute           = resource.MakeRouteConfig(routeName, clusterName)
	testEmbeddedRoute   = resource.MakeRouteConfig(embeddedRouteName, clusterName)
	testScopedRoute     = resource.MakeScopedRouteConfig(scopedRouteName, routeName, []string{"1.2.3.4"})
	testVirtualHost     = resource.MakeVirtualHost(virtualHostName, clusterName)
	testListener        = resource.MakeRouteHTTPListener(resource.Ads, listenerName, 80, routeName)
	testListenerDefault = resource.MakeRouteHTTPListenerDefaultFilterChain(resource.Ads, listenerName, 80, routeName)
	testScopedListener  = resource.MakeScopedRouteHTTPListenerForRoute(resource.Ads, scopedListenerName, 80, embeddedRouteName)
	testRuntime         = resource.MakeRuntime(runtimeName)
	testSecret          = resource.MakeSecrets(tlsName, rootName)
	testExtensionConfig = resource.MakeExtensionConfig(resource.Ads, extensionConfigName, routeName)
)

func TestValidate(t *testing.T) {
	require.NoError(t, testEndpoint.Validate())
	require.NoError(t, testCluster.Validate())
	require.NoError(t, testRoute.Validate())
	require.NoError(t, testScopedRoute.Validate())
	require.NoError(t, testVirtualHost.Validate())
	require.NoError(t, testListener.Validate())
	require.NoError(t, testScopedListener.Validate())
	require.NoError(t, testRuntime.Validate())
	require.NoError(t, testExtensionConfig.Validate())

	invalidRoute := &route.RouteConfiguration{
		Name: "test",
		VirtualHosts: []*route.VirtualHost{{
			Name:    "test",
			Domains: []string{},
		}},
	}

	err := invalidRoute.Validate()
	require.Errorf(t, err, "expected an error")
	err = invalidRoute.GetVirtualHosts()[0].Validate()
	assert.Errorf(t, err, "expected an error")
}

type customResource struct {
	cluster.Filter // Any proto would work here.
}

const customName = "test-name"

func (cs *customResource) GetName() string { return customName }

var _ types.ResourceWithName = &customResource{}

func TestGetResourceName(t *testing.T) {
	name := cache.GetResourceName(testEndpoint)
	assert.Equalf(t, clusterName, name, "GetResourceName(%v) => got %q, want %q", testEndpoint, name, clusterName)
	name = cache.GetResourceName(testCluster)
	assert.Equalf(t, clusterName, name, "GetResourceName(%v) => got %q, want %q", testCluster, name, clusterName)
	name = cache.GetResourceName(testRoute)
	assert.Equalf(t, routeName, name, "GetResourceName(%v) => got %q, want %q", testRoute, name, routeName)
	name = cache.GetResourceName(testScopedRoute)
	assert.Equalf(t, scopedRouteName, name, "GetResourceName(%v) => got %q, want %q", testScopedRoute, name, scopedRouteName)
	name = cache.GetResourceName(testVirtualHost)
	assert.Equalf(t, virtualHostName, name, "GetResourceName(%v) => got %q, want %q", testVirtualHost, name, virtualHostName)
	name = cache.GetResourceName(testListener)
	assert.Equalf(t, listenerName, name, "GetResourceName(%v) => got %q, want %q", testListener, name, listenerName)
	name = cache.GetResourceName(testRuntime)
	assert.Equalf(t, runtimeName, name, "GetResourceName(%v) => got %q, want %q", testRuntime, name, runtimeName)
	name = cache.GetResourceName(&customResource{})
	assert.Equalf(t, customName, name, "GetResourceName(nil) => got %q, want %q", name, customName)
	name = cache.GetResourceName(nil)
	assert.Emptyf(t, name, "GetResourceName(nil) => got %q, want none", name)
}

func TestGetResourceNames(t *testing.T) {
	tests := []struct {
		name  string
		input []types.ResourceWithTTL
		want  []string
	}{
		{
			name:  "empty",
			input: []types.ResourceWithTTL{},
			want:  []string{},
		},
		{
			name:  "many",
			input: []types.ResourceWithTTL{{Resource: testRuntime}, {Resource: testListener}, {Resource: testListenerDefault}, {Resource: testVirtualHost}},
			want:  []string{runtimeName, listenerName, listenerName, virtualHostName},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got := cache.GetResourceNames(test.input)
			assert.ElementsMatch(t, test.want, got)
		})
	}
}

func TestGetResourceReferences(t *testing.T) {
	cases := []struct {
		in  types.Resource
		out map[rsrc.Type]map[string]bool
	}{
		{
			in:  nil,
			out: map[rsrc.Type]map[string]bool{},
		},
		{
			in:  testCluster,
			out: map[rsrc.Type]map[string]bool{rsrc.EndpointType: {clusterName: true}},
		},
		{
			in: &cluster.Cluster{
				Name: clusterName, ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_EDS},
				EdsClusterConfig: &cluster.Cluster_EdsClusterConfig{ServiceName: "test"},
			},
			out: map[rsrc.Type]map[string]bool{rsrc.EndpointType: {"test": true}},
		},
		{
			in:  resource.MakeScopedRouteHTTPListener(resource.Xds, listenerName, 80),
			out: map[rsrc.Type]map[string]bool{},
		},
		{
			in:  resource.MakeScopedRouteHTTPListenerForRoute(resource.Xds, listenerName, 80, routeName),
			out: map[rsrc.Type]map[string]bool{rsrc.RouteType: {routeName: true}},
		},
		{
			in:  resource.MakeRouteHTTPListener(resource.Ads, listenerName, 80, routeName),
			out: map[rsrc.Type]map[string]bool{rsrc.RouteType: {routeName: true}},
		},
		{
			in:  resource.MakeRouteHTTPListenerDefaultFilterChain(resource.Ads, listenerName, 80, routeName),
			out: map[rsrc.Type]map[string]bool{rsrc.RouteType: {routeName: true}},
		},
		{
			in:  resource.MakeTCPListener(listenerName, 80, clusterName),
			out: map[rsrc.Type]map[string]bool{},
		},
		{
			in:  testRoute,
			out: map[rsrc.Type]map[string]bool{},
		},
		{
			in:  resource.MakeVHDSRouteConfig(resource.Ads, routeName),
			out: map[rsrc.Type]map[string]bool{},
		},
		{
			in:  testScopedRoute,
			out: map[rsrc.Type]map[string]bool{rsrc.RouteType: {routeName: true}},
		},
		{
			in:  testVirtualHost,
			out: map[rsrc.Type]map[string]bool{},
		},
		{
			in:  testEndpoint,
			out: map[rsrc.Type]map[string]bool{},
		},
		{
			in:  testRuntime,
			out: map[rsrc.Type]map[string]bool{},
		},
	}
	for _, cs := range cases {
		names := cache.GetResourceReferences(cache.IndexResourcesByName([]types.ResourceWithTTL{{Resource: cs.in}}))
		assert.Truef(t, reflect.DeepEqual(names, cs.out), "GetResourceReferences(%v) => got %v, want %v", cs.in, names, cs.out)
	}
}

func TestGetAllResourceReferencesReturnsExpectedRefs(t *testing.T) {
	expected := map[rsrc.Type]map[string]bool{
		rsrc.RouteType:    {routeName: true, embeddedRouteName: true},
		rsrc.EndpointType: {clusterName: true},
	}

	resources := [types.UnknownType]cache.Resources{}
	resources[types.Endpoint] = cache.NewResources("1", []types.Resource{testEndpoint})
	resources[types.Cluster] = cache.NewResources("1", []types.Resource{testCluster})
	resources[types.Route] = cache.NewResources("1", []types.Resource{testRoute})
	resources[types.Listener] = cache.NewResources("1", []types.Resource{testListener, testScopedListener})
	resources[types.ScopedRoute] = cache.NewResources("1", []types.Resource{testScopedRoute})
	actual := cache.GetAllResourceReferences(resources)

	assert.Truef(t, reflect.DeepEqual(actual, expected), "GetAllResourceReferences(%v) => got %v, want %v", resources, actual, expected)
}
