package ixo

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//_______________________________________________________________________
// Define the IxoTx

type IxoTx struct {
	Msgs       []sdk.Msg      `json:"payload"`
	Signatures []IxoSignature `json:"signatures"`
}

type IxoSignature struct {
	SignatureValue [64]byte  `json:"signatureValue"`
	Created        time.Time `json:"created"`
}

func NewSignature(created time.Time, signature [64]byte) IxoSignature {
	return IxoSignature{
		SignatureValue: signature,
		Created:        created,
	}
}

func NewIxoTx(msgs []sdk.Msg, sigs []IxoSignature) IxoTx {
	return IxoTx{
		Msgs:       msgs,
		Signatures: sigs,
	}
}

func NewIxoTxSingleMsg(msg sdk.Msg, signature IxoSignature) IxoTx {
	sigs := make([]IxoSignature, 0)
	sigs = append(sigs, signature)

	msgs := make([]sdk.Msg, 0)
	msgs = append(msgs, msg)

	return IxoTx{
		Msgs:       msgs,
		Signatures: sigs,
	}
}

func (tx IxoTx) GetMsgs() []sdk.Msg { return tx.Msgs }

//nolint
func (tx IxoTx) GetMemo() string { return "" }

func (tx IxoTx) GetSignatures() []IxoSignature {
	return tx.Signatures
}

func (tx IxoTx) String() string {
	output, err := json.MarshalIndent(tx, "", "  ")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", string(output))
}

// FeePayer returns the address responsible for paying the fees
// for the transactions. It's the first address returned by msg.GetSigners().
// If GetSigners() is empty, this panics.
func FeePayer(tx sdk.Tx) sdk.AccAddress {
	return tx.GetMsgs()[0].GetSigners()[0]
}

// enforce the msg type at compile time
var _ sdk.Tx = IxoTx{}

// Define Did as an Address
type Did = string

type DidDoc interface {
	SetDid(did Did) error
	GetDid() Did
	SetPubKey(pubkey string) error
	GetPubKey() string
}

// Define Project
type Project = string
