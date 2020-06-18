package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
	"strings"
)

var (
	_ ixo.IxoMsg = MsgCreateBond{}
	_ ixo.IxoMsg = MsgEditBond{}
	_ ixo.IxoMsg = MsgBuy{}
	_ ixo.IxoMsg = MsgSell{}
	_ ixo.IxoMsg = MsgSwap{}
)

type MsgCreateBond struct {
	BondDid                ixo.Did        `json:"bond_did" yaml:"bond_did"`
	Token                  string         `json:"token" yaml:"token"`
	Name                   string         `json:"name" yaml:"name"`
	Description            string         `json:"description" yaml:"description"`
	FunctionType           string         `json:"function_type" yaml:"function_type"`
	FunctionParameters     FunctionParams `json:"function_parameters" yaml:"function_parameters"`
	CreatorDid             ixo.Did        `json:"creator_did" yaml:"creator_did"`
	CreatorPubKey          string         `json:"pub_key" yaml:"pub_key"`
	ReserveTokens          []string       `json:"reserve_tokens" yaml:"reserve_tokens"`
	TxFeePercentage        sdk.Dec        `json:"tx_fee_percentage" yaml:"tx_fee_percentage"`
	ExitFeePercentage      sdk.Dec        `json:"exit_fee_percentage" yaml:"exit_fee_percentage"`
	FeeAddress             sdk.AccAddress `json:"fee_address" yaml:"fee_address"`
	MaxSupply              sdk.Coin       `json:"max_supply" yaml:"max_supply"`
	OrderQuantityLimits    sdk.Coins      `json:"order_quantity_limits" yaml:"order_quantity_limits"`
	SanityRate             sdk.Dec        `json:"sanity_rate" yaml:"sanity_rate"`
	SanityMarginPercentage sdk.Dec        `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
	AllowSells             string         `json:"allow_sells" yaml:"allow_sells"`
	BatchBlocks            sdk.Uint       `json:"batch_blocks" yaml:"batch_blocks"`
}

func NewMsgCreateBond(token, name, description string, creatorDid sovrin.SovrinDid,
	functionType string, functionParameters FunctionParams, reserveTokens []string,
	txFeePercentage, exitFeePercentage sdk.Dec, feeAddress sdk.AccAddress, maxSupply sdk.Coin,
	orderQuantityLimits sdk.Coins, sanityRate, sanityMarginPercentage sdk.Dec,
	allowSell string, batchBlocks sdk.Uint, bondDid ixo.Did) MsgCreateBond {
	return MsgCreateBond{
		BondDid:                bondDid,
		Token:                  token,
		Name:                   name,
		Description:            description,
		CreatorDid:             creatorDid.Did,
		CreatorPubKey:          creatorDid.VerifyKey,
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
		AllowSells:             strings.ToLower(allowSell),
		BatchBlocks:            batchBlocks,
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
	} else if strings.TrimSpace(msg.CreatorPubKey) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "CreatorPubKey")
	} else if len(msg.ReserveTokens) == 0 {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Reserve token")
	} else if msg.FeeAddress.Empty() {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Fee address")
	} else if strings.TrimSpace(msg.FunctionType) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "Function type")
	} else if strings.TrimSpace(msg.AllowSells) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "AllowSells")
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
		return sdk.ErrInternal("max supply is invalid")
	} else if !msg.OrderQuantityLimits.IsValid() {
		return sdk.ErrInternal("order quantity limits are invalid")
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

	// Check that true or false
	if msg.AllowSells != TRUE && msg.AllowSells != FALSE {
		return ErrArgumentMissingOrNonBoolean(DefaultCodespace, "AllowSells")
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

	// Note: uniqueness of reserve tokens checked when parsing

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !ixo.IsValidDid(msg.CreatorDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "creator did is invalid")
	}

	return nil
}

func (msg MsgCreateBond) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgCreateBond) GetSignerDid() ixo.Did {
	return msg.CreatorDid
}

func (msg MsgCreateBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgCreateBond) Route() string { return RouterKey }

func (msg MsgCreateBond) Type() string { return "create_bond" }

type MsgEditBond struct {
	BondDid                ixo.Did `json:"bond_did" yaml:"bond_did"`
	Token                  string  `json:"token" yaml:"token"`
	Name                   string  `json:"name" yaml:"name"`
	Description            string  `json:"description" yaml:"description"`
	OrderQuantityLimits    string  `json:"order_quantity_limits" yaml:"order_quantity_limits"`
	SanityRate             string  `json:"sanity_rate" yaml:"sanity_rate"`
	SanityMarginPercentage string  `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
	EditorDid              ixo.Did `json:"editor_did" yaml:"editor_did"`
	EditorPubKey           string  `json:"pub_key" yaml:"pub_key"`
}

func NewMsgEditBond(token, name, description, orderQuantityLimits, sanityRate,
	sanityMarginPercentage string, editorDid sovrin.SovrinDid, bondDid ixo.Did) MsgEditBond {
	return MsgEditBond{
		BondDid:                bondDid,
		Token:                  token,
		Name:                   name,
		Description:            description,
		OrderQuantityLimits:    orderQuantityLimits,
		SanityRate:             sanityRate,
		SanityMarginPercentage: sanityMarginPercentage,
		EditorDid:              editorDid.Did,
		EditorPubKey:           editorDid.VerifyKey,
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
	} else if strings.TrimSpace(msg.EditorPubKey) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "EditorPubKey")
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
	if !ixo.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !ixo.IsValidDid(msg.EditorDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "editor did is invalid")
	}

	return nil
}

