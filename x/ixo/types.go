package ixo

import (
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/crypto"
	"gopkg.in/yaml.v2"
	"regexp"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

var (
	IxoDecimals = sdk.NewDec(100000000)

	maxGasWanted = uint64((1 << 63) - 1)

	ValidDid   = regexp.MustCompile(`^did:(ixo:|sov:)([a-zA-Z0-9]){21,22}([/][a-zA-Z0-9:]+|)$`)
	IsValidDid = ValidDid.MatchString
	// https://sovrin-foundation.github.io/sovrin/spec/did-method-spec-template.html
	// IsValidDid adapted from the above link but assumes no sub-namespaces
	// TODO: ValidDid needs to be updated once we no longer want to be able
	//   to consider project accounts as DIDs (especially in treasury module),
	//   possibly should just be `^did:(ixo:|sov:)([a-zA-Z0-9]){21,22}$`.
)

const IxoNativeToken = "ixo"

func StringToAddr(str string) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(str)))
}

func DidToAddr(did Did) sdk.AccAddress {
	return StringToAddr(did)
}

type IxoTx struct {
	Msgs       []sdk.Msg      `json:"payload" yaml:"payload"`
	Fee        auth.StdFee    `json:"fee" yaml:"fee"`
	Signatures []IxoSignature `json:"signatures" yaml:"signatures"`
	Memo       string         `json:"memo" yaml:"memo"`
}

type IxoSignature struct {
	SignatureValue [64]byte  `json:"signatureValue" yaml:"signatureValue"`
	Created        time.Time `json:"created" yaml:"created"`
}

// MarshalYAML returns the YAML representation of the signature.
func (is IxoSignature) MarshalYAML() (interface{}, error) {
	var (
		bz  []byte
		err error
	)

	bz, err = yaml.Marshal(struct {
		SignatureValue string
		Created        string
	}{
		SignatureValue: fmt.Sprintf("%s", is.SignatureValue),
		Created:        is.Created.String(),
	})
	if err != nil {
		return nil, err
	}

	return string(bz), err
}

type IxoMsg interface {
	sdk.Msg
	GetSignerDid() Did
}

func NewSignature(created time.Time, signature [64]byte) IxoSignature {
	return IxoSignature{
		SignatureValue: signature,
		Created:        created,
	}
}

func NewIxoTx(msgs []sdk.Msg, fee auth.StdFee, sigs []IxoSignature, memo string) IxoTx {
	return IxoTx{
		Msgs:       msgs,
		Fee:        fee,
		Signatures: sigs,
		Memo:       memo,
	}
}

func NewIxoTxSingleMsg(msg sdk.Msg, fee auth.StdFee, signature IxoSignature, memo string) IxoTx {
	return NewIxoTx([]sdk.Msg{msg}, fee, []IxoSignature{signature}, memo)
}

func (tx IxoTx) GetMsgs() []sdk.Msg { return tx.Msgs }

func (tx IxoTx) GetMemo() string { return "" }

func (tx IxoTx) ValidateBasic() sdk.Error {
	// Fee validation
	if tx.Fee.Gas > maxGasWanted {
		return sdk.ErrGasOverflow(fmt.Sprintf("invalid gas supplied; %d > %d", tx.Fee.Gas, maxGasWanted))
	}
	if tx.Fee.Amount.IsAnyNegative() {
		return sdk.ErrInsufficientFee(fmt.Sprintf("invalid fee %s amount provided", tx.Fee.Amount))
	}

	// Signatures validation
	var ixoSigs = tx.GetSignatures()
	if len(ixoSigs) == 0 {
		return sdk.ErrNoSignatures("no signers")
	}
	if len(ixoSigs) != 1 {
		return sdk.ErrUnauthorized("there can only be one signer")
	}

	// Messages validation
	if len(tx.Msgs) != 1 {
		return sdk.ErrUnauthorized("there can only be one message")
	}

	return nil
}

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

func (tx IxoTx) GetSigner() sdk.AccAddress {
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

		if string(txBytes[0:1]) == "{" {
			var upTx map[string]interface{}
			err := json.Unmarshal(txBytes, &upTx)
			if err != nil {
				return nil, sdk.ErrTxDecode(err.Error())
			}

			payloadArray := upTx["payload"].([]interface{})
			if len(payloadArray) != 1 {
				return nil, sdk.ErrTxDecode("Multiple messages not supported")
			}

			var tx IxoTx
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
