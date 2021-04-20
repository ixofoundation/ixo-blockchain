package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	ixotypes "github.com/ixofoundation/ixo-blockchain/x/ixo/types"
	"strings"
)

const (
	TypeMsgCreateBond         = "create_bond"
	TypeMsgEditBond           = "edit_bond"
	TypeMsgSetNextAlpha       = "set_next_alpha"
	TypeMsgUpdateBondState    = "update_bond_state"
	TypeMsgBuy                = "buy"
	TypeMsgSell               = "sell"
	TypeMsgSwap               = "swap"
	TypeMsgMakeOutcomePayment = "make_outcome_payment"
	TypeMsgWithdrawShare      = "withdraw_share"
)

var (
	_ ixotypes.IxoMsg = &MsgCreateBond{}
	_ ixotypes.IxoMsg = &MsgEditBond{}
	_ ixotypes.IxoMsg = &MsgSetNextAlpha{}
	_ ixotypes.IxoMsg = &MsgBuy{}
	_ ixotypes.IxoMsg = &MsgSell{}
	_ ixotypes.IxoMsg = &MsgSwap{}
)

//type MsgCreateBond struct {
//	BondDid                did.Did        `json:"bond_did" yaml:"bond_did"`
//	Token                  string         `json:"token" yaml:"token"`
//	Name                   string         `json:"name" yaml:"name"`
//	Description            string         `json:"description" yaml:"description"`
//	FunctionType           string         `json:"function_type" yaml:"function_type"`
//	FunctionParameters     FunctionParams `json:"function_parameters" yaml:"function_parameters"`
//	CreatorDid             did.Did        `json:"creator_did" yaml:"creator_did"`
//	ControllerDid          did.Did        `json:"controller_did" yaml:"controller_did"`
//	ReserveTokens          []string       `json:"reserve_tokens" yaml:"reserve_tokens"`
//	TxFeePercentage        sdk.Dec        `json:"tx_fee_percentage" yaml:"tx_fee_percentage"`
//	ExitFeePercentage      sdk.Dec        `json:"exit_fee_percentage" yaml:"exit_fee_percentage"`
//	FeeAddress             sdk.AccAddress `json:"fee_address" yaml:"fee_address"`
//	MaxSupply              sdk.Coin       `json:"max_supply" yaml:"max_supply"`
//	OrderQuantityLimits    sdk.Coins      `json:"order_quantity_limits" yaml:"order_quantity_limits"`
//	SanityRate             sdk.Dec        `json:"sanity_rate" yaml:"sanity_rate"`
//	SanityMarginPercentage sdk.Dec        `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
//	AllowSells             bool           `json:"allow_sells" yaml:"allow_sells"`
//	AlphaBond              bool           `json:"alpha_bond" yaml:"alpha_bond"`
//	BatchBlocks            sdk.Uint       `json:"batch_blocks" yaml:"batch_blocks"`
//	OutcomePayment         sdk.Int        `json:"outcome_payment" yaml:"outcome_payment"`
//}

func NewMsgCreateBond(token, name, description string, creatorDid, controllerDid did.Did,
	functionType string, functionParameters FunctionParams, reserveTokens []string,
	txFeePercentage, exitFeePercentage sdk.Dec, feeAddress sdk.AccAddress, maxSupply sdk.Coin,
	orderQuantityLimits sdk.Coins, sanityRate, sanityMarginPercentage sdk.Dec,
	allowSell, alphaBond bool, batchBlocks sdk.Uint, outcomePayment sdk.Int, bondDid did.Did) *MsgCreateBond {

	return &MsgCreateBond{
		BondDid:                bondDid,
		Token:                  token,
		Name:                   name,
		Description:            description,
		CreatorDid:             creatorDid,
		ControllerDid:          controllerDid,
		FunctionType:           functionType,
		FunctionParameters:     functionParameters,
		ReserveTokens:          reserveTokens,
		TxFeePercentage:        txFeePercentage,
		ExitFeePercentage:      exitFeePercentage,
		FeeAddress: 			feeAddress.String(),
		MaxSupply:              maxSupply,
		OrderQuantityLimits:    orderQuantityLimits,
		SanityRate:             sanityRate,
		SanityMarginPercentage: sanityMarginPercentage,
		AllowSells:             allowSell,
		AlphaBond:              alphaBond,
		BatchBlocks:            batchBlocks,
		OutcomePayment:         outcomePayment,
	}
}

