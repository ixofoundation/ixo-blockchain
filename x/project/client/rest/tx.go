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
	r.HandleFunc("/project/project", addProjectRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/project/update_project_status", updateProjectStatusRequestHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/project/create_agent", addAgentRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/project/create_claim", addClaimRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/project/create_evaluation", addEvaluationRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/project/withdraw_funds", withdrawFundsRequestHandler(cliCtx)).Methods("POST")
}

type CreateProjectReq struct {
	BaseReq    rest.BaseReq    `json:"base_req" yaml:"base_req"`
	TxHash     string          `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did         `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did         `json:"projectDid" yaml:"projectDid"`
	PubKey     string          `json:"pubKey" yaml:"pubKey"`
	Data       json.RawMessage `json:"data" yaml:"data"`
}

func addProjectRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateProjectReq
		req.BaseReq = req.BaseReq.Sanitize()
		ProjectDid, err := did.UnmarshalIxoDid(req.ProjectDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		msg := types.NewMsgCreateProject(req.SenderDid, req.Data, ProjectDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type UpdateProjectStatusReq struct {
	BaseReq    rest.BaseReq                 `json:"base_req" yaml:"base_req"`
	TxHash     string                       `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did                      `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did                      `json:"projectDid" yaml:"projectDid"`
	Data       types.UpdateProjectStatusDoc `json:"data" yaml:"data"`
}

func updateProjectStatusRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateProjectStatusReq
		req.BaseReq = req.BaseReq.Sanitize()
		ProjectDid, err := did.UnmarshalIxoDid(req.ProjectDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		msg := types.NewMsgUpdateProjectStatus(req.SenderDid, req.Data, ProjectDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type CreateAgentReq struct {
	BaseReq    rest.BaseReq         `json:"base_req" yaml:"base_req"`
	TxHash     string               `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did              `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did              `json:"projectDid" yaml:"projectDid"`
	Data       types.CreateAgentDoc `json:"data" yaml:"data"`
}

func addAgentRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateAgentReq
		req.BaseReq = req.BaseReq.Sanitize()
		ProjectDid, err := did.UnmarshalIxoDid(req.ProjectDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		msg := types.NewMsgCreateAgent(req.TxHash, req.SenderDid, req.Data, ProjectDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type CreateClaimReq struct {
	BaseReq    rest.BaseReq         `json:"base_req" yaml:"base_req"`
	TxHash     string               `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did              `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did              `json:"projectDid" yaml:"projectDid"`
	Data       types.CreateClaimDoc `json:"data" yaml:"data"`
}

func addClaimRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateClaimReq
		req.BaseReq = req.BaseReq.Sanitize()
		ProjectDid, err := did.UnmarshalIxoDid(req.ProjectDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		msg := types.NewMsgCreateClaim(req.TxHash, req.SenderDid, req.Data, ProjectDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})

	}
}

type CreateEvaluationReq struct {
	BaseReq    rest.BaseReq              `json:"base_req" yaml:"base_req"`
	TxHash     string                    `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did                   `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did                   `json:"projectDid" yaml:"projectDid"`
	Data       types.CreateEvaluationDoc `json:"data" yaml:"data"`
}

func addEvaluationRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateEvaluationReq
		req.BaseReq = req.BaseReq.Sanitize()
		ProjectDid, err := did.UnmarshalIxoDid(req.ProjectDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		msg := types.NewMsgCreateEvaluation(req.TxHash, req.SenderDid, req.Data, ProjectDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type WithdrawFundsReq struct {
	BaseReq   rest.BaseReq           `json:"base_req" yaml:"base_req"`
	SenderDid did.Did                `json:"senderDid" yaml:"senderDid"`
	Data      types.WithdrawFundsDoc `json:"data" yaml:"data"`
}

func withdrawFundsRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req WithdrawFundsReq
		req.BaseReq = req.BaseReq.Sanitize()
		msg := types.NewMsgWithdrawFunds(req.SenderDid, req.Data)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
