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

func (msg MsgCreateToken) Type() string { return TypeMsgCreateToken }

func (msg MsgCreateToken) Route() string { return RouterKey }

// --------------------------
// MINT TOKEN
// --------------------------
const TypeMsgMintToken = "mint_token"

var _ sdk.Msg = &MsgMintToken{}

func (msg MsgMintToken) Type() string { return TypeMsgMintToken }

func (msg MsgMintToken) Route() string { return RouterKey }

// --------------------------
// TRANSFER TOKEN
// --------------------------
const TypeMsgTransferToken = "transfer_token"

var _ sdk.Msg = &MsgTransferToken{}

func (msg MsgTransferToken) Type() string { return TypeMsgTransferToken }

func (msg MsgTransferToken) Route() string { return RouterKey }

// --------------------------
// RETIRE TOKEN
// --------------------------
const TypeMsgRetireToken = "retire_token"

var _ sdk.Msg = &MsgRetireToken{}

func (msg MsgRetireToken) Type() string { return TypeMsgRetireToken }

func (msg MsgRetireToken) Route() string { return RouterKey }

// --------------------------
// RETIRE TOKEN
// --------------------------
const TypeMsgTransferCredit = "transfer_credit"

var _ sdk.Msg = &MsgTransferCredit{}

func (msg MsgTransferCredit) Type() string { return TypeMsgTransferCredit }

func (msg MsgTransferCredit) Route() string { return RouterKey }

// --------------------------
// CANCEL TOKEN
// --------------------------
const TypeMsgCancelToken = "cancel_token"

var _ sdk.Msg = &MsgCancelToken{}

func (msg MsgCancelToken) Type() string { return TypeMsgCancelToken }

func (msg MsgCancelToken) Route() string { return RouterKey }

// --------------------------
// PAUSE TOKEN
// --------------------------
const TypeMsgPauseToken = "pause_token"

var _ sdk.Msg = &MsgPauseToken{}

func (msg MsgPauseToken) Type() string { return TypeMsgPauseToken }

func (msg MsgPauseToken) Route() string { return RouterKey }

// --------------------------
// STOP TOKEN
// --------------------------
const TypeMsgStopToken = "stop_token"

var _ sdk.Msg = &MsgStopToken{}

func (msg MsgStopToken) Type() string { return TypeMsgStopToken }

func (msg MsgStopToken) Route() string { return RouterKey }
