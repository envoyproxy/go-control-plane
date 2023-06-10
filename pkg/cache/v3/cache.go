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

// Package cache defines a configuration cache for the server.
package cache

import (
	"context"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
)

type Request interface {
	GetNode() *core.Node
	GetTypeUrl() string
	GetVersionInfo() string
}

type DeltaRequest interface {
	GetNode() *core.Node
	GetTypeUrl() string
}

// SubscriptionState provides additional data on the client knowledge for the type matching the request
// This allows proper implementation of stateful aspects of the protocol (e.g. returning only some updated resources)
// Though the methods may return mutable parts of the state for performance reasons,
// the cache is expected to consider this state as immutable and thread safe between a watch creation and its cancellation
type SubscriptionState interface {
	// GetKnownResources returns a list of resources that the client has ACK'd and their associated version.
	// The versions are:
	//  - delta protocol: version of the specific resource set in the response
	//  - sotw protocol: version of the global response when the resource was last ACKed
	GetKnownResources() map[string]string

	// GetSubscribedResources returns the list of resources currently subscribed to by the client for the type.
	// For delta it keeps track of subscription updates across requests
	// For sotw it is a normalized view of the last request resources
	GetSubscribedResources() map[string]struct{}

	// IsWildcard returns whether the client has a wildcard watch.
	// This considers subtleties related to the current migration of wildcard definitions within the protocol.
	// More details on the behavior of wildcard are present at https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#how-the-client-specifies-what-resources-to-return
	IsWildcard() bool

	// IsFirstResponse returns whether the subscription has already been replied to at least once
	// It is needed for some edge cases when a response must be returned (even if empty) at first
	IsFirstResponse() bool

	// WatchesResources returns whether at least one of the resources provided is currently being watched by the subscription.
	// It is currently only applicable to delta-xds.
	// If the request is wildcard, it will always return true,
	// otherwise it will compare the provided resources to the list of resources currently subscribed
	WatchesResources(resourceNames map[string]struct{}) bool
}

// ConfigWatcher requests watches for configuration resources by a node, last
// applied version identifier, and resource names hint. The watch should send
// the responses when they are ready. The watch can be canceled by the
// consumer, in effect terminating the watch for the request.
// ConfigWatcher implementation must be thread-safe.
type ConfigWatcher interface {
	// CreateWatch returns a new open watch from a non-empty request.
	// This is the entrypoint to propagate configuration changes the
	// provided Response channel. State from the gRPC server is utilized
	// to make sure consuming cache implementations can see what the server has sent to clients.
	//
	// An individual consumer normally issues a single open watch by each type URL.
	//
	// The provided channel produces requested resources as responses, once they are available.
	//
	// Cancel is an optional function to release resources in the producer. If
	// provided, the consumer may call this function multiple times.
	CreateWatch(Request, SubscriptionState, chan Response) (cancel func(), err error)

	// CreateDeltaWatch returns a new open incremental xDS watch.
	// This is the entrypoint to propagate configuration changes the
	// provided DeltaResponse channel. State from the gRPC server is utilized
	// to make sure consuming cache implementations can see what the server has sent to clients.
	//
	// The provided channel produces requested resources as responses, or spontaneous updates in accordance
	// with the incremental xDS specification.
	//
	// Cancel is an optional function to release resources in the producer. If
	// provided, the consumer may call this function multiple times.
	CreateDeltaWatch(DeltaRequest, SubscriptionState, chan DeltaResponse) (cancel func(), err error)
}

type FetchRequest interface {
	Request
	GetResourceNames() []string
}

// ConfigFetcher fetches configuration resources from cache
type ConfigFetcher interface {
	// Fetch implements the polling method of the config cache using a non-empty request.
	Fetch(context.Context, FetchRequest) (Response, error)
}

// Cache is a generic config cache with a watcher.
type Cache interface {
	ConfigWatcher
	ConfigFetcher
}

// Response is a wrapper around Envoy's DiscoveryResponse.
type Response interface {
	GetRequest() Request

	// Get the version in the Response.
	GetVersion() (string, error)

	// Get the list of resources included in the response.
	GetReturnedResources() []string

	// Get the context provided during response creation.
	GetContext() context.Context

	// Get the Constructed DiscoveryResponse
	GetDiscoveryResponse() (*discovery.DiscoveryResponse, error)

	// Get the request that created the watch that we're now responding to. This is provided to allow the caller to correlate the
	// response with a request.
	GetDiscoveryRequest() (*discovery.DiscoveryRequest, error)
}

// DeltaResponse is a wrapper around Envoy's DeltaDiscoveryResponse
type DeltaResponse interface {
	GetDeltaRequest() DeltaRequest

	// Get the version in the DeltaResponse. This field is generally used for debugging purposes as noted by the Envoy documentation.
	GetSystemVersion() (string, error)

	// Get the list of resources included in the response.
	// If the resource was set in Resources, its version is provided
	// If the version was set in RemovedResources, its version is ""
	GetReturnedResources() map[string]string

	// Get the context provided during response creation
	GetContext() context.Context

	// Get the constructed DeltaDiscoveryResponse
	GetDeltaDiscoveryResponse() (*discovery.DeltaDiscoveryResponse, error)

	// Get the request that created the watch that we're now responding to. This is provided to allow the caller to correlate the
	// response with a request.
	GetDeltaDiscoveryRequest() (*discovery.DeltaDiscoveryRequest, error)
}
