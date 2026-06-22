package v8

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"

	"github.com/ixofoundation/ixo-blockchain/v7/app/keepers"
)

// icaHostAllowMessages is the explicit allow-list of message type URLs that an
// Interchain-Accounts host account may execute on ixo.
//
// SECURITY (2026-06 audit): the previous value was `["*"]` (allow-all). Because
// ICA-host-dispatched messages BYPASS the ante handler — including the IID
// controller check — allow-all let any IBC counterparty invoke ixo identity /
// asset messages while bypassing ante authorization. We restrict it to standard
// Cosmos SDK operations, whose authorization is bound to the (ICA-forced)
// signer field and therefore safe.
//
// Every ixo custom message (and authz.MsgExec) is intentionally OMITTED below —
// see the commented block. Re-add an entry ONLY after confirming the message's
// keeper handler re-verifies authorization against its proto-signer field (so
// it is safe without the ante).
var icaHostAllowMessages = []string{
	// --- bank ---
	"/cosmos.bank.v1beta1.MsgSend",
	"/cosmos.bank.v1beta1.MsgMultiSend",
	// --- staking ---
	"/cosmos.staking.v1beta1.MsgDelegate",
	"/cosmos.staking.v1beta1.MsgUndelegate",
	"/cosmos.staking.v1beta1.MsgBeginRedelegate",
	"/cosmos.staking.v1beta1.MsgCancelUnbondingDelegation",
	// --- distribution ---
	"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
	"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission",
	"/cosmos.distribution.v1beta1.MsgSetWithdrawAddress",
	"/cosmos.distribution.v1beta1.MsgFundCommunityPool",
	// --- gov ---
	"/cosmos.gov.v1.MsgVote",
	"/cosmos.gov.v1.MsgVoteWeighted",
	"/cosmos.gov.v1.MsgDeposit",
	"/cosmos.gov.v1.MsgSubmitProposal",
	"/cosmos.gov.v1beta1.MsgVote",
	"/cosmos.gov.v1beta1.MsgVoteWeighted",
	"/cosmos.gov.v1beta1.MsgDeposit",
	// --- slashing (validator ops) ---
	"/cosmos.slashing.v1beta1.MsgUnjail",
	// --- feegrant ---
	"/cosmos.feegrant.v1beta1.MsgGrantAllowance",
	"/cosmos.feegrant.v1beta1.MsgRevokeAllowance",
	// --- authz grants (NOT MsgExec — see below) ---
	"/cosmos.authz.v1beta1.MsgGrant",
	"/cosmos.authz.v1beta1.MsgRevoke",
	// --- IBC transfer ---
	"/ibc.applications.transfer.v1.MsgTransfer",

	// --- INTENTIONALLY OMITTED until each is proven safe without the ante ---
	// authz.MsgExec is omitted: it wraps inner messages that the ante never
	// re-inspects on the ICA route, which would re-open the nested-bypass class.
	//   "/cosmos.authz.v1beta1.MsgExec",
	//
	// ixo identity / asset modules are omitted. Their keeper handlers must be
	// confirmed to re-verify the signer→DID/owner binding in-keeper (not via the
	// ante) before any are enabled here:
	//   "/ixo.iid.v1beta1.Msg...",          // iid: keeper-protected, but review ICA exposure first
	//   "/ixo.entity.v1beta1.Msg...",       // entity: fixed in v8, but keep off ICA pending review
	//   "/ixo.claims.v1beta1.Msg...",       // claims: keeper-protected, review ICA exposure first
	//   "/ixo.token.v1beta1.Msg...",        // token
	//   "/ixo.bonds.v1beta1.Msg...",        // bonds: DISABLED
	//   "/ixo.smartaccount.v1beta1.Msg...", // smart-account
}

// CreateUpgradeHandler builds the v8 emergency-security upgrade handler.
//
// Most of the v8 hardening (bonds disable, IID ante MsgExec recursion, entity
// keeper signer checks, token batch checks) is compiled into the binary every
// validator runs and needs no state migration. The one state change here is
// tightening the ICA-host AllowMessages param from allow-all to an explicit,
// ante-safe allow-list (see icaHostAllowMessages).
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	appKeepers keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(c context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(c)
		ctx.Logger().Info("🚀 executing Ixo " + UpgradeName + " emergency security upgrade (bonds disabled) 🚀")

		// -------------------------------------------------
		// Run migrations before applying any other state changes.
		// NOTE: DO NOT PUT ANY STATE CHANGES BEFORE RunMigrations().
		// -------------------------------------------------
		ctx.Logger().Info("Run module migrations")
		newVM, err := mm.RunMigrations(c, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		// -------------------------------------------------
		// Restrict ICA-host AllowMessages (was allow-all "*").
		// -------------------------------------------------
		ctx.Logger().Info("Restrict ICA host AllowMessages to an ante-safe allow-list")
		appKeepers.ICAHostKeeper.SetParams(ctx, icahosttypes.Params{
			HostEnabled:   true,
			AllowMessages: icaHostAllowMessages,
		})

		return newVM, nil
	}
}
