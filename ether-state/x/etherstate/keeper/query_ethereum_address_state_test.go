package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/nelsonstr/o3n1/ether-state/testutil/keeper"
	"github.com/nelsonstr/o3n1/ether-state/testutil/nullify"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestEthereumAddressStateQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EtherstateKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEthereumAddressState(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetEthereumAddressStateRequest
		response *types.QueryGetEthereumAddressStateResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetEthereumAddressStateRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetEthereumAddressStateResponse{EthereumAddressState: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetEthereumAddressStateRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetEthereumAddressStateResponse{EthereumAddressState: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetEthereumAddressStateRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.EthereumAddressState(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestEthereumAddressStateQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EtherstateKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEthereumAddressState(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllEthereumAddressStateRequest {
		return &types.QueryAllEthereumAddressStateRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EthereumAddressStateAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.EthereumAddressState), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.EthereumAddressState),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EthereumAddressStateAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.EthereumAddressState), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.EthereumAddressState),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.EthereumAddressStateAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.EthereumAddressState),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.EthereumAddressStateAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
