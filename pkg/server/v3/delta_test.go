package server_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"google.golang.org/grpc"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
)

func (config *mockConfigWatcher) CreateDeltaWatch(req *discovery.DeltaDiscoveryRequest, st *stream.StreamState) (chan cache.DeltaResponse, func()) {
	config.deltaCounts[req.TypeUrl] = config.deltaCounts[req.TypeUrl] + 1

	// Create our out watch channel to return with a buffer of one
	out := make(chan cache.DeltaResponse, 1)

	if len(config.deltaResponses[req.TypeUrl]) > 0 {
		res := config.deltaResponses[req.TypeUrl][0]
		// In subscribed, we only want to send back what's changed if we detect changes
		var subscribed []types.Resource

		// Here we set our current version map to what the stream state says it is
		versionMap := st.ResourceVersions

		r, _ := res.GetDeltaDiscoveryResponse()
		if st.IsWildcard {
			for _, resource := range r.Resources {
				name := resource.GetName()
				marshaledResource, _ := cache.MarshalResource(resource)
				oldV := versionMap[name]
				if v := cache.HashResource(marshaledResource); v != oldV {
					versionMap[name] = v
				}
				subscribed = append(subscribed, resource)
			}
		} else {
			if len(req.GetResourceNamesSubscribe()) > 0 {
				for _, resource := range r.Resources {
					for _, alias := range req.GetResourceNamesSubscribe() {
						if name := resource.GetName(); name == alias {
							marshaledResource, _ := cache.MarshalResource(resource)
							oldV := versionMap[name]
							if v := cache.HashResource(marshaledResource); v != oldV {
								versionMap[name] = v
							}
							subscribed = append(subscribed, resource)
						}
					}
				}
			}
		}

		out <- &cache.RawDeltaResponse{
			DeltaRequest:      req,
			Resources:         subscribed,
			SystemVersionInfo: "",
			NextVersionMap:    versionMap,
		}
	} else if config.closeWatch {
		close(out)
	} else {
		config.deltaWatches += 1
		return out, func() {
			close(out)
			config.deltaWatches -= 1
		}
	}

	return out, nil
}

type mockDeltaStream struct {
	t         *testing.T
	ctx       context.Context
	recv      chan *discovery.DeltaDiscoveryRequest
	sent      chan *discovery.DeltaDiscoveryResponse
	nonce     int
	sendError bool
	grpc.ServerStream
}

func (stream *mockDeltaStream) Context() context.Context {
	return stream.ctx
}

func (stream *mockDeltaStream) Send(resp *discovery.DeltaDiscoveryResponse) error {
	// check that nonce is monotonically incrementing
	stream.nonce = stream.nonce + 1
	if resp.Nonce != fmt.Sprintf("%d", stream.nonce) {
		stream.t.Errorf("Nonce => got %q, want %d", resp.Nonce, stream.nonce)
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
		if res.Resource.TypeUrl != resp.TypeUrl {
			stream.t.Errorf("TypeUrl => got %q, want %q", res.Resource.TypeUrl, resp.TypeUrl)
		}
	}

	stream.sent <- resp
	if stream.sendError {
		return errors.New("send error")
	}
	return nil
}

func (stream *mockDeltaStream) Recv() (*discovery.DeltaDiscoveryRequest, error) {
	req, more := <-stream.recv
	if !more {
		return nil, errors.New("empty")
	}
	return req, nil
}

func makeMockDeltaStream(t *testing.T) *mockDeltaStream {
	return &mockDeltaStream{
		t:    t,
		ctx:  context.Background(),
		sent: make(chan *discovery.DeltaDiscoveryResponse, 10),
		recv: make(chan *discovery.DeltaDiscoveryRequest, 10),
	}
}

