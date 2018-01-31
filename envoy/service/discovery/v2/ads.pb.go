// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/discovery/v2/ads.proto

/*
	Package discovery is a generated protocol buffer package.

	It is generated from these files:
		envoy/service/discovery/v2/ads.proto
		envoy/service/discovery/v2/cds.proto
		envoy/service/discovery/v2/common.proto
		envoy/service/discovery/v2/eds.proto
		envoy/service/discovery/v2/hds.proto
		envoy/service/discovery/v2/lds.proto
		envoy/service/discovery/v2/rds.proto
		envoy/service/discovery/v2/sds.proto

	It has these top-level messages:
		AdsDummy
		CdsDummy
		DiscoveryRequest
		DiscoveryResponse
		EdsDummy
		ClusterLoadAssignment
		HdsDummy
		Capability
		HealthCheckRequest
		EndpointHealth
		EndpointHealthResponse
		HealthCheckRequestOrEndpointHealthResponse
		LocalityEndpoints
		ClusterHealthCheck
		HealthCheckSpecifier
		LdsDummy
		RdsDummy
		SdsDummy
*/
package discovery

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221
type AdsDummy struct {
}

func (m *AdsDummy) Reset()                    { *m = AdsDummy{} }
func (m *AdsDummy) String() string            { return proto.CompactTextString(m) }
func (*AdsDummy) ProtoMessage()               {}
func (*AdsDummy) Descriptor() ([]byte, []int) { return fileDescriptorAds, []int{0} }

func init() {
	proto.RegisterType((*AdsDummy)(nil), "envoy.service.discovery.v2.AdsDummy")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AggregatedDiscoveryService service

type AggregatedDiscoveryServiceClient interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error)
}

type aggregatedDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewAggregatedDiscoveryServiceClient(cc *grpc.ClientConn) AggregatedDiscoveryServiceClient {
	return &aggregatedDiscoveryServiceClient{cc}
}

func (c *aggregatedDiscoveryServiceClient) StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.service.discovery.v2.AggregatedDiscoveryService/StreamAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceStreamAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_StreamAggregatedResourcesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AggregatedDiscoveryService service

type AggregatedDiscoveryServiceServer interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(AggregatedDiscoveryService_StreamAggregatedResourcesServer) error
}

func RegisterAggregatedDiscoveryServiceServer(s *grpc.Server, srv AggregatedDiscoveryServiceServer) {
	s.RegisterService(&_AggregatedDiscoveryService_serviceDesc, srv)
}

func _AggregatedDiscoveryService_StreamAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).StreamAggregatedResources(&aggregatedDiscoveryServiceStreamAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_StreamAggregatedResourcesServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AggregatedDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.AggregatedDiscoveryService",
	HandlerType: (*AggregatedDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAggregatedResources",
			Handler:       _AggregatedDiscoveryService_StreamAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/ads.proto",
}

func (m *AdsDummy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdsDummy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintAds(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AdsDummy) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovAds(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAds(x uint64) (n int) {
	return sovAds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AdsDummy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAds
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
			return fmt.Errorf("proto: AdsDummy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdsDummy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAds
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
func skipAds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAds
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
					return 0, ErrIntOverflowAds
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
					return 0, ErrIntOverflowAds
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
				return 0, ErrInvalidLengthAds
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAds
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
				next, err := skipAds(dAtA[start:])
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
	ErrInvalidLengthAds = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAds   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/service/discovery/v2/ads.proto", fileDescriptorAds) }

var fileDescriptorAds = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0xc9, 0x2c, 0x4e, 0xce, 0x2f,
	0x4b, 0x2d, 0xaa, 0xd4, 0x2f, 0x33, 0xd2, 0x4f, 0x4c, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x02, 0xab, 0xd2, 0x83, 0xaa, 0xd2, 0x83, 0xab, 0xd2, 0x2b, 0x33, 0x92, 0x52, 0xc7,
	0x63, 0x42, 0x72, 0x7e, 0x6e, 0x6e, 0x7e, 0x1e, 0xc4, 0x10, 0x25, 0x2e, 0x2e, 0x0e, 0xc7, 0x94,
	0x62, 0x97, 0xd2, 0xdc, 0xdc, 0x4a, 0xa3, 0x39, 0x8c, 0x5c, 0x52, 0x8e, 0xe9, 0xe9, 0x45, 0xa9,
	0xe9, 0x89, 0x25, 0xa9, 0x29, 0x2e, 0x30, 0x3d, 0xc1, 0x10, 0x43, 0x84, 0xea, 0xb8, 0x24, 0x83,
	0x4b, 0x8a, 0x52, 0x13, 0x73, 0x11, 0x6a, 0x82, 0x52, 0x8b, 0xf3, 0x4b, 0x8b, 0x92, 0x53, 0x8b,
	0x85, 0x74, 0xf4, 0x70, 0xbb, 0x46, 0x0f, 0x6e, 0x54, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89,
	0x94, 0x2e, 0x91, 0xaa, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0x18, 0x34, 0x18, 0x0d, 0x18,
	0x9d, 0x78, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x24,
	0x36, 0xb0, 0xfb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xe6, 0x61, 0x44, 0x2c, 0x01,
	0x00, 0x00,
}
