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

		EthAddressList: []types.EthAddress{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StorepositionetherKeeper(t)
	storepositionether.InitGenesis(ctx, *k, genesisState)
	got := storepositionether.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EthAddressList, got.EthAddressList)
	// this line is used by starport scaffolding # genesis/test/assert
}
