package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/nelsonstr/o3n1/statether/testutil/keeper"
	"github.com/nelsonstr/o3n1/statether/testutil/nullify"
	"github.com/nelsonstr/o3n1/statether/x/statether/keeper"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStatether(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Statether {
	items := make([]types.Statether, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStatether(ctx, items[i])
	}
	return items
}

func TestStatetherGet(t *testing.T) {
	keeper, ctx := keepertest.StatetherKeeper(t)
	items := createNStatether(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStatether(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStatetherRemove(t *testing.T) {
	keeper, ctx := keepertest.StatetherKeeper(t)
	items := createNStatether(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStatether(ctx,
			item.Index,
		)
		_, found := keeper.GetStatether(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStatetherGetAll(t *testing.T) {
	keeper, ctx := keepertest.StatetherKeeper(t)
	items := createNStatether(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStatether(ctx)),
	)
}
