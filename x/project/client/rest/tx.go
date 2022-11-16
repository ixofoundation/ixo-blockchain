package rest

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	didexported "github.com/ixofoundation/ixo-blockchain/lib/legacydid"
	"github.com/ixofoundation/ixo-blockchain/x/project/types"
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/project/project", createProjectRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/project/update_project_status", updateProjectStatusRequestHandler(clientCtx)).Methods("PUT")
	r.HandleFunc("/project/create_agent", createAgentRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/project/update_agent", updateAgentRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/project/create_claim", createClaimRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/project/create_evaluation", createEvaluationRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/project/withdraw_funds", withdrawFundsRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/project/update_project_doc", updateProjectDocRequestHandler(clientCtx)).Methods("PUT")
}

type createProjectReq struct {
	BaseReq    rest.BaseReq    `json:"base_req" yaml:"base_req"`
	TxHash     string          `json:"txHash" yaml:"txHash"`
	SenderDid  didexported.Did `json:"senderDid" yaml:"senderDid"`
	ProjectDid didexported.Did `json:"projectDid" yaml:"projectDid"`
	PubKey     string          `json:"pubKey" yaml:"pubKey"`
	Data       json.RawMessage `json:"iid" yaml:"iid"`
}

func createProjectRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createProjectReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateProject(req.SenderDid, req.Data, req.ProjectDid, req.PubKey, "")
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateProjectStatusReq struct {
	BaseReq    rest.BaseReq                 `json:"base_req" yaml:"base_req"`
	TxHash     string                       `json:"txHash" yaml:"txHash"`
	SenderDid  didexported.Did              `json:"senderDid" yaml:"senderDid"`
	ProjectDid didexported.Did              `json:"projectDid" yaml:"projectDid"`
	Data       types.UpdateProjectStatusDoc `json:"iid" yaml:"iid"`
}

func updateProjectStatusRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateProjectStatusReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgUpdateProjectStatus(req.SenderDid, req.Data, req.ProjectDid, "")
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type createAgentReq struct {
	BaseReq    rest.BaseReq         `json:"base_req" yaml:"base_req"`
	TxHash     string               `json:"txHash" yaml:"txHash"`
	SenderDid  didexported.Did      `json:"senderDid" yaml:"senderDid"`
	ProjectDid didexported.Did      `json:"projectDid" yaml:"projectDid"`
	Data       types.CreateAgentDoc `json:"iid" yaml:"iid"`
}

func createAgentRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createAgentReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateAgent(req.TxHash, req.SenderDid, req.Data, req.ProjectDid, "")
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateAgentReq struct {
	BaseReq    rest.BaseReq         `json:"base_req" yaml:"base_req"`
	TxHash     string               `json:"txHash" yaml:"txHash"`
	SenderDid  didexported.Did      `json:"senderDid" yaml:"senderDid"`
	ProjectDid didexported.Did      `json:"projectDid" yaml:"projectDid"`
	Data       types.UpdateAgentDoc `json:"iid" yaml:"iid"`
}

func updateAgentRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateAgentReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgUpdateAgent(req.TxHash, req.SenderDid, req.Data, req.ProjectDid, "")
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type createClaimReq struct {
	BaseReq    rest.BaseReq         `json:"base_req" yaml:"base_req"`
	TxHash     string               `json:"txHash" yaml:"txHash"`
	SenderDid  didexported.Did      `json:"senderDid" yaml:"senderDid"`
	ProjectDid didexported.Did      `json:"projectDid" yaml:"projectDid"`
	Data       types.CreateClaimDoc `json:"iid" yaml:"iid"`
}

func createClaimRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createClaimReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateClaim(req.TxHash, req.SenderDid, req.Data, req.ProjectDid, "")
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)

	}
}

type createEvaluationReq struct {
	BaseReq    rest.BaseReq              `json:"base_req" yaml:"base_req"`
	TxHash     string                    `json:"txHash" yaml:"txHash"`
	SenderDid  didexported.Did           `json:"senderDid" yaml:"senderDid"`
	ProjectDid didexported.Did           `json:"projectDid" yaml:"projectDid"`
	Data       types.CreateEvaluationDoc `json:"iid" yaml:"iid"`
}

func createEvaluationRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createEvaluationReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateEvaluation(req.TxHash, req.SenderDid, req.Data, req.ProjectDid, "")
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type withdrawFundsReq struct {
	BaseReq   rest.BaseReq           `json:"base_req" yaml:"base_req"`
	SenderDid didexported.Did        `json:"senderDid" yaml:"senderDid"`
	Data      types.WithdrawFundsDoc `json:"iid" yaml:"iid"`
}

func withdrawFundsRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req withdrawFundsReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgWithdrawFunds(req.SenderDid, req.Data, "")
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateProjectDocReq struct {
	BaseReq    rest.BaseReq    `json:"base_req" yaml:"base_req"`
	TxHash     string          `json:"txHash" yaml:"txHash"`
	SenderDid  didexported.Did `json:"senderDid" yaml:"senderDid"`
	ProjectDid didexported.Did `json:"projectDid" yaml:"projectDid"`
	Data       json.RawMessage `json:"iid" yaml:"iid"`
}

func updateProjectDocRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateProjectDocReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgUpdateProjectDoc(req.SenderDid, req.Data, req.ProjectDid, "")
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
