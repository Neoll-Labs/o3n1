package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StatetherAll(goCtx context.Context, req *types.QueryAllStatetherRequest) (*types.QueryAllStatetherResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var statethers []types.Statether
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	statetherStore := prefix.NewStore(store, types.KeyPrefix(types.StatetherKeyPrefix))

	pageRes, err := query.Paginate(statetherStore, req.Pagination, func(key []byte, value []byte) error {
		var statether types.Statether
		if err := k.cdc.Unmarshal(value, &statether); err != nil {
			return err
		}

		statethers = append(statethers, statether)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStatetherResponse{Statether: statethers, Pagination: pageRes}, nil
}

func (k Keeper) Statether(goCtx context.Context, req *types.QueryGetStatetherRequest) (*types.QueryGetStatetherResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetStatether(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStatetherResponse{Statether: val}, nil
}
