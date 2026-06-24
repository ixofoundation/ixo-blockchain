package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v8/x/names/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns a names MsgServer wrapped around k.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{Keeper: k}
}

var _ types.MsgServer = msgServer{}

// CreateNamespace persists a new Namespace. Authority-only.
func (k msgServer) CreateNamespace(goCtx context.Context, msg *types.MsgCreateNamespace) (*types.MsgCreateNamespaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Authority != k.authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidAuthority, "expected %s, got %s", k.authority, msg.Authority)
	}
	if msg.Namespace == nil {
		return nil, errorsmod.Wrap(types.ErrInvalidRequest, "namespace is required")
	}
	if k.HasNamespace(ctx, msg.Namespace.Name) {
		return nil, errorsmod.Wrapf(types.ErrNamespaceExists, "namespace %q already exists", msg.Namespace.Name)
	}
	if err := types.ValidateNamespace(*msg.Namespace); err != nil {
		return nil, err
	}
	k.SetNamespace(ctx, *msg.Namespace)

	if err := ctx.EventManager().EmitTypedEvent(types.NewNamespaceCreatedEvent(*msg.Namespace, msg.Authority)); err != nil {
		return nil, err
	}
	return &types.MsgCreateNamespaceResponse{}, nil
}

// UpdateNamespace replaces an existing Namespace with the supplied
// configuration. Authority-only.
func (k msgServer) UpdateNamespace(goCtx context.Context, msg *types.MsgUpdateNamespace) (*types.MsgUpdateNamespaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Authority != k.authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidAuthority, "expected %s, got %s", k.authority, msg.Authority)
	}
	if msg.Namespace == nil {
		return nil, errorsmod.Wrap(types.ErrInvalidRequest, "namespace is required")
	}
	if !k.HasNamespace(ctx, msg.Namespace.Name) {
		return nil, errorsmod.Wrapf(types.ErrNamespaceNotFound, "namespace %q does not exist", msg.Namespace.Name)
	}
	if err := types.ValidateNamespace(*msg.Namespace); err != nil {
		return nil, err
	}
	k.SetNamespace(ctx, *msg.Namespace)

	if err := ctx.EventManager().EmitTypedEvent(types.NewNamespaceUpdatedEvent(*msg.Namespace, msg.Authority)); err != nil {
		return nil, err
	}
	return &types.MsgUpdateNamespaceResponse{}, nil
}

