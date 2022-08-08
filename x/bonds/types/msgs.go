package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	didexported "github.com/ixofoundation/ixo-blockchain/x/did/exported"
	didtypes "github.com/ixofoundation/ixo-blockchain/x/did/types"
	ixotypes "github.com/ixofoundation/ixo-blockchain/x/ixo/types"
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
	TypeMsgWithdrawReserve    = "withdraw_reserve"
)

var (
	_ ixotypes.IxoMsg = &MsgCreateBond{}
	_ ixotypes.IxoMsg = &MsgEditBond{}
	_ ixotypes.IxoMsg = &MsgSetNextAlpha{}
	_ ixotypes.IxoMsg = &MsgUpdateBondState{}
	_ ixotypes.IxoMsg = &MsgBuy{}
	_ ixotypes.IxoMsg = &MsgSell{}
	_ ixotypes.IxoMsg = &MsgSwap{}
	_ ixotypes.IxoMsg = &MsgMakeOutcomePayment{}
	_ ixotypes.IxoMsg = &MsgWithdrawShare{}
	_ ixotypes.IxoMsg = &MsgWithdrawReserve{}
)

func NewMsgCreateBond(token, name, description string, creatorDid, controllerDid didexported.Did,
	functionType string, functionParameters FunctionParams, reserveTokens []string,
	txFeePercentage, exitFeePercentage sdk.Dec, feeAddress, reserveWithdrawalAddress sdk.AccAddress,
	maxSupply sdk.Coin, orderQuantityLimits sdk.Coins, sanityRate, sanityMarginPercentage sdk.Dec,
	allowSell, allowReserveWithdrawals, alphaBond bool, batchBlocks sdk.Uint, outcomePayment sdk.Int,
	bondDid didexported.Did) *MsgCreateBond {

	return &MsgCreateBond{
		BondDid:                  bondDid,
		Token:                    token,
		Name:                     name,
		Description:              description,
		CreatorDid:               creatorDid,
		ControllerDid:            controllerDid,
		FunctionType:             functionType,
		FunctionParameters:       functionParameters,
		ReserveTokens:            reserveTokens,
		TxFeePercentage:          txFeePercentage,
		ExitFeePercentage:        exitFeePercentage,
		FeeAddress:               feeAddress.String(),
		ReserveWithdrawalAddress: reserveWithdrawalAddress.String(),
		MaxSupply:                maxSupply,
		OrderQuantityLimits:      orderQuantityLimits,
		SanityRate:               sanityRate,
		SanityMarginPercentage:   sanityMarginPercentage,
		AllowSells:               allowSell,
		AllowReserveWithdrawals:  allowReserveWithdrawals,
		AlphaBond:                alphaBond,
		BatchBlocks:              batchBlocks,
		OutcomePayment:           outcomePayment,
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
	} else if strings.TrimSpace(msg.ReserveWithdrawalAddress) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "reserve withdrawal address")
	} else if strings.TrimSpace(msg.FunctionType) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "function type")
	}
	// Note: FunctionParameters can be empty

	// Checks that the bond token starts with 'x'
	if !strings.HasPrefix(msg.Token, "x") {
		return ErrInvalidBondTokenPrefix
	}

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
	if msg.AlphaBond && (msg.FunctionType != AugmentedFunction || msg.FunctionType != BondingFunction) {
		return sdkerrors.Wrap(ErrFunctionNotAvailableForFunctionType,
			"only augmented bonding curves can be alpha bonds")
	}

	// Check that outcome payment not negative
	if msg.OutcomePayment.IsNegative() {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "outcome payment")
	}

	// Note: uniqueness of reserve tokens checked when parsing

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "bond DID")
	} else if !didtypes.IsValidDid(msg.CreatorDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "creator DID")
	}

	// Check that allowSells and allowReserveWithdrawals are not both True
	if msg.AllowSells && msg.AllowReserveWithdrawals {
		return ErrCannotAllowSellsAndWithdrawals
	}

	return nil
}

func (msg MsgCreateBond) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateBond) GetSignerDid() didexported.Did { return msg.CreatorDid }
func (msg MsgCreateBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreateBond) Route() string { return RouterKey }

func (msg MsgCreateBond) Type() string { return TypeMsgCreateBond }

func NewMsgEditBond(name, description, orderQuantityLimits, sanityRate,
	sanityMarginPercentage string, editorDid, bondDid didexported.Did) *MsgEditBond {
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
	if !didtypes.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "bond DID")
	} else if !didtypes.IsValidDid(msg.EditorDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "editor DID")
	}

	return nil
}

func (msg MsgEditBond) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgEditBond) GetSignerDid() didexported.Did { return msg.EditorDid }
func (msg MsgEditBond) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgEditBond) Route() string { return RouterKey }

func (msg MsgEditBond) Type() string { return TypeMsgEditBond }

