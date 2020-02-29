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
	"bytes"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/golang/protobuf/jsonpb"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	cache "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/log"
)

// HTTPGateway is a custom implementation of [gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)
// specialized to Envoy xDS API.
type HTTPGateway struct {
	// Log is an optional log for errors in response write
	Log log.Logger

	// Server is the underlying gRPC server
	Server Server
}

func (h *HTTPGateway) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	p := path.Clean(req.URL.Path)

	typeURL := ""
	switch p {
	case "/v3/discovery:endpoints":
		typeURL = cache.EndpointType
	case "/v3/discovery:clusters":
		typeURL = cache.ClusterType
	case "/v3/discovery:listeners":
		typeURL = cache.ListenerType
	case "/v3/discovery:routes":
		typeURL = cache.RouteType
	case "/v3/discovery:secrets":
		typeURL = cache.SecretType
	case "/v3/discovery:runtime":
		typeURL = cache.RuntimeType
	default:
		http.Error(resp, "no endpoint", http.StatusNotFound)
		return
	}

	if req.Body == nil {
		http.Error(resp, "empty body", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(resp, "cannot read body", http.StatusBadRequest)
		return
	}

	// parse as JSON
	out := &discovery.DiscoveryRequest{}
	err = jsonpb.UnmarshalString(string(body), out)
	if err != nil {
		http.Error(resp, "cannot parse JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}
	out.TypeUrl = typeURL

	// fetch results
	res, err := h.Server.Fetch(req.Context(), out)
	if err != nil {
		// SkipFetchErrors will return a 304 which will signify to the envoy client that
		// it is already at the latest version; all other errors will 500 with a message.
		if _, ok := err.(*cache.SkipFetchError); ok {
			resp.WriteHeader(http.StatusNotModified)
		} else {
			http.Error(resp, "fetch error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	buf := &bytes.Buffer{}
	if err := (&jsonpb.Marshaler{OrigName: true}).Marshal(buf, res); err != nil {
		http.Error(resp, "marshal error: "+err.Error(), http.StatusInternalServerError)
	}

	if _, err = resp.Write(buf.Bytes()); err != nil && h.Log != nil {
		h.Log.Errorf("gateway error: %v", err)
	}
}
