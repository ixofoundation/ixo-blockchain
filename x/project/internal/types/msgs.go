package types

import (
	"encoding/json"
	"github.com/ixofoundation/ixo-blockchain/x/did"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

var (
	_ ixo.IxoMsg = MsgCreateProject{}
	_ ixo.IxoMsg = MsgUpdateProjectStatus{}
	_ ixo.IxoMsg = MsgCreateAgent{}
	_ ixo.IxoMsg = MsgUpdateAgent{}
	_ ixo.IxoMsg = MsgCreateClaim{}
	_ ixo.IxoMsg = MsgCreateEvaluation{}
	_ ixo.IxoMsg = MsgWithdrawFunds{}

	_ StoredProjectDoc = (*MsgCreateProject)(nil)
)

type MsgCreateProject struct {
	TxHash     string     `json:"txHash" yaml:"txHash"`
	SenderDid  ixo.Did    `json:"senderDid" yaml:"senderDid"`
	ProjectDid ixo.Did    `json:"projectDid" yaml:"projectDid"`
	PubKey     string     `json:"pubKey" yaml:"pubKey"`
	Data       ProjectDoc `json:"data" yaml:"data"`
}

func (msg MsgCreateProject) Type() string { return "create-project" }

func (msg MsgCreateProject) Route() string { return RouterKey }

func (msg MsgCreateProject) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Data.NodeDid, "NodeDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Data.RequiredClaims, "RequiredClaims"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Data.CreatedBy, "CreatedBy"); !valid {
		return err
	}

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !ixo.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	return nil
}

func (msg MsgCreateProject) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg MsgCreateProject) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg MsgCreateProject) GetSignerDid() ixo.Did  { return msg.ProjectDid }

func (msg MsgCreateProject) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgCreateProject) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreateProject) GetPubKey() string        { return msg.PubKey }
func (msg MsgCreateProject) GetEvaluatorPay() int64   { return msg.Data.GetEvaluatorPay() }
func (msg MsgCreateProject) GetStatus() ProjectStatus { return msg.Data.Status }
func (msg *MsgCreateProject) SetStatus(status ProjectStatus) {
	msg.Data.Status = status
}

func (msg MsgCreateProject) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

type MsgUpdateProjectStatus struct {
	TxHash     string                 `json:"txHash" yaml:"txHash"`
	SenderDid  ixo.Did                `json:"senderDid" yaml:"senderDid"`
	ProjectDid ixo.Did                `json:"projectDid" yaml:"projectDid"`
	Data       UpdateProjectStatusDoc `json:"data" yaml:"data"`
}

func (msg MsgUpdateProjectStatus) Type() string  { return "update-project-status" }
func (msg MsgUpdateProjectStatus) Route() string { return RouterKey }

func (msg MsgUpdateProjectStatus) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// TODO: perform some checks on the Data (of type UpdateProjectStatusDoc)

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !ixo.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	// IsValidProgressionFrom checked by the handler

	return nil
}

func (msg MsgUpdateProjectStatus) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgUpdateProjectStatus) GetSignerDid() ixo.Did {
	return msg.ProjectDid
}

func (msg MsgUpdateProjectStatus) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

type MsgCreateAgent struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid" yaml:"projectDid"`
	Data       CreateAgentDoc `json:"data" yaml:"data"`
}

func (msg MsgCreateAgent) Type() string  { return "create-agent" }
func (msg MsgCreateAgent) Route() string { return RouterKey }
func (msg MsgCreateAgent) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// TODO: perform some checks on the Data (of type CreateAgentDoc)

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !ixo.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	} else if !ixo.IsValidDid(msg.Data.AgentDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "agent did is invalid")
	}

	return nil
}

func (msg MsgCreateAgent) GetSignerDid() ixo.Did {
	return msg.ProjectDid
}

func (msg MsgCreateAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgCreateAgent) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgCreateAgent) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

