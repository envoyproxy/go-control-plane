// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: contrib/envoy/extensions/filters/http/checksum/v3alpha/checksum.proto

package v3alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
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

type ChecksumConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A set of matcher and checksum pairs for which, if a path matching ``path_matcher``
	// is requested and the checksum of the response body does not match the ``sha256``, the
	// response will be replaced with a 403 Forbidden status.
	//
	// If multiple matchers match the same path, the first to match takes precedence.
	Checksums []*ChecksumConfig_Checksum `protobuf:"bytes,1,rep,name=checksums,proto3" json:"checksums,omitempty"`
	// If a request doesn't match any of the specified checksum paths and reject_unmatched is
	// true, the request is rejected immediately with 403 Forbidden.
	RejectUnmatched bool `protobuf:"varint,2,opt,name=reject_unmatched,json=rejectUnmatched,proto3" json:"reject_unmatched,omitempty"`
}

func (x *ChecksumConfig) Reset() {
	*x = ChecksumConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChecksumConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChecksumConfig) ProtoMessage() {}

func (x *ChecksumConfig) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChecksumConfig.ProtoReflect.Descriptor instead.
func (*ChecksumConfig) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDescGZIP(), []int{0}
}

func (x *ChecksumConfig) GetChecksums() []*ChecksumConfig_Checksum {
	if x != nil {
		return x.Checksums
	}
	return nil
}

func (x *ChecksumConfig) GetRejectUnmatched() bool {
	if x != nil {
		return x.RejectUnmatched
	}
	return false
}

type ChecksumConfig_Checksum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Matcher:
	//	*ChecksumConfig_Checksum_PathMatcher
	Matcher isChecksumConfig_Checksum_Matcher `protobuf_oneof:"matcher"`
	// A hex-encoded sha256 string required to match the sha256sum of the response body
	// of the path specified in the ``path_matcher`` field.
	Sha256 string `protobuf:"bytes,2,opt,name=sha256,proto3" json:"sha256,omitempty"`
}

func (x *ChecksumConfig_Checksum) Reset() {
	*x = ChecksumConfig_Checksum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChecksumConfig_Checksum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChecksumConfig_Checksum) ProtoMessage() {}

func (x *ChecksumConfig_Checksum) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChecksumConfig_Checksum.ProtoReflect.Descriptor instead.
func (*ChecksumConfig_Checksum) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ChecksumConfig_Checksum) GetMatcher() isChecksumConfig_Checksum_Matcher {
	if m != nil {
		return m.Matcher
	}
	return nil
}

func (x *ChecksumConfig_Checksum) GetPathMatcher() *v3.StringMatcher {
	if x, ok := x.GetMatcher().(*ChecksumConfig_Checksum_PathMatcher); ok {
		return x.PathMatcher
	}
	return nil
}

func (x *ChecksumConfig_Checksum) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

type isChecksumConfig_Checksum_Matcher interface {
	isChecksumConfig_Checksum_Matcher()
}

type ChecksumConfig_Checksum_PathMatcher struct {
	// A matcher for a path that is expected to have a specific checksum, as specified
	// in the ``sha256`` field.
	PathMatcher *v3.StringMatcher `protobuf:"bytes,1,opt,name=path_matcher,json=pathMatcher,proto3,oneof"`
}

func (*ChecksumConfig_Checksum_PathMatcher) isChecksumConfig_Checksum_Matcher() {}

var File_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x2e,
	0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75,
	0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x65, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x73, 0x12, 0x29,
	0x0a, 0x10, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x6e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x55, 0x6e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x1a, 0x9b, 0x01, 0x0a, 0x08, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x53, 0x0a, 0x0c, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0b,
	0x70, 0x61, 0x74, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x73,
	0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xfa, 0x42, 0x14,
	0x72, 0x12, 0x32, 0x10, 0x5e, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30, 0x2d, 0x39, 0x5d,
	0x7b, 0x36, 0x34, 0x7d, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x42, 0x09, 0x0a, 0x07,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0xb6, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a, 0x3c, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73,
	0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDescData = file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDesc
)

func file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDescData
}

var file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_goTypes = []interface{}{
	(*ChecksumConfig)(nil),          // 0: envoy.extensions.filters.http.checksum.v3alpha.ChecksumConfig
	(*ChecksumConfig_Checksum)(nil), // 1: envoy.extensions.filters.http.checksum.v3alpha.ChecksumConfig.Checksum
	(*v3.StringMatcher)(nil),        // 2: envoy.type.matcher.v3.StringMatcher
}
var file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.http.checksum.v3alpha.ChecksumConfig.checksums:type_name -> envoy.extensions.filters.http.checksum.v3alpha.ChecksumConfig.Checksum
	2, // 1: envoy.extensions.filters.http.checksum.v3alpha.ChecksumConfig.Checksum.path_matcher:type_name -> envoy.type.matcher.v3.StringMatcher
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_init() }
func file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_init() {
	if File_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChecksumConfig); i {
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
		file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChecksumConfig_Checksum); i {
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
	file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ChecksumConfig_Checksum_PathMatcher)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_depIdxs,
		MessageInfos:      file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto = out.File
	file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_rawDesc = nil
	file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_goTypes = nil
	file_contrib_envoy_extensions_filters_http_checksum_v3alpha_checksum_proto_depIdxs = nil
}
