package types

import (
	"encoding/json"
	"fmt"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	didexported "github.com/ixofoundation/ixo-blockchain/lib/legacydid"
	didtypes "github.com/ixofoundation/ixo-blockchain/lib/legacydid"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ixotypes "github.com/ixofoundation/ixo-blockchain/lib/ixo"
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
	_ ixotypes.IxoMsg = &MsgCreatePaymentTemplate{}
	_ ixotypes.IxoMsg = &MsgCreatePaymentContract{}
	_ ixotypes.IxoMsg = &MsgCreateSubscription{}
	_ ixotypes.IxoMsg = &MsgSetPaymentContractAuthorisation{}
	_ ixotypes.IxoMsg = &MsgGrantDiscount{}
	_ ixotypes.IxoMsg = &MsgRevokeDiscount{}
	_ ixotypes.IxoMsg = &MsgEffectPayment{}
)

func NewMsgCreatePaymentTemplate(template PaymentTemplate,
	creatorDid didexported.Did) *MsgCreatePaymentTemplate {
	return &MsgCreatePaymentTemplate{
		CreatorDid:      creatorDid,
		PaymentTemplate: template,
	}
}
func (msg MsgCreatePaymentTemplate) GetIidController() string { return msg.CreatorDid }

func (msg MsgCreatePaymentTemplate) Type() string  { return TypeMsgCreatePaymentTemplate }
func (msg MsgCreatePaymentTemplate) Route() string { return RouterKey }
func (msg MsgCreatePaymentTemplate) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.CreatorDid, "CreatorDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.CreatorDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "creator DID is invalid")
	}

	// Validate PaymentTemplate
	if err := msg.PaymentTemplate.Validate(); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreatePaymentTemplate) GetSignerDid() didexported.Did { return msg.CreatorDid }
func (msg MsgCreatePaymentTemplate) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.CreatorAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgCreatePaymentTemplate) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreatePaymentTemplate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func NewMsgCreatePaymentContract(templateId, contractId string,
	payer sdk.AccAddress, recipients Distribution, canDeauthorise bool,
	discountId sdk.Uint, creatorDid didexported.Did) *MsgCreatePaymentContract {
	return &MsgCreatePaymentContract{
		CreatorDid:        creatorDid,
		PaymentTemplateId: templateId,
		PaymentContractId: contractId,
		Payer:             payer.String(),
		Recipients:        recipients,
		CanDeauthorise:    canDeauthorise,
		DiscountId:        discountId,
	}
}
func (msg MsgCreatePaymentContract) GetIidController() string { return msg.CreatorDid }

func (msg MsgCreatePaymentContract) Type() string  { return TypeMsgCreatePaymentContract }
func (msg MsgCreatePaymentContract) Route() string { return RouterKey }
func (msg MsgCreatePaymentContract) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.CreatorDid, "CreatorDid"); !valid {
		return err
	} else if strings.TrimSpace(msg.Payer) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "payer address is empty")
	}

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.CreatorDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "creator DID is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract ID invalid")
	} else if !IsValidPaymentTemplateId(msg.PaymentTemplateId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment template ID invalid")
	}

	// Validate recipient distribution
	var recipients Distribution = msg.Recipients
	if err := recipients.Validate(); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreatePaymentContract) GetSignerDid() didexported.Did { return msg.CreatorDid }
func (msg MsgCreatePaymentContract) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.CreatorAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgCreatePaymentContract) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreatePaymentContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func NewMsgCreateSubscription(subscriptionId, contractId string, maxPeriods sdk.Uint,
	period Period, creatorDid didexported.Did) *MsgCreateSubscription {
	msg := &MsgCreateSubscription{
		CreatorDid:        creatorDid,
		SubscriptionId:    subscriptionId,
		PaymentContractId: contractId,
		MaxPeriods:        maxPeriods,
	}

	err := msg.SetPeriod(period)
	if err != nil {
		panic(err)
	}

	return msg
}

func (msg MsgCreateSubscription) GetIidController() string { return msg.CreatorDid }

func (msg *MsgCreateSubscription) SetPeriod(period Period) error {
	m, ok := period.(proto.Message)
	if !ok {
		return fmt.Errorf("can't proto marshal %T", m)
	}

	any, err := codectypes.NewAnyWithValue(m)
	if err != nil {
		return err
	}

	msg.Period = any
	return nil
}

func (msg *MsgCreateSubscription) GetPeriod() Period {
	period, ok := msg.Period.GetCachedValue().(Period)
	if !ok {
		return nil
	}
	return period
}

