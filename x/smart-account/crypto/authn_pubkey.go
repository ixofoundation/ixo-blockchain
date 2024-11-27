package crypto

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/gogoproto/proto"
)

// Ensure AuthnPubKey implements cryptotypes.PubKey
var _ cryptotypes.PubKey = (*AuthnPubKey)(nil)

const (
	keyType = "authn"
)

// Address returns the address of the public key
func (pubKey *AuthnPubKey) Address() cryptotypes.Address {
	return address.Hash(pubKey.Type(), pubKey.Key)
}

// Bytes returns the raw bytes of the public key
func (pubKey *AuthnPubKey) Bytes() []byte {
	return pubKey.Key
}

// String returns a string representation of the public key
func (pubKey *AuthnPubKey) String() string {
	return fmt.Sprintf("PubKeyAuthn{KeyId:%s, CoseAlgorithm:%d, Key:%X}", pubKey.KeyId, pubKey.CoseAlgorithm, pubKey.Key)
}

// Type returns the type of the public key
func (pubKey *AuthnPubKey) Type() string {
	return keyType
}

// Equals checks if two public keys are equal
func (pubKey *AuthnPubKey) Equals(other cryptotypes.PubKey) bool {
	return pubKey.Type() == other.Type() && bytes.Equal(pubKey.Bytes(), other.Bytes())
}

// VerifySignature verifies a signature over the given message
func (pubKey *AuthnPubKey) VerifySignature(msg []byte, sigBytes []byte) bool {
	// The signature bytes should be JSON-encoded WebAuthn signature data
	var sig Signature
	err := json.Unmarshal(sigBytes, &sig)
	if err != nil {
		return false
	}

	clientDataJSON, err := base64.RawURLEncoding.DecodeString(sig.ClientDataJSON)
	if err != nil {
		return false
	}
	signatureBytes, err := base64.RawURLEncoding.DecodeString(sig.Signature)
	if err != nil {
		return false
	}
	authenticatorData, err := base64.RawURLEncoding.DecodeString(sig.AuthenticatorData)
	if err != nil {
		return false
	}

	// check authenticatorData length for early validation
	if len(authenticatorData) < 37 {
		return false
	}

	// Parse clientDataJSON
	var clientData ClientData
	err = json.Unmarshal(clientDataJSON, &clientData)
	if err != nil {
		return false
	}

	challenge, err := base64.RawURLEncoding.DecodeString(clientData.Challenge)
	if err != nil {
		return false
	}

	// Compute the SHA-256 hash of the transaction sign bytes
	expectedChallenge := sha256.Sum256(msg)

	// Verify that the challenge matches the expected value (the message)
	if !bytes.Equal(challenge, expectedChallenge[:]) {
		return false
	}

	clientDataHash := sha256.Sum256(clientDataJSON)
	payload := append(authenticatorData, clientDataHash[:]...)
	hash := sha256.Sum256(payload)

	// Determine the cose algorithm and verify the signature accordingly
	switch pubKey.CoseAlgorithm {
	case -7: // ES256 (ECDSA with SHA-256)
		publicKeyInterface, err := x509.ParsePKIXPublicKey(pubKey.Key)
		if err != nil {
			return false
		}
		publicKey, ok := publicKeyInterface.(*ecdsa.PublicKey)
		if !ok {
			return false
		}
		return ecdsa.VerifyASN1(publicKey, hash[:], signatureBytes)

	case -257: // RS256 (RSASSA-PKCS1-v1_5 with SHA-256)
		publicKeyInterface, err := x509.ParsePKIXPublicKey(pubKey.Key)
		if err != nil {
			return false
		}
		publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
		if !ok {
			return false
		}
		err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signatureBytes)
		return err == nil

	default:
		// Unsupported algorithm
		return false
	}
}

// Signature represents the WebAuthn signature data
type Signature struct {
	AuthenticatorData string `json:"authenticatorData"`
	ClientDataJSON    string `json:"clientDataJSON"`
	Signature         string `json:"signature"`
}

// ClientData represents the WebAuthn client data
type ClientData struct {
	Type      string `json:"type"`
	Challenge string `json:"challenge"`
	Origin    string `json:"origin"`
}

// MarshalAuthnPubKey marshals the AuthnPubKey to bytes
func MarshalAuthnPubKey(pubKey *AuthnPubKey) ([]byte, error) {
	return proto.Marshal(pubKey)
}

// UnmarshalAuthnPubKey unmarshals bytes to an AuthnPubKey
func UnmarshalAuthnPubKey(bz []byte) (*AuthnPubKey, error) {
	var pubKey AuthnPubKey
	err := proto.Unmarshal(bz, &pubKey)
	if err != nil {
		return nil, err
	}
	return &pubKey, nil
}
