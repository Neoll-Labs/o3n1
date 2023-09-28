package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
)

func (k msgServer) GetEthaddressState(goCtx context.Context, msg *types.MsgGetEthaddressState) (*types.MsgGetEthaddressStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgGetEthaddressStateResponse{}, nil
}
