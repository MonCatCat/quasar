package types

import (
	balancerpool "github.com/MonCatCat/quasar/osmosis/gamm/pool-models/balancer"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/gogo/protobuf/proto"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*proto.Message)(nil), &balancerpool.Pool{})
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
