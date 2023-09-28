package keeper_test

import (
	"testing"

	testkeeper "github.com/nelsonstr/o3n1/statether/testutil/keeper"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.StatetherKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