func (msg MsgEditBond) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgEditBond) GetSignerDid() ixo.Did {
	return msg.EditorDid
}

func (msg MsgEditBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgEditBond) Route() string { return RouterKey }

func (msg MsgEditBond) Type() string { return "edit_bond" }

type MsgBuy struct {
	BuyerDid  ixo.Did   `json:"buyer_did" yaml:"buyer_did"`
	PubKey    string    `json:"pub_key" yaml:"pub_key"`
	Amount    sdk.Coin  `json:"amount" yaml:"amount"`
	MaxPrices sdk.Coins `json:"max_prices" yaml:"max_prices"`
	BondDid   ixo.Did   `json:"bond_did" yaml:"bond_did"`
}

func NewMsgBuy(buyerDid sovrin.SovrinDid, amount sdk.Coin, maxPrices sdk.Coins,
	bondDid ixo.Did) MsgBuy {
	return MsgBuy{
		BuyerDid:  buyerDid.Did,
		PubKey:    buyerDid.VerifyKey,
		Amount:    amount,
		MaxPrices: maxPrices,
		BondDid:   bondDid,
	}
}

func (msg MsgBuy) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.BuyerDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BuyerDid")
	} else if strings.TrimSpace(msg.PubKey) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "PubKey")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	}

	// Check that amount valid and non zero
	if !msg.Amount.IsValid() {
		return sdk.ErrInternal("amount is invalid")
	} else if msg.Amount.Amount.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "Amount")
	}

	// Check that maxPrices valid
	if !msg.MaxPrices.IsValid() {
		return sdk.ErrInternal("maxprices is invalid")
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !ixo.IsValidDid(msg.BuyerDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "buyer did is invalid")
	}

	return nil
}

func (msg MsgBuy) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgBuy) GetSignerDid() ixo.Did {
	return msg.BuyerDid
}

func (msg MsgBuy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgBuy) Route() string { return RouterKey }

func (msg MsgBuy) Type() string { return "buy" }

type MsgSell struct {
	SellerDid ixo.Did  `json:"seller_did" yaml:"seller_did"`
	PubKey    string   `json:"pub_key" yaml:"pub_key"`
	Amount    sdk.Coin `json:"amount" yaml:"amount"`
	BondDid   ixo.Did  `json:"bond_did" yaml:"bond_did"`
}

func NewMsgSell(sellerDid sovrin.SovrinDid, amount sdk.Coin, bondDid ixo.Did) MsgSell {
	return MsgSell{
		SellerDid: sellerDid.Did,
		PubKey:    sellerDid.VerifyKey,
		Amount:    amount,
		BondDid:   bondDid,
	}
}

func (msg MsgSell) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.SellerDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "SellerDid")
	} else if strings.TrimSpace(msg.PubKey) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "PubKey")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	}

	// Check that amount valid and non zero
	if !msg.Amount.IsValid() {
		return sdk.ErrInternal("amount is invalid")
	} else if msg.Amount.Amount.IsZero() {
		return ErrArgumentMustBePositive(DefaultCodespace, "Amount")
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !ixo.IsValidDid(msg.SellerDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "seller did is invalid")
	}

	return nil
}

func (msg MsgSell) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgSell) GetSignerDid() ixo.Did {
	return msg.SellerDid
}

func (msg MsgSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgSell) Route() string { return RouterKey }

func (msg MsgSell) Type() string { return "sell" }

type MsgSwap struct {
	SwapperDid ixo.Did  `json:"swapper_did" yaml:"swapper_did"`
	PubKey     string   `json:"pub_key" yaml:"pub_key"`
	BondDid    ixo.Did  `json:"bond_did" yaml:"bond_did"`
	From       sdk.Coin `json:"from" yaml:"from"`
	ToToken    string   `json:"to_token" yaml:"to_token"`
}

func NewMsgSwap(swapperDid sovrin.SovrinDid, from sdk.Coin, toToken string,
	bondDid ixo.Did) MsgSwap {
	return MsgSwap{
		SwapperDid: swapperDid.Did,
		PubKey:     swapperDid.VerifyKey,
		From:       from,
		ToToken:    toToken,
		BondDid:    bondDid,
	}
}

func (msg MsgSwap) ValidateBasic() sdk.Error {
	// Check if empty
	if strings.TrimSpace(msg.SwapperDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "SwapperDid")
	} else if strings.TrimSpace(msg.PubKey) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "PubKey")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "BondDid")
	} else if strings.TrimSpace(msg.ToToken) == "" {
		return ErrArgumentCannotBeEmpty(DefaultCodespace, "ToToken")
	}

	// Validate from amount
	if !msg.From.IsValid() {
		return sdk.ErrInternal("from amount is invalid")
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
	if !ixo.IsValidDid(msg.BondDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "bond did is invalid")
	} else if !ixo.IsValidDid(msg.SwapperDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "swapper did is invalid")
	}

	return nil
}

func (msg MsgSwap) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgSwap) GetSignerDid() ixo.Did {
	return msg.SwapperDid
}

func (msg MsgSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgSwap) Route() string { return RouterKey }

func (msg MsgSwap) Type() string { return "swap" }
