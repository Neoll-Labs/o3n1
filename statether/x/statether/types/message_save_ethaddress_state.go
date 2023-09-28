package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSaveEthaddressState = "save_ethaddress_state"

var _ sdk.Msg = &MsgSaveEthaddressState{}

func NewMsgSaveEthaddressState(creator string, ethAddress string, block uint64, nonce uint64, storagePosition uint64, active bool) *MsgSaveEthaddressState {
	return &MsgSaveEthaddressState{
		Creator:         creator,
		EthAddress:      ethAddress,
		Block:           block,
		Nonce:           nonce,
		StoragePosition: storagePosition,
		Active:          active,
	}
}

func (msg *MsgSaveEthaddressState) Route() string {
	return RouterKey
}

func (msg *MsgSaveEthaddressState) Type() string {
	return TypeMsgSaveEthaddressState
}

func (msg *MsgSaveEthaddressState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSaveEthaddressState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSaveEthaddressState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
