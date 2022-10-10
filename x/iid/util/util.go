package util

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

func GetAccountForVerificationMethod(ctx sdk.Context, accountKeeper authante.AccountKeeper, iidDoc iidtypes.IidDocument, methodId string) (authtypes.AccountI, error) {

	fmt.Println("failed----------------------------------")
	// var verificationMethod *iidtypes.VerificationMethod

	fmt.Println(methodId)

	fmt.Printf("%+v", iidDoc)

	addr, err := iidDoc.GetVerificationMethodBlockchainAddress(fmt.Sprintf("%s#%s", methodId, methodId))
	if err != nil {
		fmt.Println("failed5----------------------------------", err)

		return nil, err
	}

	return accountKeeper.GetAccount(ctx, addr), nil

	// for _, vm := range iidDoc.GetVerificationMethodBlockchainAddress()//VerificationMethod {
	// 	fmt.Printf("%+v", vm)

	// 	if vm.Id == methodId {
	// 		verificationMethod = vm
	// 		break
	// 	}
	// }

	// fmt.Println("failed2----------------------------------")

	// if verificationMethod == nil {
	// 	return nil, errors.New("iid doc does  not exists")
	// }

	// fmt.Println("failed3----------------------------------")

	// switch m := verificationMethod.GetVerificationMaterial().(type) {
	// case *iidtypes.VerificationMethod_PublicKeyHex:

	// 	pubKey, err := hex.DecodeString(m.PublicKeyHex)
	// 	if err != nil {
	// 		fmt.Println("failed4----------------------------------", err)
	// 		return nil, err

	// 	}

	// 	payerPubKey := ed25519.PubKey{Key: pubKey}
	// 	return accountKeeper.GetAccount(ctx, payerPubKey.Address().Bytes()), nil

	// case *iidtypes.VerificationMethod_BlockchainAccountID:
	// 	addr, err := sdk.AccAddressFromBech32(iidtypes.BlockchainAccountID(m.BlockchainAccountID).GetAddress())
	// 	if err != nil {
	// 		fmt.Println("failed5----------------------------------", err)

	// 		return nil, err
	// 	}

	// 	return accountKeeper.GetAccount(ctx, addr), nil

	// default:

	// 	fmt.Println("failed4----------------------------------")

	// 	return nil, errors.New("iid doc does  not exists")
	// }
}
