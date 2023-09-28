package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nelsonstr/o3n1/ether-state/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgDisableEthAddress_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDisableEthAddress
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDisableEthAddress{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid eth address",
			msg: MsgDisableEthAddress{
				Creator: sample.AccAddress(),
				Address: "invalid_address",
			},
			err: ErrInvalidEthereumAddress,
		},
		{
			name: "valid addresses",
			msg: MsgDisableEthAddress{
				Creator: sample.AccAddress(),
				Address: "0xAAAcE30a1329520a24c6B569ea57989ffa90086A",
			},
			err: ErrInvalidEthereumAddress,
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
