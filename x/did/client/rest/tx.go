package rest

import (
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/ixofoundation/ixo-blockchain/x/did/types"
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/did/add_did", newAddDidRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/did/add_credential", newAddCredentialRequestHandler(clientCtx)).Methods("POST")
}

type addDidReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Did     exported.Did `json:"did" yaml:"did"`
	PubKey  string       `json:"pubKey" yaml:"pubKey"`
}

func newAddDidRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req addDidReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgAddDid(req.Did, req.PubKey)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) { //err := msg.ValidateBasic(); err != nil {
			//rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type addCredentialReq struct {
	BaseReq       rest.BaseReq        `json:"base_req" yaml:"base_req"`
	Did           exported.Did        `json:"did" yaml:"did"`
	DidCredential types.DidCredential `json:"credential" yaml:"credential"`
}

func newAddCredentialRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req addCredentialReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgAddCredential(req.Did, req.DidCredential.Credtype, req.Did, req.DidCredential.Issued)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) { //err := msg.ValidateBasic(); err != nil {
			//rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
