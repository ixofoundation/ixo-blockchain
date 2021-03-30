package keeper

// TODO tests now generate app (simapp.Setup) instead of using CreateTestInput()

//import (
//	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
//	"testing"
//
//	"github.com/stretchr/testify/require"
//
//	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
//)
//
//func TestKeeper(t *testing.T) {
//	ctx, k, cdc := CreateTestInput()
//	cdc.RegisterInterface((*exported.DidDoc)(nil), nil)
//	_, err := k.GetDidDoc(ctx, types.EmptyDid)
//	require.NotNil(t, err)
//
//	//err = k.SetDidDoc(ctx, &types.ValidDidDoc)
//	err = k.SetDidDoc(ctx, types.ValidDidDoc)
//	require.Nil(t, err)
//
//	_, err = k.GetDidDoc(ctx, types.ValidDidDoc.GetDid())
//	require.Nil(t, err)
//}