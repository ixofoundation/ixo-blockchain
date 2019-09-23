package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/project", createProjectRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/updateProjectStatus", updateProjectStatusRequestHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/createAgent", CreateAgentRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/createClaim", CreateClaimRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/createEvaluation", CreateEvaluationRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/withdrawFunds", WithDrawFundsRequestHandler(cliCtx)).Methods("POST")
}

func createProjectRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		projectDocParam := r.URL.Query().Get("projectDoc")
		didDocParam := r.URL.Query().Get("didDoc")
		mode := r.URL.Query().Get("mode")
		var projectDoc types.ProjectDoc
		err := json.Unmarshal([]byte(projectDocParam), &projectDoc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall projectDoc into struct. Error: %s", err.Error())))
			
			return
		}
		
		var didDoc sovrin.SovrinDid
		err = json.Unmarshal([]byte(didDocParam), &didDoc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))
			
			return
		}
		
		cliCtx = cliCtx.WithBroadcastMode(mode)
		msg := types.NewCreateProjectMsg(projectDoc, didDoc)
		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(didDoc.Secret.SignKey))
		copy(privKey[32:], base58.Decode(didDoc.VerifyKey))
		
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall msg to json. Error: %s", err.Error())))
			return
		}
		
		signature := ixo.SignIxoMessage(msgBytes, didDoc.Did, privKey)
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

func updateProjectStatusRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		txHash := r.URL.Query().Get("txHash")
		senderDid := r.URL.Query().Get("senderDid")
		status := r.URL.Query().Get("status")
		sovrinDidParam := r.URL.Query().Get("sovrinDid")
		mode := r.URL.Query().Get("mode")
		
		var sovrinDid sovrin.SovrinDid
		err := json.Unmarshal([]byte(sovrinDidParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall sovrinDid into struct. Error: %s", err.Error())))
			return
		}
		
		cliCtx = cliCtx.WithBroadcastMode(mode)
		
		projectStatus := types.ProjectStatus(status)
		if projectStatus != types.CreatedProject &&
			projectStatus != types.PendingStatus &&
			projectStatus != types.FundedStatus &&
			projectStatus != types.StartedStatus &&
			projectStatus != types.StoppedStatus &&
			projectStatus != types.PaidoutStatus {
			_, _ = w.Write([]byte("The status must be one of 'CREATED', 'PENDING', 'FUNDED', 'STARTED', 'STOPPED' or 'PAIDOUT'"))
			
			return
		}
		
		updateProjectStatusDoc := types.UpdateProjectStatusDoc{
			Status:          projectStatus,
			EthFundingTxnID: txHash,
		}
		
		msg := types.NewUpdateProjectStatusMsg(txHash, senderDid, updateProjectStatusDoc, sovrinDid)
		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))
		
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall msg to json. Error: %s", err.Error())))
			
			return
		}
		signature := ixo.SignIxoMessage(msgBytes, sovrinDid.Did, privKey)
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

func CreateAgentRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		txHash := r.URL.Query().Get("txHash")
		senderDid := r.URL.Query().Get("senderDid")
		agentDid := r.URL.Query().Get("agentDid")
		role := r.URL.Query().Get("role")
		projectDidParam := r.URL.Query().Get("projectDid")
		mode := r.URL.Query().Get("mode")
		
		var projectDid sovrin.SovrinDid
		err := json.Unmarshal([]byte(projectDidParam), &projectDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall projectDid into struct. Error: %s", err.Error())))
			
			return
		}
		
		cliCtx = cliCtx.WithBroadcastMode(mode)
		
		if role != "SA" && role != "EA" && role != "IA" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "The role must be one of 'SA', 'EA' or 'IA'")
			
			return
		}
		
		createAgentDoc := types.CreateAgentDoc{
			AgentDid: agentDid,
			Role:     role,
		}
		
		msg := types.NewCreateAgentMsg(txHash, senderDid, createAgentDoc, projectDid)
		
		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(projectDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(projectDid.VerifyKey))
		
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall msg to json. Error: %s", err.Error())))
			
			return
		}
		signature := ixo.SignIxoMessage(msgBytes, projectDid.Did, privKey)
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

func UpdateAgent() {
	// TODO
}

func CreateClaimRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		txHash := r.URL.Query().Get("txHash")
		senderDid := r.URL.Query().Get("senderDid")
		claimId := r.URL.Query().Get("claimId")
		sovrinDidParam := r.URL.Query().Get("sovrinDid")
		mode := r.URL.Query().Get("mode")
		
		var sovrinDid sovrin.SovrinDid
		err := json.Unmarshal([]byte(sovrinDidParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall SovrinDid into struct. Error: %s", err.Error())))
			
			return
		}
		
		createClaimDoc := types.CreateClaimDoc{
			ClaimID: claimId,
		}
		
		cliCtx = cliCtx.WithBroadcastMode(mode)
		
		msg := types.NewCreateClaimMsg(txHash, senderDid, createClaimDoc, sovrinDid)
		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))
		
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall msg to json. Error: %s", err.Error())))
			
			return
		}
		signature := ixo.SignIxoMessage(msgBytes, sovrinDid.Did, privKey)
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

func CreateEvaluationRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		txHash := r.URL.Query().Get("txHash")
		senderDid := r.URL.Query().Get("senderDid")
		claimDid := r.URL.Query().Get("claimDid")
		status := r.URL.Query().Get("status")
		sovrinDidParam := r.URL.Query().Get("sovrinDid")
		mode := r.URL.Query().Get("mode")
		
		var sovrinDid sovrin.SovrinDid
		err := json.Unmarshal([]byte(sovrinDidParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall sovrinDid into struct. Error: %s", err.Error())))
			
			return
		}
		
		cliCtx = cliCtx.WithBroadcastMode(mode)
		
		claimStatus := types.ClaimStatus(status)
		if claimStatus != types.PendingClaim && claimStatus != types.ApprovedClaim && claimStatus != types.RejectedClaim {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "The status must be one of '0' (Pending), '1' (Approved) or '2' (Rejected)")
			
			return
		}
		
		createEvaluationDoc := types.CreateEvaluationDoc{
			ClaimID: claimDid,
			Status:  claimStatus,
		}
		
		msg := types.NewCreateEvaluationMsg(txHash, senderDid, createEvaluationDoc, sovrinDid)
		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))
		
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall msg to json. Error: %s", err.Error())))
			
			return
		}
		signature := ixo.SignIxoMessage(msgBytes, sovrinDid.Did, privKey)
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

func WithDrawFundsRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		senderDidParams := r.URL.Query().Get("senderDid")
		dataParams := r.URL.Query().Get("data")
		mode := r.URL.Query().Get("mode")
		
		var senderDid sovrin.SovrinDid
		err := json.Unmarshal([]byte(senderDidParams), &senderDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall sovrinDid into struct. Error: %s", err.Error())))
			return
		}
		
		var data types.WithdrawFundsDoc
		err = json.Unmarshal([]byte(dataParams), &data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall data into struct. Error: %s", err.Error())))
			return
		}
		
		cliCtx = cliCtx.WithBroadcastMode(mode)
		
		msg := types.NewWithDrawFundsMsg(senderDid.Did, data)
		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(senderDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(senderDid.VerifyKey))
		
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		signature := ixo.SignIxoMessage(msgBytes, senderDid.Did, privKey)
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
