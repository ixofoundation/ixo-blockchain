package project

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//_______________________________________________________________________
// Define the ProjectTx

type ProjectTx struct {
	Msg sdk.Msg
}

func NewProjectTx(msg sdk.Msg) sdk.Tx {
	return ProjectTx{
		Msg: msg,
	}
}

func (tx ProjectTx) GetMsg() sdk.Msg                   { return tx.Msg }
func (tx ProjectTx) GetSignatures() []sdk.StdSignature { return nil }

// enforce the msg type at compile time
var _ sdk.Tx = ProjectTx{}

// Define Did as an Address
type Did = sdk.Address

// Define the project message type
type ProjectMsg struct {
	Payload string `json:"payload"`
}

// New Project message
func NewProjectMsg(did string, payload string) ProjectMsg {
	return ProjectMsg{
		Payload: payload,
	}
}

// enforce the msg type at compile time
var _ sdk.Msg = ProjectMsg{}

// nolint
func (msg ProjectMsg) Type() string                            { return "project" }
func (msg ProjectMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg ProjectMsg) GetSigners() []sdk.Address               { return nil }
func (msg ProjectMsg) String() string {
	return fmt.Sprintf("ProjectMsg{Payload: %v}", msg.Payload)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg ProjectMsg) ValidateBasic() sdk.Error {
	return nil
}

// Get the bytes for the message signer to sign on
func (msg ProjectMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
