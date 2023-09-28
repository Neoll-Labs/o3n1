package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGetEthaddressState = "get_ethaddress_state"

var _ sdk.Msg = &MsgGetEthaddressState{}

func NewMsgGetEthaddressState(creator string, ethAddress string, block uint64, nonce uint64, storagePosition uint64, active bool) *MsgGetEthaddressState {
	return &MsgGetEthaddressState{
		Creator:         creator,
		EthAddress:      ethAddress,
		Block:           block,
		Nonce:           nonce,
		StoragePosition: storagePosition,
		Active:          active,
	}
}

func (msg *MsgGetEthaddressState) Route() string {
	return RouterKey
}

func (msg *MsgGetEthaddressState) Type() string {
	return TypeMsgGetEthaddressState
}

func (msg *MsgGetEthaddressState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGetEthaddressState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGetEthaddressState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
