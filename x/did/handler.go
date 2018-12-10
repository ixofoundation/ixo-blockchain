package did

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const didPrefix = "did:sov:"

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case AddDidMsg:
			return handleAddDidDocMsg(ctx, k, msg)
		case AddCredentialMsg:
			return handleAddCredentialMsg(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleAddDidDocMsg(ctx sdk.Context, k Keeper, msg AddDidMsg) sdk.Result {
	newDidDoc := msg.DidDoc

	if len(newDidDoc.Credentials) > 0 {
		return sdk.ErrUnknownRequest("Cannot add a new DID with existing Credentials").Result()
	}

	didDoc, err := k.AddDidDoc(ctx, newDidDoc)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: k.encodeDid(didDoc),
	}
}

func handleAddCredentialMsg(ctx sdk.Context, k Keeper, msg AddCredentialMsg) sdk.Result {
	didDoc, err := k.AddCredential(ctx, msg.DidCredential.Claim.Id, msg.DidCredential)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: k.encodeDid(didDoc),
	}
}

func PrefixDid(did ixo.Did) ixo.Did {
	didString := string(did)
	if strings.HasPrefix(didString, didPrefix) {
		return did
	} else {
		newDid := didPrefix + didString
		return ixo.Did(newDid)
	}
}
