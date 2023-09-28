package etherstate_test

import (
	"testing"

	keepertest "ether-state/testutil/keeper"
	"ether-state/testutil/nullify"
	"ether-state/x/etherstate"
	"ether-state/x/etherstate/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EtherstateKeeper(t)
	etherstate.InitGenesis(ctx, *k, genesisState)
	got := etherstate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
