package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	//"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/tx"
	//"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
)

// TODO We can copy WriteGenerateStdTxResponse from cosmos-sdk/x/auth/client/utils v0.39.1

func registerTxRoutes(cliCtx /*context*/client.Context, r *mux.Router) {
	r.HandleFunc("/did/add_did", addDidRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/did/add_credential", addCredentialRequestHandler(cliCtx)).Methods("POST")
}

type addDidReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Did     exported.Did `json:"did" yaml:"did"`
	PubKey  string       `json:"pubKey" yaml:"pubKey"`
}

func addDidRequestHandler(cliCtx /*context*/client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req addDidReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino/*Codec*/, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgAddDid(req.Did, req.PubKey)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// /*utils.*/WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
		tx.WriteGeneratedTxResponse(cliCtx, w, req.BaseReq, []sdk.Msg{msg}...)
	}
}

type addCredentialReq struct {
	BaseReq       rest.BaseReq           `json:"base_req" yaml:"base_req"`
	Did           exported.Did           `json:"did" yaml:"did"`
	DidCredential exported.DidCredential `json:"credential" yaml:"credential"`
}

func addCredentialRequestHandler(cliCtx /*context*/client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req addCredentialReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino/*Codec*/, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgAddCredential(req.Did, req.DidCredential.CredType, req.Did, req.DidCredential.Issued)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// /*utils.*/WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
		tx.WriteGeneratedTxResponse(cliCtx, w, req.BaseReq, []sdk.Msg{msg}...)
	}
}
