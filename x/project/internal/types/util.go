package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
)

func NewCreateProjectMsg(projectDoc ProjectDoc, projectDid sovrin.SovrinDid) CreateProjectMsg {
	return CreateProjectMsg{
		SignBytes:  "",
		TxHash:     "",
		SenderDid:  "",
		ProjectDid: projectDid.Did,
		PubKey:     projectDid.VerifyKey,
		Data:       projectDoc,
	}
}

func NewUpdateProjectStatusMsg(txHash string, senderDid string, updateProjectStatusDoc UpdateProjectStatusDoc, projectDid sovrin.SovrinDid) UpdateProjectStatusMsg {
	return UpdateProjectStatusMsg{
		SignBytes:  "",
		TxHash:     txHash,
		SenderDid:  senderDid,
		ProjectDid: projectDid.Did,
		Data:       updateProjectStatusDoc,
	}
}

func NewCreateAgentMsg(txHash string, senderDid string, createAgentDoc CreateAgentDoc, projectDid sovrin.SovrinDid) CreateAgentMsg {
	return CreateAgentMsg{
		SignBytes:  "",
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createAgentDoc,
	}
}

func NewUpdateAgentMsg(txHash string, senderDid string, updateAgentDoc UpdateAgentDoc, projectDid sovrin.SovrinDid) UpdateAgentMsg {
	return UpdateAgentMsg{
		SignBytes:  "",
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       updateAgentDoc,
	}
}

func NewCreateClaimMsg(txHash string, senderDid string, createClaimDoc CreateClaimDoc, projectDid sovrin.SovrinDid) CreateClaimMsg {
	return CreateClaimMsg{
		SignBytes:  "",
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createClaimDoc,
	}
}

func NewCreateEvaluationMsg(txHash string, senderDid string, createEvaluationDoc CreateEvaluationDoc, projectDid sovrin.SovrinDid) CreateEvaluationMsg {
	return CreateEvaluationMsg{
		SignBytes:  "",
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createEvaluationDoc,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if len(value) == 0 {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}

func NewWithDrawFundsMsg(senderDid ixo.Did, data WithdrawFundsDoc) WithdrawFundsMsg {
	return WithdrawFundsMsg{
		SignBytes: "",
		SenderDid: senderDid,
		Data:      data,
	}
}
