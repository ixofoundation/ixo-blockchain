package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

func GetIidControllers(tx signing.SigVerifiableTx) []IidTxMsg {
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

func VerifyIidControllersAgainstSigniture(tx signing.SigVerifiableTx, ctx sdk.Context, iidKeeper iidkeeper.Keeper) error {
	pubKeys, err := tx.GetPubKeys()
	if err != nil {
		return sdkerrors.Wrap(err, "TX must be signed with pubkey")
	}

	iidHasPubKey := false
	controllers := GetIidControllers(tx)

	for _, iidMsg := range controllers {
		iid := iidMsg.GetIidController().Did()
		iidDoc, exists := iidKeeper.GetDidDocument(ctx, []byte(iid))

		if !exists {
			return sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "did document %s not found", iid)
		}

		for _, pk := range pubKeys {
			if iidHasPubKey = iidDoc.HasPublicKey(pk); iidHasPubKey {
				break
			}
		}
	}

	if !iidHasPubKey && len(controllers) > 0 {
		return sdkerrors.Wrap(iidtypes.ErrDidPubKeyMismatch, "one of the dids provided mismatch with signed pubkey")
	}

	return nil
}

type IidTxMsg interface {
	GetIidController() iidtypes.DIDFragment
}
