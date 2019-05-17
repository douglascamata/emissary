// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/lua/v2/lua.proto

package v2

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
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

type Lua struct {
	// The Lua code that Envoy will execute. This can be a very small script that
	// further loads code from disk if desired. Note that if JSON configuration is used, the code must
	// be properly escaped. YAML configuration may be easier to read since YAML supports multi-line
	// strings so complex scripts can be easily expressed inline in the configuration.
	InlineCode           string   `protobuf:"bytes,1,opt,name=inline_code,json=inlineCode,proto3" json:"inline_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lua) Reset()         { *m = Lua{} }
func (m *Lua) String() string { return proto.CompactTextString(m) }
func (*Lua) ProtoMessage()    {}
func (*Lua) Descriptor() ([]byte, []int) {
	return fileDescriptor_f59dca3e63e33613, []int{0}
}
func (m *Lua) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lua) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lua.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lua) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lua.Merge(m, src)
}
func (m *Lua) XXX_Size() int {
	return m.Size()
}
func (m *Lua) XXX_DiscardUnknown() {
	xxx_messageInfo_Lua.DiscardUnknown(m)
}

var xxx_messageInfo_Lua proto.InternalMessageInfo

func (m *Lua) GetInlineCode() string {
	if m != nil {
		return m.InlineCode
	}
	return ""
}

func init() {
	proto.RegisterType((*Lua)(nil), "envoy.config.filter.http.lua.v2.Lua")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/lua/v2/lua.proto", fileDescriptor_f59dca3e63e33613)
}

var fileDescriptor_f59dca3e63e33613 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0xcf, 0x29, 0x4d, 0xd4, 0x2f, 0x33, 0x02, 0x51, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0xf2, 0x60, 0xa5, 0x7a, 0x10, 0xa5, 0x7a, 0x10, 0xa5, 0x7a, 0x20, 0xa5,
	0x7a, 0x20, 0x35, 0x65, 0x46, 0x52, 0xe2, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa,
	0x30, 0x06, 0x44, 0xa7, 0x92, 0x21, 0x17, 0xb3, 0x4f, 0x69, 0xa2, 0x90, 0x16, 0x17, 0x77, 0x66,
	0x5e, 0x4e, 0x66, 0x5e, 0x6a, 0x7c, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7,
	0x13, 0xe7, 0xae, 0x97, 0x07, 0x98, 0x59, 0x8a, 0x98, 0x14, 0x18, 0x83, 0xb8, 0x20, 0xb2, 0xce,
	0xf9, 0x29, 0xa9, 0x4e, 0xde, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0x23, 0x97, 0x6e, 0x66, 0xbe, 0x1e, 0xd8, 0xf6, 0x82, 0xa2, 0xfc, 0x8a, 0x4a, 0x3d, 0x02,
	0x0e, 0x71, 0xe2, 0xf0, 0x29, 0x4d, 0x0c, 0x00, 0xd9, 0x1c, 0xc0, 0x18, 0xc5, 0x54, 0x66, 0x94,
	0xc4, 0x06, 0x76, 0x86, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x11, 0x53, 0xce, 0xed, 0x00,
	0x00, 0x00,
}

func (m *Lua) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lua) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.InlineCode) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLua(dAtA, i, uint64(len(m.InlineCode)))
		i += copy(dAtA[i:], m.InlineCode)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintLua(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Lua) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InlineCode)
	if l > 0 {
		n += 1 + l + sovLua(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLua(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLua(x uint64) (n int) {
	return sovLua(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lua) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLua
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
			return fmt.Errorf("proto: Lua: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lua: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InlineCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLua
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLua
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLua
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InlineCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLua(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLua
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLua
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
func skipLua(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLua
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
					return 0, ErrIntOverflowLua
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
					return 0, ErrIntOverflowLua
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
				return 0, ErrInvalidLengthLua
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLua
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLua
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
				next, err := skipLua(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLua
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
	ErrInvalidLengthLua = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLua   = fmt.Errorf("proto: integer overflow")
)
