package types

import (
	"encoding/json"
	"github.com/ixofoundation/ixo-blockchain/x/did"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

type TreasuryMessage interface {
	sdk.Msg
	GetPubKey() string
	GetSenderDid() ixo.Did
}

type MsgSend struct {
	PubKey  string    `json:"pub_key" yaml:"pub_key"`
	FromDid ixo.Did   `json:"from_did" yaml:"from_did"`
	ToDid   ixo.Did   `json:"to_did" yaml:"to_did"`
	Amount  sdk.Coins `json:"amount" yaml:"amount"`
}

var _ TreasuryMessage = MsgSend{}

func (msg MsgSend) Type() string  { return "send" }
func (msg MsgSend) Route() string { return RouterKey }
func (msg MsgSend) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err = CheckNotEmpty(msg.FromDid, "FromDid"); !valid {
		return err
	} else if valid, err = CheckNotEmpty(msg.ToDid, "ToDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.FromDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "from did is invalid")
	} else if !ixo.IsValidDid(msg.ToDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "to did is invalid")
	}

	// Check amount (note: validity also checks that coins are positive)
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
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
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return bz
	}
}

type MsgOracleTransfer struct {
	PubKey    string    `json:"pub_key" yaml:"pub_key"`
	OracleDid ixo.Did   `json:"oracle_did" yaml:"oracle_did"`
	FromDid   ixo.Did   `json:"from_did" yaml:"from_did"`
	ToDid     ixo.Did   `json:"to_did" yaml:"to_did"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
	Proof     string    `json:"proof" yaml:"proof"`
}

var _ TreasuryMessage = MsgOracleTransfer{}

func (msg MsgOracleTransfer) Type() string  { return "oracle-transfer" }
func (msg MsgOracleTransfer) Route() string { return RouterKey }
func (msg MsgOracleTransfer) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.OracleDid, "OracleDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.FromDid, "FromDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.ToDid, "ToDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.OracleDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "oracle did is invalid")
	} else if !ixo.IsValidDid(msg.FromDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "from did is invalid")
	} else if !ixo.IsValidDid(msg.ToDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "to did is invalid")
	}

	// Check amount (note: validity also checks that coins are positive)
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
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
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return bz
	}
}

type MsgOracleMint struct {
	PubKey    string    `json:"pub_key" yaml:"pub_key"`
	OracleDid ixo.Did   `json:"oracle_did" yaml:"oracle_did"`
	ToDid     ixo.Did   `json:"to_did" yaml:"to_did"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
	Proof     string    `json:"proof" yaml:"proof"`
}

var _ TreasuryMessage = MsgOracleMint{}

func (msg MsgOracleMint) Type() string  { return "oracle-mint" }
func (msg MsgOracleMint) Route() string { return RouterKey }
func (msg MsgOracleMint) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.OracleDid, "OracleDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.ToDid, "ToDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.OracleDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "oracle did is invalid")
	} else if !ixo.IsValidDid(msg.ToDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "to did is invalid")
	}

	// Check amount (note: validity also checks that coins are positive)
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
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
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return bz
	}
}

type MsgOracleBurn struct {
	PubKey    string    `json:"pub_key" yaml:"pub_key"`
	OracleDid ixo.Did   `json:"oracle_did" yaml:"oracle_did"`
	FromDid   ixo.Did   `json:"from_did" yaml:"from_did"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
	Proof     string    `json:"proof" yaml:"proof"`
}

var _ TreasuryMessage = MsgOracleBurn{}

func (msg MsgOracleBurn) Type() string  { return "oracle-burn" }
func (msg MsgOracleBurn) Route() string { return RouterKey }
func (msg MsgOracleBurn) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.OracleDid, "OracleDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.FromDid, "FromDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.OracleDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "oracle did is invalid")
	} else if !ixo.IsValidDid(msg.FromDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "from did is invalid")
	}

	// Check amount (note: validity also checks that coins are positive)
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
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
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return bz
	}
}
