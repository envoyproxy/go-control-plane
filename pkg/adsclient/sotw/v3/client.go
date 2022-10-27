// Copyright 2022 Envoyproxy Authors
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

// Package sotw provides an implementation of GRPC SoTW (State of The World) part of XDS client
package sotw

import (
	"context"
	"errors"
	"io"
	"sync"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/golang/protobuf/ptypes/any"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

var (
	errInit    = errors.New("ads client: grpc connection is not initialized (use InitConnect() method to initialize connection)")
	errNilResp = errors.New("ads client: nil response from xds management server")
)

type AdsClient interface {
	// Initialize the gRPC connection with management server and send the initial Discovery Request.
	InitConnect(clientConn grpc.ClientConnInterface, opts ...grpc.CallOption) error
	// Fetch waits for a response from management server and returns response or error.
	Fetch() (*Response, error)
	// Ack acknowledge the validity of the last received response to management server.
	Ack() error
	// Nack acknowledge the invalidity of the last received response to management server.
	Nack(message string) error
}

type Response struct {
	Resources []*any.Any
}

type adsClient struct {
	ctx     context.Context
	mu      sync.RWMutex
	nodeID  string
	typeURL string

	// streamClient is the ADS discovery client
	streamClient discovery.AggregatedDiscoveryService_StreamAggregatedResourcesClient
	// lastAckedResponse is the last response acked by the ADS client
	lastAckedResponse *discovery.DiscoveryResponse
	// lastReceivedResponse is the last response received from management server
	lastReceivedResponse *discovery.DiscoveryResponse
}

func NewAdsClient(ctx context.Context, nodeID string, typeURL string) AdsClient {
	return &adsClient{
		ctx:     ctx,
		nodeID:  nodeID,
		typeURL: typeURL,
	}
}

func (c *adsClient) InitConnect(clientConn grpc.ClientConnInterface, opts ...grpc.CallOption) error {
	streamClient, err := discovery.NewAggregatedDiscoveryServiceClient(clientConn).StreamAggregatedResources(c.ctx, opts...)
	if err != nil {
		return err
	}
	c.streamClient = streamClient
	return c.Ack()
}

func (c *adsClient) Fetch() (*Response, error) {
	if c.streamClient == nil {
		return nil, errInit
	}
	resp, err := c.streamClient.Recv()
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errNilResp
	}

	c.mu.Lock()
	c.lastReceivedResponse = resp
	c.mu.Unlock()

	return &Response{
		Resources: resp.GetResources(),
	}, err
}

func (c *adsClient) Ack() error {
	c.mu.Lock()
	c.lastAckedResponse = c.lastReceivedResponse
	c.mu.Unlock()
	return c.send(nil)
}

func (c *adsClient) Nack(message string) error {
	errorDetail := &status.Status{
		Message: message,
	}
	return c.send(errorDetail)
}

// IsConnError checks the provided error is due to the gRPC connection
// and returns true if it is due to the gRPC connection.
//
// In this case the gRPC connection with the server should be re initialized with the
// AdsClient.InitConnect method.
func IsConnError(err error) bool {
	if err == nil {
		return false
	}
	if err == io.EOF {
		return true
	}
	errStatus, ok := grpcStatus.FromError(err)
	if !ok {
		return false
	}
	return errStatus.Code() == codes.Unavailable || errStatus.Code() == codes.Canceled
}

func (c *adsClient) send(errorDetail *status.Status) error {
	c.mu.RLock()
	req := &discovery.DiscoveryRequest{
		Node:          &core.Node{Id: c.nodeID},
		VersionInfo:   c.lastAckedResponse.GetVersionInfo(),
		TypeUrl:       c.typeURL,
		ResponseNonce: c.lastReceivedResponse.GetNonce(),
		ErrorDetail:   errorDetail,
	}
	c.mu.RUnlock()

	if c.streamClient == nil {
		return errInit
	}
	return c.streamClient.Send(req)
}
