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

func (msg MsgSend) Type() string  { return "send" }
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

type MsgOracleTransfer struct {
	SignBytes string    `json:"signBytes" yaml:"signBytes"`
	PubKey    string    `json:"pubKey" yaml:"pubKey"`
	OracleDid ixo.Did   `json:"oracleDid" yaml:"oracleDid"`
	FromDid   ixo.Did   `json:"fromDid" yaml:"fromDid"`
	ToDid     ixo.Did   `json:"toDid" yaml:"toDid"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
}

var _ TreasuryMessage = MsgOracleTransfer{}

func (msg MsgOracleTransfer) Type() string  { return "oracle-transfer" }
func (msg MsgOracleTransfer) Route() string { return RouterKey }
func (msg MsgOracleTransfer) ValidateBasic() sdk.Error {
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

func (msg MsgOracleTransfer) GetSenderDid() ixo.Did { return msg.OracleDid }
func (msg MsgOracleTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgOracleTransfer) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgOracleTransfer) GetPubKey() string { return msg.PubKey }

func (msg MsgOracleTransfer) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

type MsgOracleMint struct {
	SignBytes string    `json:"signBytes" yaml:"signBytes"`
	PubKey    string    `json:"pubKey" yaml:"pubKey"`
	OracleDid ixo.Did   `json:"oracleDid" yaml:"oracleDid"`
	ToDid     ixo.Did   `json:"toDid" yaml:"toDid"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
}

var _ TreasuryMessage = MsgOracleMint{}

func (msg MsgOracleMint) Type() string  { return "oracle-mint" }
func (msg MsgOracleMint) Route() string { return RouterKey }
func (msg MsgOracleMint) ValidateBasic() sdk.Error {
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

func (msg MsgOracleMint) GetSenderDid() ixo.Did { return msg.OracleDid }
func (msg MsgOracleMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgOracleMint) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgOracleMint) GetPubKey() string { return msg.PubKey }

func (msg MsgOracleMint) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

type MsgOracleBurn struct {
	SignBytes string    `json:"signBytes" yaml:"signBytes"`
	PubKey    string    `json:"pubKey" yaml:"pubKey"`
	OracleDid ixo.Did   `json:"oracleDid" yaml:"oracleDid"`
	FromDid   ixo.Did   `json:"fromDid" yaml:"fromDid"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
}

var _ TreasuryMessage = MsgOracleBurn{}

func (msg MsgOracleBurn) Type() string  { return "oracle-burn" }
func (msg MsgOracleBurn) Route() string { return RouterKey }
func (msg MsgOracleBurn) ValidateBasic() sdk.Error {
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

func (msg MsgOracleBurn) GetSenderDid() ixo.Did { return msg.OracleDid }
func (msg MsgOracleBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgOracleBurn) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgOracleBurn) GetPubKey() string { return msg.PubKey }

func (msg MsgOracleBurn) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}
