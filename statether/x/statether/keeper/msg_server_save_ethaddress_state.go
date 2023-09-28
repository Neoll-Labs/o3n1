package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
)

func (k msgServer) SaveEthaddressState(goCtx context.Context, msg *types.MsgSaveEthaddressState) (*types.MsgSaveEthaddressStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	k.Keeper.SetStatether(ctx, types.Statether{
		Index:           msg.EthAddress,
		EthAddress:      msg.EthAddress,
		Block:           msg.Block,
		Nonce:           msg.Nonce,
		StoragePosition: msg.StoragePosition,
	})

	return &types.MsgSaveEthaddressStateResponse{}, nil
}
