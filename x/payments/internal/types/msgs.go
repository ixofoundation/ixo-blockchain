package types

import (
	"encoding/json"
	"github.com/ixofoundation/ixo-blockchain/x/did"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

var (
	_ ixo.IxoMsg = MsgCreatePaymentTemplate{}
	_ ixo.IxoMsg = MsgCreatePaymentContract{}
	_ ixo.IxoMsg = MsgCreateSubscription{}
	_ ixo.IxoMsg = MsgSetPaymentContractAuthorisation{}
	_ ixo.IxoMsg = MsgGrantDiscount{}
	_ ixo.IxoMsg = MsgRevokeDiscount{}
	_ ixo.IxoMsg = MsgEffectPayment{}
)

type MsgCreatePaymentTemplate struct {
	PubKey          string          `json:"pub_key" yaml:"pub_key"`
	CreatorDid      ixo.Did         `json:"creator_did" yaml:"creator_did"`
	PaymentTemplate PaymentTemplate `json:"payment_template" yaml:"payment_template"`
}

func (msg MsgCreatePaymentTemplate) Type() string  { return "create-payment-template" }
func (msg MsgCreatePaymentTemplate) Route() string { return RouterKey }
func (msg MsgCreatePaymentTemplate) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err = CheckNotEmpty(msg.CreatorDid, "CreatorDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.CreatorDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "creator did is invalid")
	}

	// Validate PaymentTemplate
	if err := msg.PaymentTemplate.Validate(); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreatePaymentTemplate) GetSignerDid() ixo.Did { return msg.CreatorDid }
func (msg MsgCreatePaymentTemplate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgCreatePaymentTemplate) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreatePaymentTemplate) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

