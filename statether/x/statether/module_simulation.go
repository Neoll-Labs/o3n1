package statether

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"statether/testutil/sample"
	statethersimulation "statether/x/statether/simulation"
	"statether/x/statether/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = statethersimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgAddAddress = "op_weight_msg_add_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddAddress int = 100

	opWeightMsgSaveEthaddressStoragePosition = "op_weight_msg_save_ethaddress_storage_position"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSaveEthaddressStoragePosition int = 100

	opWeightMsgRemoveAddress = "op_weight_msg_remove_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveAddress int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	statetherGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&statetherGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddAddress, &weightMsgAddAddress, nil,
		func(_ *rand.Rand) {
			weightMsgAddAddress = defaultWeightMsgAddAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddAddress,
		statethersimulation.SimulateMsgAddAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSaveEthaddressStoragePosition int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSaveEthaddressStoragePosition, &weightMsgSaveEthaddressStoragePosition, nil,
		func(_ *rand.Rand) {
			weightMsgSaveEthaddressStoragePosition = defaultWeightMsgSaveEthaddressStoragePosition
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSaveEthaddressStoragePosition,
		statethersimulation.SimulateMsgSaveEthaddressStoragePosition(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveAddress, &weightMsgRemoveAddress, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveAddress = defaultWeightMsgRemoveAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveAddress,
		statethersimulation.SimulateMsgRemoveAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddAddress,
			defaultWeightMsgAddAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				statethersimulation.SimulateMsgAddAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSaveEthaddressStoragePosition,
			defaultWeightMsgSaveEthaddressStoragePosition,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				statethersimulation.SimulateMsgSaveEthaddressStoragePosition(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRemoveAddress,
			defaultWeightMsgRemoveAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				statethersimulation.SimulateMsgRemoveAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
