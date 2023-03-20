package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidante "github.com/ixofoundation/ixo-blockchain/x/iid/ante"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
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
	_ iidante.IidTxMsg = &MsgCreateBond{}
	_ iidante.IidTxMsg = &MsgEditBond{}
	_ iidante.IidTxMsg = &MsgSetNextAlpha{}
	_ iidante.IidTxMsg = &MsgUpdateBondState{}
	_ iidante.IidTxMsg = &MsgBuy{}
	_ iidante.IidTxMsg = &MsgSell{}
	_ iidante.IidTxMsg = &MsgSwap{}
	_ iidante.IidTxMsg = &MsgMakeOutcomePayment{}
	_ iidante.IidTxMsg = &MsgWithdrawShare{}
	_ iidante.IidTxMsg = &MsgWithdrawReserve{}
)

func NewMsgCreateBond(token, name, description string, creatorDid, controllerDid, oracleDid iidtypes.DIDFragment,
	functionType string, functionParameters FunctionParams, reserveTokens []string,
	txFeePercentage, exitFeePercentage sdk.Dec, feeAddress, reserveWithdrawalAddress sdk.AccAddress,
	maxSupply sdk.Coin, orderQuantityLimits sdk.Coins, sanityRate, sanityMarginPercentage sdk.Dec,
	allowSell, allowReserveWithdrawals, alphaBond bool, batchBlocks sdk.Uint, outcomePayment sdk.Int,
	bondDid string, creatorAddress string) *MsgCreateBond {

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
		CreatorAddress:           creatorAddress,
	}
}

func (msg MsgCreateBond) GetIidController() iidtypes.DIDFragment { return msg.CreatorDid }

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
	} else if strings.TrimSpace(msg.CreatorDid.String()) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "creator DID")
	} else if strings.TrimSpace(msg.ControllerDid.String()) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "controller DID")
	} else if strings.TrimSpace(msg.OracleDid.String()) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "oracle DID")
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
	if !iidtypes.IsValidDID(msg.BondDid) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.CreatorDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.CreatorDid.Did())
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

// func (msg MsgCreateBond) GetSignerDid() string { return msg.CreatorDid.Did() }
func (msg MsgCreateBond) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.CreatorAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgCreateBond) Route() string { return RouterKey }

func (msg MsgCreateBond) Type() string { return TypeMsgCreateBond }

func NewMsgEditBond(name, description, orderQuantityLimits, sanityRate,
	sanityMarginPercentage string, editorDid iidtypes.DIDFragment, bondDid string, editorAddress string) *MsgEditBond {
	return &MsgEditBond{
		BondDid:                bondDid,
		Name:                   name,
		Description:            description,
		OrderQuantityLimits:    orderQuantityLimits,
		SanityRate:             sanityRate,
		SanityMarginPercentage: sanityMarginPercentage,
		EditorDid:              editorDid,
		EditorAddress:          editorAddress,
	}
}

func (msg MsgEditBond) GetIidController() iidtypes.DIDFragment { return msg.EditorDid }

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
	} else if strings.TrimSpace(msg.EditorDid.String()) == "" {
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
	if !iidtypes.IsValidDID(msg.BondDid) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.EditorDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.EditorDid.Did())
	}

	return nil
}

func (msg MsgEditBond) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgEditBond) GetSignerDid() string { return msg.EditorDid }
func (msg MsgEditBond) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.EditorAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgEditBond) Route() string { return RouterKey }

func (msg MsgEditBond) Type() string { return TypeMsgEditBond }

func NewMsgSetNextAlpha(alpha sdk.Dec, oracleDid iidtypes.DIDFragment, bondDid string, oracleAddress string) *MsgSetNextAlpha {
	return &MsgSetNextAlpha{
		BondDid:       bondDid,
		Alpha:         alpha,
		OracleDid:     oracleDid,
		OracleAddress: oracleAddress,
	}
}

func (msg MsgSetNextAlpha) GetIidController() iidtypes.DIDFragment { return msg.OracleDid }

