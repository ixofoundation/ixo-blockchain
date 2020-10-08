package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"strings"
)

const (
	TypeMsgCreateBond         = "create_bond"
	TypeMsgEditBond           = "edit_bond"
	TypeMsgBuy                = "buy"
	TypeMsgSell               = "sell"
	TypeMsgSwap               = "swap"
	TypeMsgMakeOutcomePayment = "make_outcome_payment"
	TypeMsgWithdrawShare      = "withdraw_share"
)

var (
	_ ixo.IxoMsg = MsgCreateBond{}
	_ ixo.IxoMsg = MsgEditBond{}
	_ ixo.IxoMsg = MsgBuy{}
	_ ixo.IxoMsg = MsgSell{}
	_ ixo.IxoMsg = MsgSwap{}
)

type MsgCreateBond struct {
	BondDid                did.Did        `json:"bond_did" yaml:"bond_did"`
	Token                  string         `json:"token" yaml:"token"`
	Name                   string         `json:"name" yaml:"name"`
	Description            string         `json:"description" yaml:"description"`
	FunctionType           string         `json:"function_type" yaml:"function_type"`
	FunctionParameters     FunctionParams `json:"function_parameters" yaml:"function_parameters"`
	CreatorDid             did.Did        `json:"creator_did" yaml:"creator_did"`
	ReserveTokens          []string       `json:"reserve_tokens" yaml:"reserve_tokens"`
	TxFeePercentage        sdk.Dec        `json:"tx_fee_percentage" yaml:"tx_fee_percentage"`
	ExitFeePercentage      sdk.Dec        `json:"exit_fee_percentage" yaml:"exit_fee_percentage"`
	FeeAddress             sdk.AccAddress `json:"fee_address" yaml:"fee_address"`
	MaxSupply              sdk.Coin       `json:"max_supply" yaml:"max_supply"`
	OrderQuantityLimits    sdk.Coins      `json:"order_quantity_limits" yaml:"order_quantity_limits"`
	SanityRate             sdk.Dec        `json:"sanity_rate" yaml:"sanity_rate"`
	SanityMarginPercentage sdk.Dec        `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
	AllowSells             bool           `json:"allow_sells" yaml:"allow_sells"`
	BatchBlocks            sdk.Uint       `json:"batch_blocks" yaml:"batch_blocks"`
	OutcomePayment         sdk.Int        `json:"outcome_payment" yaml:"outcome_payment"`
}

func NewMsgCreateBond(token, name, description string, creatorDid did.Did,
	functionType string, functionParameters FunctionParams, reserveTokens []string,
	txFeePercentage, exitFeePercentage sdk.Dec, feeAddress sdk.AccAddress, maxSupply sdk.Coin,
	orderQuantityLimits sdk.Coins, sanityRate, sanityMarginPercentage sdk.Dec,
	allowSell bool, batchBlocks sdk.Uint, outcomePayment sdk.Int, bondDid did.Did) MsgCreateBond {
	return MsgCreateBond{
		BondDid:                bondDid,
		Token:                  token,
		Name:                   name,
		Description:            description,
		CreatorDid:             creatorDid,
		FunctionType:           functionType,
		FunctionParameters:     functionParameters,
		ReserveTokens:          reserveTokens,
		TxFeePercentage:        txFeePercentage,
		ExitFeePercentage:      exitFeePercentage,
		FeeAddress:             feeAddress,
		MaxSupply:              maxSupply,
		OrderQuantityLimits:    orderQuantityLimits,
		SanityRate:             sanityRate,
		SanityMarginPercentage: sanityMarginPercentage,
		AllowSells:             allowSell,
		BatchBlocks:            batchBlocks,
		OutcomePayment:         outcomePayment,
	}
}

func (msg MsgCreateBond) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	} else if strings.TrimSpace(msg.Token) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Token")
	} else if strings.TrimSpace(msg.Name) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Name")
	} else if strings.TrimSpace(msg.Description) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Description")
	} else if strings.TrimSpace(msg.CreatorDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "CreatorDid")
	} else if len(msg.ReserveTokens) == 0 {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Reserve token")
	} else if msg.FeeAddress.Empty() {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Fee address")
	} else if strings.TrimSpace(msg.FunctionType) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Function type")
	}
	// Note: FunctionParameters can be empty

	// Check that bond token is a valid token name
	err := CheckCoinDenom(msg.Token)
	if err != nil {
		return ErrInvalidCoinDenomination(DefaultCodespace, msg.Token)
	}

	// Validate function parameters
	if err := msg.FunctionParameters.Validate(msg.FunctionType); err != nil {
		return err
	}

	// Validate reserve tokens
	if err = CheckReserveTokenNames(msg.ReserveTokens, msg.Token); err != nil {
		return err
	} else if err = CheckNoOfReserveTokens(msg.ReserveTokens, msg.FunctionType); err != nil {
		return err
	}

	// Validate coins
	if !msg.MaxSupply.IsValid() {
		return sdk.ErrInvalidCoins("max supply is invalid")
	} else if !msg.OrderQuantityLimits.IsValid() {
		return sdk.ErrInvalidCoins("order quantity limits are invalid")
	}

	// Check that max supply denom matches token denom
	if msg.MaxSupply.Denom != msg.Token {
		return ErrMaxSupplyDenomDoesNotMatchTokenDenom(DefaultCodespace)
	}

	// Check that Sanity values not negative
	if msg.SanityRate.IsNegative() {
		return ErrArgumentCannotBeNegative(DefaultCodespace, "SanityRate")
	} else if msg.SanityMarginPercentage.IsNegative() {
		return ErrArgumentCannotBeNegative(DefaultCodespace, "SanityMarginPercentage")
	}

	// Check FeePercentages not negative and don't add up to 100
	if msg.TxFeePercentage.IsNegative() {
		return ErrArgumentCannotBeNegative(DefaultCodespace, "TxFeePercentage")
	} else if msg.ExitFeePercentage.IsNegative() {
		return ErrArgumentCannotBeNegative(DefaultCodespace, "ExitFeePercentage")
	} else if msg.TxFeePercentage.Add(msg.ExitFeePercentage).GTE(sdk.NewDec(100)) {
		return ErrFeesCannotBeOrExceed100Percent(DefaultCodespace)
	}

	// Check that not zero
	if msg.BatchBlocks.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "BatchBlocks")
	} else if msg.MaxSupply.Amount.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "MaxSupply")
	}

	// Check that outcome payment not negative
	if msg.OutcomePayment.IsNegative() {
		return ErrArgumentMustBePositive(DefaultCodespace, "OutcomePayment")
	}

	// Note: uniqueness of reserve tokens checked when parsing

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !did.IsValidDid(msg.CreatorDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "creator did is invalid")
	}

	return nil
}

func (msg MsgCreateBond) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateBond) GetSignerDid() did.Did { return msg.CreatorDid }
func (msg MsgCreateBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreateBond) Route() string { return RouterKey }

func (msg MsgCreateBond) Type() string { return TypeMsgCreateBond }

type MsgEditBond struct {
	BondDid                did.Did `json:"bond_did" yaml:"bond_did"`
	Token                  string  `json:"token" yaml:"token"`
	Name                   string  `json:"name" yaml:"name"`
	Description            string  `json:"description" yaml:"description"`
	OrderQuantityLimits    string  `json:"order_quantity_limits" yaml:"order_quantity_limits"`
	SanityRate             string  `json:"sanity_rate" yaml:"sanity_rate"`
	SanityMarginPercentage string  `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
	EditorDid              did.Did `json:"editor_did" yaml:"editor_did"`
}

