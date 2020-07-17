package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/client"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"

	"net/http"
	"strings"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/bonds/create_bond",
		createBondHandler(cliCtx),
	).Methods("POST")

	r.HandleFunc(
		"/bonds/edit_bond",
		editBondHandler(cliCtx),
	).Methods("POST")

	r.HandleFunc(
		"/bonds/buy",
		buyHandler(cliCtx),
	).Methods("POST")

	r.HandleFunc(
		"/bonds/sell",
		sellHandler(cliCtx),
	).Methods("POST")

	r.HandleFunc(
		"/bonds/swap",
		swapHandler(cliCtx),
	).Methods("POST")
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
	BondDid                string       `json:"bond_did" yaml:"bond_did"`
	CreatorDid             string       `json:"creator_did" yaml:"creator_did"`
}

func createBondHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createBondReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
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
			err = types.ErrArgumentMissingOrNonFloat(types.DefaultCodespace, "tx fee percentage")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse exit fee percentage
		exitFeePercentageDec, err := sdk.NewDecFromStr(req.ExitFeePercentage)
		if err != nil {
			err = types.ErrArgumentMissingOrNonFloat(types.DefaultCodespace, "exit fee percentage")
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

		// Parse batch blocks
		batchBlocks, err2 := sdk.ParseUint(req.BatchBlocks)
		if err2 != nil {
			err := types.ErrArgumentMissingOrNonUInteger(types.DefaultCodespace, "max batch blocks")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse creator's ixo DID
		creatorDid, err2 := did.UnmarshalIxoDid(req.CreatorDid)
		if err2 != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err2.Error())
			return
		}

		msg := types.NewMsgCreateBond(req.Token, req.Name, req.Description,
			creatorDid.Did, req.FunctionType, functionParams, reserveTokens,
			txFeePercentageDec, exitFeePercentageDec, feeAddress, maxSupply,
			orderQuantityLimits, sanityRate, sanityMarginPercentage,
			req.AllowSells, batchBlocks, req.BondDid)

		output, err2 := ixo.CompleteAndBroadcastTxRest(cliCtx, msg, creatorDid)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err2.Error()))
			return
		}

		rest.PostProcessResponse(w, cliCtx, output)
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

func editBondHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req editBondReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// Parse editor's ixo DID
		editorDid, err := did.UnmarshalIxoDid(req.EditorDid)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgEditBond(req.Token, req.Name, req.Description,
			req.OrderQuantityLimits, req.SanityRate,
			req.SanityMarginPercentage, editorDid.Did, req.BondDid)

		output, err := ixo.CompleteAndBroadcastTxRest(cliCtx, msg, editorDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		rest.PostProcessResponse(w, cliCtx, output)
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

func buyHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req buyReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
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

		// Parse buyer's ixo DID
		buyerDid, err := did.UnmarshalIxoDid(req.BuyerDid)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgBuy(buyerDid.Did, bondCoin, maxPrices, req.BondDid)

		output, err := ixo.CompleteAndBroadcastTxRest(cliCtx, msg, buyerDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		rest.PostProcessResponse(w, cliCtx, output)
	}
}

type sellReq struct {
	BaseReq    rest.BaseReq `json:"base_req" yaml:"base_req"`
	BondToken  string       `json:"bond_token" yaml:"bond_token"`
	BondAmount string       `json:"bond_amount" yaml:"bond_amount"`
	BondDid    string       `json:"bond_did" yaml:"bond_did"`
	SellerDid  string       `json:"seller_did" yaml:"seller_did"`
}

func sellHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req sellReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
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

		// Parse seller's ixo DID
		sellerDid, err := did.UnmarshalIxoDid(req.SellerDid)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgSell(sellerDid.Did, bondCoin, req.BondDid)

		output, err := ixo.CompleteAndBroadcastTxRest(cliCtx, msg, sellerDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		rest.PostProcessResponse(w, cliCtx, output)
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

func swapHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req swapReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
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

		// Parse swapper's ixo DID
		swapperDid, err := did.UnmarshalIxoDid(req.SwapperDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		msg := types.NewMsgSwap(swapperDid.Did, fromCoin, req.ToToken, req.BondDid)

		output, err := ixo.CompleteAndBroadcastTxRest(cliCtx, msg, swapperDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		rest.PostProcessResponse(w, cliCtx, output)
	}
}
