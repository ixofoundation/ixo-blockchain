package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgCreateIidDocument{}
var _ sdk.Msg = &MsgUpdateIidDocument{}
var _ sdk.Msg = &MsgAddVerification{}
var _ sdk.Msg = &MsgRevokeVerification{}
var _ sdk.Msg = &MsgSetVerificationRelationships{}
var _ sdk.Msg = &MsgAddService{}
var _ sdk.Msg = &MsgDeleteService{}
var _ sdk.Msg = &MsgAddController{}
var _ sdk.Msg = &MsgDeleteController{}
var _ sdk.Msg = &MsgAddLinkedResource{}
var _ sdk.Msg = &MsgDeleteLinkedResource{}
var _ sdk.Msg = &MsgAddLinkedClaim{}
var _ sdk.Msg = &MsgDeleteLinkedClaim{}
var _ sdk.Msg = &MsgAddLinkedEntity{}
var _ sdk.Msg = &MsgDeleteLinkedEntity{}
var _ sdk.Msg = &MsgAddAccordedRight{}
var _ sdk.Msg = &MsgDeleteAccordedRight{}
var _ sdk.Msg = &MsgAddIidContext{}
var _ sdk.Msg = &MsgDeactivateIID{}
var _ sdk.Msg = &MsgDeleteIidContext{}

// --------------------------
// CREATE IDENTIFIER
// --------------------------
const TypeMsgCreateDidDocument = "create-did"

// NewMsgCreateDidDocument creates a new MsgCreateDidDocument instance
func NewMsgCreateIidDocument(
	id string,
	verifications []*Verification,
	controllers []string,
	services []*Service,
	rights []*AccordedRight,
	resources []*LinkedResource,
	entity []*LinkedEntity,
	signerAccount string,
	didContexts []*Context,
) *MsgCreateIidDocument {
	return &MsgCreateIidDocument{
		Id:             id,
		Verifications:  verifications,
		Services:       services,
		AccordedRight:  rights,
		LinkedResource: resources,
		LinkedEntity:   entity,
		Context:        didContexts,
		Signer:         signerAccount,
		Controllers:    controllers,
	}
}

// Route implements sdk.Msg
func (MsgCreateIidDocument) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgCreateIidDocument) Type() string {
	return TypeMsgCreateDidDocument
}

func (msg MsgCreateIidDocument) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgCreateIidDocument) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// UPDATE IDENTIFIER
// --------------------------
const TypeMsgUpdateDidDocument = "update-did"

// Route implements sdk.Msg
func (MsgUpdateIidDocument) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgUpdateIidDocument) Type() string {
	return TypeMsgUpdateDidDocument
}

func (msg MsgUpdateIidDocument) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateIidDocument) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD VERIFICATION
// --------------------------
const TypeMsgAddVerification = "add-verification"

// NewMsgAddVerification creates a new MsgAddVerification instance
func NewMsgAddVerification(
	id string,
	verification *Verification,
	signerAccount string,
) *MsgAddVerification {
	return &MsgAddVerification{
		Id:           id,
		Verification: verification,
		Signer:       signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddVerification) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddVerification) Type() string {
	return TypeMsgAddVerification
}

func (msg MsgAddVerification) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAddVerification) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// REVOKE VERIFICATION
// --------------------------
const TypeMsgRevokeVerification = "revoke-verification"

// NewMsgRevokeVerification creates a new MsgRevokeVerification instance
func NewMsgRevokeVerification(
	id string,
	methodID string,
	signerAccount string,
) *MsgRevokeVerification {
	return &MsgRevokeVerification{
		Id:       id,
		MethodId: methodID,
		Signer:   signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgRevokeVerification) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgRevokeVerification) Type() string {
	return TypeMsgRevokeVerification
}

func (msg MsgRevokeVerification) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgRevokeVerification) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// SET VERIFICATION RELATIONSHIPS
// --------------------------
const TypeMsgSetVerificationRelationships = "set-verification-relationships"

func NewMsgSetVerificationRelationships(
	id string,
	methodID string,
	relationships []string,
	signerAccount string,
) *MsgSetVerificationRelationships {
	return &MsgSetVerificationRelationships{
		Id:            id,
		MethodId:      methodID,
		Relationships: relationships,
		Signer:        signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgSetVerificationRelationships) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgSetVerificationRelationships) Type() string {
	return TypeMsgSetVerificationRelationships
}

