package keeper

import (
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper(t *testing.T) {
	legacyAmino, appl, ctx := CreateTestInput()
	legacyAmino.RegisterInterface((*exported.DidDoc)(nil), nil)
	_, err := appl.DidKeeper.GetDidDoc(ctx, types.EmptyDid)
	require.NotNil(t, err)

	err = appl.DidKeeper.SetDidDoc(ctx, &types.ValidDidDoc)
	require.Nil(t, err)

	_, err = appl.DidKeeper.GetDidDoc(ctx, types.ValidDidDoc.GetDid())
	require.Nil(t, err)
}