func (msg MsgCreateBond) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	} else if strings.TrimSpace(msg.Token) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "token")
	} else if strings.TrimSpace(msg.Name) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "name")
	} else if strings.TrimSpace(msg.Description) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "description")
	} else if strings.TrimSpace(msg.CreatorDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "creator DID")
	} else if strings.TrimSpace(msg.ControllerDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "controller DID")
	} else if len(msg.ReserveTokens) == 0 {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "reserve tokens")
	} else if strings.TrimSpace(msg.FeeAddress) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "fee address")
	} else if strings.TrimSpace(msg.FunctionType) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "function type")
	}
	// Note: FunctionParameters can be empty

	// Check that bond token is a valid token name
	err := CheckCoinDenom(msg.Token)
	if err != nil {
		return err
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
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "max supply")
	} else if !msg.OrderQuantityLimits.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "order quantity limits")
	}

	// Check that max supply denom matches token denom
	if msg.MaxSupply.Denom != msg.Token {
		return ErrMaxSupplyDenomDoesNotMatchTokenDenom
	}

	// Check that Sanity values not negative
	if msg.SanityRate.IsNegative() {
		return sdkerrors.Wrap(ErrArgumentCannotBeNegative, "sanity rate")
	} else if msg.SanityMarginPercentage.IsNegative() {
		return sdkerrors.Wrap(ErrArgumentCannotBeNegative, "sanity margin percentage")
	}

	// Check FeePercentages not negative and don't add up to 100
	if msg.TxFeePercentage.IsNegative() {
		return sdkerrors.Wrap(ErrArgumentCannotBeNegative, "tx fee percentage")
	} else if msg.ExitFeePercentage.IsNegative() {
		return sdkerrors.Wrap(ErrArgumentCannotBeNegative, "exit fee percentage")
	} else if msg.TxFeePercentage.Add(msg.ExitFeePercentage).GTE(sdk.NewDec(100)) {
		return ErrFeesCannotBeOrExceed100Percent
	}

	// Check that not zero
	if msg.BatchBlocks.IsZero() {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "batch blocks")
	} else if msg.MaxSupply.Amount.IsZero() {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "max supply")
	}

	// Alpha bonds have to be augmented bonding curves
	if msg.AlphaBond && msg.FunctionType != AugmentedFunction {
		return sdkerrors.Wrap(ErrFunctionNotAvailableForFunctionType,
			"only augmented bonding curves can be alpha bonds")
	}

	// Check that outcome payment not negative
	if msg.OutcomePayment.IsNegative() {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "outcome payment")
	}

	// Note: uniqueness of reserve tokens checked when parsing

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "bond DID")
	} else if !did.IsValidDid(msg.CreatorDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "creator DID")
	}

	return nil
}

func (msg MsgCreateBond) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateBond) GetSignerDid() did.Did { return msg.CreatorDid }
func (msg MsgCreateBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreateBond) Route() string { return RouterKey }

func (msg MsgCreateBond) Type() string { return TypeMsgCreateBond }

//type MsgEditBond struct {
//	BondDid                did.Did `json:"bond_did" yaml:"bond_did"`
//	Name                   string  `json:"name" yaml:"name"`
//	Description            string  `json:"description" yaml:"description"`
//	OrderQuantityLimits    string  `json:"order_quantity_limits" yaml:"order_quantity_limits"`
//	SanityRate             string  `json:"sanity_rate" yaml:"sanity_rate"`
//	SanityMarginPercentage string  `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
//	EditorDid              did.Did `json:"editor_did" yaml:"editor_did"`
//}

func NewMsgEditBond(name, description, orderQuantityLimits, sanityRate,
	sanityMarginPercentage string, editorDid, bondDid did.Did) *MsgEditBond {
	return &MsgEditBond{
		BondDid:                bondDid,
		Name:                   name,
		Description:            description,
		OrderQuantityLimits:    orderQuantityLimits,
		SanityRate:             sanityRate,
		SanityMarginPercentage: sanityMarginPercentage,
		EditorDid:              editorDid,
	}
}

func (msg MsgEditBond) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	} else if strings.TrimSpace(msg.Name) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "name")
	} else if strings.TrimSpace(msg.Description) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "description")
	} else if strings.TrimSpace(msg.SanityRate) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "sanity rate")
	} else if strings.TrimSpace(msg.SanityMarginPercentage) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "sanity margin percentage")
	} else if strings.TrimSpace(msg.EditorDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "editor DID")
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
		return ErrDidNotEditAnything
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "bond DID")
	} else if !did.IsValidDid(msg.EditorDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "editor DID")
	}

	return nil
}

func (msg MsgEditBond) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgEditBond) GetSignerDid() did.Did { return msg.EditorDid }
func (msg MsgEditBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgEditBond) Route() string { return RouterKey }

func (msg MsgEditBond) Type() string { return TypeMsgEditBond }

