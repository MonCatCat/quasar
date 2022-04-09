// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmolpv/epoch_lp_info.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EpochLPInfo contains information about LP positions created at an epoch
type EpochLPInfo struct {
	EpochDay uint64     `protobuf:"varint,1,opt,name=epochDay,proto3" json:"epochDay,omitempty"`
	TotalLps uint64     `protobuf:"varint,2,opt,name=totalLps,proto3" json:"totalLps,omitempty"`
	TotalTVL types.Coin `protobuf:"bytes,3,opt,name=totalTVL,proto3" json:"totalTVL"`
}

func (m *EpochLPInfo) Reset()         { *m = EpochLPInfo{} }
func (m *EpochLPInfo) String() string { return proto.CompactTextString(m) }
func (*EpochLPInfo) ProtoMessage()    {}
func (*EpochLPInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_23ae8252961abea9, []int{0}
}
func (m *EpochLPInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochLPInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochLPInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochLPInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochLPInfo.Merge(m, src)
}
func (m *EpochLPInfo) XXX_Size() int {
	return m.Size()
}
func (m *EpochLPInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochLPInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EpochLPInfo proto.InternalMessageInfo

func (m *EpochLPInfo) GetEpochDay() uint64 {
	if m != nil {
		return m.EpochDay
	}
	return 0
}

func (m *EpochLPInfo) GetTotalLps() uint64 {
	if m != nil {
		return m.TotalLps
	}
	return 0
}

func (m *EpochLPInfo) GetTotalTVL() types.Coin {
	if m != nil {
		return m.TotalTVL
	}
	return types.Coin{}
}

// EpochDayInfo contains generic info about an epoch day.
// AUDIT NOTE - As of now it is not used, And not sure how it will be used.
// A possible use case is to verify and adjust the positions in case of network failures.
// We can add more application specific information in this struct.
type EpochDayInfo struct {
	EpochDay         uint64    `protobuf:"varint,1,opt,name=epochDay,proto3" json:"epochDay,omitempty"`
	StartBlockheight uint64    `protobuf:"varint,2,opt,name=startBlockheight,proto3" json:"startBlockheight,omitempty"`
	EndBlockheight   uint64    `protobuf:"varint,3,opt,name=endBlockheight,proto3" json:"endBlockheight,omitempty"`
	StartBlockTime   time.Time `protobuf:"bytes,4,opt,name=startBlockTime,proto3,stdtime" json:"startBlockTime" yaml:"blockTime"`
	EndBlockTime     time.Time `protobuf:"bytes,5,opt,name=endBlockTime,proto3,stdtime" json:"endBlockTime" yaml:"blockTime"`
}

func (m *EpochDayInfo) Reset()         { *m = EpochDayInfo{} }
func (m *EpochDayInfo) String() string { return proto.CompactTextString(m) }
func (*EpochDayInfo) ProtoMessage()    {}
func (*EpochDayInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_23ae8252961abea9, []int{1}
}
func (m *EpochDayInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochDayInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochDayInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochDayInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochDayInfo.Merge(m, src)
}
func (m *EpochDayInfo) XXX_Size() int {
	return m.Size()
}
func (m *EpochDayInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochDayInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EpochDayInfo proto.InternalMessageInfo

func (m *EpochDayInfo) GetEpochDay() uint64 {
	if m != nil {
		return m.EpochDay
	}
	return 0
}

func (m *EpochDayInfo) GetStartBlockheight() uint64 {
	if m != nil {
		return m.StartBlockheight
	}
	return 0
}

func (m *EpochDayInfo) GetEndBlockheight() uint64 {
	if m != nil {
		return m.EndBlockheight
	}
	return 0
}

func (m *EpochDayInfo) GetStartBlockTime() time.Time {
	if m != nil {
		return m.StartBlockTime
	}
	return time.Time{}
}

func (m *EpochDayInfo) GetEndBlockTime() time.Time {
	if m != nil {
		return m.EndBlockTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*EpochLPInfo)(nil), "abag.quasarnode.osmolpv.EpochLPInfo")
	proto.RegisterType((*EpochDayInfo)(nil), "abag.quasarnode.osmolpv.EpochDayInfo")
}

func init() { proto.RegisterFile("osmolpv/epoch_lp_info.proto", fileDescriptor_23ae8252961abea9) }

var fileDescriptor_23ae8252961abea9 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x8e, 0xda, 0x40,
	0x10, 0xc6, 0xbd, 0x40, 0xa2, 0x68, 0x41, 0x08, 0x59, 0x91, 0xe2, 0x38, 0x91, 0x8d, 0x5c, 0x44,
	0x88, 0x62, 0x57, 0x24, 0x5d, 0xd2, 0x39, 0x50, 0x44, 0xa2, 0x88, 0x10, 0x4a, 0x11, 0x45, 0x22,
	0x6b, 0xb3, 0xd8, 0x56, 0x6c, 0x8f, 0xc3, 0x2e, 0xe8, 0x68, 0xef, 0x09, 0x78, 0xa2, 0xab, 0x29,
	0x29, 0xaf, 0xe2, 0x4e, 0xf0, 0x06, 0xf7, 0x04, 0x27, 0xff, 0x83, 0x3b, 0xae, 0xb8, 0xe2, 0xba,
	0x9d, 0xfd, 0xbe, 0xf9, 0xf6, 0xb7, 0xa3, 0xc1, 0x1f, 0x40, 0x44, 0x10, 0x26, 0x4b, 0xca, 0x13,
	0x70, 0xfd, 0x49, 0x98, 0x4c, 0x82, 0x78, 0x06, 0x24, 0x99, 0x83, 0x04, 0xf5, 0x1d, 0x73, 0x98,
	0x47, 0xfe, 0x2f, 0x98, 0x60, 0xf3, 0x18, 0xa6, 0x9c, 0x14, 0x66, 0xfd, 0xad, 0x07, 0x1e, 0x64,
	0x1e, 0x9a, 0x9e, 0x72, 0xbb, 0x6e, 0xb8, 0xa9, 0x2e, 0xa8, 0xc3, 0x04, 0xa7, 0xcb, 0x9e, 0xc3,
	0x25, 0xeb, 0x51, 0x17, 0x82, 0xb8, 0xd0, 0x4d, 0x0f, 0xc0, 0x0b, 0x39, 0xcd, 0x2a, 0x67, 0x31,
	0xa3, 0x32, 0x88, 0xb8, 0x90, 0x2c, 0x4a, 0x72, 0x83, 0x75, 0x89, 0x70, 0x7d, 0x90, 0x72, 0x0c,
	0x7f, 0xfe, 0x88, 0x67, 0xa0, 0xea, 0xf8, 0x4d, 0x86, 0xd5, 0x67, 0x2b, 0x0d, 0xb5, 0x51, 0xa7,
	0x36, 0x3a, 0xd6, 0xa9, 0x26, 0x41, 0xb2, 0x70, 0x98, 0x08, 0xad, 0x92, 0x6b, 0x65, 0xad, 0x7e,
	0x2b, 0xb4, 0xf1, 0xaf, 0xa1, 0x56, 0x6d, 0xa3, 0x4e, 0xfd, 0xf3, 0x7b, 0x92, 0xb3, 0x91, 0x94,
	0x8d, 0x14, 0x6c, 0xe4, 0x3b, 0x04, 0xb1, 0x5d, 0xdb, 0xec, 0x4c, 0x65, 0x74, 0x6c, 0xb0, 0xae,
	0x2a, 0xb8, 0x31, 0x28, 0x5e, 0x79, 0x96, 0xa2, 0x8b, 0x5b, 0x42, 0xb2, 0xb9, 0xb4, 0x43, 0x70,
	0xff, 0xf9, 0x3c, 0xf0, 0x7c, 0x59, 0xd0, 0x3c, 0xb9, 0x57, 0x3f, 0xe1, 0x26, 0x8f, 0xa7, 0x0f,
	0x9d, 0xd5, 0xcc, 0x79, 0x76, 0xab, 0xfe, 0xc5, 0xcd, 0x53, 0xef, 0x38, 0x88, 0xb8, 0x56, 0xcb,
	0xfe, 0xa0, 0x93, 0x7c, 0x7e, 0xa4, 0x9c, 0x1f, 0x19, 0x97, 0xf3, 0xb3, 0x3f, 0xa6, 0x9f, 0xb8,
	0xdb, 0x99, 0xad, 0x15, 0x8b, 0xc2, 0xaf, 0x96, 0x53, 0xb6, 0x5a, 0xeb, 0x1b, 0x13, 0x8d, 0xce,
	0xf2, 0xd4, 0x3f, 0xb8, 0x51, 0xbe, 0x99, 0xe5, 0xbf, 0x7a, 0x61, 0xfe, 0xa3, 0x34, 0xbb, 0xbf,
	0xd9, 0x1b, 0x68, 0xbb, 0x37, 0xd0, 0xed, 0xde, 0x40, 0xeb, 0x83, 0xa1, 0x6c, 0x0f, 0x86, 0x72,
	0x7d, 0x30, 0x94, 0xdf, 0x5d, 0x2f, 0x90, 0xfe, 0xc2, 0x21, 0x2e, 0x44, 0x34, 0x5d, 0x2d, 0x7a,
	0x5a, 0x2d, 0x7a, 0x41, 0xcb, 0x4d, 0x94, 0xab, 0x84, 0x0b, 0xe7, 0x75, 0x46, 0xf1, 0xe5, 0x3e,
	0x00, 0x00, 0xff, 0xff, 0x51, 0x7b, 0x49, 0x29, 0xa1, 0x02, 0x00, 0x00,
}

