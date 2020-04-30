// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/gzip/v3/gzip.proto

package envoy_extensions_filters_http_gzip_v3

import (
	fmt "fmt"
	math "math"

	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/compressor/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT  Gzip_CompressionStrategy = 0
	Gzip_FILTERED Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN  Gzip_CompressionStrategy = 2
	Gzip_RLE      Gzip_CompressionStrategy = 3
)

var Gzip_CompressionStrategy_name = map[int32]string{
	0: "DEFAULT",
	1: "FILTERED",
	2: "HUFFMAN",
	3: "RLE",
}

var Gzip_CompressionStrategy_value = map[string]int32{
	"DEFAULT":  0,
	"FILTERED": 1,
	"HUFFMAN":  2,
	"RLE":      3,
}

func (x Gzip_CompressionStrategy) String() string {
	return proto.EnumName(Gzip_CompressionStrategy_name, int32(x))
}

func (Gzip_CompressionStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0, 0}
}

type Gzip_CompressionLevel_Enum int32

const (
	Gzip_CompressionLevel_DEFAULT Gzip_CompressionLevel_Enum = 0
	Gzip_CompressionLevel_BEST    Gzip_CompressionLevel_Enum = 1
	Gzip_CompressionLevel_SPEED   Gzip_CompressionLevel_Enum = 2
)

var Gzip_CompressionLevel_Enum_name = map[int32]string{
	0: "DEFAULT",
	1: "BEST",
	2: "SPEED",
}

var Gzip_CompressionLevel_Enum_value = map[string]int32{
	"DEFAULT": 0,
	"BEST":    1,
	"SPEED":   2,
}

func (x Gzip_CompressionLevel_Enum) String() string {
	return proto.EnumName(Gzip_CompressionLevel_Enum_name, int32(x))
}

func (Gzip_CompressionLevel_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0, 0, 0}
}

