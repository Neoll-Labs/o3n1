package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	helpers2 "github.com/nelsonstr/o3n1/ether-state/helpers"
)

const TypeMsgSaveEthereumAddressState = "save_ethereum_address_state"

var _ sdk.Msg = &MsgSaveEthereumAddressState{}

func NewMsgSaveEthereumAddressState(creator string, address string, blockNumber uint64, nonce uint64, storagePosition uint64) *MsgSaveEthereumAddressState {
	return &MsgSaveEthereumAddressState{
		Creator:         creator,
		Address:         address,
		BlockNumber:     blockNumber,
		Nonce:           nonce,
		StoragePosition: storagePosition,
	}
}

func (msg *MsgSaveEthereumAddressState) Route() string {
	return RouterKey
}

func (msg *MsgSaveEthereumAddressState) Type() string {
	return TypeMsgSaveEthereumAddressState
}

func (msg *MsgSaveEthereumAddressState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSaveEthereumAddressState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSaveEthereumAddressState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !helpers2.IsValidEthAddress(msg.Address) {
		return errors.Wrapf(ErrInvalidEthereumAddress, "invalid address (%s)", msg.Address)
	}

	return nil
}
