// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v4alpha/http_tracer.proto

package envoy_config_trace_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/struct"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Tracing struct {
	Http                 *Tracing_Http `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Tracing) Reset()         { *m = Tracing{} }
func (m *Tracing) String() string { return proto.CompactTextString(m) }
func (*Tracing) ProtoMessage()    {}
func (*Tracing) Descriptor() ([]byte, []int) {
	return fileDescriptor_156d1cd230b110a3, []int{0}
}

func (m *Tracing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing.Unmarshal(m, b)
}
func (m *Tracing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing.Marshal(b, m, deterministic)
}
func (m *Tracing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing.Merge(m, src)
}
func (m *Tracing) XXX_Size() int {
	return xxx_messageInfo_Tracing.Size(m)
}
func (m *Tracing) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing proto.InternalMessageInfo

func (m *Tracing) GetHttp() *Tracing_Http {
	if m != nil {
		return m.Http
	}
	return nil
}

type Tracing_Http struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*Tracing_Http_TypedConfig
	ConfigType           isTracing_Http_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Tracing_Http) Reset()         { *m = Tracing_Http{} }
func (m *Tracing_Http) String() string { return proto.CompactTextString(m) }
func (*Tracing_Http) ProtoMessage()    {}
func (*Tracing_Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_156d1cd230b110a3, []int{0, 0}
}

func (m *Tracing_Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing_Http.Unmarshal(m, b)
}
func (m *Tracing_Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing_Http.Marshal(b, m, deterministic)
}
func (m *Tracing_Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing_Http.Merge(m, src)
}
func (m *Tracing_Http) XXX_Size() int {
	return xxx_messageInfo_Tracing_Http.Size(m)
}
func (m *Tracing_Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing_Http proto.InternalMessageInfo

func (m *Tracing_Http) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTracing_Http_ConfigType interface {
	isTracing_Http_ConfigType()
}

type Tracing_Http_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Tracing_Http_TypedConfig) isTracing_Http_ConfigType() {}

func (m *Tracing_Http) GetConfigType() isTracing_Http_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *Tracing_Http) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Tracing_Http_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Tracing_Http) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Tracing_Http_TypedConfig)(nil),
	}
}

func init() {
	proto.RegisterType((*Tracing)(nil), "envoy.config.trace.v4alpha.Tracing")
	proto.RegisterType((*Tracing_Http)(nil), "envoy.config.trace.v4alpha.Tracing.Http")
}

func init() {
	proto.RegisterFile("envoy/config/trace/v4alpha/http_tracer.proto", fileDescriptor_156d1cd230b110a3)
}

var fileDescriptor_156d1cd230b110a3 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x6f, 0xd2, 0xd0, 0xde, 0x3b, 0xbd, 0x62, 0x09, 0x82, 0x35, 0x5a, 0xad, 0xa5, 0x8b,
	0x0a, 0x32, 0x03, 0xd6, 0x8d, 0xc5, 0x8d, 0x11, 0xa1, 0xb8, 0x2a, 0xa1, 0xfb, 0x32, 0x4d, 0xd3,
	0x74, 0xa0, 0xce, 0x0c, 0x93, 0x93, 0x60, 0x76, 0x2e, 0x7d, 0x02, 0x17, 0x3e, 0x80, 0x0f, 0xe1,
	0x5e, 0x70, 0xeb, 0xeb, 0xb8, 0x92, 0xcc, 0xa4, 0x0b, 0xff, 0x75, 0x97, 0x9c, 0xef, 0x77, 0xbe,
	0x73, 0xbe, 0xc3, 0xa0, 0xe3, 0x88, 0x67, 0x22, 0x27, 0xa1, 0xe0, 0x73, 0x16, 0x13, 0x50, 0x34,
	0x8c, 0x48, 0x76, 0x4a, 0x97, 0x72, 0x41, 0xc9, 0x02, 0x40, 0x4e, 0x74, 0x49, 0x61, 0xa9, 0x04,
	0x08, 0xd7, 0xd3, 0x34, 0x36, 0x34, 0xd6, 0x12, 0x2e, 0x69, 0x6f, 0x27, 0x16, 0x22, 0x5e, 0x46,
	0x44, 0x93, 0xd3, 0x74, 0x4e, 0x28, 0xcf, 0x4d, 0x9b, 0xb7, 0xf7, 0x55, 0x4a, 0x40, 0xa5, 0x21,
	0x94, 0x6a, 0x2b, 0x9d, 0x49, 0x4a, 0x28, 0xe7, 0x02, 0x28, 0x30, 0xc1, 0x13, 0x92, 0x00, 0x85,
	0x34, 0x29, 0xe5, 0xc3, 0x6f, 0x72, 0x16, 0xa9, 0x84, 0x09, 0xce, 0x78, 0x5c, 0x22, 0xdb, 0x19,
	0x5d, 0xb2, 0x19, 0x85, 0x88, 0xac, 0x3e, 0x8c, 0xd0, 0x79, 0xb0, 0x51, 0x6d, 0xac, 0x68, 0xc8,
	0x78, 0xec, 0x9e, 0x23, 0xa7, 0x08, 0xd4, 0xb4, 0xda, 0x56, 0xaf, 0x7e, 0xd2, 0xc3, 0xbf, 0x47,
	0xc1, 0x65, 0x0b, 0x1e, 0x02, 0xc8, 0x40, 0x77, 0x79, 0x4f, 0x16, 0x72, 0x8a, 0x5f, 0x77, 0x17,
	0x39, 0x9c, 0xde, 0x44, 0xda, 0xe6, 0x9f, 0x5f, 0x7b, 0xf7, 0x1d, 0x65, 0xb7, 0xad, 0x40, 0x17,
	0xdd, 0x33, 0xf4, 0x1f, 0x72, 0x19, 0xcd, 0x26, 0xc6, 0xb6, 0x59, 0xd1, 0xb3, 0xb6, 0xb0, 0xc9,
	0x8f, 0x57, 0xf9, 0xf1, 0x05, 0xcf, 0x87, 0x7f, 0x82, 0xba, 0x66, 0x2f, 0x35, 0x3a, 0x38, 0x7a,
	0x7c, 0xb9, 0xdf, 0xef, 0xa2, 0xce, 0x4f, 0x6b, 0xf5, 0x3f, 0x6d, 0xe4, 0x6f, 0xa0, 0xba, 0xd1,
	0x27, 0x85, 0xc1, 0xb5, 0xf3, 0xd7, 0x6e, 0x54, 0x82, 0xaa, 0x29, 0x0d, 0xba, 0x85, 0xcf, 0x01,
	0x6a, 0xad, 0xf5, 0xf1, 0xaf, 0x9e, 0xef, 0x5e, 0xdf, 0xaa, 0x76, 0xa3, 0x82, 0x7a, 0x4c, 0x98,
	0x53, 0x48, 0x25, 0x6e, 0xf3, 0x35, 0x57, 0xf1, 0x37, 0x8b, 0xe1, 0x63, 0xfd, 0x1c, 0x46, 0x45,
	0x90, 0x91, 0x35, 0xad, 0xea, 0x44, 0xfd, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0xc3, 0x5d,
	0xfb, 0x46, 0x02, 0x00, 0x00,
}
