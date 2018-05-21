package project

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

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
	TxHash    string  `json:"txHash"`
	SenderDid string  `json:"senderDid"`
	Did       ixo.Did `json:"did"`
	Role      string  `json:"role"`
}

type AgentStatus int

const (
	PendingAgent  AgentStatus = 0
	ApprovedAgent AgentStatus = 1
	RevokedAgent  AgentStatus = 2
)

// Define CreateAgent
type UpdateAgentDoc struct {
	TxHash    string      `json:"txHash"`
	SenderDid string      `json:"senderDid"`
	Did       ixo.Did     `json:"did"`
	Status    AgentStatus `json:"status"`
}

// Define CreateAgent
type CreateClaimDoc struct {
	TxHash    string `json:"txHash"`
	SenderDid string `json:"senderDid"`
	ClaimID   string `json:"claimID"`
}

type ClaimStatus int

const (
	PendingClaim  ClaimStatus = 0
	ApprovedClaim ClaimStatus = 1
	RejectedClaim ClaimStatus = 2
)

// Define CreateAgent
type CreateEvaluationDoc struct {
	TxHash    string      `json:"txHash"`
	SenderDid string      `json:"senderDid"`
	ClaimID   string      `json:"claimID"`
	Status    ClaimStatus `json:"status"`
}

//**************************************************************************************
// Message

type ProjectMsg interface {
	IsNewDid() bool
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
func (msg CreateProjectMsg) IsNewDid() bool {
	return true
}
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
	return []sdk.Address{}
}
func (msg CreateAgentMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg CreateAgentMsg) IsNewDid() bool {
	return false
}
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
	return []sdk.Address{}
}
func (msg UpdateAgentMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg UpdateAgentMsg) IsNewDid() bool {
	return false
}
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
	return []sdk.Address{}
}
func (msg CreateClaimMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg CreateClaimMsg) IsNewDid() bool {
	return false
}
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
	return []sdk.Address{}
}
func (msg CreateEvaluationMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg CreateEvaluationMsg) IsNewDid() bool {
	return false
}
func (msg CreateEvaluationMsg) String() string {
	msgString := string(msg.GetSignBytes())
	return fmt.Sprintf("CreateEvaluationMsg: %v", msgString)
}
