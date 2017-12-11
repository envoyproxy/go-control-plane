// Copyright 2017 Envoyproxy Authors
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

package util_test

import (
	"reflect"
	"testing"

	"github.com/envoyproxy/go-control-plane/api"
	"github.com/envoyproxy/go-control-plane/pkg/util"
	"github.com/gogo/protobuf/types"
)

func TestMessageToStruct(t *testing.T) {
	pb := &api.DiscoveryRequest{
		VersionInfo: "test",
		Node:        &api.Node{Id: "proxy"},
	}
	st, err := util.MessageToStruct(pb)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	want := map[string]*types.Value{
		"version_info": &types.Value{Kind: &types.Value_StringValue{StringValue: "test"}},
		"node": &types.Value{Kind: &types.Value_StructValue{StructValue: &types.Struct{
			Fields: map[string]*types.Value{
				"id": &types.Value{Kind: &types.Value_StringValue{StringValue: "proxy"}},
			},
		}}},
	}
	if !reflect.DeepEqual(st.Fields, want) {
		t.Errorf("MessageToStruct(%v) => got %v, want %v", pb, st.Fields, want)
	}

	if _, err = util.MessageToStruct(nil); err == nil {
		t.Error("MessageToStruct(nil) => got no error")
	}
}
