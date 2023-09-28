package keeper_test

import (
	"testing"

	testkeeper "ether-state/testutil/keeper"
	"ether-state/x/etherstate/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EtherstateKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
