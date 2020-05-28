package types

import (
	"encoding/json"
	"github.com/ixofoundation/ixo-blockchain/x/did"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

type FeesMessage interface {
	sdk.Msg
	GetPubKey() string
	GetSenderDid() ixo.Did
}

type MsgSetFeeContractAuthorisation struct {
	SignBytes     string  `json:"signBytes" yaml:"signBytes"`
	PubKey        string  `json:"pub_key" yaml:"pub_key"`
	PayerDid      ixo.Did `json:"payer_did" yaml:"payer_did"`
	FeeContractId string  `json:"fee_contract_id" yaml:"fee_contract_id"`
	Authorised    bool    `json:"authorised" yaml:"authorised"`
}

var _ FeesMessage = MsgSetFeeContractAuthorisation{}

func (msg MsgSetFeeContractAuthorisation) Type() string  { return "set-fee-contract-authorisation" }
func (msg MsgSetFeeContractAuthorisation) Route() string { return RouterKey }
func (msg MsgSetFeeContractAuthorisation) ValidateBasic() sdk.Error {
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
	if !IsValidFeeContractId(msg.FeeContractId) {
		return ErrInvalidId(DefaultCodespace, "fee contract id invalid")
	}

	return nil
}

func (msg MsgSetFeeContractAuthorisation) GetSenderDid() ixo.Did { return msg.PayerDid }
func (msg MsgSetFeeContractAuthorisation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgSetFeeContractAuthorisation) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgSetFeeContractAuthorisation) GetPubKey() string { return msg.PubKey }

func (msg MsgSetFeeContractAuthorisation) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

type MsgCreateFee struct {
	SignBytes  string     `json:"signBytes" yaml:"signBytes"`
	PubKey     string     `json:"pub_key" yaml:"pub_key"`
	CreatorDid ixo.Did    `json:"creator_did" yaml:"creator_did"`
	FeeId      string     `json:"fee_id" yaml:"fee_id"`
	FeeContent FeeContent `json:"fee_content" yaml:"fee_content"`
}

var _ FeesMessage = MsgCreateFee{}

func (msg MsgCreateFee) Type() string  { return "create-fee" }
func (msg MsgCreateFee) Route() string { return RouterKey }
func (msg MsgCreateFee) ValidateBasic() sdk.Error {
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
	if !IsValidFeeId(msg.FeeId) {
		return ErrInvalidId(DefaultCodespace, "fee id invalid")
	}

	// Validate FeeContent
	if err := msg.FeeContent.Validate(); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreateFee) GetSenderDid() ixo.Did { return msg.CreatorDid }
func (msg MsgCreateFee) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgCreateFee) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreateFee) GetPubKey() string { return msg.PubKey }

func (msg MsgCreateFee) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

type MsgCreateFeeContract struct {
	SignBytes      string         `json:"signBytes" yaml:"signBytes"`
	PubKey         string         `json:"pub_key" yaml:"pub_key"`
	CreatorDid     ixo.Did        `json:"creator_did" yaml:"creator_did"`
	FeeId          string         `json:"fee_id" yaml:"fee_id"`
	FeeContractId  string         `json:"fee_contract_id" yaml:"fee_contract_id"`
	Payer          sdk.AccAddress `json:"payer" yaml:"payer"`
	CanDeauthorise bool           `json:"can_deauthorise" yaml:"can_deauthorise"`
	DiscountId     sdk.Uint       `json:"discount_id" yaml:"discount_id"`
}

var _ FeesMessage = MsgCreateFeeContract{}

func (msg MsgCreateFeeContract) Type() string  { return "create-fee-contract" }
func (msg MsgCreateFeeContract) Route() string { return RouterKey }
func (msg MsgCreateFeeContract) ValidateBasic() sdk.Error {
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
	if !IsValidFeeId(msg.FeeId) {
		return ErrInvalidId(DefaultCodespace, "fee id invalid")
	} else if !IsValidFeeContractId(msg.FeeContractId) {
		return ErrInvalidId(DefaultCodespace, "fee contract id invalid")
	}

	return nil
}

