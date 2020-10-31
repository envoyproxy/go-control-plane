// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.13.0
// source: envoy/extensions/transport_sockets/tls/v3/secret.proto

package envoy_extensions_transport_sockets_tls_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v1 "github.com/cncf/udpa/go/udpa/core/v1"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

type GenericSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Secret of generic type and is available to filters.
	Secret *v3.DataSource `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *GenericSecret) Reset() {
	*x = GenericSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericSecret) ProtoMessage() {}

func (x *GenericSecret) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericSecret.ProtoReflect.Descriptor instead.
func (*GenericSecret) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDescGZIP(), []int{0}
}

func (x *GenericSecret) GetSecret() *v3.DataSource {
	if x != nil {
		return x.Secret
	}
	return nil
}

type SdsSecretConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name (FQDN, UUID, SPKI, SHA256, etc.) by which the secret can be uniquely referred to.
	// When both name and config are specified, then secret can be fetched and/or reloaded via
	// SDS. When only name is specified, then secret will be loaded from static resources.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Resource locator for SDS. This is mutually exclusive to *name*.
	// [#not-implemented-hide:]
	SdsResourceLocator *v1.ResourceLocator `protobuf:"bytes,3,opt,name=sds_resource_locator,json=sdsResourceLocator,proto3" json:"sds_resource_locator,omitempty"`
	SdsConfig          *v3.ConfigSource    `protobuf:"bytes,2,opt,name=sds_config,json=sdsConfig,proto3" json:"sds_config,omitempty"`
}

func (x *SdsSecretConfig) Reset() {
	*x = SdsSecretConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SdsSecretConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SdsSecretConfig) ProtoMessage() {}

func (x *SdsSecretConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SdsSecretConfig.ProtoReflect.Descriptor instead.
func (*SdsSecretConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDescGZIP(), []int{1}
}

func (x *SdsSecretConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SdsSecretConfig) GetSdsResourceLocator() *v1.ResourceLocator {
	if x != nil {
		return x.SdsResourceLocator
	}
	return nil
}

func (x *SdsSecretConfig) GetSdsConfig() *v3.ConfigSource {
	if x != nil {
		return x.SdsConfig
	}
	return nil
}

// [#next-free-field: 6]
type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name (FQDN, UUID, SPKI, SHA256, etc.) by which the secret can be uniquely referred to.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Type:
	//	*Secret_TlsCertificate
	//	*Secret_SessionTicketKeys
	//	*Secret_ValidationContext
	//	*Secret_GenericSecret
	Type isSecret_Type `protobuf_oneof:"type"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDescGZIP(), []int{2}
}

func (x *Secret) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *Secret) GetType() isSecret_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Secret) GetTlsCertificate() *TlsCertificate {
	if x, ok := x.GetType().(*Secret_TlsCertificate); ok {
		return x.TlsCertificate
	}
	return nil
}

func (x *Secret) GetSessionTicketKeys() *TlsSessionTicketKeys {
	if x, ok := x.GetType().(*Secret_SessionTicketKeys); ok {
		return x.SessionTicketKeys
	}
	return nil
}

func (x *Secret) GetValidationContext() *CertificateValidationContext {
	if x, ok := x.GetType().(*Secret_ValidationContext); ok {
		return x.ValidationContext
	}
	return nil
}

func (x *Secret) GetGenericSecret() *GenericSecret {
	if x, ok := x.GetType().(*Secret_GenericSecret); ok {
		return x.GenericSecret
	}
	return nil
}

type isSecret_Type interface {
	isSecret_Type()
}

type Secret_TlsCertificate struct {
	TlsCertificate *TlsCertificate `protobuf:"bytes,2,opt,name=tls_certificate,json=tlsCertificate,proto3,oneof"`
}

type Secret_SessionTicketKeys struct {
	SessionTicketKeys *TlsSessionTicketKeys `protobuf:"bytes,3,opt,name=session_ticket_keys,json=sessionTicketKeys,proto3,oneof"`
}

type Secret_ValidationContext struct {
	ValidationContext *CertificateValidationContext `protobuf:"bytes,4,opt,name=validation_context,json=validationContext,proto3,oneof"`
}

type Secret_GenericSecret struct {
	GenericSecret *GenericSecret `protobuf:"bytes,5,opt,name=generic_secret,json=genericSecret,proto3,oneof"`
}

func (*Secret_TlsCertificate) isSecret_Type() {}

func (*Secret_SessionTicketKeys) isSecret_Type() {}

func (*Secret_ValidationContext) isSecret_Type() {}

func (*Secret_GenericSecret) isSecret_Type() {}

var File_envoy_extensions_transport_sockets_tls_v3_secret_proto protoreflect.FileDescriptor

