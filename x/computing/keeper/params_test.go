package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "exabits/testutil/keeper"
	"exabits/x/computing/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ComputingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