func (msg MsgCreateFeeContract) GetSenderDid() ixo.Did { return msg.CreatorDid }
func (msg MsgCreateFeeContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgCreateFeeContract) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreateFeeContract) GetPubKey() string { return msg.PubKey }

func (msg MsgCreateFeeContract) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

type MsgGrantFeeDiscount struct {
	SignBytes     string         `json:"signBytes" yaml:"signBytes"`
	PubKey        string         `json:"pub_key" yaml:"pub_key"`
	SenderDid     ixo.Did        `json:"sender_did" yaml:"sender_did"`
	FeeContractId string         `json:"fee_contract_id" yaml:"fee_contract_id"`
	DiscountId    sdk.Uint       `json:"discount_id" yaml:"discount_id"`
	Recipient     sdk.AccAddress `json:"recipient" yaml:"recipient"`
}

var _ FeesMessage = MsgGrantFeeDiscount{}

func (msg MsgGrantFeeDiscount) Type() string  { return "grant-fee-discount" }
func (msg MsgGrantFeeDiscount) Route() string { return RouterKey }
func (msg MsgGrantFeeDiscount) ValidateBasic() sdk.Error {
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
	if !IsValidFeeContractId(msg.FeeContractId) {
		return ErrInvalidId(DefaultCodespace, "fee contract id invalid")
	}

	return nil
}

func (msg MsgGrantFeeDiscount) GetSenderDid() ixo.Did { return msg.SenderDid }
func (msg MsgGrantFeeDiscount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgGrantFeeDiscount) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgGrantFeeDiscount) GetPubKey() string { return msg.PubKey }

func (msg MsgGrantFeeDiscount) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

type MsgRevokeFeeDiscount struct {
	SignBytes     string         `json:"signBytes" yaml:"signBytes"`
	PubKey        string         `json:"pub_key" yaml:"pub_key"`
	SenderDid     ixo.Did        `json:"sender_did" yaml:"sender_did"`
	FeeContractId string         `json:"fee_contract_id" yaml:"fee_contract_id"`
	Holder        sdk.AccAddress `json:"holder" yaml:"holder"`
}

var _ FeesMessage = MsgRevokeFeeDiscount{}

func (msg MsgRevokeFeeDiscount) Type() string  { return "revoke-fee-discount" }
func (msg MsgRevokeFeeDiscount) Route() string { return RouterKey }
func (msg MsgRevokeFeeDiscount) ValidateBasic() sdk.Error {
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
	if !IsValidFeeContractId(msg.FeeContractId) {
		return ErrInvalidId(DefaultCodespace, "fee contract id invalid")
	}

	return nil
}

func (msg MsgRevokeFeeDiscount) GetSenderDid() ixo.Did { return msg.SenderDid }
func (msg MsgRevokeFeeDiscount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgRevokeFeeDiscount) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgRevokeFeeDiscount) GetPubKey() string { return msg.PubKey }

func (msg MsgRevokeFeeDiscount) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

type MsgChargeFee struct {
	SignBytes     string  `json:"signBytes" yaml:"signBytes"`
	PubKey        string  `json:"pub_key" yaml:"pub_key"`
	SenderDid     ixo.Did `json:"sender_did" yaml:"sender_did"`
	FeeContractId string  `json:"fee_contract_id" yaml:"fee_contract_id"`
}

var _ FeesMessage = MsgChargeFee{}

func (msg MsgChargeFee) Type() string  { return "charge-fee" }
func (msg MsgChargeFee) Route() string { return RouterKey }
func (msg MsgChargeFee) ValidateBasic() sdk.Error {
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
	if !IsValidFeeContractId(msg.FeeContractId) {
		return ErrInvalidId(DefaultCodespace, "fee contract id invalid")
	}

	return nil
}

func (msg MsgChargeFee) GetSenderDid() ixo.Did { return msg.SenderDid }
func (msg MsgChargeFee) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg MsgChargeFee) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgChargeFee) GetPubKey() string { return msg.PubKey }

func (msg MsgChargeFee) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}
