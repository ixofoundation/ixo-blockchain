package types

import (
	"encoding/json"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

type CreateProjectMsg struct {
	SignBytes  string     `json:"signBytes"`
	TxHash     string     `json:"txHash"`
	SenderDid  ixo.Did    `json:"senderDid"`
	ProjectDid ixo.Did    `json:"projectDid"`
	PubKey     string     `json:"pubKey"`
	Data       ProjectDoc `json:"data"`
}

var _ sdk.Msg = CreateProjectMsg{}

func (msg CreateProjectMsg) Type() string                            { return ModuleName }
func (msg CreateProjectMsg) Route() string                           { return RouterKey }
func (msg CreateProjectMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateProjectMsg) ValidateBasic() sdk.Error {
	valid, err := CheckNotEmpty(msg.PubKey, "PubKey")
	if !valid {
		return err
	}
	
	valid, err = CheckNotEmpty(msg.ProjectDid, "ProjectDid")
	if !valid {
		return err
	}
	
	valid, err = CheckNotEmpty(msg.Data.NodeDid, "NodeDid")
	if !valid {
		return err
	}
	
	valid, err = CheckNotEmpty(msg.Data.RequiredClaims, "RequiredClaims")
	if !valid {
		return err
	}
	
	valid, err = CheckNotEmpty(msg.Data.CreatedBy, "CreatedBy")
	if !valid {
		return err
	}
	
	return nil
}

func (msg CreateProjectMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateProjectMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateProjectMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}

func (msg CreateProjectMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg CreateProjectMsg) GetPubKey() string        { return msg.PubKey }
func (msg CreateProjectMsg) GetEvaluatorPay() int64   { return msg.Data.GetEvaluatorPay() }
func (msg CreateProjectMsg) GetStatus() ProjectStatus { return msg.Data.Status }
func (msg *CreateProjectMsg) SetStatus(status ProjectStatus) {
	msg.Data.Status = status
}

func (msg CreateProjectMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg CreateProjectMsg) IsNewDid() bool     { return true }
func (msg CreateProjectMsg) IsWithdrawal() bool { return false }

var _ StoredProjectDoc = (*CreateProjectMsg)(nil)

type UpdateProjectStatusMsg struct {
	SignBytes  string                 `json:"signBytes"`
	TxHash     string                 `json:"txHash"`
	SenderDid  ixo.Did                `json:"senderDid"`
	ProjectDid ixo.Did                `json:"projectDid"`
	Data       UpdateProjectStatusDoc `json:"data"`
}

func (msg UpdateProjectStatusMsg) Type() string                            { return ModuleName }
func (msg UpdateProjectStatusMsg) Route() string                           { return RouterKey }
func (msg UpdateProjectStatusMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg UpdateProjectStatusMsg) ValidateBasic() sdk.Error                { return nil }
func (msg UpdateProjectStatusMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg UpdateProjectStatusMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}

func (ups UpdateProjectStatusMsg) GetProjectDid() ixo.Did {
	return ups.ProjectDid
}

func (ups UpdateProjectStatusMsg) GetStatus() ProjectStatus {
	return ups.Data.Status
}

func (msg UpdateProjectStatusMsg) IsNewDid() bool     { return false }
func (msg UpdateProjectStatusMsg) IsWithdrawal() bool { return false }
func (msg UpdateProjectStatusMsg) GetEthFundingTxnID() string {
	return msg.Data.EthFundingTxnID
}

type CreateAgentMsg struct {
	SignBytes  string         `json:"signBytes"`
	TxHash     string         `json:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid"`
	Data       CreateAgentDoc `json:"data"`
}

func (msg CreateAgentMsg) IsNewDid() bool                          { return false }
func (msg CreateAgentMsg) IsWithdrawal() bool                      { return false }
func (msg CreateAgentMsg) Type() string                            { return ModuleName }
func (msg CreateAgentMsg) Route() string                           { return RouterKey }
func (msg CreateAgentMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateAgentMsg) ValidateBasic() sdk.Error {
	return nil
}

func (msg CreateAgentMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateAgentMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateAgentMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}

func (msg CreateAgentMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg CreateAgentMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

var _ sdk.Msg = CreateAgentMsg{}

type UpdateAgentMsg struct {
	SignBytes  string         `json:"signBytes"`
	TxHash     string         `json:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid"`
	Data       UpdateAgentDoc `json:"data"`
}

func (msg UpdateAgentMsg) IsNewDid() bool                          { return false }
func (msg UpdateAgentMsg) IsWithdrawal() bool                      { return false }
func (msg UpdateAgentMsg) Type() string                            { return ModuleName }
func (msg UpdateAgentMsg) Route() string                           { return RouterKey }
func (msg UpdateAgentMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg UpdateAgentMsg) ValidateBasic() sdk.Error {
	return nil
}

func (msg UpdateAgentMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg UpdateAgentMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg UpdateAgentMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}

func (msg UpdateAgentMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg UpdateAgentMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	
	return string(b)
}

var _ sdk.Msg = UpdateAgentMsg{}

type CreateClaimMsg struct {
	SignBytes  string         `json:"signBytes"`
	TxHash     string         `json:"txHash"`
	SenderDid  ixo.Did        `json:"senderDid"`
	ProjectDid ixo.Did        `json:"projectDid"`
	Data       CreateClaimDoc `json:"data"`
}

func (msg CreateClaimMsg) IsNewDid() bool                          { return false }
func (msg CreateClaimMsg) IsWithdrawal() bool                      { return false }
func (msg CreateClaimMsg) Type() string                            { return ModuleName }
func (msg CreateClaimMsg) Route() string                           { return RouterKey }
func (msg CreateClaimMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateClaimMsg) ValidateBasic() sdk.Error {
	return nil
}

func (msg CreateClaimMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateClaimMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateClaimMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}

func (msg CreateClaimMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg CreateClaimMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	
	return string(b)
}

var _ sdk.Msg = CreateClaimMsg{}

type CreateEvaluationMsg struct {
	SignBytes  string              `json:"signBytes"`
	TxHash     string              `json:"txHash"`
	SenderDid  ixo.Did             `json:"senderDid"`
	ProjectDid ixo.Did             `json:"projectDid"`
	Data       CreateEvaluationDoc `json:"data"`
}

func (msg CreateEvaluationMsg) IsNewDid() bool                          { return false }
func (msg CreateEvaluationMsg) IsWithdrawal() bool                      { return false }
func (msg CreateEvaluationMsg) Type() string                            { return ModuleName }
func (msg CreateEvaluationMsg) Route() string                           { return RouterKey }
func (msg CreateEvaluationMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg CreateEvaluationMsg) ValidateBasic() sdk.Error {
	return nil
}

func (msg CreateEvaluationMsg) GetProjectDid() ixo.Did { return msg.ProjectDid }
func (msg CreateEvaluationMsg) GetSenderDid() ixo.Did  { return msg.SenderDid }
func (msg CreateEvaluationMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetProjectDid())}
}

