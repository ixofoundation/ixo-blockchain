package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/types"
)

// *****IssueFiat

// IssueFiat - transaction input
type IssueFiat struct {
	IssuerAddress sdk.AccAddress `json:"issuerAddress"`
	ToAddress     sdk.AccAddress `json:"toAddress"`
	FiatPeg       types.FiatPeg  `json:"fiatPeg"`
}

// NewIssueFiat : initializer
func NewIssueFiat(issuerAddress sdk.AccAddress, toAddress sdk.AccAddress, fiatPeg types.FiatPeg) IssueFiat {
	return IssueFiat{issuerAddress, toAddress, fiatPeg}
}

// GetSignBytes : get bytes to sign
func (in IssueFiat) GetSignBytes() []byte {
	bin, err := ModuleCdc.MarshalJSON(struct {
		IssuerAddress string        `json:"issuerAddress"`
		ToAddress     string        `json:"toAddress"`
		FiatPeg       types.FiatPeg `json:"fiatPeg"`
	}{
		IssuerAddress: in.IssuerAddress.String(),
		ToAddress:     in.ToAddress.String(),
		FiatPeg:       in.FiatPeg,
	})
	if err != nil {
		panic(err)
	}
	return bin
}

func (in IssueFiat) ValidateBasic() sdk.Error {
	if len(in.IssuerAddress) == 0 {
		return sdk.ErrInvalidAddress(in.IssuerAddress.String())
	} else if len(in.ToAddress) == 0 {
		return sdk.ErrInvalidAddress(in.ToAddress.String())
	} else if in.FiatPeg.GetTransactionAmount() < 0 {
		return ErrNegativeAmount(DefaultCodeSpace, "Transaction amount should be grater than 0.")
	} else if in.FiatPeg.GetTransactionID() == "" {
		return sdk.ErrUnknownRequest("Transaction should not be empty")
	}
	return nil
}

// #####IssueFiat

// *****MsgIssueFiats

// MsgIssueFiats : high level issuance of fiats module
type MsgIssueFiats struct {
	IssueFiats []IssueFiat `json:"issueFiats"`
}

// NewMsgIssueFiats : initilizer
func NewMsgIssueFiats(issueFiats []IssueFiat) MsgIssueFiats {
	return MsgIssueFiats{issueFiats}
}

// ***** Implementing sdk.Msg

var _ sdk.Msg = MsgIssueFiats{}

// Type : implements msg
func (msg MsgIssueFiats) Type() string { return "bank" }

func (msg MsgIssueFiats) Route() string { return RouterKey }

// ValidateBasic : implements msg
func (msg MsgIssueFiats) ValidateBasic() sdk.Error {
	if len(msg.IssueFiats) == 0 {
		return ErrNoOutputs(DefaultCodeSpace).TraceSDK("")
	}
	for _, in := range msg.IssueFiats {
		if err := in.ValidateBasic(); err != nil {
			return err.TraceSDK("")
		}
	}
	return nil
}

// GetSignBytes : implements msg
func (msg MsgIssueFiats) GetSignBytes() []byte {
	var issueFiats []json.RawMessage
	for _, issueFiat := range msg.IssueFiats {
		issueFiats = append(issueFiats, issueFiat.GetSignBytes())
	}

	b, err := ModuleCdc.MarshalJSON(struct {
		IssueFiats []json.RawMessage `json:"issueFiats"`
	}{
		IssueFiats: issueFiats,
	})
	if err != nil {
		panic(err)
	}
	return b
}

// GetSigners : implements msg
func (msg MsgIssueFiats) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, len(msg.IssueFiats))
	for i, in := range msg.IssueFiats {
		addrs[i] = in.IssuerAddress
	}
	return addrs
}

// ##### Implement sdk.Msg

// #####MsgIssueFiats

// ****RedeemFiat

// RedeemFiat : transaction input
type RedeemFiat struct {
	RedeemerAddress sdk.AccAddress `json:"redeemerAddress"`
	IssuerAddress   sdk.AccAddress `json:"issuerAddress"`
	Amount          int64          `json:"amount"`
}

// NewRedeemFiat : initializer
func NewRedeemFiat(redeemerAddress sdk.AccAddress, issuerAddress sdk.AccAddress, amount int64) RedeemFiat {
	return RedeemFiat{redeemerAddress, issuerAddress, amount}
}

// GetSignBytes : get bytes to sign
func (in RedeemFiat) GetSignBytes() []byte {
	bin, err := ModuleCdc.MarshalJSON(struct {
		RedeemerAddress string `json:"redeemerAddress"`
		IssuerAddress   string `json:"issuerAddress"`
		Amount          int64  `json:"amount"`
	}{
		RedeemerAddress: in.RedeemerAddress.String(),
		IssuerAddress:   in.IssuerAddress.String(),
		Amount:          in.Amount,
	})
	if err != nil {
		panic(err)
	}
	return bin
}
func (in RedeemFiat) ValidateBasic() sdk.Error {
	if len(in.IssuerAddress) == 0 {
		return sdk.ErrInvalidAddress(in.IssuerAddress.String())
	} else if len(in.RedeemerAddress) == 0 {
		return sdk.ErrInvalidAddress(in.RedeemerAddress.String())
	} else if in.Amount <= 0 {
		return sdk.ErrUnknownRequest("Amount should be Positive")
	}
	return nil
}

// #####RedeemFiat

// *****MsgRedeemFiats

