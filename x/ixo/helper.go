package ixo

import (
	"fmt"
	"os"
	"time"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/ed25519"
)

func SignIxoMessage(signBytes []byte, did string, privKey [64]byte) IxoSignature {
	
	signatureBytes := ed25519.Sign(&privKey, signBytes)
	signature := *signatureBytes
	
	return NewSignature(time.Now(), signature)
}

func VerifySignature(msg sdk.Msg, publicKey [32]byte, sig IxoSignature) bool {
	signatureBytes := [64]byte(sig.SignatureValue)
	result := ed25519.Verify(&publicKey, msg.GetSignBytes(), &signatureBytes)
	
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
