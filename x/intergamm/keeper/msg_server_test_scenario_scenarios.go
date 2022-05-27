//go:build !prod

package keeper

import (
	"testing"
	"time"

	"github.com/abag/quasarnode/x/intergamm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	gammbalancer "github.com/osmosis-labs/osmosis/v7/x/gamm/pool-models/balancer"
	gammtypes "github.com/osmosis-labs/osmosis/v7/x/gamm/types"
	lockuptypes "github.com/osmosis-labs/osmosis/v7/x/lockup/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	owner        string = "quasar1sqlsc5024sszglyh7pswk5hfpc5xtl77gqjwec"
	connectionId string = "connection-0"
)

var testHooksState map[string]bool

func init() {
	testHooksState = make(map[string]bool)

	scenarios["registerIca"] = testRegisterIca
	scenarios["createPool"] = testCreatePool
	scenarios["createPoolChecks"] = testCreatePoolChecks
	scenarios["createPoolTimeout"] = testCreatePoolTimeout
	scenarios["createPoolTimeoutChecks"] = testCreatePoolTimeoutChecks
	scenarios["joinPool"] = testJoinPool
	scenarios["joinPoolChecks"] = testJoinPoolChecks
	scenarios["joinPoolTimeout"] = testJoinPoolTimeout
	scenarios["joinPoolTimeoutChecks"] = testJoinPoolTimeoutChecks
	scenarios["joinPoolSingleDenom"] = testJoinPoolSingleDenom
	scenarios["joinPoolSingleDenomChecks"] = testJoinPoolSingleDenomChecks
	scenarios["joinPoolSingleDenomTimeout"] = testJoinPoolSingleDenomTimeout
	scenarios["joinPoolSingleDenomTimeoutChecks"] = testJoinPoolSingleDenomTimeoutChecks
	scenarios["exitPool"] = testExitPool
	scenarios["exitPoolChecks"] = testExitPoolChecks
	scenarios["exitPoolTimeout"] = testExitPoolTimeout
	scenarios["exitPoolTimeoutChecks"] = testExitPoolTimeoutChecks
	scenarios["lockTokens"] = testLockTokens
	scenarios["lockTokensChecks"] = testLockTokensChecks
	scenarios["lockTokensTimeout"] = testLockTokensTimeout
	scenarios["lockTokensTimeoutChecks"] = testLockTokensTimeoutChecks
}

func createTestPoolParams() *gammbalancer.PoolParams {
	swapFee, err := sdk.NewDecFromStr("0.01")
	if err != nil {
		panic(err)
	}

	exitFee, err := sdk.NewDecFromStr("0.01")
	if err != nil {
		panic(err)
	}

	return &gammbalancer.PoolParams{
		SwapFee: swapFee,
		ExitFee: exitFee,
	}
}

func createTestPoolAssets() []gammtypes.PoolAsset {
	return []gammtypes.PoolAsset{
		{
			Weight: sdk.NewInt(100),
			Token:  sdk.NewCoin("uatom", sdk.NewInt(10000)),
		},
		{
			Weight: sdk.NewInt(100),
			Token:  sdk.NewCoin("uosmo", sdk.NewInt(10000)),
		},
	}
}

func joinPoolTestCoins() []sdk.Coin {
	return []sdk.Coin{
		sdk.NewCoin("uatom", sdk.NewInt(1000)),
		sdk.NewCoin("uosmo", sdk.NewInt(1000)),
	}
}

func joinPoolSingleDenomTestCoin() sdk.Coin {
	return sdk.NewCoin("uatom", sdk.NewInt(1000))
}

func lockTokensTestCoins() []sdk.Coin {
	return []sdk.Coin{
		sdk.NewCoin("gamm/pool/1", sdk.NewInt(42000)),
	}
}

