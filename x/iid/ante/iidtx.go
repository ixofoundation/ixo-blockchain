package ante

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/v5/x/iid/keeper"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v5/x/iid/types"
)

// VerifyIidControllersAgainstSignature verifies that the controllers of the IID for all IidTxMsgs
// in the SigVerifiableTx are authorized to control the IID. This works for both traditional
// signature-based authentication and smart account authentication methods.
func VerifyIidControllersAgainstSignature(tx signing.SigVerifiableTx, ctx sdk.Context, iidKeeper iidkeeper.Keeper, cdc codec.Codec) error {
	// The constraints for the verification relationships that can be used to act on behalf of the IID
	constraints := []string{
		iidtypes.Authentication,
		iidtypes.AssertionMethod,
		// iidtypes.KeyAgreement,
		// iidtypes.CapabilityInvocation,
		// iidtypes.CapabilityDelegation,
	}

	// For each message in the transaction check if this is an IidTxMsg, and if so validate
	for _, msg := range tx.GetMsgs() {
		iidMsg, ok := msg.(IidTxMsg)
		if !ok {
			continue // Skip non-IID messages
		}

		// Get the IID from the IidTxMsg
		iid := iidMsg.GetIidController().Did()
		iidDoc, exists := iidKeeper.GetDidDocument(ctx, []byte(iid))
		if !exists {
			return errorsmod.Wrapf(iidtypes.ErrDidDocumentNotFound, "did document %s not found", iid)
		}

		// Get the signer for this specific message using the provided codec
		signers, _, err := cdc.GetMsgV1Signers(msg)
		if err != nil {
			return errorsmod.Wrap(err, "failed to get signers")
		}
		// Enforce only one signer per message for IID messages, this is enforced by smart accounts as well as sdk v0.50 onwards
		if len(signers) != 1 {
			return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "IID messages must have exactly one signer")
		}

		// Convert the signer address to BlockchainAccountID format for verification
		signerAddr := sdk.AccAddress(signers[0])

		// Check if the signer has any of the verification relationships in the IID document
		// We check all possible relationships - the specific ones to check can be configured as needed
		if !iidDoc.HasRelationship(iidtypes.NewBlockchainAccountID(signerAddr.String()), constraints...) {
			return errorsmod.Wrapf(iidtypes.ErrUnauthorized,
				"signer account %s not authorized to act on behalf of the did %s",
				signerAddr.String(), iid)
		}
	}

	return nil
}

// IidTxMsg is an interface that is implemented by all the messages that
// can be used to control or authorize an IID
type IidTxMsg interface {
	GetIidController() iidtypes.DIDFragment
}
