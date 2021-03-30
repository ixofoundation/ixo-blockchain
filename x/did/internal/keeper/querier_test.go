package keeper

//import (
//	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
//	"testing"
//
//	"github.com/stretchr/testify/require"
//	abci "github.com/tendermint/tendermint/abci/types"
//
//	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
//)
//
//func TestQueryDidDocs(t *testing.T) {
//	ctx, k, cdc := CreateTestInput()
//	cdc.RegisterInterface((*exported.DidDoc)(nil), nil)
//	//err := k.SetDidDoc(ctx, &types.ValidDidDoc)
//	err := k.SetDidDoc(ctx, types.ValidDidDoc)
//	require.Nil(t, err)
//
//	query := abci.RequestQuery{
//		Path: "",
//		Data: []byte{},
//	}
//
//	querier := NewQuerier(k, cdc)
//	res, err := querier(ctx, []string{"queryDidDoc", types.ValidDidDoc.Did}, query)
//	require.Nil(t, err)
//
//	var a types.BaseDidDoc
//	if err := cdc.UnmarshalJSON(res, &a); err != nil {
//		t.Log(err)
//	}
//	_, _ = cdc.MarshalJSONIndent(a, "", " ")
//	resD, err := querier(ctx, []string{"queryAllDidDocs"}, query)
//	require.Nil(t, err)
//
//	var b []types.BaseDidDoc
//	if err := cdc.UnmarshalJSON(resD, &b); err != nil {
//		t.Log(err)
//	}
//
//	_, _ = cdc.MarshalJSONIndent(b, "", " ")
//
//}
