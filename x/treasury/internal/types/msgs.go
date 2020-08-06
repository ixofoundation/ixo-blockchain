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
	FromDid     did.Did   `json:"from_did" yaml:"from_did"`
	ToDidOrAddr did.Did   `json:"to_did_or_addr" yaml:"to_did_or_addr"`
	Amount      sdk.Coins `json:"amount" yaml:"amount"`
}

func (msg MsgSend) Type() string  { return TypeMsgSend }
func (msg MsgSend) Route() string { return RouterKey }
func (msg MsgSend) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.FromDid, "FromDid"); !valid {
		return err
	} else if valid, err = CheckNotEmpty(msg.ToDidOrAddr, "ToDidOrAddr"); !valid {
		return err
	}

	// Check that DIDs/addresses valid
	if !did.IsValidDid(msg.FromDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "from did is invalid")
	}
	_, err := sdk.AccAddressFromBech32(msg.ToDidOrAddr)
	if err != nil && !did.IsValidDid(msg.ToDidOrAddr) {
		return sdk.ErrInvalidAddress("recipient is neither a did nor an address")
	}

	// Check amount (note: validity also checks that coins are positive)
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}

	return nil
}

func (msg MsgSend) GetSignerDid() did.Did { return msg.FromDid }
func (msg MsgSend) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgSend) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgSend) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

type MsgOracleTransfer struct {
	OracleDid   did.Did   `json:"oracle_did" yaml:"oracle_did"`
	FromDid     did.Did   `json:"from_did" yaml:"from_did"`
	ToDidOrAddr did.Did   `json:"to_did_or_addr" yaml:"to_did_or_addr"`
	Amount      sdk.Coins `json:"amount" yaml:"amount"`
	Proof       string    `json:"proof" yaml:"proof"`
}

func (msg MsgOracleTransfer) Type() string  { return TypeMsgOracleTransfer }
func (msg MsgOracleTransfer) Route() string { return RouterKey }
func (msg MsgOracleTransfer) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.OracleDid, "OracleDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.FromDid, "FromDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.ToDidOrAddr, "ToDidOrAddr"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Proof, "Proof"); !valid {
		return err
	}

	// Check that DIDs/addresses valid
	if !did.IsValidDid(msg.OracleDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "oracle did is invalid")
	} else if !did.IsValidDid(msg.FromDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "from did is invalid")
	}
	_, err := sdk.AccAddressFromBech32(msg.ToDidOrAddr)
	if err != nil && !did.IsValidDid(msg.ToDidOrAddr) {
		return sdk.ErrInvalidAddress("recipient is neither a did nor an address")
	}

	// Check amount (note: validity also checks that coins are positive)
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}

	return nil
}

func (msg MsgOracleTransfer) GetSignerDid() did.Did { return msg.OracleDid }
func (msg MsgOracleTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgOracleTransfer) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgOracleTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

type MsgOracleMint struct {
	OracleDid   did.Did   `json:"oracle_did" yaml:"oracle_did"`
	ToDidOrAddr did.Did   `json:"to_did_or_addr" yaml:"to_did_or_addr"`
	Amount      sdk.Coins `json:"amount" yaml:"amount"`
	Proof       string    `json:"proof" yaml:"proof"`
}

func (msg MsgOracleMint) Type() string  { return TypeMsgOracleMint }
func (msg MsgOracleMint) Route() string { return RouterKey }
func (msg MsgOracleMint) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.OracleDid, "OracleDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.ToDidOrAddr, "ToDidOrAddr"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Proof, "Proof"); !valid {
		return err
	}

	// Check that DIDs/addresses valid
	if !did.IsValidDid(msg.OracleDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "oracle did is invalid")
	}
	_, err := sdk.AccAddressFromBech32(msg.ToDidOrAddr)
	if err != nil && !did.IsValidDid(msg.ToDidOrAddr) {
		return sdk.ErrInvalidAddress("recipient is neither a did nor an address")
	}

	// Check amount (note: validity also checks that coins are positive)
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}

	return nil
}

func (msg MsgOracleMint) GetSignerDid() did.Did { return msg.OracleDid }
func (msg MsgOracleMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgOracleMint) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgOracleMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
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
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgOracleBurn) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgOracleBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
