// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/extensions/common/tap/v3/common.proto

package envoy_extensions_common_tap_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v1 "github.com/cncf/udpa/go/udpa/core/v1"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/tap/v3"
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

// Common configuration for all tap extensions.
type CommonExtensionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConfigType:
	//	*CommonExtensionConfig_AdminConfig
	//	*CommonExtensionConfig_StaticConfig
	//	*CommonExtensionConfig_TapdsConfig
	ConfigType isCommonExtensionConfig_ConfigType `protobuf_oneof:"config_type"`
}

func (x *CommonExtensionConfig) Reset() {
	*x = CommonExtensionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_tap_v3_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonExtensionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonExtensionConfig) ProtoMessage() {}

func (x *CommonExtensionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_tap_v3_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonExtensionConfig.ProtoReflect.Descriptor instead.
func (*CommonExtensionConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_tap_v3_common_proto_rawDescGZIP(), []int{0}
}

func (m *CommonExtensionConfig) GetConfigType() isCommonExtensionConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (x *CommonExtensionConfig) GetAdminConfig() *AdminConfig {
	if x, ok := x.GetConfigType().(*CommonExtensionConfig_AdminConfig); ok {
		return x.AdminConfig
	}
	return nil
}

func (x *CommonExtensionConfig) GetStaticConfig() *v3.TapConfig {
	if x, ok := x.GetConfigType().(*CommonExtensionConfig_StaticConfig); ok {
		return x.StaticConfig
	}
	return nil
}

func (x *CommonExtensionConfig) GetTapdsConfig() *CommonExtensionConfig_TapDSConfig {
	if x, ok := x.GetConfigType().(*CommonExtensionConfig_TapdsConfig); ok {
		return x.TapdsConfig
	}
	return nil
}

type isCommonExtensionConfig_ConfigType interface {
	isCommonExtensionConfig_ConfigType()
}

type CommonExtensionConfig_AdminConfig struct {
	// If specified, the tap filter will be configured via an admin handler.
	AdminConfig *AdminConfig `protobuf:"bytes,1,opt,name=admin_config,json=adminConfig,proto3,oneof"`
}

type CommonExtensionConfig_StaticConfig struct {
	// If specified, the tap filter will be configured via a static configuration that cannot be
	// changed.
	StaticConfig *v3.TapConfig `protobuf:"bytes,2,opt,name=static_config,json=staticConfig,proto3,oneof"`
}

type CommonExtensionConfig_TapdsConfig struct {
	// [#not-implemented-hide:] Configuration to use for TapDS updates for the filter.
	TapdsConfig *CommonExtensionConfig_TapDSConfig `protobuf:"bytes,3,opt,name=tapds_config,json=tapdsConfig,proto3,oneof"`
}

func (*CommonExtensionConfig_AdminConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_StaticConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_TapdsConfig) isCommonExtensionConfig_ConfigType() {}

// Configuration for the admin handler. See :ref:`here <config_http_filters_tap_admin_handler>` for
// more information.
type AdminConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Opaque configuration ID. When requests are made to the admin handler, the passed opaque ID is
	// matched to the configured filter opaque ID to determine which filter to configure.
	ConfigId string `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
}

func (x *AdminConfig) Reset() {
	*x = AdminConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_tap_v3_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminConfig) ProtoMessage() {}

func (x *AdminConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_tap_v3_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminConfig.ProtoReflect.Descriptor instead.
func (*AdminConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_tap_v3_common_proto_rawDescGZIP(), []int{1}
}

func (x *AdminConfig) GetConfigId() string {
	if x != nil {
		return x.ConfigId
	}
	return ""
}

// [#not-implemented-hide:]
type CommonExtensionConfig_TapDSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for the source of TapDS updates for this Cluster.
	ConfigSource *v31.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	// Tap config to request from XDS server.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Resource locator for TAP. This is mutually exclusive to *name*.
	// [#not-implemented-hide:]
	TapResourceLocator *v1.ResourceLocator `protobuf:"bytes,3,opt,name=tap_resource_locator,json=tapResourceLocator,proto3" json:"tap_resource_locator,omitempty"`
}

func (x *CommonExtensionConfig_TapDSConfig) Reset() {
	*x = CommonExtensionConfig_TapDSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_tap_v3_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonExtensionConfig_TapDSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonExtensionConfig_TapDSConfig) ProtoMessage() {}

func (x *CommonExtensionConfig_TapDSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_tap_v3_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonExtensionConfig_TapDSConfig.ProtoReflect.Descriptor instead.
func (*CommonExtensionConfig_TapDSConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_tap_v3_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CommonExtensionConfig_TapDSConfig) GetConfigSource() *v31.ConfigSource {
	if x != nil {
		return x.ConfigSource
	}
	return nil
}

func (x *CommonExtensionConfig_TapDSConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommonExtensionConfig_TapDSConfig) GetTapResourceLocator() *v1.ResourceLocator {
	if x != nil {
		return x.TapResourceLocator
	}
	return nil
}

var File_envoy_extensions_common_tap_v3_common_proto protoreflect.FileDescriptor

var file_envoy_extensions_common_tap_v3_common_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x76, 0x33,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33, 0x1a, 0x28, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x05, 0x0a, 0x15, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x45, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x76, 0x33, 0x2e, 0x54, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0c,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x66, 0x0a, 0x0c,
	0x74, 0x61, 0x70, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x41, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70,
	0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x61, 0x70, 0x44, 0x53, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x61, 0x70, 0x64, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0xbf, 0x02, 0x0a, 0x0b, 0x54, 0x61, 0x70, 0x44, 0x53, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x51, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x10, 0x12, 0x0e, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x67, 0x0a, 0x14, 0x74, 0x61, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72,
	0x42, 0x16, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x10, 0x12, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x12, 0x74, 0x61, 0x70, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x3a, 0x48, 0x9a, 0xc5,
	0x88, 0x1e, 0x43, 0x0a, 0x41, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x61, 0x70, 0x44, 0x53,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3c, 0x9a, 0xc5, 0x88, 0x1e, 0x37, 0x0a, 0x35, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x12, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x67, 0x0a, 0x0b, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x20, 0x01, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x3a, 0x32, 0x9a,
	0xc5, 0x88, 0x1e, 0x2d, 0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x45, 0x0a, 0x2c, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76,
	0x33, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_common_tap_v3_common_proto_rawDescOnce sync.Once
	file_envoy_extensions_common_tap_v3_common_proto_rawDescData = file_envoy_extensions_common_tap_v3_common_proto_rawDesc
)

func file_envoy_extensions_common_tap_v3_common_proto_rawDescGZIP() []byte {
	file_envoy_extensions_common_tap_v3_common_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_common_tap_v3_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_common_tap_v3_common_proto_rawDescData)
	})
	return file_envoy_extensions_common_tap_v3_common_proto_rawDescData
}

var file_envoy_extensions_common_tap_v3_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_common_tap_v3_common_proto_goTypes = []interface{}{
	(*CommonExtensionConfig)(nil),             // 0: envoy.extensions.common.tap.v3.CommonExtensionConfig
	(*AdminConfig)(nil),                       // 1: envoy.extensions.common.tap.v3.AdminConfig
	(*CommonExtensionConfig_TapDSConfig)(nil), // 2: envoy.extensions.common.tap.v3.CommonExtensionConfig.TapDSConfig
	(*v3.TapConfig)(nil),                      // 3: envoy.config.tap.v3.TapConfig
	(*v31.ConfigSource)(nil),                  // 4: envoy.config.core.v3.ConfigSource
	(*v1.ResourceLocator)(nil),                // 5: udpa.core.v1.ResourceLocator
}
var file_envoy_extensions_common_tap_v3_common_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.common.tap.v3.CommonExtensionConfig.admin_config:type_name -> envoy.extensions.common.tap.v3.AdminConfig
	3, // 1: envoy.extensions.common.tap.v3.CommonExtensionConfig.static_config:type_name -> envoy.config.tap.v3.TapConfig
	2, // 2: envoy.extensions.common.tap.v3.CommonExtensionConfig.tapds_config:type_name -> envoy.extensions.common.tap.v3.CommonExtensionConfig.TapDSConfig
	4, // 3: envoy.extensions.common.tap.v3.CommonExtensionConfig.TapDSConfig.config_source:type_name -> envoy.config.core.v3.ConfigSource
	5, // 4: envoy.extensions.common.tap.v3.CommonExtensionConfig.TapDSConfig.tap_resource_locator:type_name -> udpa.core.v1.ResourceLocator
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_common_tap_v3_common_proto_init() }
func file_envoy_extensions_common_tap_v3_common_proto_init() {
	if File_envoy_extensions_common_tap_v3_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_common_tap_v3_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonExtensionConfig); i {
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
		file_envoy_extensions_common_tap_v3_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminConfig); i {
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
		file_envoy_extensions_common_tap_v3_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonExtensionConfig_TapDSConfig); i {
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
	file_envoy_extensions_common_tap_v3_common_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CommonExtensionConfig_AdminConfig)(nil),
		(*CommonExtensionConfig_StaticConfig)(nil),
		(*CommonExtensionConfig_TapdsConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_common_tap_v3_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_common_tap_v3_common_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_common_tap_v3_common_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_common_tap_v3_common_proto_msgTypes,
	}.Build()
	File_envoy_extensions_common_tap_v3_common_proto = out.File
	file_envoy_extensions_common_tap_v3_common_proto_rawDesc = nil
	file_envoy_extensions_common_tap_v3_common_proto_goTypes = nil
	file_envoy_extensions_common_tap_v3_common_proto_depIdxs = nil
}