type Gzip struct {
	MemoryLevel                                     *wrappers.UInt32Value      `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel,proto3" json:"memory_level,omitempty"`
	CompressionLevel                                Gzip_CompressionLevel_Enum `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.extensions.filters.http.gzip.v3.Gzip_CompressionLevel_Enum" json:"compression_level,omitempty"`
	CompressionStrategy                             Gzip_CompressionStrategy   `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.extensions.filters.http.gzip.v3.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	WindowBits                                      *wrappers.UInt32Value      `protobuf:"bytes,9,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	Compressor                                      *v3.Compressor             `protobuf:"bytes,10,opt,name=compressor,proto3" json:"compressor,omitempty"`
	ChunkSize                                       *wrappers.UInt32Value      `protobuf:"bytes,11,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	HiddenEnvoyDeprecatedContentLength              *wrappers.UInt32Value      `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_content_length,json=hiddenEnvoyDeprecatedContentLength,proto3" json:"hidden_envoy_deprecated_content_length,omitempty"`                                             // Deprecated: Do not use.
	HiddenEnvoyDeprecatedContentType                []string                   `protobuf:"bytes,6,rep,name=hidden_envoy_deprecated_content_type,json=hiddenEnvoyDeprecatedContentType,proto3" json:"hidden_envoy_deprecated_content_type,omitempty"`                                                   // Deprecated: Do not use.
	HiddenEnvoyDeprecatedDisableOnEtagHeader        bool                       `protobuf:"varint,7,opt,name=hidden_envoy_deprecated_disable_on_etag_header,json=hiddenEnvoyDeprecatedDisableOnEtagHeader,proto3" json:"hidden_envoy_deprecated_disable_on_etag_header,omitempty"`                      // Deprecated: Do not use.
	HiddenEnvoyDeprecatedRemoveAcceptEncodingHeader bool                       `protobuf:"varint,8,opt,name=hidden_envoy_deprecated_remove_accept_encoding_header,json=hiddenEnvoyDeprecatedRemoveAcceptEncodingHeader,proto3" json:"hidden_envoy_deprecated_remove_accept_encoding_header,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral                            struct{}                   `json:"-"`
	XXX_unrecognized                                []byte                     `json:"-"`
	XXX_sizecache                                   int32                      `json:"-"`
}

func (m *Gzip) Reset()         { *m = Gzip{} }
func (m *Gzip) String() string { return proto.CompactTextString(m) }
func (*Gzip) ProtoMessage()    {}
func (*Gzip) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0}
}

func (m *Gzip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip.Unmarshal(m, b)
}
func (m *Gzip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip.Marshal(b, m, deterministic)
}
func (m *Gzip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip.Merge(m, src)
}
func (m *Gzip) XXX_Size() int {
	return xxx_messageInfo_Gzip.Size(m)
}
func (m *Gzip) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip proto.InternalMessageInfo

func (m *Gzip) GetMemoryLevel() *wrappers.UInt32Value {
	if m != nil {
		return m.MemoryLevel
	}
	return nil
}

func (m *Gzip) GetCompressionLevel() Gzip_CompressionLevel_Enum {
	if m != nil {
		return m.CompressionLevel
	}
	return Gzip_CompressionLevel_DEFAULT
}

func (m *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if m != nil {
		return m.CompressionStrategy
	}
	return Gzip_DEFAULT
}

func (m *Gzip) GetWindowBits() *wrappers.UInt32Value {
	if m != nil {
		return m.WindowBits
	}
	return nil
}

func (m *Gzip) GetCompressor() *v3.Compressor {
	if m != nil {
		return m.Compressor
	}
	return nil
}

func (m *Gzip) GetChunkSize() *wrappers.UInt32Value {
	if m != nil {
		return m.ChunkSize
	}
	return nil
}

// Deprecated: Do not use.
func (m *Gzip) GetHiddenEnvoyDeprecatedContentLength() *wrappers.UInt32Value {
	if m != nil {
		return m.HiddenEnvoyDeprecatedContentLength
	}
	return nil
}

// Deprecated: Do not use.
func (m *Gzip) GetHiddenEnvoyDeprecatedContentType() []string {
	if m != nil {
		return m.HiddenEnvoyDeprecatedContentType
	}
	return nil
}

// Deprecated: Do not use.
func (m *Gzip) GetHiddenEnvoyDeprecatedDisableOnEtagHeader() bool {
	if m != nil {
		return m.HiddenEnvoyDeprecatedDisableOnEtagHeader
	}
	return false
}

// Deprecated: Do not use.
func (m *Gzip) GetHiddenEnvoyDeprecatedRemoveAcceptEncodingHeader() bool {
	if m != nil {
		return m.HiddenEnvoyDeprecatedRemoveAcceptEncodingHeader
	}
	return false
}

type Gzip_CompressionLevel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gzip_CompressionLevel) Reset()         { *m = Gzip_CompressionLevel{} }
func (m *Gzip_CompressionLevel) String() string { return proto.CompactTextString(m) }
func (*Gzip_CompressionLevel) ProtoMessage()    {}
func (*Gzip_CompressionLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0, 0}
}

func (m *Gzip_CompressionLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip_CompressionLevel.Unmarshal(m, b)
}
func (m *Gzip_CompressionLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip_CompressionLevel.Marshal(b, m, deterministic)
}
func (m *Gzip_CompressionLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip_CompressionLevel.Merge(m, src)
}
func (m *Gzip_CompressionLevel) XXX_Size() int {
	return xxx_messageInfo_Gzip_CompressionLevel.Size(m)
}
func (m *Gzip_CompressionLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip_CompressionLevel.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip_CompressionLevel proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.extensions.filters.http.gzip.v3.Gzip_CompressionStrategy", Gzip_CompressionStrategy_name, Gzip_CompressionStrategy_value)
	proto.RegisterEnum("envoy.extensions.filters.http.gzip.v3.Gzip_CompressionLevel_Enum", Gzip_CompressionLevel_Enum_name, Gzip_CompressionLevel_Enum_value)
	proto.RegisterType((*Gzip)(nil), "envoy.extensions.filters.http.gzip.v3.Gzip")
	proto.RegisterType((*Gzip_CompressionLevel)(nil), "envoy.extensions.filters.http.gzip.v3.Gzip.CompressionLevel")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/gzip/v3/gzip.proto", fileDescriptor_75f828f0702c619c)
}

var fileDescriptor_75f828f0702c619c = []byte{
	// 722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x5f, 0x6f, 0xda, 0x56,
	0x18, 0xc6, 0x67, 0x20, 0x80, 0x0f, 0xd1, 0xe6, 0x39, 0x93, 0x66, 0x45, 0x5b, 0xc4, 0xd0, 0xb2,
	0x59, 0xd1, 0x64, 0x4f, 0xa0, 0xfd, 0x51, 0xb4, 0x69, 0xc2, 0xc1, 0x2c, 0x89, 0x68, 0x1b, 0x19,
	0xd2, 0x5e, 0x5a, 0xc6, 0x7e, 0x63, 0x8e, 0x6a, 0xce, 0xb1, 0xec, 0x83, 0x13, 0x68, 0x2f, 0x50,
	0xaf, 0xfa, 0x19, 0xfa, 0x51, 0x7a, 0x5f, 0xa9, 0x77, 0x55, 0xbf, 0x4e, 0xae, 0x2a, 0x1f, 0xf3,
	0xaf, 0x09, 0x49, 0x50, 0xaf, 0xe0, 0xe8, 0x3d, 0xcf, 0xef, 0x79, 0x7c, 0xf4, 0x3e, 0xe8, 0x77,
	0x20, 0x09, 0x1d, 0xeb, 0x70, 0xc5, 0x80, 0xc4, 0x98, 0x92, 0x58, 0xbf, 0xc0, 0x01, 0x83, 0x28,
	0xd6, 0x07, 0x8c, 0x85, 0xba, 0x3f, 0xc1, 0xa1, 0x9e, 0x34, 0xf8, 0xaf, 0x16, 0x46, 0x94, 0x51,
	0x79, 0x9f, 0x2b, 0xb4, 0xa5, 0x42, 0x9b, 0x29, 0xb4, 0x54, 0xa1, 0xf1, 0x9b, 0x49, 0x63, 0xf7,
	0x9f, 0xfb, 0xc1, 0x2e, 0x1d, 0x86, 0x11, 0xc4, 0x31, 0x8d, 0x52, 0xfc, 0xf2, 0x94, 0x99, 0xec,
	0xee, 0xf9, 0x94, 0xfa, 0x01, 0xe8, 0xfc, 0xd4, 0x1f, 0x5d, 0xe8, 0x97, 0x91, 0x13, 0x86, 0xa9,
	0x49, 0x36, 0xff, 0x71, 0xe4, 0x85, 0x8e, 0xee, 0x10, 0x42, 0x99, 0xc3, 0x38, 0x3d, 0x66, 0x0e,
	0x1b, 0xcd, 0xc7, 0x3f, 0xdd, 0x1a, 0x27, 0x10, 0xa5, 0x29, 0x30, 0xf1, 0x67, 0x57, 0xbe, 0x4f,
	0x9c, 0x00, 0x7b, 0x0e, 0x03, 0x7d, 0xfe, 0x27, 0x1b, 0xd4, 0x3e, 0x88, 0xa8, 0xf0, 0xff, 0x04,
	0x87, 0xf2, 0x29, 0xda, 0x1e, 0xc2, 0x90, 0x46, 0x63, 0x3b, 0x80, 0x04, 0x02, 0x45, 0xa8, 0x0a,
	0x6a, 0xa5, 0xfe, 0x83, 0x96, 0x45, 0xd3, 0xe6, 0xd1, 0xb4, 0xf3, 0x13, 0xc2, 0x1a, 0xf5, 0xa7,
	0x4e, 0x30, 0x02, 0x43, 0xbc, 0x36, 0x8a, 0x07, 0x05, 0x45, 0x54, 0x05, 0xab, 0x92, 0x89, 0x3b,
	0xa9, 0x56, 0xbe, 0x42, 0xdf, 0xce, 0xbf, 0x11, 0x53, 0x32, 0x03, 0xe6, 0xab, 0x82, 0xfa, 0x75,
	0xbd, 0xa9, 0x6d, 0xf4, 0xa0, 0x5a, 0x9a, 0x49, 0x3b, 0x5a, 0x42, 0x38, 0x58, 0x33, 0xc9, 0x68,
	0x68, 0x94, 0xaf, 0x8d, 0xad, 0x57, 0x42, 0x4e, 0x12, 0x2c, 0xc9, 0xbd, 0x71, 0x41, 0x7e, 0x89,
	0xbe, 0x5b, 0x75, 0x8e, 0x59, 0xe4, 0x30, 0xf0, 0xc7, 0x4a, 0x81, 0x9b, 0xff, 0xf7, 0x85, 0xe6,
	0xdd, 0x19, 0x66, 0xc5, 0x7a, 0xc7, 0xbd, 0x3d, 0x96, 0x8f, 0x51, 0xe5, 0x12, 0x13, 0x8f, 0x5e,
	0xda, 0x7d, 0xcc, 0x62, 0x45, 0xdc, 0xfc, 0x09, 0xbf, 0x51, 0x45, 0x0b, 0x65, 0x5a, 0x03, 0xb3,
	0x58, 0x7e, 0x86, 0xd0, 0x72, 0x4b, 0x14, 0xc4, 0x41, 0x7f, 0x3d, 0x90, 0x7e, 0x65, 0xad, 0x92,
	0xc6, 0x22, 0x3e, 0x8d, 0xac, 0x15, 0x94, 0x7c, 0x82, 0x90, 0x3b, 0x18, 0x91, 0xe7, 0x76, 0x8c,
	0x27, 0xa0, 0x54, 0x36, 0x48, 0xb8, 0x7d, 0x6d, 0x88, 0x07, 0x25, 0x65, 0x3a, 0x2d, 0xa8, 0xd3,
	0xaa, 0x25, 0x72, 0x75, 0x17, 0x4f, 0x40, 0x66, 0xe8, 0x97, 0x01, 0xf6, 0x3c, 0x20, 0x36, 0xcf,
	0x65, 0x7b, 0x10, 0x46, 0xe0, 0x3a, 0x0c, 0x3c, 0xdb, 0xa5, 0x84, 0x01, 0x61, 0x76, 0x00, 0xc4,
	0x67, 0x03, 0x25, 0xb7, 0x81, 0x4d, 0x4e, 0x11, 0xac, 0x5a, 0xc6, 0x33, 0x53, 0x5c, 0x6b, 0x41,
	0x3b, 0xca, 0x60, 0x1d, 0xce, 0x92, 0x2d, 0xf4, 0xf3, 0x43, 0xae, 0x6c, 0x1c, 0x82, 0x52, 0xac,
	0xe6, 0x55, 0x91, 0x53, 0xab, 0xf7, 0x51, 0x7b, 0xe3, 0x10, 0x64, 0x0f, 0x69, 0x77, 0x31, 0x3d,
	0x1c, 0x3b, 0xfd, 0x00, 0x6c, 0x4a, 0x6c, 0x60, 0x8e, 0x6f, 0x0f, 0xc0, 0xf1, 0x20, 0x52, 0x4a,
	0x55, 0x41, 0x2d, 0x73, 0xba, 0xba, 0x96, 0xde, 0xca, 0x64, 0x4f, 0x88, 0xc9, 0x1c, 0xff, 0x98,
	0x6b, 0xe4, 0x08, 0xfd, 0x71, 0x97, 0x4b, 0x04, 0x43, 0x9a, 0x80, 0xed, 0xb8, 0x2e, 0x84, 0xcc,
	0x06, 0xe2, 0x52, 0x0f, 0x93, 0x85, 0x59, 0x79, 0x61, 0xa6, 0xaf, 0x35, 0xb3, 0xb8, 0xba, 0xc9,
	0xc5, 0xe6, 0x4c, 0x9b, 0x79, 0xee, 0xbe, 0x40, 0xd2, 0xcd, 0x12, 0xd5, 0x54, 0x54, 0x48, 0x7b,
	0x24, 0x57, 0x50, 0xa9, 0x65, 0xb6, 0x9b, 0xe7, 0x9d, 0x9e, 0xf4, 0x95, 0x5c, 0x46, 0x05, 0xc3,
	0xec, 0xf6, 0x24, 0x41, 0x16, 0xd1, 0x56, 0xf7, 0xcc, 0x34, 0x5b, 0x52, 0xee, 0xf0, 0xdf, 0x37,
	0xef, 0x5e, 0xef, 0xfd, 0x8d, 0xfe, 0xcc, 0xf6, 0xce, 0xa5, 0xe4, 0x02, 0xfb, 0xb3, 0x9d, 0x5b,
	0x2d, 0x4c, 0x7d, 0x7d, 0x5b, 0x6b, 0x6d, 0xb4, 0xb3, 0xa6, 0x44, 0x9f, 0xfb, 0x6e, 0xa3, 0x72,
	0xfb, 0xa4, 0xd3, 0x33, 0x2d, 0xb3, 0x25, 0x09, 0xe9, 0xe8, 0xf8, 0xbc, 0xdd, 0x7e, 0xd4, 0x7c,
	0x2c, 0xe5, 0xe4, 0x12, 0xca, 0x5b, 0x1d, 0x53, 0xca, 0x1f, 0xfe, 0x96, 0xc6, 0xf8, 0x15, 0xed,
	0x6f, 0x14, 0xc3, 0x38, 0x7d, 0x3b, 0x7d, 0xff, 0xb1, 0x98, 0x93, 0x72, 0xa8, 0x81, 0x69, 0x56,
	0x99, 0x30, 0xa2, 0x57, 0xe3, 0xcd, 0xba, 0x6f, 0x88, 0x29, 0xe4, 0x2c, 0xdd, 0xd0, 0x33, 0xa1,
	0x5f, 0xe4, 0xab, 0xda, 0xf8, 0x14, 0x00, 0x00, 0xff, 0xff, 0x81, 0x9a, 0xfa, 0x77, 0x38, 0x06,
	0x00, 0x00,
}
