package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgEnableEthAddress{}, "etherstate/EnableEthAddress", nil)
	cdc.RegisterConcrete(&MsgDisableEthAddress{}, "etherstate/DisableEthAddress", nil)
	cdc.RegisterConcrete(&MsgSaveEthereumAddressState{}, "etherstate/SaveEthereumAddressState", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnableEthAddress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDisableEthAddress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSaveEthereumAddressState{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
