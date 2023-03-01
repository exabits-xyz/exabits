package computing

import (
	"exabits/x/computing/keeper"
	"exabits/x/computing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the devices
	for _, elem := range genState.DevicesList {
		k.SetDevices(ctx, elem)
	}

	// Set devices count
	k.SetDevicesCount(ctx, genState.DevicesCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DevicesList = k.GetAllDevices(ctx)
	genesis.DevicesCount = k.GetDevicesCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
