package sotw

import (
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

// watches for all xDS resource types
type watches struct {
	endpoints        chan cache.Response
	clusters         chan cache.Response
	routes           chan cache.Response
	listeners        chan cache.Response
	secrets          chan cache.Response
	runtimes         chan cache.Response
	extensionConfigs chan cache.Response

	endpointCancel        func()
	clusterCancel         func()
	routeCancel           func()
	listenerCancel        func()
	secretCancel          func()
	runtimeCancel         func()
	extensionConfigCancel func()

	endpointNonce        string
	clusterNonce         string
	routeNonce           string
	listenerNonce        string
	secretNonce          string
	runtimeNonce         string
	extensionConfigNonce string

	// Opaque resources share a muxed channel. Nonces and watch cancellations are indexed by type URL.
	responses     chan cache.Response
	cancellations map[string]func()
	nonces        map[string]string
}

func newWatches() watches {
	return watches{
		responses:     make(chan cache.Response, 5), // muxed channel needs a buffer to release go-routines populating it
		cancellations: make(map[string]func()),
		nonces:        make(map[string]string),
	}
}

// Cancel all watches
func (values *watches) Cancel() {
	if values.endpointCancel != nil {
		values.endpointCancel()
	}
	if values.clusterCancel != nil {
		values.clusterCancel()
	}
	if values.routeCancel != nil {
		values.routeCancel()
	}
	if values.listenerCancel != nil {
		values.listenerCancel()
	}
	if values.secretCancel != nil {
		values.secretCancel()
	}
	if values.runtimeCancel != nil {
		values.runtimeCancel()
	}
	if values.extensionConfigCancel != nil {
		values.extensionConfigCancel()
	}
	for _, cancel := range values.cancellations {
		if cancel != nil {
			cancel()
		}
	}
}

// process handles a bi-di stream request
func (s *server) process(stream Stream, reqCh <-chan *discovery.DiscoveryRequest, defaultTypeURL string) error {
	sw := streamWrapper{
		stream:    stream,
		id:        atomic.AddInt64(&s.streamCount, 1), // increment stream count
		callbacks: s.callbacks,
	}

	// a collection of stack allocated watches per request type
	values := newWatches()
	defer values.Cancel()
	defer func() {
		if s.callbacks != nil {
			s.callbacks.OnStreamClosed(sw.id)
		}
	}()

	if s.callbacks != nil {
		if err := s.callbacks.OnStreamOpen(stream.Context(), sw.id, defaultTypeURL); err != nil {
			return err
		}
	}

	// node may only be set on the first discovery request
	var node = &core.Node{}

	for {
		select {
		case <-s.ctx.Done():
			return nil
		// config watcher can send the requested resources types in any order
		case resp, more := <-values.endpoints:
			if !more {
				return status.Errorf(codes.Unavailable, "endpoints watch failed")
			}
			nonce, err := sw.send(resp)
			if err != nil {
				return err
			}
			values.endpointNonce = nonce

		case resp, more := <-values.clusters:
			if !more {
				return status.Errorf(codes.Unavailable, "clusters watch failed")
			}
			nonce, err := sw.send(resp)
			if err != nil {
				return err
			}
			values.clusterNonce = nonce

		case resp, more := <-values.routes:
			if !more {
				return status.Errorf(codes.Unavailable, "routes watch failed")
			}
			nonce, err := sw.send(resp)
			if err != nil {
				return err
			}
			values.routeNonce = nonce

		case resp, more := <-values.listeners:
			if !more {
				return status.Errorf(codes.Unavailable, "listeners watch failed")
			}
			nonce, err := sw.send(resp)
			if err != nil {
				return err
			}
			values.listenerNonce = nonce

		case resp, more := <-values.secrets:
			if !more {
				return status.Errorf(codes.Unavailable, "secrets watch failed")
			}
			nonce, err := sw.send(resp)
			if err != nil {
				return err
			}
			values.secretNonce = nonce

		case resp, more := <-values.runtimes:
			if !more {
				return status.Errorf(codes.Unavailable, "runtimes watch failed")
			}
			nonce, err := sw.send(resp)
			if err != nil {
				return err
			}
			values.runtimeNonce = nonce

		case resp, more := <-values.extensionConfigs:
			if !more {
				return status.Errorf(codes.Unavailable, "extensionConfigs watch failed")
			}
			nonce, err := sw.send(resp)
			if err != nil {
				return err
			}
			values.extensionConfigNonce = nonce

		case resp, more := <-values.responses:
			if more {
				if resp == errorResponse {
					return status.Errorf(codes.Unavailable, "resource watch failed")
				}
				typeUrl := resp.GetRequest().TypeUrl
				nonce, err := sw.send(resp)
				if err != nil {
					return err
				}
				values.nonces[typeUrl] = nonce
			}

		case req, more := <-reqCh:
			// input stream ended or errored out
			if !more {
				return nil
			}
			if req == nil {
				return status.Errorf(codes.Unavailable, "empty request")
			}

			// node field in discovery request is delta-compressed
			if req.Node != nil {
				node = req.Node
			} else {
				req.Node = node
			}

			// nonces can be reused across streams; we verify nonce only if nonce is not initialized
			nonce := req.GetResponseNonce()

			// type URL is required for ADS but is implicit for xDS
			if defaultTypeURL == resource.AnyType {
				if req.TypeUrl == "" {
					return status.Errorf(codes.InvalidArgument, "type URL is required for ADS")
				}
			} else if req.TypeUrl == "" {
				req.TypeUrl = defaultTypeURL
			}

			if s.callbacks != nil {
				if err := s.callbacks.OnStreamRequest(sw.id, req); err != nil {
					return err
				}
			}

			// cancel existing watches to (re-)request a newer version
			switch {
			case req.TypeUrl == resource.EndpointType:
				if values.endpointNonce == "" || values.endpointNonce == nonce {
					if values.endpointCancel != nil {
						values.endpointCancel()
					}
					values.endpoints = make(chan cache.Response, 1)
					values.endpointCancel = s.cache.CreateWatch(req, values.endpoints)
				}
			case req.TypeUrl == resource.ClusterType:
				if values.clusterNonce == "" || values.clusterNonce == nonce {
					if values.clusterCancel != nil {
						values.clusterCancel()
					}
					values.clusters = make(chan cache.Response, 1)
					values.clusterCancel = s.cache.CreateWatch(req, values.clusters)
				}
			case req.TypeUrl == resource.RouteType:
				if values.routeNonce == "" || values.routeNonce == nonce {
					if values.routeCancel != nil {
						values.routeCancel()
					}
					values.routes = make(chan cache.Response, 1)
					values.routeCancel = s.cache.CreateWatch(req, values.routes)
				}
			case req.TypeUrl == resource.ListenerType:
				if values.listenerNonce == "" || values.listenerNonce == nonce {
					if values.listenerCancel != nil {
						values.listenerCancel()
					}
					values.listeners = make(chan cache.Response, 1)
					values.listenerCancel = s.cache.CreateWatch(req, values.listeners)
				}
			case req.TypeUrl == resource.SecretType:
				if values.secretNonce == "" || values.secretNonce == nonce {
					if values.secretCancel != nil {
						values.secretCancel()
					}
					values.secrets = make(chan cache.Response, 1)
					values.secretCancel = s.cache.CreateWatch(req, values.secrets)
				}
			case req.TypeUrl == resource.RuntimeType:
				if values.runtimeNonce == "" || values.runtimeNonce == nonce {
					if values.runtimeCancel != nil {
						values.runtimeCancel()
					}
					values.runtimes = make(chan cache.Response, 1)
					values.runtimeCancel = s.cache.CreateWatch(req, values.runtimes)
				}
			case req.TypeUrl == resource.ExtensionConfigType:
				if values.extensionConfigNonce == "" || values.extensionConfigNonce == nonce {
					if values.extensionConfigCancel != nil {
						values.extensionConfigCancel()
					}
					values.extensionConfigs = make(chan cache.Response, 1)
					values.extensionConfigCancel = s.cache.CreateWatch(req, values.extensionConfigs)
				}
			default:
				typeUrl := req.TypeUrl
				responseNonce, seen := values.nonces[typeUrl]
				if !seen || responseNonce == nonce {
					if cancel, seen := values.cancellations[typeUrl]; seen && cancel != nil {
						cancel()
					}
					values.cancellations[typeUrl] = s.cache.CreateWatch(req, values.responses)
				}
			}
		}
	}
}
