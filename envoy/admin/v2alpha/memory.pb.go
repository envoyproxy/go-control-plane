// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v2alpha/memory.proto

package envoy_admin_v2alpha

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Proto representation of the internal memory consumption of an Envoy instance. These represent
// values extracted from an internal TCMalloc instance. For more information, see the section of the
// docs entitled ["Generic Tcmalloc Status"](http://gperftools.github.io/gperftools/tcmalloc.html).
type Memory struct {
	// The number of bytes allocated by the heap for Envoy. This is an alias for
	// `generic.current_allocated_bytes`.
	Allocated uint64 `protobuf:"varint,1,opt,name=allocated,proto3" json:"allocated,omitempty"`
	// The number of bytes reserved by the heap but not necessarily allocated. This is an alias for
	// `generic.heap_size`.
	HeapSize uint64 `protobuf:"varint,2,opt,name=heap_size,json=heapSize,proto3" json:"heap_size,omitempty"`
	// The number of bytes in free, unmapped pages in the page heap. These bytes always count towards
	// virtual memory usage, and depending on the OS, typically do not count towards physical memory
	// usage. This is an alias for `tcmalloc.pageheap_unmapped_bytes`.
	PageheapUnmapped uint64 `protobuf:"varint,3,opt,name=pageheap_unmapped,json=pageheapUnmapped,proto3" json:"pageheap_unmapped,omitempty"`
	// The number of bytes in free, mapped pages in the page heap. These bytes always count towards
	// virtual memory usage, and unless the underlying memory is swapped out by the OS, they also
	// count towards physical memory usage. This is an alias for `tcmalloc.pageheap_free_bytes`.
	PageheapFree uint64 `protobuf:"varint,4,opt,name=pageheap_free,json=pageheapFree,proto3" json:"pageheap_free,omitempty"`
	// The amount of memory used by the TCMalloc thread caches (for small objects). This is an alias
	// for `tcmalloc.current_total_thread_cache_bytes`.
	TotalThreadCache     uint64   `protobuf:"varint,5,opt,name=total_thread_cache,json=totalThreadCache,proto3" json:"total_thread_cache,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Memory) Reset()         { *m = Memory{} }
func (m *Memory) String() string { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()    {}
func (*Memory) Descriptor() ([]byte, []int) {
	return fileDescriptor_memory_154bd10535beba13, []int{0}
}
func (m *Memory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Memory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Memory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Memory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memory.Merge(dst, src)
}
func (m *Memory) XXX_Size() int {
	return m.Size()
}
func (m *Memory) XXX_DiscardUnknown() {
	xxx_messageInfo_Memory.DiscardUnknown(m)
}

var xxx_messageInfo_Memory proto.InternalMessageInfo

func (m *Memory) GetAllocated() uint64 {
	if m != nil {
		return m.Allocated
	}
	return 0
}

func (m *Memory) GetHeapSize() uint64 {
	if m != nil {
		return m.HeapSize
	}
	return 0
}

func (m *Memory) GetPageheapUnmapped() uint64 {
	if m != nil {
		return m.PageheapUnmapped
	}
	return 0
}

func (m *Memory) GetPageheapFree() uint64 {
	if m != nil {
		return m.PageheapFree
	}
	return 0
}

func (m *Memory) GetTotalThreadCache() uint64 {
	if m != nil {
		return m.TotalThreadCache
	}
	return 0
}

func init() {
	proto.RegisterType((*Memory)(nil), "envoy.admin.v2alpha.Memory")
}
func (m *Memory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Memory) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Allocated != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMemory(dAtA, i, uint64(m.Allocated))
	}
	if m.HeapSize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMemory(dAtA, i, uint64(m.HeapSize))
	}
	if m.PageheapUnmapped != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMemory(dAtA, i, uint64(m.PageheapUnmapped))
	}
	if m.PageheapFree != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMemory(dAtA, i, uint64(m.PageheapFree))
	}
	if m.TotalThreadCache != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMemory(dAtA, i, uint64(m.TotalThreadCache))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMemory(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Memory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Allocated != 0 {
		n += 1 + sovMemory(uint64(m.Allocated))
	}
	if m.HeapSize != 0 {
		n += 1 + sovMemory(uint64(m.HeapSize))
	}
	if m.PageheapUnmapped != 0 {
		n += 1 + sovMemory(uint64(m.PageheapUnmapped))
	}
	if m.PageheapFree != 0 {
		n += 1 + sovMemory(uint64(m.PageheapFree))
	}
	if m.TotalThreadCache != 0 {
		n += 1 + sovMemory(uint64(m.TotalThreadCache))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMemory(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMemory(x uint64) (n int) {
	return sovMemory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Memory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMemory
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
			return fmt.Errorf("proto: Memory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Memory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocated", wireType)
			}
			m.Allocated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Allocated |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeapSize", wireType)
			}
			m.HeapSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeapSize |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageheapUnmapped", wireType)
			}
			m.PageheapUnmapped = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageheapUnmapped |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageheapFree", wireType)
			}
			m.PageheapFree = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageheapFree |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalThreadCache", wireType)
			}
			m.TotalThreadCache = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalThreadCache |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMemory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMemory
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
func skipMemory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMemory
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
					return 0, ErrIntOverflowMemory
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
					return 0, ErrIntOverflowMemory
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
				return 0, ErrInvalidLengthMemory
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMemory
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
				next, err := skipMemory(dAtA[start:])
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
	ErrInvalidLengthMemory = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMemory   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/memory.proto", fileDescriptor_memory_154bd10535beba13)
}

var fileDescriptor_memory_154bd10535beba13 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc7, 0x71, 0xa2, 0xeb, 0xe2, 0x0e, 0x0a, 0x1a, 0x2f, 0x01, 0xa5, 0xac, 0x7a, 0x11, 0x94,
	0x14, 0xf5, 0x0d, 0x14, 0xbc, 0x09, 0xe2, 0x9f, 0x73, 0x19, 0xdb, 0xd1, 0x16, 0x92, 0x4e, 0x88,
	0x71, 0xb1, 0xfb, 0x7a, 0x5e, 0x3c, 0xfa, 0x08, 0xd2, 0x27, 0x91, 0x1d, 0xb7, 0x7a, 0xf1, 0xfa,
	0xfd, 0x7d, 0x08, 0x4c, 0x60, 0x4a, 0xed, 0x8c, 0xbb, 0x1c, 0x2b, 0xdf, 0xb4, 0xf9, 0xec, 0x0c,
	0x5d, 0xa8, 0x31, 0xf7, 0xe4, 0x39, 0x76, 0x36, 0x44, 0x4e, 0xac, 0x77, 0x44, 0x58, 0x11, 0x76,
	0x29, 0x0e, 0xde, 0x15, 0x8c, 0xaf, 0x45, 0xe9, 0x3d, 0x98, 0xa0, 0x73, 0x5c, 0x62, 0xa2, 0xca,
	0xa8, 0xa9, 0x3a, 0x1a, 0xdd, 0xfe, 0x05, 0xbd, 0x0b, 0x93, 0x9a, 0x30, 0x14, 0x2f, 0xcd, 0x9c,
	0xcc, 0x8a, 0xac, 0xeb, 0x8b, 0x70, 0xd7, 0xcc, 0x49, 0x1f, 0xc3, 0x76, 0xc0, 0x67, 0x12, 0xf0,
	0xda, 0x7a, 0x0c, 0x81, 0x2a, 0xb3, 0x2a, 0x68, 0x6b, 0x18, 0x1e, 0x96, 0x5d, 0x1f, 0xc2, 0xe6,
	0x2f, 0x7e, 0x8a, 0x44, 0x66, 0x24, 0x70, 0x63, 0x88, 0x57, 0x91, 0x48, 0x9f, 0x80, 0x4e, 0x9c,
	0xd0, 0x15, 0xa9, 0x8e, 0x84, 0x55, 0x51, 0x62, 0x59, 0x93, 0x59, 0xfb, 0x79, 0x52, 0x96, 0x7b,
	0x19, 0x2e, 0x17, 0xfd, 0xe2, 0xf4, 0xa3, 0xcf, 0xd4, 0x67, 0x9f, 0xa9, 0xaf, 0x3e, 0x53, 0xb0,
	0xdf, 0xb0, 0x95, 0x5b, 0x43, 0xe4, 0xb7, 0xce, 0xfe, 0x73, 0xf6, 0x8d, 0x7a, 0x1c, 0xcb, 0xa7,
	0x9c, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xb4, 0x7d, 0x66, 0x38, 0x01, 0x00, 0x00,
}
