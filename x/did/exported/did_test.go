package exported

import (
	"crypto/sha256"
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

	// Correct signature and correct/incorrect msg
	require.True(t, validIxoDid.VerifySignedMessage(bz1, sig1))
	require.False(t, validIxoDid.VerifySignedMessage(bz2, sig1))

	// Incorrect signature and correct/incorrect msg
	require.False(t, validIxoDid.VerifySignedMessage(bz1, sig2))
	require.False(t, validIxoDid.VerifySignedMessage(bz2, sig2))
}
