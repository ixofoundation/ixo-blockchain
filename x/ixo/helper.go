package ixo

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ed25519 "github.com/tendermint/ed25519"
	crypto "github.com/tendermint/go-crypto"
)

func SignIxoMessage(msg sdk.Msg, did string, privKey [64]byte) IxoSignature {
	signatureBytes := ed25519.Sign(&privKey, msg.GetSignBytes())
	signature := crypto.SignatureEd25519(*signatureBytes).Wrap()

	return NewSignature(time.Now(), did, signature)
}

func VerifySignature(msg sdk.Msg, publicKey [32]byte, sig sdk.StdSignature) bool {

	// First we unwrap the crypto.Signature to the crypto.SignatureEd25519 then we cast it to bytes
	innerSignature := sig.Signature.Unwrap().(crypto.SignatureEd25519)
	signatureBytes := [64]byte(innerSignature)
	result := ed25519.Verify(&publicKey, msg.GetSignBytes(), &signatureBytes)

	return result
}
