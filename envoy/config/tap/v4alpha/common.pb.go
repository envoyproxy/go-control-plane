// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/tap/v4alpha/common.proto

package envoy_config_tap_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/config/core/v4alpha"
	v4alpha1 "github.com/envoyproxy/go-control-plane/envoy/config/route/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type OutputSink_Format int32

const (
	OutputSink_JSON_BODY_AS_BYTES            OutputSink_Format = 0
	OutputSink_JSON_BODY_AS_STRING           OutputSink_Format = 1
	OutputSink_PROTO_BINARY                  OutputSink_Format = 2
	OutputSink_PROTO_BINARY_LENGTH_DELIMITED OutputSink_Format = 3
	OutputSink_PROTO_TEXT                    OutputSink_Format = 4
)

var OutputSink_Format_name = map[int32]string{
	0: "JSON_BODY_AS_BYTES",
	1: "JSON_BODY_AS_STRING",
	2: "PROTO_BINARY",
	3: "PROTO_BINARY_LENGTH_DELIMITED",
	4: "PROTO_TEXT",
}

var OutputSink_Format_value = map[string]int32{
	"JSON_BODY_AS_BYTES":            0,
	"JSON_BODY_AS_STRING":           1,
	"PROTO_BINARY":                  2,
	"PROTO_BINARY_LENGTH_DELIMITED": 3,
	"PROTO_TEXT":                    4,
}

func (x OutputSink_Format) String() string {
	return proto.EnumName(OutputSink_Format_name, int32(x))
}

func (OutputSink_Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0226aafcd28718a, []int{4, 0}
}

type TapConfig struct {
	MatchConfig          *MatchPredicate                   `protobuf:"bytes,1,opt,name=match_config,json=matchConfig,proto3" json:"match_config,omitempty"`
	OutputConfig         *OutputConfig                     `protobuf:"bytes,2,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	TapEnabled           *v4alpha.RuntimeFractionalPercent `protobuf:"bytes,3,opt,name=tap_enabled,json=tapEnabled,proto3" json:"tap_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TapConfig) Reset()         { *m = TapConfig{} }
func (m *TapConfig) String() string { return proto.CompactTextString(m) }
func (*TapConfig) ProtoMessage()    {}
func (*TapConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0226aafcd28718a, []int{0}
}

func (m *TapConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapConfig.Unmarshal(m, b)
}
func (m *TapConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapConfig.Marshal(b, m, deterministic)
}
func (m *TapConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapConfig.Merge(m, src)
}
func (m *TapConfig) XXX_Size() int {
	return xxx_messageInfo_TapConfig.Size(m)
}
func (m *TapConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TapConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TapConfig proto.InternalMessageInfo

func (m *TapConfig) GetMatchConfig() *MatchPredicate {
	if m != nil {
		return m.MatchConfig
	}
	return nil
}

func (m *TapConfig) GetOutputConfig() *OutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

func (m *TapConfig) GetTapEnabled() *v4alpha.RuntimeFractionalPercent {
	if m != nil {
		return m.TapEnabled
	}
	return nil
}

type MatchPredicate struct {
	// Types that are valid to be assigned to Rule:
	//	*MatchPredicate_OrMatch
	//	*MatchPredicate_AndMatch
	//	*MatchPredicate_NotMatch
	//	*MatchPredicate_AnyMatch
	//	*MatchPredicate_HttpRequestHeadersMatch
	//	*MatchPredicate_HttpRequestTrailersMatch
	//	*MatchPredicate_HttpResponseHeadersMatch
	//	*MatchPredicate_HttpResponseTrailersMatch
	Rule                 isMatchPredicate_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MatchPredicate) Reset()         { *m = MatchPredicate{} }
func (m *MatchPredicate) String() string { return proto.CompactTextString(m) }
func (*MatchPredicate) ProtoMessage()    {}
func (*MatchPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0226aafcd28718a, []int{1}
}

func (m *MatchPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchPredicate.Unmarshal(m, b)
}
func (m *MatchPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchPredicate.Marshal(b, m, deterministic)
}
func (m *MatchPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchPredicate.Merge(m, src)
}
func (m *MatchPredicate) XXX_Size() int {
	return xxx_messageInfo_MatchPredicate.Size(m)
}
func (m *MatchPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_MatchPredicate proto.InternalMessageInfo

type isMatchPredicate_Rule interface {
	isMatchPredicate_Rule()
}

type MatchPredicate_OrMatch struct {
	OrMatch *MatchPredicate_MatchSet `protobuf:"bytes,1,opt,name=or_match,json=orMatch,proto3,oneof"`
}

type MatchPredicate_AndMatch struct {
	AndMatch *MatchPredicate_MatchSet `protobuf:"bytes,2,opt,name=and_match,json=andMatch,proto3,oneof"`
}

type MatchPredicate_NotMatch struct {
	NotMatch *MatchPredicate `protobuf:"bytes,3,opt,name=not_match,json=notMatch,proto3,oneof"`
}

type MatchPredicate_AnyMatch struct {
	AnyMatch bool `protobuf:"varint,4,opt,name=any_match,json=anyMatch,proto3,oneof"`
}

type MatchPredicate_HttpRequestHeadersMatch struct {
	HttpRequestHeadersMatch *HttpHeadersMatch `protobuf:"bytes,5,opt,name=http_request_headers_match,json=httpRequestHeadersMatch,proto3,oneof"`
}

type MatchPredicate_HttpRequestTrailersMatch struct {
	HttpRequestTrailersMatch *HttpHeadersMatch `protobuf:"bytes,6,opt,name=http_request_trailers_match,json=httpRequestTrailersMatch,proto3,oneof"`
}

type MatchPredicate_HttpResponseHeadersMatch struct {
	HttpResponseHeadersMatch *HttpHeadersMatch `protobuf:"bytes,7,opt,name=http_response_headers_match,json=httpResponseHeadersMatch,proto3,oneof"`
}

type MatchPredicate_HttpResponseTrailersMatch struct {
	HttpResponseTrailersMatch *HttpHeadersMatch `protobuf:"bytes,8,opt,name=http_response_trailers_match,json=httpResponseTrailersMatch,proto3,oneof"`
}

func (*MatchPredicate_OrMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_AndMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_NotMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_AnyMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpRequestHeadersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpRequestTrailersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpResponseHeadersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpResponseTrailersMatch) isMatchPredicate_Rule() {}

func (m *MatchPredicate) GetRule() isMatchPredicate_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *MatchPredicate) GetOrMatch() *MatchPredicate_MatchSet {
	if x, ok := m.GetRule().(*MatchPredicate_OrMatch); ok {
		return x.OrMatch
	}
	return nil
}

func (m *MatchPredicate) GetAndMatch() *MatchPredicate_MatchSet {
	if x, ok := m.GetRule().(*MatchPredicate_AndMatch); ok {
		return x.AndMatch
	}
	return nil
}

func (m *MatchPredicate) GetNotMatch() *MatchPredicate {
	if x, ok := m.GetRule().(*MatchPredicate_NotMatch); ok {
		return x.NotMatch
	}
	return nil
}

func (m *MatchPredicate) GetAnyMatch() bool {
	if x, ok := m.GetRule().(*MatchPredicate_AnyMatch); ok {
		return x.AnyMatch
	}
	return false
}

func (m *MatchPredicate) GetHttpRequestHeadersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpRequestHeadersMatch); ok {
		return x.HttpRequestHeadersMatch
	}
	return nil
}

