package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"statether/x/statether/types"
)

func (k msgServer) AddAddress(goCtx context.Context, msg *types.MsgAddAddress) (*types.MsgAddAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddAddressResponse{}, nil
}
