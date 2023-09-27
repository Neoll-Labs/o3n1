package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"storepositionether/x/storepositionether/types"
)

func (k msgServer) AddAddress(goCtx context.Context, msg *types.MsgAddAddress) (*types.MsgAddAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddAddressResponse{}, nil
}
