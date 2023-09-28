package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nelsonstr/o3n1/statether/helpers"
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if helpers.IsValidEthAddress(msg.EthAddress) {

		return errors.Wrapf(ErrInvalidEthereumAddress, "invalid ethereum address (%s)", msg.EthAddress)
	}

	return nil
}