func NewMsgEditBond(token, name, description, orderQuantityLimits, sanityRate,
	sanityMarginPercentage string, editorDid, bondDid did.Did) MsgEditBond {
	return MsgEditBond{
		BondDid:                bondDid,
		Token:                  token,
		Name:                   name,
		Description:            description,
		OrderQuantityLimits:    orderQuantityLimits,
		SanityRate:             sanityRate,
		SanityMarginPercentage: sanityMarginPercentage,
		EditorDid:              editorDid,
	}
}

func (msg MsgEditBond) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	} else if strings.TrimSpace(msg.Token) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Token")
	} else if strings.TrimSpace(msg.Name) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Name")
	} else if strings.TrimSpace(msg.Description) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Description")
	} else if strings.TrimSpace(msg.SanityRate) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "SanityRate")
	} else if strings.TrimSpace(msg.SanityMarginPercentage) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "SanityMarginPercentage")
	} else if strings.TrimSpace(msg.EditorDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "EditorDid")
	}
	// Note: order quantity limits can be blank

	// Check that at least one editable was edited. Fields that will not
	// be edited should be "DoNotModifyField", and not an empty string
	inputList := []string{
		msg.Name, msg.Description, msg.OrderQuantityLimits,
		msg.SanityRate, msg.SanityMarginPercentage,
	}
	atLeaseOneEdit := false
	for _, e := range inputList {
		if e != DoNotModifyField {
			atLeaseOneEdit = true
			break
		}
	}
	if !atLeaseOneEdit {
		return ErrDidNotEditAnything(DefaultCodespace)
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !did.IsValidDid(msg.EditorDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "editor did is invalid")
	}

	return nil
}

