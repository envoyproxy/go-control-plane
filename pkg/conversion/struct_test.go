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

package conversion_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/conversion"
)

func TestConversion(t *testing.T) {
	pb := &discovery.DiscoveryRequest{
		VersionInfo: "test",
		Node:        &core.Node{Id: "proxy"},
	}
	st, err := conversion.MessageToStruct(pb)
	require.NoErrorf(t, err, "unexpected error")
	pbst := map[string]*structpb.Value{
		"version_info": {Kind: &structpb.Value_StringValue{StringValue: "test"}},
		"node": {Kind: &structpb.Value_StructValue{StructValue: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"id": {Kind: &structpb.Value_StringValue{StringValue: "proxy"}},
			},
		}}},
	}
	if !cmp.Equal(st.GetFields(), pbst, cmp.Comparer(proto.Equal)) {
		t.Errorf("MessageToStruct(%v) => got %v, want %v", pb, st.GetFields(), pbst)
	}

	out := &discovery.DiscoveryRequest{}
	require.NoErrorf(t, conversion.StructToMessage(st, out), "unexpected error")
	if !cmp.Equal(pb, out, cmp.Comparer(proto.Equal)) {
		t.Errorf("StructToMessage(%v) => got %v, want %v", st, out, pb)
	}

	_, err = conversion.MessageToStruct(nil)
	require.Errorf(t, err, "MessageToStruct(nil) => got no error")

	assert.Errorf(t, conversion.StructToMessage(nil, &discovery.DiscoveryRequest{}), "StructToMessage(nil) => got no error")
}
