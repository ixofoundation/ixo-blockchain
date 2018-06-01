package project

import (
	"encoding/json"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const COIN_DENOM = "ixo"

//DOC SETUP

//Define ProjectDoc
type ProjectDoc struct {
	Title            string   `json:"title"`
	OwnerName        string   `json:"ownerName"`
	OwnerEmail       string   `json:"ownerEmail"`
	ShortDescription string   `json:"shortDescription"`
	LongDescription  string   `json:"longDescription"`
	ImpactAction     string   `json:"impactAction"`
	CreatedOn        string   `json:"createdOn"`
	CreatedBy        string   `json:"createdBy"`
	ProjectLocation  string   `json:"projectLocation"`
	Sdgs             []string `json:"sdgs"`
	Claims           struct {
		Required          int `json:"required"`
		CurrentSuccessful int `json:"currentSuccessful"`
		CurrentRejected   int `json:"currentRejected"`
	} `json:"claims"`
	Templates struct {
		Claim string `json:"claim"`
	} `json:"templates"`
	Agents struct {
		Evaluators              int `json:"evaluators"`
		EvaluatorsPending       int `json:"evaluatorsPending"`
		ServiceProviders        int `json:"serviceProviders"`
		ServiceProvidersPending int `json:"serviceProvidersPending"`
		Investors               int `json:"investors"`
	} `json:"agents"`
	EvaluatorPayPerClaim string `json:"evaluatorPayPerClaim"`
	SocialMedia          struct {
		FacebookLink  string `json:"facebookLink"`
		InstagramLink string `json:"instagramLink"`
		TwitterLink   string `json:"twitterLink"`
		WebLink       string `json:"webLink"`
	} `json:"socialMedia"`
	Ixo struct {
		TotalStaked int `json:"totalStaked"`
		TotalUsed   int `json:"totalUsed"`
	} `json:"ixo"`
	ServiceEndpoint string `json:"serviceEndpoint"`
	ImageLink       string `json:"imageLink"`
	Founder         struct {
		Name             string `json:"name"`
		Email            string `json:"email"`
		CountryOfOrigin  string `json:"countryOfOrigin"`
		ShortDescription string `json:"shortDescription"`
		WebsiteURL       string `json:"websiteURL"`
		LogoLink         string `json:"logoLink"`
	} `json:"founder"`
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
type ProjectDocDecoder func(projectEntryBytes []byte) (ixo.StoredProjectDoc, error)

func GetProjectDocDecoder(cdc *wire.Codec) ProjectDocDecoder {
	return func(projectDocBytes []byte) (res ixo.StoredProjectDoc, err error) {
		if len(projectDocBytes) == 0 {
			return nil, sdk.ErrTxDecode("projectDocBytes are empty")
		}
		projectDoc := StoredProjectDoc{}
		err = cdc.UnmarshalBinary(projectDocBytes, &projectDoc)
		if err != nil {
			panic(err)
		}
		return projectDoc, err
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

type FundProjectDoc struct {
	Signer     ixo.Did `json:"signer"`
	EthTxHash  string  `json:"ethTxHash"`
	ProjectDid ixo.Did `json:"projectDid"`
	Amount     string  `json:"amount"`
}

func (fd FundProjectDoc) GetSigner() ixo.Did { return fd.Signer }
func (fd FundProjectDoc) GetAmount() int64 {
	i, err := strconv.ParseInt(fd.Amount, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

type WithdrawFundsDoc struct {
	ProjectDid ixo.Did `json:"projectDid"`
	EthWallet  string  `json:"ethWallet"`
	Amount     string  `json:"amount"`
}

func (wd WithdrawFundsDoc) GetProjectDid() ixo.Did { return wd.ProjectDid }

//**************************************************************************************
// Message

type ProjectMsg interface {
	IsNewDid() bool
	IsPegMsg() bool
}

//CreateProjectMsg
type CreateProjectMsg struct {
	Data       ProjectDoc `json:"data"`
	TxHash     string     `json:"txHash"`
	SenderDid  ixo.Did    `json:"senderDid"`
	ProjectDid ixo.Did    `json:"projectDid"`
	PubKey     string     `json:"pubKey"`
}

var _ sdk.Msg = CreateProjectMsg{}

func (msg CreateProjectMsg) Type() string                            { return "project" }
func (msg CreateProjectMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateProjectMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg CreateProjectMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateProjectMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateProjectMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.GetProjectDid())}
}
func (msg CreateProjectMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg CreateProjectMsg) IsNewDid() bool { return true }
func (msg CreateProjectMsg) IsPegMsg() bool { return false }
func (msg CreateProjectMsg) String() string {
	return fmt.Sprintf("CreateProjectMsg{Did: %v, projectDoc: %v}", string(msg.GetProjectDid()), msg.Data)
}
func (msg CreateProjectMsg) GetPubKey() string      { return msg.PubKey }
func (msg CreateProjectMsg) GetEvaluatorPay() int64 { return msg.Data.GetEvaluatorPay() }

type StoredProjectDoc = CreateProjectMsg

var _ ixo.StoredProjectDoc = (*StoredProjectDoc)(nil)

//CreateAgentMsg
type CreateAgentMsg struct {
	Data       CreateAgentDoc `json:"data"`
	TxHash     string         `json:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid"`
}

var _ sdk.Msg = CreateAgentMsg{}

func (msg CreateAgentMsg) Type() string                            { return "project" }
func (msg CreateAgentMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateAgentMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg CreateAgentMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateAgentMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateAgentMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.GetProjectDid())}
}
func (msg CreateAgentMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg CreateAgentMsg) IsNewDid() bool { return false }
func (msg CreateAgentMsg) IsPegMsg() bool { return false }
func (msg CreateAgentMsg) String() string {
	msgString := string(msg.GetSignBytes())
	return fmt.Sprintf("CreateAgentMsg: %v", msgString)
}

//UpdateAgentMsg
type UpdateAgentMsg struct {
	Data       UpdateAgentDoc `json:"data"`
	TxHash     string         `json:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid"`
}

var _ sdk.Msg = UpdateAgentMsg{}

func (msg UpdateAgentMsg) Type() string                            { return "project" }
func (msg UpdateAgentMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg UpdateAgentMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg UpdateAgentMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg UpdateAgentMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg UpdateAgentMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.GetProjectDid())}
}
func (msg UpdateAgentMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg UpdateAgentMsg) IsNewDid() bool { return false }
func (msg UpdateAgentMsg) IsPegMsg() bool { return false }

