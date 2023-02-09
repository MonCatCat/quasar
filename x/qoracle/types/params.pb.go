// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qoracle/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	DenomPricesExpDuration uint64 `protobuf:"varint,2,opt,name=denom_prices_exp_duration,json=denomPricesExpDuration,proto3" json:"denom_prices_exp_duration,omitempty" yaml:"denom_prices_exp_duration"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_95f4338228c132f2, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDenomPricesExpDuration() uint64 {
	if m != nil {
		return m.DenomPricesExpDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "quasarlabs.quasarnode.qoracle.Params")
}

func init() { proto.RegisterFile("qoracle/params.proto", fileDescriptor_95f4338228c132f2) }

var fileDescriptor_95f4338228c132f2 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xcc, 0x2f, 0x4a,
	0x4c, 0xce, 0x49, 0xd5, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x2d, 0x2c, 0x4d, 0x2c, 0x4e, 0x2c, 0xca, 0x49, 0x4c, 0x2a, 0xd6, 0x83, 0x30, 0xf3,
	0xf2, 0x53, 0x52, 0xf5, 0xa0, 0x6a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x2a, 0xf5, 0x41,
	0x2c, 0x88, 0x26, 0x29, 0x25, 0x98, 0x51, 0x29, 0xa9, 0x79, 0xf9, 0xb9, 0xf1, 0xc5, 0x95, 0xb9,
	0x49, 0xf9, 0x39, 0xf1, 0xb9, 0x89, 0x05, 0x05, 0x99, 0x79, 0xe9, 0x10, 0x35, 0x4a, 0xf9, 0x5c,
	0x6c, 0x01, 0x60, 0x8b, 0x84, 0xe2, 0xb9, 0x24, 0x21, 0xea, 0x0a, 0x8a, 0x32, 0x93, 0x53, 0x8b,
	0xe3, 0x53, 0x2b, 0x0a, 0xe2, 0x53, 0x4a, 0x8b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x24, 0x98, 0x14,
	0x18, 0x35, 0x58, 0x9c, 0x54, 0x3e, 0xdd, 0x93, 0x57, 0xa8, 0x4c, 0xcc, 0xcd, 0xb1, 0x52, 0xc2,
	0xa9, 0x54, 0x29, 0x48, 0x0c, 0x2c, 0x17, 0x00, 0x96, 0x72, 0xad, 0x28, 0x70, 0x81, 0x4a, 0x58,
	0xb1, 0xcc, 0x58, 0x20, 0xcf, 0xe0, 0xe4, 0x75, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x51, 0x06, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x08, 0xef,
	0xea, 0x23, 0xbc, 0xab, 0x5f, 0xa1, 0x0f, 0xf3, 0x51, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b,
	0xd8, 0x0f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x48, 0x49, 0x5b, 0x34, 0x01, 0x00,
	0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DenomPricesExpDuration != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DenomPricesExpDuration))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DenomPricesExpDuration != 0 {
		n += 1 + sovParams(uint64(m.DenomPricesExpDuration))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomPricesExpDuration", wireType)
			}
			m.DenomPricesExpDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DenomPricesExpDuration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
