package project

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

//DOC SETUP
var _ ixo.ProjectDoc = (*BaseProjectDoc)(nil)

type ProjectDoc struct {
	Did              string    `json:"did"`
	PubKey           string    `json:"pubKey"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"shortDescription"`
	LongDescription  string    `json:"longDescription"`
	ImpactAction     string    `json:"impactAction"`
	CreatedOn        time.Time `json:"createdOn"`
	CreatedBy        string    `json:"createdBy"`
	Country          string    `json:"country"`
	Sdgs             []string  `json:"sdgs"`
	ImpactsRequired  string    `json:"impactsRequired"`
	ClaimTemplate    string    `json:"claimTemplate"`
	SocialMedia      struct {
		FacebookLink  string `json:"facebookLink"`
		InstagramLink string `json:"instagramLink"`
		TwitterLink   string `json:"twitterLink"`
	} `json:"socialMedia"`
	WebLink string `json:"webLink"`
	Image   string `json:"image"`
}

//GETTERS
func (pd ProjectDoc) GetDid() ixo.Did { return pd.Did }
func (pd ProjectDoc) GetData() string {
	data, err := json.Marshal(pd)
	if err != nil {
		panic(err)
	}
	return string(data)
}

//SETTERS
func (pd ProjectDoc) SetDid(did ixo.Did) error {
	if len(pd.Did) != 0 {
		return errors.New("cannot override BaseProjectDoc did")
	}
	pd.Did = did
	return nil
}
func (pd ProjectDoc) SetCreatedBy(did ixo.Did) error {
	if len(pd.CreatedBy) != 0 {
		return errors.New("cannot override BaseProjectDoc data")
	}
	pd.CreatedBy = did
	return nil
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

//**************************************************************************************

//ADD ProjectDoc
type AddProjectMsg struct {
	ProjectDoc ProjectDoc
}

func NewAddProjectMsg(projectDoc ProjectDoc) AddProjectMsg {
	return AddProjectMsg{
		ProjectDoc: projectDoc,
	}
}

var _ sdk.Msg = AddProjectMsg{}

func (msg AddProjectMsg) Type() string                            { return "project" }
func (msg AddProjectMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg AddProjectMsg) GetSigners() []sdk.Address {
	return []sdk.Address{[]byte(msg.ProjectDoc.GetDid())}
}
func (msg AddProjectMsg) String() string {
	return fmt.Sprintf("AddProjectMsg{Did: %v, data: %v}", string(msg.ProjectDoc.GetDid()), msg.ProjectDoc.GetData())
}

func (msg AddProjectMsg) ValidateBasic() sdk.Error {
	return nil
}

func (msg AddProjectMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

//**************************************************************************************

//GET ProjectDoc

type GetProjectMsg struct {
	Did ixo.Did `json:"did"`
}

func NewGetDidMsg(did string) GetProjectMsg {
	return GetProjectMsg{
		Did: did,
	}
}

var _ sdk.Msg = GetProjectMsg{}

func (msg GetProjectMsg) Type() string                            { return "project" }
func (msg GetProjectMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg GetProjectMsg) GetSigners() []sdk.Address               { return nil }
func (msg GetProjectMsg) String() string {
	return fmt.Sprintf("ProjectMsg{ProjectDoc: %v}", msg.Did)
}

func (msg GetProjectMsg) ValidateBasic() sdk.Error {
	return nil
}

func (msg GetProjectMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
