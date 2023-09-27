package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "storepositionether/testutil/keeper"
	"storepositionether/testutil/nullify"
	"storepositionether/x/storepositionether/keeper"
	"storepositionether/x/storepositionether/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNEthAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EthAddress {
	items := make([]types.EthAddress, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEthAddress(ctx, items[i])
	}
	return items
}

func TestEthAddressGet(t *testing.T) {
	keeper, ctx := keepertest.StorepositionetherKeeper(t)
	items := createNEthAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEthAddress(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEthAddressRemove(t *testing.T) {
	keeper, ctx := keepertest.StorepositionetherKeeper(t)
	items := createNEthAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEthAddress(ctx,
			item.Index,
		)
		_, found := keeper.GetEthAddress(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestEthAddressGetAll(t *testing.T) {
	keeper, ctx := keepertest.StorepositionetherKeeper(t)
	items := createNEthAddress(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEthAddress(ctx)),
	)
}
