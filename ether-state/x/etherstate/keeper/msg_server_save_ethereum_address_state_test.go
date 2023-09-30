package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/ether-state/testutil/sample"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSaveEthereumAddressState(t *testing.T) {

	address := "0xe8aCaaB95d1102D099F82F03f6106289ee19abA1"

	msgSrvr, context := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	tests := []struct {
		name string
		msg  types.MsgSaveEthereumAddressState
		err  error
	}{

		{
			name: "missing address",
			msg: types.MsgSaveEthereumAddressState{
				Creator: sample.AccAddress(),
				Address: address,
			},
			err: types.ErrEthereumAddressNotRegistered,
		},
		{
			name: "address is disabled",
			msg: types.MsgSaveEthereumAddressState{
				Creator: sample.AccAddress(),
				Address: address,
			},
			err: types.ErrEthereumAddressDisabled,
		},
		{
			name: "address enabled",
			msg: types.MsgSaveEthereumAddressState{
				Creator: sample.AccAddress(),
				Address: address,
			},
		},
		{
			name: "save state",
			msg: types.MsgSaveEthereumAddressState{
				Creator:         sample.AccAddress(),
				Address:         address,
				BlockNumber:     10,
				Nonce:           10,
				StoragePosition: 10,
			},
		},
		{
			name: "invalid nonce",
			msg: types.MsgSaveEthereumAddressState{
				Creator:         sample.AccAddress(),
				Address:         address,
				BlockNumber:     10,
				Nonce:           0,
				StoragePosition: 10,
			},
			err: types.ErrInvalidNonce,
		},
		{
			name: "invalid block number",
			msg: types.MsgSaveEthereumAddressState{
				Creator:         sample.AccAddress(),
				Address:         address,
				BlockNumber:     1,
				Nonce:           0,
				StoragePosition: 10,
			},
			err: types.ErrInvalidBlockNumber,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.name == "address is disabled" {
				msgSrvr.DisableEthAddress(context, &types.MsgDisableEthAddress{
					Creator: sample.AccAddress(),
					Address: address,
				})
			}
			if tt.name == "address enabled" {
				msgSrvr.EnableEthAddress(context, &types.MsgEnableEthAddress{
					Creator: sample.AccAddress(),
					Address: address,
				})
			}

			validMsg := types.NewMsgSaveEthereumAddressState(tt.msg.Creator, tt.msg.Address,
				tt.msg.BlockNumber, tt.msg.Nonce, tt.msg.StoragePosition)

			_, err := msgSrvr.SaveEthereumAddressState(ctx, validMsg)

			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}

}
