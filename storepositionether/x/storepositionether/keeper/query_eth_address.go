package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"storepositionether/x/storepositionether/types"
)

func (k Keeper) EthAddressAll(goCtx context.Context, req *types.QueryAllEthAddressRequest) (*types.QueryAllEthAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ethAddresss []types.EthAddress
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	ethAddressStore := prefix.NewStore(store, types.KeyPrefix(types.EthAddressKeyPrefix))

	pageRes, err := query.Paginate(ethAddressStore, req.Pagination, func(key []byte, value []byte) error {
		var ethAddress types.EthAddress
		if err := k.cdc.Unmarshal(value, &ethAddress); err != nil {
			return err
		}

		ethAddresss = append(ethAddresss, ethAddress)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEthAddressResponse{EthAddress: ethAddresss, Pagination: pageRes}, nil
}

func (k Keeper) EthAddress(goCtx context.Context, req *types.QueryGetEthAddressRequest) (*types.QueryGetEthAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetEthAddress(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEthAddressResponse{EthAddress: val}, nil
}
