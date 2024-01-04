package v0

import (
	qvestingkeeper "github.com/MonCatCat/quasar/x/qvesting/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SetQVestingParams(ctx sdk.Context, icqKeeper *qvestingkeeper.Keeper) {
	setQVestingParams(ctx, icqKeeper)
}
