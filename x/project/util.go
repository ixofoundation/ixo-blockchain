package project

import (
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
)

func NewCreateProjectMsg(projectDoc ProjectDoc, projectDid sovrin.SovrinDid) CreateProjectMsg {
	projectDoc.ProjectDid = projectDid.Did
	projectDoc.PubKey = projectDid.VerifyKey
	return CreateProjectMsg{
		projectDoc,
	}
}

func NewCreateAgentMsg(createAgentDoc CreateAgentDoc, projectDid sovrin.SovrinDid) CreateAgentMsg {
	createAgentDoc.ProjectDid = projectDid.Did
	return CreateAgentMsg{
		createAgentDoc,
	}
}

func NewUpdateAgentMsg(updateAgentDoc UpdateAgentDoc, projectDid sovrin.SovrinDid) UpdateAgentMsg {
	updateAgentDoc.ProjectDid = projectDid.Did
	return UpdateAgentMsg{
		updateAgentDoc,
	}
}

func NewCreateClaimMsg(createClaimDoc CreateClaimDoc, projectDid sovrin.SovrinDid) CreateClaimMsg {
	createClaimDoc.ProjectDid = projectDid.Did
	return CreateClaimMsg{
		createClaimDoc,
	}
}

func NewCreateEvaluationMsg(createEvaluationDoc CreateEvaluationDoc, projectDid sovrin.SovrinDid) CreateEvaluationMsg {
	createEvaluationDoc.ProjectDid = projectDid.Did
	return CreateEvaluationMsg{
		createEvaluationDoc,
	}
}

func NewFundProjectMsg(fundProjectDoc FundProjectDoc) FundProjectMsg {
	// This did is initialized when the node is initialized
	fundProjectDoc.Signer = "ETH_PEG"
	return FundProjectMsg{
		fundProjectDoc,
	}
}
