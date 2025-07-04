// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.3
// source: envoy/extensions/router/cluster_specifiers/matcher/v3/matcher.proto

package matcherv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/cncf/xds/go/xds/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClusterAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates the upstream cluster to which the request should be routed
	// to.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *ClusterAction) Reset() {
	*x = ClusterAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterAction) ProtoMessage() {}

func (x *ClusterAction) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterAction.ProtoReflect.Descriptor instead.
func (*ClusterAction) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterAction) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

type MatcherClusterSpecifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The matcher for cluster selection after the route has been selected. This is used when the
	// route has multiple clusters (like multiple clusters for different users) and the matcher
	// is used to select the cluster to use for the request.
	//
	// The match tree to use for grouping incoming requests into buckets.
	//
	// Example:
	//
	// .. validated-code-block:: yaml
	//
	//	:type-name: xds.type.matcher.v3.Matcher
	//
	//	matcher_list:
	//	  matchers:
	//	  - predicate:
	//	      single_predicate:
	//	        input:
	//	          typed_config:
	//	            '@type': type.googleapis.com/envoy.type.matcher.v3.HttpRequestHeaderMatchInput
	//	            header_name: env
	//	        value_match:
	//	          exact: staging
	//	    on_match:
	//	      action:
	//	        typed_config:
	//	          '@type': type.googleapis.com/envoy.extensions.router.cluster_specifiers.matcher.v3.ClusterAction
	//	          cluster: "staging-cluster"
	//
	//	  - predicate:
	//	      single_predicate:
	//	        input:
	//	          typed_config:
	//	            '@type': type.googleapis.com/envoy.type.matcher.v3.HttpRequestHeaderMatchInput
	//	            header_name: env
	//	        value_match:
	//	          exact: prod
	//	    on_match:
	//	      action:
	//	        typed_config:
	//	          '@type': type.googleapis.com/envoy.extensions.router.cluster_specifiers.matcher.v3.ClusterAction
	//	          cluster: "prod-cluster"
	//
	//	# Catch-all with a default cluster.
	//	on_no_match:
	//	  action:
	//	    typed_config:
	//	      '@type': type.googleapis.com/envoy.extensions.router.cluster_specifiers.matcher.v3.ClusterAction
	//	      cluster: "default-cluster"
	ClusterMatcher *v3.Matcher `protobuf:"bytes,1,opt,name=cluster_matcher,json=clusterMatcher,proto3" json:"cluster_matcher,omitempty"`
}

func (x *MatcherClusterSpecifier) Reset() {
	*x = MatcherClusterSpecifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatcherClusterSpecifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatcherClusterSpecifier) ProtoMessage() {}

func (x *MatcherClusterSpecifier) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatcherClusterSpecifier.ProtoReflect.Descriptor instead.
func (*MatcherClusterSpecifier) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDescGZIP(), []int{1}
}

func (x *MatcherClusterSpecifier) GetClusterMatcher() *v3.Matcher {
	if x != nil {
		return x.ClusterMatcher
	}
	return nil
}

var File_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto protoreflect.FileDescriptor

var file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDesc = []byte{
	0x0a, 0x43, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x21, 0x78, 0x64,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x6a, 0x0a, 0x17, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x0f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0xc5, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x0a, 0x43, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x66, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x76, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDescOnce sync.Once
	file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDescData = file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDesc
)

func file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDescGZIP() []byte {
	file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDescData)
	})
	return file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDescData
}

var file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_goTypes = []interface{}{
	(*ClusterAction)(nil),           // 0: envoy.extensions.router.cluster_specifiers.matcher.v3.ClusterAction
	(*MatcherClusterSpecifier)(nil), // 1: envoy.extensions.router.cluster_specifiers.matcher.v3.MatcherClusterSpecifier
	(*v3.Matcher)(nil),              // 2: xds.type.matcher.v3.Matcher
}
var file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.router.cluster_specifiers.matcher.v3.MatcherClusterSpecifier.cluster_matcher:type_name -> xds.type.matcher.v3.Matcher
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_init() }
func file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_init() {
	if File_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatcherClusterSpecifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_msgTypes,
	}.Build()
	File_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto = out.File
	file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_rawDesc = nil
	file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_goTypes = nil
	file_envoy_extensions_router_cluster_specifiers_matcher_v3_matcher_proto_depIdxs = nil
}
