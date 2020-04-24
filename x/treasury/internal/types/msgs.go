package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

type MsgSend struct {
	SignBytes string    `json:"signBytes" yaml:"signBytes"`
	PubKey    string    `json:"pubKey" yaml:"pubKey"`
	FromDid   ixo.Did   `json:"fromDid" yaml:"fromDid"`
	ToDid     ixo.Did   `json:"toDid" yaml:"toDid"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
}

var _ sdk.Msg = MsgSend{}

func (msg MsgSend) Type() string                            { return ModuleName }
func (msg MsgSend) Route() string                           { return RouterKey }
func (msg MsgSend) Get(key interface{}) (value interface{}) { return nil }
func (msg MsgSend) ValidateBasic() sdk.Error {
	valid, err := CheckNotEmpty(msg.PubKey, "PubKey")
	if !valid {
		return err
	}
	valid, err = CheckNotEmpty(msg.FromDid, "FromDid")
	if !valid {
		return err
	}

	valid, err = CheckNotEmpty(msg.ToDid, "ToDid")
	if !valid {
		return err
	}
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}
	if !msg.Amount.IsAllPositive() {
		return sdk.ErrInsufficientCoins("send amount must be positive")
	}

	return nil
}

func (msg MsgSend) GetSenderDid() ixo.Did { return msg.FromDid }
func (msg MsgSend) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgSend) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgSend) GetPubKey() string { return msg.PubKey }

func (msg MsgSend) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}
