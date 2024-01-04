package testutil

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"testing"

	"github.com/MonCatCat/quasar/app"
	"github.com/MonCatCat/quasar/testutil/keeper"
	"github.com/MonCatCat/quasar/testutil/mock"
	epochskeeper "github.com/MonCatCat/quasar/x/epochs/keeper"
	qoraclekeeper "github.com/MonCatCat/quasar/x/qoracle/keeper"
	qosmokeeper "github.com/MonCatCat/quasar/x/qoracle/osmosis/keeper"
	qosmotypes "github.com/MonCatCat/quasar/x/qoracle/osmosis/types"
	qtransferkeeper "github.com/MonCatCat/quasar/x/qtransfer/keeper"
	qvestingkeeper "github.com/MonCatCat/quasar/x/qvesting/keeper"
	tfkeeper "github.com/MonCatCat/quasar/x/tokenfactory/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func init() {
	// Set prefixes
	accountPubKeyPrefix := app.AccountAddressPrefix + "pub"
	validatorAddressPrefix := app.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := app.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := app.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := app.AccountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()
}
func CreateRandomAccounts(numAccts int) []sdk.AccAddress {
	testAddrs := make([]sdk.AccAddress, numAccts)
	for i := 0; i < numAccts; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddrs[i] = sdk.AccAddress(pk.Address())
	}

	return testAddrs
}

// FundAcc funds target address with specified amount.
func (ts *TestSetup) FundAcc(t testing.TB, acc sdk.AccAddress, amounts sdk.Coins) {
	err := simapp.FundAccount(ts.Keepers.BankKeeper, ts.Ctx, acc, amounts)
	require.NoError(t, err)
}

// FundModuleAcc funds target modules with specified amount.
func (ts *TestSetup) FundModuleAcc(t testing.TB, moduleName string, amounts sdk.Coins) {
	err := simapp.FundModuleAccount(ts.Keepers.BankKeeper, ts.Ctx, moduleName, amounts)
	require.NoError(t, err)
}

func (ts *TestSetup) MintCoins(t testing.TB, coins sdk.Coins) {
	err := ts.Keepers.BankKeeper.MintCoins(ts.Ctx, minttypes.ModuleName, coins)
	require.NoError(t, err)
}