func ensureIcaRegistered(ctx sdk.Context, k *Keeper, owner string, connectionId string) error {
	var err error

	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "could not generate port for address: %s", err)
	}

	_, found := k.icaControllerKeeper.GetOpenActiveChannel(ctx, connectionId, portID)
	if !found {
		err = k.RegisterInterchainAccount(ctx, connectionId, owner)
		if err != nil {
			return err
		}
	}

	return nil
}

func testRegisterIca(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		err := ensureIcaRegistered(ctx, k, owner, connectionId)
		require.NoError(t, err)
	}
}

func testCreatePool(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		var err error

		// Setup hook
		k.Hooks.Osmosis.AddHooksAckMsgCreateBalancerPool(func(sdk.Context, types.AckExchange[*gammbalancer.MsgCreateBalancerPool, *gammbalancer.MsgCreateBalancerPoolResponse]) {
			testHooksState["testCreatePool_hook"] = true
		})

		timestamp := uint64(99999999999999)
		futureGovernor := "168h"
		poolParams := createTestPoolParams()
		poolAssets := createTestPoolAssets()

		err = k.TransmitIbcCreatePool(
			ctx,
			owner,
			connectionId,
			timestamp,
			poolParams,
			poolAssets,
			futureGovernor,
		)
		require.NoError(t, err)
	}
}

func testCreatePoolChecks(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		require.True(t, testHooksState["testCreatePool_hook"])
	}
}

func testCreatePoolTimeout(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		var err error

		// Setup hook
		k.Hooks.Osmosis.AddHooksTimeoutMsgCreateBalancerPool(func(sdk.Context, types.TimeoutExchange[*gammbalancer.MsgCreateBalancerPool]) {
			testHooksState["testCreatePoolTimeout_hook"] = true
		})

		timestamp := uint64(99999999999999)
		futureGovernor := "168h"

		poolParams := createTestPoolParams()
		poolAssets := createTestPoolAssets()

		// Replace timeout to trigger timeout hooks
		tmpDefaultSendTxRelativeTimeoutTimestamp := DefaultSendTxRelativeTimeoutTimestamp
		DefaultSendTxRelativeTimeoutTimestamp = uint64((time.Duration(200) * time.Millisecond).Nanoseconds())
		defer func() {
			DefaultSendTxRelativeTimeoutTimestamp = tmpDefaultSendTxRelativeTimeoutTimestamp
		}()

		err = k.TransmitIbcCreatePool(
			ctx,
			owner,
			connectionId,
			timestamp,
			poolParams,
			poolAssets,
			futureGovernor,
		)
		require.NoError(t, err)
	}
}

func testCreatePoolTimeoutChecks(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		require.True(t, testHooksState["testCreatePoolTimeout_hook"])
	}
}

func testJoinPool(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		var err error

		// Setup hook
		k.Hooks.Osmosis.AddHooksAckMsgJoinPool(func(sdk.Context, types.AckExchange[*gammtypes.MsgJoinPool, *gammtypes.MsgJoinPoolResponse]) {
			testHooksState["testJoinPool_hook"] = true
		})

		poolId := uint64(1)
		timestamp := uint64(99999999999999)
		testCoins := joinPoolTestCoins()
		shares, ok := sdk.NewIntFromString("1000000000000000000")
		require.True(t, ok)

		err = k.TransmitIbcJoinPool(
			ctx,
			owner,
			connectionId,
			timestamp,
			poolId,
			shares,
			testCoins,
		)
		require.NoError(t, err)
	}
}

func testJoinPoolChecks(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		require.True(t, testHooksState["testJoinPool_hook"])
	}
}

