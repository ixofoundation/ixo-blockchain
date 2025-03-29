package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	iidante "github.com/ixofoundation/ixo-blockchain/v5/x/iid/ante"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v5/x/iid/types"
)

// func didToAddressSplitter(did string) (sdk.AccAddress, error) {
// 	bech32 := strings.Split(did, ":")
// 	address, err := sdk.AccAddressFromBech32(bech32[len(bech32)-1])
// 	if err != nil {
// 		return sdk.AccAddress{}, err
// 	}
// 	return address, nil
// }

var (
	_ iidante.IidTxMsg = &MsgCreateEntity{}
	_ iidante.IidTxMsg = &MsgTransferEntity{}
	_ iidante.IidTxMsg = &MsgUpdateEntity{}
	_ iidante.IidTxMsg = &MsgUpdateEntityVerified{}
)

// --------------------------
// CREATE ENTITY
// --------------------------
const TypeMsgCreateEntity = "create_entity"

var _ sdk.Msg = &MsgCreateEntity{}

func NewMsgCreateEntity(
	entityType string,
	entityStatus int32,
	controller []string,
	context []*iidtypes.Context,
	verification []*iidtypes.Verification,
	service []*iidtypes.Service,
	accordedRight []*iidtypes.AccordedRight,
	linkedResource []*iidtypes.LinkedResource,
	linkedEntity []*iidtypes.LinkedEntity,
	startDate *time.Time,
	endDate *time.Time,
	relayerNode string,
	credentials []string,
	ownerDid iidtypes.DIDFragment,
	ownerAddress string,
	data []byte,
) *MsgCreateEntity {
	return &MsgCreateEntity{
		EntityType:     entityType,
		EntityStatus:   entityStatus,
		Controller:     controller,
		Context:        context,
		Verification:   verification,
		Service:        service,
		AccordedRight:  accordedRight,
		LinkedResource: linkedResource,
		LinkedEntity:   linkedEntity,
		StartDate:      startDate,
		EndDate:        endDate,
		RelayerNode:    relayerNode,
		Credentials:    credentials,
		OwnerDid:       ownerDid,
		OwnerAddress:   ownerAddress,
		Data:           data,
	}
}

func (msg MsgCreateEntity) Type() string {
	return TypeMsgCreateEntity
}

func (msg MsgCreateEntity) GetIidController() iidtypes.DIDFragment { return msg.OwnerDid }

func (msg MsgCreateEntity) Route() string { return RouterKey }

// --------------------------
// UPDATE ENTITY
// --------------------------
const TypeMsgUpdateEntity = "update_entity"

var _ sdk.Msg = &MsgUpdateEntity{}

func NewMsgUpdateEntity(
	id string,
	entityStatus int32,
	startDate *time.Time,
	endDate *time.Time,
	credentials []string,
	controllerDid iidtypes.DIDFragment,
	controllerAddress string,
) *MsgUpdateEntity {
	return &MsgUpdateEntity{
		Id:                id,
		EntityStatus:      entityStatus,
		StartDate:         startDate,
		EndDate:           endDate,
		Credentials:       credentials,
		ControllerDid:     controllerDid,
		ControllerAddress: controllerAddress,
	}
}

func (msg MsgUpdateEntity) Type() string {
	return TypeMsgUpdateEntity
}

func (msg MsgUpdateEntity) GetIidController() iidtypes.DIDFragment { return msg.ControllerDid }

func (msg MsgUpdateEntity) Route() string { return RouterKey }

// --------------------------
// UPDATE ENTITY VERIFIED
// --------------------------
const TypeMsgUpdateEntityVerified = "update_entity_verification"

var _ sdk.Msg = &MsgUpdateEntityVerified{}

