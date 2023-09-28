package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
)

func (k msgServer) GetEthaddressState(goCtx context.Context, msg *types.MsgGetEthaddressState) (*types.MsgGetEthaddressStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	fmt.Println("GetEthaddressState--------------------------------------------------------")

	k.Keeper.SetStatether(ctx, types.Statether{
		//Index:           "",
		EthAddress:      "",
		Block:           0,
		Nonce:           0,
		StoragePosition: 0,
		//Active:          false,
	})
	return &types.MsgGetEthaddressStateResponse{}, nil
}
