package ante_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"

	bondsante "github.com/ixofoundation/ixo-blockchain/v7/x/bonds/ante"
	bondstypes "github.com/ixofoundation/ixo-blockchain/v7/x/bonds/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v7/x/iid/types"
)

// stubTx is a minimal sdk.Tx — RejectBondsMsgDecorator only calls GetMsgs().
type stubTx struct{ msgs []sdk.Msg }

func (t *stubTx) GetMsgs() []sdk.Msg                               { return t.msgs }
func (t *stubTx) GetMsgsV2() ([]protoreflect.ProtoMessage, error) { return nil, nil }

func run(t *testing.T, msgs []sdk.Msg) error {
	t.Helper()
	dec := bondsante.NewRejectBondsMsgDecorator()
	called := false
	next := func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
		called = true
		return ctx, nil
	}
	_, err := dec.AnteHandle(sdk.Context{}, &stubTx{msgs: msgs}, false, next)
	if err == nil {
		require.True(t, called, "next must be called when tx is allowed")
	} else {
		require.False(t, called, "next must NOT be called when tx is rejected")
	}
	return err
}

func bondsMsg() sdk.Msg {
	return &bondstypes.MsgMakeOutcomePayment{
		SenderDid:     iidtypes.DIDFragment("did:ixo:someone#v1"),
		Amount:        math.NewInt(1),
		BondDid:       "did:ixo:bond-x",
		SenderAddress: "ixo1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	}
}

func nonBondsMsg() sdk.Msg {
	return &banktypes.MsgSend{FromAddress: "a", ToAddress: "b"}
}

func TestRejectBonds_TopLevel(t *testing.T) {
	err := run(t, []sdk.Msg{bondsMsg()})
	require.ErrorIs(t, err, bondstypes.ErrBondsModuleDisabled)
}

func TestRejectBonds_InsideMsgExec(t *testing.T) {
	exec := authz.NewMsgExec(sdk.AccAddress("grantee"), []sdk.Msg{bondsMsg()})
	err := run(t, []sdk.Msg{&exec})
	require.ErrorIs(t, err, bondstypes.ErrBondsModuleDisabled)
}

func TestRejectBonds_NestedMsgExec(t *testing.T) {
	inner := authz.NewMsgExec(sdk.AccAddress("g1"), []sdk.Msg{bondsMsg()})
	outer := authz.NewMsgExec(sdk.AccAddress("g2"), []sdk.Msg{&inner})
	err := run(t, []sdk.Msg{&outer})
	require.ErrorIs(t, err, bondstypes.ErrBondsModuleDisabled)
}

func TestRejectBonds_NonBondsAllowed(t *testing.T) {
	err := run(t, []sdk.Msg{nonBondsMsg()})
	require.NoError(t, err)
}

func TestRejectBonds_NonBondsInsideMsgExecAllowed(t *testing.T) {
	exec := authz.NewMsgExec(sdk.AccAddress("grantee"), []sdk.Msg{nonBondsMsg()})
	err := run(t, []sdk.Msg{&exec})
	require.NoError(t, err)
}
