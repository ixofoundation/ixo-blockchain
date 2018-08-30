package project

import (
	"encoding/json"
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
	ProjectLocation  string   `json:"projectLocation"`
	RequiredClaims   int      `json:"requiredClaims"`
	Sdgs             []string `json:"sdgs"`
	Templates        struct {
		Claim struct {
			Schema string `json:"schema"`
			Form   string `json:"form"`
		} `json:"claim"`
	} `json:"templates"`
	EvaluatorPayPerClaim string `json:"evaluatorPayPerClaim"`
	SocialMedia          struct {
		FacebookLink  string `json:"facebookLink"`
		InstagramLink string `json:"instagramLink"`
		TwitterLink   string `json:"twitterLink"`
		WebLink       string `json:"webLink"`
	} `json:"socialMedia"`
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
	CreatedOn string `json:"createdOn"`
	CreatedBy string `json:"createdBy"`
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

type ProjectMsg struct {
	Original   string  `json:"original"`
	TxHash     string  `json:"txHash"`
	SenderDid  ixo.Did `json:"senderDid"`
	ProjectDid ixo.Did `json:"projectDid"`
}

func (msg ProjectMsg) IsNewDid() bool                          { return false }
func (msg ProjectMsg) Type() string                            { return "project" }
func (msg ProjectMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg ProjectMsg) ValidateBasic() sdk.Error {
	return nil
}
func (msg ProjectMsg) GetSignBytes() []byte {
	return []byte(msg.Original)
}
func (msg ProjectMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg ProjectMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg ProjectMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.GetProjectDid())}
}
func (msg ProjectMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

//CreateProjectMsg
type CreateProjectMsg struct {
	ProjectMsg
	PubKey string     `json:"pubKey"`
	Data   ProjectDoc `json:"data"`
}

var _ sdk.Msg = CreateProjectMsg{}

func (msg CreateProjectMsg) IsNewDid() bool         { return true }
func (msg CreateProjectMsg) GetPubKey() string      { return msg.PubKey }
func (msg CreateProjectMsg) GetEvaluatorPay() int64 { return msg.Data.GetEvaluatorPay() }

type StoredProjectDoc = CreateProjectMsg

var _ ixo.StoredProjectDoc = (*StoredProjectDoc)(nil)

//CreateAgentMsg
type CreateAgentMsg struct {
	ProjectMsg
	Data CreateAgentDoc `json:"data"`
}

var _ sdk.Msg = CreateAgentMsg{}

//UpdateAgentMsg
type UpdateAgentMsg struct {
	ProjectMsg
	Data UpdateAgentDoc `json:"data"`
}

var _ sdk.Msg = UpdateAgentMsg{}

//CreateClaimMsg
type CreateClaimMsg struct {
	ProjectMsg
	Data CreateClaimDoc `json:"data"`
}

var _ sdk.Msg = CreateClaimMsg{}

//CreateEvaluationMsg
type CreateEvaluationMsg struct {
	ProjectMsg
	Data CreateEvaluationDoc `json:"data"`
}

var _ sdk.Msg = CreateEvaluationMsg{}

//FundProjectMsg
type FundProjectMsg struct {
	ProjectMsg
	Data FundProjectDoc `json:"data"`
}

var _ sdk.Msg = FundProjectMsg{}

//WithdrawFundsMsg
type WithdrawFundsMsg struct {
	ProjectMsg
	Data WithdrawFundsDoc `json:"data"`
}

var _ sdk.Msg = WithdrawFundsMsg{}
