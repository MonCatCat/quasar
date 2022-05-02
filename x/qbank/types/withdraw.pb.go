// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qbank/withdraw.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Withdraw defines the withdraw object to be stored in the KV store.
type Withdraw struct {
	Id                  uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RiskProfile         string     `protobuf:"bytes,2,opt,name=riskProfile,proto3" json:"riskProfile,omitempty"`
	VaultID             string     `protobuf:"bytes,3,opt,name=vaultID,proto3" json:"vaultID,omitempty"`
	DepositorAccAddress string     `protobuf:"bytes,4,opt,name=depositorAccAddress,proto3" json:"depositorAccAddress,omitempty"`
	Coin                types.Coin `protobuf:"bytes,5,opt,name=coin,proto3" json:"coin"`
}

func (m *Withdraw) Reset()         { *m = Withdraw{} }
func (m *Withdraw) String() string { return proto.CompactTextString(m) }
func (*Withdraw) ProtoMessage()    {}
func (*Withdraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_77bcdb9426b09e8b, []int{0}
}
func (m *Withdraw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Withdraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Withdraw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Withdraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Withdraw.Merge(m, src)
}
func (m *Withdraw) XXX_Size() int {
	return m.Size()
}
func (m *Withdraw) XXX_DiscardUnknown() {
	xxx_messageInfo_Withdraw.DiscardUnknown(m)
}

var xxx_messageInfo_Withdraw proto.InternalMessageInfo

func (m *Withdraw) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Withdraw) GetRiskProfile() string {
	if m != nil {
		return m.RiskProfile
	}
	return ""
}

func (m *Withdraw) GetVaultID() string {
	if m != nil {
		return m.VaultID
	}
	return ""
}

func (m *Withdraw) GetDepositorAccAddress() string {
	if m != nil {
		return m.DepositorAccAddress
	}
	return ""
}

func (m *Withdraw) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Withdraw)(nil), "abag.quasarnode.qbank.Withdraw")
}

func init() { proto.RegisterFile("qbank/withdraw.proto", fileDescriptor_77bcdb9426b09e8b) }

var fileDescriptor_77bcdb9426b09e8b = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0xfb, 0x30,
	0x1c, 0xc4, 0xe3, 0xfe, 0xf2, 0xe3, 0x8f, 0x2b, 0x31, 0x84, 0x22, 0x99, 0x0e, 0x26, 0x62, 0xca,
	0x64, 0x53, 0xfa, 0x04, 0x2d, 0x2c, 0x6c, 0xa8, 0x0b, 0x12, 0x9b, 0x1d, 0x9b, 0xd4, 0x6a, 0x9b,
	0x6f, 0x6b, 0xbb, 0x2d, 0xbc, 0x05, 0x2f, 0xc4, 0xde, 0xb1, 0x23, 0x13, 0x42, 0xcd, 0x8b, 0xa0,
	0x38, 0x45, 0x30, 0xb0, 0x9d, 0xef, 0xce, 0x96, 0xef, 0x83, 0x3b, 0x0b, 0x29, 0xca, 0x09, 0x5f,
	0x1b, 0x3f, 0x56, 0x56, 0xac, 0xd9, 0xdc, 0x82, 0x87, 0xe4, 0x4c, 0x48, 0x51, 0xb0, 0xc5, 0x52,
	0x38, 0x61, 0x4b, 0x50, 0x9a, 0x85, 0x56, 0x97, 0xe6, 0xe0, 0x66, 0xe0, 0xb8, 0x14, 0x4e, 0xf3,
	0x55, 0x4f, 0x6a, 0x2f, 0x7a, 0x3c, 0x07, 0x53, 0x36, 0xd7, 0xba, 0x9d, 0x02, 0x0a, 0x08, 0x92,
	0xd7, 0xaa, 0x71, 0x2f, 0xdf, 0x10, 0x3e, 0x7a, 0xd8, 0xbf, 0x9f, 0x9c, 0xe0, 0x96, 0x51, 0x04,
	0xa5, 0x28, 0x8b, 0x47, 0x2d, 0xa3, 0x92, 0x14, 0xb7, 0xad, 0x71, 0x93, 0x7b, 0x0b, 0x4f, 0x66,
	0xaa, 0x49, 0x2b, 0x45, 0xd9, 0xf1, 0xe8, 0xb7, 0x95, 0x10, 0x7c, 0xb8, 0x12, 0xcb, 0xa9, 0xbf,
	0xbb, 0x25, 0xff, 0x42, 0xfa, 0x7d, 0x4c, 0xae, 0xf0, 0xa9, 0xd2, 0x73, 0x70, 0xc6, 0x83, 0x1d,
	0xe4, 0xf9, 0x40, 0x29, 0xab, 0x9d, 0x23, 0x71, 0x68, 0xfd, 0x15, 0x25, 0x7d, 0x1c, 0xd7, 0xdf,
	0x25, 0xff, 0x53, 0x94, 0xb5, 0xaf, 0xcf, 0x59, 0xb3, 0x87, 0xd5, 0x7b, 0xd8, 0x7e, 0x0f, 0xbb,
	0x01, 0x53, 0x0e, 0xe3, 0xcd, 0xc7, 0x45, 0x34, 0x0a, 0xe5, 0xe1, 0x70, 0xb3, 0xa3, 0x68, 0xbb,
	0xa3, 0xe8, 0x73, 0x47, 0xd1, 0x6b, 0x45, 0xa3, 0x6d, 0x45, 0xa3, 0xf7, 0x8a, 0x46, 0x8f, 0x59,
	0x61, 0xfc, 0x78, 0x29, 0x59, 0x0e, 0x33, 0x5e, 0x13, 0xe3, 0x3f, 0xc4, 0xf8, 0x33, 0x6f, 0xc8,
	0xfa, 0x97, 0xb9, 0x76, 0xf2, 0x20, 0xa0, 0xe8, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x2a,
	0xa4, 0x5f, 0x6f, 0x01, 0x00, 0x00,
}

func (m *Withdraw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Withdraw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Withdraw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWithdraw(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.DepositorAccAddress) > 0 {
		i -= len(m.DepositorAccAddress)
		copy(dAtA[i:], m.DepositorAccAddress)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.DepositorAccAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.VaultID) > 0 {
		i -= len(m.VaultID)
		copy(dAtA[i:], m.VaultID)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.VaultID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RiskProfile) > 0 {
		i -= len(m.RiskProfile)
		copy(dAtA[i:], m.RiskProfile)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.RiskProfile)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintWithdraw(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintWithdraw(dAtA []byte, offset int, v uint64) int {
	offset -= sovWithdraw(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Withdraw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovWithdraw(uint64(m.Id))
	}
	l = len(m.RiskProfile)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = len(m.VaultID)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = len(m.DepositorAccAddress)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovWithdraw(uint64(l))
	return n
}

func sovWithdraw(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWithdraw(x uint64) (n int) {
	return sovWithdraw(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Withdraw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWithdraw
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
			return fmt.Errorf("proto: Withdraw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Withdraw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RiskProfile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RiskProfile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VaultID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositorAccAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositorAccAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWithdraw(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWithdraw
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
func skipWithdraw(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWithdraw
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
					return 0, ErrIntOverflowWithdraw
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
					return 0, ErrIntOverflowWithdraw
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
				return 0, ErrInvalidLengthWithdraw
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWithdraw
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWithdraw
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWithdraw        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWithdraw          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWithdraw = fmt.Errorf("proto: unexpected end of group")
)
