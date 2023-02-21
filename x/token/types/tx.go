package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//	func didToAddressSplitter(did string) (sdk.AccAddress, error) {
//		bech32 := strings.Split(did, ":")
//		address, err := sdk.AccAddressFromBech32(bech32[len(bech32)-1])
//		if err != nil {
//			return sdk.AccAddress{}, err
//		}
//		return address, nil
//	}

// --------------------------
// CREATE TOKEN
// --------------------------
const TypeMsgCreateToken = "create_token"

var _ sdk.Msg = &MsgCreateToken{}

func (msg MsgCreateToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgCreateToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateToken) Type() string { return TypeMsgCreateToken }

func (msg MsgCreateToken) Route() string { return RouterKey }

// --------------------------
// MINT TOKEN
// --------------------------
const TypeMsgMintToken = "mint_token"

var _ sdk.Msg = &MsgMintToken{}

func (msg MsgMintToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgMintToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgMintToken) Type() string { return TypeMsgMintToken }

func (msg MsgMintToken) Route() string { return RouterKey }

// --------------------------
// TRANSFER TOKEN
// --------------------------
const TypeMsgTransferToken = "transfer_token"

var _ sdk.Msg = &MsgTransferToken{}

func (msg MsgTransferToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgTransferToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgTransferToken) Type() string { return TypeMsgTransferToken }

func (msg MsgTransferToken) Route() string { return RouterKey }

// --------------------------
// RETIRE TOKEN
// --------------------------
const TypeMsgRetireToken = "retire_token"

var _ sdk.Msg = &MsgRetireToken{}

func (msg MsgRetireToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgRetireToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgRetireToken) Type() string { return TypeMsgRetireToken }

func (msg MsgRetireToken) Route() string { return RouterKey }

// --------------------------
// CANCEL TOKEN
// --------------------------
const TypeMsgCancelToken = "cancel_token"

var _ sdk.Msg = &MsgCancelToken{}

func (msg MsgCancelToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgCancelToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCancelToken) Type() string { return TypeMsgCancelToken }

func (msg MsgCancelToken) Route() string { return RouterKey }

// --------------------------
// PAUSE TOKEN
// --------------------------
const TypeMsgPauseToken = "pause_token"

var _ sdk.Msg = &MsgPauseToken{}

func (msg MsgPauseToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgPauseToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgPauseToken) Type() string { return TypeMsgPauseToken }

func (msg MsgPauseToken) Route() string { return RouterKey }

// --------------------------
// STOP TOKEN
// --------------------------
const TypeMsgStopToken = "stop_token"

var _ sdk.Msg = &MsgStopToken{}

func (msg MsgStopToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgStopToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgStopToken) Type() string { return TypeMsgStopToken }

func (msg MsgStopToken) Route() string { return RouterKey }
