package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "storepositionether/testutil/keeper"
	"storepositionether/x/storepositionether/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.StorepositionetherKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
