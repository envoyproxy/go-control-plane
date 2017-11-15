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

package server

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/envoyproxy/go-control-plane/api"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/envoyproxy/go-control-plane/pkg/test/resource"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type mockConfigWatcher struct {
	counts     map[cache.ResponseType]int
	responses  map[cache.ResponseType][]cache.Response
	closeWatch bool
}

func (config *mockConfigWatcher) Watch(typ cache.ResponseType, _ *api.Node, version string, names []string) cache.Watch {
	config.counts[typ] = config.counts[typ] + 1
	out := make(chan cache.Response, 1)
	if len(config.responses[typ]) > 0 {
		out <- config.responses[typ][0]
		config.responses[typ] = config.responses[typ][1:]
	} else if config.closeWatch {
		close(out)
	}
	return cache.Watch{
		Value: out,
		Type:  typ,
		Names: names,
	}
}

func makeMockConfigWatcher() *mockConfigWatcher {
	return &mockConfigWatcher{
		counts: make(map[cache.ResponseType]int),
	}
}

type mockStream struct {
	t         *testing.T
	recv      chan *api.DiscoveryRequest
	sent      chan *api.DiscoveryResponse
	nonce     int
	sendError bool
	grpc.ServerStream
}

func (stream *mockStream) Send(resp *api.DiscoveryResponse) error {
	// check that nonce is monotonically incrementing
	stream.nonce = stream.nonce + 1
	if resp.Nonce != fmt.Sprintf("%d", stream.nonce) {
		stream.t.Errorf("Nonce => got %q, want %d", resp.Nonce, stream.nonce)
	}
	// check that version is set
	if resp.VersionInfo == "" {
		stream.t.Error("VersionInfo => got none, want non-empty")
	}
	// check resources are non-empty
	if len(resp.Resources) == 0 {
		stream.t.Error("Resources => got none, want non-empty")
	}
	// check that type URL matches in resources
	if resp.TypeUrl == "" {
		stream.t.Error("TypeUrl => got none, want non-empty")
	}
	for _, res := range resp.Resources {
		if res.TypeUrl != resp.TypeUrl {
			stream.t.Errorf("TypeUrl => got %q, want %q", res.TypeUrl, resp.TypeUrl)
		}
	}
	stream.sent <- resp
	if stream.sendError {
		return errors.New("send error")
	}
	return nil
}

func (stream *mockStream) Recv() (*api.DiscoveryRequest, error) {
	req, more := <-stream.recv
	if !more {
		return nil, errors.New("empty")
	}
	return req, nil
}

func makeMockStream(t *testing.T) *mockStream {
	return &mockStream{
		t:    t,
		sent: make(chan *api.DiscoveryResponse, 10),
		recv: make(chan *api.DiscoveryRequest, 10),
	}
}

const (
	clusterName  = "cluster0"
	routeName    = "route0"
	listenerName = "listener0"
)

var (
	node = &api.Node{
		Id:      "test-id",
		Cluster: "test-cluster",
	}
	endpoint = resource.MakeEndpoint(clusterName, 8080)
	cluster  = resource.MakeCluster(true, clusterName)
	route    = resource.MakeRoute(routeName, clusterName)
	listener = resource.MakeListener(true, listenerName, 80, routeName)
)

func makeResponses() map[cache.ResponseType][]cache.Response {
	return map[cache.ResponseType][]cache.Response{
		cache.EndpointResponse: []cache.Response{{
			Version:   "1",
			Resources: []proto.Message{endpoint},
		}},
		cache.ClusterResponse: []cache.Response{{
			Version:   "1",
			Resources: []proto.Message{cluster},
		}},
		cache.RouteResponse: []cache.Response{{
			Version:   "1",
			Resources: []proto.Message{route},
		}},
		cache.ListenerResponse: []cache.Response{{
			Version:   "1",
			Resources: []proto.Message{listener},
		}},
	}
}

func TestResponseHandlers(t *testing.T) {
	for _, typ := range cache.ResponseTypes {
		t.Run(typ.String(), func(t *testing.T) {
			config := makeMockConfigWatcher()
			config.responses = makeResponses()
			s := NewServer(config)

			// make a request
			resp := makeMockStream(t)
			resp.recv <- &api.DiscoveryRequest{Node: node}
			go func() {
				var err error
				switch typ {
				case cache.EndpointResponse:
					err = s.StreamEndpoints(resp)
				case cache.ClusterResponse:
					err = s.StreamClusters(resp)
				case cache.RouteResponse:
					err = s.StreamRoutes(resp)
				case cache.ListenerResponse:
					err = s.StreamListeners(resp)
				}
				if err != nil {
					t.Errorf("Stream() => got %v, want no error", err)
				}
			}()

			// check a response
			select {
			case <-resp.sent:
				close(resp.recv)
				if want := map[cache.ResponseType]int{typ: 1}; !reflect.DeepEqual(want, config.counts) {
					t.Errorf("watch counts => got %v, want %v", config.counts, want)
				}
			case <-time.After(1 * time.Second):
				t.Fatalf("got no response")
			}
		})
	}
}