func NewMsgSetNextAlpha(alpha sdk.Dec, editorDid, bondDid didexported.Did) *MsgSetNextAlpha {
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
	if !didtypes.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "bond did is invalid")
	} else if !didtypes.IsValidDid(msg.EditorDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "editor did is invalid")
	}

	return nil
}

func (msg MsgSetNextAlpha) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgSetNextAlpha) GetSignerDid() didexported.Did { return msg.EditorDid }
func (msg MsgSetNextAlpha) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgSetNextAlpha) Route() string { return RouterKey }

func (msg MsgSetNextAlpha) Type() string { return TypeMsgSetNextAlpha }

func NewMsgUpdateBondState(state BondState, editorDid, bondDid didexported.Did) *MsgUpdateBondState {
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
	if !didtypes.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "bond did is invalid")
	} else if !didtypes.IsValidDid(msg.EditorDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "editor did is invalid")
	}

	return nil
}

func (msg MsgUpdateBondState) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUpdateBondState) GetSignerDid() didexported.Did { return msg.EditorDid }
func (msg MsgUpdateBondState) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgUpdateBondState) Route() string { return RouterKey }

func (msg MsgUpdateBondState) Type() string { return TypeMsgUpdateBondState }

func NewMsgBuy(buyerDid didexported.Did, amount sdk.Coin, maxPrices sdk.Coins,
	bondDid didexported.Did) *MsgBuy {
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
	if !didtypes.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "bond DID")
	} else if !didtypes.IsValidDid(msg.BuyerDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "buyer DID")
	}

	return nil
}

func (msg MsgBuy) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgBuy) GetSignerDid() didexported.Did { return msg.BuyerDid }
func (msg MsgBuy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgBuy) Route() string { return RouterKey }

func (msg MsgBuy) Type() string { return TypeMsgBuy }

func NewMsgSell(sellerDid didexported.Did, amount sdk.Coin, bondDid didexported.Did) *MsgSell {
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
	if !didtypes.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "bond DID")
	} else if !didtypes.IsValidDid(msg.SellerDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "seller DID")
	}

	return nil
}

func (msg MsgSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgSell) GetSignerDid() didexported.Did { return msg.SellerDid }
func (msg MsgSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgSell) Route() string { return RouterKey }

func (msg MsgSell) Type() string { return TypeMsgSell }

func NewMsgSwap(swapperDid didexported.Did, from sdk.Coin, toToken string,
	bondDid didexported.Did) *MsgSwap {
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
	if !didtypes.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "bond DID is invalid")
	} else if !didtypes.IsValidDid(msg.SwapperDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "swapper DID is invalid")
	}

	return nil
}

func (msg MsgSwap) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgSwap) GetSignerDid() didexported.Did { return msg.SwapperDid }
func (msg MsgSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgSwap) Route() string { return RouterKey }

func (msg MsgSwap) Type() string { return TypeMsgSwap }

func NewMsgMakeOutcomePayment(senderDid didexported.Did, amount sdk.Int, bondDid didexported.Did) *MsgMakeOutcomePayment {
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
	if !didtypes.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "bond DID")
	} else if !didtypes.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender DID")
	}

	return nil
}

func (msg MsgMakeOutcomePayment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgMakeOutcomePayment) GetSignerDid() didexported.Did { return msg.SenderDid }
func (msg MsgMakeOutcomePayment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgMakeOutcomePayment) Route() string { return RouterKey }

func (msg MsgMakeOutcomePayment) Type() string { return TypeMsgMakeOutcomePayment }

func NewMsgWithdrawShare(recipientDid, bondDid didexported.Did) *MsgWithdrawShare {
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
	if !didtypes.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "bond DID")
	} else if !didtypes.IsValidDid(msg.RecipientDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "recipient DID")
	}

	return nil
}

func (msg MsgWithdrawShare) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdrawShare) GetSignerDid() didexported.Did { return msg.RecipientDid }
func (msg MsgWithdrawShare) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgWithdrawShare) Route() string { return RouterKey }

func (msg MsgWithdrawShare) Type() string { return TypeMsgWithdrawShare }

func NewMsgWithdrawReserve(withdrawerDid didexported.Did, amount sdk.Coins,
	bondDid didexported.Did) *MsgWithdrawReserve {
	return &MsgWithdrawReserve{
		WithdrawerDid: withdrawerDid,
		Amount:        amount,
		BondDid:       bondDid,
	}
}

func (msg MsgWithdrawReserve) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.WithdrawerDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "withdrawer DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Validate amount
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.BondDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "bond DID")
	} else if !didtypes.IsValidDid(msg.WithdrawerDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "withdrawer DID")
	}

	return nil
}

func (msg MsgWithdrawReserve) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdrawReserve) GetSignerDid() didexported.Did { return msg.WithdrawerDid }
func (msg MsgWithdrawReserve) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgWithdrawReserve) Route() string { return RouterKey }

func (msg MsgWithdrawReserve) Type() string { return TypeMsgWithdrawReserve }
