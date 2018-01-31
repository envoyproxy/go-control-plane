// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/ratelimit/v2/rls.proto

/*
	Package ratelimit is a generated protocol buffer package.

	It is generated from these files:
		envoy/service/ratelimit/v2/rls.proto

	It has these top-level messages:
		RateLimitRequest
		RateLimitResponse
*/
package ratelimit

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_ratelimit "github.com/envoyproxy/go-control-plane/envoy/api/v2/ratelimit"
import _ "github.com/lyft/protoc-gen-validate/validate"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RateLimitResponse_Code int32

const (
	RateLimitResponse_UNKNOWN    RateLimitResponse_Code = 0
	RateLimitResponse_OK         RateLimitResponse_Code = 1
	RateLimitResponse_OVER_LIMIT RateLimitResponse_Code = 2
)

var RateLimitResponse_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "OVER_LIMIT",
}
var RateLimitResponse_Code_value = map[string]int32{
	"UNKNOWN":    0,
	"OK":         1,
	"OVER_LIMIT": 2,
}

func (x RateLimitResponse_Code) String() string {
	return proto.EnumName(RateLimitResponse_Code_name, int32(x))
}
func (RateLimitResponse_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptorRls, []int{1, 0} }

type RateLimitResponse_RateLimit_Unit int32

const (
	RateLimitResponse_RateLimit_UNKNOWN RateLimitResponse_RateLimit_Unit = 0
	RateLimitResponse_RateLimit_SECOND  RateLimitResponse_RateLimit_Unit = 1
	RateLimitResponse_RateLimit_MINUTE  RateLimitResponse_RateLimit_Unit = 2
	RateLimitResponse_RateLimit_HOUR    RateLimitResponse_RateLimit_Unit = 3
	RateLimitResponse_RateLimit_DAY     RateLimitResponse_RateLimit_Unit = 4
)

var RateLimitResponse_RateLimit_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
	4: "DAY",
}
var RateLimitResponse_RateLimit_Unit_value = map[string]int32{
	"UNKNOWN": 0,
	"SECOND":  1,
	"MINUTE":  2,
	"HOUR":    3,
	"DAY":     4,
}

func (x RateLimitResponse_RateLimit_Unit) String() string {
	return proto.EnumName(RateLimitResponse_RateLimit_Unit_name, int32(x))
}
func (RateLimitResponse_RateLimit_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorRls, []int{1, 0, 0}
}

// Main message for a rate limit request. The rate limit service is designed to be fully generic
// in the sense that it can operate on arbitrary hierarchical key/value pairs. The loaded
// configuration will parse the request and find the most specific limit to apply. In addition,
// a RateLimitRequest can contain multiple "descriptors" to limit on. When multiple descriptors
// are provided, the server will limit on *ALL* of them and return an OVER_LIMIT response if any
// of them are over limit. This enables more complex application level rate limiting scenarios
// if desired.
// [#not-implemented-hide:] Hiding API for now.
type RateLimitRequest struct {
	// All rate limit requests must specify a domain. This enables the configuration to be per
	// application without fear of overlap. E.g., "envoy".
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// All rate limit requests must specify at least one RateLimitDescriptor. Each descriptor is
	// processed by the service (see below). If any of the descriptors are over limit, the entire
	// request is considered to be over limit.
	Descriptors []*envoy_api_v2_ratelimit.RateLimitDescriptor `protobuf:"bytes,2,rep,name=descriptors" json:"descriptors,omitempty"`
	// Rate limit requests can optionally specify the number of hits a request adds to the matched
	// limit. If the value is not set in the message, a request increases the matched limit by 1.
	HitsAddend uint32 `protobuf:"varint,3,opt,name=hits_addend,json=hitsAddend,proto3" json:"hits_addend,omitempty"`
}

