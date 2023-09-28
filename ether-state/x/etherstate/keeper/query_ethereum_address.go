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

func (k Keeper) EthereumAddressAll(goCtx context.Context, req *types.QueryAllEthereumAddressRequest) (*types.QueryAllEthereumAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ethereumAddresss []types.EthereumAddress
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	ethereumAddressStore := prefix.NewStore(store, types.KeyPrefix(types.EthereumAddressKeyPrefix))

	pageRes, err := query.Paginate(ethereumAddressStore, req.Pagination, func(key []byte, value []byte) error {
		var ethereumAddress types.EthereumAddress
		if err := k.cdc.Unmarshal(value, &ethereumAddress); err != nil {
			return err
		}

		ethereumAddresss = append(ethereumAddresss, ethereumAddress)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEthereumAddressResponse{EthereumAddress: ethereumAddresss, Pagination: pageRes}, nil
}

func (k Keeper) EthereumAddress(goCtx context.Context, req *types.QueryGetEthereumAddressRequest) (*types.QueryGetEthereumAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetEthereumAddress(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEthereumAddressResponse{EthereumAddress: val}, nil
}
