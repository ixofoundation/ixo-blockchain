package project

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//_______________________________________________________________________
// Define the IxoTx

type IxoTx struct {
	Msg sdk.Msg
}

func NewIxoTx(msg sdk.Msg) sdk.Tx {
	return IxoTx{
		Msg: msg,
	}
}

func (tx IxoTx) GetMsg() sdk.Msg                   { return tx.Msg }
func (tx IxoTx) GetSignatures() []sdk.StdSignature { return nil }

// enforce the msg type at compile time
var _ sdk.Tx = IxoTx{}

// Define Did as an Address
type Did = sdk.Address

// A message type to quiz how cool you are. these fields are can be entirely
// arbitrary and custom to your message
type NewProjectMsg struct {
	Payload string `json:"payload"`
}

// New Project message
func CreateProjectMsg(did string, projectData string) NewProjectMsg {
	return NewProjectMsg{
		Payload: projectData,
	}
}

// enforce the msg type at compile time
var _ sdk.Msg = NewProjectMsg{}

// nolint
func (msg NewProjectMsg) Type() string                            { return "project" }
func (msg NewProjectMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg NewProjectMsg) GetSigners() []sdk.Address               { return nil }
func (msg NewProjectMsg) String() string {
	return fmt.Sprintf("NewProjectMsg{Data: %v}", msg.Payload)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg NewProjectMsg) ValidateBasic() sdk.Error {
	return nil
}

// Get the bytes for the message signer to sign on
func (msg NewProjectMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
