package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"net/http"

	"github.com/ixofoundation/ixo-blockchain/x/treasury/internal/types"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/treasury/send", sendRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/treasury/oracle_transfer", oracleTransferRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/treasury/oracle_mint", oracleMintRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/treasury/oracle_burn", oracleBurnRequestHandler(cliCtx)).Methods("POST")
}

type sendReq struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	FromDid     did.Did      `json:"from_did" yaml:"from_did"`
	ToDidOrAddr did.Did      `json:"to_did" yaml:"to_did"`
	Amount      sdk.Coins    `json:"amount" yaml:"amount"`
}

func sendRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req sendReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgSend(req.ToDidOrAddr, req.Amount, req.FromDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type oracleTransferReq struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	OracleDid   did.Did      `json:"oracle_did" yaml:"oracle_did"`
	FromDid     did.Did      `json:"from_did" yaml:"from_did"`
	ToDidOrAddr did.Did      `json:"to_did" yaml:"to_did"`
	Amount      sdk.Coins    `json:"amount" yaml:"amount"`
	Proof       string       `json:"proof" yaml:"proof"`
}

func oracleTransferRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req oracleTransferReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgOracleTransfer(req.FromDid, req.ToDidOrAddr, req.Amount, req.OracleDid, req.Proof)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type oracleMintReq struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	OracleDid   did.Did      `json:"oracle_did" yaml:"oracle_did"`
	ToDidOrAddr did.Did      `json:"to_did" yaml:"to_did"`
	Amount      sdk.Coins    `json:"amount" yaml:"amount"`
	Proof       string       `json:"proof" yaml:"proof"`
}

func oracleMintRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req oracleMintReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgOracleMint(
			req.ToDidOrAddr, req.Amount, req.OracleDid, req.Proof)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type oracleBurnReq struct {
	BaseReq   rest.BaseReq `json:"base_req" yaml:"base_req"`
	OracleDid did.Did      `json:"oracle_did" yaml:"oracle_did"`
	FromDid   did.Did      `json:"from_did" yaml:"from_did"`
	Amount    sdk.Coins    `json:"amount" yaml:"amount"`
	Proof     string       `json:"proof" yaml:"proof"`
}

func oracleBurnRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req oracleBurnReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgOracleBurn(req.FromDid, req.Amount, req.OracleDid, req.Proof)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
