package rest

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/project/project", createProjectRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/project/update_project_status", updateProjectStatusRequestHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/project/create_agent", createAgentRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/project/update_agent", updateAgentRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/project/create_claim", createClaimRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/project/create_evaluation", createEvaluationRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/project/withdraw_funds", withdrawFundsRequestHandler(cliCtx)).Methods("POST")
}

type createProjectReq struct {
	BaseReq    rest.BaseReq    `json:"base_req" yaml:"base_req"`
	TxHash     string          `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did         `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did         `json:"projectDid" yaml:"projectDid"`
	PubKey     string          `json:"pubKey" yaml:"pubKey"`
	Data       json.RawMessage `json:"data" yaml:"data"`
}

func createProjectRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createProjectReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateProject(req.SenderDid, req.Data, req.ProjectDid, req.PubKey)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type updateProjectStatusReq struct {
	BaseReq    rest.BaseReq                 `json:"base_req" yaml:"base_req"`
	TxHash     string                       `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did                      `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did                      `json:"projectDid" yaml:"projectDid"`
	Data       types.UpdateProjectStatusDoc `json:"data" yaml:"data"`
}

func updateProjectStatusRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateProjectStatusReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgUpdateProjectStatus(req.SenderDid, req.Data, req.ProjectDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type createAgentReq struct {
	BaseReq    rest.BaseReq         `json:"base_req" yaml:"base_req"`
	TxHash     string               `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did              `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did              `json:"projectDid" yaml:"projectDid"`
	Data       types.CreateAgentDoc `json:"data" yaml:"data"`
}

func createAgentRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createAgentReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateAgent(req.TxHash, req.SenderDid, req.Data, req.ProjectDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type updateAgentReq struct {
	BaseReq    rest.BaseReq         `json:"base_req" yaml:"base_req"`
	TxHash     string               `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did              `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did              `json:"projectDid" yaml:"projectDid"`
	Data       types.UpdateAgentDoc `json:"data" yaml:"data"`
}

func updateAgentRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateAgentReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgUpdateAgent(req.TxHash, req.SenderDid, req.Data, req.ProjectDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type createClaimReq struct {
	BaseReq    rest.BaseReq         `json:"base_req" yaml:"base_req"`
	TxHash     string               `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did              `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did              `json:"projectDid" yaml:"projectDid"`
	Data       types.CreateClaimDoc `json:"data" yaml:"data"`
}

func createClaimRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createClaimReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateClaim(req.TxHash, req.SenderDid, req.Data, req.ProjectDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})

	}
}

type createEvaluationReq struct {
	BaseReq    rest.BaseReq              `json:"base_req" yaml:"base_req"`
	TxHash     string                    `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did                   `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did                   `json:"projectDid" yaml:"projectDid"`
	Data       types.CreateEvaluationDoc `json:"data" yaml:"data"`
}

func createEvaluationRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createEvaluationReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateEvaluation(req.TxHash, req.SenderDid, req.Data, req.ProjectDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type withdrawFundsReq struct {
	BaseReq   rest.BaseReq           `json:"base_req" yaml:"base_req"`
	SenderDid did.Did                `json:"senderDid" yaml:"senderDid"`
	Data      types.WithdrawFundsDoc `json:"data" yaml:"data"`
}

func withdrawFundsRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req withdrawFundsReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgWithdrawFunds(req.SenderDid, req.Data)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
