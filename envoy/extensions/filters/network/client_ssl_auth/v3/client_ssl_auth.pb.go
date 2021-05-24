// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: envoy/extensions/filters/network/client_ssl_auth/v3/client_ssl_auth.proto

package envoy_extensions_filters_network_client_ssl_auth_v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ClientSSLAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The :ref:`cluster manager <arch_overview_cluster_manager>` cluster that runs
	// the authentication service. The filter will connect to the service every 60s to fetch the list
	// of principals. The service must support the expected :ref:`REST API
	// <config_network_filters_client_ssl_auth_rest_api>`.
	AuthApiCluster string `protobuf:"bytes,1,opt,name=auth_api_cluster,json=authApiCluster,proto3" json:"auth_api_cluster,omitempty"`
	// The prefix to use when emitting :ref:`statistics
	// <config_network_filters_client_ssl_auth_stats>`.
	StatPrefix string `protobuf:"bytes,2,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Time in milliseconds between principal refreshes from the
	// authentication service. Default is 60000 (60s). The actual fetch time
	// will be this value plus a random jittered value between
	// 0-refresh_delay_ms milliseconds.
	RefreshDelay *duration.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	// An optional list of IP address and subnet masks that should be white
	// listed for access by the filter. If no list is provided, there is no
	// IP allowlist.
	IpWhiteList []*v3.CidrRange `protobuf:"bytes,4,rep,name=ip_white_list,json=ipWhiteList,proto3" json:"ip_white_list,omitempty"`
}

func (x *ClientSSLAuth) Reset() {
	*x = ClientSSLAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientSSLAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientSSLAuth) ProtoMessage() {}

func (x *ClientSSLAuth) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientSSLAuth.ProtoReflect.Descriptor instead.
func (*ClientSSLAuth) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ClientSSLAuth) GetAuthApiCluster() string {
	if x != nil {
		return x.AuthApiCluster
	}
	return ""
}

func (x *ClientSSLAuth) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *ClientSSLAuth) GetRefreshDelay() *duration.Duration {
	if x != nil {
		return x.RefreshDelay
	}
	return nil
}

func (x *ClientSSLAuth) GetIpWhiteList() []*v3.CidrRange {
	if x != nil {
		return x.IpWhiteList
	}
	return nil
}

var File_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDesc = []byte{
	0x0a, 0x49, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x73, 0x6c, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x73, 0x6c,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x73, 0x6c, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33,
	0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd2, 0x02, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x53, 0x4c, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x37, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a,
	0x72, 0x08, 0x10, 0x01, 0xc0, 0x01, 0x02, 0xc8, 0x01, 0x00, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68,
	0x41, 0x70, 0x69, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x3e, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x12, 0x59, 0x0a, 0x0d, 0x69, 0x70, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x43, 0x69, 0x64, 0x72, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x14, 0xf2, 0x98,
	0xfe, 0x8f, 0x05, 0x0e, 0x0a, 0x0c, 0x69, 0x70, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x0b, 0x69, 0x70, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x3a,
	0x43, 0x9a, 0xc5, 0x88, 0x1e, 0x3e, 0x0a, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x73, 0x6c, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x53, 0x4c,
	0x41, 0x75, 0x74, 0x68, 0x42, 0x61, 0x0a, 0x41, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x73,
	0x6c, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x42, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x73, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDescData = file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDesc
)

func file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDescData
}

var file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_goTypes = []interface{}{
	(*ClientSSLAuth)(nil),     // 0: envoy.extensions.filters.network.client_ssl_auth.v3.ClientSSLAuth
	(*duration.Duration)(nil), // 1: google.protobuf.Duration
	(*v3.CidrRange)(nil),      // 2: envoy.config.core.v3.CidrRange
}
var file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.network.client_ssl_auth.v3.ClientSSLAuth.refresh_delay:type_name -> google.protobuf.Duration
	2, // 1: envoy.extensions.filters.network.client_ssl_auth.v3.ClientSSLAuth.ip_white_list:type_name -> envoy.config.core.v3.CidrRange
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_init() }
func file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_init() {
	if File_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientSSLAuth); i {
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
			RawDescriptor: file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto = out.File
	file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_rawDesc = nil
	file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_goTypes = nil
	file_envoy_extensions_filters_network_client_ssl_auth_v3_client_ssl_auth_proto_depIdxs = nil
}
