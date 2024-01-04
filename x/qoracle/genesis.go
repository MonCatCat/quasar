package qoracle

import (
	genesistypes "github.com/MonCatCat/quasar/x/qoracle/genesis/types"
	qoraclekeeper "github.com/MonCatCat/quasar/x/qoracle/keeper"
	qosmokeeper "github.com/MonCatCat/quasar/x/qoracle/osmosis/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(
	ctx sdk.Context,
	qKeeper qoraclekeeper.Keeper,
	osmoKeeper qosmokeeper.Keeper,
	state genesistypes.GenesisState,
) {
	qKeeper.SetParams(ctx, state.Params)
	qosmokeeper.InitGenesis(ctx, osmoKeeper, state.OsmosisGenesisState)
}

func ExportGenesis(
	ctx sdk.Context,
	qKeeper qoraclekeeper.Keeper,
	osmoKeeper qosmokeeper.Keeper,
) *genesistypes.GenesisState {
	return genesistypes.NewGenesisState(
		qKeeper.GetParams(ctx),
		qosmokeeper.ExportGenesis(ctx, osmoKeeper),
	)
}
