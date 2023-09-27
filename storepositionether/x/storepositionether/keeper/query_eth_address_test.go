package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "storepositionether/testutil/keeper"
	"storepositionether/testutil/nullify"
	"storepositionether/x/storepositionether/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestEthAddressQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.StorepositionetherKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEthAddress(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetEthAddressRequest
		response *types.QueryGetEthAddressResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetEthAddressRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetEthAddressResponse{EthAddress: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetEthAddressRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetEthAddressResponse{EthAddress: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetEthAddressRequest{
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
			response, err := keeper.EthAddress(wctx, tc.request)
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

func TestEthAddressQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.StorepositionetherKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEthAddress(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllEthAddressRequest {
		return &types.QueryAllEthAddressRequest{
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
			resp, err := keeper.EthAddressAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.EthAddress), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.EthAddress),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EthAddressAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.EthAddress), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.EthAddress),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.EthAddressAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.EthAddress),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.EthAddressAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
