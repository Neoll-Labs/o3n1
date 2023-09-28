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

func createNEthereumAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EthereumAddress {
	items := make([]types.EthereumAddress, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEthereumAddress(ctx, items[i])
	}
	return items
}

func TestEthereumAddressGet(t *testing.T) {
	keeper, ctx := keepertest.EtherstateKeeper(t)
	items := createNEthereumAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEthereumAddress(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEthereumAddressRemove(t *testing.T) {
	keeper, ctx := keepertest.EtherstateKeeper(t)
	items := createNEthereumAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEthereumAddress(ctx,
			item.Index,
		)
		_, found := keeper.GetEthereumAddress(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestEthereumAddressGetAll(t *testing.T) {
	keeper, ctx := keepertest.EtherstateKeeper(t)
	items := createNEthereumAddress(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEthereumAddress(ctx)),
	)
}
