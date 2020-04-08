package types

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
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
	ActionID            string `json:"actionID"`
	ProjectEthWallet    string `json:"projectEthWallet"`
	RecipientEthAddress string `json:"recipientEthAddress"`
	Amount              int64  `json:"amount"`
}

type UpdateProjectStatusDoc struct {
	Status          ProjectStatus `json:"status"`
	EthFundingTxnID string        `json:"ethFundingTxnID"`
}

func (ps UpdateProjectStatusDoc) GetEthFundingTxnID() string {
	return ps.EthFundingTxnID
}

type ProjectDoc struct {
	NodeDid              string        `json:"nodeDid"`
	RequiredClaims       string        `json:"requiredClaims"`
	EvaluatorPayPerClaim string        `json:"evaluatorPayPerClaim"`
	ServiceEndpoint      string        `json:"serviceEndpoint"`
	CreatedOn            string        `json:"createdOn"`
	CreatedBy            string        `json:"createdBy"`
	Status               ProjectStatus `json:"status"`
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

func GetProjectDocDecoder(cdc *codec.Codec) ProjectDocDecoder {
	return func(projectDocBytes []byte) (res StoredProjectDoc, err error) {
		if len(projectDocBytes) == 0 {
			return nil, sdk.ErrTxDecode("projectDocBytes are empty")
		}
		projectDoc := MsgCreateProject{}
		err = cdc.UnmarshalBinaryLengthPrefixed(projectDocBytes, &projectDoc)
		if err != nil {
			panic(err)
		}

		return &projectDoc, err
	}
}

type CreateAgentDoc struct {
	AgentDid ixo.Did `json:"did"`
	Role     string  `json:"role"`
}

type AgentStatus string

const (
	PendingAgent  AgentStatus = "0"
	ApprovedAgent AgentStatus = "1"
	RevokedAgent  AgentStatus = "2"
)

type UpdateAgentDoc struct {
	Did    ixo.Did     `json:"did"`
	Status AgentStatus `json:"status"`
	Role   string      `json:"role"`
}

type CreateClaimDoc struct {
	ClaimID string `json:"claimID"`
}

type ClaimStatus string

const (
	PendingClaim  ClaimStatus = "0"
	ApprovedClaim ClaimStatus = "1"
	RejectedClaim ClaimStatus = "2"
)

type CreateEvaluationDoc struct {
	ClaimID string      `json:"claimID"`
	Status  ClaimStatus `json:"status"`
}

type WithdrawFundsDoc struct {
	ProjectDid ixo.Did `json:"projectDid"`
	EthWallet  string  `json:"ethWallet"`
	Amount     string  `json:"amount"`
	IsRefund   bool    `json:"isRefund"`
}

func (wd WithdrawFundsDoc) GetProjectDid() ixo.Did { return wd.ProjectDid }
func (wd WithdrawFundsDoc) GetEthWallet() string   { return wd.EthWallet }
func (wd WithdrawFundsDoc) GetIsRefund() bool      { return wd.IsRefund }

type ProjectMsg interface {
	sdk.Msg
	IsNewDid() bool
	IsWithdrawal() bool
}