func NewTestSetup(t testing.TB, controller ...*gomock.Controller) *TestSetup {
	// Test setup params

	logger := log.TestingLogger()
	// Use nop logger if logging becomes too verbose for test output
	// logger := log.NewNopLogger()
	logger.Debug("creating test setup")

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)
	encodingConfig := app.MakeEncodingConfig()

	// Mocks

	// If no controller is given, no mock is needed so we don't need to check that mocks were called
	var ctl *gomock.Controller
	switch len(controller) {
	case 0:
		ctl = gomock.NewController(t)
	default:
		ctl = controller[0]
	}
	ibcClientKeeperMock := mock.NewMockClientKeeper(ctl)
	ibcChannelKeeperMock := mock.NewMockChannelKeeper(ctl)
	icaControllerKeeperMock := mock.NewMockICAControllerKeeper(ctl)
	ics4WrapperMock := mock.NewMockICS4Wrapper(ctl)
	ibcPortKeeperMock := mock.NewMockPortKeeper(ctl)
	// Set BindPort method for mock and return a mock capability
	ibcPortKeeperMock.EXPECT().BindPort(gomock.Any(), gomock.Any()).AnyTimes().Return(capabilitytypes.NewCapability(1))
	// ibcClientKeeperMock := mock.NewMockClientKeeper(ctl)

	// Keepers

	// Create a factory first to easily create keepers
	factory := keeper.NewKeeperFactory(db, stateStore, ctx, encodingConfig)

	maccPerms := factory.TestModuleAccountPerms()
	blockedMaccAddresses := factory.BlockedModuleAccountAddrs(maccPerms)

	paramsKeeper := factory.ParamsKeeper()
	epochsKeeper := factory.EpochsKeeper(paramsKeeper)
	accountKeeper := factory.AccountKeeper(paramsKeeper, maccPerms)
	bankKeeper := factory.BankKeeper(paramsKeeper, accountKeeper, blockedMaccAddresses)
	capabilityKeeper := factory.CapabilityKeeper()
	capabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	stakingKeeper := factory.StakingKeeper(paramsKeeper, accountKeeper, bankKeeper)
	distrKeeper := factory.DistributionKeeper(paramsKeeper, accountKeeper, bankKeeper, stakingKeeper,
		"feeCollectorName", blockedMaccAddresses)
	qosmoScopedKeeper := capabilityKeeper.ScopeToModule(qosmotypes.SubModuleName)

	qoracleKeeper := factory.QoracleKeeper(paramsKeeper, authtypes.NewModuleAddress(govtypes.ModuleName).String())
	qosmosisKeeper := factory.QosmosisKeeper(paramsKeeper, authtypes.NewModuleAddress(govtypes.ModuleName).String(), ibcClientKeeperMock, ics4WrapperMock, ibcChannelKeeperMock, ibcPortKeeperMock, qosmoScopedKeeper, qoracleKeeper)
	qoracleKeeper.RegisterPoolOracle(qosmosisKeeper)
	qoracleKeeper.Seal()
	qtransferkeeper := factory.QTransferKeeper(paramsKeeper, accountKeeper)
	qvestingKeeper := factory.QVestingKeeper(paramsKeeper, accountKeeper, bankKeeper)
	tfKeeper := factory.TfKeeper(paramsKeeper, accountKeeper, bankKeeper, distrKeeper)

	// Note: the relative order of LoadLatestVersion and Set*DefaultParams is important.
	// Setting params before loading stores causes store does not exist error.
	// LoadLatestVersion must not be called again after setting params, as reloading stores clears set params.

	require.NoError(t, factory.StateStore.LoadLatestVersion())

	factory.SetQoracleDefaultParams(qoracleKeeper)
	factory.SetQosmosisDefaultParams(qosmosisKeeper)
	testAccts := CreateRandomAccounts(3)

	//  Init Genesis of Keepers

	distrGendata := distrtypes.GenesisState{Params: distrtypes.DefaultParams()}
	distrKeeper.InitGenesis(ctx, distrGendata)
	return &TestSetup{
		Ctx: ctx,
		Cdc: encodingConfig.Marshaler,

		Mocks: &testMocks{
			ICAControllerKeeperMock: icaControllerKeeperMock,
		},

		Keepers: &testKeepers{
			ParamsKeeper:     paramsKeeper,
			EpochsKeeper:     epochsKeeper,
			AccountKeeper:    accountKeeper,
			BankKeeper:       bankKeeper,
			CapabilityKeeper: capabilityKeeper,
			QoracleKeeper:    qoracleKeeper,
			QosmosisKeeper:   qosmosisKeeper,
			QTransfer:        qtransferkeeper,
			QVestingKeeper:   qvestingKeeper,
			TfKeeper:         tfKeeper,
		},
		TestAccs: testAccts,
	}
}

type TestSetup struct {
	Ctx sdk.Context
	Cdc codec.Codec

	Keepers  *testKeepers
	Mocks    *testMocks
	TestAccs []sdk.AccAddress
}

type testMocks struct {
	ICAControllerKeeperMock *mock.MockICAControllerKeeper
}

type testKeepers struct {
	ParamsKeeper      paramskeeper.Keeper
	EpochsKeeper      *epochskeeper.Keeper
	AccountKeeper     authkeeper.AccountKeeper
	BankKeeper        bankkeeper.Keeper
	StakingKeeper     stakingKeeper.Keeper
	DistributedKeeper distrkeeper.Keeper
	CapabilityKeeper  capabilitykeeper.Keeper
	QoracleKeeper     qoraclekeeper.Keeper
	QosmosisKeeper    qosmokeeper.Keeper
	QTransfer         qtransferkeeper.Keeper
	QVestingKeeper    qvestingkeeper.Keeper
	TfKeeper          tfkeeper.Keeper
}
