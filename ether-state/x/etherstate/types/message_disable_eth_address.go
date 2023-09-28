package types

import (
	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nelsonstr/o3n1/ether-state/helpers"
)

const TypeMsgDisableEthAddress = "disable_eth_address"

var _ sdk.Msg = &MsgDisableEthAddress{}

func NewMsgDisableEthAddress(creator string, address string) *MsgDisableEthAddress {
	return &MsgDisableEthAddress{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgDisableEthAddress) Route() string {
	return RouterKey
}

func (msg *MsgDisableEthAddress) Type() string {
	return TypeMsgDisableEthAddress
}

func (msg *MsgDisableEthAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDisableEthAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDisableEthAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if helpers.IsValidEthAddress(msg.Address) == false {
		return errors.Wrapf(ErrInvalidEthereumAddress, "invalid address (%s)", msg.Address)
	}

	return nil
}
