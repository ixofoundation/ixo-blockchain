package types

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
)

var (
	_ StoredProjectDoc = (*ProjectDoc)(nil)
)

type (
	InternalAccountID          string
	AccountMap                 map[InternalAccountID]sdk.AccAddress
	GenesisAccountMap          map[string]sdk.AccAddress
	ProjectStatus              string
	ProjectStatusTransitionMap map[ProjectStatus][]ProjectStatus
	ProjectDataMap             map[string]json.RawMessage
)

func (id InternalAccountID) ToAddressKey(projectDid did.Did) string {
	return projectDid + "/" + string(id)
}

type StoredProjectDoc interface {
	GetClaimerPay() sdk.Coins
	GetEvaluatorPay() sdk.Coins
	GetProjectDid() did.Did
	GetSenderDid() did.Did
	GetPubKey() string
	GetStatus() ProjectStatus
	SetStatus(status ProjectStatus)
}

const (
	NullStatus     ProjectStatus = ""
	CreatedProject ProjectStatus = "CREATED"
	PendingStatus  ProjectStatus = "PENDING"
	FundedStatus   ProjectStatus = "FUNDED"
	StartedStatus  ProjectStatus = "STARTED"
	StoppedStatus  ProjectStatus = "STOPPED"
	PaidoutStatus  ProjectStatus = "PAIDOUT"
)

var StateTransitions = initStateTransitions()

func initStateTransitions() ProjectStatusTransitionMap {
	return ProjectStatusTransitionMap{
		NullStatus:     {CreatedProject},
		CreatedProject: {PendingStatus},
		PendingStatus:  {CreatedProject, FundedStatus},
		FundedStatus:   {StartedStatus},
		StartedStatus:  {StoppedStatus},
		StoppedStatus:  {PaidoutStatus},
	}
}

func (next ProjectStatus) IsValidProgressionFrom(prev ProjectStatus) bool {
	validStatuses := StateTransitions[prev]
	for _, v := range validStatuses {
		if v == next {
			return true
		}
	}
	return false
}

type WithdrawalInfo struct {
	ActionID     string   `json:"actionID" yaml:"actionID"`
	ProjectDid   did.Did  `json:"projectDid" yaml:"projectDid"`
	RecipientDid did.Did  `json:"recipientDid" yaml:"recipientDid"`
	Amount       sdk.Coin `json:"amount" yaml:"amount"`
}

type UpdateProjectStatusDoc struct {
	Status          ProjectStatus `json:"status" yaml:"status"`
	EthFundingTxnID string        `json:"ethFundingTxnID" yaml:"ethFundingTxnID"`
}

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

func (pd ProjectDoc) GetProjectDid() did.Did          { return pd.ProjectDid }
func (pd ProjectDoc) GetSenderDid() did.Did           { return pd.SenderDid }
func (pd ProjectDoc) GetPubKey() string               { return pd.PubKey }
func (pd ProjectDoc) GetStatus() ProjectStatus        { return pd.Status }
func (pd *ProjectDoc) SetStatus(status ProjectStatus) { pd.Status = status }
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

func (pd ProjectDoc) GetEvaluatorPay() sdk.Coins {
	return pd.getPay("evaluatorPayPerClaim")
}

func (pd ProjectDoc) GetClaimVerifiedPay() sdk.Coins {
	return pd.getPay("claimerPayPerVerifiedClaim")
}

type CreateAgentDoc struct {
	AgentDid did.Did `json:"did" yaml:"did"`
	Role     string  `json:"role" yaml:"role"`
}

type AgentStatus = string

const (
	PendingAgent  AgentStatus = "0"
	ApprovedAgent AgentStatus = "1"
	RevokedAgent  AgentStatus = "2"
)

type UpdateAgentDoc struct {
	Did    did.Did     `json:"did" yaml:"did"`
	Status AgentStatus `json:"status" yaml:"status"`
	Role   string      `json:"role" yaml:"role"`
}

type CreateClaimDoc struct {
	ClaimID string `json:"claimID" yaml:"claimID"`
}

type ClaimStatus string

const (
	PendingClaim  ClaimStatus = "0"
	ApprovedClaim ClaimStatus = "1"
	RejectedClaim ClaimStatus = "2"
)

type CreateEvaluationDoc struct {
	ClaimID string      `json:"claimID" yaml:"claimID"`
	Status  ClaimStatus `json:"status" yaml:"status"`
}

type WithdrawFundsDoc struct {
	ProjectDid   did.Did `json:"projectDid" yaml:"projectDid"`
	RecipientDid did.Did `json:"recipientDid" yaml:"recipientDid"`
	Amount       sdk.Int `json:"amount" yaml:"amount"`
	IsRefund     bool    `json:"isRefund" yaml:"isRefund"`
}
