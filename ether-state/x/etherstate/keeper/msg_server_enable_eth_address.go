package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
)

func (k msgServer) EnableEthAddress(goCtx context.Context, msg *types.MsgEnableEthAddress) (*types.MsgEnableEthAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: ✓ Handling the message
	_ = ctx

	k.Keeper.SetEthereumAddress(ctx, types.EthereumAddress{
		Index:  msg.Address,
		Active: true,
	})

	return &types.MsgEnableEthAddressResponse{Success: true}, nil
}
