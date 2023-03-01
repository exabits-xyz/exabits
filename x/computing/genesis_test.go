package computing_test

import (
	"testing"

	keepertest "exabits/testutil/keeper"
	"exabits/testutil/nullify"
	"exabits/x/computing"
	"exabits/x/computing/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DevicesList: []types.Devices{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		DevicesCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ComputingKeeper(t)
	computing.InitGenesis(ctx, *k, genesisState)
	got := computing.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DevicesList, got.DevicesList)
	require.Equal(t, genesisState.DevicesCount, got.DevicesCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