func (msg MsgEditBond) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgEditBond) GetSignerDid() did.Did { return msg.EditorDid }
func (msg MsgEditBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgEditBond) Route() string { return RouterKey }

func (msg MsgEditBond) Type() string { return TypeMsgEditBond }

type MsgBuy struct {
	BuyerDid  did.Did   `json:"buyer_did" yaml:"buyer_did"`
	Amount    sdk.Coin  `json:"amount" yaml:"amount"`
	MaxPrices sdk.Coins `json:"max_prices" yaml:"max_prices"`
	BondDid   did.Did   `json:"bond_did" yaml:"bond_did"`
}

func NewMsgBuy(buyerDid did.Did, amount sdk.Coin, maxPrices sdk.Coins,
	bondDid did.Did) MsgBuy {
	return MsgBuy{
		BuyerDid:  buyerDid,
		Amount:    amount,
		MaxPrices: maxPrices,
		BondDid:   bondDid,
	}
}

func (msg MsgBuy) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.BuyerDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BuyerDid")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	}

	// Check that amount valid and non zero
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("amount is invalid")
	} else if msg.Amount.Amount.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "Amount")
	}

	// Check that maxPrices valid
	if !msg.MaxPrices.IsValid() {
		return sdk.ErrInvalidCoins("maxprices is invalid")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !did.IsValidDid(msg.BuyerDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "buyer did is invalid")
	}

	return nil
}

func (msg MsgBuy) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgBuy) GetSignerDid() did.Did { return msg.BuyerDid }
func (msg MsgBuy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgBuy) Route() string { return RouterKey }

func (msg MsgBuy) Type() string { return TypeMsgBuy }

type MsgSell struct {
	SellerDid did.Did  `json:"seller_did" yaml:"seller_did"`
	Amount    sdk.Coin `json:"amount" yaml:"amount"`
	BondDid   did.Did  `json:"bond_did" yaml:"bond_did"`
}

func NewMsgSell(sellerDid did.Did, amount sdk.Coin, bondDid did.Did) MsgSell {
	return MsgSell{
		SellerDid: sellerDid,
		Amount:    amount,
		BondDid:   bondDid,
	}
}

