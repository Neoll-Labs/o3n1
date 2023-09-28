package statether_test

import (
	"testing"

	keepertest "github.com/nelsonstr/o3n1/statether/testutil/keeper"
	"github.com/nelsonstr/o3n1/statether/testutil/nullify"
	"github.com/nelsonstr/o3n1/statether/x/statether"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
	"github.com/stretchr/testify/require"
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
