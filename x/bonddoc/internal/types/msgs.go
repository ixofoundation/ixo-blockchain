package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

type MsgCreateBond struct {
	SignBytes string  `json:"signBytes" yaml:"signBytes"`
	TxHash    string  `json:"txHash" yaml:"txHash"`
	SenderDid ixo.Did `json:"senderDid" yaml:"senderDid"`
	BondDid   ixo.Did `json:"bondDid" yaml:"bondDid"`
	PubKey    string  `json:"pubKey" yaml:"pubKey"`
	Data      BondDoc `json:"data" yaml:"data"`
}

var _ sdk.Msg = MsgCreateBond{}

func (msg MsgCreateBond) Type() string  { return "create-bond" }
func (msg MsgCreateBond) Route() string { return RouterKey }
func (msg MsgCreateBond) ValidateBasic() sdk.Error {
	valid, err := CheckNotEmpty(msg.PubKey, "PubKey")
	if !valid {
		return err
	}

	valid, err = CheckNotEmpty(msg.BondDid, "BondDid")
	if !valid {
		return err
	}

	valid, err = CheckNotEmpty(msg.Data.CreatedBy, "CreatedBy")
	if !valid {
		return err
	}

	return nil
}

func (msg MsgCreateBond) GetBondDid() ixo.Did   { return msg.BondDid }
func (msg MsgCreateBond) GetSenderDid() ixo.Did { return msg.SenderDid }
func (msg MsgCreateBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetBondDid())}
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
	return []byte(msg.SignBytes)
}

func (msg MsgCreateBond) IsNewDid() bool { return true }

var _ StoredBondDoc = (*MsgCreateBond)(nil)

type MsgUpdateBondStatus struct {
	SignBytes string              `json:"signBytes" yaml:"signBytes"`
	SenderDid ixo.Did             `json:"senderDid" yaml:"senderDid"`
	BondDid   ixo.Did             `json:"bondDid" yaml:"bondDid"`
	Data      UpdateBondStatusDoc `json:"data" yaml:"data"`
}

func (msg MsgUpdateBondStatus) Type() string             { return "update-bond-status" }
func (msg MsgUpdateBondStatus) Route() string            { return RouterKey }
func (msg MsgUpdateBondStatus) ValidateBasic() sdk.Error { return nil }
func (msg MsgUpdateBondStatus) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg MsgUpdateBondStatus) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetBondDid())}
}

func (ups MsgUpdateBondStatus) GetBondDid() ixo.Did {
	return ups.BondDid
}

func (ups MsgUpdateBondStatus) GetStatus() BondStatus {
	return ups.Data.Status
}

func (msg MsgUpdateBondStatus) IsNewDid() bool     { return false }
func (msg MsgUpdateBondStatus) IsWithdrawal() bool { return false }
