// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: envoy/extensions/matching/input_matchers/runtime_fraction/v3/runtime_fraction.proto

package runtime_fractionv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

// The runtime fraction matchers computes a hash from the input and matches if runtime feature is enabled
// for the the resulting hash. Every time the input is considered for a match, its hash must fall within
// the percentage of matches indicated by this field. For a fraction N/D, a number is computed as a hash
// of the input on a field in the range [0,D). If the number is less than or equal to the value of the
// numerator N, the matcher evaluates to true. A runtime_fraction input matcher can be used to gradually
// roll out matcher changes without requiring full code or configuration deployments.
// Note that distribution of matching results is only as good as one of the input.
type RuntimeFraction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Match the input against the given runtime key. The specified default value is used if key is not
	// present in the runtime configuration.
	RuntimeFraction *v3.RuntimeFractionalPercent `protobuf:"bytes,1,opt,name=runtime_fraction,json=runtimeFraction,proto3" json:"runtime_fraction,omitempty"`
	// Optional seed passed through the hash function. This allows using additional information when computing
	// the hash value: by changing the seed value, a potentially different outcome can be achieved for the same input.
	Seed uint64 `protobuf:"varint,2,opt,name=seed,proto3" json:"seed,omitempty"`
}

func (x *RuntimeFraction) Reset() {
	*x = RuntimeFraction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuntimeFraction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeFraction) ProtoMessage() {}

func (x *RuntimeFraction) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeFraction.ProtoReflect.Descriptor instead.
func (*RuntimeFraction) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDescGZIP(), []int{0}
}

func (x *RuntimeFraction) GetRuntimeFraction() *v3.RuntimeFractionalPercent {
	if x != nil {
		return x.RuntimeFraction
	}
	return nil
}

func (x *RuntimeFraction) GetSeed() uint64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

var File_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto protoreflect.FileDescriptor

var file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDesc = []byte{
	0x0a, 0x53, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67,
	0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a,
	0x0f, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x63, 0x0a, 0x10, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x42, 0xe4, 0x01, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x0a, 0x4a, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33,
	0x42, 0x14, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x76, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x3b, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x33,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDescOnce sync.Once
	file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDescData = file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDesc
)

func file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDescGZIP() []byte {
	file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDescData)
	})
	return file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDescData
}

var file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_goTypes = []interface{}{
	(*RuntimeFraction)(nil),             // 0: envoy.extensions.matching.input_matchers.runtime_fraction.v3.RuntimeFraction
	(*v3.RuntimeFractionalPercent)(nil), // 1: envoy.config.core.v3.RuntimeFractionalPercent
}
var file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.matching.input_matchers.runtime_fraction.v3.RuntimeFraction.runtime_fraction:type_name -> envoy.config.core.v3.RuntimeFractionalPercent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_init()
}
func file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_init() {
	if File_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuntimeFraction); i {
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
			RawDescriptor: file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_msgTypes,
	}.Build()
	File_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto = out.File
	file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_rawDesc = nil
	file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_goTypes = nil
	file_envoy_extensions_matching_input_matchers_runtime_fraction_v3_runtime_fraction_proto_depIdxs = nil
}
