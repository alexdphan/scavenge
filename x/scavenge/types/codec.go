package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitScavenge{}, "scavenge/SubmitScavenge", nil)
	cdc.RegisterConcrete(&MsgRevealSolution{}, "scavenge/RevealSolution", nil)
	cdc.RegisterConcrete(&MsgCommitSolution{}, "scavenge/CommitSolution", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitScavenge{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevealSolution{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCommitSolution{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
