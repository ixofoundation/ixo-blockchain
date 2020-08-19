package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/client"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"
	"net/http"
	"strings"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/bonds/create_bond", createBondRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/bonds/edit_bond", editBondRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/bonds/buy", buyRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/bonds/sell", sellRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/bonds/swap", swapRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/bonds/make_outcome_payment", makeOutcomePaymentRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/bonds/withdraw_share", withdrawShareRequestHandler(cliCtx)).Methods("POST")
}

type createBondReq struct {
	BaseReq                rest.BaseReq `json:"base_req" yaml:"base_req"`
	Token                  string       `json:"token" yaml:"token"`
	Name                   string       `json:"name" yaml:"name"`
	Description            string       `json:"description" yaml:"description"`
	FunctionType           string       `json:"function_type" yaml:"function_type"`
	FunctionParameters     string       `json:"function_parameters" yaml:"function_parameters"`
	ReserveTokens          string       `json:"reserve_tokens" yaml:"reserve_tokens"`
	TxFeePercentage        string       `json:"tx_fee_percentage" yaml:"tx_fee_percentage"`
	ExitFeePercentage      string       `json:"exit_fee_percentage" yaml:"exit_fee_percentage"`
	FeeAddress             string       `json:"fee_address" yaml:"fee_address"`
	MaxSupply              string       `json:"max_supply" yaml:"max_supply"`
	OrderQuantityLimits    string       `json:"order_quantity_limits" yaml:"order_quantity_limits"`
	SanityRate             string       `json:"sanity_rate" yaml:"sanity_rate"`
	SanityMarginPercentage string       `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
	AllowSells             string       `json:"allow_sells" yaml:"allow_sells"`
	BatchBlocks            string       `json:"batch_blocks" yaml:"batch_blocks"`
	OutcomePayment         string       `json:"outcome_payment" yaml:"outcome_payment"`
	BondDid                string       `json:"bond_did" yaml:"bond_did"`
	CreatorDid             string       `json:"creator_did" yaml:"creator_did"`
}

func createBondRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createBondReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// Parse function parameters
		functionParams, err := client.ParseFunctionParams(req.FunctionParameters)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
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
		feeAddress, err2 := sdk.AccAddressFromBech32(req.FeeAddress)
		if err2 != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err2.Error())
			return
		}

		// Parse max supply
		maxSupply, err2 := sdk.ParseCoin(req.MaxSupply)
		if err2 != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err2.Error())
			return
		}

		// Parse order quantity limits
		orderQuantityLimits, err2 := sdk.ParseCoins(req.OrderQuantityLimits)
		if err2 != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err2.Error())
			return
		}

		// Parse sanity rate
		sanityRate, err := sdk.NewDecFromStr(req.SanityRate)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse sanity margin percentage
		sanityMarginPercentage, err := sdk.NewDecFromStr(req.SanityMarginPercentage)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
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
			err := sdkerrors.Wrap(types.ErrArgumentMissingOrNonBooleant, "allow_sells")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse batch blocks
		batchBlocks, err2 := sdk.ParseUint(req.BatchBlocks)
		if err2 != nil {
			err := sdkerrors.Wrap(types.ErrArgumentMissingOrNonUInteger, "max batch blocks")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse outcome payment
		outcomePayment, err2 := sdk.ParseCoins(req.OutcomePayment)
		if err2 != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err2.Error())
			return
		}

		msg := types.NewMsgCreateBond(req.Token, req.Name, req.Description,
			req.CreatorDid, req.FunctionType, functionParams, reserveTokens,
			txFeePercentageDec, exitFeePercentageDec, feeAddress, maxSupply,
			orderQuantityLimits, sanityRate, sanityMarginPercentage,
			allowSells, batchBlocks, outcomePayment, req.BondDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type editBondReq struct {
	BaseReq                rest.BaseReq `json:"base_req" yaml:"base_req"`
	Token                  string       `json:"token" yaml:"token"`
	Name                   string       `json:"name" yaml:"name"`
	Description            string       `json:"description" yaml:"description"`
	OrderQuantityLimits    string       `json:"order_quantity_limits" yaml:"order_quantity_limits"`
	SanityRate             string       `json:"sanity_rate" yaml:"sanity_rate"`
	SanityMarginPercentage string       `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
	BondDid                string       `json:"bond_did" yaml:"bond_did"`
	EditorDid              string       `json:"editor_did" yaml:"editor_did"`
}

func editBondRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req editBondReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgEditBond(req.Token, req.Name, req.Description,
			req.OrderQuantityLimits, req.SanityRate,
			req.SanityMarginPercentage, req.EditorDid, req.BondDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
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

func buyRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req buyReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		bondCoin, err := client.ParseTwoPartCoin(req.BondAmount, req.BondToken)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		maxPrices, err := sdk.ParseCoins(req.MaxPrices)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgBuy(req.BuyerDid, bondCoin, maxPrices, req.BondDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type sellReq struct {
	BaseReq    rest.BaseReq `json:"base_req" yaml:"base_req"`
	BondToken  string       `json:"bond_token" yaml:"bond_token"`
	BondAmount string       `json:"bond_amount" yaml:"bond_amount"`
	BondDid    string       `json:"bond_did" yaml:"bond_did"`
	SellerDid  string       `json:"seller_did" yaml:"seller_did"`
}

func sellRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req sellReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		bondCoin, err := client.ParseTwoPartCoin(req.BondAmount, req.BondToken)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgSell(req.SellerDid, bondCoin, req.BondDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
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

func swapRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req swapReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// Check that from amount and token can be parsed to a coin
		fromCoin, err := client.ParseTwoPartCoin(req.FromAmount, req.FromToken)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgSwap(req.SwapperDid, fromCoin, req.ToToken, req.BondDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type makeOutcomePaymentReq struct {
	BaseReq   rest.BaseReq `json:"base_req" yaml:"base_req"`
	BondDid   string       `json:"bond_did" yaml:"bond_did"`
	SenderDid string       `json:"sender_did" yaml:"sender_did"`
}

func makeOutcomePaymentRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req makeOutcomePaymentReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgMakeOutcomePayment(req.SenderDid, req.BondDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type withdrawShareReq struct {
	BaseReq      rest.BaseReq `json:"base_req" yaml:"base_req"`
	BondDid      string       `json:"bond_did" yaml:"bond_did"`
	RecipientDid string       `json:"recipient_did" yaml:"recipient_did"`
}

func withdrawShareRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req withdrawShareReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgWithdrawShare(req.RecipientDid, req.BondDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
