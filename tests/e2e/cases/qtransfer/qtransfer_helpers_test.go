package qtransfer

import (
	"context"
	"fmt"
	testsuite "github.com/quasarlabs/quasarnode/tests/e2e/suite"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
)

// deployPrimitives stores the contract, initiates it and returns the contract address.
func (s *Qtransfer) deployPrimitives(ctx context.Context, acc *ibc.Wallet, filePath, label string, initArgs1, initArgs2, initArgs3 any) {
	accAddress := acc.Bech32Address(s.Quasar().Config().Bech32Prefix)

	// Read the contract from os file
	contract, err := os.ReadFile(filePath)
	s.Require().NoError(err)

	// Store the contract in chain
	codeID := s.StoreContractCode(ctx, s.Quasar(), acc.KeyName, contract)
	s.PrimitiveStoreID = codeID

	// instantiate the contracts
	res := s.InstantiateContract(ctx, s.Quasar(), acc.KeyName, codeID, label, accAddress, sdk.NewCoins(), initArgs1)
	s.Require().NotEmpty(res.Address)
	s.LpStrategyContractAddress1 = res.Address

	// create channels for all the instantiated contracts address 1
	s.CreateChannel(
		ctx,
		testsuite.Quasar2OsmosisPath,
		fmt.Sprintf("wasm.%s", s.LpStrategyContractAddress1),
		"icqhost",
		ibc.Unordered,
		"icq-1",
	)

	s.CreateChannel(
		ctx,
		testsuite.Quasar2OsmosisPath,
		fmt.Sprintf("wasm.%s", s.LpStrategyContractAddress1),
		"icahost",
		ibc.Ordered,
		fmt.Sprintf(
			`{"version":"ics27-1","encoding":"proto3","tx_type":"sdk_multi_msg","controller_connection_id":"%s","host_connection_id":"%s"}`,
			s.Quasar2OsmosisConn.Id,
			s.Quasar2OsmosisConn.Counterparty.ConnectionId,
		),
	)

	res = s.InstantiateContract(ctx, s.Quasar(), acc.KeyName, codeID, label, accAddress, sdk.NewCoins(), initArgs2)
	s.Require().NotEmpty(res.Address)
	s.LpStrategyContractAddress2 = res.Address

	// create channels for all the instantiated contracts address 2
	s.CreateChannel(
		ctx,
		testsuite.Quasar2OsmosisPath,
		fmt.Sprintf("wasm.%s", s.LpStrategyContractAddress2),
		"icqhost",
		ibc.Unordered,
		"icq-1",
	)

	s.CreateChannel(
		ctx,
		testsuite.Quasar2OsmosisPath,
		fmt.Sprintf("wasm.%s", s.LpStrategyContractAddress2),
		"icahost",
		ibc.Ordered,
		fmt.Sprintf(
			`{"version":"ics27-1","encoding":"proto3","tx_type":"sdk_multi_msg","controller_connection_id":"%s","host_connection_id":"%s"}`,
			s.Quasar2OsmosisConn.Id,
			s.Quasar2OsmosisConn.Counterparty.ConnectionId,
		),
	)

	res = s.InstantiateContract(ctx, s.Quasar(), acc.KeyName, codeID, label, accAddress, sdk.NewCoins(), initArgs3)
	s.Require().NotEmpty(res.Address)
	s.LpStrategyContractAddress3 = res.Address

	// create channels for all the instantiated contracts address 3
	s.CreateChannel(
		ctx,
		testsuite.Quasar2OsmosisPath,
		fmt.Sprintf("wasm.%s", s.LpStrategyContractAddress3),
		"icqhost",
		ibc.Unordered,
		"icq-1",
	)

	s.CreateChannel(
		ctx,
		testsuite.Quasar2OsmosisPath,
		fmt.Sprintf("wasm.%s", s.LpStrategyContractAddress3),
		"icahost",
		ibc.Ordered,
		fmt.Sprintf(
			`{"version":"ics27-1","encoding":"proto3","tx_type":"sdk_multi_msg","controller_connection_id":"%s","host_connection_id":"%s"}`,
			s.Quasar2OsmosisConn.Id,
			s.Quasar2OsmosisConn.Counterparty.ConnectionId,
		),
	)
}

// deployRewardsContract stores the contract
func (s *Qtransfer) deployRewardsContract(ctx context.Context, acc *ibc.Wallet, filePath string) {
	// Read the contract from os file
	contract, err := os.ReadFile(filePath)
	s.Require().NoError(err)

	// Store the contract in chain
	codeID := s.StoreContractCode(ctx, s.Quasar(), acc.KeyName, contract)
	s.RewardsStoreID = codeID
}

// deployVault stores the contract, initiates it and returns the contract address.
func (s *Qtransfer) deployVault(ctx context.Context, acc *ibc.Wallet, filePath, label string, initArgs any) string {
	accAddress := acc.Bech32Address(s.Quasar().Config().Bech32Prefix)

	// Read the contract from os file
	contract, err := os.ReadFile(filePath)
	s.Require().NoError(err)

	// Store the contract in chain
	codeID := s.StoreContractCode(ctx, s.Quasar(), acc.KeyName, contract)
	s.VaultStoreID = codeID

	res := s.InstantiateContract(ctx, s.Quasar(), acc.KeyName, codeID, label, accAddress, sdk.NewCoins(), initArgs)
	s.Require().NotEmpty(res.Address)

	return res.Address
}