func testJoinPoolTimeout(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		var err error

		// Setup hook
		k.Hooks.Osmosis.AddHooksTimeoutMsgJoinPool(func(sdk.Context, types.TimeoutExchange[*gammtypes.MsgJoinPool]) {
			testHooksState["testJoinPoolTimeout_hook"] = true
		})

		poolId := uint64(1)
		timestamp := uint64(99999999999999)
		testCoins := joinPoolTestCoins()
		shares, ok := sdk.NewIntFromString("1000000000000000000")
		require.True(t, ok)

		// Replace timeout to trigger timeout hooks
		tmpDefaultSendTxRelativeTimeoutTimestamp := DefaultSendTxRelativeTimeoutTimestamp
		DefaultSendTxRelativeTimeoutTimestamp = uint64((time.Duration(200) * time.Millisecond).Nanoseconds())
		defer func() {
			DefaultSendTxRelativeTimeoutTimestamp = tmpDefaultSendTxRelativeTimeoutTimestamp
		}()

		err = k.TransmitIbcJoinPool(
			ctx,
			owner,
			connectionId,
			timestamp,
			poolId,
			shares,
			testCoins,
		)
		require.NoError(t, err)
	}
}

func testJoinPoolTimeoutChecks(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		require.True(t, testHooksState["testJoinPoolTimeout_hook"])
	}
}

func testJoinPoolSingleDenom(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		var err error

		// Setup hook
		k.Hooks.Osmosis.AddHooksAckMsgJoinPoolSingleDenom(func(sdk.Context, types.AckExchange[*gammtypes.MsgJoinSwapExternAmountIn, *gammtypes.MsgJoinSwapExternAmountInResponse]) {
			testHooksState["testJoinPoolSingleDenom_hook"] = true
		})

		poolId := uint64(1)
		timestamp := uint64(99999999999999)
		testCoin := joinPoolSingleDenomTestCoin()
		shares, ok := sdk.NewIntFromString("500000000000000000")
		require.True(t, ok)

		err = k.TransmitIbcJoinSwapExternAmountIn(
			ctx,
			owner,
			connectionId,
			timestamp,
			poolId,
			testCoin,
			shares,
		)
		require.NoError(t, err)
	}
}

func testJoinPoolSingleDenomChecks(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		require.True(t, testHooksState["testJoinPoolSingleDenom_hook"])
	}
}

func testJoinPoolSingleDenomTimeout(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		var err error

		// Setup hook
		k.Hooks.Osmosis.AddHooksTimeoutMsgJoinPoolSingleDenom(func(sdk.Context, types.TimeoutExchange[*gammtypes.MsgJoinSwapExternAmountIn]) {
			testHooksState["testJoinPoolSingleDenomTimeout_hook"] = true
		})

		poolId := uint64(1)
		timestamp := uint64(99999999999999)
		testCoin := joinPoolSingleDenomTestCoin()
		shares, ok := sdk.NewIntFromString("500000000000000000")
		require.True(t, ok)

		// Replace timeout to trigger timeout hooks
		tmpDefaultSendTxRelativeTimeoutTimestamp := DefaultSendTxRelativeTimeoutTimestamp
		DefaultSendTxRelativeTimeoutTimestamp = uint64((time.Duration(200) * time.Millisecond).Nanoseconds())
		defer func() {
			DefaultSendTxRelativeTimeoutTimestamp = tmpDefaultSendTxRelativeTimeoutTimestamp
		}()

		err = k.TransmitIbcJoinSwapExternAmountIn(
			ctx,
			owner,
			connectionId,
			timestamp,
			poolId,
			testCoin,
			shares,
		)
		require.NoError(t, err)
	}
}

func testJoinPoolSingleDenomTimeoutChecks(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		require.True(t, testHooksState["testJoinPoolSingleDenomTimeout_hook"])
	}
}

func testExitPool(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		var err error

		// Setup hook
		k.Hooks.Osmosis.AddHooksAckMsgExitPool(func(sdk.Context, types.AckExchange[*gammtypes.MsgExitPool, *gammtypes.MsgExitPoolResponse]) {
			testHooksState["testExitPool_hook"] = true
		})

		poolId := uint64(1)
		timestamp := uint64(99999999999999)
		testCoins := joinPoolTestCoins()
		shares, ok := sdk.NewIntFromString("1000000000000000000")
		require.True(t, ok)

		err = k.TransmitIbcExitPool(
			ctx,
			owner,
			connectionId,
			timestamp,
			poolId,
			shares,
			testCoins,
		)
		require.NoError(t, err)
	}
}

