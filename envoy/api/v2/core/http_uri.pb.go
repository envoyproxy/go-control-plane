// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/core/http_uri.proto

package core

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Envoy external URI descriptor
type HttpUri struct {
	// The HTTP server URI. It should be a full FQDN with protocol, host and path.
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//    uri: https://www.googleapis.com/oauth2/v1/certs
	//
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Specify how `uri` is to be fetched. Today, this requires an explicit
	// cluster, but in the future we may support dynamic cluster creation or
	// inline DNS resolution. See `issue
	// <https://github.com/envoyproxy/envoy/issues/1606>`_.
	//
	// Types that are valid to be assigned to HttpUpstreamType:
	//	*HttpUri_Cluster
	HttpUpstreamType isHttpUri_HttpUpstreamType `protobuf_oneof:"http_upstream_type"`
	// Sets the maximum duration in milliseconds that a response can take to arrive upon request.
	Timeout              *time.Duration `protobuf:"bytes,3,opt,name=timeout,proto3,stdduration" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HttpUri) Reset()         { *m = HttpUri{} }
func (m *HttpUri) String() string { return proto.CompactTextString(m) }
func (*HttpUri) ProtoMessage()    {}
func (*HttpUri) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_uri_a8729bd15bcc4a68, []int{0}
}
func (m *HttpUri) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpUri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpUri.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HttpUri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpUri.Merge(dst, src)
}
func (m *HttpUri) XXX_Size() int {
	return m.Size()
}
func (m *HttpUri) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpUri.DiscardUnknown(m)
}

var xxx_messageInfo_HttpUri proto.InternalMessageInfo

type isHttpUri_HttpUpstreamType interface {
	isHttpUri_HttpUpstreamType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type HttpUri_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*HttpUri_Cluster) isHttpUri_HttpUpstreamType() {}

func (m *HttpUri) GetHttpUpstreamType() isHttpUri_HttpUpstreamType {
	if m != nil {
		return m.HttpUpstreamType
	}
	return nil
}

func (m *HttpUri) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *HttpUri) GetCluster() string {
	if x, ok := m.GetHttpUpstreamType().(*HttpUri_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *HttpUri) GetTimeout() *time.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HttpUri) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HttpUri_OneofMarshaler, _HttpUri_OneofUnmarshaler, _HttpUri_OneofSizer, []interface{}{
		(*HttpUri_Cluster)(nil),
	}
}

func _HttpUri_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HttpUri)
	// http_upstream_type
	switch x := m.HttpUpstreamType.(type) {
	case *HttpUri_Cluster:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Cluster)
	case nil:
	default:
		return fmt.Errorf("HttpUri.HttpUpstreamType has unexpected type %T", x)
	}
	return nil
}

func _HttpUri_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HttpUri)
	switch tag {
	case 2: // http_upstream_type.cluster
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.HttpUpstreamType = &HttpUri_Cluster{x}
		return true, err
	default:
		return false, nil
	}
}

func _HttpUri_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HttpUri)
	// http_upstream_type
	switch x := m.HttpUpstreamType.(type) {
	case *HttpUri_Cluster:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Cluster)))
		n += len(x.Cluster)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*HttpUri)(nil), "envoy.api.v2.core.HttpUri")
}
func (m *HttpUri) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpUri) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uri) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHttpUri(dAtA, i, uint64(len(m.Uri)))
		i += copy(dAtA[i:], m.Uri)
	}
	if m.HttpUpstreamType != nil {
		nn1, err := m.HttpUpstreamType.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.Timeout != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHttpUri(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*m.Timeout)))
		n2, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.Timeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *HttpUri_Cluster) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x12
	i++
	i = encodeVarintHttpUri(dAtA, i, uint64(len(m.Cluster)))
	i += copy(dAtA[i:], m.Cluster)
	return i, nil
}
func encodeVarintHttpUri(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HttpUri) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovHttpUri(uint64(l))
	}
	if m.HttpUpstreamType != nil {
		n += m.HttpUpstreamType.Size()
	}
	if m.Timeout != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.Timeout)
		n += 1 + l + sovHttpUri(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HttpUri_Cluster) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cluster)
	n += 1 + l + sovHttpUri(uint64(l))
	return n
}

func sovHttpUri(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHttpUri(x uint64) (n int) {
	return sovHttpUri(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HttpUri) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpUri
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
			return fmt.Errorf("proto: HttpUri: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpUri: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpUri
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
				return ErrInvalidLengthHttpUri
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpUri
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
				return ErrInvalidLengthHttpUri
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HttpUpstreamType = &HttpUri_Cluster{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpUri
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
				return ErrInvalidLengthHttpUri
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timeout == nil {
				m.Timeout = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.Timeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttpUri(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpUri
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHttpUri(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttpUri
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
					return 0, ErrIntOverflowHttpUri
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
					return 0, ErrIntOverflowHttpUri
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
				return 0, ErrInvalidLengthHttpUri
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHttpUri
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
				next, err := skipHttpUri(dAtA[start:])
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
	ErrInvalidLengthHttpUri = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttpUri   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/api/v2/core/http_uri.proto", fileDescriptor_http_uri_a8729bd15bcc4a68)
}

var fileDescriptor_http_uri_a8729bd15bcc4a68 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4e, 0xf3, 0x30,
	0x14, 0xc7, 0xfb, 0x92, 0xaf, 0x5f, 0xa8, 0x61, 0xc1, 0x42, 0x22, 0xb4, 0x52, 0x88, 0x90, 0x90,
	0x3a, 0xd9, 0x52, 0xb8, 0x41, 0xc4, 0xd0, 0x85, 0x25, 0x12, 0x73, 0xe5, 0xb6, 0x26, 0x58, 0x4a,
	0x6b, 0xcb, 0x7d, 0x8e, 0xd4, 0x9b, 0x70, 0x06, 0x4e, 0x00, 0x4c, 0x1d, 0x19, 0xb9, 0x01, 0x28,
	0x5b, 0x6f, 0x81, 0x92, 0xd4, 0x0b, 0xdb, 0x4f, 0xef, 0xfd, 0xfc, 0xf7, 0x5f, 0x8f, 0xa4, 0x72,
	0x53, 0xeb, 0x1d, 0x17, 0x46, 0xf1, 0x3a, 0xe3, 0x4b, 0x6d, 0x25, 0x7f, 0x46, 0x34, 0x73, 0x67,
	0x15, 0x33, 0x56, 0xa3, 0xa6, 0xe7, 0x9d, 0xc1, 0x84, 0x51, 0xac, 0xce, 0x58, 0x6b, 0x8c, 0x93,
	0x52, 0xeb, 0xb2, 0x92, 0xbc, 0x13, 0x16, 0xee, 0x89, 0xaf, 0x9c, 0x15, 0xa8, 0xf4, 0xa6, 0x7f,
	0x32, 0xbe, 0x28, 0x75, 0xa9, 0x3b, 0xe4, 0x2d, 0x1d, 0xa7, 0x97, 0xb5, 0xa8, 0xd4, 0x4a, 0xa0,
	0xe4, 0x1e, 0xfa, 0xc5, 0xcd, 0x3b, 0x90, 0x68, 0x86, 0x68, 0x1e, 0xad, 0xa2, 0x13, 0x12, 0x3a,
	0xab, 0x62, 0x48, 0x61, 0x3a, 0xca, 0x47, 0x1f, 0x87, 0x7d, 0xf8, 0xcf, 0x06, 0x29, 0x14, 0xed,
	0x94, 0xde, 0x92, 0x68, 0x59, 0xb9, 0x2d, 0x4a, 0x1b, 0x07, 0x7f, 0x84, 0xd9, 0xa0, 0xf0, 0x3b,
	0xfa, 0x40, 0x22, 0x54, 0x6b, 0xa9, 0x1d, 0xc6, 0x61, 0x0a, 0xd3, 0xd3, 0xec, 0x8a, 0xf5, 0x85,
	0x99, 0x2f, 0xcc, 0xee, 0x8f, 0x85, 0xf3, 0xf8, 0xe5, 0xfb, 0x1a, 0xda, 0x94, 0xe1, 0x2b, 0x04,
	0xd9, 0xc0, 0xd3, 0x09, 0x14, 0x3e, 0x23, 0x9f, 0x10, 0xda, 0x9f, 0xc4, 0x6c, 0xd1, 0x4a, 0xb1,
	0x9e, 0xe3, 0xce, 0x48, 0x3a, 0x7c, 0x3b, 0xec, 0x43, 0xc8, 0xcf, 0x3e, 0x9b, 0x04, 0xbe, 0x9a,
	0x04, 0x7e, 0x9a, 0x04, 0x16, 0xff, 0xbb, 0x0f, 0xee, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xee,
	0x33, 0x45, 0x4d, 0x56, 0x01, 0x00, 0x00,
}