//type MsgSetNextAlpha struct {
//	BondDid   did.Did `json:"bond_did" yaml:"bond_did"`
//	Alpha     sdk.Dec `json:"alpha" yaml:"alpha"`
//	EditorDid did.Did `json:"editor_did" yaml:"editor_did"`
//}

func NewMsgSetNextAlpha(alpha sdk.Dec, editorDid, bondDid did.Did) *MsgSetNextAlpha {
	return &MsgSetNextAlpha{
		BondDid:   bondDid,
		Alpha:     alpha,
		EditorDid: editorDid,
	}
}

func (msg MsgSetNextAlpha) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "BondDid")
	} else if strings.TrimSpace(msg.EditorDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "EditorDid")
	}

	// Check that 0.0001 <= alpha <= 0.9999. Note that we cannot set public
	// alpha to 0 or 1, because these are edge cases which cause the system
	// alpha to get stuck if we try to change the value of public alpha again.
	minNextAlpha := sdk.MustNewDecFromStr("0.0001")
	maxNextAlpha := sdk.MustNewDecFromStr("0.9999")
	if msg.Alpha.LT(minNextAlpha) || msg.Alpha.GT(maxNextAlpha) {
		return sdkerrors.Wrap(ErrInvalidAlpha, "0.0001 <= alpha <= 0.9999")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "bond did is invalid")
	} else if !did.IsValidDid(msg.EditorDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "editor did is invalid")
	}

	return nil
}

func (msg MsgSetNextAlpha) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgSetNextAlpha) GetSignerDid() did.Did { return msg.EditorDid }
func (msg MsgSetNextAlpha) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgSetNextAlpha) Route() string { return RouterKey }

func (msg MsgSetNextAlpha) Type() string { return TypeMsgSetNextAlpha }

//type MsgUpdateBondState struct {
//	BondDid   did.Did   `json:"bond_did" yaml:"bond_did"`
//	State     BondState `json:"state" yaml:"state"`
//	EditorDid did.Did   `json:"editor_did" yaml:"editor_did"`
//}

func NewMsgUpdateBondState(state BondState, editorDid, bondDid did.Did) *MsgUpdateBondState {
	return &MsgUpdateBondState{
		BondDid:   bondDid,
		State:     state.String(),
		EditorDid: editorDid,
	}
}

func (msg MsgUpdateBondState) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "BondDid")
	} else if strings.TrimSpace(msg.EditorDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "EditorDid")
	}

	// Bond status can only be updated to SETTLE or FAILED
	if msg.State != SettleState.String() && msg.State != FailedState.String() {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "cannot transition to that state")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "bond did is invalid")
	} else if !did.IsValidDid(msg.EditorDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "editor did is invalid")
	}

	return nil
}

func (msg MsgUpdateBondState) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUpdateBondState) GetSignerDid() did.Did { return msg.EditorDid }
func (msg MsgUpdateBondState) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgUpdateBondState) Route() string { return RouterKey }

func (msg MsgUpdateBondState) Type() string { return TypeMsgUpdateBondState }

//type MsgBuy struct {
//	BuyerDid  did.Did   `json:"buyer_did" yaml:"buyer_did"`
//	Amount    sdk.Coin  `json:"amount" yaml:"amount"`
//	MaxPrices sdk.Coins `json:"max_prices" yaml:"max_prices"`
//	BondDid   did.Did   `json:"bond_did" yaml:"bond_did"`
//}

func NewMsgBuy(buyerDid did.Did, amount sdk.Coin, maxPrices sdk.Coins,
	bondDid did.Did) *MsgBuy {
	return &MsgBuy{
		BuyerDid:  buyerDid,
		Amount:    amount,
		MaxPrices: maxPrices,
		BondDid:   bondDid,
	}
}

func (msg MsgBuy) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BuyerDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "buyer DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Check that amount valid and non zero
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	} else if msg.Amount.Amount.IsZero() {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "amount")
	}

	// Check that maxPrices valid
	if !msg.MaxPrices.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "max prices")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "bond DID")
	} else if !did.IsValidDid(msg.BuyerDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "buyer DID")
	}

	return nil
}

func (msg MsgBuy) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgBuy) GetSignerDid() did.Did { return msg.BuyerDid }
func (msg MsgBuy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgBuy) Route() string { return RouterKey }

func (msg MsgBuy) Type() string { return TypeMsgBuy }

//type MsgSell struct {
//	SellerDid did.Did  `json:"seller_did" yaml:"seller_did"`
//	Amount    sdk.Coin `json:"amount" yaml:"amount"`
//	BondDid   did.Did  `json:"bond_did" yaml:"bond_did"`
//}