func testExitPoolChecks(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		require.True(t, testHooksState["testExitPool_hook"])
	}
}

func testExitPoolTimeout(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		var err error

		// Setup hook
		k.Hooks.Osmosis.AddHooksTimeoutMsgExitPool(func(sdk.Context, types.TimeoutExchange[*gammtypes.MsgExitPool]) {
			testHooksState["testExitPoolTimeout_hook"] = true
		})

		poolId := uint64(1)
		timestamp := uint64(99999999999999)
		testCoins := joinPoolTestCoins()
		shares, ok := sdk.NewIntFromString("1000000000000000000")
		require.True(t, ok)

		// Replace timeout to trigger timeout hooks
		tmpDefaultSendTxRelativeTimeoutTimestamp := DefaultSendTxRelativeTimeoutTimestamp
		DefaultSendTxRelativeTimeoutTimestamp = uint64((time.Duration(200) * time.Millisecond).Nanoseconds())
		defer func() {
			DefaultSendTxRelativeTimeoutTimestamp = tmpDefaultSendTxRelativeTimeoutTimestamp
		}()

		err = k.TransmitIbcExitPool(
			ctx,
			owner,
			connectionId,
			timestamp,
			poolId,
			shares,
			testCoins,
		)
		require.NoError(t, err)
	}
}

func testExitPoolTimeoutChecks(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		require.True(t, testHooksState["testExitPoolTimeout_hook"])
	}
}

func testLockTokens(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		var err error

		// Setup hook
		k.Hooks.Osmosis.AddHooksAckMsgLockTokens(func(sdk.Context, types.AckExchange[*lockuptypes.MsgLockTokens, *lockuptypes.MsgLockTokensResponse]) {
			testHooksState["testLockTokens_hook"] = true
		})

		timestamp := uint64(99999999999999)
		lockupPeriod := 1 * time.Hour
		testCoins := lockTokensTestCoins()

		err = k.TransmitIbcLockTokens(
			ctx,
			owner,
			connectionId,
			timestamp,
			lockupPeriod,
			testCoins,
		)
		require.NoError(t, err)
	}
}

func testLockTokensChecks(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		require.True(t, testHooksState["testLockTokens_hook"])
	}
}

func testLockTokensTimeout(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		var err error

		// Setup hook
		k.Hooks.Osmosis.AddHooksTimeoutMsgLockTokens(func(sdk.Context, types.TimeoutExchange[*lockuptypes.MsgLockTokens]) {
			testHooksState["testLockTokensTimeout_hook"] = true
		})

		timestamp := uint64(99999999999999)
		lockupPeriod := 1 * time.Hour
		testCoins := lockTokensTestCoins()

		// Replace timeout to trigger timeout hooks
		tmpDefaultSendTxRelativeTimeoutTimestamp := DefaultSendTxRelativeTimeoutTimestamp
		DefaultSendTxRelativeTimeoutTimestamp = uint64((time.Duration(200) * time.Millisecond).Nanoseconds())
		defer func() {
			DefaultSendTxRelativeTimeoutTimestamp = tmpDefaultSendTxRelativeTimeoutTimestamp
		}()

		err = k.TransmitIbcLockTokens(
			ctx,
			owner,
			connectionId,
			timestamp,
			lockupPeriod,
			testCoins,
		)
		require.NoError(t, err)
	}
}

func testLockTokensTimeoutChecks(ctx sdk.Context, k *Keeper) func(t *testing.T) {
	return func(t *testing.T) {
		require.True(t, testHooksState["testLockTokensTimeout_hook"])
	}
}