func (m *RateLimitRequest) Reset()                    { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string            { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()               {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) { return fileDescriptorRls, []int{0} }

func (m *RateLimitRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitRequest) GetDescriptors() []*envoy_api_v2_ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimitRequest) GetHitsAddend() uint32 {
	if m != nil {
		return m.HitsAddend
	}
	return 0
}

// A response from a ShouldRateLimit call.
// [#not-implemented-hide:] Hiding API for now.
type RateLimitResponse struct {
	// The overall response code which takes into account all of the descriptors that were passed
	// in the RateLimitRequest message.
	OverallCode RateLimitResponse_Code `protobuf:"varint,1,opt,name=overall_code,json=overallCode,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"overall_code,omitempty"`
	// A list of DescriptorStatus messages which matches the length of the descriptor list passed
	// in the RateLimitRequest. This can be used by the caller to determine which individual
	// descriptors failed and/or what the currently configured limits are for all of them.
	Statuses []*RateLimitResponse_DescriptorStatus `protobuf:"bytes,2,rep,name=statuses" json:"statuses,omitempty"`
}

func (m *RateLimitResponse) Reset()                    { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string            { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()               {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) { return fileDescriptorRls, []int{1} }

func (m *RateLimitResponse) GetOverallCode() RateLimitResponse_Code {
	if m != nil {
		return m.OverallCode
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse) GetStatuses() []*RateLimitResponse_DescriptorStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

// Defines an actual rate limit in terms of requests per unit of time and the unit itself.
type RateLimitResponse_RateLimit struct {
	RequestsPerUnit uint32                           `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	Unit            RateLimitResponse_RateLimit_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit" json:"unit,omitempty"`
}

func (m *RateLimitResponse_RateLimit) Reset()         { *m = RateLimitResponse_RateLimit{} }
func (m *RateLimitResponse_RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_RateLimit) ProtoMessage()    {}
func (*RateLimitResponse_RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptorRls, []int{1, 0}
}

func (m *RateLimitResponse_RateLimit) GetRequestsPerUnit() uint32 {
	if m != nil {
		return m.RequestsPerUnit
	}
	return 0
}

func (m *RateLimitResponse_RateLimit) GetUnit() RateLimitResponse_RateLimit_Unit {
	if m != nil {
		return m.Unit
	}
	return RateLimitResponse_RateLimit_UNKNOWN
}

type RateLimitResponse_DescriptorStatus struct {
	// The response code for an individual descriptor.
	Code RateLimitResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"code,omitempty"`
	// The current limit as configured by the server. Useful for debugging, etc.
	CurrentLimit *RateLimitResponse_RateLimit `protobuf:"bytes,2,opt,name=current_limit,json=currentLimit" json:"current_limit,omitempty"`
	// The limit remaining in the current time unit.
	LimitRemaining uint32 `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining,proto3" json:"limit_remaining,omitempty"`
}

func (m *RateLimitResponse_DescriptorStatus) Reset()         { *m = RateLimitResponse_DescriptorStatus{} }
func (m *RateLimitResponse_DescriptorStatus) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_DescriptorStatus) ProtoMessage()    {}
func (*RateLimitResponse_DescriptorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptorRls, []int{1, 1}
}

func (m *RateLimitResponse_DescriptorStatus) GetCode() RateLimitResponse_Code {
	if m != nil {
		return m.Code
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse_DescriptorStatus) GetCurrentLimit() *RateLimitResponse_RateLimit {
	if m != nil {
		return m.CurrentLimit
	}
	return nil
}

func (m *RateLimitResponse_DescriptorStatus) GetLimitRemaining() uint32 {
	if m != nil {
		return m.LimitRemaining
	}
	return 0
}

func init() {
	proto.RegisterType((*RateLimitRequest)(nil), "envoy.service.ratelimit.v2.RateLimitRequest")
	proto.RegisterType((*RateLimitResponse)(nil), "envoy.service.ratelimit.v2.RateLimitResponse")
	proto.RegisterType((*RateLimitResponse_RateLimit)(nil), "envoy.service.ratelimit.v2.RateLimitResponse.RateLimit")
	proto.RegisterType((*RateLimitResponse_DescriptorStatus)(nil), "envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus")
	proto.RegisterEnum("envoy.service.ratelimit.v2.RateLimitResponse_Code", RateLimitResponse_Code_name, RateLimitResponse_Code_value)
	proto.RegisterEnum("envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit", RateLimitResponse_RateLimit_Unit_name, RateLimitResponse_RateLimit_Unit_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RateLimitService service

type RateLimitServiceClient interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitServiceClient(cc *grpc.ClientConn) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := grpc.Invoke(ctx, "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RateLimitService service

type RateLimitServiceServer interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ratelimit.v2.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/ratelimit/v2/rls.proto",
}

func (m *RateLimitRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Domain) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRls(dAtA, i, uint64(len(m.Domain)))
		i += copy(dAtA[i:], m.Domain)
	}
	if len(m.Descriptors) > 0 {
		for _, msg := range m.Descriptors {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRls(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.HitsAddend != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.HitsAddend))
	}
	return i, nil
}

func (m *RateLimitResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OverallCode != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.OverallCode))
	}
	if len(m.Statuses) > 0 {
		for _, msg := range m.Statuses {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRls(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *RateLimitResponse_RateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitResponse_RateLimit) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RequestsPerUnit != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.RequestsPerUnit))
	}
	if m.Unit != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.Unit))
	}
	return i, nil
}

