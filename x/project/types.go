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
var _ ixo.ProjectDoc = (*ProjectDoc)(nil)

type ProjectDoc struct {
	TxHash           string   `json:"txHash"`
	SenderDid        string   `json:"senderDid"`
	ProjectDid       string   `json:"projectDid"`
	PubKey           string   `json:"pubKey"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"shortDescription"`
	LongDescription  string   `json:"longDescription"`
	ImpactAction     string   `json:"impactAction"`
	CreatedOn        string   `json:"createdOn"`
	CreatedBy        string   `json:"createdBy"`
	Country          string   `json:"country"`
	Sdgs             []string `json:"sdgs"`
	ImpactsRequired  int      `json:"impactsRequired"`
	ClaimTemplate    string   `json:"claimTemplate"`
	EvaluatorPay     string   `json:"evaluatorPay"`
	SocialMedia      struct {
		FacebookLink  string `json:"facebookLink"`
		InstagramLink string `json:"instagramLink"`
		TwitterLink   string `json:"twitterLink"`
		WebLink       string `json:"webLink"`
	} `json:"socialMedia"`
	ServiceEndpoint string `json:"serviceEndpoint"`
	ImageLink       string `json:"imageLink"`
}

//GETTERS
func (pd ProjectDoc) GetProjectDid() ixo.Did { return pd.ProjectDid }
func (pd ProjectDoc) GetPubKey() string      { return pd.PubKey }

func (pd ProjectDoc) GetEvaluatorPay() int64 {
	i, err := strconv.ParseInt(pd.EvaluatorPay, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

//DECODERS
type ProjectDocDecoder func(projectEntryBytes []byte) (ixo.ProjectDoc, error)

func GetProjectDocDecoder(cdc *wire.Codec) ProjectDocDecoder {
	return func(projectDocBytes []byte) (res ixo.ProjectDoc, err error) {
		if len(projectDocBytes) == 0 {
			return nil, sdk.ErrTxDecode("projectDocBytes are empty")
		}
		projectDoc := ProjectDoc{}
		err = cdc.UnmarshalBinary(projectDocBytes, &projectDoc)
		if err != nil {
			panic(err)
		}
		return projectDoc, err
	}
}

// Define CreateAgent
type CreateAgentDoc struct {
	TxHash     string  `json:"txHash"`
	SenderDid  ixo.Did `json:"senderDid"`
	ProjectDid ixo.Did `json:"projectDid"`
	AgentDid   ixo.Did `json:"did"`
	Role       string  `json:"role"`
}

func (cd CreateAgentDoc) GetProjectDid() ixo.Did { return cd.ProjectDid }

type AgentStatus int

const (
	PendingAgent  AgentStatus = 0
	ApprovedAgent AgentStatus = 1
	RevokedAgent  AgentStatus = 2
)

// Define CreateAgent
type UpdateAgentDoc struct {
	TxHash     string      `json:"txHash"`
	SenderDid  ixo.Did     `json:"senderDid"`
	ProjectDid ixo.Did     `json:"projectDid"`
	Did        ixo.Did     `json:"did"`
	Status     AgentStatus `json:"status"`
}

func (ud UpdateAgentDoc) GetProjectDid() ixo.Did { return ud.ProjectDid }
func (ud UpdateAgentDoc) GetSenderDid() ixo.Did  { return ud.SenderDid }

// Define CreateAgent
type CreateClaimDoc struct {
	TxHash     string  `json:"txHash"`
	SenderDid  ixo.Did `json:"senderDid"`
	ProjectDid ixo.Did `json:"projectDid"`
	ClaimID    string  `json:"claimID"`
}

func (cd CreateClaimDoc) GetProjectDid() ixo.Did { return cd.ProjectDid }
func (cd CreateClaimDoc) GetSenderDid() ixo.Did  { return cd.SenderDid }

type ClaimStatus int

const (
	PendingClaim  ClaimStatus = 0
	ApprovedClaim ClaimStatus = 1
	RejectedClaim ClaimStatus = 2
)

// Define CreateAgent
type CreateEvaluationDoc struct {
	TxHash     string      `json:"txHash"`
	SenderDid  ixo.Did     `json:"senderDid"`
	ProjectDid ixo.Did     `json:"projectDid"`
	ClaimID    string      `json:"claimID"`
	Status     ClaimStatus `json:"status"`
}

func (cd CreateEvaluationDoc) GetProjectDid() ixo.Did { return cd.ProjectDid }
func (cd CreateEvaluationDoc) GetSenderDid() ixo.Did  { return cd.SenderDid }

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
	ProjectDoc ProjectDoc `json:"data"`
}

var _ sdk.Msg = CreateProjectMsg{}

func (msg CreateProjectMsg) Type() string                            { return "project" }
func (msg CreateProjectMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateProjectMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg CreateProjectMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.ProjectDoc.GetProjectDid())}
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
	return fmt.Sprintf("CreateProjectMsg{Did: %v, projectDoc: %v}", string(msg.ProjectDoc.GetProjectDid()), msg.ProjectDoc)
}

//CreateAgentMsg
type CreateAgentMsg struct {
	Data CreateAgentDoc `json:"data"`
}

var _ sdk.Msg = CreateAgentMsg{}

func (msg CreateAgentMsg) Type() string                            { return "project" }
func (msg CreateAgentMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateAgentMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg CreateAgentMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.Data.GetProjectDid())}
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
	Data UpdateAgentDoc `json:"data"`
}

var _ sdk.Msg = UpdateAgentMsg{}

func (msg UpdateAgentMsg) Type() string                            { return "project" }
func (msg UpdateAgentMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg UpdateAgentMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg UpdateAgentMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.Data.GetProjectDid())}
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
	Data CreateClaimDoc `json:"data"`
}

var _ sdk.Msg = CreateClaimMsg{}

func (msg CreateClaimMsg) Type() string                            { return "project" }
func (msg CreateClaimMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateClaimMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg CreateClaimMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.Data.GetProjectDid())}
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
	Data CreateEvaluationDoc `json:"data"`
}

var _ sdk.Msg = CreateEvaluationMsg{}

func (msg CreateEvaluationMsg) Type() string                            { return "project" }
func (msg CreateEvaluationMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateEvaluationMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg CreateEvaluationMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.Data.GetProjectDid())}
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
