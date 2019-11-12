// Copyright 2019 Envoyproxy Authors
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

// Package proxy implements an xDS proxy
package proxy

import (
	"context"
	"errors"
	"sync"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	discoverygrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/envoyproxy/go-control-plane/pkg/server"
)

// Proxy forwards xDS requests over ADS but keeps a version of the received resources to respond with.
type Proxy struct {
	hash  cache.NodeHash
	cache cache.SnapshotCache
	pool  *Pool
}

// Update function updates an existing snapshot of configuration resources with a new
// discovery response.
func (p *Proxy) Update(id string, resp *v2.DiscoveryResponse) {
	snapshot, err := p.cache.GetSnapshot(id)

}

// Pool of client streams to xDS servers grouped by node IDs.
// This pool spawns a go routine per receiving client stream.
// Any error occurring in forwarding is propagated downstream.
type Pool struct {
	ctx    context.Context
	hash   cache.NodeHash
	client discoverygrpc.AggregatedDiscoveryServiceClient
	recv   func(string, *v2.DiscoveryResponse)

	pool map[string]discoverygrpc.AggregatedDiscoveryService_StreamAggregatedResourcesClient
	mu   sync.RWMutex
}

func (p *Pool) Send(req *v2.DiscoveryRequest) error {
	id := p.hash.ID(req.Node)

	p.mu.Lock()
	defer p.mu.Unlock()
	active, found := p.pool[id]
	if !found {
		var err error
		active, err = p.client.StreamAggregatedResources(p.ctx /** TODO: add options */)
		if err != nil {
			return err
		}

		p.pool[id] = active
		go func() {
			for {
				msg, err := active.Recv()
				if err != nil {
					p.mu.Lock()
					delete(p.pool, id)
					p.mu.Unlock()
					return
				}
				p.recv(id, msg)
			}
		}()
	}
	return active.Send(req)
}

var _ server.Callbacks = &Proxy{}

func (p *Proxy) OnStreamRequest(id int64, req *v2.DiscoveryRequest) error {
	return p.pool.Send(req)
}
func (p *Proxy) OnFetchRequest(context.Context, *v2.DiscoveryRequest) error {
	return errors.New("proxy from REST to ADS not implemented")
}
func (p *Proxy) OnStreamResponse(int64, *v2.DiscoveryRequest, *v2.DiscoveryResponse) {}
func (p *Proxy) OnFetchResponse(*v2.DiscoveryRequest, *v2.DiscoveryResponse)         {}
func (p *Proxy) OnStreamOpen(context.Context, int64, string) error {
	return nil
}
func (p *Proxy) OnStreamClosed(int64) {}
