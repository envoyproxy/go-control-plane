// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/v3alpha/circuit_breaker.proto

package envoy_config_cluster_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/core/v3alpha"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
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

type CircuitBreakers struct {
	Thresholds           []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CircuitBreakers) Reset()         { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()    {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cdfac73f383a10, []int{0}
}

func (m *CircuitBreakers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers.Unmarshal(m, b)
}
func (m *CircuitBreakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers.Merge(m, src)
}
func (m *CircuitBreakers) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers.Size(m)
}
func (m *CircuitBreakers) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers proto.InternalMessageInfo

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

type CircuitBreakers_Thresholds struct {
	Priority             v3alpha.RoutingPriority                 `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.config.core.v3alpha.RoutingPriority" json:"priority,omitempty"`
	MaxConnections       *wrappers.UInt32Value                   `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	MaxPendingRequests   *wrappers.UInt32Value                   `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	MaxRequests          *wrappers.UInt32Value                   `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	MaxRetries           *wrappers.UInt32Value                   `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	RetryBudget          *CircuitBreakers_Thresholds_RetryBudget `protobuf:"bytes,8,opt,name=retry_budget,json=retryBudget,proto3" json:"retry_budget,omitempty"`
	TrackRemaining       bool                                    `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	MaxConnectionPools   *wrappers.UInt32Value                   `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CircuitBreakers_Thresholds) Reset()         { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cdfac73f383a10, []int{0, 0}
}

func (m *CircuitBreakers_Thresholds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Unmarshal(m, b)
}
func (m *CircuitBreakers_Thresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers_Thresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Size(m)
}
func (m *CircuitBreakers_Thresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds) GetPriority() v3alpha.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return v3alpha.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetRetryBudget() *CircuitBreakers_Thresholds_RetryBudget {
	if m != nil {
		return m.RetryBudget
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetTrackRemaining() bool {
	if m != nil {
		return m.TrackRemaining
	}
	return false
}

func (m *CircuitBreakers_Thresholds) GetMaxConnectionPools() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectionPools
	}
	return nil
}

type CircuitBreakers_Thresholds_RetryBudget struct {
	BudgetPercent        *v3alpha1.Percent     `protobuf:"bytes,1,opt,name=budget_percent,json=budgetPercent,proto3" json:"budget_percent,omitempty"`
	MinRetryConcurrency  *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=min_retry_concurrency,json=minRetryConcurrency,proto3" json:"min_retry_concurrency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CircuitBreakers_Thresholds_RetryBudget) Reset() {
	*m = CircuitBreakers_Thresholds_RetryBudget{}
}
func (m *CircuitBreakers_Thresholds_RetryBudget) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds_RetryBudget) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds_RetryBudget) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cdfac73f383a10, []int{0, 0, 0}
}

func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Unmarshal(m, b)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Size(m)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds_RetryBudget) GetBudgetPercent() *v3alpha1.Percent {
	if m != nil {
		return m.BudgetPercent
	}
	return nil
}

