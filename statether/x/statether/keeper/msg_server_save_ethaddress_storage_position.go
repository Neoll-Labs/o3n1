package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"statether/x/statether/types"
)

func (k msgServer) SaveEthaddressStoragePosition(goCtx context.Context, msg *types.MsgSaveEthaddressStoragePosition) (*types.MsgSaveEthaddressStoragePositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSaveEthaddressStoragePositionResponse{}, nil
}
