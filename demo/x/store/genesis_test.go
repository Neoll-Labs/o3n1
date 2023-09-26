package store_test

import (
	"testing"

	keepertest "demo/testutil/keeper"
	"demo/testutil/nullify"
	"demo/x/store"
	"demo/x/store/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StoreKeeper(t)
	store.InitGenesis(ctx, *k, genesisState)
	got := store.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