func (msg MsgSetVerificationRelationships) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgSetVerificationRelationships) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD SERVICE
// --------------------------
const TypeMsgAddService = "add-service"

// NewMsgAddService creates a new MsgAddService instance
func NewMsgAddService(
	id string,
	service *Service,
	signerAccount string,
) *MsgAddService {
	return &MsgAddService{
		Id:          id,
		ServiceData: service,
		Signer:      signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddService) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddService) Type() string {
	return TypeMsgAddService
}

func (msg MsgAddService) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAddService) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD LINKED RESOURCE
// --------------------------
const TypeMsgAddLinkedResource = "add-linked-resource"

func NewMsgAddLinkedResource(
	id string,
	linkedResource *LinkedResource,
	signerAccount string,
) *MsgAddLinkedResource {
	return &MsgAddLinkedResource{
		Id:             id,
		LinkedResource: linkedResource,
		Signer:         signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddLinkedResource) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddLinkedResource) Type() string {
	return TypeMsgAddLinkedResource
}

func (msg MsgAddLinkedResource) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAddLinkedResource) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD LINKED CLAIM
// --------------------------
const TypeMsgAddLinkedClaim = "add-linked-claim"

func NewMsgAddLinkedClaim(
	id string,
	linkedClaim *LinkedClaim,
	signerAccount string,
) *MsgAddLinkedClaim {
	return &MsgAddLinkedClaim{
		Id:          id,
		LinkedClaim: linkedClaim,
		Signer:      signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddLinkedClaim) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddLinkedClaim) Type() string {
	return TypeMsgAddLinkedClaim
}

func (msg MsgAddLinkedClaim) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAddLinkedClaim) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD LINKED ENTITY
// --------------------------
const TypeMsgAddLinkedEntity = "add-linked-entity"

func NewMsgAddLinkedEntity(
	id string,
	linkedResource *LinkedEntity,
	signerAccount string,
) *MsgAddLinkedEntity {
	return &MsgAddLinkedEntity{
		Id:           id,
		LinkedEntity: linkedResource,
		Signer:       signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddLinkedEntity) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddLinkedEntity) Type() string {
	return TypeMsgAddLinkedEntity
}

func (msg MsgAddLinkedEntity) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

func (msg MsgAddLinkedEntity) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE SERVICE
// --------------------------
const TypeMsgDeleteService = "delete-service"

func NewMsgDeleteService(
	id string,
	serviceID string,
	signerAccount string,
) *MsgDeleteService {
	return &MsgDeleteService{
		Id:        id,
		ServiceId: serviceID,
		Signer:    signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgDeleteService) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteService) Type() string {
	return TypeMsgDeleteService
}

func (msg MsgDeleteService) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteService) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE LINKED RESOURCE
// --------------------------
const TypeMsgDeleteLinkedResource = "delete-linked-resource"

func NewMsgDeleteLinkedResource(
	id string,
	resourceID string,
	signerAccount string,
) *MsgDeleteLinkedResource {
	return &MsgDeleteLinkedResource{
		Id:         id,
		ResourceId: resourceID,
		Signer:     signerAccount,
	}
}

func (MsgDeleteLinkedResource) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteLinkedResource) Type() string {
	return TypeMsgDeleteLinkedResource
}

func (msg MsgDeleteLinkedResource) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteLinkedResource) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE LINKED CLAIM
// --------------------------
const TypeMsgDeleteLinkedClaim = "delete-linked-claim"

func NewMsgDeleteLinkedClaim(
	id string,
	claimID string,
	signerAccount string,
) *MsgDeleteLinkedClaim {
	return &MsgDeleteLinkedClaim{
		Id:      id,
		ClaimId: claimID,
		Signer:  signerAccount,
	}
}

func (MsgDeleteLinkedClaim) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteLinkedClaim) Type() string {
	return TypeMsgDeleteLinkedClaim
}

func (msg MsgDeleteLinkedClaim) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteLinkedClaim) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE LINKED ENTITY
// --------------------------
const TypeMsgDeleteLinkedEntity = "delete-linked-entity"

