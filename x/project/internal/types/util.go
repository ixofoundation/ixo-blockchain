package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"strings"
)

func NewMsgCreateProject(senderDid did.Did, projectData json.RawMessage,
	projectDid did.Did, pubKey string) MsgCreateProject {
	return MsgCreateProject{
		TxHash:     "",
		SenderDid:  senderDid,
		ProjectDid: projectDid,
		PubKey:     pubKey,
		Data:       projectData,
	}
}

func NewMsgUpdateProjectStatus(senderDid did.Did, updateProjectStatusDoc UpdateProjectStatusDoc, projectDid did.Did) MsgUpdateProjectStatus {
	return MsgUpdateProjectStatus{
		TxHash:     "",
		SenderDid:  senderDid,
		ProjectDid: projectDid,
		Data:       updateProjectStatusDoc,
	}
}

func NewMsgCreateAgent(txHash string, senderDid did.Did, createAgentDoc CreateAgentDoc, projectDid did.Did) MsgCreateAgent {
	return MsgCreateAgent{
		ProjectDid: projectDid,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createAgentDoc,
	}
}

func NewMsgUpdateAgent(txHash string, senderDid did.Did, updateAgentDoc UpdateAgentDoc, projectDid did.Did) MsgUpdateAgent {
	return MsgUpdateAgent{
		ProjectDid: projectDid,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       updateAgentDoc,
	}
}

func NewMsgCreateClaim(txHash string, senderDid did.Did, createClaimDoc CreateClaimDoc, projectDid did.Did) MsgCreateClaim {
	return MsgCreateClaim{
		ProjectDid: projectDid,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createClaimDoc,
	}
}

func NewMsgCreateEvaluation(txHash string, senderDid did.Did, createEvaluationDoc CreateEvaluationDoc, projectDid did.Did) MsgCreateEvaluation {
	return MsgCreateEvaluation{
		ProjectDid: projectDid,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createEvaluationDoc,
	}
}

func NewMsgWithdrawFunds(senderDid did.Did, data WithdrawFundsDoc) MsgWithdrawFunds {
	return MsgWithdrawFunds{
		SenderDid: senderDid,
		Data:      data,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if strings.TrimSpace(value) == "" {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}
