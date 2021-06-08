package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/app"
	"github.com/ixofoundation/ixo-blockchain/cmd"
	"github.com/ixofoundation/ixo-blockchain/x/did/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)


func CreateTestInput() (*codec.LegacyAmino, *app.IxoApp, sdk.Context) {
	appl := cmd.Setup(false)
	ctx := appl.BaseApp.NewContext(false, tmproto.Header{})

	appl.DidKeeper = keeper.NewKeeper(appl.AppCodec(), appl.GetKey(types.StoreKey))

	return appl.LegacyAmino(), appl, ctx
}