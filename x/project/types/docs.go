package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	didexported "github.com/ixofoundation/ixo-blockchain/lib/legacydid"
)

func NewProjectDoc(txHash string, projectDid, senderDid didexported.Did,
	pubKey string, status ProjectStatus, data json.RawMessage) ProjectDoc {
	return ProjectDoc{
		TxHash:     txHash,
		ProjectDid: projectDid,
		SenderDid:  senderDid,
		PubKey:     pubKey,
		Status:     string(status),
		Data:       data,
	}
}

func (pd ProjectDoc) GetProjectData() (dataMap ProjectDataMap) {
	err := json.Unmarshal(pd.Data, &dataMap)
	if err != nil {
		panic(err)
	}
	return dataMap
}

func (pd ProjectDoc) GetProjectFeesMap() (feesMap ProjectFeesMap) {
	feesMapRaw := pd.GetProjectData()["fees"]
	err := json.Unmarshal(feesMapRaw, &feesMap)
	if err != nil {
		panic(err)
	}
	return feesMap
}

func NewUpdateProjectStatusDoc(status ProjectStatus, ethFundingTxnID string) UpdateProjectStatusDoc {
	return UpdateProjectStatusDoc{
		Status:          string(status),
		EthFundingTxnId: ethFundingTxnID,
	}
}

func NewCreateAgentDoc(agentDid didexported.Did, role string) CreateAgentDoc {
	return CreateAgentDoc{
		AgentDid: agentDid,
		Role:     role,
	}
}

func NewUpdateAgentDoc(did didexported.Did, status AgentStatus, role string) UpdateAgentDoc {
	return UpdateAgentDoc{
		Did:    did,
		Status: status,
		Role:   role,
	}
}

func NewCreateClaimDoc(claimId string, claimTemplateID string) CreateClaimDoc {
	return CreateClaimDoc{
		ClaimId:         claimId,
		ClaimTemplateId: claimTemplateID,
	}
}

func NewCreateEvaluationDoc(claimId string, status ClaimStatus) CreateEvaluationDoc {
	return CreateEvaluationDoc{
		ClaimId: claimId,
		Status:  string(status),
	}
}

func NewWithdrawalInfoDoc(projectDid, recipientDid didexported.Did, amount sdk.Coin) WithdrawalInfoDoc {
	return WithdrawalInfoDoc{
		ProjectDid:   projectDid,
		RecipientDid: recipientDid,
		Amount:       amount,
	}
}

func NewWithdrawFundsDoc(projectDid, recipientDid didexported.Did, amount sdk.Int, isRefund bool) WithdrawFundsDoc {
	return WithdrawFundsDoc{
		ProjectDid:   projectDid,
		RecipientDid: recipientDid,
		Amount:       amount,
		IsRefund:     isRefund,
	}
}
