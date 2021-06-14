package types

import (
	"encoding/json"
	"fmt"

	didexported "github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

type (
	InternalAccountID string
	//AccountMap                 map[InternalAccountID]sdk.AccAddress
	//GenesisAccountMap          map[string]sdk.AccAddress
	ProjectStatus              string
	ProjectStatusTransitionMap map[ProjectStatus][]ProjectStatus
	ProjectDataMap             map[string]json.RawMessage
	ProjectFeesMap             struct {
		Context string `json:"@context" yaml:"@context"`
		Items   []ProjectFeesMapItem
	}
	ProjectFeesMapItem struct {
		Type              FeeType `json:"@type" yaml:"@type"`
		PaymentTemplateId string  `json:"id" yaml:"id"`
	}
	FeeType string
)

const (
	FeeForService      FeeType = "FeeForService"
	OracleFee          FeeType = "OracleFee"
	Subscription       FeeType = "Subscription"
	RentalFee          FeeType = "RentalFee"
	OutcomePayment     FeeType = "OutcomePayment"
	InterestRepayment  FeeType = "InterestRepayment"
	LoanRepayment      FeeType = "LoanRepayment"
	IncomeDistribution FeeType = "IncomeDistribution"
	DisputeSettlement  FeeType = "DisputeSettlement"
)

func (pfm ProjectFeesMap) GetPayTemplateId(feeType FeeType) (string, error) {
	for _, v := range pfm.Items {
		if v.Type == feeType {
			return v.PaymentTemplateId, nil
		}
	}
	return "", fmt.Errorf("fee '%s' not found in fees map", feeType)
}

func (id InternalAccountID) ToAddressKey(projectDid didexported.Did) string {
	return projectDid + "/" + string(id)
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

func ProjectStatusFromString(s string) ProjectStatus {
	return ProjectStatus(s)
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

type AgentStatus = string

const (
	PendingAgent  AgentStatus = "0"
	ApprovedAgent AgentStatus = "1"
	RevokedAgent  AgentStatus = "2"
)

type ClaimStatus string

const (
	PendingClaim  ClaimStatus = "0"
	ApprovedClaim ClaimStatus = "1"
	RejectedClaim ClaimStatus = "2"
)

func NewClaim(id string, templateId string, claimerDid didexported.Did) Claim {
	return Claim{
		Id:         id,
		TemplateId: templateId,
		ClaimerDid: claimerDid,
		Status:     string(PendingClaim),
	}
}

func AppendClaims(list Claims, claim Claim) Claims {
	appended := append(list.ClaimsList, claim)
	return Claims{ClaimsList: appended}
}

func AppendWithdrawalInfoDocs(list WithdrawalInfoDocs, doc WithdrawalInfoDoc) WithdrawalInfoDocs {
	appended := append(list.DocsList, doc)
	return WithdrawalInfoDocs{DocsList: appended}
}
