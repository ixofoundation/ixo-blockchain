package ixo

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

type DidDoc interface {
	SetDid(did Did) error
	GetDid() Did
	SetPubKey(pubkey string) error
	GetPubKey() string
}

// Define the ixo message type
type IxoMsg struct {
	Payload string `json:"payload"`
}

// New Ixo message
func NewIxoMsg(did string, payload string) IxoMsg {
	return IxoMsg{
		Payload: payload,
	}
}

// enforce the msg type at compile time
var _ sdk.Msg = IxoMsg{}

// nolint
func (msg IxoMsg) Type() string                            { return "ixo" }
func (msg IxoMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg IxoMsg) GetSigners() []sdk.Address               { return nil }
func (msg IxoMsg) String() string {
	return fmt.Sprintf("IxoMsg{Payload: %v}", msg.Payload)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg IxoMsg) ValidateBasic() sdk.Error {
	return nil
}

// Get the bytes for the message signer to sign on
func (msg IxoMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
