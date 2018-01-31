// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/filter/http/gzip.proto

package http

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT  Gzip_CompressionStrategy = 0
	Gzip_FILTERED Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN  Gzip_CompressionStrategy = 2
	Gzip_RLE      Gzip_CompressionStrategy = 3
)

var Gzip_CompressionStrategy_name = map[int32]string{
	0: "DEFAULT",
	1: "FILTERED",
	2: "HUFFMAN",
	3: "RLE",
}
var Gzip_CompressionStrategy_value = map[string]int32{
	"DEFAULT":  0,
	"FILTERED": 1,
	"HUFFMAN":  2,
	"RLE":      3,
}

func (x Gzip_CompressionStrategy) String() string {
	return proto.EnumName(Gzip_CompressionStrategy_name, int32(x))
}
func (Gzip_CompressionStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorGzip, []int{0, 0}
}

type Gzip_CompressionLevel_Enum int32

const (
	Gzip_CompressionLevel_DEFAULT Gzip_CompressionLevel_Enum = 0
	Gzip_CompressionLevel_BEST    Gzip_CompressionLevel_Enum = 1
	Gzip_CompressionLevel_SPEED   Gzip_CompressionLevel_Enum = 2
)

var Gzip_CompressionLevel_Enum_name = map[int32]string{
	0: "DEFAULT",
	1: "BEST",
	2: "SPEED",
}
var Gzip_CompressionLevel_Enum_value = map[string]int32{
	"DEFAULT": 0,
	"BEST":    1,
	"SPEED":   2,
}

func (x Gzip_CompressionLevel_Enum) String() string {
	return proto.EnumName(Gzip_CompressionLevel_Enum_name, int32(x))
}
func (Gzip_CompressionLevel_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorGzip, []int{0, 0, 0}
}