func (msg MsgSetNextAlpha) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "BondDid")
	} else if strings.TrimSpace(msg.OracleDid.String()) == "" {
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
	if !iidtypes.IsValidDID(msg.BondDid) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.OracleDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.OracleDid.Did())
	}

	return nil
}

func (msg MsgSetNextAlpha) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgSetNextAlpha) GetSignerDid() string { return msg.EditorDid }
func (msg MsgSetNextAlpha) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.OracleAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgSetNextAlpha) Route() string { return RouterKey }

func (msg MsgSetNextAlpha) Type() string { return TypeMsgSetNextAlpha }

func NewMsgUpdateBondState(state BondState, editorDid iidtypes.DIDFragment, bondDid string, editorAddress string) *MsgUpdateBondState {
	return &MsgUpdateBondState{
		BondDid:       bondDid,
		State:         state.String(),
		EditorDid:     editorDid,
		EditorAddress: editorAddress,
	}
}

func (msg MsgUpdateBondState) GetIidController() iidtypes.DIDFragment { return msg.EditorDid }

func (msg MsgUpdateBondState) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "BondDid")
	} else if strings.TrimSpace(msg.EditorDid.String()) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "EditorDid")
	}

	// Bond status can only be updated to SETTLE or FAILED
	if msg.State != SettleState.String() && msg.State != FailedState.String() {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "cannot transition to that state")
	}

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.EditorDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.EditorDid.Did())
	}

	return nil
}

func (msg MsgUpdateBondState) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgUpdateBondState) GetSignerDid() string { return msg.EditorDid }
func (msg MsgUpdateBondState) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.EditorAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgUpdateBondState) Route() string { return RouterKey }

func (msg MsgUpdateBondState) Type() string { return TypeMsgUpdateBondState }

func NewMsgBuy(buyerDid iidtypes.DIDFragment, amount sdk.Coin, maxPrices sdk.Coins,
	bondDid string, buyerAddress string) *MsgBuy {
	return &MsgBuy{
		BuyerDid:     buyerDid,
		Amount:       amount,
		MaxPrices:    maxPrices,
		BondDid:      bondDid,
		BuyerAddress: buyerAddress,
	}
}

func (msg MsgBuy) GetIidController() iidtypes.DIDFragment { return msg.BuyerDid }

func (msg MsgBuy) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.BuyerDid.String()) == "" {
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
	if !iidtypes.IsValidDID(msg.BondDid) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.BuyerDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BuyerDid.Did())
	}

	return nil
}

func (msg MsgBuy) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgBuy) GetSignerDid() string { return msg.BuyerDid }
func (msg MsgBuy) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.BuyerAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgBuy) Route() string { return RouterKey }

func (msg MsgBuy) Type() string { return TypeMsgBuy }

func NewMsgSell(sellerDid iidtypes.DIDFragment, amount sdk.Coin, bondDid string, sellerAddress string) *MsgSell {
	return &MsgSell{
		SellerDid:     sellerDid,
		Amount:        amount,
		BondDid:       bondDid,
		SellerAddress: sellerAddress,
	}
}

func (msg MsgSell) GetIidController() iidtypes.DIDFragment { return msg.SellerDid }

func (msg MsgSell) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.SellerDid.String()) == "" {
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
	if !iidtypes.IsValidDID(msg.BondDid) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.SellerDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.SellerDid.Did())
	}

	return nil
}

func (msg MsgSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgSell) GetSignerDid() string { return msg.SellerDid }
func (msg MsgSell) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.SellerAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgSell) Route() string { return RouterKey }

func (msg MsgSell) Type() string { return TypeMsgSell }

func NewMsgSwap(swapperDid iidtypes.DIDFragment, from sdk.Coin, toToken string,
	bondDid string, swapperAddress string) *MsgSwap {
	return &MsgSwap{
		SwapperDid:     swapperDid,
		From:           from,
		ToToken:        toToken,
		BondDid:        bondDid,
		SwapperAddress: swapperAddress,
	}
}

func (msg MsgSwap) GetIidController() iidtypes.DIDFragment { return msg.SwapperDid }

