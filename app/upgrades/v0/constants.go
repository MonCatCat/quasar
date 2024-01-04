package v0

import (
	"github.com/MonCatCat/quasar/app/upgrades"
	qvestingtypes "github.com/MonCatCat/quasar/x/qvesting/types"
	tftypes "github.com/MonCatCat/quasar/x/tokenfactory/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz/keeper"
)

// UpgradeName defines the on-chain upgrade name for the Quasar chain v1.0.0 upgrade.
const UpgradeName = "v1"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{qvestingtypes.StoreKey, tftypes.StoreKey, authztypes.StoreKey},
		Deleted: []string{},
	},
}
