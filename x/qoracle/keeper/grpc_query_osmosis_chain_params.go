package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/quasarlabs/quasarnode/x/qoracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OsmosisChainParams(goCtx context.Context, req *types.QueryOsmosisChainParamsRequest) (*types.QueryOsmosisChainParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryOsmosisChainParamsResponse{
		EpochsInfo:          k.GetOsmosisEpochsInfo(ctx),
		LockableDurations:   k.GetOsmosisLockableDurations(ctx),
		MintParams:          k.GetOsmosisMintParams(ctx),
		MintEpochProvisions: k.GetOsmosisMintEpochProvisions(ctx),
		DistrInfo:           k.GetOsmosisDistrInfo(ctx),
	}, nil
}
