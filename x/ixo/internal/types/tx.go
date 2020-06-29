package types

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"gopkg.in/yaml.v2"
	"time"
)

const (
	maxGasWanted = uint64((1 << 63) - 1)
)

type IxoMsg interface {
	sdk.Msg
	GetSignerDid() exported.Did
}

var _ sdk.Tx = (*IxoTx)(nil)

type IxoTx struct {
	Msgs       []sdk.Msg      `json:"payload" yaml:"payload"`
	Fee        auth.StdFee    `json:"fee" yaml:"fee"`
	Signatures []IxoSignature `json:"signatures" yaml:"signatures"`
	Memo       string         `json:"memo" yaml:"memo"`
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

type IxoSignature struct {
	SignatureValue []byte    `json:"signatureValue" yaml:"signatureValue"`
	Created        time.Time `json:"created" yaml:"created"`
}

func NewIxoSignature(
	created time.Time, signature []byte,
) IxoSignature {
	return IxoSignature{
		SignatureValue: signature,
		Created:        created,
	}
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

func DefaultTxDecoder(cdc *codec.Codec) sdk.TxDecoder {
	return func(txBytes []byte) (sdk.Tx, sdk.Error) {

		if len(txBytes) == 0 {
			return nil, sdk.ErrTxDecode("txBytes are empty")
		}

		if string(txBytes[0:1]) == "{" {
			var tx IxoTx
			err := cdc.UnmarshalJSON(txBytes, &tx)
			if err != nil {
				return nil, sdk.ErrTxDecode("").TraceSDK(err.Error())
			}
			return tx, nil
		} else {
			var tx auth.StdTx
			err := cdc.UnmarshalBinaryLengthPrefixed(txBytes, &tx)
			if err != nil {
				return nil, sdk.ErrTxDecode("").TraceSDK(err.Error())
			}
			return tx, nil
		}
	}
}