func (m *RateLimitResponse_DescriptorStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitResponse_DescriptorStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.Code))
	}
	if m.CurrentLimit != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.CurrentLimit.Size()))
		n1, err := m.CurrentLimit.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.LimitRemaining != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRls(dAtA, i, uint64(m.LimitRemaining))
	}
	return i, nil
}

func encodeVarintRls(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RateLimitRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovRls(uint64(l))
	}
	if len(m.Descriptors) > 0 {
		for _, e := range m.Descriptors {
			l = e.Size()
			n += 1 + l + sovRls(uint64(l))
		}
	}
	if m.HitsAddend != 0 {
		n += 1 + sovRls(uint64(m.HitsAddend))
	}
	return n
}

func (m *RateLimitResponse) Size() (n int) {
	var l int
	_ = l
	if m.OverallCode != 0 {
		n += 1 + sovRls(uint64(m.OverallCode))
	}
	if len(m.Statuses) > 0 {
		for _, e := range m.Statuses {
			l = e.Size()
			n += 1 + l + sovRls(uint64(l))
		}
	}
	return n
}

func (m *RateLimitResponse_RateLimit) Size() (n int) {
	var l int
	_ = l
	if m.RequestsPerUnit != 0 {
		n += 1 + sovRls(uint64(m.RequestsPerUnit))
	}
	if m.Unit != 0 {
		n += 1 + sovRls(uint64(m.Unit))
	}
	return n
}

func (m *RateLimitResponse_DescriptorStatus) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovRls(uint64(m.Code))
	}
	if m.CurrentLimit != nil {
		l = m.CurrentLimit.Size()
		n += 1 + l + sovRls(uint64(l))
	}
	if m.LimitRemaining != 0 {
		n += 1 + sovRls(uint64(m.LimitRemaining))
	}
	return n
}

