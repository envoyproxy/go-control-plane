// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/cluster/circuit_breaker.proto

package envoy_api_v3alpha_cluster

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/v2/envoy/api/v3alpha/core"
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
	return fileDescriptor_9f46fa043d540e79, []int{0}
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
	Priority             core.RoutingPriority  `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.api.v3alpha.core.RoutingPriority" json:"priority,omitempty"`
	MaxConnections       *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	MaxPendingRequests   *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	MaxRequests          *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	MaxRetries           *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	TrackRemaining       bool                  `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	MaxConnectionPools   *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CircuitBreakers_Thresholds) Reset()         { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f46fa043d540e79, []int{0, 0}
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

func (m *CircuitBreakers_Thresholds) GetPriority() core.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return core.RoutingPriority_DEFAULT
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

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.api.v3alpha.cluster.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.api.v3alpha.cluster.CircuitBreakers.Thresholds")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/cluster/circuit_breaker.proto", fileDescriptor_9f46fa043d540e79)
}

var fileDescriptor_9f46fa043d540e79 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0x87, 0x95, 0xed, 0xb2, 0xac, 0x5c, 0xd4, 0x4a, 0x01, 0xa1, 0x50, 0x21, 0x54, 0xb8, 0xb4,
	0x27, 0x5b, 0x4a, 0xc5, 0x11, 0x21, 0x5a, 0xed, 0x81, 0xcb, 0x2a, 0x0a, 0xec, 0x5e, 0x23, 0x37,
	0x3b, 0xa4, 0xd6, 0x26, 0x1e, 0x33, 0x76, 0x96, 0xf4, 0x95, 0x78, 0x16, 0x5e, 0x82, 0x2b, 0x4f,
	0x81, 0x1a, 0xf7, 0x0f, 0xac, 0x5a, 0x29, 0xc7, 0x78, 0xe6, 0xfb, 0xc6, 0x3f, 0x67, 0x98, 0x00,
	0xfd, 0x80, 0x6b, 0x21, 0x8d, 0x12, 0x0f, 0x33, 0x59, 0x9a, 0x95, 0x14, 0x79, 0x59, 0x5b, 0x07,
	0x24, 0x72, 0x45, 0x79, 0xad, 0x5c, 0xb6, 0x24, 0x90, 0xf7, 0x40, 0xdc, 0x10, 0x3a, 0x0c, 0x5f,
	0xb5, 0x00, 0x97, 0x46, 0xf1, 0x2d, 0xc0, 0xb7, 0xc0, 0xe8, 0xed, 0x11, 0x17, 0x12, 0x88, 0xa5,
	0xb4, 0xe0, 0xe9, 0xd1, 0x9b, 0x02, 0xb1, 0x28, 0x41, 0xb4, 0x5f, 0xcb, 0xfa, 0x9b, 0xf8, 0x41,
	0xd2, 0x18, 0x20, 0xeb, 0xeb, 0xef, 0x7e, 0x9d, 0xb3, 0xe1, 0xc2, 0xcf, 0x9d, 0xfb, 0xb1, 0x36,
	0xbc, 0x61, 0xcc, 0xad, 0x08, 0xec, 0x0a, 0xcb, 0x3b, 0x1b, 0x05, 0xe3, 0xde, 0xb4, 0x1f, 0xbf,
	0xe7, 0x27, 0xaf, 0xc1, 0x1f, 0xf1, 0xfc, 0xeb, 0x1e, 0x4e, 0xff, 0x11, 0x8d, 0x7e, 0xf7, 0x18,
	0x3b, 0x94, 0xc2, 0x05, 0xbb, 0x34, 0xa4, 0x90, 0x94, 0x5b, 0x47, 0xc1, 0x38, 0x98, 0x0e, 0xe2,
	0xc9, 0xb1, 0x19, 0x48, 0xc0, 0x53, 0xac, 0x9d, 0xd2, 0x45, 0xb2, 0x6d, 0x4f, 0xf7, 0x60, 0x78,
	0xc5, 0x86, 0x95, 0x6c, 0xb2, 0x1c, 0xb5, 0x86, 0xdc, 0x29, 0xd4, 0x36, 0x3a, 0x1b, 0x07, 0xd3,
	0x7e, 0xfc, 0x9a, 0xfb, 0xe0, 0x7c, 0x17, 0x9c, 0xdf, 0x7c, 0xd6, 0x6e, 0x16, 0xdf, 0xca, 0xb2,
	0x86, 0x74, 0x50, 0xc9, 0x66, 0x71, 0x60, 0xc2, 0x6b, 0xf6, 0x62, 0xa3, 0x31, 0xa0, 0xef, 0x94,
	0x2e, 0x32, 0x82, 0xef, 0x35, 0x58, 0x67, 0xa3, 0x5e, 0x07, 0x57, 0x58, 0xc9, 0x26, 0xf1, 0x60,
	0xba, 0xe5, 0xc2, 0x8f, 0xec, 0xd9, 0xc6, 0xb7, 0xf7, 0x9c, 0x77, 0xf0, 0xf4, 0x2b, 0xd9, 0xec,
	0x05, 0x1f, 0x58, 0xdf, 0x0b, 0x1c, 0x29, 0xb0, 0xd1, 0x93, 0x0e, 0x3c, 0x6b, 0xf9, 0xb6, 0x3f,
	0x9c, 0xb0, 0xa1, 0x23, 0x99, 0xdf, 0x67, 0x04, 0x95, 0x54, 0x5a, 0xe9, 0x22, 0xba, 0x18, 0x07,
	0xd3, 0xcb, 0x74, 0xd0, 0x1e, 0xa7, 0xbb, 0xd3, 0x5d, 0xf0, 0xc3, 0xfb, 0x65, 0x06, 0xb1, 0xb4,
	0xd1, 0xd3, 0x8e, 0xc1, 0x0f, 0x8f, 0x98, 0x6c, 0xb8, 0x39, 0xb1, 0x89, 0x42, 0xff, 0x1b, 0x0d,
	0x61, 0xb3, 0x3e, 0xbd, 0x35, 0xf3, 0xe7, 0xff, 0xaf, 0x4d, 0xb2, 0x19, 0x91, 0x04, 0x3f, 0xcf,
	0x5e, 0x5e, 0xb5, 0xc8, 0x27, 0xa3, 0xf8, 0x6d, 0xcc, 0x17, 0xbe, 0xfb, 0xfa, 0xcb, 0x9f, 0x53,
	0x85, 0xe5, 0x45, 0x7b, 0xbb, 0xd9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x35, 0x7c, 0xd8,
	0x5a, 0x03, 0x00, 0x00,
}
