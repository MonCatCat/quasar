package keeper

import (
	epochstypes "github.com/MonCatCat/quasar/x/epochs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Hooks wrapper struct for qoracle bandchain keeper.
type EpochHooks struct {
	k Keeper
}

var _ epochstypes.EpochHooks = EpochHooks{}

// Return the wrapper struct.
func (k Keeper) EpochHooks() EpochHooks {
	return EpochHooks{k}
}

func (h EpochHooks) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {}

func (h EpochHooks) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	h.k.AfterEpochEnd(ctx, epochIdentifier, epochNumber)
}
