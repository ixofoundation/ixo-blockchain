package project

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

func GetPubKeyGetter(keeper Keeper, didKeeper did.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) ([32]byte, sdk.Result) {
		// Message must be a ProjectMsg
		projectMsg := msg.(types.ProjectMsg)

		// Get signer PubKey
		var pubKey [32]byte
		if projectMsg.IsNewDid() {
			createProjectMsg := msg.(types.MsgCreateProject)
			copy(pubKey[:], base58.Decode(createProjectMsg.GetPubKey()))
		} else {
			if projectMsg.IsWithdrawal() {
				signerDid := msg.GetSignerDid()
				didDoc, _ := didKeeper.GetDidDoc(ctx, signerDid)
				if didDoc == nil {
					return pubKey, sdk.ErrUnauthorized("signer did not found").Result()
				}
				copy(pubKey[:], base58.Decode(didDoc.GetPubKey()))
			} else {
				projectDid := msg.GetSignerDid()
				projectDoc, err := keeper.GetProjectDoc(ctx, projectDid)
				if err != nil {
					return pubKey, sdk.ErrInternal("project did not found").Result()
				}
				copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
			}
		}
		return pubKey, sdk.Result{}
	}
}
