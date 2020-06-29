package types

import (
	"encoding/json"
	"github.com/ixofoundation/ixo-blockchain/x/did"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

const (
	TypeMsgCreateBond       = "create-bond"
	TypeMsgUpdateBondStatus = "update-bond-status"
)

var (
	_ ixo.IxoMsg = MsgCreateBond{}
	_ ixo.IxoMsg = MsgUpdateBondStatus{}

	_ StoredBondDoc = (*MsgCreateBond)(nil)
)

type MsgCreateBond struct {
	TxHash    string  `json:"tx_hash" yaml:"tx_hash"`
	SenderDid did.Did `json:"sender_did" yaml:"sender_did"`
	BondDid   did.Did `json:"bond_did" yaml:"bond_did"`
	PubKey    string  `json:"pub_key" yaml:"pub_key"`
	Data      BondDoc `json:"data" yaml:"data"`
}

func (msg MsgCreateBond) Type() string  { return TypeMsgCreateBond }
func (msg MsgCreateBond) Route() string { return RouterKey }

func (msg MsgCreateBond) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.BondDid, "BondDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Data.CreatedBy, "CreatedBy"); !valid {
		return err
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !did.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	// No need for extra checks on Data since a blank status is valid

	return nil
}
func (msg MsgCreateBond) GetBondDid() did.Did   { return msg.BondDid }
func (msg MsgCreateBond) GetSignerDid() did.Did { return msg.GetBondDid() }
func (msg MsgCreateBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{did.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgCreateBond) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}
func (msg MsgCreateBond) GetPubKey() string     { return msg.PubKey }
func (msg MsgCreateBond) GetStatus() BondStatus { return msg.Data.Status }

func (msg *MsgCreateBond) SetStatus(status BondStatus) {
	msg.Data.Status = status
}

func (msg MsgCreateBond) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

type MsgUpdateBondStatus struct {
	SenderDid did.Did             `json:"sender_did" yaml:"sender_did"`
	BondDid   did.Did             `json:"bond_did" yaml:"bond_did"`
	Data      UpdateBondStatusDoc `json:"data" yaml:"data"`
}

func (msg MsgUpdateBondStatus) Type() string  { return TypeMsgUpdateBondStatus }
func (msg MsgUpdateBondStatus) Route() string { return RouterKey }

func (msg MsgUpdateBondStatus) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.BondDid, "BondDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !did.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	// No need for extra checks on Data since a blank status is valid
	// IsValidProgressionFrom checked by the handler

	return nil
}

func (msg MsgUpdateBondStatus) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgUpdateBondStatus) GetSignerDid() did.Did { return msg.BondDid }
func (msg MsgUpdateBondStatus) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{did.DidToAddr(msg.GetSignerDid())}
}
