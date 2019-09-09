// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/core/health_check.proto

package envoy_api_v3alpha_core

import (
	fmt "fmt"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type HealthStatus int32

const (
	HealthStatus_UNKNOWN   HealthStatus = 0
	HealthStatus_HEALTHY   HealthStatus = 1
	HealthStatus_UNHEALTHY HealthStatus = 2
	HealthStatus_DRAINING  HealthStatus = 3
	HealthStatus_TIMEOUT   HealthStatus = 4
	HealthStatus_DEGRADED  HealthStatus = 5
)

var HealthStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "HEALTHY",
	2: "UNHEALTHY",
	3: "DRAINING",
	4: "TIMEOUT",
	5: "DEGRADED",
}

var HealthStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"HEALTHY":   1,
	"UNHEALTHY": 2,
	"DRAINING":  3,
	"TIMEOUT":   4,
	"DEGRADED":  5,
}

func (x HealthStatus) String() string {
	return proto.EnumName(HealthStatus_name, int32(x))
}

func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0}
}

type HealthCheck struct {
	Timeout               *duration.Duration    `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Interval              *duration.Duration    `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	InitialJitter         *duration.Duration    `protobuf:"bytes,20,opt,name=initial_jitter,json=initialJitter,proto3" json:"initial_jitter,omitempty"`
	IntervalJitter        *duration.Duration    `protobuf:"bytes,3,opt,name=interval_jitter,json=intervalJitter,proto3" json:"interval_jitter,omitempty"`
	IntervalJitterPercent uint32                `protobuf:"varint,18,opt,name=interval_jitter_percent,json=intervalJitterPercent,proto3" json:"interval_jitter_percent,omitempty"`
	UnhealthyThreshold    *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=unhealthy_threshold,json=unhealthyThreshold,proto3" json:"unhealthy_threshold,omitempty"`
	HealthyThreshold      *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=healthy_threshold,json=healthyThreshold,proto3" json:"healthy_threshold,omitempty"`
	AltPort               *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=alt_port,json=altPort,proto3" json:"alt_port,omitempty"`
	ReuseConnection       *wrappers.BoolValue   `protobuf:"bytes,7,opt,name=reuse_connection,json=reuseConnection,proto3" json:"reuse_connection,omitempty"`
	// Types that are valid to be assigned to HealthChecker:
	//	*HealthCheck_HttpHealthCheck_
	//	*HealthCheck_TcpHealthCheck_
	//	*HealthCheck_GrpcHealthCheck_
	//	*HealthCheck_CustomHealthCheck_
	HealthChecker                isHealthCheck_HealthChecker `protobuf_oneof:"health_checker"`
	NoTrafficInterval            *duration.Duration          `protobuf:"bytes,12,opt,name=no_traffic_interval,json=noTrafficInterval,proto3" json:"no_traffic_interval,omitempty"`
	UnhealthyInterval            *duration.Duration          `protobuf:"bytes,14,opt,name=unhealthy_interval,json=unhealthyInterval,proto3" json:"unhealthy_interval,omitempty"`
	UnhealthyEdgeInterval        *duration.Duration          `protobuf:"bytes,15,opt,name=unhealthy_edge_interval,json=unhealthyEdgeInterval,proto3" json:"unhealthy_edge_interval,omitempty"`
	HealthyEdgeInterval          *duration.Duration          `protobuf:"bytes,16,opt,name=healthy_edge_interval,json=healthyEdgeInterval,proto3" json:"healthy_edge_interval,omitempty"`
	EventLogPath                 string                      `protobuf:"bytes,17,opt,name=event_log_path,json=eventLogPath,proto3" json:"event_log_path,omitempty"`
	AlwaysLogHealthCheckFailures bool                        `protobuf:"varint,19,opt,name=always_log_health_check_failures,json=alwaysLogHealthCheckFailures,proto3" json:"always_log_health_check_failures,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                    `json:"-"`
	XXX_unrecognized             []byte                      `json:"-"`
	XXX_sizecache                int32                       `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *HealthCheck) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *HealthCheck) GetInitialJitter() *duration.Duration {
	if m != nil {
		return m.InitialJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitter() *duration.Duration {
	if m != nil {
		return m.IntervalJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitterPercent() uint32 {
	if m != nil {
		return m.IntervalJitterPercent
	}
	return 0
}

func (m *HealthCheck) GetUnhealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetHealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.HealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetAltPort() *wrappers.UInt32Value {
	if m != nil {
		return m.AltPort
	}
	return nil
}

func (m *HealthCheck) GetReuseConnection() *wrappers.BoolValue {
	if m != nil {
		return m.ReuseConnection
	}
	return nil
}

type isHealthCheck_HealthChecker interface {
	isHealthCheck_HealthChecker()
}

type HealthCheck_HttpHealthCheck_ struct {
	HttpHealthCheck *HealthCheck_HttpHealthCheck `protobuf:"bytes,8,opt,name=http_health_check,json=httpHealthCheck,proto3,oneof"`
}

type HealthCheck_TcpHealthCheck_ struct {
	TcpHealthCheck *HealthCheck_TcpHealthCheck `protobuf:"bytes,9,opt,name=tcp_health_check,json=tcpHealthCheck,proto3,oneof"`
}

type HealthCheck_GrpcHealthCheck_ struct {
	GrpcHealthCheck *HealthCheck_GrpcHealthCheck `protobuf:"bytes,11,opt,name=grpc_health_check,json=grpcHealthCheck,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_ struct {
	CustomHealthCheck *HealthCheck_CustomHealthCheck `protobuf:"bytes,13,opt,name=custom_health_check,json=customHealthCheck,proto3,oneof"`
}

func (*HealthCheck_HttpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_TcpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_GrpcHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_CustomHealthCheck_) isHealthCheck_HealthChecker() {}

func (m *HealthCheck) GetHealthChecker() isHealthCheck_HealthChecker {
	if m != nil {
		return m.HealthChecker
	}
	return nil
}

func (m *HealthCheck) GetHttpHealthCheck() *HealthCheck_HttpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_HttpHealthCheck_); ok {
		return x.HttpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetTcpHealthCheck() *HealthCheck_TcpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_TcpHealthCheck_); ok {
		return x.TcpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetGrpcHealthCheck() *HealthCheck_GrpcHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_GrpcHealthCheck_); ok {
		return x.GrpcHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetCustomHealthCheck() *HealthCheck_CustomHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_CustomHealthCheck_); ok {
		return x.CustomHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetNoTrafficInterval() *duration.Duration {
	if m != nil {
		return m.NoTrafficInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetHealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.HealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetEventLogPath() string {
	if m != nil {
		return m.EventLogPath
	}
	return ""
}

func (m *HealthCheck) GetAlwaysLogHealthCheckFailures() bool {
	if m != nil {
		return m.AlwaysLogHealthCheckFailures
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_HttpHealthCheck_)(nil),
		(*HealthCheck_TcpHealthCheck_)(nil),
		(*HealthCheck_GrpcHealthCheck_)(nil),
		(*HealthCheck_CustomHealthCheck_)(nil),
	}
}

type HealthCheck_Payload struct {
	// Types that are valid to be assigned to Payload:
	//	*HealthCheck_Payload_Text
	//	*HealthCheck_Payload_Binary
	Payload              isHealthCheck_Payload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *HealthCheck_Payload) Reset()         { *m = HealthCheck_Payload{} }
func (m *HealthCheck_Payload) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_Payload) ProtoMessage()    {}
func (*HealthCheck_Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 0}
}

func (m *HealthCheck_Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_Payload.Unmarshal(m, b)
}
func (m *HealthCheck_Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_Payload.Marshal(b, m, deterministic)
}
func (m *HealthCheck_Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_Payload.Merge(m, src)
}
func (m *HealthCheck_Payload) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_Payload.Size(m)
}
func (m *HealthCheck_Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_Payload proto.InternalMessageInfo

type isHealthCheck_Payload_Payload interface {
	isHealthCheck_Payload_Payload()
}

type HealthCheck_Payload_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type HealthCheck_Payload_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*HealthCheck_Payload_Text) isHealthCheck_Payload_Payload() {}

func (*HealthCheck_Payload_Binary) isHealthCheck_Payload_Payload() {}

func (m *HealthCheck_Payload) GetPayload() isHealthCheck_Payload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HealthCheck_Payload) GetText() string {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Text); ok {
		return x.Text
	}
	return ""
}

func (m *HealthCheck_Payload) GetBinary() []byte {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Binary); ok {
		return x.Binary
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_Payload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_Payload_Text)(nil),
		(*HealthCheck_Payload_Binary)(nil),
	}
}

type HealthCheck_HttpHealthCheck struct {
	Host                   string               `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Path                   string               `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Send                   *HealthCheck_Payload `protobuf:"bytes,3,opt,name=send,proto3" json:"send,omitempty"`
	Receive                *HealthCheck_Payload `protobuf:"bytes,4,opt,name=receive,proto3" json:"receive,omitempty"`
	ServiceName            string               `protobuf:"bytes,5,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	RequestHeadersToAdd    []*HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove []string             `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	UseHttp2               bool                 `protobuf:"varint,7,opt,name=use_http2,json=useHttp2,proto3" json:"use_http2,omitempty"`
	ExpectedStatuses       []*_type.Int64Range  `protobuf:"bytes,9,rep,name=expected_statuses,json=expectedStatuses,proto3" json:"expected_statuses,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}             `json:"-"`
	XXX_unrecognized       []byte               `json:"-"`
	XXX_sizecache          int32                `json:"-"`
}

func (m *HealthCheck_HttpHealthCheck) Reset()         { *m = HealthCheck_HttpHealthCheck{} }
func (m *HealthCheck_HttpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_HttpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_HttpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 1}
}

func (m *HealthCheck_HttpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Size(m)
}
func (m *HealthCheck_HttpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_HttpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_HttpHealthCheck) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetReceive() *HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToAdd() []*HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetUseHttp2() bool {
	if m != nil {
		return m.UseHttp2
	}
	return false
}

func (m *HealthCheck_HttpHealthCheck) GetExpectedStatuses() []*_type.Int64Range {
	if m != nil {
		return m.ExpectedStatuses
	}
	return nil
}

type HealthCheck_TcpHealthCheck struct {
	Send                 *HealthCheck_Payload   `protobuf:"bytes,1,opt,name=send,proto3" json:"send,omitempty"`
	Receive              []*HealthCheck_Payload `protobuf:"bytes,2,rep,name=receive,proto3" json:"receive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheck_TcpHealthCheck) Reset()         { *m = HealthCheck_TcpHealthCheck{} }
func (m *HealthCheck_TcpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_TcpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_TcpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 2}
}

func (m *HealthCheck_TcpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Size(m)
}
func (m *HealthCheck_TcpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_TcpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_TcpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_TcpHealthCheck) GetReceive() []*HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

type HealthCheck_RedisHealthCheck struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_RedisHealthCheck) Reset()         { *m = HealthCheck_RedisHealthCheck{} }
func (m *HealthCheck_RedisHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_RedisHealthCheck) ProtoMessage()    {}
func (*HealthCheck_RedisHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 3}
}

func (m *HealthCheck_RedisHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.Merge(m, src)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Size(m)
}
func (m *HealthCheck_RedisHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_RedisHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_RedisHealthCheck) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type HealthCheck_GrpcHealthCheck struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Authority            string   `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_GrpcHealthCheck) Reset()         { *m = HealthCheck_GrpcHealthCheck{} }
func (m *HealthCheck_GrpcHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_GrpcHealthCheck) ProtoMessage()    {}
func (*HealthCheck_GrpcHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 4}
}

func (m *HealthCheck_GrpcHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.Merge(m, src)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Size(m)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_GrpcHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_GrpcHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *HealthCheck_GrpcHealthCheck) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

type HealthCheck_CustomHealthCheck struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*HealthCheck_CustomHealthCheck_Config
	//	*HealthCheck_CustomHealthCheck_TypedConfig
	ConfigType           isHealthCheck_CustomHealthCheck_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *HealthCheck_CustomHealthCheck) Reset()         { *m = HealthCheck_CustomHealthCheck{} }
func (m *HealthCheck_CustomHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_CustomHealthCheck) ProtoMessage()    {}
func (*HealthCheck_CustomHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 5}
}

func (m *HealthCheck_CustomHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.Merge(m, src)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Size(m)
}
func (m *HealthCheck_CustomHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_CustomHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_CustomHealthCheck) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isHealthCheck_CustomHealthCheck_ConfigType interface {
	isHealthCheck_CustomHealthCheck_ConfigType()
}

type HealthCheck_CustomHealthCheck_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*HealthCheck_CustomHealthCheck_Config) isHealthCheck_CustomHealthCheck_ConfigType() {}

func (*HealthCheck_CustomHealthCheck_TypedConfig) isHealthCheck_CustomHealthCheck_ConfigType() {}

func (m *HealthCheck_CustomHealthCheck) GetConfigType() isHealthCheck_CustomHealthCheck_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *HealthCheck_CustomHealthCheck) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_Config); ok {
		return x.Config
	}
	return nil
}

func (m *HealthCheck_CustomHealthCheck) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_CustomHealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_CustomHealthCheck_Config)(nil),
		(*HealthCheck_CustomHealthCheck_TypedConfig)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.api.v3alpha.core.HealthStatus", HealthStatus_name, HealthStatus_value)
	proto.RegisterType((*HealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck")
	proto.RegisterType((*HealthCheck_Payload)(nil), "envoy.api.v3alpha.core.HealthCheck.Payload")
	proto.RegisterType((*HealthCheck_HttpHealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck.HttpHealthCheck")
	proto.RegisterType((*HealthCheck_TcpHealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck.TcpHealthCheck")
	proto.RegisterType((*HealthCheck_RedisHealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck.RedisHealthCheck")
	proto.RegisterType((*HealthCheck_GrpcHealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck.GrpcHealthCheck")
	proto.RegisterType((*HealthCheck_CustomHealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck.CustomHealthCheck")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/core/health_check.proto", fileDescriptor_092b49335160c834)
}

var fileDescriptor_092b49335160c834 = []byte{
	// 1183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcf, 0x6e, 0xdb, 0xc6,
	0x13, 0xc7, 0x45, 0x4b, 0xb1, 0xa4, 0x91, 0x2c, 0x53, 0xeb, 0xd8, 0x66, 0x14, 0xff, 0x7e, 0x50,
	0x8a, 0x1c, 0x9c, 0x14, 0x90, 0x50, 0xbb, 0x4d, 0x91, 0x5e, 0x52, 0xd3, 0x56, 0x22, 0xa5, 0x89,
	0x62, 0x6c, 0xe4, 0x06, 0x45, 0x81, 0xb2, 0x6b, 0x72, 0x2c, 0xb1, 0xa1, 0xb9, 0xec, 0x72, 0xa9,
	0x44, 0x4f, 0xd0, 0x7b, 0x4f, 0x7d, 0x86, 0x00, 0x7d, 0xaf, 0xbe, 0x40, 0x81, 0xc2, 0xa7, 0x82,
	0x4b, 0xea, 0xbf, 0xd3, 0x28, 0x45, 0x6f, 0xe4, 0xec, 0x7c, 0x3f, 0xb3, 0xc3, 0xf9, 0x03, 0xc2,
	0x3d, 0xf4, 0x87, 0x7c, 0xd4, 0x64, 0x81, 0xdb, 0x1c, 0x1e, 0x32, 0x2f, 0x18, 0xb0, 0xa6, 0xcd,
	0x05, 0x36, 0x07, 0xc8, 0x3c, 0x39, 0xb0, 0xec, 0x01, 0xda, 0xaf, 0x1b, 0x81, 0xe0, 0x92, 0x93,
	0x1d, 0xe5, 0xda, 0x60, 0x81, 0xdb, 0x48, 0x5d, 0x1b, 0xb1, 0x6b, 0xed, 0xce, 0x7b, 0x10, 0xe7,
	0x2c, 0xc4, 0x44, 0x5a, 0x4b, 0xa4, 0x4d, 0x39, 0x0a, 0xb0, 0x29, 0x98, 0xdf, 0x1f, 0xdb, 0x6f,
	0xf5, 0x39, 0xef, 0x7b, 0xd8, 0x54, 0x6f, 0xe7, 0xd1, 0x45, 0x93, 0xf9, 0xa3, 0xf4, 0xe8, 0xff,
	0x8b, 0x47, 0x4e, 0x24, 0x98, 0x74, 0xb9, 0x9f, 0x9e, 0xef, 0x2d, 0x9e, 0x87, 0x52, 0x44, 0xb6,
	0x7c, 0x9f, 0xfa, 0x8d, 0x60, 0x41, 0x80, 0x22, 0x4c, 0xcf, 0x77, 0x87, 0xcc, 0x73, 0x1d, 0x26,
	0xb1, 0x39, 0x7e, 0x48, 0x0e, 0x3e, 0xf9, 0xe5, 0x26, 0x94, 0xda, 0x2a, 0xf7, 0xe3, 0x38, 0x75,
	0xf2, 0x08, 0xf2, 0xd2, 0xbd, 0x44, 0x1e, 0x49, 0x43, 0xab, 0x6b, 0xfb, 0xa5, 0x83, 0x5b, 0x8d,
	0x04, 0xdd, 0x18, 0xa3, 0x1b, 0x27, 0xe9, 0xc5, 0x4c, 0xb8, 0x32, 0xf3, 0xef, 0xb4, 0x5c, 0x41,
	0xbb, 0x9f, 0xa1, 0x63, 0x15, 0x39, 0x82, 0x82, 0xeb, 0x4b, 0x14, 0x43, 0xe6, 0x19, 0x6b, 0x1f,
	0x43, 0x98, 0xc8, 0xc8, 0xd7, 0x50, 0x71, 0x7d, 0x57, 0xba, 0xcc, 0xb3, 0x7e, 0x72, 0xa5, 0x44,
	0x61, 0xdc, 0xfc, 0x00, 0x88, 0x6e, 0xa4, 0x82, 0xa7, 0xca, 0x9f, 0x98, 0xb0, 0x39, 0xa6, 0x8d,
	0x11, 0xd9, 0x0f, 0x21, 0x2a, 0x63, 0x45, 0xca, 0x78, 0x00, 0xbb, 0x0b, 0x0c, 0x2b, 0x40, 0x61,
	0xa3, 0x2f, 0x0d, 0x52, 0xd7, 0xf6, 0x37, 0xe8, 0xf6, 0xbc, 0xe0, 0x34, 0x39, 0x24, 0xcf, 0x61,
	0x2b, 0xf2, 0x93, 0x76, 0x1a, 0x59, 0x72, 0x20, 0x30, 0x1c, 0x70, 0xcf, 0x31, 0x72, 0x2a, 0xfe,
	0xde, 0x52, 0xfc, 0xb3, 0x8e, 0x2f, 0x0f, 0x0f, 0xbe, 0x65, 0x5e, 0x84, 0x94, 0x4c, 0x84, 0xbd,
	0xb1, 0x8e, 0x74, 0xa0, 0xba, 0x0c, 0xbb, 0xb1, 0x02, 0x4c, 0x5f, 0x42, 0x7d, 0x09, 0x05, 0xe6,
	0x49, 0x2b, 0xe0, 0x42, 0x1a, 0xeb, 0x2b, 0x10, 0xf2, 0xcc, 0x93, 0xa7, 0x5c, 0x48, 0xd2, 0x02,
	0x5d, 0x60, 0x14, 0xa2, 0x65, 0x73, 0xdf, 0x47, 0x3b, 0xfe, 0x5c, 0x46, 0x5e, 0x01, 0x6a, 0x4b,
	0x00, 0x93, 0x73, 0x2f, 0x91, 0x6f, 0x2a, 0xcd, 0xf1, 0x44, 0x42, 0x18, 0x54, 0x07, 0x52, 0x06,
	0xd6, 0xec, 0xac, 0x19, 0x05, 0xc5, 0x39, 0x6c, 0x5c, 0x3f, 0x6c, 0x8d, 0x99, 0xde, 0x6c, 0xb4,
	0xa5, 0x0c, 0x66, 0xde, 0xdb, 0x19, 0xba, 0x39, 0x98, 0x37, 0x91, 0x1f, 0x40, 0x97, 0xf6, 0x42,
	0x84, 0xa2, 0x8a, 0x70, 0xb0, 0x4a, 0x84, 0x9e, 0xbd, 0x10, 0xa0, 0x22, 0xe7, 0x2c, 0x71, 0x0a,
	0x7d, 0x11, 0xd8, 0xf3, 0x01, 0x4a, 0xab, 0xa7, 0xf0, 0x44, 0x04, 0xf6, 0x42, 0x0a, 0xfd, 0x79,
	0x13, 0xe9, 0xc3, 0x96, 0x1d, 0x85, 0x92, 0x5f, 0xce, 0x07, 0xd9, 0x50, 0x41, 0xbe, 0x58, 0x25,
	0xc8, 0xb1, 0x92, 0xcf, 0x87, 0xa9, 0xda, 0x8b, 0x46, 0xf2, 0x12, 0xb6, 0x7c, 0x6e, 0x49, 0xc1,
	0x2e, 0x2e, 0x5c, 0xdb, 0x9a, 0x0c, 0x6d, 0xf9, 0x43, 0x43, 0x5b, 0xb8, 0x32, 0x6f, 0xbc, 0xd3,
	0xd6, 0xee, 0x67, 0x68, 0xd5, 0xe7, 0xbd, 0x44, 0xde, 0x19, 0xcf, 0x2e, 0x85, 0x69, 0x13, 0x4f,
	0x99, 0x95, 0x8f, 0x60, 0x4e, 0xe4, 0x13, 0xe6, 0xf7, 0xb0, 0x3b, 0x65, 0xa2, 0xd3, 0xc7, 0x29,
	0x78, 0x73, 0x75, 0xf0, 0xf6, 0x84, 0xd1, 0x72, 0xfa, 0x38, 0x81, 0xbf, 0x82, 0xed, 0xeb, 0xd1,
	0xfa, 0xea, 0xe8, 0xad, 0xeb, 0xc0, 0x77, 0xa1, 0x82, 0x43, 0xf4, 0xa5, 0xe5, 0xf1, 0xbe, 0x15,
	0x30, 0x39, 0x30, 0xaa, 0x75, 0x6d, 0xbf, 0x48, 0xcb, 0xca, 0xfa, 0x8c, 0xf7, 0x4f, 0x99, 0x1c,
	0x90, 0xc7, 0x50, 0x67, 0xde, 0x1b, 0x36, 0x0a, 0x95, 0xdb, 0x6c, 0xc5, 0xad, 0x0b, 0xe6, 0x7a,
	0x91, 0xc0, 0xd0, 0xd8, 0xaa, 0x6b, 0xfb, 0x05, 0xba, 0x97, 0xf8, 0x3d, 0xe3, 0xfd, 0x99, 0x22,
	0x3e, 0x4e, 0x7d, 0x6a, 0x14, 0xf2, 0xa7, 0x6c, 0xe4, 0x71, 0xe6, 0x90, 0xff, 0x41, 0x4e, 0xe2,
	0xdb, 0x64, 0x7f, 0x17, 0xcd, 0xfc, 0x95, 0x99, 0x13, 0x6b, 0x75, 0xad, 0x9d, 0xa1, 0xca, 0x4c,
	0x0c, 0x58, 0x3f, 0x77, 0x7d, 0x26, 0x46, 0x6a, 0x3d, 0x97, 0xdb, 0x19, 0x9a, 0xbe, 0x9b, 0x15,
	0xc8, 0x07, 0x29, 0x23, 0xfb, 0x97, 0xa9, 0xd5, 0xfe, 0xcc, 0xc2, 0xe6, 0xc2, 0xcc, 0x11, 0x02,
	0xb9, 0x01, 0x0f, 0x53, 0x38, 0x55, 0xcf, 0xe4, 0x36, 0xe4, 0x54, 0x7e, 0x6b, 0x73, 0x01, 0xa9,
	0x32, 0x92, 0x47, 0x90, 0x0b, 0xd1, 0x77, 0xd2, 0xfd, 0xfb, 0xe9, 0x2a, 0xfd, 0x9b, 0x26, 0x42,
	0x95, 0x90, 0xb4, 0x20, 0x2f, 0xd0, 0x46, 0x77, 0x88, 0xe9, 0x0e, 0xfd, 0x28, 0xc6, 0x58, 0x4b,
	0xee, 0x40, 0x39, 0x44, 0x31, 0x74, 0x6d, 0xb4, 0x7c, 0x76, 0x89, 0x6a, 0x85, 0x16, 0x69, 0x29,
	0xb5, 0x75, 0xd9, 0x25, 0x12, 0x17, 0x76, 0x04, 0xfe, 0x1c, 0x61, 0x28, 0xe3, 0x42, 0x38, 0x28,
	0x42, 0x4b, 0x72, 0x8b, 0x39, 0x8e, 0xb1, 0x5e, 0xcf, 0xee, 0x97, 0x0e, 0xee, 0xfd, 0x43, 0x60,
	0x07, 0x85, 0xda, 0x7a, 0x2f, 0x02, 0xd5, 0x1b, 0xc5, 0x2b, 0x73, 0xfd, 0x57, 0x2d, 0xab, 0xff,
	0x91, 0xa7, 0x5b, 0x29, 0x33, 0x71, 0x0a, 0x7b, 0xfc, 0xc8, 0x71, 0xc8, 0x43, 0xb8, 0x75, 0x4d,
	0x28, 0x81, 0x97, 0x7c, 0x88, 0x46, 0xa1, 0x9e, 0xdd, 0x2f, 0xd2, 0x9d, 0x45, 0x1d, 0x55, 0xa7,
	0xe4, 0x36, 0x14, 0xe3, 0x55, 0x1c, 0x6f, 0xbe, 0x03, 0xb5, 0x85, 0x0b, 0xb4, 0x10, 0x85, 0x18,
	0x17, 0xea, 0x80, 0x1c, 0x43, 0x15, 0xdf, 0x06, 0x68, 0x4b, 0x74, 0xac, 0x50, 0x32, 0x19, 0x85,
	0x18, 0x1a, 0x45, 0x75, 0xfb, 0x9d, 0xf4, 0xf6, 0xf1, 0x4f, 0x49, 0xa3, 0xe3, 0xcb, 0x07, 0x9f,
	0xd3, 0xf8, 0xcf, 0x84, 0xea, 0x63, 0xc1, 0xcb, 0xd4, 0xbf, 0xf6, 0x9b, 0x06, 0x95, 0xf9, 0x4d,
	0x38, 0xa9, 0xa2, 0xf6, 0x1f, 0x54, 0x71, 0x4d, 0x5d, 0xe7, 0x5f, 0x55, 0xb1, 0x76, 0x17, 0x74,
	0x8a, 0x8e, 0x1b, 0xce, 0xde, 0x4d, 0x87, 0xec, 0x6b, 0x1c, 0xa5, 0x1d, 0x19, 0x3f, 0xd6, 0x28,
	0x6c, 0x2e, 0x2c, 0xda, 0xa5, 0xf2, 0x6b, 0xcb, 0xe5, 0xdf, 0x83, 0x22, 0x8b, 0xe4, 0x80, 0x0b,
	0x57, 0x26, 0xb3, 0x51, 0xa4, 0x53, 0x43, 0xed, 0x77, 0x0d, 0xaa, 0x4b, 0x8b, 0x35, 0x6e, 0xfd,
	0x29, 0x6e, 0xa6, 0xf5, 0x63, 0x23, 0xf9, 0x0c, 0xd6, 0x6d, 0xee, 0x5f, 0xb8, 0xfd, 0xf4, 0x47,
	0x68, 0x77, 0x69, 0x97, 0xbc, 0x54, 0xff, 0x70, 0xf1, 0x08, 0x26, 0x8e, 0xe4, 0x21, 0x94, 0xe3,
	0xfa, 0x38, 0x56, 0x2a, 0x4c, 0xa6, 0xe6, 0xe6, 0x92, 0xf0, 0xc8, 0x1f, 0xb5, 0x33, 0xb4, 0xa4,
	0x7c, 0x8f, 0x95, 0xab, 0xb9, 0x01, 0xa5, 0x44, 0x64, 0xc5, 0x56, 0x73, 0x1b, 0x2a, 0xb3, 0xdb,
	0x04, 0x85, 0x9a, 0xe9, 0xa7, 0xb9, 0x02, 0xe8, 0x25, 0x4a, 0x44, 0xfc, 0x11, 0xe7, 0xd6, 0xcd,
	0xfd, 0x1f, 0xa1, 0x9c, 0x64, 0x96, 0xf4, 0x01, 0x29, 0x41, 0xfe, 0xac, 0xfb, 0x4d, 0xf7, 0xc5,
	0xab, 0xae, 0x9e, 0x89, 0x5f, 0xda, 0xad, 0xa3, 0x67, 0xbd, 0xf6, 0x77, 0xba, 0x46, 0x36, 0xa0,
	0x78, 0xd6, 0x1d, 0xbf, 0xae, 0x91, 0x32, 0x14, 0x4e, 0xe8, 0x51, 0xa7, 0xdb, 0xe9, 0x3e, 0xd1,
	0xb3, 0xb1, 0x67, 0xaf, 0xf3, 0xbc, 0xf5, 0xe2, 0xac, 0xa7, 0xe7, 0xd4, 0x51, 0xeb, 0x09, 0x3d,
	0x3a, 0x69, 0x9d, 0xe8, 0x37, 0xcc, 0xaf, 0xe0, 0xae, 0xcb, 0x93, 0xb2, 0x07, 0x82, 0xbf, 0x1d,
	0xbd, 0xa7, 0x03, 0x4c, 0x7d, 0xe6, 0x0b, 0x9f, 0xc6, 0x19, 0x9f, 0x6a, 0xe7, 0xeb, 0x2a, 0xf5,
	0xc3, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x9a, 0xea, 0xf4, 0xc0, 0x0b, 0x00, 0x00,
}