func (m *MatchPredicate) GetHttpRequestTrailersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpRequestTrailersMatch); ok {
		return x.HttpRequestTrailersMatch
	}
	return nil
}

func (m *MatchPredicate) GetHttpResponseHeadersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpResponseHeadersMatch); ok {
		return x.HttpResponseHeadersMatch
	}
	return nil
}

func (m *MatchPredicate) GetHttpResponseTrailersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpResponseTrailersMatch); ok {
		return x.HttpResponseTrailersMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MatchPredicate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MatchPredicate_OrMatch)(nil),
		(*MatchPredicate_AndMatch)(nil),
		(*MatchPredicate_NotMatch)(nil),
		(*MatchPredicate_AnyMatch)(nil),
		(*MatchPredicate_HttpRequestHeadersMatch)(nil),
		(*MatchPredicate_HttpRequestTrailersMatch)(nil),
		(*MatchPredicate_HttpResponseHeadersMatch)(nil),
		(*MatchPredicate_HttpResponseTrailersMatch)(nil),
	}
}

type MatchPredicate_MatchSet struct {
	Rules                []*MatchPredicate `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MatchPredicate_MatchSet) Reset()         { *m = MatchPredicate_MatchSet{} }
func (m *MatchPredicate_MatchSet) String() string { return proto.CompactTextString(m) }
func (*MatchPredicate_MatchSet) ProtoMessage()    {}
func (*MatchPredicate_MatchSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0226aafcd28718a, []int{1, 0}
}

func (m *MatchPredicate_MatchSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchPredicate_MatchSet.Unmarshal(m, b)
}
func (m *MatchPredicate_MatchSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchPredicate_MatchSet.Marshal(b, m, deterministic)
}
func (m *MatchPredicate_MatchSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchPredicate_MatchSet.Merge(m, src)
}
func (m *MatchPredicate_MatchSet) XXX_Size() int {
	return xxx_messageInfo_MatchPredicate_MatchSet.Size(m)
}
func (m *MatchPredicate_MatchSet) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchPredicate_MatchSet.DiscardUnknown(m)
}

var xxx_messageInfo_MatchPredicate_MatchSet proto.InternalMessageInfo

func (m *MatchPredicate_MatchSet) GetRules() []*MatchPredicate {
	if m != nil {
		return m.Rules
	}
	return nil
}

type HttpHeadersMatch struct {
	Headers              []*v4alpha1.HeaderMatcher `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *HttpHeadersMatch) Reset()         { *m = HttpHeadersMatch{} }
