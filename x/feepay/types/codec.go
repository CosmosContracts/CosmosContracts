package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global erc20 module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding.
	//
	// The actual codec used for serialization should be provided to modules/erc20 and
	// defined at the application level.
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

	// AminoCdc is a amino codec created to support amino JSON compatible msgs.
	AminoCdc = codec.NewAminoCodec(amino)
)

const (
	// Amino names
	registerFeePayContract   = "juno/MsgRegisterFeePayContract"
	unregisterFeePayContract = "juno/MsgUnregisterFeePayContract"
	fundFeePayContract       = "juno/MsgFundFeePayContract"
	updateFeeShareParams     = "juno/MsgFeePayUpdateParams"
)

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterLegacyAminoCodec(amino)

	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec
	// so that this can later be used to properly serialize MsgGrant and MsgExec
	// instances.
	RegisterLegacyAminoCodec(authzcodec.Amino)
}

// RegisterInterfaces register implementations
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterFeePayContract{},
		&MsgUnregisterFeePayContract{},
		&MsgFundFeePayContract{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers the necessary x/FeeShare interfaces and
// concrete types on the provided LegacyAmino codec. These types are used for
// Amino JSON serialization and EIP-712 compatibility.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterFeePayContract{}, registerFeePayContract, nil)
	cdc.RegisterConcrete(&MsgUnregisterFeePayContract{}, unregisterFeePayContract, nil)
	cdc.RegisterConcrete(&MsgFundFeePayContract{}, fundFeePayContract, nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, updateFeeShareParams, nil)
}
