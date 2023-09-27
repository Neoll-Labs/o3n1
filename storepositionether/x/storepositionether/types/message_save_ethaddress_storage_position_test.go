package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"storepositionether/testutil/sample"
)

func TestMsgSaveEthaddressStoragePosition_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSaveEthaddressStoragePosition
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSaveEthaddressStoragePosition{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid eth address",
			msg: MsgSaveEthaddressStoragePosition{
				Creator: sample.AccAddress(),
				Data: &EthaddressStoragePosition{
					EthAddress: "e8aCaaB95d1102D099F82F03f6106289ee19abA8",
				},
			},
			err: ErrInvalidEthAddress,
		}, {
			name: "valid request",
			msg: MsgSaveEthaddressStoragePosition{
				Creator: sample.AccAddress(),
				Data: &EthaddressStoragePosition{
					EthAddress:      "0xe8aCaaB95d1102D099F82F03f6106289ee19abA8",
					Nonce:           0,
					Block:           0,
					StoragePosition: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