func (m *HttpHeadersMatch) String() string { return proto.CompactTextString(m) }
func (*HttpHeadersMatch) ProtoMessage()    {}
func (*HttpHeadersMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0226aafcd28718a, []int{2}
}

func (m *HttpHeadersMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpHeadersMatch.Unmarshal(m, b)
}
func (m *HttpHeadersMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpHeadersMatch.Marshal(b, m, deterministic)
}
func (m *HttpHeadersMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpHeadersMatch.Merge(m, src)
}
func (m *HttpHeadersMatch) XXX_Size() int {
	return xxx_messageInfo_HttpHeadersMatch.Size(m)
}
func (m *HttpHeadersMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpHeadersMatch.DiscardUnknown(m)
}

var xxx_messageInfo_HttpHeadersMatch proto.InternalMessageInfo

func (m *HttpHeadersMatch) GetHeaders() []*v4alpha1.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

type OutputConfig struct {
	Sinks                []*OutputSink         `protobuf:"bytes,1,rep,name=sinks,proto3" json:"sinks,omitempty"`
	MaxBufferedRxBytes   *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_buffered_rx_bytes,json=maxBufferedRxBytes,proto3" json:"max_buffered_rx_bytes,omitempty"`
	MaxBufferedTxBytes   *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=max_buffered_tx_bytes,json=maxBufferedTxBytes,proto3" json:"max_buffered_tx_bytes,omitempty"`
	Streaming            bool                  `protobuf:"varint,4,opt,name=streaming,proto3" json:"streaming,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OutputConfig) Reset()         { *m = OutputConfig{} }
func (m *OutputConfig) String() string { return proto.CompactTextString(m) }
func (*OutputConfig) ProtoMessage()    {}
func (*OutputConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0226aafcd28718a, []int{3}
}

func (m *OutputConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputConfig.Unmarshal(m, b)
}
func (m *OutputConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputConfig.Marshal(b, m, deterministic)
}
func (m *OutputConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputConfig.Merge(m, src)
}
func (m *OutputConfig) XXX_Size() int {
	return xxx_messageInfo_OutputConfig.Size(m)
}
func (m *OutputConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OutputConfig proto.InternalMessageInfo

func (m *OutputConfig) GetSinks() []*OutputSink {
	if m != nil {
		return m.Sinks
	}
	return nil
}

func (m *OutputConfig) GetMaxBufferedRxBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxBufferedRxBytes
	}
	return nil
}

func (m *OutputConfig) GetMaxBufferedTxBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxBufferedTxBytes
	}
	return nil
}

func (m *OutputConfig) GetStreaming() bool {
	if m != nil {
		return m.Streaming
	}
	return false
}

type OutputSink struct {
	Format OutputSink_Format `protobuf:"varint,1,opt,name=format,proto3,enum=envoy.config.tap.v4alpha.OutputSink_Format" json:"format,omitempty"`
	// Types that are valid to be assigned to OutputSinkType:
	//	*OutputSink_StreamingAdmin
	//	*OutputSink_FilePerTap
	//	*OutputSink_StreamingGrpc
	OutputSinkType       isOutputSink_OutputSinkType `protobuf_oneof:"output_sink_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *OutputSink) Reset()         { *m = OutputSink{} }
func (m *OutputSink) String() string { return proto.CompactTextString(m) }
func (*OutputSink) ProtoMessage()    {}
func (*OutputSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0226aafcd28718a, []int{4}
}

func (m *OutputSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputSink.Unmarshal(m, b)
}
func (m *OutputSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputSink.Marshal(b, m, deterministic)
}
func (m *OutputSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputSink.Merge(m, src)
}
func (m *OutputSink) XXX_Size() int {
	return xxx_messageInfo_OutputSink.Size(m)
}
func (m *OutputSink) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputSink.DiscardUnknown(m)
}

