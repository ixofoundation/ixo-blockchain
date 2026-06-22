package ante

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"

	"github.com/ixofoundation/ixo-blockchain/v7/x/bonds/types"
)

// bondsMsgTypePrefix is the proto type-URL prefix shared by every x/bonds
// message (e.g. "/ixo.bonds.v1beta1.MsgMakeOutcomePayment"). Matching on the
// prefix keeps this guard correct even if new bonds messages are added.
const bondsMsgTypePrefix = "/ixo.bonds."

// maxMsgExecNestingDepth caps how deep we recurse into nested authz.MsgExec
// wrappers. Cosmos allows MsgExec to wrap further MsgExecs; a crafted tx could
// nest them to force pathological recursion. A depth of 6 is far beyond any
// legitimate use.
const maxMsgExecNestingDepth = 6

// RejectBondsMsgDecorator rejects any transaction carrying an x/bonds message —
// at the top level or nested inside any authz.MsgExec — with a clear
// ErrBondsModuleDisabled error, before it reaches the message router.
//
// This is the front-line layer of the v8 emergency disable of the bonds module.
// It is intentionally redundant with the
// disabled msg server registered in x/bonds/module.go: the disabled msg server
// is the authoritative backstop for routes that bypass the ante (CosmWasm
// stargate, ICA-host), while this decorator gives normal and authz transactions
// an early, explicit rejection.
type RejectBondsMsgDecorator struct{}

func NewRejectBondsMsgDecorator() RejectBondsMsgDecorator {
	return RejectBondsMsgDecorator{}
}

func (d RejectBondsMsgDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if err := assertNoBondsMsg(tx.GetMsgs(), 0); err != nil {
		return ctx, err
	}
	return next(ctx, tx, simulate)
}

// assertNoBondsMsg returns ErrBondsModuleDisabled if any message in msgs is an
// x/bonds message, recursing through authz.MsgExec wrappers up to
// maxMsgExecNestingDepth.
func assertNoBondsMsg(msgs []sdk.Msg, depth int) error {
	if depth > maxMsgExecNestingDepth {
		return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "authz.MsgExec nesting too deep")
	}
	for _, msg := range msgs {
		if exec, ok := msg.(*authz.MsgExec); ok {
			inner, err := exec.GetMessages()
			if err != nil {
				return errorsmod.Wrap(err, "failed to unpack authz.MsgExec messages")
			}
			if err := assertNoBondsMsg(inner, depth+1); err != nil {
				return err
			}
			continue
		}
		if strings.HasPrefix(sdk.MsgTypeURL(msg), bondsMsgTypePrefix) {
			return errorsmod.Wrapf(types.ErrBondsModuleDisabled,
				"message %s rejected: the bonds module is disabled", sdk.MsgTypeURL(msg))
		}
	}
	return nil
}