func NewMsgUpdateEntityVerified(
	id string,
	entityVerified bool,
	relayerNodeDid iidtypes.DIDFragment,
	relayerNodeAddress string,
) *MsgUpdateEntityVerified {
	return &MsgUpdateEntityVerified{
		Id:                 id,
		EntityVerified:     entityVerified,
		RelayerNodeDid:     relayerNodeDid,
		RelayerNodeAddress: relayerNodeAddress,
	}
}

func (msg MsgUpdateEntityVerified) Type() string {
	return TypeMsgUpdateEntityVerified
}

func (msg MsgUpdateEntityVerified) GetIidController() iidtypes.DIDFragment { return msg.RelayerNodeDid }

func (msg MsgUpdateEntityVerified) Route() string { return RouterKey }

// --------------------------
// TRANSFER ENTITY
// --------------------------
const TypeMsgTransferEntity = "transfer_entity"

var _ sdk.Msg = &MsgTransferEntity{}

func NewMsgTransferEntity(
	id string,
	ownerDid iidtypes.DIDFragment,
	ownerAddress string,
	recipientDid iidtypes.DIDFragment,
) *MsgTransferEntity {
	return &MsgTransferEntity{
		Id:           id,
		OwnerDid:     ownerDid,
		OwnerAddress: ownerAddress,
		RecipientDid: recipientDid,
	}
}

func (msg MsgTransferEntity) Type() string {
	return TypeMsgTransferEntity
}

func (msg MsgTransferEntity) GetIidController() iidtypes.DIDFragment { return msg.OwnerDid }

func (msg MsgTransferEntity) Route() string { return RouterKey }

// --------------------------
// CREATE ENTITY ACCOUNT
// --------------------------
const TypeMsgCreateEntityAccount = "create_entity_account"

var _ sdk.Msg = &MsgCreateEntityAccount{}

func NewMsgCreateEntityAccount(
	id, name string,
	ownerAddress string,
) *MsgCreateEntityAccount {
	return &MsgCreateEntityAccount{
		Id:           id,
		Name:         name,
		OwnerAddress: ownerAddress,
	}
}

func (msg MsgCreateEntityAccount) Type() string {
	return TypeMsgCreateEntityAccount
}

func (msg MsgCreateEntityAccount) Route() string { return RouterKey }

// --------------------------
// GRANT ENTITY ACCOUNT AUTHZ
// --------------------------
const TypeMsgGrantEntityAccountAuthz = "grant_entity_account_authz"

var _ sdk.Msg = &MsgGrantEntityAccountAuthz{}

func NewMsgGrantEntityAccountAuthz(
	id, name, ownerAddress, granteeAddress string,
	grant authz.Grant,
) *MsgGrantEntityAccountAuthz {
	return &MsgGrantEntityAccountAuthz{
		Id:             id,
		Name:           name,
		OwnerAddress:   ownerAddress,
		GranteeAddress: granteeAddress,
		Grant:          grant,
	}
}

func (msg MsgGrantEntityAccountAuthz) Type() string {
	return TypeMsgGrantEntityAccountAuthz
}

func (msg MsgGrantEntityAccountAuthz) Route() string { return RouterKey }

// --------------------------
// REVOKE ENTITY ACCOUNT AUTHZ
// --------------------------
const TypeMsgRevokeEntityAccountAuthz = "revoke_entity_account_authz"

var _ sdk.Msg = &MsgRevokeEntityAccountAuthz{}

func NewMsgRevokeEntityAccountAuthz(
	id, name, ownerAddress, granteeAddress, msgTypeUrl string,
) *MsgRevokeEntityAccountAuthz {
	return &MsgRevokeEntityAccountAuthz{
		Id:             id,
		Name:           name,
		OwnerAddress:   ownerAddress,
		GranteeAddress: granteeAddress,
		MsgTypeUrl:     msgTypeUrl,
	}
}

func (msg MsgRevokeEntityAccountAuthz) Type() string {
	return TypeMsgRevokeEntityAccountAuthz
}

func (msg MsgRevokeEntityAccountAuthz) Route() string { return RouterKey }
