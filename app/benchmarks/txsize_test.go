package app

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ixofoundation/ixo-blockchain/app"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// This will fail half the time with the second output being 173
// This is due to secp256k1 signatures not being constant size.
// nolint: vet
func ExampleTxSendSizeSecp256k1() {
	cdc := app.MakeTestEncodingConfig()
	var gas uint64 = 1

	priv1 := secp256k1.GenPrivKeySecp256k1([]byte{0})
	addr1 := sdk.AccAddress(priv1.PubKey().Address())
	priv2 := secp256k1.GenPrivKeySecp256k1([]byte{1})
	addr2 := sdk.AccAddress(priv2.PubKey().Address())
	coins := sdk.Coins{sdk.NewCoin("denom", sdk.NewInt(10))}
	msg1 := banktypes.MsgMultiSend{
		Inputs:  []banktypes.Input{banktypes.NewInput(addr1, coins)},
		Outputs: []banktypes.Output{banktypes.NewOutput(addr2, coins)},
	}
	fee := legacytx.NewStdFee(gas,coins)
	signBytes := legacytx.StdSignBytes("example-chain-ID",
		1, 1, 0, fee, []sdk.Msg{&msg1}, "")
	sig, _ := priv1.Sign(signBytes)
	sigs := legacytx.StdSignature{nil, sig}
	tx := legacytx.NewStdTx([]sdk.Msg{&msg1}, fee, []legacytx.StdSignature{sigs}, "")
	fmt.Println(len(cdc.Amino.MustMarshalBinaryBare([]sdk.Msg{&msg1})))
	fmt.Println(len(cdc.Amino.MustMarshalBinaryBare(tx)))
	// output: 80
	// 169
}

// nolint: vet
func ExampleTxSendSizeEd25519() {
	cdc := app.MakeTestEncodingConfig()
	var gas uint64 = 1

	priv1 := ed25519.GenPrivKeyFromSecret([]byte{0})
	addr1 := sdk.AccAddress(priv1.PubKey().Address())
	priv2 := ed25519.GenPrivKeyFromSecret([]byte{1})
	addr2 := sdk.AccAddress(priv2.PubKey().Address())
	coins := sdk.Coins{sdk.NewCoin("denom", sdk.NewInt(10))}
	msg1 := banktypes.MsgMultiSend{
		Inputs:  []banktypes.Input{banktypes.NewInput(addr1, coins)},
		Outputs: []banktypes.Output{banktypes.NewOutput(addr2, coins)},
	}
	fee := legacytx.NewStdFee(gas, coins)
	signBytes := legacytx.StdSignBytes("example-chain-ID",
		1, 1, 0, fee, []sdk.Msg{&msg1}, "")
	sig, _ := priv1.Sign(signBytes)
	sigs := []legacytx.StdSignature{{nil, sig}}
	tx := legacytx.NewStdTx([]sdk.Msg{&msg1}, fee, sigs, "")
	fmt.Println(len(cdc.Amino.MustMarshalBinaryBare([]sdk.Msg{&msg1})))
	fmt.Println(len(cdc.Amino.MustMarshalBinaryBare(tx)))
	// output: 80
	// 169
}
