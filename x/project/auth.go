package project

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

func GetPubKeyGetter(keeper Keeper, didKeeper did.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) ([32]byte, sdk.Result) {

		// Get signer PubKey
		var pubKey [32]byte
		switch msg := msg.(type) {
		case MsgCreateProject:
			copy(pubKey[:], base58.Decode(msg.GetPubKey()))
		case MsgUpdateProjectStatus:
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.ProjectDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("project did not found").Result()
			}
			copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
		case MsgCreateAgent:
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.ProjectDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("project did not found").Result()
			}
			copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
		case MsgUpdateAgent:
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.ProjectDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("project did not found").Result()
			}
			copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
		case MsgCreateClaim:
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.ProjectDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("project did not found").Result()
			}
			copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
		case MsgCreateEvaluation:
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.ProjectDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("project did not found").Result()
			}
			copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
		case MsgWithdrawFunds:
			didDoc, _ := didKeeper.GetDidDoc(ctx, msg.Data.RecipientDid)
			if didDoc == nil {
				return pubKey, sdk.ErrUnauthorized("signer did not found").Result()
			}
			copy(pubKey[:], base58.Decode(didDoc.GetPubKey()))
		default:
			return pubKey, sdk.ErrUnknownRequest("No match for message type.").Result()
		}
		return pubKey, sdk.Result{}
	}
}
