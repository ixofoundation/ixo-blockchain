package ante

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/authz"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/v8/x/iid/keeper"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
)

// maxMsgExecNestingDepth caps how deep VerifyIidControllersAgainstSignature
// recurses into nested authz.MsgExec wrappers. Cosmos allows a MsgExec to wrap
// further MsgExecs; without a bound a crafted tx could nest them to force
// pathological recursion. A depth of 6 is far beyond any legitimate use.
const maxMsgExecNestingDepth = 6

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

	// Collect every message in the transaction, unwrapping authz.MsgExec
	// wrappers so that IID-controlled messages nested inside a MsgExec are
	// subject to the same controller check as top-level messages.
	msgs, err := flattenMsgsForIidCheck(tx.GetMsgs(), 0)
	if err != nil {
		return err
	}

	// For each message in the transaction check if this is an IidTxMsg, and if so validate
	for _, msg := range msgs {
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

// flattenMsgsForIidCheck returns msgs with any authz.MsgExec wrappers replaced
// by the messages they carry, recursively, up to maxMsgExecNestingDepth. This
// ensures IID-controlled messages hidden inside one or more authz.MsgExec
// layers are still subject to the controller check.
//
// For a message nested inside a MsgExec, cdc.GetMsgV1Signers (called by the
// caller) returns the message's declared signer field — the authz *granter*,
// the principal on whose behalf the action runs. Requiring that principal to
// control the DID matches the top-level semantics; authz separately enforces
// that the grantee who actually signed the tx holds a grant from the granter.
func flattenMsgsForIidCheck(msgs []sdk.Msg, depth int) ([]sdk.Msg, error) {
	if depth > maxMsgExecNestingDepth {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "authz.MsgExec nesting too deep")
	}
	out := make([]sdk.Msg, 0, len(msgs))
	for _, msg := range msgs {
		exec, ok := msg.(*authz.MsgExec)
		if !ok {
			out = append(out, msg)
			continue
		}
		inner, err := exec.GetMessages()
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to unpack authz.MsgExec messages")
		}
		flat, err := flattenMsgsForIidCheck(inner, depth+1)
		if err != nil {
			return nil, err
		}
		out = append(out, flat...)
	}
	return out, nil
}

// IidTxMsg is an interface that is implemented by all the messages that
// can be used to control or authorize an IID
type IidTxMsg interface {
	GetIidController() iidtypes.DIDFragment
}