func TestWatchClosed(t *testing.T) {
	for _, typ := range cache.ResponseTypes {
		t.Run(typ.String(), func(t *testing.T) {
			config := makeMockConfigWatcher()
			config.closeWatch = true
			s := NewServer(config)

			// make a request
			resp := makeMockStream(t)
			resp.recv <- &api.DiscoveryRequest{
				Node:    node,
				TypeUrl: GetTypeURL(typ),
			}

			// check that response fails since watch gets closed
			if err := s.StreamAggregatedResources(resp); err == nil {
				t.Error("Stream() => got no error, want watch failed")
			}

			close(resp.recv)
		})
	}
}

func TestSendError(t *testing.T) {
	for _, typ := range cache.ResponseTypes {
		t.Run(typ.String(), func(t *testing.T) {
			config := makeMockConfigWatcher()
			config.responses = makeResponses()
			s := NewServer(config)

			// make a request
			resp := makeMockStream(t)
			resp.sendError = true
			resp.recv <- &api.DiscoveryRequest{
				Node:    node,
				TypeUrl: GetTypeURL(typ),
			}

			// check that response fails since watch gets closed
			if err := s.StreamAggregatedResources(resp); err == nil {
				t.Error("Stream() => got no error, want send error")
			}

			close(resp.recv)
		})
	}
}

func TestStaleNonce(t *testing.T) {
	for _, typ := range cache.ResponseTypes {
		t.Run(typ.String(), func(t *testing.T) {
			config := makeMockConfigWatcher()
			config.responses = makeResponses()
			s := NewServer(config)

			resp := makeMockStream(t)
			resp.recv <- &api.DiscoveryRequest{
				Node:    node,
				TypeUrl: GetTypeURL(typ),
			}
			stop := make(chan struct{})
			go func() {
				if err := s.StreamAggregatedResources(resp); err != nil {
					t.Errorf("StreamAggregatedResources() => got %v, want no error", err)
				}
				// should be two watches called
				if want := map[cache.ResponseType]int{typ: 2}; !reflect.DeepEqual(want, config.counts) {
					t.Errorf("watch counts => got %v, want %v", config.counts, want)
				}
				close(stop)
			}()
			select {
			case <-resp.sent:
				// stale request
				resp.recv <- &api.DiscoveryRequest{
					Node:          node,
					TypeUrl:       GetTypeURL(typ),
					ResponseNonce: "xyz",
				}
				// fresh request
				resp.recv <- &api.DiscoveryRequest{
					VersionInfo:   "1",
					Node:          node,
					TypeUrl:       GetTypeURL(typ),
					ResponseNonce: "1",
				}
				close(resp.recv)
			case <-time.After(1 * time.Second):
				t.Fatalf("got %d messages on the stream, not 4", resp.nonce)
			}
			<-stop
		})
	}
}

func TestAggregatedHandlers(t *testing.T) {
	config := makeMockConfigWatcher()
	config.responses = makeResponses()
	resp := makeMockStream(t)

	resp.recv <- &api.DiscoveryRequest{
		Node:    node,
		TypeUrl: ListenerType,
	}
	resp.recv <- &api.DiscoveryRequest{
		Node:    node,
		TypeUrl: ClusterType,
	}
	resp.recv <- &api.DiscoveryRequest{
		Node:          node,
		TypeUrl:       EndpointType,
		ResourceNames: []string{clusterName},
	}
	resp.recv <- &api.DiscoveryRequest{
		Node:          node,
		TypeUrl:       RouteType,
		ResourceNames: []string{routeName},
	}

	s := NewServer(config)
	go func() {
		if err := s.StreamAggregatedResources(resp); err != nil {
			t.Errorf("StreamAggregatedResources() => got %v, want no error", err)
		}
	}()

	count := 0
	for {
		select {
		case <-resp.sent:
			count++
			if count >= 4 {
				close(resp.recv)
				if want := map[cache.ResponseType]int{
					cache.EndpointResponse: 1,
					cache.ClusterResponse:  1,
					cache.RouteResponse:    1,
					cache.ListenerResponse: 1,
				}; !reflect.DeepEqual(want, config.counts) {
					t.Errorf("watch counts => got %v, want %v", config.counts, want)
				}

				// got all messages
				return
			}
		case <-time.After(1 * time.Second):
			t.Fatalf("got %d messages on the stream, not 4", count)
		}
	}
}

func TestAggregateRequestType(t *testing.T) {
	config := makeMockConfigWatcher()
	s := NewServer(config)
	resp := makeMockStream(t)
	resp.recv <- &api.DiscoveryRequest{Node: node}
	if err := s.StreamAggregatedResources(resp); err == nil {
		t.Error("StreamAggregatedResources() => got nil, want an error")
	}
}
