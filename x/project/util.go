package project

import (
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func NewCreateProjectMsg(projectDoc ProjectDoc, projectDid ixo.SovrinDid) CreateProjectMsg {
	projectDoc.ProjectDid = projectDid.Did
	projectDoc.PubKey = projectDid.VerifyKey
	return CreateProjectMsg{
		projectDoc,
	}
}

func NewCreateAgentMsg(createAgentDoc CreateAgentDoc) CreateAgentMsg {
	return CreateAgentMsg{
		createAgentDoc,
	}
}

func NewUpdateAgentMsg(updateAgentDoc UpdateAgentDoc) UpdateAgentMsg {
	return UpdateAgentMsg{
		updateAgentDoc,
	}
}

func NewCreateClaimMsg(createClaimDoc CreateClaimDoc) CreateClaimMsg {
	return CreateClaimMsg{
		createClaimDoc,
	}
}

func NewCreateEvaluationMsg(createEvaluationDoc CreateEvaluationDoc) CreateEvaluationMsg {
	return CreateEvaluationMsg{
		createEvaluationDoc,
	}
}
