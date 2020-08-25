package types

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
)

type ProjectDoc struct {
	TxHash     string          `json:"txHash" yaml:"txHash"`
	ProjectDid did.Did         `json:"projectDid" yaml:"projectDid"`
	SenderDid  did.Did         `json:"senderDid" yaml:"senderDid"`
	PubKey     string          `json:"pubKey" yaml:"pubKey"`
	Status     ProjectStatus   `json:"status" yaml:"status"`
	Data       json.RawMessage `json:"data" yaml:"data"`
}

func NewProjectDoc(txHash string, projectDid, senderDid did.Did,
	pubKey string, status ProjectStatus, data json.RawMessage) ProjectDoc {
	return ProjectDoc{
		TxHash:     txHash,
		ProjectDid: projectDid,
		SenderDid:  senderDid,
		PubKey:     pubKey,
		Status:     status,
		Data:       data,
	}
}

func (pd ProjectDoc) GetProjectData() ProjectDataMap {
	var dataMap ProjectDataMap
	err := json.Unmarshal(pd.Data, &dataMap)
	if err != nil {
		panic(err)
	}
	return dataMap
}

func (pd ProjectDoc) getPay(key string) sdk.Coins {
	payBz, found := pd.GetProjectData()[key]
	if !found {
		panic(fmt.Sprintf("%s not found", key))
	}
	payCoins, err := sdk.ParseCoins(withoutQuotes(string(payBz)))
	if err != nil {
		panic(err)
	}
	return payCoins
}

func (pd ProjectDoc) GetClaimerPay() sdk.Coins {
	return pd.getPay("claimerPayPerClaim")
}

func (pd ProjectDoc) GetClaimApprovedPay() sdk.Coins {
	return pd.getPay("claimerPayPerApprovedClaim")
}

func (pd ProjectDoc) GetEvaluatorPay() sdk.Coins {
	return pd.getPay("evaluatorPayPerClaim")
}

type UpdateProjectStatusDoc struct {
	Status          ProjectStatus `json:"status" yaml:"status"`
	EthFundingTxnID string        `json:"ethFundingTxnID" yaml:"ethFundingTxnID"`
}

func NewUpdateProjectStatusDoc(status ProjectStatus, ethFundingTxnID string) UpdateProjectStatusDoc {
	return UpdateProjectStatusDoc{
		Status:          status,
		EthFundingTxnID: ethFundingTxnID,
	}
}

type CreateAgentDoc struct {
	AgentDid did.Did `json:"did" yaml:"did"`
	Role     string  `json:"role" yaml:"role"`
}

func NewCreateAgentDoc(agentDid did.Did, role string) CreateAgentDoc {
	return CreateAgentDoc{
		AgentDid: agentDid,
		Role:     role,
	}
}

type UpdateAgentDoc struct {
	Did    did.Did     `json:"did" yaml:"did"`
	Status AgentStatus `json:"status" yaml:"status"`
	Role   string      `json:"role" yaml:"role"`
}

func NewUpdateAgentDoc(did did.Did, status AgentStatus, role string) UpdateAgentDoc {
	return UpdateAgentDoc{
		Did:    did,
		Status: status,
		Role:   role,
	}
}

type CreateClaimDoc struct {
	ClaimID string `json:"claimID" yaml:"claimID"`
}

func NewCreateClaimDoc(claimId string) CreateClaimDoc {
	return CreateClaimDoc{
		ClaimID: claimId,
	}
}

type CreateEvaluationDoc struct {
	ClaimID string      `json:"claimID" yaml:"claimID"`
	Status  ClaimStatus `json:"status" yaml:"status"`
}

func NewCreateEvaluationDoc(claimId string, status ClaimStatus) CreateEvaluationDoc {
	return CreateEvaluationDoc{
		ClaimID: claimId,
		Status:  status,
	}
}

type WithdrawalInfoDoc struct {
	ProjectDid   did.Did  `json:"projectDid" yaml:"projectDid"`
	RecipientDid did.Did  `json:"recipientDid" yaml:"recipientDid"`
	Amount       sdk.Coin `json:"amount" yaml:"amount"`
}

func NewWithdrawalInfoDoc(projectDid, recipientDid did.Did, amount sdk.Coin) WithdrawalInfoDoc {
	return WithdrawalInfoDoc{
		ProjectDid:   projectDid,
		RecipientDid: recipientDid,
		Amount:       amount,
	}
}

type WithdrawFundsDoc struct {
	ProjectDid   did.Did `json:"projectDid" yaml:"projectDid"`
	RecipientDid did.Did `json:"recipientDid" yaml:"recipientDid"`
	Amount       sdk.Int `json:"amount" yaml:"amount"`
	IsRefund     bool    `json:"isRefund" yaml:"isRefund"`
}

func NewWithdrawFundsDoc(projectDid, recipientDid did.Did, amount sdk.Int, isRefund bool) WithdrawFundsDoc {
	return WithdrawFundsDoc{
		ProjectDid:   projectDid,
		RecipientDid: recipientDid,
		Amount:       amount,
		IsRefund:     isRefund,
	}
}
