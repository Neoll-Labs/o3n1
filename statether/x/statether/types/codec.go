package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddAddress{}, "statether/AddAddress", nil)
	cdc.RegisterConcrete(&MsgSaveEthaddressStoragePosition{}, "statether/SaveEthaddressStoragePosition", nil)
	cdc.RegisterConcrete(&MsgRemoveAddress{}, "statether/RemoveAddress", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddAddress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSaveEthaddressStoragePosition{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveAddress{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
