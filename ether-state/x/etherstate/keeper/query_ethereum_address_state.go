package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EthereumAddressStateAll(goCtx context.Context, req *types.QueryAllEthereumAddressStateRequest) (*types.QueryAllEthereumAddressStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ethereumAddressStates []types.EthereumAddressState
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	ethereumAddressStateStore := prefix.NewStore(store, types.KeyPrefix(types.EthereumAddressStateKeyPrefix))

	pageRes, err := query.Paginate(ethereumAddressStateStore, req.Pagination, func(key []byte, value []byte) error {
		var ethereumAddressState types.EthereumAddressState
		if err := k.cdc.Unmarshal(value, &ethereumAddressState); err != nil {
			return err
		}

		ethereumAddressStates = append(ethereumAddressStates, ethereumAddressState)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEthereumAddressStateResponse{EthereumAddressState: ethereumAddressStates, Pagination: pageRes}, nil
}

func (k Keeper) EthereumAddressState(goCtx context.Context, req *types.QueryGetEthereumAddressStateRequest) (*types.QueryGetEthereumAddressStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetEthereumAddressState(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEthereumAddressStateResponse{EthereumAddressState: val}, nil
}
