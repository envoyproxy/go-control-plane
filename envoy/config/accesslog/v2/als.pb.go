// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/accesslog/v2/als.proto

/*
	Package accesslog is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/accesslog/v2/als.proto

	It has these top-level messages:
		TcpGrpcAccessLogConfig
		HttpGrpcAccessLogConfig
		CommonGrpcAccessLogConfig
*/
package accesslog

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v22 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
import _ "github.com/lyft/protoc-gen-validate/validate"

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

// Configuration for the built-in *envoy.tcp_grpc_access_log* type. This configuration will
// populate *StreamAccessLogsMessage.tcp_logs*.
// [#not-implemented-hide:]
// [#comment:TODO(mattklein123): Block type in non-tcp proxy cases?]
type TcpGrpcAccessLogConfig struct {
	CommonConfig *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig" json:"common_config,omitempty"`
}

func (m *TcpGrpcAccessLogConfig) Reset()                    { *m = TcpGrpcAccessLogConfig{} }
func (m *TcpGrpcAccessLogConfig) String() string            { return proto.CompactTextString(m) }
func (*TcpGrpcAccessLogConfig) ProtoMessage()               {}
func (*TcpGrpcAccessLogConfig) Descriptor() ([]byte, []int) { return fileDescriptorAls, []int{0} }