// RegisterName lets a user self-register a name. The signer must control
// owner_did and the namespace must allow self-registration.
func (k msgServer) RegisterName(goCtx context.Context, msg *types.MsgRegisterName) (*types.MsgRegisterNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ns, err := k.requireNamespace(ctx, msg.Namespace)
	if err != nil {
		return nil, err
	}
	if !ns.AllowSelfRegister {
		return nil, errorsmod.Wrapf(types.ErrSelfRegisterNotAllowed, "namespace %q is registrar-only", ns.Name)
	}
	if err := k.verifyDidController(ctx, msg.OwnerDid, msg.Signer); err != nil {
		return nil, err
	}

	normalized, err := k.normalizeAndCheckAvailable(ctx, ns, msg.Name)
	if err != nil {
		return nil, err
	}

	now := ctx.BlockTime().Unix()
	record := types.NameRecord{
		Namespace:      ns.Name,
		NormalizedName: normalized,
		DisplayName:    msg.Name,
		OwnerDid:       msg.OwnerDid,
		Verified:       false,
		ValidUntil:     0,
		Status:         types.NAME_STATUS_ACTIVE,
		Source:         "self",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	k.SetNameRecord(ctx, record)

	if err := ctx.EventManager().EmitTypedEvent(types.NewNameRegisteredEvent(record, msg.Signer)); err != nil {
		return nil, err
	}
	return &types.MsgRegisterNameResponse{NormalizedName: normalized}, nil
}

// RegisterNameByRegistrar lets a registrar register on behalf of any DID.
// The signer must be a registrar of the namespace.
func (k msgServer) RegisterNameByRegistrar(goCtx context.Context, msg *types.MsgRegisterNameByRegistrar) (*types.MsgRegisterNameByRegistrarResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Defence-in-depth length checks. ValidateBasic catches these for normal
	// txs via the antehandler, but Wasm sub-message dispatch bypasses ante.
	if err := types.ValidateRecordMetadata(msg.EvidenceHash, msg.Source); err != nil {
		return nil, err
	}
	ns, err := k.requireNamespace(ctx, msg.Namespace)
	if err != nil {
		return nil, err
	}
	if !types.HasRegistrar(ns, msg.Registrar) {
		return nil, errorsmod.Wrapf(types.ErrNotRegistrar, "address %s is not a registrar of namespace %q", msg.Registrar, ns.Name)
	}

	// owner_did must reference an existing DID document so the bound DID is
	// real; we do not require the registrar to control it.
	if _, found := k.iidKeeper.GetDidDocument(ctx, []byte(msg.OwnerDid)); !found {
		return nil, errorsmod.Wrapf(types.ErrInvalidDID, "owner DID %q not found", msg.OwnerDid)
	}

	normalized, err := k.normalizeAndCheckAvailable(ctx, ns, msg.Name)
	if err != nil {
		return nil, err
	}

	now := ctx.BlockTime().Unix()
	record := types.NameRecord{
		Namespace:      ns.Name,
		NormalizedName: normalized,
		DisplayName:    msg.Name,
		OwnerDid:       msg.OwnerDid,
		Verified:       msg.Verified,
		ValidUntil:     0,
		Status:         types.NAME_STATUS_ACTIVE,
		VerifiedBy:     msg.Registrar,
		EvidenceHash:   msg.EvidenceHash,
		Source:         msg.Source,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	k.SetNameRecord(ctx, record)

	if err := ctx.EventManager().EmitTypedEvent(types.NewNameRegisteredEvent(record, msg.Registrar)); err != nil {
		return nil, err
	}
	return &types.MsgRegisterNameByRegistrarResponse{NormalizedName: normalized}, nil
}

// UpdateNameByRegistrar updates the verification metadata of an existing
// record. Owner is unchanged; status is unchanged.
func (k msgServer) UpdateNameByRegistrar(goCtx context.Context, msg *types.MsgUpdateNameByRegistrar) (*types.MsgUpdateNameByRegistrarResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Defence-in-depth length checks (Wasm sub-msg bypass).
	if err := types.ValidateRecordMetadata(msg.EvidenceHash, msg.Source); err != nil {
		return nil, err
	}
	ns, err := k.requireNamespace(ctx, msg.Namespace)
	if err != nil {
		return nil, err
	}
	if !types.HasRegistrar(ns, msg.Registrar) {
		return nil, errorsmod.Wrapf(types.ErrNotRegistrar, "address %s is not a registrar of namespace %q", msg.Registrar, ns.Name)
	}

	record, found := k.GetNameRecord(ctx, ns.Name, msg.NormalizedName)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrNameNotFound, "name %q in namespace %q not found", msg.NormalizedName, ns.Name)
	}

	record.Verified = msg.Verified
	record.VerifiedBy = msg.Registrar
	record.EvidenceHash = msg.EvidenceHash
	record.Source = msg.Source
	record.UpdatedAt = ctx.BlockTime().Unix()
	k.SetNameRecord(ctx, record)

	if err := ctx.EventManager().EmitTypedEvent(types.NewNameUpdatedEvent(record, msg.Registrar)); err != nil {
		return nil, err
	}
	return &types.MsgUpdateNameByRegistrarResponse{}, nil
}

// TransferName changes the owner_did of a record. Permitted for the current
// owner; permitted for a registrar when the namespace has
// allow_registrar_override = true.
func (k msgServer) TransferName(goCtx context.Context, msg *types.MsgTransferName) (*types.MsgTransferNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ns, err := k.requireNamespace(ctx, msg.Namespace)
	if err != nil {
		return nil, err
	}
	record, found := k.GetNameRecord(ctx, ns.Name, msg.NormalizedName)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrNameNotFound, "name %q in namespace %q not found", msg.NormalizedName, ns.Name)
	}

	// Authorisation: signer is current owner, or registrar with override.
	ownerOk := k.verifyDidController(ctx, record.OwnerDid, msg.Signer) == nil
	registrarOk := types.HasRegistrar(ns, msg.Signer) && ns.AllowRegistrarOverride
	if !ownerOk && !registrarOk {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "signer is not the current owner and registrar override is not allowed")
	}

	// New owner DID must exist.
	if _, foundDid := k.iidKeeper.GetDidDocument(ctx, []byte(msg.NewOwnerDid)); !foundDid {
		return nil, errorsmod.Wrapf(types.ErrInvalidDID, "new owner DID %q not found", msg.NewOwnerDid)
	}

	oldOwner := record.OwnerDid
	if oldOwner == msg.NewOwnerDid {
		return nil, errorsmod.Wrap(types.ErrInvalidRequest, "new owner is the same as current owner")
	}

	k.removeOwnerIndex(ctx, oldOwner, ns.Name, record.NormalizedName)
	record.OwnerDid = msg.NewOwnerDid
	record.UpdatedAt = ctx.BlockTime().Unix()
	k.SetNameRecord(ctx, record)

	// Dual-emit: NameUpdatedEvent carries the new full record so an indexer can
	// refresh its NameRecord state table from a single hook regardless of which
	// action mutated the row; NameTransferredEvent feeds an audit/history table
	// with the from→to delta and the actor.
	if err := ctx.EventManager().EmitTypedEvents(
		types.NewNameUpdatedEvent(record, msg.Signer),
		types.NewNameTransferredEvent(ns.Name, record.NormalizedName, oldOwner, msg.NewOwnerDid, msg.Signer),
	); err != nil {
		return nil, err
	}
	return &types.MsgTransferNameResponse{}, nil
}

