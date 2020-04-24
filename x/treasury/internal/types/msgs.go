package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

type TreasuryMessage interface {
	sdk.Msg
	GetPubKey() string
	GetSenderDid() ixo.Did
}

type MsgSend struct {
	SignBytes string    `json:"signBytes" yaml:"signBytes"`
	PubKey    string    `json:"pubKey" yaml:"pubKey"`
	FromDid   ixo.Did   `json:"fromDid" yaml:"fromDid"`
	ToDid     ixo.Did   `json:"toDid" yaml:"toDid"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
}

var _ TreasuryMessage = MsgSend{}

func (msg MsgSend) Type() string  { return ModuleName }
func (msg MsgSend) Route() string { return RouterKey }
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

type MsgMint struct {
	SignBytes string    `json:"signBytes" yaml:"signBytes"`
	PubKey    string    `json:"pubKey" yaml:"pubKey"`
	OracleDid ixo.Did   `json:"oracleDid" yaml:"oracleDid"`
	ToDid     ixo.Did   `json:"toDid" yaml:"toDid"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
}

var _ TreasuryMessage = MsgMint{}

func (msg MsgMint) Type() string  { return ModuleName }
func (msg MsgMint) Route() string { return RouterKey }
func (msg MsgMint) ValidateBasic() sdk.Error {
	valid, err := CheckNotEmpty(msg.PubKey, "PubKey")
	if !valid {
		return err
	}
	valid, err = CheckNotEmpty(msg.OracleDid, "OracleDid")
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

func (msg MsgMint) GetSenderDid() ixo.Did { return msg.OracleDid }
func (msg MsgMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgMint) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgMint) GetPubKey() string { return msg.PubKey }

func (msg MsgMint) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

type MsgBurn struct {
	SignBytes string    `json:"signBytes" yaml:"signBytes"`
	PubKey    string    `json:"pubKey" yaml:"pubKey"`
	OracleDid ixo.Did   `json:"oracleDid" yaml:"oracleDid"`
	FromDid   ixo.Did   `json:"fromDid" yaml:"fromDid"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
}

var _ TreasuryMessage = MsgBurn{}

func (msg MsgBurn) Type() string  { return ModuleName }
func (msg MsgBurn) Route() string { return RouterKey }
func (msg MsgBurn) ValidateBasic() sdk.Error {
	valid, err := CheckNotEmpty(msg.PubKey, "PubKey")
	if !valid {
		return err
	}
	valid, err = CheckNotEmpty(msg.OracleDid, "OracleDid")
	if !valid {
		return err
	}

	valid, err = CheckNotEmpty(msg.FromDid, "FromDid")
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

func (msg MsgBurn) GetSenderDid() ixo.Did { return msg.OracleDid }
func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgBurn) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgBurn) GetPubKey() string { return msg.PubKey }

func (msg MsgBurn) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}