var xxx_messageInfo_OutputSink proto.InternalMessageInfo

func (m *OutputSink) GetFormat() OutputSink_Format {
	if m != nil {
		return m.Format
	}
	return OutputSink_JSON_BODY_AS_BYTES
}

type isOutputSink_OutputSinkType interface {
	isOutputSink_OutputSinkType()
}

type OutputSink_StreamingAdmin struct {
	StreamingAdmin *StreamingAdminSink `protobuf:"bytes,2,opt,name=streaming_admin,json=streamingAdmin,proto3,oneof"`
}

type OutputSink_FilePerTap struct {
	FilePerTap *FilePerTapSink `protobuf:"bytes,3,opt,name=file_per_tap,json=filePerTap,proto3,oneof"`
}

type OutputSink_StreamingGrpc struct {
	StreamingGrpc *StreamingGrpcSink `protobuf:"bytes,4,opt,name=streaming_grpc,json=streamingGrpc,proto3,oneof"`
}

func (*OutputSink_StreamingAdmin) isOutputSink_OutputSinkType() {}

func (*OutputSink_FilePerTap) isOutputSink_OutputSinkType() {}

func (*OutputSink_StreamingGrpc) isOutputSink_OutputSinkType() {}

func (m *OutputSink) GetOutputSinkType() isOutputSink_OutputSinkType {
	if m != nil {
		return m.OutputSinkType
	}
	return nil
}

func (m *OutputSink) GetStreamingAdmin() *StreamingAdminSink {
	if x, ok := m.GetOutputSinkType().(*OutputSink_StreamingAdmin); ok {
		return x.StreamingAdmin
	}
	return nil
}

func (m *OutputSink) GetFilePerTap() *FilePerTapSink {
	if x, ok := m.GetOutputSinkType().(*OutputSink_FilePerTap); ok {
		return x.FilePerTap
	}
	return nil
}

func (m *OutputSink) GetStreamingGrpc() *StreamingGrpcSink {
	if x, ok := m.GetOutputSinkType().(*OutputSink_StreamingGrpc); ok {
		return x.StreamingGrpc
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OutputSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OutputSink_StreamingAdmin)(nil),
		(*OutputSink_FilePerTap)(nil),
		(*OutputSink_StreamingGrpc)(nil),
	}
}

type StreamingAdminSink struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingAdminSink) Reset()         { *m = StreamingAdminSink{} }
func (m *StreamingAdminSink) String() string { return proto.CompactTextString(m) }
func (*StreamingAdminSink) ProtoMessage()    {}
func (*StreamingAdminSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0226aafcd28718a, []int{5}
}

