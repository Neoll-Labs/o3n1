package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSaveEthaddressStoragePosition = "save_ethaddress_storage_position"

var _ sdk.Msg = &MsgSaveEthaddressStoragePosition{}

func NewMsgSaveEthaddressStoragePosition(creator string, data *EthaddressStoragePosition) *MsgSaveEthaddressStoragePosition {
	return &MsgSaveEthaddressStoragePosition{
		Creator: creator,
		Data:    data,
	}
}

func (msg *MsgSaveEthaddressStoragePosition) Route() string {
	return RouterKey
}

func (msg *MsgSaveEthaddressStoragePosition) Type() string {
	return TypeMsgSaveEthaddressStoragePosition
}

func (msg *MsgSaveEthaddressStoragePosition) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSaveEthaddressStoragePosition) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSaveEthaddressStoragePosition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !isValidEthAddress(msg.Data.EthAddress) {
		return errors.Wrapf(ErrInvalidEthAddress, "invalid eth address (%s)", msg.Data.EthAddress)
	}

	return nil
}
