package project

import (
	"encoding/json"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

//DOC SETUP
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
		NullStatus:     []ProjectStatus{CreatedProject},
		CreatedProject: []ProjectStatus{PendingStatus},
		PendingStatus:  []ProjectStatus{CreatedProject, FundedStatus},
		FundedStatus:   []ProjectStatus{StartedStatus},
		StartedStatus:  []ProjectStatus{StoppedStatus},
		StoppedStatus:  []ProjectStatus{PaidoutStatus},
	}

}

//IsValidProgressionFrom encapsulates legal ProjectStatus prpgression
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
	ActionID            [32]byte `json:"actionID"`
	ProjectEthWallet    string   `json:"projectEthWallet"`
	RecipientEthAddress string   `json:"recipientEthAddress"`
	Amount              int64    `json:"amount"`
}

//UpdateProjectStatusDoc defined
type UpdateProjectStatusDoc struct {
	Status          ProjectStatus `json:"status"`
	EthFundingTxnID string        `json:"ethFundingTxnID"`
}

func (ps UpdateProjectStatusDoc) GetEthFundingTxnID() string {
	return ps.EthFundingTxnID
}

//Define ProjectDoc
type ProjectDoc struct {
	NodeDid              string        `json:"nodeDid"`
	RequiredClaims       string        `json:"requiredClaims"`
	EvaluatorPayPerClaim string        `json:"evaluatorPayPerClaim"`
	ServiceEndpoint      string        `json:"serviceEndpoint"`
	CreatedOn            string        `json:"createdOn"`
	CreatedBy            string        `json:"createdBy"`
	Status               ProjectStatus `json:"status"`
}

//GETTERS

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

//DECODERS
type ProjectDocDecoder func(projectEntryBytes []byte) (StoredProjectDoc, error)

func GetProjectDocDecoder(cdc *wire.Codec) ProjectDocDecoder {
	return func(projectDocBytes []byte) (res StoredProjectDoc, err error) {
		if len(projectDocBytes) == 0 {
			return nil, sdk.ErrTxDecode("projectDocBytes are empty")
		}
		projectDoc := CreateProjectMsg{}
		err = cdc.UnmarshalBinary(projectDocBytes, &projectDoc)
		if err != nil {
			panic(err)
		}
		return &projectDoc, err
	}
}

// Define CreateAgent
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

// Define CreateAgent
type UpdateAgentDoc struct {
	Did    ixo.Did     `json:"did"`
	Status AgentStatus `json:"status"`
	Role   string      `json:"role"`
}

// Define CreateAgent
type CreateClaimDoc struct {
	ClaimID string `json:"claimID"`
}

type ClaimStatus string

const (
	PendingClaim  ClaimStatus = "0"
	ApprovedClaim ClaimStatus = "1"
	RejectedClaim ClaimStatus = "2"
)

// Define CreateAgent
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

//**************************************************************************************
// Message
type ProjectMsg interface {
	sdk.Msg
	IsNewDid() bool
	IsWithdrawal() bool
}

//CreateProjectMsg
type CreateProjectMsg struct {
	SignBytes  string     `json:"signBytes"`
	TxHash     string     `json:"txHash"`
	SenderDid  ixo.Did    `json:"senderDid"`
	ProjectDid ixo.Did    `json:"projectDid"`
	PubKey     string     `json:"pubKey"`
	Data       ProjectDoc `json:"data"`
}

var _ sdk.Msg = CreateProjectMsg{}

func (msg CreateProjectMsg) Type() string                            { return "project" }
func (msg CreateProjectMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateProjectMsg) ValidateBasic() sdk.Error {
	valid, err := CheckNotEmpty(msg.PubKey, "PubKey")
	if !valid {
		return err
	}
	valid, err = CheckNotEmpty(msg.ProjectDid, "ProjectDid")
	if !valid {
		return err
	}
	valid, err = CheckNotEmpty(msg.Data.NodeDid, "NodeDid")
	if !valid {
		return err
	}
	valid, err = CheckNotEmpty(msg.Data.RequiredClaims, "RequiredClaims")
	if !valid {
		return err
	}
	valid, err = CheckNotEmpty(msg.Data.CreatedBy, "CreatedBy")
	if !valid {
		return err
	}
	return nil
}
func (msg CreateProjectMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateProjectMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateProjectMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}
func (msg CreateProjectMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}
func (msg CreateProjectMsg) GetPubKey() string        { return msg.PubKey }
func (msg CreateProjectMsg) GetEvaluatorPay() int64   { return msg.Data.GetEvaluatorPay() }
func (msg CreateProjectMsg) GetStatus() ProjectStatus { return msg.Data.Status }
func (msg *CreateProjectMsg) SetStatus(status ProjectStatus) {
	msg.Data.Status = status
}
func (msg CreateProjectMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}
func (msg CreateProjectMsg) IsNewDid() bool     { return true }
func (msg CreateProjectMsg) IsWithdrawal() bool { return false }

var _ StoredProjectDoc = (*CreateProjectMsg)(nil)

