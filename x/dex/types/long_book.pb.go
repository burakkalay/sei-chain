// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/long_book.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LongBook struct {
	Price github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"price"`
	Entry *OrderEntry                            `protobuf:"bytes,2,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (m *LongBook) Reset()         { *m = LongBook{} }
func (m *LongBook) String() string { return proto.CompactTextString(m) }
func (*LongBook) ProtoMessage()    {}
func (*LongBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_daa9d9e2359557c9, []int{0}
}
func (m *LongBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LongBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LongBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LongBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongBook.Merge(m, src)
}
func (m *LongBook) XXX_Size() int {
	return m.Size()
}
func (m *LongBook) XXX_DiscardUnknown() {
	xxx_messageInfo_LongBook.DiscardUnknown(m)
}

var xxx_messageInfo_LongBook proto.InternalMessageInfo

func (m *LongBook) GetEntry() *OrderEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func init() {
	proto.RegisterType((*LongBook)(nil), "seiprotocol.seichain.dex.LongBook")
}

func init() { proto.RegisterFile("dex/long_book.proto", fileDescriptor_daa9d9e2359557c9) }

var fileDescriptor_daa9d9e2359557c9 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x49, 0xad, 0xd0,
	0xcf, 0xc9, 0xcf, 0x4b, 0x8f, 0x4f, 0xca, 0xcf, 0xcf, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x28, 0x4e, 0xcd, 0x04, 0xb3, 0x92, 0xf3, 0x73, 0xf4, 0x8a, 0x53, 0x33, 0x93, 0x33, 0x12,
	0x33, 0xf3, 0xf4, 0x52, 0x52, 0x2b, 0xa4, 0x44, 0x41, 0xca, 0xf3, 0x8b, 0x52, 0x52, 0x8b, 0xe2,
	0x53, 0xf3, 0x4a, 0x8a, 0x2a, 0x21, 0x1a, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4c, 0x7d,
	0x10, 0x0b, 0x22, 0xaa, 0x34, 0x87, 0x91, 0x8b, 0xc3, 0x27, 0x3f, 0x2f, 0xdd, 0x29, 0x3f, 0x3f,
	0x5b, 0x28, 0x84, 0x8b, 0xb5, 0xa0, 0x28, 0x33, 0x39, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3,
	0xc9, 0xee, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0xd5, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93,
	0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xa1, 0x94, 0x6e, 0x71, 0x4a,
	0xb6, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0xb1, 0x9e, 0x4b, 0x6a, 0xf2, 0xa7, 0x7b, 0xf2, 0x3c, 0x95,
	0x89, 0xb9, 0x39, 0x56, 0x4a, 0x60, 0x43, 0x94, 0x82, 0x20, 0x86, 0x09, 0x59, 0x71, 0xb1, 0x82,
	0xdd, 0x21, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa2, 0x87, 0xcb, 0xe5, 0x7a, 0xfe, 0x20,
	0x47, 0xbb, 0x82, 0xd4, 0x06, 0x41, 0xb4, 0x38, 0xb9, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x94, 0x2e, 0x92, 0xa3, 0x8a, 0x53, 0x33, 0x75, 0x61, 0x26, 0x82, 0x39, 0x60,
	0x23, 0xf5, 0x2b, 0xf4, 0x41, 0x21, 0x01, 0x76, 0x5f, 0x12, 0x1b, 0x58, 0xde, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x9f, 0xaf, 0x45, 0x06, 0x4c, 0x01, 0x00, 0x00,
}

func (m *LongBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LongBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LongBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Entry != nil {
		{
			size, err := m.Entry.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLongBook(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLongBook(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintLongBook(dAtA []byte, offset int, v uint64) int {
	offset -= sovLongBook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LongBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Price.Size()
	n += 1 + l + sovLongBook(uint64(l))
	if m.Entry != nil {
		l = m.Entry.Size()
		n += 1 + l + sovLongBook(uint64(l))
	}
	return n
}

func sovLongBook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLongBook(x uint64) (n int) {
	return sovLongBook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LongBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLongBook
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
			return fmt.Errorf("proto: LongBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LongBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLongBook
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
				return ErrInvalidLengthLongBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLongBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLongBook
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
				return ErrInvalidLengthLongBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLongBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entry == nil {
				m.Entry = &OrderEntry{}
			}
			if err := m.Entry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLongBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLongBook
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
func skipLongBook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLongBook
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
					return 0, ErrIntOverflowLongBook
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLongBook
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
				return 0, ErrInvalidLengthLongBook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLongBook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLongBook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLongBook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLongBook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLongBook = fmt.Errorf("proto: unexpected end of group")
)
