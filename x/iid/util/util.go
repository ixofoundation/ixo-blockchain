package util

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v3/x/iid/types"
)

func GetAccountForVerificationMethod(ctx sdk.Context, accountKeeper authante.AccountKeeper, iidDoc iidtypes.IidDocument, methodId string) (authtypes.AccountI, error) {
	addr, err := iidDoc.GetVerificationMethodBlockchainAddress(iidDoc.Id)
	if err != nil {
		return nil, err
	}

	return accountKeeper.GetAccount(ctx, addr), nil

	// for _, vm := range iidDoc.GetVerificationMethodBlockchainAddress()//VerificationMethod {
	// 	if vm.Id == methodId {
	// 		verificationMethod = vm
	// 		break
	// 	}
	// }

	// if verificationMethod == nil {
	// 	return nil, errors.New("iid doc does  not exists")
	// }

	// switch m := verificationMethod.GetVerificationMaterial().(type) {
	// case *iidtypes.VerificationMethod_PublicKeyHex:

	// 	pubKey, err := hex.DecodeString(m.PublicKeyHex)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	payerPubKey := ed25519.PubKey{Key: pubKey}
	// 	return accountKeeper.GetAccount(ctx, payerPubKey.Address().Bytes()), nil

	// case *iidtypes.VerificationMethod_BlockchainAccountID:
	// 	addr, err := sdk.AccAddressFromBech32(iidtypes.BlockchainAccountID(m.BlockchainAccountID).GetAddress())
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return accountKeeper.GetAccount(ctx, addr), nil

	// default:
	// 	return nil, errors.New("iid doc does  not exists")
	// }
}
