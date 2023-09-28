package statether

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/statether/x/statether/keeper"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the statether
	for _, elem := range genState.StatetherList {
		k.SetStatether(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.StatetherList = k.GetAllStatether(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
