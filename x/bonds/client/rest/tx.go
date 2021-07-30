package rest

import (
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	bondsclient "github.com/ixofoundation/ixo-blockchain/x/bonds/client"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/bonds/create_bond", createBondRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bonds/edit_bond", editBondRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bonds/set_next_alpha", setNextAlphaRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bonds/update_bond_state", updateBondStateRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bonds/buy", buyRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bonds/sell", sellRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bonds/swap", swapRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bonds/make_outcome_payment", makeOutcomePaymentRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bonds/withdraw_share", withdrawShareRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bonds/withdraw_reserve", withdrawReserveRequestHandler(clientCtx)).Methods("POST")
}

type createBondReq struct {
	BaseReq                  rest.BaseReq `json:"base_req" yaml:"base_req"`
	Token                    string       `json:"token" yaml:"token"`
	Name                     string       `json:"name" yaml:"name"`
	Description              string       `json:"description" yaml:"description"`
	FunctionType             string       `json:"function_type" yaml:"function_type"`
	FunctionParameters       string       `json:"function_parameters" yaml:"function_parameters"`
	ReserveTokens            string       `json:"reserve_tokens" yaml:"reserve_tokens"`
	TxFeePercentage          string       `json:"tx_fee_percentage" yaml:"tx_fee_percentage"`
	ExitFeePercentage        string       `json:"exit_fee_percentage" yaml:"exit_fee_percentage"`
	FeeAddress               string       `json:"fee_address" yaml:"fee_address"`
	ReserveWithdrawalAddress string       `json:"reserve_withdrawal_address" yaml:"reserve_withdrawal_address"`
	MaxSupply                string       `json:"max_supply" yaml:"max_supply"`
	OrderQuantityLimits      string       `json:"order_quantity_limits" yaml:"order_quantity_limits"`
	SanityRate               string       `json:"sanity_rate" yaml:"sanity_rate"`
	SanityMarginPercentage   string       `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
	AllowSells               string       `json:"allow_sells" yaml:"allow_sells"`
	AllowReserveWithdrawals  string       `json:"allow_reserve_withdrawals" yaml:"allow_reserve_withdrawals"`
	AlphaBond                string       `json:"alpha_bond" yaml:"alpha_bond"`
	BatchBlocks              string       `json:"batch_blocks" yaml:"batch_blocks"`
	OutcomePayment           string       `json:"outcome_payment" yaml:"outcome_payment"`
	BondDid                  string       `json:"bond_did" yaml:"bond_did"`
	CreatorDid               string       `json:"creator_did" yaml:"creator_did"`
	ControllerDid            string       `json:"controller_did" yaml:"controller_did"`
}

func createBondRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createBondReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// Parse function parameters
		functionParams, err := bondsclient.ParseFunctionParams(req.FunctionParameters)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		// Parse reserve tokens
		reserveTokens := strings.Split(req.ReserveTokens, ",")

		// Parse tx fee percentage
		txFeePercentageDec, err := sdk.NewDecFromStr(req.TxFeePercentage)
		if err != nil {
			err = sdkerrors.Wrap(types.ErrArgumentMissingOrNonFloat, "tx fee percentage")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse exit fee percentage
		exitFeePercentageDec, err := sdk.NewDecFromStr(req.ExitFeePercentage)
		if err != nil {
			err = sdkerrors.Wrap(types.ErrArgumentMissingOrNonFloat, "exit fee percentage")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse fee address
		feeAddress, err := sdk.AccAddressFromBech32(req.FeeAddress)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		// Parse reserve withdrawal address
		reserveWithdrawalAddress, err := sdk.AccAddressFromBech32(req.ReserveWithdrawalAddress)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		// Parse max supply
		maxSupply, err := sdk.ParseCoinNormalized(req.MaxSupply)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		// Parse order quantity limits
		orderQuantityLimits, err := sdk.ParseCoinsNormalized(req.OrderQuantityLimits)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		// Parse sanity rate
		sanityRate, err := sdk.NewDecFromStr(req.SanityRate)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		// Parse sanity margin percentage
		sanityMarginPercentage, err := sdk.NewDecFromStr(req.SanityMarginPercentage)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		// Parse allowSells
		var allowSells bool
		allowSellsStrLower := strings.ToLower(req.AllowSells)
		if allowSellsStrLower == "true" {
			allowSells = true
		} else if allowSellsStrLower == "false" {
			allowSells = false
		} else {
			err := sdkerrors.Wrap(types.ErrArgumentMissingOrNonBoolean, "allow_sells")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse allowReserveWithdrawals
		var allowReserveWithdrawals bool
		allowReserveWithdrawalsStrLower := strings.ToLower(req.AllowReserveWithdrawals)
		if allowReserveWithdrawalsStrLower == "true" {
			allowReserveWithdrawals = true
		} else if allowReserveWithdrawalsStrLower == "false" {
			allowReserveWithdrawals = false
		} else {
			err := sdkerrors.Wrap(types.ErrArgumentMissingOrNonBoolean, "allow_reserve_withdrawals")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse alphaBond
		var alphaBond bool
		alphaBondStrLower := strings.ToLower(req.AlphaBond)
		if alphaBondStrLower == "true" {
			alphaBond = true
		} else if alphaBondStrLower == "false" {
			alphaBond = false
		} else {
			err := sdkerrors.Wrap(types.ErrArgumentMissingOrNonBoolean, "alpha_bond")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse batch blocks
		batchBlocks, err := sdk.ParseUint(req.BatchBlocks)
		if err != nil {
			err := sdkerrors.Wrap(types.ErrArgumentMissingOrNonUInteger, "max batch blocks")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse outcome payment
		var outcomePayment sdk.Int
		if len(req.OutcomePayment) == 0 {
			outcomePayment = sdk.ZeroInt()
		} else {
			var ok bool
			outcomePayment, ok = sdk.NewIntFromString(req.OutcomePayment)
			if !ok {
				err := sdkerrors.Wrap(types.ErrArgumentMustBeInteger, "outcome payment")
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		msg := types.NewMsgCreateBond(req.Token, req.Name, req.Description,
			req.CreatorDid, req.ControllerDid, req.FunctionType, functionParams,
			reserveTokens, txFeePercentageDec, exitFeePercentageDec, feeAddress,
			reserveWithdrawalAddress, maxSupply, orderQuantityLimits, sanityRate,
			sanityMarginPercentage, allowSells, allowReserveWithdrawals, alphaBond,
			batchBlocks, outcomePayment, req.BondDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type editBondReq struct {
	BaseReq                rest.BaseReq `json:"base_req" yaml:"base_req"`
	Name                   string       `json:"name" yaml:"name"`
	Description            string       `json:"description" yaml:"description"`
	OrderQuantityLimits    string       `json:"order_quantity_limits" yaml:"order_quantity_limits"`
	SanityRate             string       `json:"sanity_rate" yaml:"sanity_rate"`
	SanityMarginPercentage string       `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
	BondDid                string       `json:"bond_did" yaml:"bond_did"`
	EditorDid              string       `json:"editor_did" yaml:"editor_did"`
}

func editBondRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req editBondReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgEditBond(req.Name, req.Description,
			req.OrderQuantityLimits, req.SanityRate,
			req.SanityMarginPercentage, req.EditorDid, req.BondDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type setNextAlphaReq struct {
	BaseReq   rest.BaseReq `json:"base_req" yaml:"base_req"`
	NewAlpha  string       `json:"new_alpha" yaml:"new_alpha"`
	BondDid   string       `json:"bond_did" yaml:"bond_did"`
	EditorDid string       `json:"editor_did" yaml:"editor_did"`
}

func setNextAlphaRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setNextAlphaReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// Parse new alpha
		newAlpha, err := sdk.NewDecFromStr(req.NewAlpha)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		msg := types.NewMsgSetNextAlpha(newAlpha, req.EditorDid, req.BondDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateBondStateReq struct {
	BaseReq   rest.BaseReq `json:"base_req" yaml:"base_req"`
	NewState  string       `json:"new_state" yaml:"new_state"`
	BondDid   string       `json:"bond_did" yaml:"bond_did"`
	EditorDid string       `json:"editor_did" yaml:"editor_did"`
}

func updateBondStateRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateBondStateReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		msg := types.NewMsgUpdateBondState(types.BondState(req.NewState), req.EditorDid, req.BondDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type buyReq struct {
	BaseReq    rest.BaseReq `json:"base_req" yaml:"base_req"`
	BondToken  string       `json:"bond_token" yaml:"bond_token"`
	BondAmount string       `json:"bond_amount" yaml:"bond_amount"`
	MaxPrices  string       `json:"max_prices" yaml:"max_prices"`
	BondDid    string       `json:"bond_did" yaml:"bond_did"`
	BuyerDid   string       `json:"buyer_did" yaml:"buyer_did"`
}

func buyRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req buyReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		bondCoin, err := bondsclient.ParseTwoPartCoin(req.BondAmount, req.BondToken)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		maxPrices, err := sdk.ParseCoinsNormalized(req.MaxPrices)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		msg := types.NewMsgBuy(req.BuyerDid, bondCoin, maxPrices, req.BondDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type sellReq struct {
	BaseReq    rest.BaseReq `json:"base_req" yaml:"base_req"`
	BondToken  string       `json:"bond_token" yaml:"bond_token"`
	BondAmount string       `json:"bond_amount" yaml:"bond_amount"`
	BondDid    string       `json:"bond_did" yaml:"bond_did"`
	SellerDid  string       `json:"seller_did" yaml:"seller_did"`
}

func sellRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req sellReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		bondCoin, err := bondsclient.ParseTwoPartCoin(req.BondAmount, req.BondToken)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		msg := types.NewMsgSell(req.SellerDid, bondCoin, req.BondDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type swapReq struct {
	BaseReq    rest.BaseReq `json:"base_req" yaml:"base_req"`
	FromAmount string       `json:"from_amount" yaml:"from_amount"`
	FromToken  string       `json:"from_token" yaml:"from_token"`
	ToToken    string       `json:"to_token" yaml:"to_token"`
	BondDid    string       `json:"bond_did" yaml:"bond_did"`
	SwapperDid string       `json:"swapper_did" yaml:"swapper_did"`
}

func swapRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req swapReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// Check that from amount and token can be parsed to a coin
		fromCoin, err := bondsclient.ParseTwoPartCoin(req.FromAmount, req.FromToken)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		msg := types.NewMsgSwap(req.SwapperDid, fromCoin, req.ToToken, req.BondDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type makeOutcomePaymentReq struct {
	BaseReq   rest.BaseReq `json:"base_req" yaml:"base_req"`
	BondDid   string       `json:"bond_did" yaml:"bond_did"`
	Amount    string       `json:"amount" yaml:"amount"`
	SenderDid string       `json:"sender_did" yaml:"sender_did"`
}

func makeOutcomePaymentRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req makeOutcomePaymentReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		amount, ok := sdk.NewIntFromString(req.Amount)
		if !ok {
			err := sdkerrors.Wrap(types.ErrArgumentMustBeInteger, "amount")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgMakeOutcomePayment(req.SenderDid, amount, req.BondDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type withdrawShareReq struct {
	BaseReq      rest.BaseReq `json:"base_req" yaml:"base_req"`
	BondDid      string       `json:"bond_did" yaml:"bond_did"`
	RecipientDid string       `json:"recipient_did" yaml:"recipient_did"`
}

func withdrawShareRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req withdrawShareReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgWithdrawShare(req.RecipientDid, req.BondDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type withdrawReserveReq struct {
	BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
	BondDid       string       `json:"bond_did" yaml:"bond_did"`
	Amount        string       `json:"amount" yaml:"amount"`
	WithdrawerDid string       `json:"withdrawer_did" yaml:"withdrawer_did"`
}

func withdrawReserveRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req withdrawReserveReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		amount, err := sdk.ParseCoinsNormalized(req.Amount)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		msg := types.NewMsgWithdrawReserve(req.WithdrawerDid, amount, req.BondDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
