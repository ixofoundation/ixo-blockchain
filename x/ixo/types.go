package ixo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

var (
	IxoDecimals = sdk.NewDec(100000000)
	ValidDid    = regexp.MustCompile(`^did:(ixo:|sov:)([a-zA-Z0-9]){21,22}([/][a-zA-Z0-9:]+|)$`)
	IsValidDid  = ValidDid.MatchString
	// https://sovrin-foundation.github.io/sovrin/spec/did-method-spec-template.html
	// IsValidDid adapted from the above link but assumes no sub-namespaces
	// TODO: ValidDid needs to be updated once we no longer want to be able
	//   to consider project accounts as DIDs (especially in treasury module),
	//   possibly should just be `^did:(ixo:|sov:)([a-zA-Z0-9]){21,22}$`.
)

const IxoNativeToken = "ixo"

type IxoTx struct {
	Msgs       []sdk.Msg      `json:"payload" yaml:"payload"`
	Signatures []IxoSignature `json:"signatures" yaml:"signatures"`
}

type IxoSignature struct {
	SignatureValue [64]byte  `json:"signatureValue" yaml:"signatureValue"`
	Created        time.Time `json:"created" yaml:"created"`
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

func (tx IxoTx) GetMemo() string { return "" }

func (tx IxoTx) ValidateBasic() sdk.Error { return nil }

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

func FeePayer(tx sdk.Tx) sdk.AccAddress {
	return tx.GetMsgs()[0].GetSigners()[0]
}

var _ sdk.Tx = (*IxoTx)(nil)

type Did = string

type DidDoc interface {
	SetDid(did Did) error
	GetDid() Did
	SetPubKey(pubkey string) error
	GetPubKey() string
}

func DefaultTxDecoder(cdc *codec.Codec) sdk.TxDecoder {
	return func(txBytes []byte) (sdk.Tx, sdk.Error) {

		if len(txBytes) == 0 {
			return nil, sdk.ErrTxDecode("txBytes are empty")
		}

		txByteString := string(txBytes[0:1])
		if txByteString == "{" {
			var tx = IxoTx{}

			var upTx map[string]interface{}

			err := json.Unmarshal(txBytes, &upTx)
			if err != nil {
				return nil, sdk.ErrTxDecode(err.Error())
			}

			payloadArray := upTx["payload"].([]interface{})
			if len(payloadArray) != 1 {
				return nil, sdk.ErrTxDecode("Multiple messages not supported")
			}

			txBytes, _ = json.Marshal(upTx)

			err = cdc.UnmarshalJSON(txBytes, &tx)
			if err != nil {
				return nil, sdk.ErrTxDecode("").TraceSDK(err.Error())
			}

			return tx, nil

		} else {
			var tx = auth.StdTx{}
			err := cdc.UnmarshalBinaryLengthPrefixed(txBytes, &tx)
			if err != nil {
				return nil, sdk.ErrTxDecode("").TraceSDK(err.Error())
			}
			return tx, nil

		}
	}
}

func getSignBytes(txBytes []byte) string {
	const strtTxt string = "\"value\":"
	const endTxt string = "}],\"signatures\":"

	strt := bytes.Index(txBytes, []byte(strtTxt)) + len(strtTxt)
	end := bytes.Index(txBytes, []byte(endTxt))

	signBytes := txBytes[strt:end]
	return string(signBytes)
}
