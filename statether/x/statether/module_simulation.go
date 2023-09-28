package statether

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/nelsonstr/o3n1/statether/testutil/sample"
	statethersimulation "github.com/nelsonstr/o3n1/statether/x/statether/simulation"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
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
	opWeightMsgSaveEthaddressState = "op_weight_msg_save_ethaddress_state"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSaveEthaddressState int = 100

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

	var weightMsgSaveEthaddressState int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSaveEthaddressState, &weightMsgSaveEthaddressState, nil,
		func(_ *rand.Rand) {
			weightMsgSaveEthaddressState = defaultWeightMsgSaveEthaddressState
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSaveEthaddressState,
		statethersimulation.SimulateMsgSaveEthaddressState(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSaveEthaddressState,
			defaultWeightMsgSaveEthaddressState,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				statethersimulation.SimulateMsgSaveEthaddressState(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
