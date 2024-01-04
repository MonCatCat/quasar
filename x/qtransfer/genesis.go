package qtransfer

import (
	"github.com/MonCatCat/quasar/x/qtransfer/keeper"
	"github.com/MonCatCat/quasar/x/qtransfer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the qtransfer state and creates the intermediate account for wasm hooks.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	err := k.CreateIntermediateAccountAccount(ctx)
	if err != nil {
		k.Logger(ctx).Error("InitGenesis failed during CreateIntermediateAccountAccount",
			"error", err)
		panic(err)
	}

	k.SetParams(ctx, state.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	gs := types.DefaultGenesisState()
	gs.Params = k.GetParams(ctx)

	return gs
}