func sovRls(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRls(x uint64) (n int) {
	return sovRls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimitRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimitRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimitRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRls
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Descriptors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Descriptors = append(m.Descriptors, &envoy_api_v2_ratelimit.RateLimitDescriptor{})
			if err := m.Descriptors[len(m.Descriptors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HitsAddend", wireType)
			}
			m.HitsAddend = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HitsAddend |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RateLimitResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OverallCode", wireType)
			}
			m.OverallCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OverallCode |= (RateLimitResponse_Code(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statuses = append(m.Statuses, &RateLimitResponse_DescriptorStatus{})
			if err := m.Statuses[len(m.Statuses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RateLimitResponse_RateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestsPerUnit", wireType)
			}
			m.RequestsPerUnit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestsPerUnit |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			m.Unit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Unit |= (RateLimitResponse_RateLimit_Unit(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RateLimitResponse_DescriptorStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DescriptorStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DescriptorStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (RateLimitResponse_Code(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CurrentLimit == nil {
				m.CurrentLimit = &RateLimitResponse_RateLimit{}
			}
			if err := m.CurrentLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitRemaining", wireType)
			}
			m.LimitRemaining = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LimitRemaining |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRls
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRls
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRls
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRls
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRls
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRls(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRls = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRls   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/service/ratelimit/v2/rls.proto", fileDescriptorRls) }

var fileDescriptorRls = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xbb, 0x89, 0xbf, 0xb4, 0x19, 0x37, 0x89, 0xbb, 0x87, 0xaf, 0x91, 0x0f, 0x21, 0x8a,
	0x10, 0x44, 0x14, 0x1c, 0xc9, 0x1c, 0xb8, 0xa0, 0x4a, 0xa5, 0x09, 0x22, 0x6a, 0x93, 0x54, 0x9b,
	0x06, 0x44, 0x85, 0x64, 0x99, 0x78, 0x45, 0x57, 0x72, 0xbd, 0x66, 0x77, 0x63, 0x89, 0x3b, 0x4f,
	0xc1, 0x95, 0x97, 0x81, 0x1b, 0x8f, 0x00, 0x79, 0x12, 0xe4, 0xb5, 0xe3, 0x04, 0x10, 0x12, 0x81,
	0xdb, 0xee, 0xcc, 0xfc, 0x7f, 0x9a, 0xff, 0xec, 0x2c, 0xdc, 0xa6, 0x51, 0xc2, 0xdf, 0xf5, 0x24,
	0x15, 0x09, 0x9b, 0xd3, 0x9e, 0xf0, 0x15, 0x0d, 0xd9, 0x0d, 0x53, 0xbd, 0xc4, 0xed, 0x89, 0x50,
	0x3a, 0xb1, 0xe0, 0x8a, 0x63, 0x5b, 0x57, 0x39, 0x79, 0x95, 0x53, 0x54, 0x39, 0x89, 0x6b, 0xdf,
	0xc9, 0x08, 0x7e, 0xcc, 0xb4, 0xa6, 0x00, 0xac, 0x8b, 0x34, 0xc3, 0x3e, 0x4c, 0xfc, 0x90, 0x05,
	0xbe, 0xa2, 0xbd, 0xd5, 0x21, 0x4b, 0x74, 0x3e, 0x20, 0xb0, 0x88, 0xaf, 0xe8, 0x79, 0x5a, 0x4c,
	0xe8, 0xdb, 0x05, 0x95, 0x0a, 0xff, 0x0f, 0x95, 0x80, 0xdf, 0xf8, 0x2c, 0x6a, 0xa2, 0x36, 0xea,
	0x56, 0x49, 0x7e, 0xc3, 0x23, 0x30, 0x03, 0x2a, 0xe7, 0x82, 0xc5, 0x8a, 0x0b, 0xd9, 0x2c, 0xb5,
	0xcb, 0x5d, 0xd3, 0x3d, 0x72, 0xb2, 0xfe, 0xfc, 0x98, 0x39, 0x89, 0xbb, 0xd1, 0x5e, 0x81, 0xed,
	0x17, 0x1a, 0xb2, 0xa9, 0xc7, 0xb7, 0xc0, 0xbc, 0x66, 0x4a, 0x7a, 0x7e, 0x10, 0xd0, 0x28, 0x68,
	0x96, 0xdb, 0xa8, 0x5b, 0x23, 0x90, 0x86, 0x4e, 0x74, 0xa4, 0xf3, 0xf1, 0x3f, 0x38, 0xd8, 0x68,
	0x4e, 0xc6, 0x3c, 0x92, 0x14, 0xcf, 0x60, 0x9f, 0x27, 0x54, 0xf8, 0x61, 0xe8, 0xcd, 0x79, 0x40,
	0x75, 0x8f, 0x75, 0xd7, 0x75, 0x7e, 0x3f, 0x26, 0xe7, 0x17, 0x88, 0x73, 0xca, 0x03, 0x4a, 0xcc,
	0x9c, 0x93, 0x5e, 0xf0, 0x15, 0xec, 0x49, 0xe5, 0xab, 0x85, 0xa4, 0x2b, 0x67, 0xc7, 0xdb, 0x21,
	0xd7, 0x36, 0xa7, 0x9a, 0x43, 0x0a, 0x9e, 0xfd, 0x19, 0x41, 0xb5, 0x10, 0xe0, 0x7b, 0x70, 0x20,
	0xb2, 0x49, 0x4b, 0x2f, 0xa6, 0xc2, 0x5b, 0x44, 0x4c, 0x69, 0x17, 0x35, 0xd2, 0x58, 0x25, 0x2e,
	0xa8, 0x98, 0x45, 0x4c, 0xe1, 0x0b, 0x30, 0x74, 0xba, 0xa4, 0x4d, 0x3e, 0xde, 0xae, 0xa3, 0x22,
	0xe2, 0xa4, 0x2c, 0xa2, 0x49, 0x9d, 0x63, 0x30, 0x34, 0xd9, 0x84, 0xdd, 0xd9, 0xf8, 0x6c, 0x3c,
	0x79, 0x31, 0xb6, 0x76, 0x30, 0x40, 0x65, 0x3a, 0x38, 0x9d, 0x8c, 0xfb, 0x16, 0x4a, 0xcf, 0xa3,
	0xe1, 0x78, 0x76, 0x39, 0xb0, 0x4a, 0x78, 0x0f, 0x8c, 0x67, 0x93, 0x19, 0xb1, 0xca, 0x78, 0x17,
	0xca, 0xfd, 0x93, 0x97, 0x96, 0x61, 0x7f, 0x43, 0x60, 0xfd, 0x6c, 0x15, 0x3f, 0x05, 0xe3, 0x1f,
	0xdf, 0x42, 0xeb, 0xf1, 0x2b, 0xa8, 0xcd, 0x17, 0x42, 0xd0, 0x48, 0x79, 0x5a, 0xa0, 0x7d, 0x9b,
	0xee, 0xa3, 0xbf, 0xf4, 0x4d, 0xf6, 0x73, 0x5a, 0x36, 0xf8, 0xbb, 0xd0, 0xd0, 0x2a, 0x4f, 0xd0,
	0x74, 0x9f, 0x59, 0xf4, 0x26, 0x5f, 0xba, 0x7a, 0x98, 0xe9, 0xf3, 0x68, 0xe7, 0x08, 0x0c, 0xbd,
	0x13, 0x3f, 0xcc, 0xa8, 0x02, 0xa5, 0xc9, 0x99, 0x85, 0x70, 0x1d, 0x60, 0xf2, 0x7c, 0x40, 0xbc,
	0xf3, 0xe1, 0x68, 0x78, 0x69, 0x95, 0xdc, 0xf7, 0x9b, 0x5f, 0x68, 0x9a, 0x75, 0x88, 0x63, 0x68,
	0x4c, 0xaf, 0xf9, 0x22, 0x0c, 0xd6, 0xcf, 0x7e, 0xff, 0x0f, 0x4d, 0xe8, 0x05, 0xb0, 0x1f, 0x6c,
	0x65, 0xb9, 0xb3, 0xf3, 0xe4, 0xf0, 0xd3, 0xb2, 0x85, 0xbe, 0x2c, 0x5b, 0xe8, 0xeb, 0xb2, 0x85,
	0xae, 0xaa, 0x45, 0xfd, 0xeb, 0x8a, 0xfe, 0xe9, 0x0f, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x80,
	0x3f, 0xc8, 0xaa, 0x6e, 0x04, 0x00, 0x00,
}
