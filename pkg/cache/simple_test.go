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
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/envoyproxy/go-control-plane/api"
	"github.com/golang/protobuf/proto"
)

type group struct{}

const (
	key Key = "node"
)

func (group) Hash(node *api.Node) (Key, error) {
	if node == nil {
		return "", errors.New("nil node")
	}
	return key, nil
}

var (
	version  = "x"
	snapshot = NewSnapshot(version,
		[]proto.Message{endpoint},
		[]proto.Message{cluster},
		[]proto.Message{route},
		[]proto.Message{listener})
	names = map[ResponseType][]string{
		EndpointResponse: []string{clusterName},
		ClusterResponse:  nil,
		RouteResponse:    []string{routeName},
		ListenerResponse: nil,
	}
)

func TestSimpleCache(t *testing.T) {
	c := NewSimpleCache(group{}, nil)
	if err := c.SetSnapshot(key, snapshot); err != nil {
		t.Fatal(err)
	}

	// try to get endpoints with incorrect list of names
	// should not receive response
	es := c.Watch(EndpointResponse, &api.Node{}, "", []string{"none"})
	select {
	case out := <-es.Value:
		t.Errorf("watch for endpoints and mismatched names => got %v, want none", out)
	case <-time.After(time.Second / 4):
	}

	// try to get from nil node
	nilNode := c.Watch(ListenerResponse, nil, "", nil)
	if nilNode.Value != nil {
		t.Errorf("watch for nil node => got value %v, want none", nilNode.Value)
	}

	for _, typ := range ResponseTypes {
		t.Run(typ.String(), func(t *testing.T) {
			w := c.Watch(typ, &api.Node{}, "", names[typ])
			if w.Type != typ {
				t.Errorf("watch type => got %q, want %q", w.Type, typ)
			}
			if !reflect.DeepEqual(w.Names, names[typ]) {
				t.Errorf("watch names => got %q, want %q", w.Names, names[typ])
			}
			select {
			case out := <-w.Value:
				if out.Version != version {
					t.Errorf("got version %q, want %q", out.Version, version)
				}
				if !reflect.DeepEqual(out.Resources, snapshot.resources[typ]) {
					t.Errorf("get resources %v, want %v", out.Resources, snapshot.resources[typ])
				}
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}
}

func TestSimpleCacheWatch(t *testing.T) {
	c := NewSimpleCache(group{}, nil)
	watches := make(map[ResponseType]Watch)
	for _, typ := range ResponseTypes {
		watches[typ] = c.Watch(typ, &api.Node{}, "", names[typ])
	}
	if err := c.SetSnapshot(key, snapshot); err != nil {
		t.Fatal(err)
	}
	for _, typ := range ResponseTypes {
		t.Run(typ.String(), func(t *testing.T) {
			select {
			case out := <-watches[typ].Value:
				if out.Version != version {
					t.Errorf("got version %q, want %q", out.Version, version)
				}
				if !reflect.DeepEqual(out.Resources, snapshot.resources[typ]) {
					t.Errorf("get resources %v, want %v", out.Resources, snapshot.resources[typ])
				}
			case <-time.After(time.Second):
				t.Fatal("failed to receive snapshot response")
			}
		})
	}
}

func TestSimpleCacheWatchCancel(t *testing.T) {
	c := NewSimpleCache(group{}, nil)
	for _, typ := range ResponseTypes {
		watch := c.Watch(typ, &api.Node{}, "", names[typ])
		watch.Cancel()
	}
	for _, typ := range ResponseTypes {
		if count := len(c.(*SimpleCache).watches[key]); count > 0 {
			t.Errorf("watches should be released for %s", typ)
		}
	}
}

func TestSimpleCacheCallback(t *testing.T) {
	var called Key
	stop := make(chan struct{})
	c := NewSimpleCache(group{}, func(key Key) { called = key; close(stop) })
	c.Watch(ListenerResponse, &api.Node{}, "", nil)
	select {
	case <-stop:
	case <-time.After(time.Second):
		t.Fatal("callback not called")
	}
	if called != key {
		t.Errorf("got %q, callback not called for watch with missing node group", called)
	}
}
