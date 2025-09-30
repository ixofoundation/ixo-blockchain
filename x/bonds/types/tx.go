package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	iidante "github.com/ixofoundation/ixo-blockchain/v6/x/iid/ante"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
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

// --------------------------
// CREATE BOND
// --------------------------
func NewMsgCreateBond(token, name, description string, creatorDid, controllerDid, oracleDid iidtypes.DIDFragment,
	functionType string, functionParameters FunctionParams, reserveTokens []string,
	txFeePercentage, exitFeePercentage math.LegacyDec, feeAddress, reserveWithdrawalAddress sdk.AccAddress,
	maxSupply sdk.Coin, orderQuantityLimits sdk.Coins, sanityRate, sanityMarginPercentage math.LegacyDec,
	allowSell, allowReserveWithdrawals, alphaBond bool, batchBlocks math.Uint, outcomePayment math.Int,
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

func (msg MsgCreateBond) Route() string { return RouterKey }

func (msg MsgCreateBond) Type() string { return TypeMsgCreateBond }

// --------------------------
// EDIT BOND
// --------------------------
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

func (msg MsgEditBond) Route() string { return RouterKey }

func (msg MsgEditBond) Type() string { return TypeMsgEditBond }

// --------------------------
// SET NEXT ALPHA
// --------------------------
func NewMsgSetNextAlpha(alpha math.LegacyDec, oracleDid iidtypes.DIDFragment, bondDid string, oracleAddress string) *MsgSetNextAlpha {
	return &MsgSetNextAlpha{
		BondDid:       bondDid,
		Alpha:         alpha,
		OracleDid:     oracleDid,
		OracleAddress: oracleAddress,
	}
}

func (msg MsgSetNextAlpha) GetIidController() iidtypes.DIDFragment { return msg.OracleDid }

func (msg MsgSetNextAlpha) Route() string { return RouterKey }

func (msg MsgSetNextAlpha) Type() string { return TypeMsgSetNextAlpha }

// --------------------------
// UPDATE BOND STATE
// --------------------------
func NewMsgUpdateBondState(state BondState, editorDid iidtypes.DIDFragment, bondDid string, editorAddress string) *MsgUpdateBondState {
	return &MsgUpdateBondState{
		BondDid:       bondDid,
		State:         state.String(),
		EditorDid:     editorDid,
		EditorAddress: editorAddress,
	}
}

func (msg MsgUpdateBondState) GetIidController() iidtypes.DIDFragment { return msg.EditorDid }

func (msg MsgUpdateBondState) Route() string { return RouterKey }

func (msg MsgUpdateBondState) Type() string { return TypeMsgUpdateBondState }

// --------------------------
// BUY
// --------------------------
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

func (msg MsgBuy) Route() string { return RouterKey }

func (msg MsgBuy) Type() string { return TypeMsgBuy }

// --------------------------
// SELL
// --------------------------
func NewMsgSell(sellerDid iidtypes.DIDFragment, amount sdk.Coin, bondDid string, sellerAddress string) *MsgSell {
	return &MsgSell{
		SellerDid:     sellerDid,
		Amount:        amount,
		BondDid:       bondDid,
		SellerAddress: sellerAddress,
	}
}

func (msg MsgSell) GetIidController() iidtypes.DIDFragment { return msg.SellerDid }

func (msg MsgSell) Route() string { return RouterKey }

func (msg MsgSell) Type() string { return TypeMsgSell }

// --------------------------
// SWAP
// --------------------------
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

func (msg MsgSwap) Route() string { return RouterKey }

func (msg MsgSwap) Type() string { return TypeMsgSwap }

// --------------------------
// MAKE OUTCOME PAYMENT
// --------------------------
func NewMsgMakeOutcomePayment(senderDid iidtypes.DIDFragment, amount math.Int, bondDid string, senderAddress string) *MsgMakeOutcomePayment {
	return &MsgMakeOutcomePayment{
		SenderDid:     senderDid,
		Amount:        amount,
		BondDid:       bondDid,
		SenderAddress: senderAddress,
	}
}

func (msg MsgMakeOutcomePayment) GetIidController() iidtypes.DIDFragment { return msg.SenderDid }

func (msg MsgMakeOutcomePayment) Route() string { return RouterKey }

func (msg MsgMakeOutcomePayment) Type() string { return TypeMsgMakeOutcomePayment }

// --------------------------
// WITHDRAW SHARE
// --------------------------
func NewMsgWithdrawShare(recipientDid iidtypes.DIDFragment, bondDid string, recipientAddress string) *MsgWithdrawShare {
	return &MsgWithdrawShare{
		RecipientDid:     recipientDid,
		BondDid:          bondDid,
		RecipientAddress: recipientAddress,
	}
}

func (msg MsgWithdrawShare) GetIidController() iidtypes.DIDFragment { return msg.RecipientDid }

func (msg MsgWithdrawShare) Route() string { return RouterKey }

func (msg MsgWithdrawShare) Type() string { return TypeMsgWithdrawShare }

// --------------------------
// WITHDRAW RESERVE
// --------------------------
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

func (msg MsgWithdrawReserve) Route() string { return RouterKey }

func (msg MsgWithdrawReserve) Type() string { return TypeMsgWithdrawReserve }
