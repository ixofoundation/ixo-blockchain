package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
)

func NewMsgCreateProject(projectDoc ProjectDoc, projectDid sovrin.SovrinDid) MsgCreateProject {
	return MsgCreateProject{
		SignBytes:  "",
		TxHash:     "",
		SenderDid:  "",
		ProjectDid: projectDid.Did,
		PubKey:     projectDid.VerifyKey,
		Data:       projectDoc,
	}
}

func NewMsgUpdateProjectStatus(senderDid string, updateProjectStatusDoc UpdateProjectStatusDoc, projectDid sovrin.SovrinDid) MsgUpdateProjectStatus {
	return MsgUpdateProjectStatus{
		SignBytes:  "",
		TxHash:     "",
		SenderDid:  senderDid,
		ProjectDid: projectDid.Did,
		Data:       updateProjectStatusDoc,
	}
}

func NewMsgCreateAgent(txHash string, senderDid string, createAgentDoc CreateAgentDoc, projectDid sovrin.SovrinDid) MsgCreateAgent {
	return MsgCreateAgent{
		SignBytes:  "",
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createAgentDoc,
	}
}

func NewMsgUpdateAgent(txHash string, senderDid string, updateAgentDoc UpdateAgentDoc, projectDid sovrin.SovrinDid) MsgUpdateAgent {
	return MsgUpdateAgent{
		SignBytes:  "",
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       updateAgentDoc,
	}
}

func NewMsgCreateClaim(txHash string, senderDid string, createClaimDoc CreateClaimDoc, projectDid sovrin.SovrinDid) MsgCreateClaim {
	return MsgCreateClaim{
		SignBytes:  "",
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createClaimDoc,
	}
}

func NewMsgCreateEvaluation(txHash string, senderDid string, createEvaluationDoc CreateEvaluationDoc, projectDid sovrin.SovrinDid) MsgCreateEvaluation {
	return MsgCreateEvaluation{
		SignBytes:  "",
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createEvaluationDoc,
	}
}

func NewMsgWithdrawFunds(senderDid ixo.Did, data WithdrawFundsDoc) MsgWithdrawFunds {
	return MsgWithdrawFunds{
		SignBytes: "",
		SenderDid: senderDid,
		Data:      data,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if len(value) == 0 {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}
