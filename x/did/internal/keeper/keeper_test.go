package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

func TestKeeper(t *testing.T) {
	ctx, k, cdc := CreateTestInput()
	cdc.RegisterInterface((*ixo.DidDoc)(nil), nil)
	_, err := k.GetDidDoc(ctx, types.EmptyDid)
	require.NotNil(t, err)

	err = k.SetDidDoc(ctx, &types.ValidDidDoc)
	require.Nil(t, err)

	_, err = k.GetDidDoc(ctx, types.ValidDidDoc.GetDid())
	require.Nil(t, err)
}