func (msg UpdateAgentMsg) String() string {
	msgString := string(msg.GetSignBytes())
	return fmt.Sprintf("UpdateAgentMsg: %v", msgString)
}

//CreateClaimMsg
type CreateClaimMsg struct {
	Data       CreateClaimDoc `json:"data"`
	TxHash     string         `json:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid"`
}

var _ sdk.Msg = CreateClaimMsg{}

func (msg CreateClaimMsg) Type() string                            { return "project" }
func (msg CreateClaimMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateClaimMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg CreateClaimMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateClaimMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateClaimMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.GetProjectDid())}
}
func (msg CreateClaimMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg CreateClaimMsg) IsNewDid() bool { return false }
func (msg CreateClaimMsg) IsPegMsg() bool { return false }
func (msg CreateClaimMsg) String() string {
	msgString := string(msg.GetSignBytes())
	return fmt.Sprintf("CreateClaimMsg: %v", msgString)
}

//CreateEvaluationMsg
type CreateEvaluationMsg struct {
	Data       CreateEvaluationDoc `json:"data"`
	TxHash     string              `json:"txHash"`
	SenderDid  ixo.Did             `json:"senderDid"`
	ProjectDid ixo.Did             `json:"projectDid"`
}

var _ sdk.Msg = CreateEvaluationMsg{}

func (msg CreateEvaluationMsg) Type() string                            { return "project" }
func (msg CreateEvaluationMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateEvaluationMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg CreateEvaluationMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateEvaluationMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateEvaluationMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.GetProjectDid())}
}
func (msg CreateEvaluationMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg CreateEvaluationMsg) IsNewDid() bool { return false }
func (msg CreateEvaluationMsg) IsPegMsg() bool { return false }
func (msg CreateEvaluationMsg) String() string {
	msgString := string(msg.GetSignBytes())
	return fmt.Sprintf("CreateEvaluationMsg: %v", msgString)
}

//FundProjectMsg
type FundProjectMsg struct {
	Data FundProjectDoc `json:"data"`
}

var _ sdk.Msg = FundProjectMsg{}

func (msg FundProjectMsg) Type() string                            { return "project" }
func (msg FundProjectMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg FundProjectMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg FundProjectMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.Data.GetSigner())}
}
func (msg FundProjectMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg FundProjectMsg) IsNewDid() bool { return false }
func (msg FundProjectMsg) IsPegMsg() bool { return true }
func (msg FundProjectMsg) String() string {
	msgString := string(msg.GetSignBytes())
	return fmt.Sprintf("FundProjectMsg: %v", msgString)
}

//WithdrawFundsMsg
type WithdrawFundsMsg struct {
	Data WithdrawFundsDoc `json:"data"`
}

var _ sdk.Msg = WithdrawFundsMsg{}

func (msg WithdrawFundsMsg) Type() string                            { return "project" }
func (msg WithdrawFundsMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg WithdrawFundsMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg WithdrawFundsMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.Data.GetProjectDid())}
}
func (msg WithdrawFundsMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg WithdrawFundsMsg) IsNewDid() bool { return false }
func (msg WithdrawFundsMsg) IsPegMsg() bool { return false }
func (msg WithdrawFundsMsg) String() string {
	msgString := string(msg.GetSignBytes())
	return fmt.Sprintf("WithdrawFundsMsg: %v", msgString)
}