func NewMsgDeleteLinkedEntity(
	id string,
	entityID string,
	signerAccount string,
) *MsgDeleteLinkedEntity {
	return &MsgDeleteLinkedEntity{
		Id:       id,
		EntityId: entityID,
		Signer:   signerAccount,
	}
}

func (MsgDeleteLinkedEntity) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteLinkedEntity) Type() string {
	return TypeMsgDeleteLinkedEntity
}

func (msg MsgDeleteLinkedEntity) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

func (msg MsgDeleteLinkedEntity) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD RIGHT
// --------------------------
const TypeMsgAddRight = "add-right"

// NewMsgAddAccordedRight creates a new MsgAddAccordedright instance
func NewMsgAddAccordedRight(
	id string,
	right *AccordedRight,
	signerAccount string,
) *MsgAddAccordedRight {
	return &MsgAddAccordedRight{
		Id:            id,
		AccordedRight: right,
		Signer:        signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddAccordedRight) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddAccordedRight) Type() string {
	return TypeMsgAddRight
}

func (msg MsgAddAccordedRight) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAddAccordedRight) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE RIGHT
// --------------------------
const TypeMsgDeleteAccordedRight = "delete-right"

func NewMsgDeleteAccordedRight(
	id string,
	rightID string,
	signerAccount string,
) *MsgDeleteAccordedRight {
	return &MsgDeleteAccordedRight{
		Id:      id,
		RightId: rightID,
		Signer:  signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgDeleteAccordedRight) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteAccordedRight) Type() string {
	return TypeMsgDeleteAccordedRight
}

func (msg MsgDeleteAccordedRight) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteAccordedRight) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD CONTROLLER
// --------------------------
const TypeMsgAddController = "add-controller"

func NewMsgAddController(
	id string,
	controllerDID string,
	signerAccount string,
) *MsgAddController {
	return &MsgAddController{
		Id:            id,
		ControllerDid: controllerDID,
		Signer:        signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddController) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddController) Type() string {
	return TypeMsgAddController
}

func (msg MsgAddController) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAddController) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE CONTROLLER
// --------------------------
const TypeMsgDeleteController = "delete-controller"

func NewMsgDeleteController(
	id string,
	controllerDID string,
	signerAccount string,
) *MsgDeleteController {
	return &MsgDeleteController{
		Id:            id,
		ControllerDid: controllerDID,
		Signer:        signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgDeleteController) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteController) Type() string {
	return TypeMsgDeleteController
}

func (msg MsgDeleteController) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteController) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD CONTEXT
// --------------------------
const TypeMsgAddContext = "add-did-context"

// NewMsgAddService creates a new MsgAddService instance
func NewMsgAddDidContext(
	id string,
	context *Context,
	signerAccount string,
) *MsgAddIidContext {
	return &MsgAddIidContext{
		Id:      id,
		Context: context,
		Signer:  signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddIidContext) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddIidContext) Type() string {
	return TypeMsgAddContext
}

func (msg MsgAddIidContext) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAddIidContext) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE CONTEXT
// --------------------------
const TypeMsgDeleteDidContext = "delete-context"

func NewMsgDeleteDidContext(
	id string,
	key string,
	signerAccount string,
) *MsgDeleteIidContext {
	return &MsgDeleteIidContext{
		Id:         id,
		ContextKey: key,
		Signer:     signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgDeleteIidContext) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteIidContext) Type() string {
	return TypeMsgDeleteDidContext
}

func (msg MsgDeleteIidContext) GetSignBytes() []byte {
	// panic("IBC messages do not support amino")
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteIidContext) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DEACTIVATE DID
// --------------------------
const TypeMsgDeactivateDid = "deactivate-did"

func NewMsgDeactivateIID(
	id string,
	state bool,
	signerAccount string,
) *MsgDeactivateIID {
	return &MsgDeactivateIID{
		Id:     id,
		State:  state,
		Signer: signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgDeactivateIID) Route() string {
	return RouterKey
}

func (MsgDeactivateIID) Type() string {
	return TypeMsgDeactivateDid
}

func (msg MsgDeactivateIID) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDeactivateIID) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
