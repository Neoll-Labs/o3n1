package types

import (
	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"storepositionether/helpers"
)

const TypeMsgAddAddress = "add_address"

var _ sdk.Msg = &MsgAddAddress{}

func NewMsgAddAddress(creator string, ethAddress string) *MsgAddAddress {
	return &MsgAddAddress{
		Creator:    creator,
		EthAddress: ethAddress,
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !helpers.IsValidEthAddress(msg.EthAddress) {
		return errors.Wrapf(ErrInvalidEthAddress, "invalid eth address (%s)", msg.EthAddress)
	}

	return nil
}
