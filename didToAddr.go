package main

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func byte32(s []byte) [32]byte {
	var a [32]byte
	for i := 0; i < 32; i += 1 {
		a[i] = s[i]
	}
	return a
}

func main() {
	//bech := sdk.MustGetAccPubKeyBech32()
	//fmt.Println(bech)

	//pubkey := ed25519.GenPrivKey().PubKey()
	//fmt.Println(sdk.AccAddress(key.PubKey().Address()).String())

	_ = sovrin.SovrinDid{}
	publicKey := base58.Decode("FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW")
	did := base58.Encode(publicKey[:16])
	verifyKey := base58.Encode(publicKey)
	fmt.Println(did)
	fmt.Println(verifyKey)

	pubkey := ed25519.PubKeyEd25519(byte32(publicKey))
	mustBech32AccPub := sdk.MustBech32ifyAccPub(pubkey)
	fmt.Println(mustBech32AccPub)
	test, err := sdk.GetAccPubKeyBech32(mustBech32AccPub)
	if err != nil {
		panic(err)
	}
	fmt.Println(test)

	aa := sdk.AccAddress(pubkey.Address().Bytes()).String()
	fmt.Println(aa)

	addr, err := sdk.AccAddressFromBech32(mustBech32AccPub)
	if err != nil {
		panic(err)
	}
	fmt.Println(addr)

	//var pub ed25519.PubKeyEd25519
	//pub = publicKey
	//
	//pubkey, err := crypto.UnmarshalPubkey(publicKey)
	//if err != nil {
	//	panic(err)
	//}
	//mustBech32AccPub := sdk.MustBech32ifyAccPub(pubkey)
}
