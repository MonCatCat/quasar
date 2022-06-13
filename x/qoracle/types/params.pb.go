// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qoracle/params.proto

package types

import (
	fmt "fmt"
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

type OneHopIbcDenomMapping struct {
	OriginName string `protobuf:"bytes,1,opt,name=originName,proto3" json:"originName,omitempty" yaml:"origin_name"`
	Quasar     string `protobuf:"bytes,2,opt,name=quasar,proto3" json:"quasar,omitempty" yaml:"quasar"`
	Osmo       string `protobuf:"bytes,3,opt,name=osmo,proto3" json:"osmo,omitempty" yaml:"osmo"`
}

func (m *OneHopIbcDenomMapping) Reset()         { *m = OneHopIbcDenomMapping{} }
func (m *OneHopIbcDenomMapping) String() string { return proto.CompactTextString(m) }
func (*OneHopIbcDenomMapping) ProtoMessage()    {}
func (*OneHopIbcDenomMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_95f4338228c132f2, []int{0}
}
func (m *OneHopIbcDenomMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OneHopIbcDenomMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OneHopIbcDenomMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OneHopIbcDenomMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneHopIbcDenomMapping.Merge(m, src)
}
func (m *OneHopIbcDenomMapping) XXX_Size() int {
	return m.Size()
}
func (m *OneHopIbcDenomMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_OneHopIbcDenomMapping.DiscardUnknown(m)
}

var xxx_messageInfo_OneHopIbcDenomMapping proto.InternalMessageInfo

func (m *OneHopIbcDenomMapping) GetOriginName() string {
	if m != nil {
		return m.OriginName
	}
	return ""
}

func (m *OneHopIbcDenomMapping) GetQuasar() string {
	if m != nil {
		return m.Quasar
	}
	return ""
}

func (m *OneHopIbcDenomMapping) GetOsmo() string {
	if m != nil {
		return m.Osmo
	}
	return ""
}

type BandchainParams struct {
	OraclePortId            string `protobuf:"bytes,1,opt,name=oracle_port_id,json=oraclePortId,proto3" json:"oracle_port_id,omitempty" yaml:"oracle_port_id"`
	OracleVersion           string `protobuf:"bytes,2,opt,name=oracle_version,json=oracleVersion,proto3" json:"oracle_version,omitempty" yaml:"oracle_version"`
	OracleActiveChannelPath string `protobuf:"bytes,3,opt,name=oracle_active_channel_path,json=oracleActiveChannelPath,proto3" json:"oracle_active_channel_path,omitempty" yaml:"oracle_active_channel_path"`
	CoinRatesScriptId       uint64 `protobuf:"varint,4,opt,name=coin_rates_script_id,json=coinRatesScriptId,proto3" json:"coin_rates_script_id,omitempty" yaml:"coin_rates_script_id"`
}

func (m *BandchainParams) Reset()         { *m = BandchainParams{} }
func (m *BandchainParams) String() string { return proto.CompactTextString(m) }
func (*BandchainParams) ProtoMessage()    {}
func (*BandchainParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_95f4338228c132f2, []int{1}
}
func (m *BandchainParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BandchainParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BandchainParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BandchainParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BandchainParams.Merge(m, src)
}
func (m *BandchainParams) XXX_Size() int {
	return m.Size()
}
func (m *BandchainParams) XXX_DiscardUnknown() {
	xxx_messageInfo_BandchainParams.DiscardUnknown(m)
}

var xxx_messageInfo_BandchainParams proto.InternalMessageInfo

func (m *BandchainParams) GetOraclePortId() string {
	if m != nil {
		return m.OraclePortId
	}
	return ""
}

func (m *BandchainParams) GetOracleVersion() string {
	if m != nil {
		return m.OracleVersion
	}
	return ""
}

func (m *BandchainParams) GetOracleActiveChannelPath() string {
	if m != nil {
		return m.OracleActiveChannelPath
	}
	return ""
}

func (m *BandchainParams) GetCoinRatesScriptId() uint64 {
	if m != nil {
		return m.CoinRatesScriptId
	}
	return 0
}

// Params defines the parameters for the module.
type Params struct {
	BandchainParams BandchainParams          `protobuf:"bytes,1,opt,name=bandchain_params,json=bandchainParams,proto3" json:"bandchain_params" yaml:"bandchain_params"`
	OracleAccounts  string                   `protobuf:"bytes,2,opt,name=oracleAccounts,proto3" json:"oracleAccounts,omitempty" yaml:"oracle_accounts"`
	StableDenoms    []string                 `protobuf:"bytes,3,rep,name=stableDenoms,proto3" json:"stableDenoms,omitempty" yaml:"stable_denoms"`
	OneHopDenomMap  []*OneHopIbcDenomMapping `protobuf:"bytes,4,rep,name=oneHopDenomMap,proto3" json:"oneHopDenomMap,omitempty" yaml:"onehop_ibcdenoms"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_95f4338228c132f2, []int{2}
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

func (m *Params) GetBandchainParams() BandchainParams {
	if m != nil {
		return m.BandchainParams
	}
	return BandchainParams{}
}

func (m *Params) GetOracleAccounts() string {
	if m != nil {
		return m.OracleAccounts
	}
	return ""
}

func (m *Params) GetStableDenoms() []string {
	if m != nil {
		return m.StableDenoms
	}
	return nil
}

func (m *Params) GetOneHopDenomMap() []*OneHopIbcDenomMapping {
	if m != nil {
		return m.OneHopDenomMap
	}
	return nil
}

func init() {
	proto.RegisterType((*OneHopIbcDenomMapping)(nil), "abag.quasarnode.qoracle.OneHopIbcDenomMapping")
	proto.RegisterType((*BandchainParams)(nil), "abag.quasarnode.qoracle.BandchainParams")
	proto.RegisterType((*Params)(nil), "abag.quasarnode.qoracle.Params")
}

func init() { proto.RegisterFile("qoracle/params.proto", fileDescriptor_95f4338228c132f2) }

var fileDescriptor_95f4338228c132f2 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0x8d, 0x9b, 0x28, 0x52, 0x2f, 0x6d, 0x43, 0xad, 0xb4, 0x31, 0xa9, 0xe4, 0x0b, 0x87, 0x90,
	0x02, 0x83, 0x2d, 0x15, 0x89, 0xa1, 0x42, 0x82, 0x9a, 0x0e, 0x64, 0x00, 0x22, 0x23, 0x31, 0xb0,
	0x58, 0x67, 0xfb, 0x64, 0x5b, 0x8a, 0xef, 0x5c, 0xdf, 0xa5, 0xa2, 0xff, 0x82, 0x91, 0x0d, 0x7e,
	0x4e, 0xc5, 0xd4, 0x91, 0xc9, 0x42, 0xc9, 0x3f, 0xf0, 0xca, 0x82, 0x72, 0x67, 0xd3, 0xd6, 0xa4,
	0x9b, 0xfd, 0xbd, 0x77, 0xef, 0xbe, 0xf7, 0xee, 0xfb, 0xc0, 0xe0, 0x9c, 0xe5, 0x38, 0x98, 0x13,
	0x3b, 0xc3, 0x39, 0x4e, 0xb9, 0x95, 0xe5, 0x4c, 0x30, 0x7d, 0x88, 0x7d, 0x1c, 0x59, 0xe7, 0x0b,
	0xcc, 0x71, 0x4e, 0x59, 0x48, 0xac, 0x8a, 0x35, 0x1a, 0x44, 0x2c, 0x62, 0x92, 0x63, 0xaf, 0xbf,
	0x14, 0x1d, 0x7d, 0xd7, 0xc0, 0xc1, 0x07, 0x4a, 0xde, 0xb2, 0x6c, 0xea, 0x07, 0x67, 0x84, 0xb2,
	0xf4, 0x1d, 0xce, 0xb2, 0x84, 0x46, 0xfa, 0x0b, 0x00, 0x58, 0x9e, 0x44, 0x09, 0x7d, 0x8f, 0x53,
	0x62, 0x68, 0x63, 0x6d, 0xb2, 0xed, 0x1c, 0x96, 0x05, 0xd4, 0x2f, 0x71, 0x3a, 0x3f, 0x41, 0x0a,
	0xf3, 0x28, 0x4e, 0x09, 0x72, 0x6f, 0x31, 0xf5, 0xa7, 0xa0, 0xab, 0x6e, 0x37, 0xb6, 0xe4, 0x99,
	0xfd, 0xb2, 0x80, 0xbb, 0xea, 0x8c, 0xaa, 0x23, 0xb7, 0x22, 0xe8, 0x8f, 0x41, 0x87, 0xf1, 0x94,
	0x19, 0x6d, 0x49, 0xec, 0x97, 0x05, 0xec, 0x55, 0xe2, 0x3c, 0x65, 0xc8, 0x95, 0x20, 0xfa, 0xb9,
	0x05, 0xfa, 0x0e, 0xa6, 0x61, 0x10, 0xe3, 0x84, 0xce, 0xa4, 0x55, 0xfd, 0x15, 0xd8, 0x53, 0xae,
	0xbc, 0x8c, 0xe5, 0xc2, 0x4b, 0xc2, 0xaa, 0xbf, 0x87, 0x65, 0x01, 0x0f, 0xea, 0xfe, 0x6e, 0xe3,
	0xc8, 0xdd, 0x51, 0x85, 0x19, 0xcb, 0xc5, 0x34, 0xd4, 0x5f, 0xff, 0x13, 0xb8, 0x20, 0x39, 0x4f,
	0x18, 0xad, 0x9a, 0xfd, 0x5f, 0xa0, 0xc2, 0x91, 0xbb, 0xab, 0x0a, 0x9f, 0xd4, 0xbf, 0xee, 0x83,
	0x51, 0xc5, 0xc0, 0x81, 0x48, 0x2e, 0x88, 0x17, 0xc4, 0x98, 0x52, 0x32, 0xf7, 0x32, 0x2c, 0xe2,
	0xca, 0xd1, 0x93, 0xb2, 0x80, 0x8f, 0xee, 0xa8, 0x6d, 0xe0, 0x22, 0x77, 0xa8, 0xc0, 0x53, 0x89,
	0xbd, 0x51, 0xd0, 0x0c, 0x8b, 0x58, 0x9f, 0x81, 0x41, 0xc0, 0x12, 0xea, 0xe5, 0x58, 0x10, 0xee,
	0xf1, 0x20, 0x4f, 0x32, 0x69, 0xb6, 0x33, 0xd6, 0x26, 0x1d, 0x07, 0x96, 0x05, 0x3c, 0x52, 0xea,
	0x9b, 0x58, 0xc8, 0xdd, 0x5f, 0x97, 0xdd, 0x75, 0xf5, 0xa3, 0x2c, 0x4e, 0x43, 0xf4, 0x67, 0x0b,
	0x74, 0xab, 0x0c, 0x05, 0x78, 0xe0, 0xd7, 0xb1, 0x7a, 0x6a, 0x84, 0x64, 0x8a, 0xbd, 0xe3, 0x89,
	0x75, 0xcf, 0x0c, 0x59, 0x8d, 0x77, 0x70, 0xe0, 0x55, 0x01, 0x5b, 0x65, 0x01, 0x87, 0xaa, 0x8d,
	0xa6, 0x1e, 0x72, 0xfb, 0x7e, 0xe3, 0xe5, 0x9c, 0x3a, 0xf8, 0xd3, 0x20, 0x60, 0x0b, 0x2a, 0x78,
	0x15, 0xfc, 0xa8, 0x2c, 0xe0, 0x61, 0x23, 0x2a, 0x45, 0x40, 0x6e, 0xe3, 0x84, 0xfe, 0x12, 0xec,
	0x70, 0x81, 0xfd, 0x39, 0x91, 0xf3, 0xca, 0x8d, 0xf6, 0xb8, 0x3d, 0xd9, 0x76, 0x8c, 0xb2, 0x80,
	0x03, 0xa5, 0xa0, 0x50, 0x2f, 0x94, 0x30, 0x72, 0xef, 0xb0, 0x75, 0x0e, 0xf6, 0x98, 0x1c, 0xf8,
	0x7a, 0xda, 0x8d, 0xce, 0xb8, 0x3d, 0xe9, 0x1d, 0x5b, 0xf7, 0xba, 0xde, 0xb8, 0x1f, 0xce, 0xd1,
	0x8d, 0x6f, 0x46, 0x49, 0xcc, 0x32, 0x2f, 0xf1, 0x83, 0xfa, 0xca, 0xc6, 0x15, 0x27, 0x9d, 0x6f,
	0x3f, 0x60, 0xcb, 0x39, 0xbb, 0x5a, 0x9a, 0xda, 0xf5, 0xd2, 0xd4, 0x7e, 0x2f, 0x4d, 0xed, 0xeb,
	0xca, 0x6c, 0x5d, 0xaf, 0xcc, 0xd6, 0xaf, 0x95, 0xd9, 0xfa, 0xfc, 0x2c, 0x4a, 0x44, 0xbc, 0xf0,
	0xad, 0x80, 0xa5, 0xf6, 0xba, 0x0d, 0xfb, 0xa6, 0x0d, 0xfb, 0x8b, 0x5d, 0x2f, 0xba, 0xb8, 0xcc,
	0x08, 0xf7, 0xbb, 0x72, 0x73, 0x9f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x32, 0x58, 0x4c, 0xfb,
	0x00, 0x04, 0x00, 0x00,
}

func (m *OneHopIbcDenomMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneHopIbcDenomMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OneHopIbcDenomMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Osmo) > 0 {
		i -= len(m.Osmo)
		copy(dAtA[i:], m.Osmo)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Osmo)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Quasar) > 0 {
		i -= len(m.Quasar)
		copy(dAtA[i:], m.Quasar)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Quasar)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OriginName) > 0 {
		i -= len(m.OriginName)
		copy(dAtA[i:], m.OriginName)
		i = encodeVarintParams(dAtA, i, uint64(len(m.OriginName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BandchainParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BandchainParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BandchainParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CoinRatesScriptId != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CoinRatesScriptId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.OracleActiveChannelPath) > 0 {
		i -= len(m.OracleActiveChannelPath)
		copy(dAtA[i:], m.OracleActiveChannelPath)
		i = encodeVarintParams(dAtA, i, uint64(len(m.OracleActiveChannelPath)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OracleVersion) > 0 {
		i -= len(m.OracleVersion)
		copy(dAtA[i:], m.OracleVersion)
		i = encodeVarintParams(dAtA, i, uint64(len(m.OracleVersion)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OraclePortId) > 0 {
		i -= len(m.OraclePortId)
		copy(dAtA[i:], m.OraclePortId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.OraclePortId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if len(m.OneHopDenomMap) > 0 {
		for iNdEx := len(m.OneHopDenomMap) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OneHopDenomMap[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.StableDenoms) > 0 {
		for iNdEx := len(m.StableDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.StableDenoms[iNdEx])
			copy(dAtA[i:], m.StableDenoms[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.StableDenoms[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OracleAccounts) > 0 {
		i -= len(m.OracleAccounts)
		copy(dAtA[i:], m.OracleAccounts)
		i = encodeVarintParams(dAtA, i, uint64(len(m.OracleAccounts)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.BandchainParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *OneHopIbcDenomMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OriginName)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Quasar)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Osmo)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *BandchainParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OraclePortId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.OracleVersion)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.OracleActiveChannelPath)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.CoinRatesScriptId != 0 {
		n += 1 + sovParams(uint64(m.CoinRatesScriptId))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BandchainParams.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.OracleAccounts)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.StableDenoms) > 0 {
		for _, s := range m.StableDenoms {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.OneHopDenomMap) > 0 {
		for _, e := range m.OneHopDenomMap {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OneHopIbcDenomMapping) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: OneHopIbcDenomMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OneHopIbcDenomMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quasar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quasar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Osmo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Osmo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *BandchainParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: BandchainParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BandchainParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OraclePortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OraclePortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleActiveChannelPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleActiveChannelPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinRatesScriptId", wireType)
			}
			m.CoinRatesScriptId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoinRatesScriptId |= uint64(b&0x7F) << shift
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandchainParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BandchainParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleAccounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleAccounts = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StableDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StableDenoms = append(m.StableDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OneHopDenomMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OneHopDenomMap = append(m.OneHopDenomMap, &OneHopIbcDenomMapping{})
			if err := m.OneHopDenomMap[len(m.OneHopDenomMap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
