package keeper

import (
	"context"
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
)

func (k msgServer) EnableEthAddress(goCtx context.Context, msg *types.MsgEnableEthAddress) (*types.MsgEnableEthAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	newActiveState := true

	eth, exists := k.Keeper.GetEthereumAddress(ctx, msg.Address)
	if exists {
		if newActiveState == eth.Active {
			return nil,
				errors.Wrapf(types.ErrEthereumAddressAlreadyOnRequiredState, "ethereum address (%s) already enabled.", msg.Address)
		}
	}

	k.Keeper.SetEthereumAddress(ctx, types.EthereumAddress{
		Index:  msg.Address,
		Active: newActiveState,
	})

	err := ctx.EventManager().EmitTypedEvent(msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgEnableEthAddressResponse{Success: true}, nil
}