type Gzip struct {
	// Value from 1 to 9 that controls the amount of internal memory used by zlib. Higher values
	// use more memory, but are faster and produce better compression results. Default value is 5.
	MemoryLevel *google_protobuf1.UInt32Value `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel" json:"memory_level,omitempty"`
	// Minimum response length, in bytes, which will trigger compression. Default value is 30.
	ContentLength *google_protobuf1.UInt32Value `protobuf:"bytes,2,opt,name=content_length,json=contentLength" json:"content_length,omitempty"`
	// A value used for selecting zlib's compression level. This setting will affect speed and amount
	// of compression applied to the content. "BEST" provides higher compression at the cost of higher
	// latency, "SPEED" provides lower compression with minimum impact on response time. "DEFAULT"
	// provides an optimal result between speed and compression. This field will be set to "DEFAULT"
	// if not specified.
	CompressionLevel Gzip_CompressionLevel_Enum `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.api.v2.filter.http.Gzip_CompressionLevel_Enum" json:"compression_level,omitempty"`
	// A value used for selecting zlib's compression strategy which is directly related to the
	// characteristics of the content. Most of the time "DEFAULT" will be the best choice, though
	// there are situations which changing this parameter might produce better results. For example,
	// run-length encoding (RLE) is typically used when the content is known for having sequences
	// which same data occurs many consecutive times. For more information about each strategy, please
	// refer to zlib manual.
	CompressionStrategy Gzip_CompressionStrategy `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.api.v2.filter.http.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	// Set of strings that allows specifying which mime-types yield compression; e.g.,
	// application/json, text/html, etc. When this field is not defined, compression will be applied
	// to the following mime-types: "application/javascript", "application/json",
	// "application/xhtml+xml", "image/svg+xml", "text/css", "text/html", "text/plain", "text/xml".
	ContentType []string `protobuf:"bytes,6,rep,name=content_type,json=contentType" json:"content_type,omitempty"`
	// If true, disables compression when the response contains Etag header. When it is false, the
	// filter will preserve weak Etags and remove the ones that require strong validation. Note that
	// if any of these options fit the specific scenario, an alternative solution is to disabled Etag
	// at the origin and use *last-modified* header instead.
	DisableOnEtagHeader bool `protobuf:"varint,7,opt,name=disable_on_etag_header,json=disableOnEtagHeader,proto3" json:"disable_on_etag_header,omitempty"`
	// If true, disables compression when response contains *last-modified*
	// header.
	DisableOnLastModifiedHeader bool `protobuf:"varint,8,opt,name=disable_on_last_modified_header,json=disableOnLastModifiedHeader,proto3" json:"disable_on_last_modified_header,omitempty"`
	// Value from 9 to 15 that represents the base two logarithmic of the compressor's window size.
	// Larger window results in better compression at the expense of memory usage. The default is 12
	// which will produce a 4096 bytes window. For more details about this parameter, please refer to
	// zlib manual > deflateInit2.
	WindowBits *google_protobuf1.UInt32Value `protobuf:"bytes,9,opt,name=window_bits,json=windowBits" json:"window_bits,omitempty"`
}

func (m *Gzip) Reset()                    { *m = Gzip{} }
func (m *Gzip) String() string            { return proto.CompactTextString(m) }
func (*Gzip) ProtoMessage()               {}
func (*Gzip) Descriptor() ([]byte, []int) { return fileDescriptorGzip, []int{0} }

func (m *Gzip) GetMemoryLevel() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MemoryLevel
	}
	return nil
}

func (m *Gzip) GetContentLength() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.ContentLength
	}
	return nil
}

func (m *Gzip) GetCompressionLevel() Gzip_CompressionLevel_Enum {
	if m != nil {
		return m.CompressionLevel
	}
	return Gzip_CompressionLevel_DEFAULT
}

func (m *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if m != nil {
		return m.CompressionStrategy
	}
	return Gzip_DEFAULT
}

func (m *Gzip) GetContentType() []string {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Gzip) GetDisableOnEtagHeader() bool {
	if m != nil {
		return m.DisableOnEtagHeader
	}
	return false
}

func (m *Gzip) GetDisableOnLastModifiedHeader() bool {
	if m != nil {
		return m.DisableOnLastModifiedHeader
	}
	return false
}

func (m *Gzip) GetWindowBits() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.WindowBits
	}
	return nil
}

type Gzip_CompressionLevel struct {
}

func (m *Gzip_CompressionLevel) Reset()                    { *m = Gzip_CompressionLevel{} }
func (m *Gzip_CompressionLevel) String() string            { return proto.CompactTextString(m) }
func (*Gzip_CompressionLevel) ProtoMessage()               {}
func (*Gzip_CompressionLevel) Descriptor() ([]byte, []int) { return fileDescriptorGzip, []int{0, 0} }

func init() {
	proto.RegisterType((*Gzip)(nil), "envoy.api.v2.filter.http.Gzip")
	proto.RegisterType((*Gzip_CompressionLevel)(nil), "envoy.api.v2.filter.http.Gzip.CompressionLevel")
	proto.RegisterEnum("envoy.api.v2.filter.http.Gzip_CompressionStrategy", Gzip_CompressionStrategy_name, Gzip_CompressionStrategy_value)
	proto.RegisterEnum("envoy.api.v2.filter.http.Gzip_CompressionLevel_Enum", Gzip_CompressionLevel_Enum_name, Gzip_CompressionLevel_Enum_value)
}
func (m *Gzip) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gzip) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MemoryLevel != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.MemoryLevel.Size()))
		n1, err := m.MemoryLevel.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.ContentLength != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.ContentLength.Size()))
		n2, err := m.ContentLength.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.CompressionLevel != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.CompressionLevel))
	}
	if m.CompressionStrategy != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.CompressionStrategy))
	}
	if len(m.ContentType) > 0 {
		for _, s := range m.ContentType {
			dAtA[i] = 0x32
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
	if m.DisableOnEtagHeader {
		dAtA[i] = 0x38
		i++
		if m.DisableOnEtagHeader {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DisableOnLastModifiedHeader {
		dAtA[i] = 0x40
		i++
		if m.DisableOnLastModifiedHeader {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.WindowBits != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.WindowBits.Size()))
		n3, err := m.WindowBits.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *Gzip_CompressionLevel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gzip_CompressionLevel) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintGzip(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Gzip) Size() (n int) {
	var l int
	_ = l
	if m.MemoryLevel != nil {
		l = m.MemoryLevel.Size()
		n += 1 + l + sovGzip(uint64(l))
	}
	if m.ContentLength != nil {
		l = m.ContentLength.Size()
		n += 1 + l + sovGzip(uint64(l))
	}
	if m.CompressionLevel != 0 {
		n += 1 + sovGzip(uint64(m.CompressionLevel))
	}
	if m.CompressionStrategy != 0 {
		n += 1 + sovGzip(uint64(m.CompressionStrategy))
	}
	if len(m.ContentType) > 0 {
		for _, s := range m.ContentType {
			l = len(s)
			n += 1 + l + sovGzip(uint64(l))
		}
	}
	if m.DisableOnEtagHeader {
		n += 2
	}
	if m.DisableOnLastModifiedHeader {
		n += 2
	}
	if m.WindowBits != nil {
		l = m.WindowBits.Size()
		n += 1 + l + sovGzip(uint64(l))
	}
	return n
}

func (m *Gzip_CompressionLevel) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovGzip(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGzip(x uint64) (n int) {
	return sovGzip(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Gzip) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGzip
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
			return fmt.Errorf("proto: Gzip: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gzip: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryLevel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
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
				return ErrInvalidLengthGzip
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MemoryLevel == nil {
				m.MemoryLevel = &google_protobuf1.UInt32Value{}
			}
			if err := m.MemoryLevel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentLength", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
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
				return ErrInvalidLengthGzip
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ContentLength == nil {
				m.ContentLength = &google_protobuf1.UInt32Value{}
			}
			if err := m.ContentLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompressionLevel", wireType)
			}
			m.CompressionLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompressionLevel |= (Gzip_CompressionLevel_Enum(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompressionStrategy", wireType)
			}
			m.CompressionStrategy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompressionStrategy |= (Gzip_CompressionStrategy(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
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
				return ErrInvalidLengthGzip
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentType = append(m.ContentType, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableOnEtagHeader", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableOnEtagHeader = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableOnLastModifiedHeader", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableOnLastModifiedHeader = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowBits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
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
				return ErrInvalidLengthGzip
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WindowBits == nil {
				m.WindowBits = &google_protobuf1.UInt32Value{}
			}
			if err := m.WindowBits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGzip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGzip
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
func (m *Gzip_CompressionLevel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGzip
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
			return fmt.Errorf("proto: CompressionLevel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompressionLevel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGzip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGzip
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
func skipGzip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGzip
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
					return 0, ErrIntOverflowGzip
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
					return 0, ErrIntOverflowGzip
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
				return 0, ErrInvalidLengthGzip
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGzip
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
				next, err := skipGzip(dAtA[start:])
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
	ErrInvalidLengthGzip = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGzip   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/api/v2/filter/http/gzip.proto", fileDescriptorGzip) }

var fileDescriptorGzip = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0xbf, 0x71, 0xdc, 0x34, 0x1e, 0xe7, 0x2b, 0x66, 0x82, 0xc0, 0x0a, 0x28, 0x44, 0x61,
	0x63, 0x55, 0x62, 0x2c, 0x39, 0x2c, 0xd9, 0x34, 0xc4, 0xa1, 0x45, 0x6e, 0x41, 0x4e, 0xc2, 0xd6,
	0x9a, 0xc4, 0x13, 0x67, 0x24, 0xdb, 0x63, 0xec, 0x49, 0xa2, 0x74, 0xd9, 0x47, 0xe0, 0x71, 0x58,
	0xb1, 0x64, 0xc9, 0x23, 0xa0, 0xec, 0x78, 0x0b, 0xe4, 0x3f, 0x81, 0xb6, 0x80, 0x54, 0x76, 0x96,
	0xce, 0xf9, 0x9d, 0x33, 0xf7, 0x5e, 0xc3, 0x67, 0x34, 0x5e, 0xf3, 0xad, 0x49, 0x12, 0x66, 0xae,
	0x2d, 0x73, 0xc1, 0x42, 0x41, 0x53, 0x73, 0x29, 0x44, 0x62, 0x06, 0x97, 0x2c, 0xc1, 0x49, 0xca,
	0x05, 0x47, 0x7a, 0x61, 0xc2, 0x24, 0x61, 0x78, 0x6d, 0xe1, 0xd2, 0x84, 0x73, 0x53, 0xbb, 0x13,
	0x70, 0x1e, 0x84, 0xd4, 0x2c, 0x7c, 0xb3, 0xd5, 0xc2, 0xdc, 0xa4, 0x24, 0x49, 0x68, 0x9a, 0x95,
	0x64, 0xfb, 0xd1, 0x9a, 0x84, 0xcc, 0x27, 0x82, 0x9a, 0xfb, 0x8f, 0x52, 0xe8, 0x5d, 0xd5, 0xa1,
	0xfc, 0xfa, 0x92, 0x25, 0xc8, 0x81, 0xcd, 0x88, 0x46, 0x3c, 0xdd, 0x7a, 0x21, 0x5d, 0xd3, 0x50,
	0x07, 0x5d, 0x60, 0xa8, 0xd6, 0x13, 0x5c, 0x06, 0xe3, 0x7d, 0x30, 0x9e, 0x9e, 0xc5, 0xa2, 0x6f,
	0xbd, 0x27, 0xe1, 0x8a, 0x0e, 0xd4, 0x4f, 0xdf, 0x3f, 0xd7, 0xea, 0xc7, 0xb2, 0xae, 0x18, 0xc0,
	0x55, 0x4b, 0xdc, 0xc9, 0x69, 0x74, 0x01, 0x8f, 0xe6, 0x3c, 0x16, 0x34, 0x16, 0x5e, 0x48, 0xe3,
	0x40, 0x2c, 0x75, 0xe9, 0x0e, 0x79, 0x4a, 0x9e, 0x27, 0x1f, 0x4b, 0x46, 0xc7, 0xfd, 0xbf, 0xc2,
	0x9d, 0x82, 0x46, 0x11, 0xbc, 0x3f, 0xe7, 0x51, 0x92, 0xd2, 0x2c, 0x63, 0x3c, 0xae, 0x9e, 0x58,
	0xeb, 0x02, 0xe3, 0xc8, 0x7a, 0x81, 0xff, 0xb6, 0x15, 0x9c, 0x0f, 0x86, 0x5f, 0xfd, 0xe2, 0x8a,
	0xb7, 0x61, 0x3b, 0x5e, 0x45, 0x03, 0x98, 0x57, 0x1d, 0x5c, 0x01, 0x49, 0x03, 0xae, 0x36, 0xbf,
	0x65, 0x41, 0x1f, 0xe0, 0x83, 0xeb, 0x75, 0x99, 0x48, 0x89, 0xa0, 0xc1, 0x56, 0x97, 0x8b, 0x46,
	0xeb, 0xee, 0x8d, 0xe3, 0x8a, 0xbc, 0xd1, 0xd7, 0x9a, 0xff, 0x6e, 0x40, 0xcf, 0x61, 0x73, 0xbf,
	0x31, 0xb1, 0x4d, 0xa8, 0x5e, 0xef, 0xd6, 0x0c, 0xa5, 0xc2, 0x3e, 0x02, 0x49, 0xb3, 0x5c, 0xb5,
	0xd2, 0x27, 0xdb, 0x84, 0xa2, 0x3e, 0x7c, 0xe8, 0xb3, 0x8c, 0xcc, 0x42, 0xea, 0xf1, 0xd8, 0xa3,
	0x82, 0x04, 0xde, 0x92, 0x12, 0x9f, 0xa6, 0xfa, 0x61, 0x17, 0x18, 0x0d, 0xb7, 0x55, 0xa9, 0x6f,
	0x63, 0x5b, 0x90, 0xe0, 0xb4, 0x90, 0xd0, 0x10, 0x3e, 0xbd, 0x06, 0x85, 0x24, 0x13, 0x5e, 0xc4,
	0x7d, 0xb6, 0x60, 0xd4, 0xdf, 0xd3, 0x8d, 0x82, 0x7e, 0xfc, 0x93, 0x76, 0x48, 0x26, 0xce, 0x2b,
	0x4f, 0x95, 0xf2, 0x06, 0xaa, 0x1b, 0x16, 0xfb, 0x7c, 0xe3, 0xcd, 0x98, 0xc8, 0x74, 0xe5, 0x5f,
	0x7e, 0x94, 0x7b, 0x86, 0xe2, 0xc2, 0x92, 0x1e, 0x30, 0x91, 0xb5, 0x5f, 0x42, 0xed, 0xf6, 0x7d,
	0x7a, 0x06, 0x94, 0xf3, 0x13, 0x21, 0x15, 0x1e, 0x0e, 0xed, 0xd1, 0xc9, 0xd4, 0x99, 0x68, 0xff,
	0xa1, 0x06, 0x94, 0x07, 0xf6, 0x78, 0xa2, 0x01, 0xa4, 0xc0, 0x83, 0xf1, 0x3b, 0xdb, 0x1e, 0x6a,
	0x52, 0x6f, 0x04, 0x5b, 0x7f, 0xd8, 0xf5, 0x4d, 0xb0, 0x09, 0x1b, 0xa3, 0x33, 0x67, 0x62, 0xbb,
	0xf6, 0x50, 0x03, 0xb9, 0x74, 0x3a, 0x1d, 0x8d, 0xce, 0x4f, 0x2e, 0x34, 0x09, 0x1d, 0xc2, 0x9a,
	0xeb, 0xd8, 0x5a, 0x6d, 0xd0, 0xfc, 0xb2, 0xeb, 0x80, 0xaf, 0xbb, 0x0e, 0xf8, 0xb6, 0xeb, 0x80,
	0x59, 0xbd, 0x18, 0xa1, 0xff, 0x23, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xe4, 0x3e, 0xec, 0x93, 0x03,
	0x00, 0x00,
}