type MsgUpdateAgent struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid" yaml:"projectDid"`
	Data       UpdateAgentDoc `json:"data" yaml:"data"`
}

func (msg MsgUpdateAgent) Type() string  { return "update-agent" }
func (msg MsgUpdateAgent) Route() string { return RouterKey }
func (msg MsgUpdateAgent) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// TODO: perform some checks on the Data (of type UpdateAgentDoc)

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !ixo.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	} else if !ixo.IsValidDid(msg.Data.Did) {
		return did.ErrorInvalidDid(DefaultCodespace, "agent did is invalid")
	}

	return nil
}

func (msg MsgUpdateAgent) GetSignerDid() ixo.Did {
	return msg.ProjectDid
}

func (msg MsgUpdateAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgUpdateAgent) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgUpdateAgent) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(b)
}

type MsgCreateClaim struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid" yaml:"projectDid"`
	Data       CreateClaimDoc `json:"data" yaml:"data"`
}

func (msg MsgCreateClaim) Type() string  { return "create-claim" }
func (msg MsgCreateClaim) Route() string { return RouterKey }

func (msg MsgCreateClaim) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// TODO: perform some checks on the Data (of type CreateClaimDoc)

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !ixo.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	return nil
}

func (msg MsgCreateClaim) GetSignerDid() ixo.Did {
	return msg.ProjectDid
}

func (msg MsgCreateClaim) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgCreateClaim) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgCreateClaim) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(b)
}

type MsgCreateEvaluation struct {
	TxHash     string              `json:"txHash" yaml:"txHash"`
	SenderDid  ixo.Did             `json:"senderDid" yaml:"senderDid"`
	ProjectDid ixo.Did             `json:"projectDid" yaml:"projectDid"`
	Data       CreateEvaluationDoc `json:"data" yaml:"data"`
}

func (msg MsgCreateEvaluation) Type() string  { return "create-evaluation" }
func (msg MsgCreateEvaluation) Route() string { return RouterKey }

func (msg MsgCreateEvaluation) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// TODO: perform some checks on the Data (of type CreateEvaluationDoc)

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !ixo.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	return nil
}

func (msg MsgCreateEvaluation) GetSignerDid() ixo.Did {
	return msg.ProjectDid
}

func (msg MsgCreateEvaluation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgCreateEvaluation) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgCreateEvaluation) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(b)
}

type MsgWithdrawFunds struct {
	SenderDid ixo.Did          `json:"senderDid" yaml:"senderDid"`
	Data      WithdrawFundsDoc `json:"data" yaml:"data"`
}

func (msg MsgWithdrawFunds) Type() string  { return "withdraw-funds" }
func (msg MsgWithdrawFunds) Route() string { return RouterKey }

func (msg MsgWithdrawFunds) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Data.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Data.RecipientDid, "RecipientDid"); !valid {
		return err
	}

	// TODO: perform some checks on the Data (of type WithdrawFundsDoc)

	// Check that DIDs valid
	if !ixo.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	} else if !ixo.IsValidDid(msg.Data.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !ixo.IsValidDid(msg.Data.RecipientDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "recipient did is invalid")
	}

	// Check that the sender is also the recipient
	if msg.SenderDid != msg.Data.RecipientDid {
		return sdk.ErrInternal("sender did must match recipient did")
	}

	// Check that amount is positive
	if !msg.Data.Amount.IsPositive() {
		return sdk.ErrInternal("amount should be positive")
	}

	return nil
}

func (msg MsgWithdrawFunds) GetSignerDid() ixo.Did {
	return msg.Data.RecipientDid
}

func (msg MsgWithdrawFunds) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ixo.DidToAddr(msg.GetSignerDid())}
}

func (msg MsgWithdrawFunds) GetSignBytes() []byte {
	if bz, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return sdk.MustSortJSON(bz)
	}
}

func (msg MsgWithdrawFunds) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(b)
}