//UpdateProjectStatusMsg
type UpdateProjectStatusMsg struct {
	SignBytes  string                 `json:"signBytes"`
	TxHash     string                 `json:"txHash"`
	SenderDid  ixo.Did                `json:"senderDid"`
	ProjectDid ixo.Did                `json:"projectDid"`
	Data       UpdateProjectStatusDoc `json:"data"`
}

func (msg UpdateProjectStatusMsg) Type() string                            { return "project" }
func (msg UpdateProjectStatusMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg UpdateProjectStatusMsg) ValidateBasic() sdk.Error                { return nil }
func (msg UpdateProjectStatusMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}
func (msg UpdateProjectStatusMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}
func (ups UpdateProjectStatusMsg) GetProjectDid() ixo.Did {
	return ups.ProjectDid
}
func (ups UpdateProjectStatusMsg) GetStatus() ProjectStatus {
	return ups.Data.Status
}
func (msg UpdateProjectStatusMsg) IsNewDid() bool     { return false }
func (msg UpdateProjectStatusMsg) IsWithdrawal() bool { return false }
func (msg UpdateProjectStatusMsg) GetEthFundingTxnID() string {
	return msg.Data.EthFundingTxnID
}

//CreateAgentMsg
type CreateAgentMsg struct {
	SignBytes  string         `json:"signBytes"`
	TxHash     string         `json:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid"`
	Data       CreateAgentDoc `json:"data"`
}

func (msg CreateAgentMsg) IsNewDid() bool                          { return false }
func (msg CreateAgentMsg) IsWithdrawal() bool                      { return false }
func (msg CreateAgentMsg) Type() string                            { return "project" }
func (msg CreateAgentMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateAgentMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg CreateAgentMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateAgentMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateAgentMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}
func (msg CreateAgentMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg CreateAgentMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

var _ sdk.Msg = CreateAgentMsg{}

//UpdateAgentMsg
type UpdateAgentMsg struct {
	SignBytes  string         `json:"signBytes"`
	TxHash     string         `json:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid"`
	Data       UpdateAgentDoc `json:"data"`
}

func (msg UpdateAgentMsg) IsNewDid() bool                          { return false }
func (msg UpdateAgentMsg) IsWithdrawal() bool                      { return false }
func (msg UpdateAgentMsg) Type() string                            { return "project" }
func (msg UpdateAgentMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg UpdateAgentMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg UpdateAgentMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg UpdateAgentMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg UpdateAgentMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}
func (msg UpdateAgentMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}
func (msg UpdateAgentMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

var _ sdk.Msg = UpdateAgentMsg{}

//CreateClaimMsg
type CreateClaimMsg struct {
	SignBytes  string         `json:"signBytes"`
	TxHash     string         `json:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid"`
	Data       CreateClaimDoc `json:"data"`
}

func (msg CreateClaimMsg) IsNewDid() bool                          { return false }
func (msg CreateClaimMsg) IsWithdrawal() bool                      { return false }
func (msg CreateClaimMsg) Type() string                            { return "project" }
func (msg CreateClaimMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateClaimMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg CreateClaimMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateClaimMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateClaimMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}
func (msg CreateClaimMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}
func (msg CreateClaimMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

var _ sdk.Msg = CreateClaimMsg{}

//CreateEvaluationMsg
type CreateEvaluationMsg struct {
	SignBytes  string              `json:"signBytes"`
	TxHash     string              `json:"txHash"`
	SenderDid  ixo.Did             `json:"senderDid"`
	ProjectDid ixo.Did             `json:"projectDid"`
	Data       CreateEvaluationDoc `json:"data"`
}

func (msg CreateEvaluationMsg) IsNewDid() bool                          { return false }
func (msg CreateEvaluationMsg) IsWithdrawal() bool                      { return false }
func (msg CreateEvaluationMsg) Type() string                            { return "project" }
func (msg CreateEvaluationMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateEvaluationMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg CreateEvaluationMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateEvaluationMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateEvaluationMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}
func (msg CreateEvaluationMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}
func (msg CreateEvaluationMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

var _ sdk.Msg = CreateEvaluationMsg{}

//WithdrawFundsMsg
type WithdrawFundsMsg struct {
	SignBytes string           `json:"signBytes"`
	SenderDid ixo.Did          `json:"senderDid"`
	Data      WithdrawFundsDoc `json:"data"`
}

func (msg WithdrawFundsMsg) IsNewDid() bool                          { return false }
func (msg WithdrawFundsMsg) IsWithdrawal() bool                      { return true }
func (msg WithdrawFundsMsg) Type() string                            { return "project" }
func (msg WithdrawFundsMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg WithdrawFundsMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg WithdrawFundsMsg) GetSenderDid() ixo.Did                 { return msg.SenderDid }
func (msg WithdrawFundsMsg) GetWithdrawFundsDoc() WithdrawFundsDoc { return msg.Data }
func (msg WithdrawFundsMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}
func (msg WithdrawFundsMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}
func (msg WithdrawFundsMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

var _ sdk.Msg = WithdrawFundsMsg{}