func (m *StreamingAdminSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingAdminSink.Unmarshal(m, b)
}
func (m *StreamingAdminSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingAdminSink.Marshal(b, m, deterministic)
}
func (m *StreamingAdminSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingAdminSink.Merge(m, src)
}
func (m *StreamingAdminSink) XXX_Size() int {
	return xxx_messageInfo_StreamingAdminSink.Size(m)
}
func (m *StreamingAdminSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingAdminSink.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingAdminSink proto.InternalMessageInfo

type FilePerTapSink struct {
	PathPrefix           string   `protobuf:"bytes,1,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilePerTapSink) Reset()         { *m = FilePerTapSink{} }
func (m *FilePerTapSink) String() string { return proto.CompactTextString(m) }
func (*FilePerTapSink) ProtoMessage()    {}
func (*FilePerTapSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0226aafcd28718a, []int{6}
}

func (m *FilePerTapSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilePerTapSink.Unmarshal(m, b)
}
func (m *FilePerTapSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilePerTapSink.Marshal(b, m, deterministic)
}
func (m *FilePerTapSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePerTapSink.Merge(m, src)
}
func (m *FilePerTapSink) XXX_Size() int {
	return xxx_messageInfo_FilePerTapSink.Size(m)
}
func (m *FilePerTapSink) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePerTapSink.DiscardUnknown(m)
}

var xxx_messageInfo_FilePerTapSink proto.InternalMessageInfo

func (m *FilePerTapSink) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

type StreamingGrpcSink struct {
	TapId                string               `protobuf:"bytes,1,opt,name=tap_id,json=tapId,proto3" json:"tap_id,omitempty"`
	GrpcService          *v4alpha.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StreamingGrpcSink) Reset()         { *m = StreamingGrpcSink{} }
func (m *StreamingGrpcSink) String() string { return proto.CompactTextString(m) }
func (*StreamingGrpcSink) ProtoMessage()    {}
func (*StreamingGrpcSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0226aafcd28718a, []int{7}
}

func (m *StreamingGrpcSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingGrpcSink.Unmarshal(m, b)
}
func (m *StreamingGrpcSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingGrpcSink.Marshal(b, m, deterministic)
}
func (m *StreamingGrpcSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingGrpcSink.Merge(m, src)
}
func (m *StreamingGrpcSink) XXX_Size() int {
	return xxx_messageInfo_StreamingGrpcSink.Size(m)
}
func (m *StreamingGrpcSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingGrpcSink.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingGrpcSink proto.InternalMessageInfo

func (m *StreamingGrpcSink) GetTapId() string {
	if m != nil {
		return m.TapId
	}
	return ""
}

func (m *StreamingGrpcSink) GetGrpcService() *v4alpha.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.tap.v4alpha.OutputSink_Format", OutputSink_Format_name, OutputSink_Format_value)
	proto.RegisterType((*TapConfig)(nil), "envoy.config.tap.v4alpha.TapConfig")
	proto.RegisterType((*MatchPredicate)(nil), "envoy.config.tap.v4alpha.MatchPredicate")
	proto.RegisterType((*MatchPredicate_MatchSet)(nil), "envoy.config.tap.v4alpha.MatchPredicate.MatchSet")
	proto.RegisterType((*HttpHeadersMatch)(nil), "envoy.config.tap.v4alpha.HttpHeadersMatch")
	proto.RegisterType((*OutputConfig)(nil), "envoy.config.tap.v4alpha.OutputConfig")
	proto.RegisterType((*OutputSink)(nil), "envoy.config.tap.v4alpha.OutputSink")
	proto.RegisterType((*StreamingAdminSink)(nil), "envoy.config.tap.v4alpha.StreamingAdminSink")
	proto.RegisterType((*FilePerTapSink)(nil), "envoy.config.tap.v4alpha.FilePerTapSink")
	proto.RegisterType((*StreamingGrpcSink)(nil), "envoy.config.tap.v4alpha.StreamingGrpcSink")
}

func init() {
	proto.RegisterFile("envoy/config/tap/v4alpha/common.proto", fileDescriptor_e0226aafcd28718a)
}

var fileDescriptor_e0226aafcd28718a = []byte{
	// 1144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0x4f, 0x6f, 0x1a, 0x47,
	0x14, 0xc0, 0xbd, 0x80, 0x01, 0x3f, 0x88, 0x4b, 0xa7, 0x4a, 0xed, 0xba, 0x49, 0x4a, 0x90, 0xe3,
	0x38, 0x8d, 0x03, 0x8a, 0xdd, 0x13, 0x97, 0xca, 0x6b, 0x63, 0x43, 0x65, 0x1b, 0xb4, 0x6c, 0xda,
	0xfa, 0xb4, 0x1a, 0x60, 0xc0, 0x5b, 0xc3, 0xce, 0x74, 0x76, 0x70, 0xe1, 0x52, 0x55, 0x55, 0x0f,
	0x51, 0x8f, 0x39, 0x55, 0xf9, 0x0e, 0x3d, 0xf6, 0xd2, 0x4b, 0x4f, 0x95, 0x7a, 0xed, 0xb7, 0xa9,
	0x7c, 0xaa, 0x66, 0x66, 0xf9, 0xb3, 0xc6, 0xc4, 0xb1, 0x6f, 0xec, 0x9b, 0xf7, 0x7e, 0xef, 0xcf,
	0x3c, 0xde, 0x3c, 0x78, 0x42, 0xbc, 0x0b, 0x3a, 0x2c, 0x34, 0xa9, 0xd7, 0x76, 0x3b, 0x05, 0x81,
	0x59, 0xe1, 0xe2, 0x0b, 0xdc, 0x65, 0x67, 0xb8, 0xd0, 0xa4, 0xbd, 0x1e, 0xf5, 0xf2, 0x8c, 0x53,
	0x41, 0xd1, 0xaa, 0x52, 0xcb, 0x6b, 0xb5, 0xbc, 0xc0, 0x2c, 0x1f, 0xa8, 0xad, 0xad, 0x87, 0x00,
	0x4d, 0xca, 0xc9, 0x98, 0xd0, 0xc0, 0x3e, 0xd1, 0xf6, 0x6b, 0x5b, 0xf3, 0xb5, 0x3a, 0x9c, 0x35,
	0x1d, 0x9f, 0xf0, 0x0b, 0xb7, 0x39, 0xd2, 0x7e, 0x19, 0xd2, 0xe6, 0xb4, 0x2f, 0x26, 0xea, 0xea,
	0xcb, 0x69, 0xd2, 0x1e, 0xa3, 0x1e, 0xf1, 0x84, 0x1f, 0x98, 0x3c, 0xea, 0x50, 0xda, 0xe9, 0x92,
	0x82, 0xfa, 0x6a, 0xf4, 0xdb, 0x85, 0x1f, 0x38, 0x66, 0x8c, 0xf0, 0xd1, 0xf9, 0xc3, 0x7e, 0x8b,
	0xe1, 0x02, 0xf6, 0x3c, 0x2a, 0xb0, 0x70, 0xa9, 0xe7, 0x17, 0x7c, 0x81, 0x45, 0x7f, 0x74, 0xfc,
	0x78, 0xe6, 0xf8, 0x82, 0x70, 0xdf, 0xa5, 0x9e, 0xeb, 0x75, 0x02, 0x95, 0x95, 0x0b, 0xdc, 0x75,
	0x5b, 0x58, 0x86, 0x12, 0xfc, 0xd0, 0x07, 0xb9, 0x3f, 0x22, 0xb0, 0x64, 0x63, 0xb6, 0xa7, 0xa2,
	0x45, 0xaf, 0x20, 0xdd, 0xc3, 0xa2, 0x79, 0xe6, 0xe8, 0xe8, 0x57, 0x8d, 0xac, 0xb1, 0x99, 0xda,
	0xde, 0xcc, 0xcf, 0x2b, 0x60, 0xfe, 0x58, 0x6a, 0xd7, 0x38, 0x69, 0xb9, 0x4d, 0x2c, 0x88, 0x99,
	0xbc, 0x34, 0x17, 0x7f, 0x35, 0x22, 0x19, 0xc3, 0x4a, 0x29, 0xce, 0x18, 0x7b, 0x8f, 0xf6, 0x05,
	0xeb, 0x8b, 0x11, 0x37, 0xa2, 0xb8, 0x1b, 0xf3, 0xb9, 0x55, 0xa5, 0xae, 0xcd, 0xa7, 0xa8, 0x69,
	0x3a, 0x25, 0x47, 0x36, 0xa4, 0x04, 0x66, 0x0e, 0xf1, 0x70, 0xa3, 0x4b, 0x5a, 0xab, 0x51, 0x05,
	0xdd, 0x09, 0x43, 0xe5, 0x6d, 0x8d, 0xa9, 0x56, 0xdf, 0x13, 0x6e, 0x8f, 0x1c, 0x70, 0xdc, 0x94,
	0x75, 0xc2, 0xdd, 0x1a, 0xe1, 0x4d, 0xe2, 0x09, 0x0b, 0x04, 0x66, 0x25, 0x8d, 0x29, 0xae, 0xbf,
	0xfd, 0xfb, 0xf5, 0xa3, 0xcf, 0xe0, 0xe1, 0x6c, 0x6c, 0x3b, 0xf9, 0x71, 0xa5, 0x72, 0x6f, 0x13,
	0xb0, 0x1c, 0x4e, 0x1e, 0x9d, 0x40, 0x92, 0x72, 0x47, 0xe5, 0x1d, 0x14, 0xee, 0xe5, 0xfb, 0x16,
	0x4e, 0x7f, 0xd6, 0x89, 0x28, 0x2f, 0x58, 0x09, 0xca, 0xd5, 0x17, 0xaa, 0xc1, 0x12, 0xf6, 0x5a,
	0x01, 0x30, 0x72, 0x77, 0x60, 0x12, 0x7b, 0x2d, 0x4d, 0x3c, 0x84, 0x25, 0x8f, 0x8a, 0x80, 0x18,
	0xbd, 0xdd, 0xdd, 0x4a, 0x90, 0x47, 0x85, 0x06, 0x6d, 0xc8, 0xd0, 0x86, 0x01, 0x28, 0x96, 0x35,
	0x36, 0x93, 0x66, 0xe2, 0xd2, 0x8c, 0x7d, 0x17, 0x49, 0x1a, 0xda, 0xe1, 0x50, 0xeb, 0xb9, 0xb0,
	0x76, 0x26, 0x04, 0x73, 0x38, 0xf9, 0xbe, 0x4f, 0x7c, 0xe1, 0x9c, 0x11, 0xdc, 0x22, 0xdc, 0x0f,
	0x0c, 0x17, 0x55, 0x04, 0x9f, 0xcf, 0x8f, 0xa0, 0x2c, 0x04, 0x2b, 0x6b, 0x13, 0xc5, 0x2b, 0x2f,
	0x58, 0x2b, 0x92, 0x67, 0x69, 0xdc, 0xf4, 0x11, 0x3a, 0x87, 0x4f, 0x43, 0xae, 0x04, 0xc7, 0x6e,
	0x77, 0xe2, 0x2b, 0x7e, 0x07, 0x5f, 0xab, 0x53, 0xbe, 0xec, 0x00, 0x77, 0xd5, 0x99, 0xcf, 0xa8,
	0xe7, 0x93, 0x2b, 0x89, 0x25, 0xee, 0xee, 0x4c, 0xf3, 0x42, 0x99, 0xf5, 0xe0, 0x41, 0xd8, 0xd9,
	0x95, 0xd4, 0x92, 0x77, 0xf0, 0xf6, 0xc9, 0xb4, 0xb7, 0x50, 0x6e, 0x6b, 0xaf, 0x0d, 0x48, 0x8e,
	0xba, 0x07, 0x95, 0x61, 0x91, 0xf7, 0xbb, 0xc4, 0x5f, 0x35, 0xb2, 0xd1, 0x5b, 0x4f, 0x82, 0x37,
	0x46, 0x24, 0x19, 0xb1, 0x34, 0xa0, 0xb8, 0x2d, 0xff, 0x56, 0x2f, 0xe0, 0xf9, 0x75, 0x7f, 0xab,
	0x39, 0xbd, 0x5b, 0x7c, 0x26, 0x6d, 0xd6, 0x21, 0x77, 0xb3, 0x8d, 0x99, 0x82, 0x98, 0xf4, 0x83,
	0xa2, 0xff, 0x99, 0x46, 0xee, 0x17, 0x03, 0x32, 0x57, 0x93, 0x46, 0x7b, 0x90, 0x08, 0x6e, 0x29,
	0x48, 0xe6, 0x59, 0x38, 0x19, 0x35, 0x9b, 0x27, 0x35, 0x53, 0xaa, 0xca, 0x92, 0x70, 0x6b, 0x64,
	0x59, 0x7c, 0x2e, 0x23, 0xda, 0x80, 0xf5, 0xeb, 0x22, 0xba, 0xea, 0x31, 0xf7, 0x57, 0x04, 0xd2,
	0xd3, 0x83, 0x0c, 0x1d, 0xc0, 0xa2, 0xef, 0x7a, 0xe7, 0xa3, 0x00, 0xd6, 0x6f, 0x9a, 0x7f, 0x75,
	0xd7, 0x3b, 0x37, 0xe1, 0xd2, 0x4c, 0xbc, 0x31, 0x62, 0x49, 0x23, 0x63, 0x58, 0xda, 0x1c, 0x55,
	0xe1, 0x7e, 0x0f, 0x0f, 0x9c, 0x46, 0xbf, 0xdd, 0x26, 0x9c, 0xb4, 0x1c, 0x3e, 0x70, 0x1a, 0x43,
	0x41, 0xfc, 0x60, 0x4a, 0x3c, 0xc8, 0xeb, 0xf7, 0x24, 0x3f, 0x7a, 0x4f, 0xf2, 0xaf, 0x2a, 0x9e,
	0xd8, 0xd9, 0xfe, 0x1a, 0x77, 0xfb, 0xc4, 0x42, 0x3d, 0x3c, 0x30, 0x03, 0x4b, 0x6b, 0x60, 0x4a,
	0xbb, 0x19, 0xa0, 0x18, 0x01, 0xa3, 0xb7, 0x04, 0xda, 0x01, 0xf0, 0x01, 0x2c, 0xf9, 0x82, 0x13,
	0xdc, 0x73, 0xbd, 0x8e, 0x1e, 0x10, 0xd6, 0x44, 0x50, 0x7c, 0x2a, 0xab, 0x98, 0x83, 0xec, 0x75,
	0x55, 0x9c, 0x2e, 0x58, 0xee, 0xb7, 0x18, 0xc0, 0xa4, 0x14, 0xe8, 0x18, 0xe2, 0x6d, 0xca, 0x7b,
	0x58, 0xa8, 0xf9, 0xba, 0xbc, 0xfd, 0xfc, 0x7d, 0x0a, 0x98, 0x3f, 0x50, 0x26, 0xaa, 0x23, 0x7f,
	0x56, 0xaf, 0x48, 0x00, 0x41, 0xdf, 0xc0, 0x07, 0xe3, 0x98, 0x1c, 0xdc, 0xea, 0xb9, 0x5e, 0x50,
	0xc0, 0xad, 0xf9, 0xdc, 0xfa, 0xc8, 0x60, 0x57, 0xea, 0x4b, 0x7e, 0x79, 0xc1, 0x5a, 0xf6, 0x43,
	0x52, 0x74, 0x04, 0xe9, 0xb6, 0xdb, 0x25, 0x0e, 0x23, 0xdc, 0x11, 0x98, 0xdd, 0x3c, 0x6a, 0x0f,
	0xdc, 0x2e, 0xa9, 0x11, 0x6e, 0x63, 0x16, 0x10, 0xa1, 0x3d, 0x96, 0x20, 0x1b, 0x26, 0x7c, 0x47,
	0x2e, 0x1c, 0xaa, 0xa0, 0xa9, 0x77, 0x65, 0x3f, 0x8e, 0xf2, 0x90, 0xb3, 0x66, 0x80, 0xbc, 0xe7,
	0x4f, 0x0b, 0x73, 0x3f, 0x42, 0x5c, 0x17, 0x06, 0x7d, 0x0c, 0xe8, 0xab, 0x7a, 0xf5, 0xc4, 0x31,
	0xab, 0xfb, 0xa7, 0xce, 0x6e, 0xdd, 0x31, 0x4f, 0xed, 0x52, 0x3d, 0xb3, 0x80, 0x56, 0xe0, 0xa3,
	0x90, 0xbc, 0x6e, 0x5b, 0x95, 0x93, 0xc3, 0x8c, 0x81, 0x32, 0x90, 0xae, 0x59, 0x55, 0xbb, 0xea,
	0x98, 0x95, 0x93, 0x5d, 0xeb, 0x34, 0x13, 0x41, 0x8f, 0xe1, 0xe1, 0xb4, 0xc4, 0x39, 0x2a, 0x9d,
	0x1c, 0xda, 0x65, 0x67, 0xbf, 0x74, 0x54, 0x39, 0xae, 0xd8, 0xa5, 0xfd, 0x4c, 0x14, 0x2d, 0x03,
	0x68, 0x15, 0xbb, 0xf4, 0xad, 0x9d, 0x89, 0x15, 0x9f, 0xc8, 0x1e, 0xc8, 0xc2, 0xa3, 0xf9, 0x3d,
	0xa0, 0xba, 0x7f, 0x05, 0x32, 0xc1, 0xea, 0x20, 0x5b, 0xdf, 0x11, 0x43, 0x16, 0xfc, 0xc7, 0xf7,
	0x00, 0xcd, 0xde, 0x45, 0xf1, 0x85, 0xa4, 0x6e, 0xc2, 0xc6, 0x75, 0xd4, 0x59, 0xf5, 0x1c, 0x81,
	0xe5, 0x70, 0xe9, 0xd1, 0x26, 0xa4, 0x18, 0x16, 0x67, 0x0e, 0xe3, 0xa4, 0xed, 0x0e, 0x54, 0x9f,
	0x2d, 0xa9, 0xb7, 0x8d, 0x47, 0xb2, 0x86, 0x05, 0xf2, 0xac, 0xa6, 0x8e, 0xde, 0x39, 0x9c, 0xc2,
	0xd0, 0xdc, 0xef, 0x06, 0x7c, 0x38, 0x73, 0x25, 0xe8, 0x3e, 0xc4, 0xe5, 0xfa, 0xe2, 0xb6, 0xb4,
	0x17, 0x6b, 0x51, 0x60, 0x56, 0x69, 0xa1, 0x3a, 0xa4, 0xa7, 0xb7, 0xca, 0xeb, 0x77, 0xa5, 0xd0,
	0x5a, 0xa3, 0x88, 0x5a, 0x7b, 0x7a, 0x03, 0xeb, 0x4c, 0xc4, 0xc5, 0x2d, 0x19, 0xec, 0xd3, 0x60,
	0x61, 0x9e, 0x57, 0x97, 0x51, 0x64, 0xe6, 0x97, 0x7f, 0xfe, 0xf4, 0xcf, 0xbf, 0xf1, 0x48, 0x26,
	0x0a, 0x1b, 0x2e, 0xd5, 0x8e, 0x19, 0xa7, 0x83, 0xe1, 0xdc, 0x86, 0x33, 0x53, 0x7b, 0x6a, 0xe1,
	0xae, 0xc9, 0xf9, 0x50, 0x33, 0x1a, 0x71, 0x35, 0x28, 0x76, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0x90, 0xf7, 0x91, 0x9c, 0xa1, 0x0b, 0x00, 0x00,
}
