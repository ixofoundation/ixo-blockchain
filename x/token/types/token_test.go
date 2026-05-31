package types_test

import (
	"testing"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v7/x/token/types"
)

func mkAddr() string {
	return sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address()).String()
}

func TestIsValidToken(t *testing.T) {
	good := &types.Token{
		Minter:          mkAddr(),
		ContractAddress: mkAddr(),
		Class:           "did:ixo:abc",
		Name:            "alpha",
	}
	require.True(t, types.IsValidToken(good))

	require.False(t, types.IsValidToken(nil))

	bad := *good
	bad.Name = ""
	require.False(t, types.IsValidToken(&bad), "empty name fails")

	bad = *good
	bad.ContractAddress = "not-bech32"
	require.False(t, types.IsValidToken(&bad), "non-bech32 contract address fails")

	bad = *good
	bad.Class = "not-a-did"
	require.False(t, types.IsValidToken(&bad), "non-DID class fails")
}

func TestIsValidTokenProperties(t *testing.T) {
	good := &types.TokenProperties{Id: "x", Name: "y"}
	require.True(t, types.IsValidTokenProperties(good))
	require.False(t, types.IsValidTokenProperties(nil))
	require.False(t, types.IsValidTokenProperties(&types.TokenProperties{Name: "y"}))
	require.False(t, types.IsValidTokenProperties(&types.TokenProperties{Id: "x"}))
}
