package keeper

import (
	"context"
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
)

func (k msgServer) SaveEthereumAddressState(goCtx context.Context, msg *types.MsgSaveEthereumAddressState) (*types.MsgSaveEthereumAddressStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	eth, exists := k.Keeper.GetEthereumAddress(ctx, msg.Address)
	if exists {
		if !eth.Active {
			return &types.MsgSaveEthereumAddressStateResponse{Success: false},
				errors.Wrapf(types.ErrEthereumAddressDisabled, "ethereum address disabled (%v)", msg.Address)
		}
	} else {
		return &types.MsgSaveEthereumAddressStateResponse{Success: false},
			errors.Wrapf(types.ErrEthereumAddressNotRegistered, "ethereum address not registered (%v)", msg.Address)
	}

	previousState, exists := k.Keeper.GetEthereumAddressState(ctx, msg.Address)
	if exists {
		if previousState.Block > msg.BlockNumber {
			return &types.MsgSaveEthereumAddressStateResponse{Success: false},
				errors.Wrapf(types.ErrInvalidBlockNumber, "invalid block number (%v)", msg.BlockNumber)
		}

		if previousState.Nonce > msg.Nonce {
			return &types.MsgSaveEthereumAddressStateResponse{Success: false},
				errors.Wrapf(types.ErrInvalidNonce, "invalid nonce (%v)", msg.Nonce)
		}
	}

	k.Keeper.SetEthereumAddressState(ctx, types.EthereumAddressState{
		Index: msg.Address,
		State: msg.StoragePosition,
		Block: msg.BlockNumber,
		Nonce: msg.Nonce,
	})

	return &types.MsgSaveEthereumAddressStateResponse{
		Success: true,
	}, nil
}
