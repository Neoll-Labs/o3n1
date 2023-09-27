package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"statether/x/statether/types"
)

func (k Keeper) GetEthaddressStoragePosition(goCtx context.Context, req *types.QueryGetEthaddressStoragePositionRequest) (*types.QueryGetEthaddressStoragePositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetEthaddressStoragePositionResponse{}, nil
}
