package ixo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ed25519 "github.com/tendermint/ed25519"
	crypto "github.com/tendermint/go-crypto"
)

func SignIxoMessage(msg sdk.Msg, did string, privKey [64]byte) IxoSignature {
	fmt.Println("*******SIGNING_MSG******* \n", string(msg.GetSignBytes()))

	signatureBytes := ed25519.Sign(&privKey, msg.GetSignBytes())
	signature := crypto.SignatureEd25519(*signatureBytes).Wrap()

	return NewSignature(time.Now(), signature)
}

func VerifySignature(msg sdk.Msg, publicKey [32]byte, sig sdk.StdSignature) bool {
	// First we unwrap the crypto.Signature to the crypto.SignatureEd25519 then we cast it to bytes
	innerSignature := sig.Signature.Unwrap().(crypto.SignatureEd25519)
	signatureBytes := [64]byte(innerSignature)
	result := ed25519.Verify(&publicKey, msg.GetSignBytes(), &signatureBytes)

	if !result {
		fmt.Println("******* VERIFY_MSG: Failed ******* ")
		fmt.Println(string(msg.GetSignBytes()))

	}
	return result
}

func JSONStringify(msg sdk.Msg) []byte {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "")
	err := encoder.Encode(msg)

	if err != nil {
		panic(err)
	}

	return bytes.Trim(buffer.Bytes(), " \n")

}