func (msg MsgSwap) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.SwapperDid.String()) == "" {
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
	if !iidtypes.IsValidDID(msg.BondDid) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.SwapperDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.SwapperDid.Did())
	}

	return nil
}

func (msg MsgSwap) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgSwap) GetSignerDid() string { return msg.SwapperDid }
func (msg MsgSwap) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.SwapperAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgSwap) Route() string { return RouterKey }

func (msg MsgSwap) Type() string { return TypeMsgSwap }

func NewMsgMakeOutcomePayment(senderDid iidtypes.DIDFragment, amount sdk.Int, bondDid string, senderAddress string) *MsgMakeOutcomePayment {
	return &MsgMakeOutcomePayment{
		SenderDid:     senderDid,
		Amount:        amount,
		BondDid:       bondDid,
		SenderAddress: senderAddress,
	}
}

func (msg MsgMakeOutcomePayment) GetIidController() iidtypes.DIDFragment { return msg.SenderDid }

func (msg MsgMakeOutcomePayment) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.SenderDid.String()) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "sender DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Outcome payment amount has to be greater than 0
	if msg.Amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "amount")
	}

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.SenderDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.SenderDid.Did())
	}

	return nil
}

func (msg MsgMakeOutcomePayment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgMakeOutcomePayment) GetSignerDid() string { return msg.SenderDid }
func (msg MsgMakeOutcomePayment) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.SenderAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgMakeOutcomePayment) Route() string { return RouterKey }

func (msg MsgMakeOutcomePayment) Type() string { return TypeMsgMakeOutcomePayment }

func NewMsgWithdrawShare(recipientDid iidtypes.DIDFragment, bondDid string, recipientAddress string) *MsgWithdrawShare {
	return &MsgWithdrawShare{
		RecipientDid:     recipientDid,
		BondDid:          bondDid,
		RecipientAddress: recipientAddress,
	}
}

func (msg MsgWithdrawShare) GetIidController() iidtypes.DIDFragment { return msg.RecipientDid }

func (msg MsgWithdrawShare) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.RecipientDid.String()) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "recipient DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.RecipientDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.RecipientDid.Did())
	}

	return nil
}

func (msg MsgWithdrawShare) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgWithdrawShare) GetSignerDid() string { return msg.RecipientDid }
func (msg MsgWithdrawShare) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.RecipientAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgWithdrawShare) Route() string { return RouterKey }

func (msg MsgWithdrawShare) Type() string { return TypeMsgWithdrawShare }

func NewMsgWithdrawReserve(withdrawerDid iidtypes.DIDFragment, amount sdk.Coins,
	bondDid string, withdrawerAddress string) *MsgWithdrawReserve {
	return &MsgWithdrawReserve{
		WithdrawerDid:     withdrawerDid,
		Amount:            amount,
		BondDid:           bondDid,
		WithdrawerAddress: withdrawerAddress,
	}
}

func (msg MsgWithdrawReserve) GetIidController() iidtypes.DIDFragment { return msg.WithdrawerDid }

func (msg MsgWithdrawReserve) ValidateBasic() error {
	// Check if empty
	if strings.TrimSpace(msg.WithdrawerDid.String()) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "withdrawer DID")
	} else if strings.TrimSpace(msg.BondDid) == "" {
		return sdkerrors.Wrap(ErrArgumentCannotBeEmpty, "bond DID")
	}

	// Validate amount
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	// Check that DIDs valid
	if !iidtypes.IsValidDID(msg.BondDid) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.BondDid)
	}
	if !iidtypes.IsValidDID(msg.WithdrawerDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.WithdrawerDid.Did())
	}

	return nil
}

func (msg MsgWithdrawReserve) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgWithdrawReserve) GetSignerDid() string { return msg.WithdrawerDid }
func (msg MsgWithdrawReserve) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.WithdrawerAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgWithdrawReserve) Route() string { return RouterKey }

func (msg MsgWithdrawReserve) Type() string { return TypeMsgWithdrawReserve }
