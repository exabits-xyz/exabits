package keeper_test

import (
	"testing"

	keepertest "exabits/testutil/keeper"
	"exabits/testutil/nullify"
	"exabits/x/computing/keeper"
	"exabits/x/computing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNDevices(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Devices {
	items := make([]types.Devices, n)
	for i := range items {
		items[i].Id = keeper.AppendDevices(ctx, items[i])
	}
	return items
}

func TestDevicesGet(t *testing.T) {
	keeper, ctx := keepertest.ComputingKeeper(t)
	items := createNDevices(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetDevices(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestDevicesRemove(t *testing.T) {
	keeper, ctx := keepertest.ComputingKeeper(t)
	items := createNDevices(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDevices(ctx, item.Id)
		_, found := keeper.GetDevices(ctx, item.Id)
		require.False(t, found)
	}
}

func TestDevicesGetAll(t *testing.T) {
	keeper, ctx := keepertest.ComputingKeeper(t)
	items := createNDevices(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDevices(ctx)),
	)
}

func TestDevicesCount(t *testing.T) {
	keeper, ctx := keepertest.ComputingKeeper(t)
	items := createNDevices(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetDevicesCount(ctx))
}