func (msg MsgSell) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.SellerDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "SellerDid")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	}

	// Check that amount valid and non zero
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("amount is invalid")
	} else if msg.Amount.Amount.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "Amount")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !did.IsValidDid(msg.SellerDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "seller did is invalid")
	}

	return nil
}

func (msg MsgSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSell) GetSignerDid() did.Did { return msg.SellerDid }
func (msg MsgSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgSell) Route() string { return RouterKey }

func (msg MsgSell) Type() string { return TypeMsgSell }

type MsgSwap struct {
	SwapperDid did.Did  `json:"swapper_did" yaml:"swapper_did"`
	BondDid    did.Did  `json:"bond_did" yaml:"bond_did"`
	From       sdk.Coin `json:"from" yaml:"from"`
	ToToken    string   `json:"to_token" yaml:"to_token"`
}

func NewMsgSwap(swapperDid did.Did, from sdk.Coin, toToken string,
	bondDid did.Did) MsgSwap {
	return MsgSwap{
		SwapperDid: swapperDid,
		From:       from,
		ToToken:    toToken,
		BondDid:    bondDid,
	}
}

func (msg MsgSwap) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.SwapperDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "SwapperDid")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	} else if strings.TrimSpace(msg.ToToken) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "ToToken")
	}

	// Validate from amount
	if !msg.From.IsValid() {
		return sdk.ErrInvalidCoins("from amount is invalid")
	}

	// Validate to token
	err := CheckCoinDenom(msg.ToToken)
	if err != nil {
		return err
	}

	// Check if from and to the same token
	if msg.From.Denom == msg.ToToken {
		return ErrFromAndToCannotBeTheSameToken(DefaultCodespace)
	}

	// Check that non zero
	if msg.From.Amount.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "FromAmount")
	}

	// Note: From denom and amount must be valid since sdk.Coin

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !did.IsValidDid(msg.SwapperDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "swapper did is invalid")
	}

	return nil
}

func (msg MsgSwap) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSwap) GetSignerDid() did.Did { return msg.SwapperDid }
func (msg MsgSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgSwap) Route() string { return RouterKey }

func (msg MsgSwap) Type() string { return TypeMsgSwap }

type MsgMakeOutcomePayment struct {
	SenderDid did.Did `json:"sender_did" yaml:"sender_did"`
	BondDid   did.Did `json:"bond_did" yaml:"bond_did"`
}

func NewMsgMakeOutcomePayment(senderDid, bondDid did.Did) MsgMakeOutcomePayment {
	return MsgMakeOutcomePayment{
		SenderDid: senderDid,
		BondDid:   bondDid,
	}
}

func (msg MsgMakeOutcomePayment) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.SenderDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "SenderDid")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !did.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	return nil
}

func (msg MsgMakeOutcomePayment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgMakeOutcomePayment) GetSignerDid() did.Did { return msg.SenderDid }
func (msg MsgMakeOutcomePayment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgMakeOutcomePayment) Route() string { return RouterKey }

func (msg MsgMakeOutcomePayment) Type() string { return TypeMsgMakeOutcomePayment }

type MsgWithdrawShare struct {
	RecipientDid did.Did `json:"recipient_did" yaml:"recipient_did"`
	BondDid      did.Did `json:"bond_did" yaml:"bond_did"`
}

func NewMsgWithdrawShare(recipientDid, bondDid did.Did) MsgWithdrawShare {
	return MsgWithdrawShare{
		RecipientDid: recipientDid,
		BondDid:      bondDid,
	}
}

func (msg MsgWithdrawShare) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.RecipientDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "RecipientDid")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !did.IsValidDid(msg.RecipientDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "recipient did is invalid")
	}

	return nil
}

func (msg MsgWithdrawShare) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgWithdrawShare) GetSignerDid() did.Did { return msg.RecipientDid }
func (msg MsgWithdrawShare) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgWithdrawShare) Route() string { return RouterKey }

func (msg MsgWithdrawShare) Type() string { return TypeMsgWithdrawShare }
