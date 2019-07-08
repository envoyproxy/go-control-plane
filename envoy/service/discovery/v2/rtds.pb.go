// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/discovery/v2/rtds.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
import _ "github.com/gogo/googleapis/google/api"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221
type RtdsDummy struct {
}

func (m *RtdsDummy) Reset()                    { *m = RtdsDummy{} }
func (m *RtdsDummy) String() string            { return proto.CompactTextString(m) }
func (*RtdsDummy) ProtoMessage()               {}
func (*RtdsDummy) Descriptor() ([]byte, []int) { return fileDescriptorRtds, []int{0} }

// RTDS resource type. This describes a layer in the runtime virtual filesystem.
type Runtime struct {
	// Runtime resource name. This makes the Runtime a self-describing xDS
	// resource.
	Name  string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Layer *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=layer" json:"layer,omitempty"`
}

func (m *Runtime) Reset()                    { *m = Runtime{} }
func (m *Runtime) String() string            { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()               {}
func (*Runtime) Descriptor() ([]byte, []int) { return fileDescriptorRtds, []int{1} }

func (m *Runtime) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Runtime) GetLayer() *google_protobuf1.Struct {
	if m != nil {
		return m.Layer
	}
	return nil
}

func init() {
	proto.RegisterType((*RtdsDummy)(nil), "envoy.service.discovery.v2.RtdsDummy")
	proto.RegisterType((*Runtime)(nil), "envoy.service.discovery.v2.Runtime")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RuntimeDiscoveryService service

type RuntimeDiscoveryServiceClient interface {
	StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error)
	DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error)
	FetchRuntime(ctx context.Context, in *envoy_api_v2.DiscoveryRequest, opts ...grpc.CallOption) (*envoy_api_v2.DiscoveryResponse, error)
}

type runtimeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRuntimeDiscoveryServiceClient(cc *grpc.ClientConn) RuntimeDiscoveryServiceClient {
	return &runtimeDiscoveryServiceClient{cc}
}

