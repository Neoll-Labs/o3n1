package types

import (
	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nelsonstr/o3n1/ether-state/helpers"
)

const TypeMsgEnableEthAddress = "enable_eth_address"

var _ sdk.Msg = &MsgEnableEthAddress{}

func NewMsgEnableEthAddress(creator string, address string) *MsgEnableEthAddress {
	return &MsgEnableEthAddress{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgEnableEthAddress) Route() string {
	return RouterKey
}

func (msg *MsgEnableEthAddress) Type() string {
	return TypeMsgEnableEthAddress
}

func (msg *MsgEnableEthAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnableEthAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableEthAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if helpers.IsValidEthAddress(msg.Address) == false {
		return errors.Wrapf(ErrInvalidEthereumAddress, "invalid address (%s)", msg.Address)
	}

	return nil
}
