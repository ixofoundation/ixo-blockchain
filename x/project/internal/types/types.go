package types

import (
	"github.com/tendermint/tendermint/crypto"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

type Config struct {
	AccountMapPrefix  string
	WithdrawalsPrefix string
}

type AccountMap map[string]sdk.AccAddress

type StoredProjectDoc interface {
	GetEvaluatorPay() int64
	GetProjectDid() ixo.Did
	GetPubKey() string
	GetStatus() ProjectStatus
	SetStatus(status ProjectStatus)
}

type ProjectStatus string

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

func initStateTransitions() map[ProjectStatus][]ProjectStatus {
	return map[ProjectStatus][]ProjectStatus{
		NullStatus:     {CreatedProject},
		CreatedProject: {PendingStatus},
		PendingStatus:  {CreatedProject, FundedStatus},
		FundedStatus:   {StartedStatus},
		StartedStatus:  {StoppedStatus},
		StoppedStatus:  {PaidoutStatus},
	}

}

func (nextProjectStatus ProjectStatus) IsValidProgressionFrom(previousProjectStatus ProjectStatus) bool {
	validStatuses := StateTransitions[previousProjectStatus]
	for _, v := range validStatuses {
		if v == nextProjectStatus {
			return true
		}
	}

	return false
}

type WithdrawalInfo struct {
	ActionID            string `json:"actionID" yaml:"actionID"`
	ProjectEthWallet    string `json:"projectEthWallet" yaml:"projectEthWallet"`
	RecipientEthAddress string `json:"recipientEthAddress" yaml:"recipientEthAddress"`
	Amount              int64  `json:"amount" yaml:"amount"`
}

type UpdateProjectStatusDoc struct {
	Status          ProjectStatus `json:"status" yaml:"status"`
	EthFundingTxnID string        `json:"ethFundingTxnID" yaml:"ethFundingTxnID"`
}

type ProjectDoc struct {
	NodeDid              string        `json:"nodeDid" yaml:"nodeDid"`
	RequiredClaims       string        `json:"requiredClaims" yaml:"requiredClaims"`
	EvaluatorPayPerClaim string        `json:"evaluatorPayPerClaim" yaml:"evaluatorPayPerClaim"`
	ServiceEndpoint      string        `json:"serviceEndpoint" yaml:"serviceEndpoint"`
	CreatedOn            string        `json:"createdOn" yaml:"createdOn"`
	CreatedBy            string        `json:"createdBy" yaml:"createdBy"`
	Status               ProjectStatus `json:"status" yaml:"status"`
}

func (pd ProjectDoc) GetEvaluatorPay() int64 {
	if pd.EvaluatorPayPerClaim == "" {
		return 0
	} else {
		i, err := strconv.ParseInt(pd.EvaluatorPayPerClaim, 10, 64)
		if err != nil {
			panic(err)
		}

		return i
	}
}

type ProjectDocDecoder func(projectEntryBytes []byte) (StoredProjectDoc, error)

type CreateAgentDoc struct {
	AgentDid ixo.Did `json:"did" yaml:"did"`
	Role     string  `json:"role" yaml:"role"`
}

type AgentStatus string

const (
	PendingAgent  AgentStatus = "0"
	ApprovedAgent AgentStatus = "1"
	RevokedAgent  AgentStatus = "2"
)

type UpdateAgentDoc struct {
	Did    ixo.Did     `json:"did" yaml:"did"`
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
	ProjectDid ixo.Did `json:"projectDid" yaml:"projectDid"`
	EthWallet  string  `json:"ethWallet" yaml:"ethWallet"`
	Amount     string  `json:"amount" yaml:"amount"`
	IsRefund   bool    `json:"isRefund" yaml:"isRefund"`
}

func (wd WithdrawFundsDoc) GetProjectDid() ixo.Did { return wd.ProjectDid }
func (wd WithdrawFundsDoc) GetEthWallet() string   { return wd.EthWallet }
func (wd WithdrawFundsDoc) GetIsRefund() bool      { return wd.IsRefund }

type ProjectMsg interface {
	sdk.Msg
	IsNewDid() bool
	IsWithdrawal() bool
}

func StringToAddr(str string) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(str)))
}