func (m *CircuitBreakers_Thresholds_RetryBudget) GetMinRetryConcurrency() *wrappers.UInt32Value {
	if m != nil {
		return m.MinRetryConcurrency
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.config.cluster.v3alpha.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.config.cluster.v3alpha.CircuitBreakers.Thresholds")
	proto.RegisterType((*CircuitBreakers_Thresholds_RetryBudget)(nil), "envoy.config.cluster.v3alpha.CircuitBreakers.Thresholds.RetryBudget")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/v3alpha/circuit_breaker.proto", fileDescriptor_c7cdfac73f383a10)
}

var fileDescriptor_c7cdfac73f383a10 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6b, 0xd4, 0x4c,
	0x18, 0xc7, 0x49, 0xfb, 0xb6, 0xef, 0x32, 0xa9, 0x5b, 0x49, 0x15, 0xc3, 0x5a, 0xca, 0x2a, 0x45,
	0x97, 0x0a, 0x13, 0x4c, 0x41, 0x64, 0xa5, 0x08, 0x59, 0x3d, 0x78, 0x91, 0x10, 0x54, 0xbc, 0x85,
	0xd9, 0xec, 0xd3, 0xec, 0xd0, 0x64, 0x66, 0x9c, 0x99, 0xac, 0x9b, 0xab, 0x27, 0x3f, 0x80, 0x27,
	0xbf, 0x8f, 0x9f, 0xc5, 0xcf, 0xa0, 0x27, 0x49, 0x26, 0x9b, 0xb4, 0x3d, 0x94, 0xe8, 0x2d, 0x93,
	0x79, 0x7e, 0xbf, 0xcc, 0xf3, 0x9f, 0x27, 0xc8, 0x07, 0xb6, 0xe2, 0xa5, 0x97, 0x70, 0x76, 0x4e,
	0x53, 0x2f, 0xc9, 0x0a, 0xa5, 0x41, 0x7a, 0xab, 0x53, 0x92, 0x89, 0x25, 0xf1, 0x12, 0x2a, 0x93,
	0x82, 0xea, 0x78, 0x2e, 0x81, 0x5c, 0x80, 0xc4, 0x42, 0x72, 0xcd, 0x9d, 0xc3, 0x9a, 0xc1, 0x86,
	0xc1, 0x0d, 0x83, 0x1b, 0x66, 0x74, 0x7c, 0xd5, 0xc8, 0x25, 0xb4, 0xba, 0x39, 0x51, 0x60, 0x1c,
	0xa3, 0xb1, 0xa9, 0xd2, 0xa5, 0xe8, 0xb6, 0x05, 0xc8, 0x04, 0x98, 0x6e, 0x2a, 0x8e, 0x52, 0xce,
	0xd3, 0x0c, 0xbc, 0x7a, 0x35, 0x2f, 0xce, 0xbd, 0xcf, 0x92, 0x08, 0x01, 0x52, 0x35, 0xfb, 0x0f,
	0x8a, 0x85, 0x20, 0x1e, 0x61, 0x8c, 0x6b, 0xa2, 0x29, 0x67, 0xca, 0x5b, 0x81, 0x54, 0x94, 0x33,
	0xca, 0xd2, 0xa6, 0xe4, 0xde, 0x8a, 0x64, 0x74, 0x41, 0x34, 0x78, 0x9b, 0x07, 0xb3, 0xf1, 0xf0,
	0xdb, 0x00, 0xed, 0xcf, 0x4c, 0x6f, 0x81, 0x69, 0x4d, 0x39, 0x1f, 0x11, 0xd2, 0x4b, 0x09, 0x6a,
	0xc9, 0xb3, 0x85, 0x72, 0xad, 0xf1, 0xf6, 0xc4, 0xf6, 0x9f, 0xe3, 0x9b, 0x5a, 0xc5, 0xd7, 0x14,
	0xf8, 0x5d, 0xcb, 0x47, 0x97, 0x5c, 0xa3, 0x5f, 0xbb, 0x08, 0x75, 0x5b, 0x4e, 0x88, 0x06, 0x42,
	0x52, 0x2e, 0xa9, 0x2e, 0x5d, 0x6b, 0x6c, 0x4d, 0x86, 0xfe, 0xc9, 0xb5, 0xcf, 0x70, 0x09, 0xed,
	0x37, 0x22, 0x5e, 0x68, 0xca, 0xd2, 0xb0, 0x21, 0x82, 0xc1, 0xef, 0x60, 0xe7, 0x8b, 0xb5, 0x75,
	0xdb, 0x8a, 0x5a, 0x8b, 0xf3, 0x1a, 0xed, 0xe7, 0x64, 0x1d, 0x27, 0x9c, 0x31, 0x48, 0xea, 0x2c,
	0xdc, 0xad, 0xb1, 0x35, 0xb1, 0xfd, 0x43, 0x6c, 0x42, 0xc4, 0x9b, 0x10, 0xf1, 0xfb, 0x37, 0x4c,
	0x9f, 0xfa, 0x1f, 0x48, 0x56, 0x40, 0x34, 0xcc, 0xc9, 0x7a, 0xd6, 0x31, 0xce, 0x5b, 0x74, 0xa7,
	0xd2, 0x08, 0x60, 0x0b, 0xca, 0xd2, 0x58, 0xc2, 0xa7, 0x02, 0x94, 0x56, 0xee, 0x76, 0x0f, 0x97,
	0x93, 0x93, 0x75, 0x68, 0xc0, 0xa8, 0xe1, 0x9c, 0x97, 0x68, 0xaf, 0xf2, 0xb5, 0x9e, 0xff, 0x7a,
	0x78, 0xec, 0x9c, 0xac, 0x5b, 0xc1, 0x19, 0xb2, 0x8d, 0x40, 0x4b, 0x0a, 0xca, 0xdd, 0xe9, 0xc1,
	0xa3, 0x9a, 0xaf, 0xeb, 0x9d, 0x14, 0xed, 0x55, 0x68, 0x19, 0xcf, 0x8b, 0x45, 0x0a, 0xda, 0x1d,
	0xd4, 0xfc, 0xab, 0x7f, 0xbd, 0x53, 0x5c, 0x79, 0xcb, 0xa0, 0x76, 0x45, 0xb6, 0xec, 0x16, 0xce,
	0x63, 0xb4, 0xaf, 0x25, 0x49, 0x2e, 0x62, 0x09, 0x39, 0xa1, 0xd5, 0x00, 0xba, 0xbb, 0x63, 0x6b,
	0x32, 0x88, 0x86, 0xf5, 0xeb, 0x68, 0xf3, 0x76, 0x93, 0x70, 0x77, 0x51, 0xb1, 0xe0, 0x3c, 0x53,
	0xee, 0xff, 0x3d, 0x13, 0xee, 0x6e, 0x2b, 0xac, 0xb8, 0xd1, 0x4f, 0x0b, 0xd9, 0x97, 0x4e, 0xe5,
	0x04, 0x68, 0x68, 0x7a, 0x8d, 0x9b, 0x7f, 0xa9, 0x1e, 0x30, 0xdb, 0xbf, 0xdf, 0xf4, 0x5c, 0xfd,
	0x6e, 0x6d, 0xa7, 0xa1, 0x29, 0x89, 0x6e, 0x19, 0xa4, 0x59, 0x3a, 0x21, 0xba, 0x9b, 0x53, 0x16,
	0x9b, 0xe4, 0x12, 0xce, 0x92, 0x42, 0x4a, 0x60, 0x49, 0xd9, 0x6b, 0xa4, 0x0e, 0x72, 0xca, 0xea,
	0x03, 0xcd, 0x3a, 0x70, 0x1a, 0x7c, 0xff, 0xf1, 0xf5, 0xe8, 0x0c, 0xbd, 0x30, 0x67, 0x20, 0x82,
	0xe2, 0x95, 0xdf, 0xe6, 0xde, 0x2f, 0xef, 0xe9, 0xb3, 0xca, 0xf1, 0x14, 0x79, 0x7f, 0xe9, 0x98,
	0x3e, 0xa9, 0xb8, 0x47, 0xe8, 0xb8, 0x0f, 0x17, 0xcc, 0xd0, 0x09, 0xe5, 0x26, 0x2a, 0x21, 0xf9,
	0xba, 0xbc, 0x71, 0x52, 0x82, 0x83, 0xab, 0x78, 0x58, 0xe5, 0x11, 0x5a, 0xf3, 0xdd, 0x3a, 0x98,
	0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x53, 0xf0, 0x81, 0x5a, 0x05, 0x00, 0x00,
}
