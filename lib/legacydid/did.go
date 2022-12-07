package legacydid

import (
	"bytes"
	ed25519Local "crypto/ed25519"
	cryptoRand "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	fmt "fmt"
	"io"
	"regexp"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/gogo/protobuf/proto"
	naclBox "golang.org/x/crypto/nacl/box"
)

var (
	ValidDid      = regexp.MustCompile(`^did\:[a-z0-9]+\:(?:([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])*\:)*(?:[A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])+(?:#[a-zA-Z0-9-\._]+)?$`)
	ValidPubKey   = regexp.MustCompile(`^[123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ]{43,44}$`)
	IsValidDid    = ValidDid.MatchString
	IsValidPubKey = ValidPubKey.MatchString
	// https://sovrin-foundation.github.io/sovrin/spec/did-method-spec-template.html
	// IsValidDid adapted from the above link but assumes no sub-namespaces
	// TODO: ValidDid needs to be updated once we no longer want to be able
	//   to consider project accounts as DIDs (especially in treasury module),
	//   possibly should just be `^did:(ixo:|sov:)([a-zA-Z0-9]){21,22}$`.
)

var DidPrefix = "did:ixo:"

type Did = string

func fromJsonString(jsonIxoDid string) (IxoDid, error) {
	var did IxoDid
	err := json.Unmarshal([]byte(jsonIxoDid), &did)
	if err != nil {
		err := fmt.Errorf("could not unmarshal did into struct due to error: %s", err.Error())
		return IxoDid{}, err
	}

	return did, nil
}

func UnmarshalIxoDid(jsonIxoDid string) (IxoDid, error) {
	return fromJsonString(jsonIxoDid)
}

func UnprefixedDid(did Did) string {
	// Assumes that DID is valid (check IsValidDid regex)
	// Removes 8 characters (for did:ixo: or did:sov:)
	return did[8:]
}

func UnprefixedDidFromPubKey(pubKey string) string {
	// Assumes that PubKey is valid (check IsValidPubKey regex)
	// Since result is not prefixed (did:ixo:), string returned rather than DID
	pubKeyBz := base58.Decode(pubKey)
	return base58.Encode(pubKeyBz[:16])
}

type DidDoc interface {
	proto.Message

	SetDid(did Did) error
	GetDid() Did
	SetPubKey(pubkey string) error
	GetPubKey() string
	Address() sdk.AccAddress
}

func NewSecret(seed, signKey, encryptionPrivateKey string) Secret {
	return Secret{
		Seed:                 seed,
		SignKey:              signKey,
		EncryptionPrivateKey: encryptionPrivateKey,
	}
}

func (s Secret) Equals(other Secret) bool {
	return s.Seed == other.Seed &&
		s.SignKey == other.SignKey &&
		s.EncryptionPrivateKey == other.EncryptionPrivateKey
}

// Above IxoDid modelled after Sovrin documents
// Ref: https://www.npmjs.com/package/sovrin-did
// {
//    did: "<base58 did>",
//    verifyKey: "<base58 publicKey>",
//    publicKey: "<base58 publicKey>",
//
//    secret: {
//        seed: "<hex encoded 32-byte seed>",
//        signKey: "<base58 secretKey>",
//        privateKey: "<base58 privateKey>"
//    }
// }

func NewIxoDid(did, verifyKey, encryptionPublicKey string, secret Secret) IxoDid {
	return IxoDid{
		Did:                 did,
		VerifyKey:           verifyKey,
		EncryptionPublicKey: encryptionPublicKey,
		Secret:              &secret,
	}
}

func (id IxoDid) Equals(other IxoDid) bool {
	return id.Did == other.Did &&
		id.VerifyKey == other.VerifyKey &&
		id.EncryptionPublicKey == other.EncryptionPublicKey &&
		id.Secret.Equals(*other.Secret)
}

func VerifyKeyToAddr(verifyKey string) sdk.AccAddress {
	var pubKey ed25519.PubKey
	pubKey.Key = base58.Decode(verifyKey)
	// var pkSECP secp256k1.PubKey
	// pkSECP.Key = base58.Decode(verifyKey)

	return sdk.AccAddress(pubKey.Address())
}

func (id IxoDid) Address() sdk.AccAddress {
	return VerifyKeyToAddr(id.VerifyKey)
}

func GenerateMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

func FromMnemonic(mnemonic string) (IxoDid, error) {
	seed := sha256.New()
	seed.Write([]byte(mnemonic))

	var seed32 [32]byte
	copy(seed32[:], seed.Sum(nil)[:32])

	return FromSeed(seed32)
}

func Gen() (IxoDid, error) {
	var seed [32]byte
	_, err := io.ReadFull(cryptoRand.Reader, seed[:])
	if err != nil {
		return IxoDid{}, err
	}
	return FromSeed(seed)
}

func FromSeed(seed [32]byte) (IxoDid, error) {
	publicKeyBytes, privateKeyBytes, err := ed25519Local.GenerateKey(bytes.NewReader(seed[0:32]))
	if err != nil {
		return IxoDid{}, err
	}
	publicKey := []byte(publicKeyBytes)
	privateKey := []byte(privateKeyBytes)

	signKey := base58.Encode(privateKey[:32])
	keyPairPublicKey, keyPairPrivateKey, err := naclBox.GenerateKey(bytes.NewReader(privateKey[:]))
	if err != nil {
		return IxoDid{}, err
	}

	return IxoDid{
		Did:                 DidPrefix + base58.Encode(publicKey[:16]),
		VerifyKey:           base58.Encode(publicKey),
		EncryptionPublicKey: base58.Encode(keyPairPublicKey[:]),
		Secret: &Secret{
			Seed:                 hex.EncodeToString(seed[0:32]),
			SignKey:              signKey,
			EncryptionPrivateKey: base58.Encode(keyPairPrivateKey[:]),
		},
	}, nil
}

func (id IxoDid) SignMessage(msg []byte) ([]byte, error) {
	var privateKey ed25519.PrivKey
	privateKey.Key = append(base58.Decode(id.Secret.SignKey), base58.Decode(id.VerifyKey)...)

	return privateKey.Sign(msg)
}

func (id IxoDid) VerifySignedMessage(msg []byte, sig []byte) bool {
	var publicKey ed25519.PubKey
	publicKey.Key = base58.Decode(id.VerifyKey)

	return publicKey.VerifySignature(msg, sig)
}