func (msg MsgCreateSubscription) Type() string  { return TypeMsgCreateSubscription }
func (msg MsgCreateSubscription) Route() string { return RouterKey }
func (msg MsgCreateSubscription) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.CreatorDid, "CreatorDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.CreatorDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "creator DID is invalid")
	}

	// Check that IDs valid
	if !IsValidSubscriptionId(msg.SubscriptionId) {
		return sdkerrors.Wrap(ErrInvalidId, "subscription ID invalid")
	} else if !IsValidPaymentContractId(msg.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract ID invalid")
	}

	// Validate Period
	period := msg.GetPeriod()
	if period == nil {
		return sdkerrors.Wrap(ErrInvalidPeriod, "missing period")
	}
	if err := period.Validate(); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreateSubscription) GetSignerDid() didexported.Did { return msg.CreatorDid }
func (msg MsgCreateSubscription) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.CreatorAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgCreateSubscription) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreateSubscription) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

var _ codectypes.UnpackInterfacesMessage = MsgCreateSubscription{}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m MsgCreateSubscription) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var period Period
	return unpacker.UnpackAny(m.Period, &period)
}

func NewMsgSetPaymentContractAuthorisation(contractId string, authorised bool,
	payerDid didexported.Did) *MsgSetPaymentContractAuthorisation {
	return &MsgSetPaymentContractAuthorisation{
		PayerDid:          payerDid,
		PaymentContractId: contractId,
		Authorised:        authorised,
	}
}
func (msg MsgSetPaymentContractAuthorisation) GetIidController() string { return msg.PayerDid }

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
	if !didtypes.IsValidDid(msg.PayerDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "payer DID is invalid")

	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract ID invalid")
	}

	return nil
}

func (msg MsgSetPaymentContractAuthorisation) GetSignerDid() didexported.Did { return msg.PayerDid }
func (msg MsgSetPaymentContractAuthorisation) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.PayerAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgSetPaymentContractAuthorisation) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgSetPaymentContractAuthorisation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func NewMsgGrantDiscount(contractId string, discountId sdk.Uint,
	recipient sdk.AccAddress, creatorDid didexported.Did) *MsgGrantDiscount {
	return &MsgGrantDiscount{
		SenderDid:         creatorDid,
		PaymentContractId: contractId,
		DiscountId:        discountId,
		Recipient:         recipient.String(),
	}
}

func (msg MsgGrantDiscount) GetIidController() string { return msg.SenderDid }

func (msg MsgGrantDiscount) Type() string  { return TypeMsgGrantDiscount }
func (msg MsgGrantDiscount) Route() string { return RouterKey }
func (msg MsgGrantDiscount) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	} else if strings.TrimSpace(msg.Recipient) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "recipient address is empty")
	}

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender DID is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract ID invalid")
	}

	return nil
}

func (msg MsgGrantDiscount) GetSignerDid() didexported.Did { return msg.SenderDid }
func (msg MsgGrantDiscount) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.SenderAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgGrantDiscount) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgGrantDiscount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func NewMsgRevokeDiscount(contractId string, holder sdk.AccAddress,
	creatorDid didexported.Did) *MsgRevokeDiscount {
	return &MsgRevokeDiscount{
		SenderDid:         creatorDid,
		PaymentContractId: contractId,
		Holder:            holder.String(),
	}
}

func (msg MsgRevokeDiscount) GetIidController() string { return msg.SenderDid }

func (msg MsgRevokeDiscount) Type() string  { return TypeMsgRevokeDiscount }
func (msg MsgRevokeDiscount) Route() string { return RouterKey }
func (msg MsgRevokeDiscount) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	} else if strings.TrimSpace(msg.Holder) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "holder address is empty")
	}

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender DID is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract ID invalid")
	}

	return nil
}

func (msg MsgRevokeDiscount) GetSignerDid() didexported.Did { return msg.SenderDid }
func (msg MsgRevokeDiscount) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.SenderAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgRevokeDiscount) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgRevokeDiscount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func NewMsgEffectPayment(contractId string, creatorDid didexported.Did) *MsgEffectPayment {
	return &MsgEffectPayment{
		SenderDid:         creatorDid,
		PaymentContractId: contractId,
	}
}

func (msg MsgEffectPayment) GetIidController() string { return msg.SenderDid }

func (msg MsgEffectPayment) Type() string  { return TypeMsgEffectPayment }
func (msg MsgEffectPayment) Route() string { return RouterKey }
func (msg MsgEffectPayment) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender DID is invalid")
	}

	// Check that IDs valid
	if !IsValidPaymentContractId(msg.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract ID invalid")
	}

	return nil
}

func (msg MsgEffectPayment) GetSignerDid() didexported.Did { return msg.SenderDid }
func (msg MsgEffectPayment) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.SenderAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgEffectPayment) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgEffectPayment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}
