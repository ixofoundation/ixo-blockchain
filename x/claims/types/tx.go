package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// --------------------------
// CREATE CLAIM
// --------------------------
const TypeMsgCreateClaim = "create_claim"

var _ sdk.Msg = &MsgCreateClaim{}

func (msg MsgCreateClaim) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgCreateClaim) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateClaim) Type() string { return TypeMsgCreateClaim }

func (msg MsgCreateClaim) Route() string { return RouterKey }
