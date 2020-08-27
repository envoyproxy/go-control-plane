package server_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	rsrc "github.com/envoyproxy/go-control-plane/pkg/resource/v2"
	"github.com/envoyproxy/go-control-plane/pkg/server/v2"
	"google.golang.org/grpc"
)

func (config *mockConfigWatcher) CreateDeltaWatch(req *discovery.DeltaDiscoveryRequest, version string) (chan cache.DeltaResponse, func()) {
	config.counts[req.TypeUrl] = config.counts[req.TypeUrl] + 1

	// Create our out watch channel to return with a buffer of one
	out := make(chan cache.DeltaResponse, 1)

	if len(config.deltaResponses[req.TypeUrl]) > 0 {
		res := config.deltaResponses[req.TypeUrl][0]
		var subscribed []types.Resource

		// Only return back the subscribed resources to our request type
		r, _ := res.GetDeltaDiscoveryResponse()
		for _, resource := range r.Resources {
			for _, alias := range req.GetResourceNamesSubscribe() {
				if resource.GetName() == alias {
					subscribed = append(subscribed, resource)
				}
			}
		}

		// We should only send back subscribed resources here
		out <- &cache.RawDeltaResponse{
			DeltaRequest:      req,
			Resources:         subscribed,
			SystemVersionInfo: version,
		}

	} else if config.closeWatch {
		fmt.Printf("No resources... closing watch\n")
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

	// TODO:
	// Not sure this is a valid check because I think one incremental payload can carry various types of resources
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
				Resources:         []types.Resource{deltaCluster, deltaCluster2},
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
	}
}

<<<<<<< HEAD
// TESTS =======================================================================================

=======
>>>>>>> split delta server tests out and add test for all response handlers
func TestDeltaResponseClusterHandler(t *testing.T) {
	t.Run(testTypes[1], func(t *testing.T) {
		config := makeMockConfigWatcher()
		config.deltaResponses = makeDeltaResponses()
		s := server.NewServer(context.Background(), config, server.CallbackFuncs{}, logger{t})

		// make a request
		resp := makeMockDeltaStream(t)
		resp.recv <- &discovery.DeltaDiscoveryRequest{Node: node, TypeUrl: testTypes[1], ResourceNamesSubscribe: []string{"cluster0"}}

		go func() {
			err := s.DeltaClusters(resp)

			if err != nil {
				t.Errorf("Delta() => got %v, want no error", err)
			}
		}()

		// check a response
		select {
		case res := <-resp.sent:
			close(resp.recv)

			if want := map[string]int{testTypes[1]: 1}; !reflect.DeepEqual(want, config.counts) {
				t.Errorf("watch counts => got %v, want %v", config.counts, want)
			}

			if want := res.GetSystemVersionInfo(); want == "" {
				t.Errorf("response was missing version")
			}
		case <-time.After(1 * time.Second):
			t.Fatalf("got no response")
		}
	})
}
<<<<<<< HEAD
=======

func TestDeltaResponseHandlers(t *testing.T) {
	for _, typ := range testTypes {
		t.Run(typ, func(t *testing.T) {
			config := makeMockConfigWatcher()
			config.deltaResponses = makeDeltaResponses()
			s := server.NewServer(context.Background(), config, server.CallbackFuncs{}, logger{t})

			resp := makeMockDeltaStream(t)
			// This should put through a wildcard request but the test configWatcher might not work here
			resp.recv <- &discovery.DeltaDiscoveryRequest{Node: node, TypeUrl: typ, ResourceNamesSubscribe: []string{clusterName}}

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
				}

				if err != nil {
					t.Errorf("Delta() => got %v, want no error", err)
				}
			}()

			select {
			case res := <-resp.sent:
				close(resp.recv)
				if want := map[string]int{typ: 1}; !reflect.DeepEqual(want, config.counts) {
					t.Errorf("watch counts => got %v, want %v", config.counts, want)
				}

				if v := res.GetSystemVersionInfo(); v != "" {
					t.Errorf("should've had an emtpy version for first request, got %s", v)
				}
			case <-time.After(1 * time.Second):
				t.Fatalf("got no response")
			}
		})
	}
}
>>>>>>> split delta server tests out and add test for all response handlers
