package types

import (
	"encoding/json"
	"github.com/ixofoundation/ixo-blockchain/x/did"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

const (
	TypeMsgSend           = "send"
	TypeMsgOracleTransfer = "oracle-transfer"
	TypeMsgOracleMint     = "oracle-mint"
	TypeMsgOracleBurn     = "oracle-burn"
)

var (
	_ ixo.IxoMsg = MsgSend{}
	_ ixo.IxoMsg = MsgOracleTransfer{}
	_ ixo.IxoMsg = MsgOracleMint{}
	_ ixo.IxoMsg = MsgOracleBurn{}
)

type MsgSend struct {
	FromDid did.Did   `json:"from_did" yaml:"from_did"`
	ToDid   did.Did   `json:"to_did" yaml:"to_did"`
	Amount  sdk.Coins `json:"amount" yaml:"amount"`
}

func (msg MsgSend) Type() string  { return TypeMsgSend }
func (msg MsgSend) Route() string { return RouterKey }
func (msg MsgSend) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.FromDid, "FromDid"); !valid {
		return err
	} else if valid, err = CheckNotEmpty(msg.ToDid, "ToDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.FromDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "from did is invalid")
	} else if !did.IsValidDid(msg.ToDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "to did is invalid")
	}

	// Check amount (note: validity also checks that coins are positive)
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}

	return nil
}

func (msg MsgSend) GetSignerDid() did.Did { return msg.FromDid }
func (msg MsgSend) GetSigners() []sdk.AccAddress {
	panic("tried to use unimplemented GetSigners function")
}

func (msg MsgSend) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgSend) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

type MsgOracleTransfer struct {
	OracleDid did.Did   `json:"oracle_did" yaml:"oracle_did"`
	FromDid   did.Did   `json:"from_did" yaml:"from_did"`
	ToDid     did.Did   `json:"to_did" yaml:"to_did"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
	Proof     string    `json:"proof" yaml:"proof"`
}

func (msg MsgOracleTransfer) Type() string  { return TypeMsgOracleTransfer }
func (msg MsgOracleTransfer) Route() string { return RouterKey }
func (msg MsgOracleTransfer) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.OracleDid, "OracleDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.FromDid, "FromDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.ToDid, "ToDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Proof, "Proof"); !valid {
		return err
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.OracleDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "oracle did is invalid")
	} else if !did.IsValidDid(msg.FromDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "from did is invalid")
	} else if !did.IsValidDid(msg.ToDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "to did is invalid")
	}

	// Check amount (note: validity also checks that coins are positive)
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}

	return nil
}

func (msg MsgOracleTransfer) GetSignerDid() did.Did { return msg.OracleDid }
func (msg MsgOracleTransfer) GetSigners() []sdk.AccAddress {
	panic("tried to use unimplemented GetSigners function")
}

func (msg MsgOracleTransfer) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgOracleTransfer) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

type MsgOracleMint struct {
	OracleDid did.Did   `json:"oracle_did" yaml:"oracle_did"`
	ToDid     did.Did   `json:"to_did" yaml:"to_did"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
	Proof     string    `json:"proof" yaml:"proof"`
}

func (msg MsgOracleMint) Type() string  { return TypeMsgOracleMint }
func (msg MsgOracleMint) Route() string { return RouterKey }
func (msg MsgOracleMint) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.OracleDid, "OracleDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.ToDid, "ToDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Proof, "Proof"); !valid {
		return err
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.OracleDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "oracle did is invalid")
	} else if !did.IsValidDid(msg.ToDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "to did is invalid")
	}

	// Check amount (note: validity also checks that coins are positive)
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}

	return nil
}

func (msg MsgOracleMint) GetSignerDid() did.Did { return msg.OracleDid }
func (msg MsgOracleMint) GetSigners() []sdk.AccAddress {
	panic("tried to use unimplemented GetSigners function")
}

func (msg MsgOracleMint) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgOracleMint) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

type MsgOracleBurn struct {
	OracleDid did.Did   `json:"oracle_did" yaml:"oracle_did"`
	FromDid   did.Did   `json:"from_did" yaml:"from_did"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
	Proof     string    `json:"proof" yaml:"proof"`
}

func (msg MsgOracleBurn) Type() string  { return TypeMsgOracleBurn }
func (msg MsgOracleBurn) Route() string { return RouterKey }
func (msg MsgOracleBurn) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.OracleDid, "OracleDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.FromDid, "FromDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Proof, "Proof"); !valid {
		return err
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.OracleDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "oracle did is invalid")
	} else if !did.IsValidDid(msg.FromDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "from did is invalid")
	}

	// Check amount (note: validity also checks that coins are positive)
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}

	return nil
}

func (msg MsgOracleBurn) GetSignerDid() did.Did { return msg.OracleDid }
func (msg MsgOracleBurn) GetSigners() []sdk.AccAddress {
	panic("tried to use unimplemented GetSigners function")
}

func (msg MsgOracleBurn) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgOracleBurn) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}
