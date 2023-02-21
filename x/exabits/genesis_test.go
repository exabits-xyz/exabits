package exabits_test

import (
	"testing"

	keepertest "exabits/testutil/keeper"
	"exabits/testutil/nullify"
	"exabits/x/exabits"
	"exabits/x/exabits/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExabitsKeeper(t)
	exabits.InitGenesis(ctx, *k, genesisState)
	got := exabits.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
