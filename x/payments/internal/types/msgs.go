package types

import (
	"encoding/json"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/did"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

const (
	TypeMsgCreatePaymentTemplate           = "create-payment-template"
	TypeMsgCreatePaymentContract           = "create-payment-contract"
	TypeMsgCreateSubscription              = "create-subscription"
	TypeMsgSetPaymentContractAuthorisation = "set-payment-contract-authorisation"
	TypeMsgGrantDiscount                   = "grant-discount"
	TypeMsgRevokeDiscount                  = "revoke-discount"
	TypeMsgEffectPayment                   = "effect-payment"
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
	CreatorDid      did.Did         `json:"creator_did" yaml:"creator_did"`
	PaymentTemplate PaymentTemplate `json:"payment_template" yaml:"payment_template"`
}

func (msg MsgCreatePaymentTemplate) Type() string  { return TypeMsgCreatePaymentTemplate }
func (msg MsgCreatePaymentTemplate) Route() string { return RouterKey }
func (msg MsgCreatePaymentTemplate) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.CreatorDid, "CreatorDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.CreatorDid) {
		return sdkerrors.Wrap(ErrInvalidId, "creator did is invalid")
	}

	// Validate PaymentTemplate
	if err := msg.PaymentTemplate.Validate(); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreatePaymentTemplate) GetSignerDid() did.Did { return msg.CreatorDid }
func (msg MsgCreatePaymentTemplate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreatePaymentTemplate) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreatePaymentTemplate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

type MsgCreatePaymentContract struct {
	CreatorDid        did.Did        `json:"creator_did" yaml:"creator_did"`
	PaymentTemplateId string         `json:"payment_template_id" yaml:"payment_template_id"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	Payer             sdk.AccAddress `json:"payer" yaml:"payer"`
	CanDeauthorise    bool           `json:"can_deauthorise" yaml:"can_deauthorise"`
	DiscountId        sdk.Uint       `json:"discount_id" yaml:"discount_id"`
}

func (msg MsgCreatePaymentContract) Type() string  { return TypeMsgCreatePaymentContract }
func (msg MsgCreatePaymentContract) Route() string { return RouterKey }
func (msg MsgCreatePaymentContract) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.CreatorDid, "CreatorDid"); !valid {
		return err
	} else if msg.Payer.Empty() {
		return sdkerrors.Wrap(did.ErrInvalidDid, "payer did is empty")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.CreatorDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "creator did is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentTemplateId(msg.PaymentTemplateId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment template id invalid")
	} else if !IsValidPaymentContractId(msg.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract id invalid")
	}

	return nil
}

func (msg MsgCreatePaymentContract) GetSignerDid() did.Did { return msg.CreatorDid }
func (msg MsgCreatePaymentContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreatePaymentContract) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreatePaymentContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

type MsgCreateSubscription struct {
	CreatorDid        did.Did  `json:"creator_did" yaml:"creator_did"`
	SubscriptionId    string   `json:"subscription_id" yaml:"subscription_id"`
	PaymentContractId string   `json:"payment_contract_id" yaml:"payment_contract_id"`
	MaxPeriods        sdk.Uint `json:"max_periods" yaml:"max_periods"`
	Period            Period   `json:"period" yaml:"period"`
}

func (msg MsgCreateSubscription) Type() string  { return TypeMsgCreateSubscription }
func (msg MsgCreateSubscription) Route() string { return RouterKey }
func (msg MsgCreateSubscription) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.CreatorDid, "CreatorDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.CreatorDid) {
		return sdkerrors.Wrap(ErrInvalidId, "creator did is invalid")
	}

	// Check that IDs valid
	if !IsValidSubscriptionId(msg.SubscriptionId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment template id invalid")
	}

	// Validate Period
	if err := msg.Period.Validate(); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreateSubscription) GetSignerDid() did.Did { return msg.CreatorDid }
func (msg MsgCreateSubscription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreateSubscription) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreateSubscription) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

type MsgSetPaymentContractAuthorisation struct {
	PayerDid          did.Did `json:"payer_did" yaml:"payer_did"`
	PaymentContractId string  `json:"payment_contract_id" yaml:"payment_contract_id"`
	Authorised        bool    `json:"authorised" yaml:"authorised"`
}

func (msg MsgSetPaymentContractAuthorisation) Type() string {
	return TypeMsgSetPaymentContractAuthorisation
}
func (msg MsgSetPaymentContractAuthorisation) Route() string { return RouterKey }
func (msg MsgSetPaymentContractAuthorisation) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PayerDid, "PayerDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.PayerDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "payer did is invalid")

	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract id invalid")
	}

	return nil
}

func (msg MsgSetPaymentContractAuthorisation) GetSignerDid() did.Did { return msg.PayerDid }
func (msg MsgSetPaymentContractAuthorisation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgSetPaymentContractAuthorisation) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgSetPaymentContractAuthorisation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

type MsgGrantDiscount struct {
	SenderDid         did.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	DiscountId        sdk.Uint       `json:"discount_id" yaml:"discount_id"`
	Recipient         sdk.AccAddress `json:"recipient" yaml:"recipient"`
}

func (msg MsgGrantDiscount) Type() string  { return TypeMsgGrantDiscount }
func (msg MsgGrantDiscount) Route() string { return RouterKey }
func (msg MsgGrantDiscount) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	} else if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "recipient address is empty")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "sender did is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract id invalid")
	}

	return nil
}

func (msg MsgGrantDiscount) GetSignerDid() did.Did { return msg.SenderDid }
func (msg MsgGrantDiscount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgGrantDiscount) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgGrantDiscount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

type MsgRevokeDiscount struct {
	SenderDid         did.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	Holder            sdk.AccAddress `json:"holder" yaml:"holder"`
}

func (msg MsgRevokeDiscount) Type() string  { return TypeMsgRevokeDiscount }
func (msg MsgRevokeDiscount) Route() string { return RouterKey }
func (msg MsgRevokeDiscount) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	} else if msg.Holder.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "holder address is empty")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "sender did is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract id invalid")
	}

	return nil
}

func (msg MsgRevokeDiscount) GetSignerDid() did.Did { return msg.SenderDid }
func (msg MsgRevokeDiscount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgRevokeDiscount) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgRevokeDiscount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

type MsgEffectPayment struct {
	SenderDid         did.Did `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string  `json:"payment_contract_id" yaml:"payment_contract_id"`
}

func (msg MsgEffectPayment) Type() string  { return TypeMsgEffectPayment }
func (msg MsgEffectPayment) Route() string { return RouterKey }
func (msg MsgEffectPayment) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "sender did is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract id invalid")
	}

	return nil
}

func (msg MsgEffectPayment) GetSignerDid() did.Did { return msg.SenderDid }
func (msg MsgEffectPayment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgEffectPayment) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgEffectPayment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
