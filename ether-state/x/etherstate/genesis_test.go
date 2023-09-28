package etherstate_test

import (
	"testing"

	keepertest "github.com/nelsonstr/o3n1/ether-state/testutil/keeper"
	"github.com/nelsonstr/o3n1/ether-state/testutil/nullify"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EthereumAddressList: []types.EthereumAddress{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EtherstateKeeper(t)
	etherstate.InitGenesis(ctx, *k, genesisState)
	got := etherstate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EthereumAddressList, got.EthereumAddressList)
	// this line is used by starport scaffolding # genesis/test/assert
}
