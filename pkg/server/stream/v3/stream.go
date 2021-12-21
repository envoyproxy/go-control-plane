package stream

import (
	"sync"

	"google.golang.org/grpc"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
)

// Generic RPC stream.
type Stream interface {
	grpc.ServerStream

	Send(*discovery.DiscoveryResponse) error
	Recv() (*discovery.DiscoveryRequest, error)
}

type DeltaStream interface {
	grpc.ServerStream

	Send(*discovery.DeltaDiscoveryResponse) error
	Recv() (*discovery.DeltaDiscoveryRequest, error)
}

// StreamState will keep track of resource state per type on a stream.
type StreamState struct { // nolint:golint,revive
	// Indicates whether the original DeltaRequest was a wildcard LDS/RDS request.
	wildcard bool

	// ResourceVersions contains a hash of the resource as the value and the resource name as the key.
	// This field stores the last state sent to the client.
	resourceVersions map[string]string

	// knownResourceNames contains resource names that a client has received previously
	knownResourceNames map[string]map[string]struct{}

	// indicates whether the object has beed modified since its creation
	first bool
}

func (s *StreamState) GetResourceVersions() map[string]string {
	return s.resourceVersions
}

func (s *StreamState) SetResourceVersions(resourceVersions map[string]string) {
	s.first = false
	s.resourceVersions = resourceVersions
}

func (s *StreamState) IsFirst() bool {
	return s.first
}

func (s *StreamState) IsWildcard() bool {
	return s.wildcard
}

func (s *StreamState) SetKnownResourceNames(url string, names map[string]struct{}) {
	s.knownResourceNames[url] = names
}

func (s *StreamState) SetKnownResourceNamesAsList(url string, names []string) {
	m := map[string]struct{}{}
	for _, name := range names {
		m[name] = struct{}{}
	}
	s.knownResourceNames[url] = m
}

func (s *StreamState) GetKnownResourceNames(url string) map[string]struct{} {
	return s.knownResourceNames[url]
}

// NewStreamState initializes a stream state.
func NewStreamState(wildcard bool, initialResourceVersions map[string]string) StreamState {
	state := StreamState{
		wildcard:           wildcard,
		resourceVersions:   initialResourceVersions,
		first:              true,
		knownResourceNames: map[string]map[string]struct{}{},
	}

	if initialResourceVersions == nil {
		state.resourceVersions = make(map[string]string)
	}

	return state
}

func NewLastDiscoveryResponse(nonce string, resources []string) LastDiscoveryResponse {
	resp := LastDiscoveryResponse{
		Nonce:     nonce,
		Resources: make(map[string]struct{}),
	}
	for _, r := range resources {
		resp.Resources[r] = struct{}{}
	}
	return resp
}

// LastDiscoveryResponse that is sent over GRPC stream
// We need to record what resource names are already sent to a client
// So if the client requests a new name we can respond back
// regardless current snapshot version (even if it is not changed yet)
type LastDiscoveryResponse struct {
	Nonce     string
	Resources map[string]struct{}
}

func NewSTOWStreamState() STOWStreamState {
	return STOWStreamState{
		StreamState: NewStreamState(false, map[string]string{}),
		responses:   make(map[string]LastDiscoveryResponse),
	}
}

type STOWStreamState struct {
	StreamState
	responses map[string]LastDiscoveryResponse
	mu        sync.RWMutex
}

func (l *STOWStreamState) Set(key string, value LastDiscoveryResponse) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.responses[key] = value
}

func (l *STOWStreamState) Get(key string) (value LastDiscoveryResponse, ok bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	value, ok = l.responses[key]
	return
}