func (s *Qtransfer) setDepositorForContracts(ctx context.Context, acc *ibc.Wallet, initArgs any) {
	s.SetDepositors(ctx, s.Quasar(), s.LpStrategyContractAddress1, acc.KeyName, initArgs)
	s.SetDepositors(ctx, s.Quasar(), s.LpStrategyContractAddress2, acc.KeyName, initArgs)
	s.SetDepositors(ctx, s.Quasar(), s.LpStrategyContractAddress3, acc.KeyName, initArgs)
}

func (s *Qtransfer) CreatePools(ctx context.Context) {
	// Read the pool details from os file
	poolBz, err := os.ReadFile(osmosisPool1Path)
	s.Require().NoError(err)
	s.CreatePoolsOnOsmosis(ctx, s.Osmosis(), s.E2EBuilder.OsmosisAccounts.Authority.KeyName, poolBz)

	// Read the contract from os file
	poolBz, err = os.ReadFile(osmosisPool2Path)
	s.Require().NoError(err)
	s.CreatePoolsOnOsmosis(ctx, s.Osmosis(), s.E2EBuilder.OsmosisAccounts.Authority.KeyName, poolBz)

	// Read the contract from os file
	poolBz, err = os.ReadFile(osmosisPool3Path)
	s.Require().NoError(err)
	s.CreatePoolsOnOsmosis(ctx, s.Osmosis(), s.E2EBuilder.OsmosisAccounts.Authority.KeyName, poolBz)
}

func (s *Qtransfer) SendTokensToRespectiveAccounts(ctx context.Context) {
	// Send uqsr to all the bond test accounts
	s.SendTokensToOneAddress(ctx, s.Quasar(), s.E2EBuilder.QuasarAccounts.Authority, s.E2EBuilder.QuasarAccounts.BondTest, "10000000uqsr")
	s.SendTokensToOneAddress(ctx, s.Quasar(), s.E2EBuilder.QuasarAccounts.Authority, s.E2EBuilder.QuasarAccounts.BondTest1, "10000000uqsr")
	s.SendTokensToOneAddress(ctx, s.Quasar(), s.E2EBuilder.QuasarAccounts.Authority, s.E2EBuilder.QuasarAccounts.BondTest2, "10000000uqsr")
	s.SendTokensToOneAddress(ctx, s.Quasar(), s.E2EBuilder.QuasarAccounts.Authority, s.E2EBuilder.QuasarAccounts.BondTest3, "10000000uqsr")
	s.SendTokensToOneAddress(ctx, s.Quasar(), s.E2EBuilder.QuasarAccounts.Authority, s.E2EBuilder.QuasarAccounts.BondTest4, "10000000uqsr")
	s.SendTokensToOneAddress(ctx, s.Quasar(), s.E2EBuilder.QuasarAccounts.Authority, s.E2EBuilder.QuasarAccounts.BondTest5, "10000000uqsr")
	s.SendTokensToOneAddress(ctx, s.Quasar(), s.E2EBuilder.QuasarAccounts.Authority, s.E2EBuilder.QuasarAccounts.BondTest6, "10000000uqsr")
	s.SendTokensToOneAddress(ctx, s.Quasar(), s.E2EBuilder.QuasarAccounts.Authority, s.E2EBuilder.QuasarAccounts.BondTest7, "10000000uqsr")

	// Send stake1 and uosmo and usdc to Osmosis authority account
	s.SendTokensToOneAddress(ctx, s.Osmosis(), s.E2EBuilder.OsmosisAccounts.MasterMinter, s.E2EBuilder.OsmosisAccounts.Owner, "10000000uosmo")
	s.SendTokensToOneAddress(ctx, s.Osmosis(), s.E2EBuilder.OsmosisAccounts.MasterMinter, s.E2EBuilder.OsmosisAccounts.NewOwner, "10000000uosmo")
	s.SendTokensToOneAddress(ctx, s.Osmosis(), s.E2EBuilder.OsmosisAccounts.Owner, s.E2EBuilder.OsmosisAccounts.Authority, "10000000000000000stake1")
	s.SendTokensToOneAddress(ctx, s.Osmosis(), s.E2EBuilder.OsmosisAccounts.NewOwner, s.E2EBuilder.OsmosisAccounts.Authority, "10000000000000000usdc")
	s.SendTokensToOneAddress(ctx, s.Osmosis(), s.E2EBuilder.OsmosisAccounts.MasterMinter, s.E2EBuilder.OsmosisAccounts.Authority, "1000000000000000uosmo")

	walletAmount := ibc.WalletAmount{
		Address: s.E2EBuilder.QuasarAccounts.Authority.Address,
		Denom:   "uosmo",
		Amount:  1000000000,
	}
	transfer, err := s.Osmosis().SendIBCTransfer(ctx, s.Osmosis2QuasarTransferChan.ChannelId, s.E2EBuilder.OsmosisAccounts.Authority.KeyName, walletAmount, ibc.TransferOptions{})
	s.Require().NoError(err)
	s.Require().NoError(transfer.Validate())
}