// SetNameStatus changes a record's status. Permitted for namespace registrars
// and for the gov authority. Tombstoned records are terminal: no further
// transitions are allowed.
func (k msgServer) SetNameStatus(goCtx context.Context, msg *types.MsgSetNameStatus) (*types.MsgSetNameStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Defence-in-depth: reject the proto zero value here so a Wasm sub-message
	// dispatch (which bypasses ante / ValidateBasic) can't persist a record
	// in NAME_STATUS_UNSPECIFIED state.
	switch msg.Status {
	case types.NAME_STATUS_ACTIVE,
		types.NAME_STATUS_SUSPENDED,
		types.NAME_STATUS_REVOKED,
		types.NAME_STATUS_TOMBSTONED:
	default:
		return nil, errorsmod.Wrapf(types.ErrInvalidStatusTransition, "unsupported target status %s", msg.Status)
	}
	ns, err := k.requireNamespace(ctx, msg.Namespace)
	if err != nil {
		return nil, err
	}
	if !types.HasRegistrar(ns, msg.Signer) && msg.Signer != k.authority {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "signer must be a registrar or the governance authority")
	}

	record, found := k.GetNameRecord(ctx, ns.Name, msg.NormalizedName)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrNameNotFound, "name %q in namespace %q not found", msg.NormalizedName, ns.Name)
	}
	if record.Status == types.NAME_STATUS_TOMBSTONED {
		return nil, errorsmod.Wrap(types.ErrInvalidStatusTransition, "tombstoned records are terminal")
	}
	if record.Status == msg.Status {
		return nil, errorsmod.Wrap(types.ErrInvalidStatusTransition, "status is already set to the requested value")
	}

	old := record.Status
	record.Status = msg.Status
	record.UpdatedAt = ctx.BlockTime().Unix()
	k.SetNameRecord(ctx, record)

	// Dual-emit: see TransferName for rationale.
	if err := ctx.EventManager().EmitTypedEvents(
		types.NewNameUpdatedEvent(record, msg.Signer),
		types.NewNameStatusChangedEvent(ns.Name, record.NormalizedName, old, msg.Status, msg.Signer, msg.Reason),
	); err != nil {
		return nil, err
	}
	return &types.MsgSetNameStatusResponse{}, nil
}

// requireNamespace returns the namespace or an error if it is missing.
func (k Keeper) requireNamespace(ctx sdk.Context, name string) (types.Namespace, error) {
	ns, found := k.GetNamespace(ctx, name)
	if !found {
		return types.Namespace{}, errorsmod.Wrapf(types.ErrNamespaceNotFound, "namespace %q does not exist", name)
	}
	return ns, nil
}

// normalizeAndCheckAvailable runs the v1 normalization (trim + lowercase),
// validates the result against the namespace's rules, and ensures the slot is
// free. Returns the normalized form on success.
func (k Keeper) normalizeAndCheckAvailable(ctx sdk.Context, ns types.Namespace, displayName string) (string, error) {
	normalized := types.NormalizeName(displayName)
	if err := types.ValidateNameAgainstNamespace(ns, normalized); err != nil {
		return "", err
	}
	if k.HasNameRecord(ctx, ns.Name, normalized) {
		return "", errorsmod.Wrapf(types.ErrNameTaken, "name %q is already taken in namespace %q", normalized, ns.Name)
	}
	return normalized, nil
}
