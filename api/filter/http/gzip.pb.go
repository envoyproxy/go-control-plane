// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/filter/http/gzip.proto

package envoy_api_v2_filter_http

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

type Gzip_CompressionLevel int32

const (
	Gzip_DEFAULT_LEVEL Gzip_CompressionLevel = 0
	Gzip_BEST          Gzip_CompressionLevel = 1
	Gzip_SPEED         Gzip_CompressionLevel = 2
)

var Gzip_CompressionLevel_name = map[int32]string{
	0: "DEFAULT_LEVEL",
	1: "BEST",
	2: "SPEED",
}
var Gzip_CompressionLevel_value = map[string]int32{
	"DEFAULT_LEVEL": 0,
	"BEST":          1,
	"SPEED":         2,
}

func (x Gzip_CompressionLevel) String() string {
	return proto.EnumName(Gzip_CompressionLevel_name, int32(x))
}
func (Gzip_CompressionLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptorGzip, []int{0, 0} }

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
	return fileDescriptorGzip, []int{0, 1}
}

// Gzip is an HTTP filter which enables Envoy to compress dispatched data from an upstream
// service upon client request. This is useful in situations where large payloads need to
// be transmitted without compromising the response time. Note that when compression is applied,
// this filter will set "content-encoding" and "transfer-encoding" headers to gzip and chunked,
// respectively.
// TODO(gsagula): elaborate the last part in the final documentation.
type Gzip struct {
	// Value from 1 to 9 that controls the amount of internal memory
	// used by zlib. Higher values use more memory, but are faster and produce better compression
	// results. Default value is 8.
	MemoryLevel *google_protobuf1.UInt32Value `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel" json:"memory_level,omitempty"`
	// Minimum response length, in bytes, which will trigger
	// compression. Default value is 30.
	ContentLength *google_protobuf1.UInt32Value `protobuf:"bytes,2,opt,name=content_length,json=contentLength" json:"content_length,omitempty"`
	// Allows selecting Zlib's compression level. This setting will affect
	// speed and amount of compression applied to the content. "BEST" option provides higher
	// compression at cost of higher latency, "SPEED" provides lower compression with minimum impact
	// on response time. "DEFAULT" provides an optimal result between speed and compression. This
	// field will be set to "DEFAULT" if not specified.
	CompressionLevel Gzip_CompressionLevel `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.api.v2.filter.http.Gzip_CompressionLevel" json:"compression_level,omitempty"`
	// Allows selecting zlib's compression strategy. Strategy is directly
	// related to the characteristics of the content which is being compressed. Most of the time
	// "DEFAULT" will be the best choice, however there are situations which changing the strategy
	// might produce better results. For example, Run-length encoding (RLE) is normally used when the
	// content is known for having sequences which same data occurs many consecutive times. For more
	// information about each strategy, please refer to Zlib manual. This field will be set to
	// "DEFAULT" if not specified.
	CompressionStrategy Gzip_CompressionStrategy `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.api.v2.filter.http.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	// Array of strings that allows specifying which "cache-control" header
	// values yield compression. Normally, if "cache-control" is present in the response headers,
	// compression should only occur if directives indicate that the content should not be cached;
	// e.g. no-cache or no-store.
	CacheControl []string `protobuf:"bytes,5,rep,name=cache_control,json=cacheControl" json:"cache_control,omitempty"`
	// Array of strings that allows specifying which mime-types yield compression; e.g.
	// application/json, text/html, etc. When this field is not specified, compression will be applied
	// to any "content-type".
	ContentType []string `protobuf:"bytes,6,rep,name=content_type,json=contentType" json:"content_type,omitempty"`
	// Allows disabling compression if response contains "etag" (entity tag)
	// header. Default is false.
	DisableOnEtag *google_protobuf1.BoolValue `protobuf:"bytes,7,opt,name=disable_on_etag,json=disableOnEtag" json:"disable_on_etag,omitempty"`
	// Allows disabling compression if response contains "last-modified"
	// header. Default is false.
	DisableOnLastModified *google_protobuf1.BoolValue `protobuf:"bytes,8,opt,name=disable_on_last_modified,json=disableOnLastModified" json:"disable_on_last_modified,omitempty"`
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

func (m *Gzip) GetCompressionLevel() Gzip_CompressionLevel {
	if m != nil {
		return m.CompressionLevel
	}
	return Gzip_DEFAULT_LEVEL
}

func (m *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if m != nil {
		return m.CompressionStrategy
	}
	return Gzip_DEFAULT
}

func (m *Gzip) GetCacheControl() []string {
	if m != nil {
		return m.CacheControl
	}
	return nil
}

func (m *Gzip) GetContentType() []string {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Gzip) GetDisableOnEtag() *google_protobuf1.BoolValue {
	if m != nil {
		return m.DisableOnEtag
	}
	return nil
}

func (m *Gzip) GetDisableOnLastModified() *google_protobuf1.BoolValue {
	if m != nil {
		return m.DisableOnLastModified
	}
	return nil
}

func init() {
	proto.RegisterType((*Gzip)(nil), "envoy.api.v2.filter.http.Gzip")
	proto.RegisterEnum("envoy.api.v2.filter.http.Gzip_CompressionLevel", Gzip_CompressionLevel_name, Gzip_CompressionLevel_value)
	proto.RegisterEnum("envoy.api.v2.filter.http.Gzip_CompressionStrategy", Gzip_CompressionStrategy_name, Gzip_CompressionStrategy_value)
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
	if len(m.CacheControl) > 0 {
		for _, s := range m.CacheControl {
			dAtA[i] = 0x2a
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
	if m.DisableOnEtag != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.DisableOnEtag.Size()))
		n3, err := m.DisableOnEtag.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.DisableOnLastModified != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.DisableOnLastModified.Size()))
		n4, err := m.DisableOnLastModified.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
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
	if len(m.CacheControl) > 0 {
		for _, s := range m.CacheControl {
			l = len(s)
			n += 1 + l + sovGzip(uint64(l))
		}
	}
	if len(m.ContentType) > 0 {
		for _, s := range m.ContentType {
			l = len(s)
			n += 1 + l + sovGzip(uint64(l))
		}
	}
	if m.DisableOnEtag != nil {
		l = m.DisableOnEtag.Size()
		n += 1 + l + sovGzip(uint64(l))
	}
	if m.DisableOnLastModified != nil {
		l = m.DisableOnLastModified.Size()
		n += 1 + l + sovGzip(uint64(l))
	}
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
				m.CompressionLevel |= (Gzip_CompressionLevel(b) & 0x7F) << shift
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheControl", wireType)
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
			m.CacheControl = append(m.CacheControl, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableOnEtag", wireType)
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
			if m.DisableOnEtag == nil {
				m.DisableOnEtag = &google_protobuf1.BoolValue{}
			}
			if err := m.DisableOnEtag.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableOnLastModified", wireType)
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
			if m.DisableOnLastModified == nil {
				m.DisableOnLastModified = &google_protobuf1.BoolValue{}
			}
			if err := m.DisableOnLastModified.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

func init() { proto.RegisterFile("api/filter/http/gzip.proto", fileDescriptorGzip) }

var fileDescriptorGzip = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcb, 0x6e, 0xd3, 0x4e,
	0x18, 0xc5, 0x33, 0xb9, 0x67, 0x72, 0xf9, 0xbb, 0xd3, 0x3f, 0x62, 0x14, 0x21, 0x2b, 0xca, 0x2a,
	0xea, 0x62, 0x2c, 0x92, 0x1d, 0xbb, 0xba, 0x71, 0xa0, 0x92, 0x5b, 0x90, 0x93, 0x74, 0x6b, 0x4d,
	0x9c, 0x89, 0x63, 0x69, 0xe2, 0x19, 0xec, 0x69, 0x50, 0xba, 0x64, 0xc1, 0x03, 0xf0, 0x38, 0xac,
	0x58, 0xb2, 0xe4, 0x11, 0x50, 0x76, 0xbc, 0x05, 0xf2, 0x25, 0x50, 0x1a, 0x10, 0xdd, 0x8d, 0xf4,
	0x9d, 0xdf, 0x39, 0xdf, 0x7c, 0x07, 0x76, 0xa9, 0x0c, 0x8c, 0x55, 0xc0, 0x15, 0x8b, 0x8c, 0xb5,
	0x52, 0xd2, 0xf0, 0xef, 0x02, 0x49, 0x64, 0x24, 0x94, 0x40, 0x98, 0x85, 0x5b, 0xb1, 0x23, 0x54,
	0x06, 0x64, 0x3b, 0x24, 0x99, 0x88, 0x24, 0xa2, 0xae, 0xee, 0x0b, 0xe1, 0x73, 0x66, 0xa4, 0xba,
	0xc5, 0xed, 0xca, 0x78, 0x17, 0x51, 0x29, 0x59, 0x14, 0x67, 0x64, 0xf7, 0xe9, 0x96, 0xf2, 0x60,
	0x49, 0x15, 0x33, 0x0e, 0x8f, 0x6c, 0xd0, 0xff, 0x50, 0x85, 0xe5, 0x97, 0x77, 0x81, 0x44, 0x36,
	0x6c, 0x6d, 0xd8, 0x46, 0x44, 0x3b, 0x97, 0xb3, 0x2d, 0xe3, 0x18, 0xf4, 0xc0, 0xa0, 0x39, 0x7c,
	0x46, 0x32, 0x63, 0x72, 0x30, 0x26, 0xf3, 0xcb, 0x50, 0x8d, 0x86, 0x37, 0x94, 0xdf, 0x32, 0xb3,
	0xf9, 0xe9, 0xfb, 0xe7, 0x52, 0xf5, 0xac, 0x8c, 0x1b, 0x03, 0xe0, 0x34, 0x33, 0xdc, 0x4e, 0x68,
	0x74, 0x0d, 0x3b, 0x9e, 0x08, 0x15, 0x0b, 0x95, 0xcb, 0x59, 0xe8, 0xab, 0x35, 0x2e, 0x3e, 0xc2,
	0xaf, 0x91, 0xf8, 0x95, 0xcf, 0x8a, 0x03, 0xdd, 0x69, 0xe7, 0xb8, 0x9d, 0xd2, 0x68, 0x0d, 0x4f,
	0x3c, 0xb1, 0x91, 0x11, 0x8b, 0xe3, 0x40, 0x84, 0xf9, 0x8a, 0xa5, 0x1e, 0x18, 0x74, 0x86, 0x06,
	0xf9, 0xdb, 0x55, 0x48, 0xf2, 0x31, 0x72, 0xf1, 0x8b, 0x4b, 0x77, 0x33, 0x61, 0x92, 0x52, 0x79,
	0x0f, 0x8a, 0x1a, 0x70, 0x34, 0xef, 0xc1, 0x14, 0xbd, 0x85, 0xff, 0xdf, 0x4f, 0x8a, 0x55, 0x44,
	0x15, 0xf3, 0x77, 0xb8, 0x9c, 0x86, 0x0d, 0x1f, 0x1f, 0x36, 0xcd, 0xc9, 0xdf, 0xf2, 0x4e, 0xbd,
	0x63, 0x01, 0x1a, 0xc1, 0xb6, 0x47, 0xbd, 0x35, 0x73, 0x93, 0x3f, 0x47, 0x82, 0xe3, 0x4a, 0xaf,
	0x34, 0x68, 0x98, 0x9d, 0x84, 0x6b, 0x7c, 0x04, 0xd5, 0x7a, 0x41, 0x83, 0x18, 0x38, 0xad, 0x54,
	0x74, 0x91, 0x69, 0xd0, 0x73, 0xd8, 0x3a, 0x5c, 0x58, 0xed, 0x24, 0xc3, 0xd5, 0x23, 0x46, 0xc7,
	0xc0, 0x69, 0xe6, 0x9a, 0xd9, 0x4e, 0x32, 0x64, 0xc2, 0xff, 0x96, 0x41, 0x4c, 0x17, 0x9c, 0xb9,
	0x22, 0x74, 0x99, 0xa2, 0x3e, 0xae, 0xa5, 0xad, 0x74, 0x8f, 0x5a, 0x31, 0x85, 0xe0, 0x69, 0x27,
	0x4e, 0x3b, 0x47, 0x5e, 0x87, 0x96, 0xa2, 0x3e, 0x9a, 0x42, 0x7c, 0xcf, 0x83, 0xd3, 0x58, 0xb9,
	0x1b, 0xb1, 0x0c, 0x56, 0x01, 0x5b, 0xe2, 0xfa, 0x3f, 0xcd, 0x9e, 0xfc, 0x34, 0xb3, 0x69, 0xac,
	0xae, 0x72, 0xb0, 0xff, 0x02, 0x6a, 0x0f, 0x5b, 0x42, 0x27, 0xb0, 0x3d, 0xb6, 0x26, 0xe7, 0x73,
	0x7b, 0xe6, 0xda, 0xd6, 0x8d, 0x65, 0x6b, 0x05, 0x54, 0x87, 0x65, 0xd3, 0x9a, 0xce, 0x34, 0x80,
	0x1a, 0xb0, 0x32, 0x7d, 0x63, 0x59, 0x63, 0xad, 0xd8, 0x9f, 0xc0, 0xd3, 0x3f, 0x1c, 0x1d, 0x35,
	0x61, 0x2d, 0xc7, 0xb5, 0x02, 0x6a, 0xc1, 0xfa, 0xe4, 0xd2, 0x9e, 0x59, 0x8e, 0x35, 0xd6, 0x40,
	0x32, 0x7a, 0x35, 0x9f, 0x4c, 0xae, 0xce, 0xaf, 0xb5, 0x22, 0xaa, 0xc1, 0x92, 0x63, 0x5b, 0x5a,
	0xc9, 0x6c, 0x7d, 0xd9, 0xeb, 0xe0, 0xeb, 0x5e, 0x07, 0xdf, 0xf6, 0x3a, 0x58, 0x54, 0xd3, 0xe5,
	0x47, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x0c, 0x58, 0xcb, 0x8e, 0x03, 0x00, 0x00,
}