func makeDeltaResponses() map[string][]cache.DeltaResponse {
	return map[string][]cache.DeltaResponse{
		rsrc.EndpointType: {
			&cache.RawDeltaResponse{
				Resources:         []types.Resource{endpoint},
				DeltaRequest:      &discovery.DeltaDiscoveryRequest{TypeUrl: rsrc.EndpointType},
				SystemVersionInfo: "1",
			},
		},
		rsrc.ClusterType: {
			&cache.RawDeltaResponse{
				Resources:         []types.Resource{cluster},
				DeltaRequest:      &discovery.DeltaDiscoveryRequest{TypeUrl: rsrc.ClusterType},
				SystemVersionInfo: "2",
			},
		},
		rsrc.RouteType: {
			&cache.RawDeltaResponse{
				Resources:         []types.Resource{route},
				DeltaRequest:      &discovery.DeltaDiscoveryRequest{TypeUrl: rsrc.RouteType},
				SystemVersionInfo: "3",
			},
		},
		rsrc.ListenerType: {
			&cache.RawDeltaResponse{
				Resources:         []types.Resource{listener},
				DeltaRequest:      &discovery.DeltaDiscoveryRequest{TypeUrl: rsrc.ListenerType},
				SystemVersionInfo: "4",
			},
		},
		rsrc.SecretType: {
			&cache.RawDeltaResponse{
				SystemVersionInfo: "5",
				Resources:         []types.Resource{secret},
				DeltaRequest:      &discovery.DeltaDiscoveryRequest{TypeUrl: rsrc.SecretType},
			},
		},
		rsrc.RuntimeType: {
			&cache.RawDeltaResponse{
				SystemVersionInfo: "6",
				Resources:         []types.Resource{runtime},
				DeltaRequest:      &discovery.DeltaDiscoveryRequest{TypeUrl: rsrc.RuntimeType},
			},
		},
		// Pass-through type (xDS does not exist for this type)
		opaqueType: {
			&cache.RawDeltaResponse{
				SystemVersionInfo: "7",
				Resources:         []types.Resource{opaque},
				DeltaRequest:      &discovery.DeltaDiscoveryRequest{TypeUrl: opaqueType},
			},
		},
	}
}

func TestDeltaResponseHandlersWildcard(t *testing.T) {
	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			config := makeMockConfigWatcher()
			config.deltaResponses = makeDeltaResponses()
			s := server.NewServer(context.Background(), config, server.CallbackFuncs{})

			resp := makeMockDeltaStream(t)
			// This is a wildcard request since we don't specify a list of resource subscriptions
			resp.recv <- &discovery.DeltaDiscoveryRequest{Node: node, TypeUrl: typ}

			go func() {
				var err error
				switch typ {
				case rsrc.EndpointType:
					err = s.DeltaEndpoints(resp)
				case rsrc.ClusterType:
					err = s.DeltaClusters(resp)
				case rsrc.RouteType:
					err = s.DeltaRoutes(resp)
				case rsrc.ListenerType:
					err = s.DeltaListeners(resp)
				case rsrc.SecretType:
					err = s.DeltaSecrets(resp)
				case rsrc.RuntimeType:
					err = s.DeltaRuntime(resp)
				case opaqueType:
					err = s.DeltaAggregatedResources(resp)
				}

				if err != nil {
					t.Errorf("Delta() => got \"%v\", want no error", err)
				}
			}()

			select {
			case res := <-resp.sent:
				close(resp.recv)

				// We should only have 7 watch channels initialized since that is the base map length
				if config.deltaCounts[typ] != 1 {
					t.Errorf("watch counts for typ: %s => got %v, want 1", typ, config.deltaCounts[typ])
				}

				if v := res.GetSystemVersionInfo(); v != "" {
					t.Errorf("expected emtpy version on initial request, got %s", v)
				}
			case <-time.After(1 * time.Second):
				t.Fatalf("got no response")
			}
		})
	}
}