func (m *EpochLPInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochLPInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochLPInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TotalTVL.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEpochLpInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.TotalLps != 0 {
		i = encodeVarintEpochLpInfo(dAtA, i, uint64(m.TotalLps))
		i--
		dAtA[i] = 0x10
	}
	if m.EpochDay != 0 {
		i = encodeVarintEpochLpInfo(dAtA, i, uint64(m.EpochDay))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EpochDayInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochDayInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochDayInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndBlockTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEpochLpInfo(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartBlockTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintEpochLpInfo(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	if m.EndBlockheight != 0 {
		i = encodeVarintEpochLpInfo(dAtA, i, uint64(m.EndBlockheight))
		i--
		dAtA[i] = 0x18
	}
	if m.StartBlockheight != 0 {
		i = encodeVarintEpochLpInfo(dAtA, i, uint64(m.StartBlockheight))
		i--
		dAtA[i] = 0x10
	}
	if m.EpochDay != 0 {
		i = encodeVarintEpochLpInfo(dAtA, i, uint64(m.EpochDay))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEpochLpInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpochLpInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EpochLPInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochDay != 0 {
		n += 1 + sovEpochLpInfo(uint64(m.EpochDay))
	}
	if m.TotalLps != 0 {
		n += 1 + sovEpochLpInfo(uint64(m.TotalLps))
	}
	l = m.TotalTVL.Size()
	n += 1 + l + sovEpochLpInfo(uint64(l))
	return n
}

func (m *EpochDayInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochDay != 0 {
		n += 1 + sovEpochLpInfo(uint64(m.EpochDay))
	}
	if m.StartBlockheight != 0 {
		n += 1 + sovEpochLpInfo(uint64(m.StartBlockheight))
	}
	if m.EndBlockheight != 0 {
		n += 1 + sovEpochLpInfo(uint64(m.EndBlockheight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartBlockTime)
	n += 1 + l + sovEpochLpInfo(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndBlockTime)
	n += 1 + l + sovEpochLpInfo(uint64(l))
	return n
}

func sovEpochLpInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpochLpInfo(x uint64) (n int) {
	return sovEpochLpInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EpochLPInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpochLpInfo
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
			return fmt.Errorf("proto: EpochLPInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochLPInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochDay", wireType)
			}
			m.EpochDay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochLpInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochDay |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalLps", wireType)
			}
			m.TotalLps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochLpInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalLps |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalTVL", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochLpInfo
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
				return ErrInvalidLengthEpochLpInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpochLpInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalTVL.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEpochLpInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpochLpInfo
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
func (m *EpochDayInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpochLpInfo
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
			return fmt.Errorf("proto: EpochDayInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochDayInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochDay", wireType)
			}
			m.EpochDay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochLpInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochDay |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlockheight", wireType)
			}
			m.StartBlockheight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochLpInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlockheight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockheight", wireType)
			}
			m.EndBlockheight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochLpInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlockheight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochLpInfo
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
				return ErrInvalidLengthEpochLpInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpochLpInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochLpInfo
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
				return ErrInvalidLengthEpochLpInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpochLpInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEpochLpInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpochLpInfo
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
func skipEpochLpInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpochLpInfo
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
					return 0, ErrIntOverflowEpochLpInfo
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
					return 0, ErrIntOverflowEpochLpInfo
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
				return 0, ErrInvalidLengthEpochLpInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpochLpInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpochLpInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpochLpInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpochLpInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpochLpInfo = fmt.Errorf("proto: unexpected end of group")
)