func NewMsgSell(sellerDid did.Did, amount sdk.Coin, bondDid did.Did) *MsgSell {
	return &MsgSell{
		SellerDid: sellerDid,
		Amount:    amount,
		BondDid:   bondDid,
	}
}

func (msg MsgSell) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.SellerDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "seller DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Check that amount valid and non zero
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	} else if msg.Amount.Amount.IsZero() {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "amount")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "bond DID")
	} else if !did.IsValidDid(msg.SellerDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "seller DID")
	}

	return nil
}

func (msg MsgSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgSell) GetSignerDid() did.Did { return msg.SellerDid }
func (msg MsgSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgSell) Route() string { return RouterKey }

func (msg MsgSell) Type() string { return TypeMsgSell }

//type MsgSwap struct {
//	SwapperDid did.Did  `json:"swapper_did" yaml:"swapper_did"`
//	BondDid    did.Did  `json:"bond_did" yaml:"bond_did"`
//	From       sdk.Coin `json:"from" yaml:"from"`
//	ToToken    string   `json:"to_token" yaml:"to_token"`
//}

func NewMsgSwap(swapperDid did.Did, from sdk.Coin, toToken string,
	bondDid did.Did) *MsgSwap {
	return &MsgSwap{
		SwapperDid: swapperDid,
		From:       from,
		ToToken:    toToken,
		BondDid:    bondDid,
	}
}

func (msg MsgSwap) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.SwapperDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "swapper DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	} else if strings.TrimSpace(msg.ToToken) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "to token")
	}

	// Validate from amount
	if !msg.From.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.From.String())
	}

	// Validate to token
	err := CheckCoinDenom(msg.ToToken)
	if err != nil {
		return err
	}

	// Check if from and to the same token
	if msg.From.Denom == msg.ToToken {
		return ErrFromAndToCannotBeTheSameToken
	}

	// Check that non zero
	if msg.From.Amount.IsZero() {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "amount")
	}

	// Note: From denom and amount must be valid since sdk.Coin

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "bond DID is invalid")
	} else if !did.IsValidDid(msg.SwapperDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "swapper DID is invalid")
	}

	return nil
}

func (msg MsgSwap) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgSwap) GetSignerDid() did.Did { return msg.SwapperDid }
func (msg MsgSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgSwap) Route() string { return RouterKey }

func (msg MsgSwap) Type() string { return TypeMsgSwap }

//type MsgMakeOutcomePayment struct {
//	SenderDid did.Did `json:"sender_did" yaml:"sender_did"`
//	Amount    sdk.Int `json:"amount" yaml:"amount"`
//	BondDid   did.Did `json:"bond_did" yaml:"bond_did"`
//}

func NewMsgMakeOutcomePayment(senderDid did.Did, amount sdk.Int, bondDid did.Did) *MsgMakeOutcomePayment {
	return &MsgMakeOutcomePayment{
		SenderDid: senderDid,
		Amount:    amount,
		BondDid:   bondDid,
	}
}

func (msg MsgMakeOutcomePayment) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.SenderDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "sender DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Outcome payment amount has to be greater than 0
	if msg.Amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "amount")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "bond DID")
	} else if !did.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "sender DID")
	}

	return nil
}

func (msg MsgMakeOutcomePayment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgMakeOutcomePayment) GetSignerDid() did.Did { return msg.SenderDid }
func (msg MsgMakeOutcomePayment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgMakeOutcomePayment) Route() string { return RouterKey }

func (msg MsgMakeOutcomePayment) Type() string { return TypeMsgMakeOutcomePayment }

//type MsgWithdrawShare struct {
//	RecipientDid did.Did `json:"recipient_did" yaml:"recipient_did"`
//	BondDid      did.Did `json:"bond_did" yaml:"bond_did"`
//}

func NewMsgWithdrawShare(recipientDid, bondDid did.Did) *MsgWithdrawShare {
	return &MsgWithdrawShare{
		RecipientDid: recipientDid,
		BondDid:      bondDid,
	}
}

func (msg MsgWithdrawShare) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.RecipientDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "recipient DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Check that DIDs valid
	if !did.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "bond DID")
	} else if !did.IsValidDid(msg.RecipientDid) {
		return sdkerrors.Wrap(did.ErrInvalidDid, "recipient DID")
	}

	return nil
}

func (msg MsgWithdrawShare) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdrawShare) GetSignerDid() did.Did { return msg.RecipientDid }
func (msg MsgWithdrawShare) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgWithdrawShare) Route() string { return RouterKey }

func (msg MsgWithdrawShare) Type() string { return TypeMsgWithdrawShare }
