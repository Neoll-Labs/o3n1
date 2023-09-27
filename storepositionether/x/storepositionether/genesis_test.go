package storepositionether_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "storepositionether/testutil/keeper"
	"storepositionether/testutil/nullify"
	"storepositionether/x/storepositionether"
	"storepositionether/x/storepositionether/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StorepositionetherKeeper(t)
	storepositionether.InitGenesis(ctx, *k, genesisState)
	got := storepositionether.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
