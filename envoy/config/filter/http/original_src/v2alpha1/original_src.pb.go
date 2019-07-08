// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/original_src/v2alpha1/original_src.proto

/*
	Package v2alpha1 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/http/original_src/v2alpha1/original_src.proto

	It has these top-level messages:
		OriginalSrc
*/
package v2alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"

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

// The Original Src filter binds upstream connections to the original source address determined
// for the request. This address could come from something like the Proxy Protocol filter, or it
// could come from trusted http headers.
type OriginalSrc struct {
	// Sets the SO_MARK option on the upstream connection's socket to the provided value. Used to
	// ensure that non-local addresses may be routed back through envoy when binding to the original
	// source address. The option will not be applied if the mark is 0.
	// [#proto-status: experimental]
	Mark uint32 `protobuf:"varint,1,opt,name=mark,proto3" json:"mark,omitempty"`
}

func (m *OriginalSrc) Reset()                    { *m = OriginalSrc{} }
func (m *OriginalSrc) String() string            { return proto.CompactTextString(m) }
func (*OriginalSrc) ProtoMessage()               {}
func (*OriginalSrc) Descriptor() ([]byte, []int) { return fileDescriptorOriginalSrc, []int{0} }

func (m *OriginalSrc) GetMark() uint32 {
	if m != nil {
		return m.Mark
	}
	return 0
}

func init() {
	proto.RegisterType((*OriginalSrc)(nil), "envoy.config.filter.http.original_src.v2alpha1.OriginalSrc")
}
func (m *OriginalSrc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OriginalSrc) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mark != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOriginalSrc(dAtA, i, uint64(m.Mark))
	}
	return i, nil
}

func encodeVarintOriginalSrc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *OriginalSrc) Size() (n int) {
	var l int
	_ = l
	if m.Mark != 0 {
		n += 1 + sovOriginalSrc(uint64(m.Mark))
	}
	return n
}

func sovOriginalSrc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOriginalSrc(x uint64) (n int) {
	return sovOriginalSrc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OriginalSrc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOriginalSrc
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
			return fmt.Errorf("proto: OriginalSrc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OriginalSrc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mark", wireType)
			}
			m.Mark = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOriginalSrc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mark |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOriginalSrc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOriginalSrc
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
func skipOriginalSrc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOriginalSrc
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
					return 0, ErrIntOverflowOriginalSrc
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
					return 0, ErrIntOverflowOriginalSrc
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
				return 0, ErrInvalidLengthOriginalSrc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOriginalSrc
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
				next, err := skipOriginalSrc(dAtA[start:])
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
	ErrInvalidLengthOriginalSrc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOriginalSrc   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/original_src/v2alpha1/original_src.proto", fileDescriptorOriginalSrc)
}

var fileDescriptorOriginalSrc = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0xcf, 0x2f, 0xca, 0x4c, 0xcf, 0xcc, 0x4b, 0xcc, 0x89, 0x2f, 0x2e,
	0x4a, 0xd6, 0x2f, 0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0x44, 0x11, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xd2, 0x03, 0x1b, 0xa1, 0x07, 0x31, 0x42, 0x0f, 0x62, 0x84, 0x1e, 0xc8, 0x08,
	0x3d, 0x14, 0xc5, 0x30, 0x23, 0xa4, 0xc4, 0xcb, 0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5,
	0x61, 0x0c, 0x88, 0x41, 0x4a, 0x8a, 0x5c, 0xdc, 0xfe, 0x50, 0x1d, 0xc1, 0x45, 0xc9, 0x42, 0x42,
	0x5c, 0x2c, 0xb9, 0x89, 0x45, 0xd9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x60, 0xb6, 0x53,
	0xc6, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xc8, 0x65, 0x93,
	0x99, 0x0f, 0xb1, 0xbc, 0xa0, 0x28, 0xbf, 0xa2, 0x92, 0x44, 0x77, 0x38, 0x09, 0x20, 0x59, 0x16,
	0x00, 0x72, 0x40, 0x00, 0x63, 0x14, 0x07, 0x4c, 0x36, 0x89, 0x0d, 0xec, 0x26, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbf, 0xc8, 0x17, 0xc9, 0x21, 0x01, 0x00, 0x00,
}
