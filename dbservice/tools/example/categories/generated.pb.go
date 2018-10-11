// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/kennyzhu/go-os/dbservice/tools/example/protos/github.com/kennyzhu/go-os/dbservice/tools/example/categories/generated.proto

package categories

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

func (m *CategoryOptions) Reset()         { *m = CategoryOptions{} }
func (m *CategoryOptions) String() string { return proto.CompactTextString(m) }
func (*CategoryOptions) ProtoMessage()    {}
func (*CategoryOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_b6826ce6c3ad3c1d, []int{0}
}
func (m *CategoryOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CategoryOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CategoryOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CategoryOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryOptions.Merge(dst, src)
}
func (m *CategoryOptions) XXX_Size() int {
	return m.ProtoSize()
}
func (m *CategoryOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryOptions proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CategoryOptions)(nil), "github.com.kennyzhu.goos.dbservice.tools.example.categories.CategoryOptions")
}
func (m *CategoryOptions) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CategoryOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ShowPrices {
		dAtA[i] = 0x8
		i++
		if m.ShowPrices {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CanBuy {
		dAtA[i] = 0x10
		i++
		if m.CanBuy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CategoryOptions) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ShowPrices {
		n += 2
	}
	if m.CanBuy {
		n += 2
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CategoryOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: CategoryOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CategoryOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShowPrices", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
			m.ShowPrices = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanBuy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
			m.CanBuy = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/kennyzhu/go-os/dbservice/tools/example/protos/github.com/kennyzhu/go-os/dbservice/tools/example/categories/generated.proto", fileDescriptor_generated_b6826ce6c3ad3c1d)
}

var fileDescriptor_generated_b6826ce6c3ad3c1d = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8e, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x6d, 0x86, 0x52, 0x99, 0x01, 0xa9, 0x0b, 0x15, 0x83, 0x83, 0x98, 0xba, 0xd4, 0x1e,
	0xd8, 0x60, 0x2b, 0x0f, 0x00, 0x42, 0x62, 0x61, 0xa9, 0x12, 0xf7, 0xe2, 0x58, 0xb4, 0xbe, 0x91,
	0x7f, 0x80, 0xb0, 0x23, 0x31, 0xf2, 0x08, 0xe5, 0x6d, 0x18, 0x3b, 0x32, 0xa2, 0xe4, 0x05, 0x78,
	0x04, 0x84, 0x03, 0x84, 0x95, 0xcd, 0xe7, 0x58, 0xe7, 0xbb, 0x1f, 0x7b, 0xa4, 0xda, 0x84, 0x32,
	0x16, 0x42, 0xe1, 0x4a, 0xde, 0x80, 0xb5, 0xf5, 0x43, 0x19, 0xa5, 0xc6, 0x29, 0x7a, 0xb9, 0x28,
	0x3c, 0xb8, 0x5b, 0xa3, 0x40, 0x06, 0xc4, 0xa5, 0x97, 0x70, 0x9f, 0xaf, 0xaa, 0x25, 0xc8, 0xca,
	0x61, 0x40, 0x2f, 0xff, 0x3f, 0x54, 0x79, 0x00, 0x8d, 0xce, 0x80, 0x97, 0x1a, 0x2c, 0xb8, 0x3c,
	0xc0, 0x42, 0x24, 0xda, 0xe8, 0xa4, 0xa7, 0x89, 0x1f, 0x9a, 0xd0, 0x88, 0x5e, 0xfc, 0xc2, 0x44,
	0x82, 0x89, 0x6f, 0x98, 0xe8, 0x61, 0xfb, 0xd3, 0x3f, 0x2a, 0x1a, 0x35, 0x76, 0x86, 0x45, 0xbc,
	0x4e, 0x29, 0x85, 0xf4, 0xea, 0x6e, 0x1d, 0x5e, 0xb2, 0xdd, 0xd3, 0x6e, 0x5c, 0x9f, 0x55, 0xc1,
	0xa0, 0xf5, 0xa3, 0x8c, 0xed, 0xf8, 0x12, 0xef, 0xe6, 0x95, 0x33, 0x0a, 0xfc, 0x98, 0x1e, 0xd0,
	0xc9, 0xf0, 0x82, 0x7d, 0x55, 0xe7, 0xa9, 0x19, 0xed, 0xb1, 0x6d, 0x95, 0xdb, 0x79, 0x11, 0xeb,
	0xf1, 0x56, 0xfa, 0x1c, 0xa8, 0xdc, 0xce, 0x62, 0x7d, 0x3c, 0x7c, 0x5a, 0x67, 0xe4, 0xe3, 0x25,
	0x23, 0xb3, 0xc9, 0x6b, 0xc3, 0xe9, 0xa6, 0xe1, 0xf4, 0xbd, 0xe1, 0xe4, 0xb9, 0xe5, 0x64, 0xdd,
	0x72, 0xba, 0x69, 0x39, 0x79, 0x6b, 0x39, 0xb9, 0x62, 0xbd, 0x6f, 0x31, 0x48, 0x1e, 0x47, 0x9f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xb1, 0x90, 0x85, 0x7d, 0x01, 0x00, 0x00,
}