func (m *TcpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

// Configuration for the built-in *envoy.http_grpc_access_log* type. This configuration will
// populate *StreamAccessLogsMessage.http_logs*.
// [#not-implemented-hide:]
// [#comment:TODO(mattklein123): Block type in non-http/router proxy cases?]
type HttpGrpcAccessLogConfig struct {
	CommonConfig *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig" json:"common_config,omitempty"`
	// Additional request headers to log in *HTTPRequestProperties.request_headers*.
	AdditionalRequestHeadersToLog []string `protobuf:"bytes,2,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog" json:"additional_request_headers_to_log,omitempty"`
	// Additional response headers to log in *HTTPResponseProperties.response_headers*.
	AdditionalResponseHeadersToLog []string `protobuf:"bytes,3,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog" json:"additional_response_headers_to_log,omitempty"`
}

func (m *HttpGrpcAccessLogConfig) Reset()                    { *m = HttpGrpcAccessLogConfig{} }
func (m *HttpGrpcAccessLogConfig) String() string            { return proto.CompactTextString(m) }
func (*HttpGrpcAccessLogConfig) ProtoMessage()               {}
func (*HttpGrpcAccessLogConfig) Descriptor() ([]byte, []int) { return fileDescriptorAls, []int{1} }

func (m *HttpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

// Common configuration for gRPC access logs.
// [#not-implemented-hide:]
type CommonGrpcAccessLogConfig struct {
	// The friendly name of the access log to be returned in StreamAccessLogsMessage.Identifier. This
	// allows the access log server to differentiate between different access logs coming from the
	// same Envoy.
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// The gRPC service for the access log service.
	GrpcService *envoy_api_v22.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService" json:"grpc_service,omitempty"`
}

func (m *CommonGrpcAccessLogConfig) Reset()                    { *m = CommonGrpcAccessLogConfig{} }
func (m *CommonGrpcAccessLogConfig) String() string            { return proto.CompactTextString(m) }
func (*CommonGrpcAccessLogConfig) ProtoMessage()               {}
func (*CommonGrpcAccessLogConfig) Descriptor() ([]byte, []int) { return fileDescriptorAls, []int{2} }

func (m *CommonGrpcAccessLogConfig) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *CommonGrpcAccessLogConfig) GetGrpcService() *envoy_api_v22.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*TcpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.TcpGrpcAccessLogConfig")
	proto.RegisterType((*HttpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.HttpGrpcAccessLogConfig")
	proto.RegisterType((*CommonGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.CommonGrpcAccessLogConfig")
}
func (m *TcpGrpcAccessLogConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TcpGrpcAccessLogConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CommonConfig != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAls(dAtA, i, uint64(m.CommonConfig.Size()))
		n1, err := m.CommonConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *HttpGrpcAccessLogConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpGrpcAccessLogConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CommonConfig != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAls(dAtA, i, uint64(m.CommonConfig.Size()))
		n2, err := m.CommonConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.AdditionalRequestHeadersToLog) > 0 {
		for _, s := range m.AdditionalRequestHeadersToLog {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.AdditionalResponseHeadersToLog) > 0 {
		for _, s := range m.AdditionalResponseHeadersToLog {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *CommonGrpcAccessLogConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonGrpcAccessLogConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.LogName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAls(dAtA, i, uint64(len(m.LogName)))
		i += copy(dAtA[i:], m.LogName)
	}
	if m.GrpcService != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAls(dAtA, i, uint64(m.GrpcService.Size()))
		n3, err := m.GrpcService.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeVarintAls(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TcpGrpcAccessLogConfig) Size() (n int) {
	var l int
	_ = l
	if m.CommonConfig != nil {
		l = m.CommonConfig.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	return n
}

func (m *HttpGrpcAccessLogConfig) Size() (n int) {
	var l int
	_ = l
	if m.CommonConfig != nil {
		l = m.CommonConfig.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if len(m.AdditionalRequestHeadersToLog) > 0 {
		for _, s := range m.AdditionalRequestHeadersToLog {
			l = len(s)
			n += 1 + l + sovAls(uint64(l))
		}
	}
	if len(m.AdditionalResponseHeadersToLog) > 0 {
		for _, s := range m.AdditionalResponseHeadersToLog {
			l = len(s)
			n += 1 + l + sovAls(uint64(l))
		}
	}
	return n
}

func (m *CommonGrpcAccessLogConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.LogName)
	if l > 0 {
		n += 1 + l + sovAls(uint64(l))
	}
	if m.GrpcService != nil {
		l = m.GrpcService.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	return n
}

func sovAls(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAls(x uint64) (n int) {
	return sovAls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TcpGrpcAccessLogConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
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
			return fmt.Errorf("proto: TcpGrpcAccessLogConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TcpGrpcAccessLogConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CommonConfig == nil {
				m.CommonConfig = &CommonGrpcAccessLogConfig{}
			}
			if err := m.CommonConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
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
func (m *HttpGrpcAccessLogConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
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
			return fmt.Errorf("proto: HttpGrpcAccessLogConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpGrpcAccessLogConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CommonConfig == nil {
				m.CommonConfig = &CommonGrpcAccessLogConfig{}
			}
			if err := m.CommonConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalRequestHeadersToLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalRequestHeadersToLog = append(m.AdditionalRequestHeadersToLog, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalResponseHeadersToLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalResponseHeadersToLog = append(m.AdditionalResponseHeadersToLog, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
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
func (m *CommonGrpcAccessLogConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
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
			return fmt.Errorf("proto: CommonGrpcAccessLogConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonGrpcAccessLogConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrpcService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GrpcService == nil {
				m.GrpcService = &envoy_api_v22.GrpcService{}
			}
			if err := m.GrpcService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
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
func skipAls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAls
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
					return 0, ErrIntOverflowAls
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
					return 0, ErrIntOverflowAls
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
				return 0, ErrInvalidLengthAls
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAls
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
				next, err := skipAls(dAtA[start:])
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
	ErrInvalidLengthAls = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAls   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/config/accesslog/v2/als.proto", fileDescriptorAls) }

var fileDescriptorAls = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xc7, 0xe5, 0x94, 0xaf, 0xb8, 0x45, 0x42, 0x19, 0xe8, 0x87, 0x44, 0x28, 0x81, 0xa1, 0x53,
	0x22, 0x15, 0x5e, 0x80, 0x76, 0x20, 0x42, 0x15, 0x43, 0xe8, 0xc4, 0x12, 0x19, 0xc7, 0x98, 0x48,
	0x49, 0x2e, 0xc4, 0x26, 0x12, 0x13, 0x3b, 0xec, 0x3c, 0x0c, 0x13, 0x23, 0x23, 0x8f, 0x80, 0xba,
	0xf1, 0x16, 0x28, 0x36, 0x94, 0x14, 0xd1, 0x99, 0xcd, 0xd2, 0xfd, 0xee, 0xe7, 0xbb, 0xbf, 0x0e,
	0xef, 0xb3, 0xac, 0x84, 0x3b, 0x8f, 0x42, 0x76, 0x15, 0x73, 0x8f, 0x50, 0xca, 0x84, 0x48, 0x80,
	0x7b, 0xe5, 0xd0, 0x23, 0x89, 0x70, 0xf3, 0x02, 0x24, 0x58, 0x5d, 0x05, 0xb9, 0x1a, 0x72, 0xe7,
	0x90, 0x5b, 0x0e, 0x7b, 0xbb, 0xba, 0x9f, 0xe4, 0x71, 0xd5, 0xc2, 0x8b, 0x9c, 0x86, 0x82, 0x15,
	0x65, 0x4c, 0x99, 0xee, 0xed, 0xb5, 0x4b, 0x92, 0xc4, 0x11, 0x91, 0xcc, 0xfb, 0x7e, 0xe8, 0x82,
	0x73, 0x8f, 0xb7, 0xa7, 0x34, 0x3f, 0x29, 0x72, 0x7a, 0xac, 0x84, 0x13, 0xe0, 0x63, 0xf5, 0x81,
	0xc5, 0xf0, 0x26, 0x85, 0x34, 0x85, 0x2c, 0xd4, 0x3f, 0x76, 0x50, 0x1f, 0x0d, 0x9a, 0xc3, 0x23,
	0x77, 0xe9, 0x18, 0xee, 0x58, 0xf1, 0x7f, 0xc8, 0x46, 0xf8, 0xf9, 0xe3, 0xa5, 0xb1, 0xfa, 0x80,
	0x8c, 0x2d, 0x14, 0xb4, 0xb4, 0x56, 0x57, 0x9c, 0x27, 0x03, 0xb7, 0x7d, 0x29, 0xff, 0x71, 0x04,
	0xcb, 0xc7, 0x7b, 0x24, 0x8a, 0x62, 0x19, 0x43, 0x46, 0x92, 0xb0, 0x60, 0x37, 0xb7, 0x4c, 0xc8,
	0xf0, 0x9a, 0x91, 0x88, 0x15, 0x22, 0x94, 0x10, 0x26, 0xc0, 0x3b, 0x46, 0xbf, 0x31, 0x30, 0x83,
	0x9d, 0x1f, 0x30, 0xd0, 0x9c, 0xaf, 0xb1, 0x29, 0x4c, 0x80, 0x5b, 0xa7, 0xd8, 0x59, 0x30, 0x89,
	0x1c, 0x32, 0xc1, 0x7e, 0xab, 0x1a, 0x4a, 0x65, 0xd7, 0x55, 0x1a, 0xac, 0xbb, 0x9c, 0x47, 0x84,
	0xbb, 0x4b, 0xb7, 0xb1, 0x0e, 0xf0, 0x46, 0x02, 0x3c, 0xcc, 0x48, 0xca, 0x54, 0x2a, 0xe6, 0xc8,
	0xac, 0xf6, 0x5b, 0x29, 0x8c, 0x3e, 0x0a, 0xd6, 0x13, 0xe0, 0x67, 0x24, 0x65, 0x96, 0x8f, 0x5b,
	0xf5, 0x63, 0xe8, 0x18, 0x2a, 0xbf, 0xee, 0x57, 0x7e, 0x24, 0x8f, 0xab, 0xc8, 0x2a, 0xfd, 0xb9,
	0x06, 0x16, 0x42, 0x6a, 0xf2, 0x5a, 0xa1, 0xfd, 0x3a, 0xb3, 0xd1, 0xdb, 0xcc, 0x46, 0xef, 0x33,
	0x1b, 0x5d, 0x98, 0xf3, 0xcc, 0x2f, 0xd7, 0xd4, 0x1d, 0x1d, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xcb, 0x60, 0x8b, 0x3e, 0xc3, 0x02, 0x00, 0x00,
}
