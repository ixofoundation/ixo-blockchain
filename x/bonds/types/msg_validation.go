package types

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
)

// --------------------------
// CREATE BOND
// --------------------------
func (msg MsgCreateBond) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	} else if strings.TrimSpace(msg.Token) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "token")
	} else if strings.TrimSpace(msg.Name) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "name")
	} else if strings.TrimSpace(msg.Description) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "description")
	} else if strings.TrimSpace(msg.CreatorDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "creator DID")
	} else if strings.TrimSpace(msg.ControllerDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "controller DID")
	} else if strings.TrimSpace(msg.OracleDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "oracle DID")
	} else if len(msg.ReserveTokens) == 0 {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "reserve tokens")
	} else if strings.TrimSpace(msg.FeeAddress) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "fee address")
	} else if strings.TrimSpace(msg.ReserveWithdrawalAddress) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "reserve withdrawal address")
	} else if strings.TrimSpace(msg.FunctionType) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "function type")
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
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "max supply")
	} else if !msg.OrderQuantityLimits.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "order quantity limits")
	}

	// Check that max supply denom matches token denom
	if msg.MaxSupply.Denom != msg.Token {
		return ErrMaxSupplyDenomDoesNotMatchTokenDenom
	}

	// Check that Sanity values not negative
	if msg.SanityRate.IsNegative() {
		return errorsmod.Wrap(ErrArgumentCannotBeNegative, "sanity rate")
	} else if msg.SanityMarginPercentage.IsNegative() {
		return errorsmod.Wrap(ErrArgumentCannotBeNegative, "sanity margin percentage")
	}

	// Check FeePercentages not negative and don't add up to 100
	if msg.TxFeePercentage.IsNegative() {
		return errorsmod.Wrap(ErrArgumentCannotBeNegative, "tx fee percentage")
	} else if msg.ExitFeePercentage.IsNegative() {
		return errorsmod.Wrap(ErrArgumentCannotBeNegative, "exit fee percentage")
	} else if msg.TxFeePercentage.Add(msg.ExitFeePercentage).GTE(math.LegacyNewDec(100)) {
		return ErrFeesCannotBeOrExceed100Percent
	}

	// Check that not zero
	if msg.BatchBlocks.IsZero() {
		return errorsmod.Wrap(ErrArgumentMustBePositive, "batch blocks")
	} else if msg.MaxSupply.Amount.IsZero() {
		return errorsmod.Wrap(ErrArgumentMustBePositive, "max supply")
	}

	// Alpha bonds have to be augmented bonding curves
	if msg.AlphaBond && msg.FunctionType != AugmentedFunction {
		return errorsmod.Wrap(ErrFunctionNotAvailableForFunctionType,
			"only augmented bonding curves can be alpha bonds")
	}

	// Check that outcome payment not negative
	if msg.OutcomePayment.IsNegative() {
		return errorsmod.Wrap(ErrArgumentMustBePositive, "outcome payment")
	}

	// Note: uniqueness of reserve tokens checked when parsing

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.CreatorDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.CreatorDid.Did())
	}

	// Check that allowSells and allowReserveWithdrawals are not both True
	if msg.AllowSells && msg.AllowReserveWithdrawals {
		return ErrCannotAllowSellsAndWithdrawals
	}

	return nil
}

// --------------------------
// EDIT BOND
// --------------------------
func (msg MsgEditBond) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	} else if strings.TrimSpace(msg.Name) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "name")
	} else if strings.TrimSpace(msg.Description) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "description")
	} else if strings.TrimSpace(msg.SanityRate) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "sanity rate")
	} else if strings.TrimSpace(msg.SanityMarginPercentage) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "sanity margin percentage")
	} else if strings.TrimSpace(msg.EditorDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "editor DID")
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
	if !iidtypes.IsValidDID(msg.BondDid) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.EditorDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.EditorDid.Did())
	}

	return nil
}

// --------------------------
// SET NEXT ALPHA
// --------------------------
func (msg MsgSetNextAlpha) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "BondDid")
	} else if strings.TrimSpace(msg.OracleDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "EditorDid")
	}

	// Check that 0.0001 <= alpha <= 0.9999. Note that we cannot set public
	// alpha to 0 or 1, because these are edge cases which cause the system
	// alpha to get stuck if we try to change the value of public alpha again.
	minNextAlpha := math.LegacyMustNewDecFromStr("0.0001")
	maxNextAlpha := math.LegacyMustNewDecFromStr("0.9999")
	if msg.Alpha.LT(minNextAlpha) || msg.Alpha.GT(maxNextAlpha) {
		return errorsmod.Wrap(ErrInvalidAlpha, "0.0001 <= alpha <= 0.9999")
	}

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.OracleDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.OracleDid.Did())
	}

	return nil
}

// --------------------------
// UPDATE BOND STATE
// --------------------------
func (msg MsgUpdateBondState) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "BondDid")
	} else if strings.TrimSpace(msg.EditorDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "EditorDid")
	}

	// Bond status can only be updated to SETTLE or FAILED
	if msg.State != SettleState.String() && msg.State != FailedState.String() {
		return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "cannot transition to that state")
	}

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.EditorDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.EditorDid.Did())
	}

	return nil
}

// --------------------------
// BUY
// --------------------------
func (msg MsgBuy) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BuyerDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "buyer DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Check that amount valid and non zero
	if !msg.Amount.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	} else if msg.Amount.Amount.IsZero() {
		return errorsmod.Wrap(ErrArgumentMustBePositive, "amount")
	}

	// Check that maxPrices valid
	if !msg.MaxPrices.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "max prices")
	}

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.BuyerDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BuyerDid.Did())
	}

	return nil
}

// --------------------------
// SELL
// --------------------------
func (msg MsgSell) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.SellerDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "seller DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Check that amount valid and non zero
	if !msg.Amount.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	} else if msg.Amount.Amount.IsZero() {
		return errorsmod.Wrap(ErrArgumentMustBePositive, "amount")
	}

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.SellerDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.SellerDid.Did())
	}

	return nil
}

// --------------------------
// SWAP
// --------------------------
func (msg MsgSwap) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.SwapperDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "swapper DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	} else if strings.TrimSpace(msg.ToToken) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "to token")
	}

	// Validate from amount
	if !msg.From.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.From.String())
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
		return errorsmod.Wrap(ErrArgumentMustBePositive, "amount")
	}

	// Note: From denom and amount must be valid since sdk.Coin

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.SwapperDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.SwapperDid.Did())
	}

	return nil
}

// --------------------------
// MAKE OUTCOME PAYMENT
// --------------------------
func (msg MsgMakeOutcomePayment) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.SenderDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "sender DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Outcome payment amount has to be greater than 0
	if msg.Amount.LTE(math.ZeroInt()) {
		return errorsmod.Wrap(ErrArgumentMustBePositive, "amount")
	}

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.SenderDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.SenderDid.Did())
	}

	return nil
}

// --------------------------
// WITHDRAW SHARE
// --------------------------
func (msg MsgWithdrawShare) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.RecipientDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "recipient DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.RecipientDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.RecipientDid.Did())
	}

	return nil
}

// --------------------------
// WITHDRAW RESERVE
// --------------------------
func (msg MsgWithdrawReserve) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.WithdrawerDid.String()) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "withdrawer DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return errorsmod.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Validate amount
	if !msg.Amount.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.WithdrawerDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.WithdrawerDid.Did())
	}

	return nil
}
