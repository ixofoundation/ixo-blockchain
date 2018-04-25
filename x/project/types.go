package project

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//_______________________________________________________________________

// A message type to quiz how cool you are. these fields are can be entirely
// arbitrary and custom to your message
type NewProjectMsg struct {
	Sender sdk.Address
	Data   string
}

// New Project message
func CreateProjectMsg(projectData string) NewProjectMsg {
	return NewProjectMsg{
		Sender: sdk.Address([]byte("3J56r8ZGfD6ThhwhaDv9iA")),
		Data:   projectData,
	}
}

// enforce the msg type at compile time
var _ sdk.Msg = NewProjectMsg{}

// nolint
func (msg NewProjectMsg) Type() string                            { return "project" }
func (msg NewProjectMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg NewProjectMsg) GetSigners() []sdk.Address               { return []sdk.Address{msg.Sender} }
func (msg NewProjectMsg) String() string {
	return fmt.Sprintf("NewProjectMsg{Data: %v}", msg.Data)
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