// MsgRedeemFiats : Message to redeem issued fiats
type MsgRedeemFiats struct {
	RedeemFiats []RedeemFiat `json:"redeemFiats"`
}

// NewMsgRedeemFiats : initializer
func NewMsgRedeemFiats(redeemFiats []RedeemFiat) MsgRedeemFiats {
	return MsgRedeemFiats{redeemFiats}
}

// *****Implementing sdk.Msg

var _ sdk.Msg = MsgRedeemFiats{}

// Type : implements msg
func (msg MsgRedeemFiats) Type() string { return "bank" }

func (msg MsgRedeemFiats) Route() string { return RouterKey }

// ValidateBasic : implements msg
func (msg MsgRedeemFiats) ValidateBasic() sdk.Error {
	if len(msg.RedeemFiats) == 0 {
		return ErrNoOutputs(DefaultCodeSpace).TraceSDK("")
	}
	for _, in := range msg.RedeemFiats {
		if err := in.ValidateBasic(); err != nil {
			return err.TraceSDK("")
		}
	}
	return nil
}

// GetSignBytes : implements msg
func (msg MsgRedeemFiats) GetSignBytes() []byte {
	var redeemFiats []json.RawMessage
	for _, redeemFiat := range msg.RedeemFiats {
		redeemFiats = append(redeemFiats, redeemFiat.GetSignBytes())
	}

	bz, err := ModuleCdc.MarshalJSON(struct {
		RedeemFiats []json.RawMessage `json:"redeemFiats"`
	}{
		RedeemFiats: redeemFiats,
	})
	if err != nil {
		panic(err)
	}
	return bz
}

// GetSigners : implements msg
func (msg MsgRedeemFiats) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, len(msg.RedeemFiats))
	for i, in := range msg.RedeemFiats {
		addrs[i] = in.RedeemerAddress
	}
	return addrs
}

// ##### Implement sdk.Msg

// ######MsgRedeemFiats

// *****SendFiat

// SendFiat - transaction input
type SendFiat struct {
	FromAddress sdk.AccAddress `json:"fromAddress"`
	ToAddress   sdk.AccAddress `json:"toAddress"`
	PegHash     types.PegHash  `json:"pegHash"`
	Amount      int64          `json:"amount"`
}

// NewSendFiat : initializer
func NewSendFiat(fromAddress sdk.AccAddress, toAddress sdk.AccAddress, pegHash types.PegHash, amount int64) SendFiat {
	return SendFiat{fromAddress, toAddress, pegHash, amount}
}

// GetSignBytes : get bytes to sign
func (in SendFiat) GetSignBytes() []byte {
	bin, err := ModuleCdc.MarshalJSON(struct {
		FromAddress string `json:"fromAddress"`
		ToAddress   string `json:"toAddress"`
		PegHash     string `json:"pegHash"`
		Amount      int64  `json:"amount"`
	}{
		FromAddress: in.FromAddress.String(),
		ToAddress:   in.ToAddress.String(),
		PegHash:     in.PegHash.String(),
		Amount:      in.Amount,
	})
	if err != nil {
		panic(err)
	}
	return bin
}

func (in SendFiat) ValidateBasic() sdk.Error {
	if len(in.FromAddress) == 0 {
		return sdk.ErrInvalidAddress(in.FromAddress.String())
	} else if len(in.ToAddress) == 0 {
		return sdk.ErrInvalidAddress(in.ToAddress.String())
	} else if len(in.PegHash) == 0 {
		return sdk.ErrUnknownRequest("PegHash is Empty")
	} else if in.Amount <= 0 {
		return ErrNegativeAmount(DefaultCodeSpace, "Amount should be positive")
	}
	return nil
}

// #####SendFiat

// *****MsgSendFiats

// MsgSendFiats : high level issuance of fiats module
type MsgSendFiats struct {
	SendFiats []SendFiat `json:"sendFiats"`
}

// NewMsgSendFiats : initilizer
func NewMsgSendFiats(sendFiats []SendFiat) MsgSendFiats {
	return MsgSendFiats{sendFiats}
}

// ***** Implementing sdk.Msg

var _ sdk.Msg = MsgSendFiats{}

// Type : implements msg
func (msg MsgSendFiats) Type() string { return "bank" }

func (msg MsgSendFiats) Route() string { return RouterKey }

// ValidateBasic : implements msg
func (msg MsgSendFiats) ValidateBasic() sdk.Error {
	if len(msg.SendFiats) == 0 {
		return ErrNoOutputs(DefaultCodeSpace).TraceSDK("")
	}
	for _, in := range msg.SendFiats {
		if err := in.ValidateBasic(); err != nil {
			return err.TraceSDK("")
		}
	}
	return nil
}

// GetSignBytes : implements msg
func (msg MsgSendFiats) GetSignBytes() []byte {
	var sendFiats []json.RawMessage
	for _, sendFiat := range msg.SendFiats {
		sendFiats = append(sendFiats, sendFiat.GetSignBytes())
	}

	b, err := ModuleCdc.MarshalJSON(struct {
		SendFiats []json.RawMessage `json:"sendFiats"`
	}{
		SendFiats: sendFiats,
	})
	if err != nil {
		panic(err)
	}
	return b
}

// GetSigners : implements msg
func (msg MsgSendFiats) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, len(msg.SendFiats))
	for i, in := range msg.SendFiats {
		addrs[i] = in.FromAddress
	}
	return addrs
}

// ##### Implement sdk.Msg

// #####MsgSendFiats
