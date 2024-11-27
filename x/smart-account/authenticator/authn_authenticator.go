package authenticator

import (
	"crypto/x509"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/ixofoundation/ixo-blockchain/v3/x/smart-account/crypto"
)

// Ensure AuthnVerification implements the Authenticator interface
var _ Authenticator = &AuthnVerification{}

const (
	// AuthnVerificationType is the unique type for this authenticator
	AuthnVerificationType = "AuthnVerification"
	// TODO: do benchmarking to determine more accurately the gas cost for each algorithm
	GasCostVerifyES256 = 1500
	GasCostVerifyRS256 = 2500
)

// AuthnVerification authenticator
type AuthnVerification struct {
	ak     authante.AccountKeeper
	PubKey crypto.AuthnPubKey
}

// NewAuthnVerification creates a new AuthnVerification authenticator
func NewAuthnVerification(ak authante.AccountKeeper) AuthnVerification {
	return AuthnVerification{ak: ak}
}

// Type returns the authenticator's type
func (pva AuthnVerification) Type() string {
	return AuthnVerificationType
}

// StaticGas returns the fixed gas cost for this authenticator
func (pva AuthnVerification) StaticGas() uint64 {
	// Adjust the gas cost as needed
	return 0
}

// Initialize initializes the authenticator with configuration data
func (pva AuthnVerification) Initialize(config []byte) (Authenticator, error) {
	if len(config) == 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "public key data is empty")
	}
	pubKey, err := crypto.UnmarshalAuthnPubKey(config)
	if err != nil || pubKey == nil {
		return nil, errorsmod.Wrap(err, "failed to unmarshal public key")
	}
	pva.PubKey = *pubKey
	return pva, nil
}

// Authenticate verifies the signature using WebAuthn
func (pva AuthnVerification) Authenticate(ctx sdk.Context, request AuthenticationRequest) error {
	// Consume gas for verifying the signature
	var gasCost uint64
	switch pva.PubKey.CoseAlgorithm {
	case -7: // ES256
		gasCost = GasCostVerifyES256
	case -257: // RS256
		gasCost = GasCostVerifyRS256
	default:
		return errorsmod.Wrapf(sdkerrors.ErrInvalidPubKey, "unsupported algorithm: %d", pva.PubKey.CoseAlgorithm)
	}

	ctx.GasMeter().ConsumeGas(gasCost, "authn signature verification")

	if request.Simulate || ctx.IsReCheckTx() {
		return nil
	}
	if pva.PubKey.Key == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "public key not set on authenticator")
	}

	// Verify the signature
	if !pva.PubKey.VerifySignature(request.SignModeTxData.Direct, request.Signature) {
		return errorsmod.Wrapf(
			sdkerrors.ErrUnauthorized,
			"signature verification failed; please verify account number (%d), sequence (%d) and chain-id (%s)",
			request.TxData.AccountNumber,
			request.TxData.AccountSequence,
			request.TxData.ChainID,
		)
	}
	return nil
}

// Track is a no-op for this authenticator
func (pva AuthnVerification) Track(ctx sdk.Context, request AuthenticationRequest) error {
	return nil
}

// ConfirmExecution is a no-op for this authenticator
func (pva AuthnVerification) ConfirmExecution(ctx sdk.Context, request AuthenticationRequest) error {
	return nil
}

// OnAuthenticatorAdded handles the addition of the authenticator to an account
func (pva AuthnVerification) OnAuthenticatorAdded(ctx sdk.Context, account sdk.AccAddress, config []byte, authenticatorId string) error {
	if len(config) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "public key data is empty")
	}
	pubKey, err := crypto.UnmarshalAuthnPubKey(config)
	if err != nil || pubKey == nil {
		return errorsmod.Wrap(err, "failed to unmarshal public key")
	}

	// Validate that key_id (credential id) is not empty
	if len(pubKey.KeyId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "public key id (credential id) is empty")
	}

	// Validate the CoseAlgorithm
	switch pubKey.CoseAlgorithm {
	case -7: // ES256 (ECDSA with SHA-256)
		// Early Rejection: Quickly discard keys that are clearly invalid due to extreme sizes.
		// Validate key length (approximate range for ECDSA public keys)
		if len(pubKey.Key) < 80 || len(pubKey.Key) > 120 {
			return errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "invalid ECDSA public key length")
		}
		_, err := x509.ParsePKIXPublicKey(pubKey.Key)
		if err != nil {
			return errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "failed to parse ECDSA public key")
		}

	case -257: // RS256 (RSASSA-PKCS1-v1_5 with SHA-256)
		// Early Rejection: Quickly discard keys that are clearly invalid due to extreme sizes.
		// Validate key length (RSA keys are typically larger)
		if len(pubKey.Key) < 200 || len(pubKey.Key) > 800 {
			return errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "invalid RSA public key length")
		}
		_, err := x509.ParsePKIXPublicKey(pubKey.Key)
		if err != nil {
			return errorsmod.Wrap(sdkerrors.ErrInvalidPubKey, "failed to parse RSA public key")
		}

	default:
		return errorsmod.Wrapf(sdkerrors.ErrInvalidPubKey, "unsupported algorithm: %d", pubKey.CoseAlgorithm)
	}

	return nil
}

// OnAuthenticatorRemoved handles the removal of the authenticator from an account
func (pva AuthnVerification) OnAuthenticatorRemoved(ctx sdk.Context, account sdk.AccAddress, config []byte, authenticatorId string) error {
	return nil
}
