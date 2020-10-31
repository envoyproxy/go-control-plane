// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.13.0
// source: envoy/config/listener/v4alpha/udp_gso_batch_writer_config.proto

package envoy_config_listener_v4alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

// [#not-implemented-hide:]
// Configuration specific to the Udp Gso Batch Writer.
type UdpGsoBatchWriterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UdpGsoBatchWriterOptions) Reset() {
	*x = UdpGsoBatchWriterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UdpGsoBatchWriterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UdpGsoBatchWriterOptions) ProtoMessage() {}

func (x *UdpGsoBatchWriterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UdpGsoBatchWriterOptions.ProtoReflect.Descriptor instead.
func (*UdpGsoBatchWriterOptions) Descriptor() ([]byte, []int) {
	return file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDescGZIP(), []int{0}
}

var File_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto protoreflect.FileDescriptor

var file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x75, 0x64, 0x70, 0x5f, 0x67, 0x73, 0x6f, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x54, 0x0a, 0x18, 0x55, 0x64, 0x70, 0x47, 0x73, 0x6f, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x38,
	0x9a, 0xc5, 0x88, 0x1e, 0x33, 0x0a, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x55, 0x64, 0x70, 0x47, 0x73, 0x6f, 0x42, 0x61, 0x74, 0x63, 0x68, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x55, 0x0a, 0x2b, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x1c, 0x55, 0x64, 0x70, 0x47, 0x73, 0x6f, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x03, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDescOnce sync.Once
	file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDescData = file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDesc
)

func file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDescGZIP() []byte {
	file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDescData)
	})
	return file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDescData
}

var file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_goTypes = []interface{}{
	(*UdpGsoBatchWriterOptions)(nil), // 0: envoy.config.listener.v4alpha.UdpGsoBatchWriterOptions
}
var file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_init() }
func file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_init() {
	if File_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UdpGsoBatchWriterOptions); i {
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
			RawDescriptor: file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_msgTypes,
	}.Build()
	File_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto = out.File
	file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_rawDesc = nil
	file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_goTypes = nil
	file_envoy_config_listener_v4alpha_udp_gso_batch_writer_config_proto_depIdxs = nil
}
