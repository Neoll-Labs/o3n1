package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
)

func (k msgServer) DisableEthAddress(goCtx context.Context, msg *types.MsgDisableEthAddress) (*types.MsgDisableEthAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: ✓ Handling the message
	_ = ctx

	k.Keeper.SetEthereumAddress(ctx, types.EthereumAddress{
		Index:  msg.Address,
		Active: false,
	})

	return &types.MsgDisableEthAddressResponse{Success: true}, nil
}
