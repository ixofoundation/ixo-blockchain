package project

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// A really project msg type, these fields are can be entirely arbitrary and
// custom to your message
type SetTrendMsg struct {
	Sender sdk.Address
	Project   string
}

// Genesis state - specify genesis trend
type ProjectGenesis struct {
	Trend string `json:"trend"`
}

// New project message
func NewSetTrendMsg(sender sdk.Address, project string) SetTrendMsg {
	return SetTrendMsg{
		Sender: sender,
		Project:   project,
	}
}

// enforce the msg type at compile time
var _ sdk.Msg = SetTrendMsg{}

// nolint
func (msg SetTrendMsg) Type() string                            { return "project" }
func (msg SetTrendMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg SetTrendMsg) GetSigners() []sdk.Address               { return []sdk.Address{msg.Sender} }
func (msg SetTrendMsg) String() string {
	return fmt.Sprintf("SetTrendMsg{Sender: %v, Project: %v}", msg.Sender, msg.Project)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg SetTrendMsg) ValidateBasic() sdk.Error {
	if len(msg.Sender) == 0 {
		return sdk.ErrUnknownAddress(msg.Sender.String()).Trace("")
	}
	if strings.Contains(msg.Project, "hot") {
		return sdk.ErrUnauthorized("").Trace("hot is not project")
	}
	if strings.Contains(msg.Project, "warm") {
		return sdk.ErrUnauthorized("").Trace("warm is not very project")
	}
	return nil
}

// Get the bytes for the message signer to sign on
func (msg SetTrendMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

//_______________________________________________________________________

// A message type to quiz how project you are. these fields are can be entirely
// arbitrary and custom to your message
type QuizMsg struct {
	Sender     sdk.Address
	ProjectAnswer string
}

// New project message
func NewQuizMsg(sender sdk.Address, answer string) QuizMsg {
	return QuizMsg{
		Sender:     sender,
		ProjectAnswer: answer,
	}
}

// enforce the msg type at compile time
var _ sdk.Msg = QuizMsg{}

// nolint
func (msg QuizMsg) Type() string                            { return "project" }
func (msg QuizMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg QuizMsg) GetSigners() []sdk.Address               { return []sdk.Address{msg.Sender} }
func (msg QuizMsg) String() string {
	return fmt.Sprintf("QuizMsg{Sender: %v, Answer: %v}", msg.Sender, msg.ProjectAnswer)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg QuizMsg) ValidateBasic() sdk.Error {
	if len(msg.Sender) == 0 {
		return sdk.ErrUnknownAddress(msg.Sender.String()).Trace("")
	}
	return nil
}

// Get the bytes for the message signer to sign on
func (msg QuizMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
