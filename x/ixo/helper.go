package ixo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ed25519 "github.com/tendermint/ed25519"
)

func SignIxoMessage(signBytes []byte, did string, privKey [64]byte) IxoSignature {
	fmt.Println("*******SIGNING_MSG*******")
	fmt.Println(string(signBytes))

	signatureBytes := ed25519.Sign(&privKey, signBytes)
	signature := *signatureBytes

	return NewSignature(time.Now(), signature)
}

func VerifySignature(msg sdk.Msg, publicKey [32]byte, sig IxoSignature) bool {
	// First we unwrap the crypto.Signature to the crypto.SignatureEd25519 then we cast it to bytes
	signatureBytes := [64]byte(sig.SignatureValue)
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
