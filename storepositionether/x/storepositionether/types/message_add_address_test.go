package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"storepositionether/testutil/sample"
)

func TestMsgAddAddress_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddAddress
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgAddAddress{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid creator address",
			msg: MsgAddAddress{
				Creator:    sample.AccAddress(),
				EthAddress: "0xe8aCaaB95d1102D099F82F03f6106289ee19abA8",
			},
		}, {
			name: "invalid ether address",
			msg: MsgAddAddress{
				Creator:    sample.AccAddress(),
				EthAddress: "e8aCaaB95d1102D099F82F03f6106289ee19abA8",
			},
			err: ErrInvalidEthAddress,
		}, {
			name: "valid ether address",
			msg: MsgAddAddress{
				Creator:    sample.AccAddress(),
				EthAddress: "0xe8aCaaB95d1102D099F82F03f6106289ee19abA8",
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
