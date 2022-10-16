package keeper_test

/*
import (
	"github.com/ixofoundation/ixo-blockchain/x/did/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestQueryDidDocs(t *testing.T) {
	legacyAmino, appl, ctx := CreateTestInput()
	err := appl.DidKeeper.SetDidDoc(ctx, &types.ValidDidDoc)
	require.Nil(t, err)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	querier := keeper.NewQuerier(appl.DidKeeper, legacyAmino)
	res, err := querier(ctx, []string{"queryDidDoc", types.ValidDidDoc.Did}, query)
	require.Nil(t, err)

	var a types.BaseDidDoc
	if err := legacyAmino.UnmarshalJSON(res, &a); err != nil {
		t.Log(err)
	}
	_, _ = legacyAmino.MarshalJSONIndent(a, "", " ")
	resD, err := querier(ctx, []string{"queryAllDidDocs"}, query)
	require.Nil(t, err)

	var b []types.BaseDidDoc
	if err := legacyAmino.UnmarshalJSON(resD, &b); err != nil {
		t.Log(err)
	}

	_, _ = legacyAmino.MarshalJSONIndent(b, "", " ")

}
*/
