package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "statether/testutil/keeper"
	"statether/x/statether/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.StatetherKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
