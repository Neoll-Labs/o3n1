package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/nelsonstr/o3n1/statether/testutil/keeper"
	"github.com/nelsonstr/o3n1/statether/testutil/nullify"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestStatetherQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.StatetherKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStatether(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetStatetherRequest
		response *types.QueryGetStatetherResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetStatetherRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetStatetherResponse{Statether: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetStatetherRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetStatetherResponse{Statether: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetStatetherRequest{
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
			response, err := keeper.Statether(wctx, tc.request)
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

func TestStatetherQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.StatetherKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStatether(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllStatetherRequest {
		return &types.QueryAllStatetherRequest{
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
			resp, err := keeper.StatetherAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Statether), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Statether),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StatetherAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Statether), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Statether),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.StatetherAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Statether),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.StatetherAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
