// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/csrf/v2/csrf.proto

package v2

import (
	fmt "fmt"
	io "io"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// CSRF filter config.
type CsrfPolicy struct {
	// Specify if CSRF is enabled.
	//
	// More information on how this can be controlled via runtime can be found
	// :ref:`here <csrf-runtime>`.
	//
	// .. note::
	//
	//   This field defaults to 100/:ref:`HUNDRED
	//   <envoy_api_enum_type.FractionalPercent.DenominatorType>`.
	FilterEnabled *core.RuntimeFractionalPercent `protobuf:"bytes,1,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	// Specifies that CSRF policies will be evaluated and tracked, but not enforced.
	// This is intended to be used when filter_enabled is off.
	//
	// More information on how this can be controlled via runtime can be found
	// :ref:`here <csrf-runtime>`.
	//
	// .. note::
	//
	//   This field defaults to 100/:ref:`HUNDRED
	//   <envoy_api_enum_type.FractionalPercent.DenominatorType>`.
	ShadowEnabled        *core.RuntimeFractionalPercent `protobuf:"bytes,2,opt,name=shadow_enabled,json=shadowEnabled,proto3" json:"shadow_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *CsrfPolicy) Reset()         { *m = CsrfPolicy{} }
func (m *CsrfPolicy) String() string { return proto.CompactTextString(m) }
func (*CsrfPolicy) ProtoMessage()    {}
func (*CsrfPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9146cdf92353980, []int{0}
}
func (m *CsrfPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CsrfPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CsrfPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CsrfPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsrfPolicy.Merge(m, src)
}
func (m *CsrfPolicy) XXX_Size() int {
	return m.Size()
}
func (m *CsrfPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_CsrfPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_CsrfPolicy proto.InternalMessageInfo

func (m *CsrfPolicy) GetFilterEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.FilterEnabled
	}
	return nil
}

func (m *CsrfPolicy) GetShadowEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.ShadowEnabled
	}
	return nil
}

func init() {
	proto.RegisterType((*CsrfPolicy)(nil), "envoy.config.filter.http.csrf.v2.CsrfPolicy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/csrf/v2/csrf.proto", fileDescriptor_a9146cdf92353980)
}

var fileDescriptor_a9146cdf92353980 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0xc9, 0x88, 0x2e, 0x22, 0x33, 0xca, 0x20, 0x28, 0x83, 0x94, 0xc1, 0x95, 0x30, 0xf0,
	0x02, 0xf5, 0x06, 0x15, 0x5d, 0x97, 0xba, 0x73, 0xa1, 0xa4, 0x6d, 0xda, 0x09, 0xd4, 0xbc, 0x92,
	0xc6, 0x68, 0xaf, 0xe0, 0x91, 0x5c, 0xcd, 0xd2, 0xa5, 0x47, 0x90, 0xee, 0xbc, 0x85, 0x34, 0x6f,
	0xd4, 0xa5, 0xb8, 0xea, 0x4f, 0xf3, 0xbd, 0xef, 0x25, 0x3f, 0x5f, 0x29, 0xe3, 0xb1, 0x17, 0x05,
	0x9a, 0x4a, 0xd7, 0xa2, 0xd2, 0x8d, 0x53, 0x56, 0xac, 0x9d, 0x6b, 0x45, 0xd1, 0xd9, 0x4a, 0xf8,
	0x38, 0x7c, 0xa1, 0xb5, 0xe8, 0x70, 0xbe, 0x0c, 0x30, 0x10, 0x0c, 0x04, 0xc3, 0x08, 0x43, 0x80,
	0x7c, 0xbc, 0x38, 0x25, 0x9d, 0x6c, 0x75, 0x18, 0x45, 0xab, 0x44, 0x2e, 0x3b, 0x45, 0xf3, 0x8b,
	0x63, 0x2f, 0x1b, 0x5d, 0x4a, 0xa7, 0xc4, 0x77, 0xd8, 0x1e, 0x1c, 0xd5, 0x58, 0x63, 0x88, 0x62,
	0x4c, 0xf4, 0xf7, 0x6c, 0xc3, 0x38, 0xbf, 0xec, 0x6c, 0x95, 0x62, 0xa3, 0x8b, 0x7e, 0x7e, 0xc7,
	0x67, 0xb4, 0xf2, 0x5e, 0x19, 0x99, 0x37, 0xaa, 0x3c, 0x61, 0x4b, 0x76, 0xbe, 0x1f, 0xaf, 0x80,
	0xae, 0x25, 0x5b, 0x0d, 0x3e, 0x86, 0x71, 0x29, 0x64, 0x8f, 0xc6, 0xe9, 0x07, 0x75, 0x6d, 0x65,
	0xe1, 0x34, 0x1a, 0xd9, 0xa4, 0xca, 0x16, 0xca, 0xb8, 0x84, 0xbf, 0x7e, 0x6e, 0x76, 0x76, 0x5f,
	0xd8, 0xe4, 0x90, 0x65, 0x53, 0xd2, 0x5d, 0x91, 0x6d, 0x9e, 0xf1, 0x59, 0xb7, 0x96, 0x25, 0x3e,
	0xfd, 0xf8, 0x27, 0xff, 0xf6, 0x67, 0x53, 0x52, 0x6c, 0x9d, 0xc9, 0xcd, 0xdb, 0x10, 0xb1, 0xf7,
	0x21, 0x62, 0x1f, 0x43, 0xc4, 0x38, 0x68, 0x24, 0x57, 0x6b, 0xf1, 0xb9, 0x87, 0xbf, 0xda, 0x4c,
	0x0e, 0x7e, 0x5f, 0x9f, 0x8e, 0x8d, 0xa4, 0xec, 0x76, 0xe2, 0xe3, 0x7c, 0x2f, 0xd4, 0x73, 0xf1,
	0x15, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x93, 0x22, 0x93, 0xbc, 0x01, 0x00, 0x00,
}

func (m *CsrfPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CsrfPolicy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FilterEnabled != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCsrf(dAtA, i, uint64(m.FilterEnabled.Size()))
		n1, err := m.FilterEnabled.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.ShadowEnabled != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCsrf(dAtA, i, uint64(m.ShadowEnabled.Size()))
		n2, err := m.ShadowEnabled.MarshalTo(dAtA[i:])
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

func encodeVarintCsrf(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CsrfPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FilterEnabled != nil {
		l = m.FilterEnabled.Size()
		n += 1 + l + sovCsrf(uint64(l))
	}
	if m.ShadowEnabled != nil {
		l = m.ShadowEnabled.Size()
		n += 1 + l + sovCsrf(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCsrf(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCsrf(x uint64) (n int) {
	return sovCsrf(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CsrfPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCsrf
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CsrfPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CsrfPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilterEnabled", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsrf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCsrf
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCsrf
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FilterEnabled == nil {
				m.FilterEnabled = &core.RuntimeFractionalPercent{}
			}
			if err := m.FilterEnabled.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShadowEnabled", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsrf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCsrf
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCsrf
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ShadowEnabled == nil {
				m.ShadowEnabled = &core.RuntimeFractionalPercent{}
			}
			if err := m.ShadowEnabled.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCsrf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCsrf
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCsrf
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
func skipCsrf(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCsrf
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
					return 0, ErrIntOverflowCsrf
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
					return 0, ErrIntOverflowCsrf
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
			if length < 0 {
				return 0, ErrInvalidLengthCsrf
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCsrf
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCsrf
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
				next, err := skipCsrf(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCsrf
				}
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
	ErrInvalidLengthCsrf = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCsrf   = fmt.Errorf("proto: integer overflow")
)