func (c *runtimeDiscoveryServiceClient) StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RuntimeDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.service.discovery.v2.RuntimeDiscoveryService/StreamRuntime", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceStreamRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_StreamRuntimeClient interface {
	Send(*envoy_api_v2.DiscoveryRequest) error
	Recv() (*envoy_api_v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceStreamRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Send(m *envoy_api_v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Recv() (*envoy_api_v2.DiscoveryResponse, error) {
	m := new(envoy_api_v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RuntimeDiscoveryService_serviceDesc.Streams[1], c.cc, "/envoy.service.discovery.v2.RuntimeDiscoveryService/DeltaRuntime", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceDeltaRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_DeltaRuntimeClient interface {
	Send(*envoy_api_v2.DeltaDiscoveryRequest) error
	Recv() (*envoy_api_v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceDeltaRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Send(m *envoy_api_v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Recv() (*envoy_api_v2.DeltaDiscoveryResponse, error) {
	m := new(envoy_api_v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) FetchRuntime(ctx context.Context, in *envoy_api_v2.DiscoveryRequest, opts ...grpc.CallOption) (*envoy_api_v2.DiscoveryResponse, error) {
	out := new(envoy_api_v2.DiscoveryResponse)
	err := grpc.Invoke(ctx, "/envoy.service.discovery.v2.RuntimeDiscoveryService/FetchRuntime", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RuntimeDiscoveryService service

type RuntimeDiscoveryServiceServer interface {
	StreamRuntime(RuntimeDiscoveryService_StreamRuntimeServer) error
	DeltaRuntime(RuntimeDiscoveryService_DeltaRuntimeServer) error
	FetchRuntime(context.Context, *envoy_api_v2.DiscoveryRequest) (*envoy_api_v2.DiscoveryResponse, error)
}

func RegisterRuntimeDiscoveryServiceServer(s *grpc.Server, srv RuntimeDiscoveryServiceServer) {
	s.RegisterService(&_RuntimeDiscoveryService_serviceDesc, srv)
}

func _RuntimeDiscoveryService_StreamRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).StreamRuntime(&runtimeDiscoveryServiceStreamRuntimeServer{stream})
}

type RuntimeDiscoveryService_StreamRuntimeServer interface {
	Send(*envoy_api_v2.DiscoveryResponse) error
	Recv() (*envoy_api_v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceStreamRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Send(m *envoy_api_v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Recv() (*envoy_api_v2.DiscoveryRequest, error) {
	m := new(envoy_api_v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_DeltaRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).DeltaRuntime(&runtimeDiscoveryServiceDeltaRuntimeServer{stream})
}

type RuntimeDiscoveryService_DeltaRuntimeServer interface {
	Send(*envoy_api_v2.DeltaDiscoveryResponse) error
	Recv() (*envoy_api_v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceDeltaRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Send(m *envoy_api_v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Recv() (*envoy_api_v2.DeltaDiscoveryRequest, error) {
	m := new(envoy_api_v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_FetchRuntime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(envoy_api_v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v2.RuntimeDiscoveryService/FetchRuntime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, req.(*envoy_api_v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RuntimeDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.RuntimeDiscoveryService",
	HandlerType: (*RuntimeDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRuntime",
			Handler:    _RuntimeDiscoveryService_FetchRuntime_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRuntime",
			Handler:       _RuntimeDiscoveryService_StreamRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRuntime",
			Handler:       _RuntimeDiscoveryService_DeltaRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/rtds.proto",
}

func (m *RtdsDummy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RtdsDummy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Runtime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Runtime) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRtds(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Layer != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRtds(dAtA, i, uint64(m.Layer.Size()))
		n1, err := m.Layer.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintRtds(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RtdsDummy) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Runtime) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRtds(uint64(l))
	}
	if m.Layer != nil {
		l = m.Layer.Size()
		n += 1 + l + sovRtds(uint64(l))
	}
	return n
}

func sovRtds(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRtds(x uint64) (n int) {
	return sovRtds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RtdsDummy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRtds
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
			return fmt.Errorf("proto: RtdsDummy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RtdsDummy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRtds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRtds
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
func (m *Runtime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRtds
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
			return fmt.Errorf("proto: Runtime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Runtime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRtds
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
				return ErrInvalidLengthRtds
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Layer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRtds
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
				return ErrInvalidLengthRtds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Layer == nil {
				m.Layer = &google_protobuf1.Struct{}
			}
			if err := m.Layer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRtds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRtds
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
func skipRtds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRtds
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
					return 0, ErrIntOverflowRtds
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
					return 0, ErrIntOverflowRtds
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
				return 0, ErrInvalidLengthRtds
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRtds
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
				next, err := skipRtds(dAtA[start:])
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
	ErrInvalidLengthRtds = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRtds   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/service/discovery/v2/rtds.proto", fileDescriptorRtds) }

var fileDescriptorRtds = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xbf, 0x6a, 0xe3, 0x30,
	0x1c, 0x3e, 0xf9, 0x72, 0x77, 0x44, 0xc9, 0x2d, 0x86, 0x23, 0xc1, 0xe4, 0x7c, 0xc6, 0xd7, 0x82,
	0x29, 0x54, 0x2e, 0xee, 0x96, 0x31, 0x84, 0xcc, 0xc1, 0x81, 0x76, 0x2c, 0x8a, 0xad, 0xa6, 0x02,
	0xdb, 0x72, 0x24, 0xd9, 0xd4, 0x6b, 0xa7, 0xee, 0x7d, 0x9b, 0x4e, 0x5d, 0x0a, 0x1d, 0x0b, 0x7d,
	0x81, 0x12, 0xba, 0xf4, 0x2d, 0x4a, 0xa4, 0x38, 0x25, 0x29, 0xe9, 0xd4, 0x4d, 0xd2, 0xf7, 0x4f,
	0xfa, 0x7d, 0x82, 0xfb, 0x24, 0x2b, 0x59, 0xe5, 0x0b, 0xc2, 0x4b, 0x1a, 0x11, 0x3f, 0xa6, 0x22,
	0x62, 0x25, 0xe1, 0x95, 0x5f, 0x06, 0x3e, 0x97, 0xb1, 0x40, 0x39, 0x67, 0x92, 0x99, 0x96, 0xa2,
	0xa1, 0x15, 0x0d, 0xad, 0x69, 0xa8, 0x0c, 0xac, 0x9e, 0xb6, 0xc0, 0x39, 0x5d, 0x8a, 0xde, 0x21,
	0xa5, 0xb4, 0x7a, 0x33, 0xc6, 0x66, 0x09, 0x51, 0x30, 0xce, 0x32, 0x26, 0xb1, 0xa4, 0x2c, 0x13,
	0x5b, 0xa8, 0xda, 0x4d, 0x8b, 0x73, 0x5f, 0x48, 0x5e, 0x44, 0x72, 0x85, 0x76, 0x4a, 0x9c, 0xd0,
	0x18, 0x4b, 0xe2, 0xd7, 0x0b, 0x0d, 0xb8, 0x2d, 0xd8, 0x0c, 0x65, 0x2c, 0x86, 0x45, 0x9a, 0x56,
	0xee, 0x29, 0xfc, 0x15, 0x16, 0x99, 0xa4, 0x29, 0x31, 0xff, 0xc2, 0x46, 0x86, 0x53, 0xd2, 0x05,
	0x0e, 0xf0, 0x9a, 0x83, 0xe6, 0xed, 0xeb, 0xdd, 0xf7, 0x06, 0x37, 0x1c, 0x10, 0xaa, 0x63, 0xf3,
	0x10, 0xfe, 0x48, 0x70, 0x45, 0x78, 0xd7, 0x70, 0x80, 0xd7, 0x0a, 0x3a, 0x48, 0xa7, 0xa3, 0x3a,
	0x1d, 0x4d, 0x54, 0x7a, 0xa8, 0x59, 0xc1, 0xbd, 0x01, 0x3b, 0x2b, 0xe7, 0x61, 0xfd, 0xaa, 0x89,
	0x9e, 0x80, 0x79, 0x02, 0x7f, 0x4f, 0x24, 0x27, 0x38, 0xad, 0xa3, 0x6d, 0xa4, 0x47, 0x84, 0x73,
	0x8a, 0xca, 0x00, 0xad, 0x05, 0x21, 0x99, 0x17, 0x44, 0x48, 0xeb, 0xdf, 0x4e, 0x5c, 0xe4, 0x2c,
	0x13, 0xc4, 0xfd, 0xe6, 0x81, 0x23, 0x60, 0x9e, 0xc1, 0xf6, 0x90, 0x24, 0x12, 0xd7, 0xb6, 0xff,
	0xb7, 0x64, 0x4b, 0xec, 0x83, 0xf7, 0xde, 0xe7, 0xa4, 0x8d, 0x80, 0x39, 0x6c, 0x8f, 0x88, 0x8c,
	0x2e, 0xbe, 0xec, 0xde, 0xce, 0xd5, 0xd3, 0xcb, 0x8d, 0x61, 0xb9, 0x7f, 0x36, 0xda, 0xef, 0x73,
	0xed, 0xdf, 0x07, 0x07, 0x83, 0xd1, 0xc3, 0xc2, 0x06, 0x8f, 0x0b, 0x1b, 0x3c, 0x2f, 0x6c, 0x00,
	0x3d, 0xca, 0xb4, 0x65, 0xce, 0xd9, 0x65, 0x85, 0x76, 0x7f, 0xac, 0x81, 0xea, 0x78, 0xbc, 0xec,
	0x66, 0x0c, 0xae, 0x01, 0x98, 0xfe, 0x54, 0x3d, 0x1d, 0xbf, 0x05, 0x00, 0x00, 0xff, 0xff, 0xbd,
	0xb2, 0xe2, 0x9d, 0xb4, 0x02, 0x00, 0x00,
}