var file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDesc = []byte{
	0x0a, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73,
	0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x06, 0xb8, 0xb7, 0x8b,
	0xa4, 0x02, 0x01, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x3a, 0x26, 0x9a, 0xc5, 0x88,
	0x1e, 0x21, 0x0a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x22, 0x9a, 0x02, 0x0a, 0x0f, 0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x31, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0xf2, 0x98,
	0xfe, 0x8f, 0x05, 0x10, 0x12, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x67, 0x0a, 0x14, 0x73, 0x64,
	0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x16, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x10, 0x12,
	0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52,
	0x12, 0x73, 0x64, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x0a, 0x73, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x73, 0x64, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x28, 0x9a, 0xc5, 0x88, 0x1e, 0x23, 0x0a, 0x21, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0xfb, 0x03, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x64, 0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x71, 0x0a, 0x13, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x54,
	0x6c, 0x73, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x73, 0x48, 0x00, 0x52, 0x11, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x78, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52,
	0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x61, 0x0a, 0x0e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e,
	0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x3a, 0x1f, 0x9a, 0xc5, 0x88, 0x1e, 0x1a, 0x0a, 0x18, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x50,
	0x0a, 0x37, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDescOnce sync.Once
	file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDescData = file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDesc
)

func file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDescGZIP() []byte {
	file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDescData)
	})
	return file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDescData
}

var file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_transport_sockets_tls_v3_secret_proto_goTypes = []interface{}{
	(*GenericSecret)(nil),                // 0: envoy.extensions.transport_sockets.tls.v3.GenericSecret
	(*SdsSecretConfig)(nil),              // 1: envoy.extensions.transport_sockets.tls.v3.SdsSecretConfig
	(*Secret)(nil),                       // 2: envoy.extensions.transport_sockets.tls.v3.Secret
	(*v3.DataSource)(nil),                // 3: envoy.config.core.v3.DataSource
	(*v1.ResourceLocator)(nil),           // 4: udpa.core.v1.ResourceLocator
	(*v3.ConfigSource)(nil),              // 5: envoy.config.core.v3.ConfigSource
	(*TlsCertificate)(nil),               // 6: envoy.extensions.transport_sockets.tls.v3.TlsCertificate
	(*TlsSessionTicketKeys)(nil),         // 7: envoy.extensions.transport_sockets.tls.v3.TlsSessionTicketKeys
	(*CertificateValidationContext)(nil), // 8: envoy.extensions.transport_sockets.tls.v3.CertificateValidationContext
}
var file_envoy_extensions_transport_sockets_tls_v3_secret_proto_depIdxs = []int32{
	3, // 0: envoy.extensions.transport_sockets.tls.v3.GenericSecret.secret:type_name -> envoy.config.core.v3.DataSource
	4, // 1: envoy.extensions.transport_sockets.tls.v3.SdsSecretConfig.sds_resource_locator:type_name -> udpa.core.v1.ResourceLocator
	5, // 2: envoy.extensions.transport_sockets.tls.v3.SdsSecretConfig.sds_config:type_name -> envoy.config.core.v3.ConfigSource
	6, // 3: envoy.extensions.transport_sockets.tls.v3.Secret.tls_certificate:type_name -> envoy.extensions.transport_sockets.tls.v3.TlsCertificate
	7, // 4: envoy.extensions.transport_sockets.tls.v3.Secret.session_ticket_keys:type_name -> envoy.extensions.transport_sockets.tls.v3.TlsSessionTicketKeys
	8, // 5: envoy.extensions.transport_sockets.tls.v3.Secret.validation_context:type_name -> envoy.extensions.transport_sockets.tls.v3.CertificateValidationContext
	0, // 6: envoy.extensions.transport_sockets.tls.v3.Secret.generic_secret:type_name -> envoy.extensions.transport_sockets.tls.v3.GenericSecret
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_envoy_extensions_transport_sockets_tls_v3_secret_proto_init() }
func file_envoy_extensions_transport_sockets_tls_v3_secret_proto_init() {
	if File_envoy_extensions_transport_sockets_tls_v3_secret_proto != nil {
		return
	}
	file_envoy_extensions_transport_sockets_tls_v3_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericSecret); i {
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
		file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SdsSecretConfig); i {
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
		file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
	file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Secret_TlsCertificate)(nil),
		(*Secret_SessionTicketKeys)(nil),
		(*Secret_ValidationContext)(nil),
		(*Secret_GenericSecret)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_transport_sockets_tls_v3_secret_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_transport_sockets_tls_v3_secret_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_transport_sockets_tls_v3_secret_proto_msgTypes,
	}.Build()
	File_envoy_extensions_transport_sockets_tls_v3_secret_proto = out.File
	file_envoy_extensions_transport_sockets_tls_v3_secret_proto_rawDesc = nil
	file_envoy_extensions_transport_sockets_tls_v3_secret_proto_goTypes = nil
	file_envoy_extensions_transport_sockets_tls_v3_secret_proto_depIdxs = nil
}