func TestDeltaResponseHandlers(t *testing.T) {
	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			config := makeMockConfigWatcher()
			config.deltaResponses = makeDeltaResponses()
			s := server.NewServer(context.Background(), config, server.CallbackFuncs{})

			resp := makeMockDeltaStream(t)
			// This is a wildcard request since we don't specify a list of resource subscriptions
			res, err := config.deltaResponses[typ][0].GetDeltaDiscoveryResponse()
			if err != nil {
				t.Error(err)
			}
			// We only subscribe to one resource to see if we get the appropriate number of responses back
			resp.recv <- &discovery.DeltaDiscoveryRequest{Node: node, TypeUrl: typ, ResourceNamesSubscribe: []string{res.Resources[0].Name}}

			go func() {
				var err error
				switch typ {
				case rsrc.EndpointType:
					err = s.DeltaEndpoints(resp)
				case rsrc.ClusterType:
					err = s.DeltaClusters(resp)
				case rsrc.RouteType:
					err = s.DeltaRoutes(resp)
				case rsrc.ListenerType:
					err = s.DeltaListeners(resp)
				case rsrc.SecretType:
					err = s.DeltaSecrets(resp)
				case rsrc.RuntimeType:
					err = s.DeltaRuntime(resp)
				case opaqueType:
					err = s.DeltaAggregatedResources(resp)
				}

				if err != nil {
					t.Errorf("Delta() => got \"%v\", want no error", err)
				}
			}()

			select {
			case res := <-resp.sent:
				close(resp.recv)

				// We should only have 7 watch channels initialized since that is the base map length
				if config.deltaCounts[typ] != 1 {
					t.Errorf("watch counts for typ: %s => got %v, want 1", typ, config.deltaCounts[typ])
				}
				if v := res.GetSystemVersionInfo(); v != "" {
					t.Errorf("expected emtpy version on initial request, got %s", v)
				}
			case <-time.After(1 * time.Second):
				t.Fatalf("got no response")
			}
		})
	}
}

func TestDeltaWatchClosed(t *testing.T) {
	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			config := makeMockConfigWatcher()
			config.closeWatch = true
			s := server.NewServer(context.Background(), config, server.CallbackFuncs{})

			resp := makeMockDeltaStream(t)
			resp.recv <- &discovery.DeltaDiscoveryRequest{
				Node:    node,
				TypeUrl: typ,
			}

			// Verify that the response fails when the watch is closed
			if err := s.DeltaAggregatedResources(resp); err == nil {
				t.Error("DeltaAggregatedResources() => got no error, want watch failed")
			}

			close(resp.recv)
		})
	}
}

func TestSendDeltaError(t *testing.T) {
	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			config := makeMockConfigWatcher()
			config.deltaResponses = makeDeltaResponses()
			s := server.NewServer(context.Background(), config, server.CallbackFuncs{})

			// make a request
			resp := makeMockDeltaStream(t)
			resp.sendError = true
			resp.recv <- &discovery.DeltaDiscoveryRequest{
				Node:    node,
				TypeUrl: typ,
			}

			// check that response fails since send returns error
			if err := s.DeltaAggregatedResources(resp); err == nil {
				t.Error("DeltaAggregatedResources() => got no error, want send error")
			}

			close(resp.recv)
		})
	}
}

func TestDeltaAggregatedHandlers(t *testing.T) {
	config := makeMockConfigWatcher()
	config.deltaResponses = makeDeltaResponses()
	resp := makeMockDeltaStream(t)

	resp.recv <- &discovery.DeltaDiscoveryRequest{
		Node:    node,
		TypeUrl: rsrc.ListenerType,
	}
	resp.recv <- &discovery.DeltaDiscoveryRequest{
		Node:    node,
		TypeUrl: rsrc.ClusterType,
	}
	resp.recv <- &discovery.DeltaDiscoveryRequest{
		Node:                   node,
		TypeUrl:                rsrc.EndpointType,
		ResourceNamesSubscribe: []string{clusterName},
	}
	resp.recv <- &discovery.DeltaDiscoveryRequest{
		TypeUrl:                rsrc.RouteType,
		ResourceNamesSubscribe: []string{routeName},
	}
	resp.recv <- &discovery.DeltaDiscoveryRequest{
		TypeUrl:                rsrc.SecretType,
		ResourceNamesSubscribe: []string{secretName},
	}

	s := server.NewServer(context.Background(), config, server.CallbackFuncs{})
	go func() {
		if err := s.DeltaAggregatedResources(resp); err != nil {
			t.Errorf("DeltaAggregatedResources() => got %v, want no error", err)
		}
	}()

	count := 0
	for {
		select {
		case <-resp.sent:
			count++
			if count >= 5 {
				close(resp.recv)
				if want := map[string]int{
					rsrc.EndpointType: 1,
					rsrc.ClusterType:  1,
					rsrc.RouteType:    1,
					rsrc.ListenerType: 1,
					rsrc.SecretType:   1,
				}; !reflect.DeepEqual(want, config.deltaCounts) {
					t.Errorf("watch counts => got %v, want %v", config.deltaCounts, want)
				}

				return
			}
		case <-time.After(1 * time.Second):
			t.Fatalf("got %d messages on the stream, not 5", count)
		}
	}
}
