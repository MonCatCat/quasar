// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qoracle/pool_metrics.proto

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

type PoolMetrics struct {
	HighestAPY string      `protobuf:"bytes,1,opt,name=highestAPY,proto3" json:"highestAPY,omitempty"`
	TVL        string      `protobuf:"bytes,2,opt,name=tVL,proto3" json:"tVL,omitempty"`
	GaugeAPYs  []*GaugeAPY `protobuf:"bytes,3,rep,name=gaugeAPYs,proto3" json:"gaugeAPYs,omitempty"`
}

func (m *PoolMetrics) Reset()         { *m = PoolMetrics{} }
func (m *PoolMetrics) String() string { return proto.CompactTextString(m) }
func (*PoolMetrics) ProtoMessage()    {}
func (*PoolMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fad8abc4c7c26f8, []int{0}
}
func (m *PoolMetrics) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolMetrics.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolMetrics.Merge(m, src)
}
func (m *PoolMetrics) XXX_Size() int {
	return m.Size()
}
func (m *PoolMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_PoolMetrics proto.InternalMessageInfo

func (m *PoolMetrics) GetHighestAPY() string {
	if m != nil {
		return m.HighestAPY
	}
	return ""
}

func (m *PoolMetrics) GetTVL() string {
	if m != nil {
		return m.TVL
	}
	return ""
}

func (m *PoolMetrics) GetGaugeAPYs() []*GaugeAPY {
	if m != nil {
		return m.GaugeAPYs
	}
	return nil
}

func init() {
	proto.RegisterType((*PoolMetrics)(nil), "quasarlabs.quasarnode.qoracle.PoolMetrics")
}

func init() { proto.RegisterFile("qoracle/pool_metrics.proto", fileDescriptor_6fad8abc4c7c26f8) }

var fileDescriptor_6fad8abc4c7c26f8 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xcc, 0x2f, 0x4a,
	0x4c, 0xce, 0x49, 0xd5, 0x2f, 0xc8, 0xcf, 0xcf, 0x89, 0xcf, 0x4d, 0x2d, 0x29, 0xca, 0x4c, 0x2e,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2d, 0x2c, 0x4d, 0x2c, 0x4e, 0x2c, 0xca, 0x49,
	0x4c, 0x2a, 0xd6, 0x83, 0x30, 0xf3, 0xf2, 0x53, 0x52, 0xf5, 0xa0, 0x3a, 0xa4, 0xc4, 0x61, 0x5a,
	0xd3, 0x13, 0x4b, 0xd3, 0x53, 0xe3, 0x13, 0x0b, 0x2a, 0x21, 0xfa, 0x94, 0xda, 0x18, 0xb9, 0xb8,
	0x03, 0xf2, 0xf3, 0x73, 0x7c, 0x21, 0xa6, 0x09, 0xc9, 0x71, 0x71, 0x65, 0x64, 0xa6, 0x67, 0xa4,
	0x16, 0x97, 0x38, 0x06, 0x44, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x89, 0x08, 0x09,
	0x70, 0x31, 0x97, 0x84, 0xf9, 0x48, 0x30, 0x81, 0x25, 0x40, 0x4c, 0x21, 0x57, 0x2e, 0x4e, 0xb0,
	0xa1, 0x8e, 0x01, 0x91, 0xc5, 0x12, 0xcc, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xea, 0x7a, 0x78, 0x5d,
	0xa3, 0xe7, 0x0e, 0x55, 0x1f, 0x84, 0xd0, 0xe9, 0xe4, 0x75, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0x06, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x08, 0x73, 0xf5, 0x11, 0xe6, 0xea, 0x57, 0xe8, 0xc3, 0xbc, 0x57, 0x52, 0x59, 0x90, 0x5a, 0x9c,
	0xc4, 0x06, 0xf6, 0x9b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x26, 0x8c, 0xb7, 0x98, 0x31, 0x01,
	0x00, 0x00,
}

func (m *PoolMetrics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolMetrics) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolMetrics) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GaugeAPYs) > 0 {
		for iNdEx := len(m.GaugeAPYs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GaugeAPYs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPoolMetrics(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TVL) > 0 {
		i -= len(m.TVL)
		copy(dAtA[i:], m.TVL)
		i = encodeVarintPoolMetrics(dAtA, i, uint64(len(m.TVL)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.HighestAPY) > 0 {
		i -= len(m.HighestAPY)
		copy(dAtA[i:], m.HighestAPY)
		i = encodeVarintPoolMetrics(dAtA, i, uint64(len(m.HighestAPY)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoolMetrics(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolMetrics(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolMetrics) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HighestAPY)
	if l > 0 {
		n += 1 + l + sovPoolMetrics(uint64(l))
	}
	l = len(m.TVL)
	if l > 0 {
		n += 1 + l + sovPoolMetrics(uint64(l))
	}
	if len(m.GaugeAPYs) > 0 {
		for _, e := range m.GaugeAPYs {
			l = e.Size()
			n += 1 + l + sovPoolMetrics(uint64(l))
		}
	}
	return n
}

func sovPoolMetrics(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolMetrics(x uint64) (n int) {
	return sovPoolMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolMetrics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolMetrics
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
			return fmt.Errorf("proto: PoolMetrics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolMetrics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HighestAPY", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolMetrics
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
				return ErrInvalidLengthPoolMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HighestAPY = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TVL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolMetrics
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
				return ErrInvalidLengthPoolMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TVL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeAPYs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolMetrics
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
				return ErrInvalidLengthPoolMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPoolMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GaugeAPYs = append(m.GaugeAPYs, &GaugeAPY{})
			if err := m.GaugeAPYs[len(m.GaugeAPYs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoolMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolMetrics
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
func skipPoolMetrics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolMetrics
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
					return 0, ErrIntOverflowPoolMetrics
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
					return 0, ErrIntOverflowPoolMetrics
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
				return 0, ErrInvalidLengthPoolMetrics
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolMetrics
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolMetrics
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolMetrics        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolMetrics          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolMetrics = fmt.Errorf("proto: unexpected end of group")
)
