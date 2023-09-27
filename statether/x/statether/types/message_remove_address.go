package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveAddress = "remove_address"

var _ sdk.Msg = &MsgRemoveAddress{}

func NewMsgRemoveAddress(creator string, address string) *MsgRemoveAddress {
	return &MsgRemoveAddress{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgRemoveAddress) Route() string {
	return RouterKey
}

func (msg *MsgRemoveAddress) Type() string {
	return TypeMsgRemoveAddress
}

func (msg *MsgRemoveAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
