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

package server

import (
	"errors"
	"log"
	"strconv"
	"sync/atomic"

	discovery "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v2"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type deltaStream interface {
	grpc.ServerStream

	Send(*discovery.DeltaDiscoveryResponse) error
	Recv() (*discovery.DeltaDiscoveryRequest, error)
}

func createDeltaResponse(resp *cache.DeltaResponse, typeURL string) (*v2.DeltaDiscoveryResponse, error) {
	if resp == nil {
		return nil, errors.New("missing response")
	}

	var resources []*v2.Resource
	if resp.ResourceMarshaled {
		resources = make([]*v2.Resource, len(resp.MarshaledResources))
	} else {
		resources = make([]*v2.Resource, len(resp.Resources))
	}

	for i := 0; i < len(resources); i++ {
		// Envoy relies on serialized protobuf bytes for detecting changes to the resources.
		// This requires deterministic serialization.
		if resp.ResourceMarshaled {
			resources[i] = &v2.Resource{
				Name: cache.GetResourceName(resp.Resources[i]),
				Resource: &any.Any{
					TypeUrl: typeURL,
					Value:   resp.MarshaledResources[i],
				},
			}
		} else {
			marshaledResource, err := cache.MarshalResource(resp.Resources[i])
			if err != nil {
				return nil, err
			}
			resources[i] = &v2.Resource{
				Name: cache.GetResourceName(resp.Resources[i]),
				Resource: &any.Any{
					TypeUrl: typeURL,
					Value:   marshaledResource,
				},
			}
		}
	}
	out := &v2.DeltaDiscoveryResponse{
		Resources: resources,
		TypeUrl:   typeURL,
	}
	return out, nil
}

func (s *server) deltaHandler(stream deltaStream, typeURL string) error {
	// a channel for receiving incoming delta requests
	reqCh := make(chan *v2.DeltaDiscoveryRequest)
	reqStop := int32(0)

	go func() {
		log.Printf("Started deltaHandler() request go routine for resource Type: %s\n", typeURL)
		for {
			req, err := stream.Recv()
			if atomic.LoadInt32(&reqStop) != 0 {
				return
			}
			if err != nil {
				close(reqCh)
				return
			}

			reqCh <- req
		}
	}()

	err := s.processDelta(stream, reqCh, typeURL)

	// prevents writing to a closed channel if send failed on blocked recv
	// TODO(kuat) figure out how to unblock recv through gRPC API
	atomic.StoreInt32(&reqStop, 1)

	return err
}

func (s *server) processDelta(stream deltaStream, reqCh <-chan *discovery.DeltaDiscoveryRequest, defaultTypeURL string) error {
	log.Printf("processDelta() for Type: %s\n", defaultTypeURL)

	// increment stream count
	streamID := atomic.AddInt64(&s.streamCount, 1)

	// unique nonce generator for req-resp pairs per xDS stream; the server
	// ignores stale nonces. nonce is only modified within send() function.
	var streamNonce int64

	// a collection of watches per request type
	var values watches
	defer func() {
		values.Cancel()
		if s.callbacks != nil {
			s.callbacks.OnStreamClosed(streamID)
		}
	}()

	// sends a response by serializing to protobuf Any
	send := func(resp cache.DeltaResponse, typeURL string) (string, error) {
		out, err := createDeltaResponse(&resp, typeURL)
		if err != nil {
			return "", err
		}

		// increment nonce
		streamNonce = streamNonce + 1
		out.Nonce = strconv.FormatInt(streamNonce, 10)
		if s.callbacks != nil {
			s.callbacks.OnStreamDeltaResponse(streamID, &resp.DeltaRequest, out)
		}
		return out.Nonce, stream.Send(out)
	}

	if s.callbacks != nil {
		if err := s.callbacks.OnStreamOpen(stream.Context(), streamID, defaultTypeURL); err != nil {
			return err
		}
	}

	// node may only be set on the first discovery request
	var node = &core.Node{}

	log.Printf("Starting send loop and watching for values on watches:")
	for {
		select {
		case <-s.ctx.Done():
			return nil
			// config watcher can send the requested resources types in any order
		case resp, more := <-values.deltaEndpoints:
			log.Printf("Received values over deltaEndpoints chan")
			if !more {
				return status.Errorf(codes.Unavailable, "endpoints watch failed")
			}
			nonce, err := send(resp, resource.EndpointType)
			if err != nil {
				return err
			}
			values.endpointNonce = nonce

		case resp, more := <-values.deltaClusters:
			log.Printf("Received values over deltaClusters chan")
			if !more {
				return status.Errorf(codes.Unavailable, "clusters watch failed")
			}
			nonce, err := send(resp, resource.ClusterType)
			if err != nil {
				return err
			}
			values.deltaClusterNonce = nonce

		case resp, more := <-values.deltaRoutes:
			log.Printf("Received values over deltaRoutes chan")
			if !more {
				return status.Errorf(codes.Unavailable, "routes watch failed")
			}
			nonce, err := send(resp, resource.RouteType)
			if err != nil {
				return err
			}
			values.routeNonce = nonce

		case resp, more := <-values.deltaListeners:
			log.Printf("Received values over deltaListeners chan")
			if !more {
				return status.Errorf(codes.Unavailable, "listeners watch failed")
			}
			nonce, err := send(resp, resource.ListenerType)
			if err != nil {
				return err
			}
			values.listenerNonce = nonce

		case resp, more := <-values.deltaSecrets:
			if !more {
				return status.Errorf(codes.Unavailable, "secrets watch failed")
			}
			nonce, err := send(resp, resource.SecretType)
			if err != nil {
				return err
			}
			values.secretNonce = nonce

		case resp, more := <-values.deltaRuntimes:
			if !more {
				return status.Errorf(codes.Unavailable, "runtimes watch failed")
			}
			nonce, err := send(resp, resource.RuntimeType)
			if err != nil {
				return err
			}
			values.runtimeNonce = nonce

		case req, more := <-reqCh:
			// input stream ended or errored out
			if !more {
				return nil
			}
			if req == nil {
				return status.Errorf(codes.Unavailable, "empty request")
			}

			// node field in discovery request is delta-compressed
			// nonces can be reused across streams; we verify nonce only if nonce is not initialized
			var nonce string
			if req.Node != nil {
				node = req.Node
				nonce = req.GetResponseNonce()
			} else {
				req.Node = node
				// If we have no nonce, i.e. this is the first request on a delta stream, set one
				nonce = strconv.FormatInt(streamNonce, 10)
			}

			// type URL is required for ADS but is implicit for xDS
			if defaultTypeURL == resource.AnyType {
				if req.TypeUrl == "" {
					return status.Errorf(codes.InvalidArgument, "type URL is required for ADS")
				}
			} else if req.TypeUrl == "" {
				req.TypeUrl = defaultTypeURL
			}

			if s.callbacks != nil {
				if err := s.callbacks.OnStreamDeltaRequest(streamID, req); err != nil {
					return err
				}
			}

			// cancel existing watches to (re-)request a newer version
			switch {
			case req.TypeUrl == resource.EndpointType && (values.endpointNonce == "" || values.endpointNonce == nonce):
				if values.endpointCancel != nil {
					values.endpointCancel()
				}
				values.deltaEndpoints, values.endpointCancel = s.cache.CreateDeltaWatch(*req)
			case req.TypeUrl == resource.ClusterType && (values.clusterNonce == "" || values.clusterNonce == nonce):
				if values.clusterCancel != nil {
					values.clusterCancel()
				}
				values.deltaClusters, values.clusterCancel = s.cache.CreateDeltaWatch(*req)
			case req.TypeUrl == resource.RouteType && (values.routeNonce == "" || values.routeNonce == nonce):
				if values.routeCancel != nil {
					values.routeCancel()
				}
				values.deltaRoutes, values.routeCancel = s.cache.CreateDeltaWatch(*req)
			case req.TypeUrl == resource.ListenerType && (values.listenerNonce == "" || values.listenerNonce == nonce):
				if values.listenerCancel != nil {
					values.listenerCancel()
				}
				values.deltaListeners, values.listenerCancel = s.cache.CreateDeltaWatch(*req)
			case req.TypeUrl == resource.SecretType && (values.secretNonce == "" || values.secretNonce == nonce):
				if values.secretCancel != nil {
					values.secretCancel()
				}
				values.deltaSecrets, values.secretCancel = s.cache.CreateDeltaWatch(*req)
			case req.TypeUrl == resource.RuntimeType && (values.runtimeNonce == "" || values.runtimeNonce == nonce):
				if values.runtimeCancel != nil {
					values.runtimeCancel()
				}
				values.deltaRuntimes, values.runtimeCancel = s.cache.CreateDeltaWatch(*req)
			}
		}
	}
}
