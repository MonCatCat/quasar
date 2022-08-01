// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qoracle/gauge_apy.proto

package types

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

type GaugeAPY struct {
	GaugeId  uint64 `protobuf:"varint,1,opt,name=gaugeId,proto3" json:"gaugeId,omitempty"`
	Duration string `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	APY      string `protobuf:"bytes,3,opt,name=aPY,proto3" json:"aPY,omitempty"`
}

func (m *GaugeAPY) Reset()         { *m = GaugeAPY{} }
func (m *GaugeAPY) String() string { return proto.CompactTextString(m) }
func (*GaugeAPY) ProtoMessage()    {}
func (*GaugeAPY) Descriptor() ([]byte, []int) {
	return fileDescriptor_c57539e311091ba3, []int{0}
}
func (m *GaugeAPY) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GaugeAPY) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GaugeAPY.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GaugeAPY) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaugeAPY.Merge(m, src)
}
func (m *GaugeAPY) XXX_Size() int {
	return m.Size()
}
func (m *GaugeAPY) XXX_DiscardUnknown() {
	xxx_messageInfo_GaugeAPY.DiscardUnknown(m)
}

var xxx_messageInfo_GaugeAPY proto.InternalMessageInfo

func (m *GaugeAPY) GetGaugeId() uint64 {
	if m != nil {
		return m.GaugeId
	}
	return 0
}

func (m *GaugeAPY) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *GaugeAPY) GetAPY() string {
	if m != nil {
		return m.APY
	}
	return ""
}

func init() {
	proto.RegisterType((*GaugeAPY)(nil), "quasarlabs.quasarnode.qoracle.GaugeAPY")
}

func init() { proto.RegisterFile("qoracle/gauge_apy.proto", fileDescriptor_c57539e311091ba3) }

var fileDescriptor_c57539e311091ba3 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xcc, 0x2f, 0x4a,
	0x4c, 0xce, 0x49, 0xd5, 0x4f, 0x4f, 0x2c, 0x4d, 0x4f, 0x8d, 0x4f, 0x2c, 0xa8, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2d, 0x2c, 0x4d, 0x2c, 0x4e, 0x2c, 0xca, 0x49, 0x4c, 0x2a, 0xd6,
	0x83, 0x30, 0xf3, 0xf2, 0x53, 0x52, 0xf5, 0xa0, 0xca, 0x95, 0x82, 0xb8, 0x38, 0xdc, 0x41, 0x3a,
	0x1c, 0x03, 0x22, 0x85, 0x24, 0xb8, 0xd8, 0xc1, 0xba, 0x3d, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x58, 0x82, 0x60, 0x5c, 0x21, 0x29, 0x2e, 0x8e, 0x94, 0xd2, 0xa2, 0xc4, 0x92, 0xcc, 0xfc, 0x3c,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x5f, 0x48, 0x80, 0x8b, 0x39, 0x31, 0x20, 0x52,
	0x82, 0x19, 0x2c, 0x0c, 0x62, 0x3a, 0x79, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x43, 0x94, 0x41, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc2, 0x5d,
	0xfa, 0x08, 0x77, 0xe9, 0x57, 0xe8, 0xc3, 0x3c, 0x52, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06,
	0xf6, 0x85, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x43, 0xfd, 0x64, 0x86, 0xe0, 0x00, 0x00, 0x00,
}

func (m *GaugeAPY) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GaugeAPY) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GaugeAPY) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.APY) > 0 {
		i -= len(m.APY)
		copy(dAtA[i:], m.APY)
		i = encodeVarintGaugeApy(dAtA, i, uint64(len(m.APY)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Duration) > 0 {
		i -= len(m.Duration)
		copy(dAtA[i:], m.Duration)
		i = encodeVarintGaugeApy(dAtA, i, uint64(len(m.Duration)))
		i--
		dAtA[i] = 0x12
	}
	if m.GaugeId != 0 {
		i = encodeVarintGaugeApy(dAtA, i, uint64(m.GaugeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGaugeApy(dAtA []byte, offset int, v uint64) int {
	offset -= sovGaugeApy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GaugeAPY) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GaugeId != 0 {
		n += 1 + sovGaugeApy(uint64(m.GaugeId))
	}
	l = len(m.Duration)
	if l > 0 {
		n += 1 + l + sovGaugeApy(uint64(l))
	}
	l = len(m.APY)
	if l > 0 {
		n += 1 + l + sovGaugeApy(uint64(l))
	}
	return n
}

func sovGaugeApy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGaugeApy(x uint64) (n int) {
	return sovGaugeApy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GaugeAPY) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGaugeApy
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
			return fmt.Errorf("proto: GaugeAPY: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GaugeAPY: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeId", wireType)
			}
			m.GaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaugeApy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaugeApy
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
				return ErrInvalidLengthGaugeApy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGaugeApy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Duration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APY", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaugeApy
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
				return ErrInvalidLengthGaugeApy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGaugeApy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.APY = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGaugeApy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGaugeApy
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
func skipGaugeApy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGaugeApy
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
					return 0, ErrIntOverflowGaugeApy
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
					return 0, ErrIntOverflowGaugeApy
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
				return 0, ErrInvalidLengthGaugeApy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGaugeApy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGaugeApy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGaugeApy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGaugeApy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGaugeApy = fmt.Errorf("proto: unexpected end of group")
)
