package project

import (
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
)

func NewCreateProjectMsg(projectDoc ProjectDoc, projectDid sovrin.SovrinDid) CreateProjectMsg {
	return CreateProjectMsg{
		Data:       projectDoc,
		ProjectDid: projectDid.Did,
		PubKey:     projectDid.VerifyKey,
	}
}

func NewCreateAgentMsg(txHash string, senderDid string, createAgentDoc CreateAgentDoc, projectDid sovrin.SovrinDid) CreateAgentMsg {
	return CreateAgentMsg{
		Data:       createAgentDoc,
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
	}
}

func NewUpdateAgentMsg(txHash string, senderDid string, updateAgentDoc UpdateAgentDoc, projectDid sovrin.SovrinDid) UpdateAgentMsg {
	return UpdateAgentMsg{
		Data:       updateAgentDoc,
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
	}
}

func NewCreateClaimMsg(txHash string, senderDid string, createClaimDoc CreateClaimDoc, projectDid sovrin.SovrinDid) CreateClaimMsg {
	return CreateClaimMsg{
		Data:       createClaimDoc,
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
	}
}

func NewCreateEvaluationMsg(txHash string, senderDid string, createEvaluationDoc CreateEvaluationDoc, projectDid sovrin.SovrinDid) CreateEvaluationMsg {
	return CreateEvaluationMsg{
		Data:       createEvaluationDoc,
		ProjectDid: projectDid.Did,
		TxHash:     txHash,
		SenderDid:  senderDid,
	}
}

func NewFundProjectMsg(fundProjectDoc FundProjectDoc) FundProjectMsg {
	// This did is initialized when the node is initialized
	fundProjectDoc.Signer = "ETH_PEG"
	return FundProjectMsg{
		fundProjectDoc,
	}
}
