package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/v3/x/iid/keeper"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v3/x/iid/types"
)

// GetIidTxs returns all the IidTxMsgs from a SigVerifiableTx
// and ignore any messages that are not IidTxMsgs
func GetIidTxs(tx signing.SigVerifiableTx) []IidTxMsg {
	var msgs []IidTxMsg

	for _, txMsg := range tx.GetMsgs() {
		iidMsg, ok := txMsg.(IidTxMsg)
		if !ok {
			continue
		}
		msgs = append(msgs, iidMsg)
	}

	return msgs
}

// VerifyIidControllersAgainstSignature verifies that the controllers of the IID for all IidTxMsgs
// in the SigVerifiableTx are authorized to control the IID
func VerifyIidControllersAgainstSignature(tx signing.SigVerifiableTx, ctx sdk.Context, iidKeeper iidkeeper.Keeper) error {
	pubKeys, err := tx.GetPubKeys()
	if err != nil {
		return errorsmod.Wrap(err, "TX must be signed with pubkey")
	}

	iidHasPubKey := false
	iidMsgs := GetIidTxs(tx)

	for _, iidMsg := range iidMsgs {
		iid := iidMsg.GetIidController().Did()
		iidDoc, exists := iidKeeper.GetDidDocument(ctx, []byte(iid))

		if !exists {
			return errorsmod.Wrapf(iidtypes.ErrDidDocumentNotFound, "did document %s not found", iid)
		}

		for _, pk := range pubKeys {
			if iidHasPubKey = iidDoc.HasPublicKey(pk); iidHasPubKey {
				break
			}
		}
	}

	if !iidHasPubKey && len(iidMsgs) > 0 {
		return errorsmod.Wrap(iidtypes.ErrDidPubKeyMismatch, "one of the dids provided mismatch with signed pubkey")
	}

	return nil
}

// IidTxMsg is an interface that is implemented by all the messages that
// can be used to control or authorize an IID
type IidTxMsg interface {
	GetIidController() iidtypes.DIDFragment
}
