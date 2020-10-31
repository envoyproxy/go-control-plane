// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.13.0
// source: envoy/type/matcher/v4alpha/path.proto

package envoy_type_matcher_v4alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// Specifies the way to match a path on HTTP request.
type PathMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Rule:
	//	*PathMatcher_Path
	Rule isPathMatcher_Rule `protobuf_oneof:"rule"`
}

func (x *PathMatcher) Reset() {
	*x = PathMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v4alpha_path_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathMatcher) ProtoMessage() {}

func (x *PathMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v4alpha_path_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathMatcher.ProtoReflect.Descriptor instead.
func (*PathMatcher) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v4alpha_path_proto_rawDescGZIP(), []int{0}
}

func (m *PathMatcher) GetRule() isPathMatcher_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (x *PathMatcher) GetPath() *StringMatcher {
	if x, ok := x.GetRule().(*PathMatcher_Path); ok {
		return x.Path
	}
	return nil
}

type isPathMatcher_Rule interface {
	isPathMatcher_Rule()
}

type PathMatcher_Path struct {
	// The `path` must match the URL path portion of the :path header. The query and fragment
	// string (if present) are removed in the URL path portion.
	// For example, the path */data* will match the *:path* header */data#fragment?param=value*.
	Path *StringMatcher `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

func (*PathMatcher_Path) isPathMatcher_Rule() {}

var File_envoy_type_matcher_v4alpha_path_proto protoreflect.FileDescriptor

var file_envoy_type_matcher_v4alpha_path_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x61, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x68,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x3a, 0x28, 0x9a, 0xc5, 0x88, 0x1e, 0x23, 0x0a, 0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x0b, 0x0a, 0x04,
	0x72, 0x75, 0x6c, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x3f, 0x0a, 0x28, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x34,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x09, 0x50, 0x61, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_type_matcher_v4alpha_path_proto_rawDescOnce sync.Once
	file_envoy_type_matcher_v4alpha_path_proto_rawDescData = file_envoy_type_matcher_v4alpha_path_proto_rawDesc
)

func file_envoy_type_matcher_v4alpha_path_proto_rawDescGZIP() []byte {
	file_envoy_type_matcher_v4alpha_path_proto_rawDescOnce.Do(func() {
		file_envoy_type_matcher_v4alpha_path_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_matcher_v4alpha_path_proto_rawDescData)
	})
	return file_envoy_type_matcher_v4alpha_path_proto_rawDescData
}

var file_envoy_type_matcher_v4alpha_path_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_type_matcher_v4alpha_path_proto_goTypes = []interface{}{
	(*PathMatcher)(nil),   // 0: envoy.type.matcher.v4alpha.PathMatcher
	(*StringMatcher)(nil), // 1: envoy.type.matcher.v4alpha.StringMatcher
}
var file_envoy_type_matcher_v4alpha_path_proto_depIdxs = []int32{
	1, // 0: envoy.type.matcher.v4alpha.PathMatcher.path:type_name -> envoy.type.matcher.v4alpha.StringMatcher
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_type_matcher_v4alpha_path_proto_init() }
func file_envoy_type_matcher_v4alpha_path_proto_init() {
	if File_envoy_type_matcher_v4alpha_path_proto != nil {
		return
	}
	file_envoy_type_matcher_v4alpha_string_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_matcher_v4alpha_path_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathMatcher); i {
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
	file_envoy_type_matcher_v4alpha_path_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PathMatcher_Path)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_matcher_v4alpha_path_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_matcher_v4alpha_path_proto_goTypes,
		DependencyIndexes: file_envoy_type_matcher_v4alpha_path_proto_depIdxs,
		MessageInfos:      file_envoy_type_matcher_v4alpha_path_proto_msgTypes,
	}.Build()
	File_envoy_type_matcher_v4alpha_path_proto = out.File
	file_envoy_type_matcher_v4alpha_path_proto_rawDesc = nil
	file_envoy_type_matcher_v4alpha_path_proto_goTypes = nil
	file_envoy_type_matcher_v4alpha_path_proto_depIdxs = nil
}
