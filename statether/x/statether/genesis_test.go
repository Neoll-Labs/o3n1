package statether_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "statether/testutil/keeper"
	"statether/testutil/nullify"
	"statether/x/statether"
	"statether/x/statether/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StatetherKeeper(t)
	statether.InitGenesis(ctx, *k, genesisState)
	got := statether.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
