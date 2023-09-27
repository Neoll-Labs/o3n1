package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddAddress = "add_address"

var _ sdk.Msg = &MsgAddAddress{}

func NewMsgAddAddress(creator string, address string) *MsgAddAddress {
	return &MsgAddAddress{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgAddAddress) Route() string {
	return RouterKey
}

func (msg *MsgAddAddress) Type() string {
	return TypeMsgAddAddress
}

func (msg *MsgAddAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