func (msg CreateEvaluationMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg CreateEvaluationMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	
	return string(b)
}

var _ sdk.Msg = CreateEvaluationMsg{}

type WithdrawFundsMsg struct {
	SignBytes string           `json:"signBytes"`
	SenderDid ixo.Did          `json:"senderDid"`
	Data      WithdrawFundsDoc `json:"data"`
}

func (msg WithdrawFundsMsg) IsNewDid() bool                          { return false }
func (msg WithdrawFundsMsg) IsWithdrawal() bool                      { return true }
func (msg WithdrawFundsMsg) Type() string                            { return ModuleName }
func (msg WithdrawFundsMsg) Route() string                           { return RouterKey }
func (msg WithdrawFundsMsg) Get(key interface{}) (value interface{}) { return nil }
func (msg WithdrawFundsMsg) ValidateBasic() sdk.Error {
	return nil
}

func (msg WithdrawFundsMsg) GetSenderDid() ixo.Did                 { return msg.SenderDid }
func (msg WithdrawFundsMsg) GetWithdrawFundsDoc() WithdrawFundsDoc { return msg.Data }
func (msg WithdrawFundsMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{[]byte(msg.GetSenderDid())}
}

func (msg WithdrawFundsMsg) GetSignBytes() []byte {
	return []byte(msg.SignBytes)
}

func (msg WithdrawFundsMsg) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	
	return string(b)
}

var _ sdk.Msg = WithdrawFundsMsg{}
