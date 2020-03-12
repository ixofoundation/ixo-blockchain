package rest

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/client"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"net/http"
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

		// Check that bond token is a valid token name
		err := client.CheckCoinDenom(req.Token)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse function parameters
		functionParams, err := client.ParseFunctionParams(req.FunctionParameters, req.FunctionType)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse reserve tokens
		reserveTokens, err2 := client.ParseReserveTokens(req.ReserveTokens, req.FunctionType, req.Token)
		if err2 != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err2.Error())
			return
		}

		txFeePercentageDec, err := sdk.NewDecFromStr(req.TxFeePercentage)
		if err != nil {
			err = types.ErrArgumentMissingOrNonFloat(types.DefaultCodespace, "tx fee percentage")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		exitFeePercentageDec, err := sdk.NewDecFromStr(req.ExitFeePercentage)
		if err != nil {
			err = types.ErrArgumentMissingOrNonFloat(types.DefaultCodespace, "exit fee percentage")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if txFeePercentageDec.Add(exitFeePercentageDec).GTE(sdk.NewDec(100)) {
			err = types.ErrFeesCannotBeOrExceed100Percent(types.DefaultCodespace)
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		feeAddress, err2 := sdk.AccAddressFromBech32(req.FeeAddress)
		if err2 != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err2.Error())
			return
		}

		maxSupply, err2 := client.ParseMaxSupply(req.MaxSupply, req.Token)
		if err2 != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err2.Error())
			return
		}

		orderQuantityLimits, err2 := sdk.ParseCoins(req.OrderQuantityLimits)
		if err2 != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err2.Error())
			return
		}

		// Parse sanity
		sanityRate, sanityMarginPercentage, err := client.ParseSanityValues(req.SanityRate, req.SanityMarginPercentage)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse batch blocks
		batchBlocks, err2 := client.ParseBatchBlocks(req.BatchBlocks)
		if err2 != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err2.Error())
			return
		}

		// Parse bond's sovrin DID
		bondDid := client.UnmarshalSovrinDID(req.BondDid)

		msg := types.NewMsgCreateBond(req.Token, req.Name, req.Description,
			req.CreatorDid, req.FunctionType, functionParams, reserveTokens,
			txFeePercentageDec, exitFeePercentageDec, feeAddress, maxSupply,
			orderQuantityLimits, sanityRate, sanityMarginPercentage,
			req.AllowSells, batchBlocks, bondDid)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(bondDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(bondDid.VerifyKey))

		msgBytes, err2 := json.Marshal(msg)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall msg to json. Error: %s", err2.Error())))
			return
		}

		signature := ixo.SignIxoMessage(msgBytes, bondDid.Did, privKey)
		tx := ixo.NewIxoTxSingleMsg(msg, signature)

		bz, err2 := cliCtx.Codec.MarshalJSON(tx)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall tx to binary. Error: %s", err2.Error())))

			return
		}

		res, err2 := cliCtx.BroadcastTx(bz)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not broadcast tx. Error: %s", err2.Error())))

			return
		}

		output, err2 := json.MarshalIndent(res, "", "  ")
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

		// Parse bond's sovrin DID
		bondDid := client.UnmarshalSovrinDID(req.BondDid)

		msg := types.NewMsgEditBond(req.Token, req.Name, req.Description,
			req.OrderQuantityLimits, req.SanityRate,
			req.SanityMarginPercentage, req.EditorDid, bondDid)
		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(bondDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(bondDid.VerifyKey))

		msgBytes, err2 := json.Marshal(msg)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall msg to json. Error: %s", err2.Error())))
			return
		}

		signature := ixo.SignIxoMessage(msgBytes, bondDid.Did, privKey)
		tx := ixo.NewIxoTxSingleMsg(msg, signature)

		bz, err2 := cliCtx.Codec.MarshalJSON(tx)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall tx to binary. Error: %s", err2.Error())))

			return
		}

		res, err2 := cliCtx.BroadcastTx(bz)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not broadcast tx. Error: %s", err2.Error())))

			return
		}

		output, err2 := json.MarshalIndent(res, "", "  ")
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err2.Error()))

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

		bondCoin, err := client.ParseCoin(req.BondAmount, req.BondToken)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		maxPrices, err := sdk.ParseCoins(req.MaxPrices)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse buyer's sovrin DID
		buyerDid := client.UnmarshalSovrinDID(req.BuyerDid)

		msg := types.NewMsgBuy(buyerDid, bondCoin, maxPrices, req.BondDid)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(buyerDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(buyerDid.VerifyKey))

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall msg to json. Error: %s", err.Error())))
			return
		}

		signature := ixo.SignIxoMessage(msgBytes, buyerDid.Did, privKey)
		tx := ixo.NewIxoTxSingleMsg(msg, signature)

		bz, err := cliCtx.Codec.MarshalJSON(tx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall tx to binary. Error: %s", err.Error())))

			return
		}

		res, err := cliCtx.BroadcastTx(bz)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not broadcast tx. Error: %s", err.Error())))

			return
		}

		output, err := json.MarshalIndent(res, "", "  ")
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

		bondCoin, err := client.ParseCoin(req.BondAmount, req.BondToken)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse seller's sovrin DID
		sellerDid := client.UnmarshalSovrinDID(req.SellerDid)

		msg := types.NewMsgSell(sellerDid, bondCoin, req.BondDid)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(sellerDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(sellerDid.VerifyKey))

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall msg to json. Error: %s", err.Error())))
			return
		}

		signature := ixo.SignIxoMessage(msgBytes, sellerDid.Did, privKey)
		tx := ixo.NewIxoTxSingleMsg(msg, signature)

		bz, err := cliCtx.Codec.MarshalJSON(tx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall tx to binary. Error: %s", err.Error())))

			return
		}

		res, err := cliCtx.BroadcastTx(bz)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not broadcast tx. Error: %s", err.Error())))

			return
		}

		output, err := json.MarshalIndent(res, "", "  ")
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
		fromCoin, err := client.ParseCoin(req.FromAmount, req.FromToken)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Check that ToToken is a valid token name
		err = client.CheckCoinDenom(req.ToToken)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse swapper's sovrin DID
		swapperDid := client.UnmarshalSovrinDID(req.SwapperDid)

		msg := types.NewMsgSwap(swapperDid, fromCoin, req.ToToken, req.BondDid)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(swapperDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(swapperDid.VerifyKey))

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall msg to json. Error: %s", err.Error())))
			return
		}

		signature := ixo.SignIxoMessage(msgBytes, swapperDid.Did, privKey)
		tx := ixo.NewIxoTxSingleMsg(msg, signature)

		bz, err := cliCtx.Codec.MarshalJSON(tx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall tx to binary. Error: %s", err.Error())))

			return
		}

		res, err := cliCtx.BroadcastTx(bz)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not broadcast tx. Error: %s", err.Error())))

			return
		}

		output, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))

			return
		}

		rest.PostProcessResponse(w, cliCtx, output)
	}
}
