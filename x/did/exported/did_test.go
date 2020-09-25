package exported

import (
	"crypto/sha256"
	"encoding/base64"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	validMnemonic = "" +
		"basket mechanic myself capable shoe then " +
		"home magic cream edge seminar artefact"
	validIxoDid = IxoDid{
		Did:                 DidPrefix + "CYCc2xaJKrp8Yt947Nc6jd",
		VerifyKey:           "7HjjYKd4SoBv26MqXp1SzmvDiouQxarBZ2ryscZLK22x",
		EncryptionPublicKey: "FaE44kz98vbKdKh3YWzhe7PTPZ8YsbpDFpdwveGjDgv6",
		Secret: Secret{
			Seed:                 "29a58bc799e8ce6a0ee87cc1e42107fc93e9d904f345501fcd92c20172b2603a",
			SignKey:              "3oa8GeqqCYpmdXa1TW8Q8CtU1M1PELhkTnNYbhcTamBX",
			EncryptionPrivateKey: "3oa8GeqqCYpmdXa1TW8Q8CtU1M1PELhkTnNYbhcTamBX",
		},
	}
	// Note: validIxoDid deduced from validMnemonic
)

const (
	Bech32MainPrefix     = "ixo"
	Bech32PrefixAccAddr  = Bech32MainPrefix
	Bech32PrefixAccPub   = Bech32MainPrefix + sdk.PrefixPublic
	Bech32PrefixValAddr  = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator
	Bech32PrefixValPub   = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	Bech32PrefixConsAddr = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus
	Bech32PrefixConsPub  = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

func TestVerifyKeyToAddr(t *testing.T) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	config.Seal()

	verifyKey := validIxoDid.VerifyKey
	expectedAddress := "ixo19h3lqj50uhzdrv8mkafnp55nqmz4ghc2sd3m48"
	actualAddress := VerifyKeyToAddr(verifyKey).String()

	require.Equal(t, expectedAddress, actualAddress)
}

func TestUnprefixedDid(t *testing.T) {
	expectedResult := "U7GK8p8rVhJMKhBVRCJJ8c"

	did1 := "did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"
	did2 := "did:sov:U7GK8p8rVhJMKhBVRCJJ8c"

	result1 := UnprefixedDid(did1)
	result2 := UnprefixedDid(did2)

	require.Equal(t, expectedResult, result1)
	require.Equal(t, expectedResult, result2)
}

func TestDidFromPubKey(t *testing.T) {
	expectedDid := "U7GK8p8rVhJMKhBVRCJJ8c"
	// equivalent to UnprefixedDid("U7GK8p8rVhJMKhBVRCJJ8c")

	pubKey := "FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW"
	did := UnprefixedDidFromPubKey(pubKey)

	require.Equal(t, expectedDid, did)
}

func TestIxoDid_Equals(t *testing.T) {
	ixoDid := NewIxoDid(
		validIxoDid.Did, validIxoDid.VerifyKey,
		validIxoDid.EncryptionPublicKey,
		NewSecret(
			validIxoDid.Secret.Seed, validIxoDid.Secret.SignKey,
			validIxoDid.Secret.EncryptionPrivateKey))
	require.True(t, ixoDid.Equals(validIxoDid))

	var ixoDid2 IxoDid

	ixoDid2 = ixoDid
	ixoDid2.Did += "_"
	require.False(t, ixoDid2.Equals(validIxoDid))

	ixoDid2 = ixoDid
	ixoDid2.VerifyKey += "_"
	require.False(t, ixoDid2.Equals(validIxoDid))

	ixoDid2 = ixoDid
	ixoDid2.EncryptionPublicKey += "_"
	require.False(t, ixoDid2.Equals(validIxoDid))

	ixoDid2 = ixoDid
	ixoDid2.Secret.Seed += "_"
	require.False(t, ixoDid2.Equals(validIxoDid))

	ixoDid2 = ixoDid
	ixoDid2.Secret.SignKey += "_"
	require.False(t, ixoDid2.Equals(validIxoDid))

	ixoDid2 = ixoDid
	ixoDid2.Secret.EncryptionPrivateKey += "_"
	require.False(t, ixoDid2.Equals(validIxoDid))
}

func TestSecret_Equals(t *testing.T) {
	validSecret := validIxoDid.Secret
	secret := NewSecret(
		validSecret.Seed, validSecret.SignKey,
		validSecret.EncryptionPrivateKey)
	require.True(t, secret.Equals(validSecret))

	var secret2 Secret

	secret2 = secret
	secret2.Seed += "_"
	require.False(t, secret2.Equals(validSecret))

	secret2 = secret
	secret2.SignKey += "_"
	require.False(t, secret2.Equals(validSecret))

	secret2 = secret
	secret2.EncryptionPrivateKey += "_"
	require.False(t, secret2.Equals(validSecret))
}

func TestGenerateMnemonic(t *testing.T) {
	mnemonic, err := GenerateMnemonic()
	require.Nil(t, err)
	require.True(t, bip39.IsMnemonicValid(mnemonic))
}

func TestFromMnemonic(t *testing.T) {
	ixoDid, err := FromMnemonic(validMnemonic)
	require.Nil(t, err)
	require.Equal(t, validIxoDid, ixoDid)
}

func TestGen(t *testing.T) {
	ixoDid, err := Gen()
	require.Nil(t, err)
	require.NotNil(t, ixoDid)
}

func TestFromSeed(t *testing.T) {
	seed := sha256.New()
	seed.Write([]byte(validMnemonic))
	var seed32 [32]byte
	copy(seed32[:], seed.Sum(nil)[:32])

	ixoDid, err := FromSeed(seed32)
	require.Nil(t, err)
	require.Equal(t, validIxoDid, ixoDid)
}

func TestSignAndVerify(t *testing.T) {
	bz1 := []byte("abcdefghijklmnopqrstuvwxyz1234567890")  // "correct" msg
	bz2 := []byte("abcdefghijklmnopqrstuvwxyz1234567890_") // "incorrect" msg

	sig1, err := validIxoDid.SignMessage(bz1) // "correct" signature
	require.Nil(t, err)
	sig2 := append(sig1, byte(0)) // "incorrect" signature

	// Check signature 1
	expectedSig1B64 := "vzWqSM2JMDMZYs8T4NQhXq6oZXbToqJIAflD27KFgd8ZF8khmIYQwaVrAnVxQK9oIwtn6q0xELfv5AA8Ggd4BA=="
	actualSig1B64 := base64.StdEncoding.EncodeToString(sig1)
	require.Equal(t, expectedSig1B64, actualSig1B64)

	// Check signature 2
	expectedSig2B64 := "vzWqSM2JMDMZYs8T4NQhXq6oZXbToqJIAflD27KFgd8ZF8khmIYQwaVrAnVxQK9oIwtn6q0xELfv5AA8Ggd4BAA="
	actualSig2B64 := base64.StdEncoding.EncodeToString(sig2)
	require.Equal(t, expectedSig2B64, actualSig2B64)

	// Correct signature and correct/incorrect msg
	require.True(t, validIxoDid.VerifySignedMessage(bz1, sig1))
	require.False(t, validIxoDid.VerifySignedMessage(bz2, sig1))

	// Incorrect signature and correct/incorrect msg
	require.False(t, validIxoDid.VerifySignedMessage(bz1, sig2))
	require.False(t, validIxoDid.VerifySignedMessage(bz2, sig2))
}
