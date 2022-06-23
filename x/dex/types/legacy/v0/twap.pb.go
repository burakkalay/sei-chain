// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: legacy/dex/v0/twap.proto

package v0

import (
	fmt "fmt"
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

type Twap struct {
	LastEpoch  uint64   `protobuf:"varint,1,opt,name=lastEpoch,proto3" json:"lastEpoch,omitempty"`
	Prices     []uint64 `protobuf:"varint,2,rep,packed,name=prices,proto3" json:"prices,omitempty"`
	TwapPrice  uint64   `protobuf:"varint,3,opt,name=twapPrice,proto3" json:"twapPrice,omitempty"`
	PriceDenom string   `protobuf:"bytes,4,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
	AssetDenom string   `protobuf:"bytes,5,opt,name=assetDenom,proto3" json:"assetDenom,omitempty"`
}

func (m *Twap) Reset()         { *m = Twap{} }
func (m *Twap) String() string { return proto.CompactTextString(m) }
func (*Twap) ProtoMessage()    {}
func (*Twap) Descriptor() ([]byte, []int) {
	return fileDescriptor_f13fe59824d2008e, []int{0}
}
func (m *Twap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Twap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Twap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Twap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Twap.Merge(m, src)
}
func (m *Twap) XXX_Size() int {
	return m.Size()
}
func (m *Twap) XXX_DiscardUnknown() {
	xxx_messageInfo_Twap.DiscardUnknown(m)
}

var xxx_messageInfo_Twap proto.InternalMessageInfo

func (m *Twap) GetLastEpoch() uint64 {
	if m != nil {
		return m.LastEpoch
	}
	return 0
}

func (m *Twap) GetPrices() []uint64 {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *Twap) GetTwapPrice() uint64 {
	if m != nil {
		return m.TwapPrice
	}
	return 0
}

func (m *Twap) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *Twap) GetAssetDenom() string {
	if m != nil {
		return m.AssetDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*Twap)(nil), "seiprotocol.seichain.legacy.dex.v0.Twap")
}

func init() { proto.RegisterFile("legacy/dex/v0/twap.proto", fileDescriptor_f13fe59824d2008e) }

var fileDescriptor_f13fe59824d2008e = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0x63, 0x1a, 0x2a, 0xd5, 0x63, 0x06, 0xe4, 0x01, 0x59, 0x51, 0xa7, 0x2c, 0xd8, 0x91,
	0x18, 0xd8, 0x11, 0xec, 0x10, 0x31, 0xb1, 0xb9, 0xee, 0xaf, 0xc6, 0x52, 0x5a, 0x5b, 0xfd, 0x4d,
	0x9a, 0xbe, 0x05, 0x33, 0x4f, 0xc4, 0xd8, 0x91, 0x11, 0x25, 0x2f, 0x82, 0x62, 0x17, 0xc2, 0xe6,
	0xfb, 0xee, 0x7e, 0x59, 0x77, 0x94, 0x35, 0xb0, 0x51, 0xfa, 0x28, 0xd7, 0xd0, 0xc9, 0xb6, 0x94,
	0xfe, 0xa0, 0x9c, 0x70, 0x7b, 0xeb, 0x6d, 0xb6, 0x44, 0x30, 0xe1, 0xa5, 0x6d, 0x23, 0x10, 0x8c,
	0xae, 0x95, 0xd9, 0x89, 0x18, 0x17, 0x6b, 0xe8, 0x44, 0x5b, 0x2e, 0x3f, 0x08, 0x4d, 0x5f, 0x0e,
	0xca, 0x65, 0xd7, 0x74, 0xd1, 0x28, 0xf4, 0x8f, 0xce, 0xea, 0x9a, 0x91, 0x9c, 0x14, 0x69, 0x35,
	0x81, 0xec, 0x8a, 0xce, 0xdd, 0xde, 0x68, 0x40, 0x76, 0x91, 0xcf, 0x8a, 0xb4, 0x3a, 0xab, 0xf1,
	0x6a, 0xfc, 0xf0, 0x69, 0x54, 0x6c, 0x16, 0xaf, 0xfe, 0x40, 0xc6, 0x29, 0x0d, 0xb9, 0x07, 0xd8,
	0xd9, 0x2d, 0x4b, 0x73, 0x52, 0x2c, 0xaa, 0x7f, 0x64, 0xf4, 0x15, 0x22, 0xf8, 0xe8, 0x5f, 0x46,
	0x7f, 0x22, 0xf7, 0xcf, 0x9f, 0x3d, 0x27, 0xa7, 0x9e, 0x93, 0xef, 0x9e, 0x93, 0xf7, 0x81, 0x27,
	0xa7, 0x81, 0x27, 0x5f, 0x03, 0x4f, 0x5e, 0xef, 0x36, 0xc6, 0xd7, 0x6f, 0x2b, 0xa1, 0xed, 0x56,
	0x22, 0x98, 0x9b, 0xdf, 0x9a, 0x41, 0x84, 0x9e, 0xb2, 0x0b, 0x8b, 0xf8, 0xa3, 0x03, 0x94, 0xe7,
	0x89, 0xda, 0x72, 0x35, 0x0f, 0xc9, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0x3f, 0x81,
	0x46, 0x36, 0x01, 0x00, 0x00,
}

func (m *Twap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Twap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Twap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetDenom) > 0 {
		i -= len(m.AssetDenom)
		copy(dAtA[i:], m.AssetDenom)
		i = encodeVarintTwap(dAtA, i, uint64(len(m.AssetDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintTwap(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x22
	}
	if m.TwapPrice != 0 {
		i = encodeVarintTwap(dAtA, i, uint64(m.TwapPrice))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Prices) > 0 {
		dAtA2 := make([]byte, len(m.Prices)*10)
		var j1 int
		for _, num := range m.Prices {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTwap(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if m.LastEpoch != 0 {
		i = encodeVarintTwap(dAtA, i, uint64(m.LastEpoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTwap(dAtA []byte, offset int, v uint64) int {
	offset -= sovTwap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Twap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LastEpoch != 0 {
		n += 1 + sovTwap(uint64(m.LastEpoch))
	}
	if len(m.Prices) > 0 {
		l = 0
		for _, e := range m.Prices {
			l += sovTwap(uint64(e))
		}
		n += 1 + sovTwap(uint64(l)) + l
	}
	if m.TwapPrice != 0 {
		n += 1 + sovTwap(uint64(m.TwapPrice))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovTwap(uint64(l))
	}
	l = len(m.AssetDenom)
	if l > 0 {
		n += 1 + l + sovTwap(uint64(l))
	}
	return n
}

func sovTwap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTwap(x uint64) (n int) {
	return sovTwap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Twap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTwap
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
			return fmt.Errorf("proto: Twap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Twap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastEpoch", wireType)
			}
			m.LastEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTwap
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Prices = append(m.Prices, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTwap
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTwap
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTwap
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Prices) == 0 {
					m.Prices = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTwap
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Prices = append(m.Prices, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TwapPrice", wireType)
			}
			m.TwapPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TwapPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwap
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
				return ErrInvalidLengthTwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwap
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
				return ErrInvalidLengthTwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTwap
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
func skipTwap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTwap
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
					return 0, ErrIntOverflowTwap
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
					return 0, ErrIntOverflowTwap
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
				return 0, ErrInvalidLengthTwap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTwap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTwap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTwap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTwap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTwap = fmt.Errorf("proto: unexpected end of group")
)
