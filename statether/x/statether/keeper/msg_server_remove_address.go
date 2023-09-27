package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"statether/x/statether/types"
)

func (k msgServer) RemoveAddress(goCtx context.Context, msg *types.MsgRemoveAddress) (*types.MsgRemoveAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRemoveAddressResponse{}, nil
}
