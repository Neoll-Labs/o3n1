package etherstate

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/nelsonstr/o3n1/ether-state/testutil/sample"
	etherstatesimulation "github.com/nelsonstr/o3n1/ether-state/x/etherstate/simulation"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = etherstatesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgEnableEthAddress = "op_weight_msg_enable_eth_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEnableEthAddress int = 100

	opWeightMsgDisableEthAddress = "op_weight_msg_disable_eth_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDisableEthAddress int = 100

	opWeightMsgSaveEthereumAddressState = "op_weight_msg_save_ethereum_address_state"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSaveEthereumAddressState int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	etherstateGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&etherstateGenesis)
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

	var weightMsgEnableEthAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEnableEthAddress, &weightMsgEnableEthAddress, nil,
		func(_ *rand.Rand) {
			weightMsgEnableEthAddress = defaultWeightMsgEnableEthAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEnableEthAddress,
		etherstatesimulation.SimulateMsgEnableEthAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDisableEthAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDisableEthAddress, &weightMsgDisableEthAddress, nil,
		func(_ *rand.Rand) {
			weightMsgDisableEthAddress = defaultWeightMsgDisableEthAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDisableEthAddress,
		etherstatesimulation.SimulateMsgDisableEthAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSaveEthereumAddressState int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSaveEthereumAddressState, &weightMsgSaveEthereumAddressState, nil,
		func(_ *rand.Rand) {
			weightMsgSaveEthereumAddressState = defaultWeightMsgSaveEthereumAddressState
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSaveEthereumAddressState,
		etherstatesimulation.SimulateMsgSaveEthereumAddressState(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgEnableEthAddress,
			defaultWeightMsgEnableEthAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				etherstatesimulation.SimulateMsgEnableEthAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDisableEthAddress,
			defaultWeightMsgDisableEthAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				etherstatesimulation.SimulateMsgDisableEthAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSaveEthereumAddressState,
			defaultWeightMsgSaveEthereumAddressState,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				etherstatesimulation.SimulateMsgSaveEthereumAddressState(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