type MsgCreatePaymentContract struct {
	PubKey            string         `json:"pub_key" yaml:"pub_key"`
	CreatorDid        ixo.Did        `json:"creator_did" yaml:"creator_did"`
	PaymentTemplateId string         `json:"payment_template_id" yaml:"payment_template_id"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	Payer             sdk.AccAddress `json:"payer" yaml:"payer"`
	CanDeauthorise    bool           `json:"can_deauthorise" yaml:"can_deauthorise"`
	DiscountId        sdk.Uint       `json:"discount_id" yaml:"discount_id"`
}

func (msg MsgCreatePaymentContract) Type() string  { return "create-payment-contract" }
func (msg MsgCreatePaymentContract) Route() string { return RouterKey }
func (msg MsgCreatePaymentContract) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err = CheckNotEmpty(msg.CreatorDid, "CreatorDid"); !valid {
		return err
	} else if msg.Payer.Empty() {
		return sdk.ErrInvalidAddress("payer address is empty")
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.CreatorDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "creator did is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentTemplateId(msg.PaymentTemplateId) {
		return ErrInvalidId(DefaultCodespace, "payment template id invalid")
	} else if !IsValidPaymentContractId(msg.PaymentContractId) {
		return ErrInvalidId(DefaultCodespace, "payment contract id invalid")
	}

	return nil
}

func (msg MsgCreatePaymentContract) GetSignerDid() ixo.Did { return msg.CreatorDid }
func (msg MsgCreatePaymentContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgCreatePaymentContract) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreatePaymentContract) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

type MsgCreateSubscription struct {
	PubKey            string   `json:"pub_key" yaml:"pub_key"`
	CreatorDid        ixo.Did  `json:"creator_did" yaml:"creator_did"`
	SubscriptionId    string   `json:"subscription_id" yaml:"subscription_id"`
	PaymentContractId string   `json:"payment_contract_id" yaml:"payment_contract_id"`
	MaxPeriods        sdk.Uint `json:"max_periods" yaml:"max_periods"`
	Period            Period   `json:"period" yaml:"period"`
}

func (msg MsgCreateSubscription) Type() string  { return "create-subscription" }
func (msg MsgCreateSubscription) Route() string { return RouterKey }
func (msg MsgCreateSubscription) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err = CheckNotEmpty(msg.CreatorDid, "CreatorDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.CreatorDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "creator did is invalid")
	}

	// Check that IDs valid
	if !IsValidSubscriptionId(msg.SubscriptionId) {
		return ErrInvalidId(DefaultCodespace, "payment template id invalid")
	}

	// Validate Period
	if err := msg.Period.Validate(); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreateSubscription) GetSignerDid() ixo.Did { return msg.CreatorDid }
func (msg MsgCreateSubscription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgCreateSubscription) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreateSubscription) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

type MsgSetPaymentContractAuthorisation struct {
	PubKey            string  `json:"pub_key" yaml:"pub_key"`
	PayerDid          ixo.Did `json:"payer_did" yaml:"payer_did"`
	PaymentContractId string  `json:"payment_contract_id" yaml:"payment_contract_id"`
	Authorised        bool    `json:"authorised" yaml:"authorised"`
}

func (msg MsgSetPaymentContractAuthorisation) Type() string {
	return "set-payment-contract-authorisation"
}
func (msg MsgSetPaymentContractAuthorisation) Route() string { return RouterKey }
func (msg MsgSetPaymentContractAuthorisation) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err = CheckNotEmpty(msg.PayerDid, "PayerDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.PayerDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "payer did is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return ErrInvalidId(DefaultCodespace, "payment contract id invalid")
	}

	return nil
}

func (msg MsgSetPaymentContractAuthorisation) GetSignerDid() ixo.Did { return msg.PayerDid }
func (msg MsgSetPaymentContractAuthorisation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgSetPaymentContractAuthorisation) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgSetPaymentContractAuthorisation) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

type MsgGrantDiscount struct {
	PubKey            string         `json:"pub_key" yaml:"pub_key"`
	SenderDid         ixo.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	DiscountId        sdk.Uint       `json:"discount_id" yaml:"discount_id"`
	Recipient         sdk.AccAddress `json:"recipient" yaml:"recipient"`
}

func (msg MsgGrantDiscount) Type() string  { return "grant-discount" }
func (msg MsgGrantDiscount) Route() string { return RouterKey }
func (msg MsgGrantDiscount) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err = CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	} else if msg.Recipient.Empty() {
		return sdk.ErrInvalidAddress("recipient address is empty")
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return ErrInvalidId(DefaultCodespace, "payment contract id invalid")
	}

	return nil
}

func (msg MsgGrantDiscount) GetSignerDid() ixo.Did { return msg.SenderDid }
func (msg MsgGrantDiscount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgGrantDiscount) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgGrantDiscount) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

type MsgRevokeDiscount struct {
	PubKey            string         `json:"pub_key" yaml:"pub_key"`
	SenderDid         ixo.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	Holder            sdk.AccAddress `json:"holder" yaml:"holder"`
}

func (msg MsgRevokeDiscount) Type() string  { return "revoke-discount" }
func (msg MsgRevokeDiscount) Route() string { return RouterKey }
func (msg MsgRevokeDiscount) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err = CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	} else if msg.Holder.Empty() {
		return sdk.ErrInvalidAddress("holder address is empty")
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return ErrInvalidId(DefaultCodespace, "payment contract id invalid")
	}

	return nil
}

func (msg MsgRevokeDiscount) GetSignerDid() ixo.Did { return msg.SenderDid }
func (msg MsgRevokeDiscount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgRevokeDiscount) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgRevokeDiscount) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

type MsgEffectPayment struct {
	PubKey            string  `json:"pub_key" yaml:"pub_key"`
	SenderDid         ixo.Did `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string  `json:"payment_contract_id" yaml:"payment_contract_id"`
}

func (msg MsgEffectPayment) Type() string  { return "effect-payment" }
func (msg MsgEffectPayment) Route() string { return RouterKey }
func (msg MsgEffectPayment) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err = CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return ErrInvalidId(DefaultCodespace, "payment contract id invalid")
	}

	return nil
}

func (msg MsgEffectPayment) GetSignerDid() ixo.Did { return msg.SenderDid }
func (msg MsgEffectPayment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgEffectPayment) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgEffectPayment) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}
