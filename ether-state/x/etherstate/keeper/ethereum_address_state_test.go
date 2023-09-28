package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/nelsonstr/o3n1/ether-state/testutil/keeper"
	"github.com/nelsonstr/o3n1/ether-state/testutil/nullify"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/keeper"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNEthereumAddressState(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EthereumAddressState {
	items := make([]types.EthereumAddressState, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEthereumAddressState(ctx, items[i])
	}
	return items
}

func TestEthereumAddressStateGet(t *testing.T) {
	keeper, ctx := keepertest.EtherstateKeeper(t)
	items := createNEthereumAddressState(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEthereumAddressState(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEthereumAddressStateRemove(t *testing.T) {
	keeper, ctx := keepertest.EtherstateKeeper(t)
	items := createNEthereumAddressState(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEthereumAddressState(ctx,
			item.Index,
		)
		_, found := keeper.GetEthereumAddressState(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestEthereumAddressStateGetAll(t *testing.T) {
	keeper, ctx := keepertest.EtherstateKeeper(t)
	items := createNEthereumAddressState(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEthereumAddressState(ctx)),
	)
}
