package sovrin

import (
	"bytes"
	crypto_rand "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	naclbox "golang.org/x/crypto/nacl/box"

	base58 "github.com/btcsuite/btcutil/base58"
	ed25519 "golang.org/x/crypto/ed25519"
)

type SovrinSecret struct {
	Seed                 string `json:"seed"`
	SignKey              string `json:"signKey"`
	EncryptionPrivateKey string `json:"encryptionPrivateKey"`
}

func (ss SovrinSecret) String() string {
	output, err := json.MarshalIndent(ss, "", "  ")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", string(output))
}

type SovrinDid struct {
	Did                 string       `json:"did"`
	VerifyKey           string       `json:"verifyKey"`
	Kyc                 bool         `json:"kyc"`
	EncryptionPublicKey string       `json:"encryptionPublicKey"`
	Secret              SovrinSecret `json:"secret"`
}

func (sd SovrinDid) String() string {
	output, err := json.MarshalIndent(sd, "", "  ")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", string(output))
}

func FromSeed(seed [32]byte) SovrinDid {

	publicKeyBytes, privateKeyBytes, err := ed25519.GenerateKey(bytes.NewReader(seed[0:32]))
	if err != nil {
		panic(err)
	}
	publicKey := []byte(publicKeyBytes)
	privateKey := []byte(privateKeyBytes)

	signKey := base58.Encode(privateKey[:32])
	keyPair_publicKey, keyPair_privateKey, err := naclbox.GenerateKey(bytes.NewReader(privateKey[:]))

	sovDid := SovrinDid{
		Did:                 base58.Encode(publicKey[:16]),
		VerifyKey:           base58.Encode(publicKey),
		EncryptionPublicKey: base58.Encode(keyPair_publicKey[:]),

		Secret: SovrinSecret{
			Seed:                 hex.EncodeToString(seed[0:32]),
			SignKey:              signKey,
			EncryptionPrivateKey: base58.Encode(keyPair_privateKey[:]),
		},
	}

	return sovDid
}

func Gen() SovrinDid {
	var seed [32]byte
	if _, err := io.ReadFull(crypto_rand.Reader, seed[:]); err != nil {
		panic(err)
	}
	return FromSeed(seed)
}

func SignMessage(message []byte, signKey string, verifyKey string) []byte {
	// Force the length to 64
	fmt.Println("*******SIGNING_MSG******* \n", string(message))

	privateKey := make([]byte, ed25519.PrivateKeySize)
	fullPrivKey := ed25519.PrivateKey(privateKey)
	copy(fullPrivKey[:], getArrayFromKey(signKey))
	copy(fullPrivKey[32:], getArrayFromKey(verifyKey))

	return ed25519.Sign(fullPrivKey, message)
}

func VerifySignedMessage(message []byte, signature []byte, verifyKey string) bool {
	publicKey := ed25519.PublicKey{}
	copy(publicKey[:], getArrayFromKey(verifyKey))
	result := ed25519.Verify(publicKey, message, signature)

	return result
}

func GetNonce() [24]byte {
	var nonce [24]byte
	if _, err := io.ReadFull(crypto_rand.Reader, nonce[:]); err != nil {
		panic(err)
	}
	return nonce
}

func getArrayFromKey(key string) []byte {
	fmt.Println(len(base58.Decode(key)))
	return base58.Decode(key)
}

func GetKeyPairFromSignKey(signKey string) ([32]byte, [32]byte) {
	publicKey, privateKey, err := naclbox.GenerateKey(bytes.NewReader(getArrayFromKey(signKey)))
	if err != nil {
		panic(err)
	}
	return *publicKey, *privateKey
}
