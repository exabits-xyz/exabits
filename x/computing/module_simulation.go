package computing

import (
	"math/rand"

	"exabits/testutil/sample"
	computingsimulation "exabits/x/computing/simulation"
	"exabits/x/computing/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = computingsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateDevices = "op_weight_msg_devices"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDevices int = 100

	opWeightMsgUpdateDevices = "op_weight_msg_devices"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDevices int = 100

	opWeightMsgDeleteDevices = "op_weight_msg_devices"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDevices int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	computingGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DevicesList: []types.Devices{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		DevicesCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&computingGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDevices int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDevices, &weightMsgCreateDevices, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDevices = defaultWeightMsgCreateDevices
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDevices,
		computingsimulation.SimulateMsgCreateDevices(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDevices int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDevices, &weightMsgUpdateDevices, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDevices = defaultWeightMsgUpdateDevices
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDevices,
		computingsimulation.SimulateMsgUpdateDevices(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDevices int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDevices, &weightMsgDeleteDevices, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDevices = defaultWeightMsgDeleteDevices
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDevices,
		computingsimulation.SimulateMsgDeleteDevices(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
