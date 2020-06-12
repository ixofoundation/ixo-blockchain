package ixo

import (
	"fmt"
	"os"
	"time"

	"github.com/tendermint/ed25519"
)

func SignIxoMessage(signBytes []byte, privKey [64]byte) IxoSignature {

	signatureBytes := ed25519.Sign(&privKey, signBytes)
	signature := *signatureBytes

	return NewSignature(time.Now(), signature)
}

func VerifySignature(signBytes []byte, publicKey [32]byte, sig IxoSignature) bool {
	result := ed25519.Verify(&publicKey, signBytes, &sig.SignatureValue)
	if !result {
		fmt.Println("******* VERIFY_MSG: Failed ******* ")
	}
	return result
}

func LookupEnv(name string, defaultValue string) string {
	val, found := os.LookupEnv(name)
	if found && len(val) > 0 {
		return val
	}
	return defaultValue
}
