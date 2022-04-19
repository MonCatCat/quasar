package keeper_test

import (
	"testing"

	keepertest "github.com/abag/quasarnode/testutil/keeper"
	"github.com/abag/quasarnode/x/intergamm/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	ctx, keeper := keepertest.NewTestSetup(t).GetInterGammKeeper()
	params := types.DefaultParams()

	keeper.SetParams(ctx, params)

	require.EqualValues(t, params, keeper.GetParams(ctx))
}
