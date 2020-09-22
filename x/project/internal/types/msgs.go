package types

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/spf13/viper"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

const (
	TypeMsgCreateProject       = "create-project"
	TypeMsgUpdateProjectStatus = "update-project-status"
	TypeMsgCreateAgent         = "create-agent"
	TypeMsgUpdateAgent         = "update-agent"
	TypeMsgCreateClaim         = "create-claim"
	TypeMsgCreateEvaluation    = "create-evaluation"
	TypeMsgWithdrawFunds       = "withdraw-funds"

	MsgCreateProjectFee            = int64(1000000)
	MsgCreateProjectTransactionFee = int64(10000)
)

var (
	_ ixo.IxoMsg = MsgCreateProject{}
	_ ixo.IxoMsg = MsgUpdateProjectStatus{}
	_ ixo.IxoMsg = MsgCreateAgent{}
	_ ixo.IxoMsg = MsgUpdateAgent{}
	_ ixo.IxoMsg = MsgCreateClaim{}
	_ ixo.IxoMsg = MsgCreateEvaluation{}
	_ ixo.IxoMsg = MsgWithdrawFunds{}
)

type MsgCreateProject struct {
	TxHash     string          `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did         `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did         `json:"projectDid" yaml:"projectDid"`
	PubKey     string          `json:"pubKey" yaml:"pubKey"`
	Data       json.RawMessage `json:"data" yaml:"data"`
}

func (msg MsgCreateProject) ToStdSignMsg(fee int64) auth.StdSignMsg {
	chainID := viper.GetString(flags.FlagChainID)
	accNum, accSeq := uint64(0), uint64(0)
	stdFee := auth.NewStdFee(0, sdk.NewCoins(sdk.NewCoin(
		ixo.IxoNativeToken, sdk.NewInt(fee))))
	memo := viper.GetString(flags.FlagMemo)

	return auth.StdSignMsg{
		ChainID:       chainID,
		AccountNumber: accNum,
		Sequence:      accSeq,
		Fee:           stdFee,
		Msgs:          []sdk.Msg{msg},
		Memo:          memo,
	}
}

func (msg MsgCreateProject) Type() string { return TypeMsgCreateProject }

func (msg MsgCreateProject) Route() string { return RouterKey }

func (msg MsgCreateProject) ValidateBasic() sdk.Error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.PubKey, "PubKey"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	}

	// Check that data marshallable to map[string]json.RawMessage
	var dataMap ProjectDataMap
	err := json.Unmarshal(msg.Data, &dataMap)
	if err != nil {
		return sdk.ErrInternal(err.Error())
	}

	// Check that DIDs and PubKey valid
	if !did.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !did.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	} else if !did.IsValidPubKey(msg.PubKey) {
		return did.ErrorInvalidPubKey(DefaultCodespace, "pubKey is invalid")
	}

	// Check that project DID matches the PubKey
	unprefixedDid := exported.UnprefixedDid(msg.ProjectDid)
	expectedUnprefixedDid := exported.UnprefixedDidFromPubKey(msg.PubKey)
	if unprefixedDid != expectedUnprefixedDid {
		return did.ErrorDidPubKeyMismatch(DefaultCodespace, fmt.Sprintf(
			"did not deducable from pubKey; expected: %s received: %s",
			expectedUnprefixedDid, unprefixedDid))
	}

	return nil
}

func (msg MsgCreateProject) GetSignerDid() did.Did { return msg.ProjectDid }
func (msg MsgCreateProject) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreateProject) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (msg MsgCreateProject) GetPubKey() string { return msg.PubKey }

func (msg MsgCreateProject) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

type MsgUpdateProjectStatus struct {
	TxHash     string                 `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did                `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did                `json:"projectDid" yaml:"projectDid"`
	Data       UpdateProjectStatusDoc `json:"data" yaml:"data"`
}

func (msg MsgUpdateProjectStatus) Type() string  { return TypeMsgUpdateProjectStatus }
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
	if !did.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !did.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	// IsValidProgressionFrom checked by the handler

	return nil
}

func (msg MsgUpdateProjectStatus) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgUpdateProjectStatus) GetSignerDid() did.Did { return msg.ProjectDid }
func (msg MsgUpdateProjectStatus) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

type MsgCreateAgent struct {
	TxHash     string         `json:"txHash" yaml:"txHash"`
	SenderDid  did.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did        `json:"projectDid" yaml:"projectDid"`
	Data       CreateAgentDoc `json:"data" yaml:"data"`
}

func (msg MsgCreateAgent) Type() string  { return TypeMsgCreateAgent }
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
	if !did.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !did.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	} else if !did.IsValidDid(msg.Data.AgentDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "agent did is invalid")
	}

	return nil
}

func (msg MsgCreateAgent) GetSignerDid() did.Did { return msg.ProjectDid }
func (msg MsgCreateAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreateAgent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
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
	SenderDid  did.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did        `json:"projectDid" yaml:"projectDid"`
	Data       UpdateAgentDoc `json:"data" yaml:"data"`
}

func (msg MsgUpdateAgent) Type() string  { return TypeMsgUpdateAgent }
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
	if !did.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !did.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	} else if !did.IsValidDid(msg.Data.Did) {
		return did.ErrorInvalidDid(DefaultCodespace, "agent did is invalid")
	}

	return nil
}

func (msg MsgUpdateAgent) GetSignerDid() did.Did { return msg.ProjectDid }
func (msg MsgUpdateAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgUpdateAgent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
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
	SenderDid  did.Did        `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did        `json:"projectDid" yaml:"projectDid"`
	Data       CreateClaimDoc `json:"data" yaml:"data"`
}

func (msg MsgCreateClaim) Type() string  { return TypeMsgCreateClaim }
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
	if !did.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !did.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	return nil
}

func (msg MsgCreateClaim) GetSignerDid() did.Did { return msg.ProjectDid }
func (msg MsgCreateClaim) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreateClaim) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
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
	SenderDid  did.Did             `json:"senderDid" yaml:"senderDid"`
	ProjectDid did.Did             `json:"projectDid" yaml:"projectDid"`
	Data       CreateEvaluationDoc `json:"data" yaml:"data"`
}

func (msg MsgCreateEvaluation) Type() string  { return TypeMsgCreateEvaluation }
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
	if !did.IsValidDid(msg.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !did.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	}

	return nil
}

func (msg MsgCreateEvaluation) GetSignerDid() did.Did { return msg.ProjectDid }
func (msg MsgCreateEvaluation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreateEvaluation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateEvaluation) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(b)
}

type MsgWithdrawFunds struct {
	SenderDid did.Did          `json:"senderDid" yaml:"senderDid"`
	Data      WithdrawFundsDoc `json:"data" yaml:"data"`
}

func (msg MsgWithdrawFunds) Type() string  { return TypeMsgWithdrawFunds }
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
	if !did.IsValidDid(msg.SenderDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "sender did is invalid")
	} else if !did.IsValidDid(msg.Data.ProjectDid) {
		return did.ErrorInvalidDid(DefaultCodespace, "project did is invalid")
	} else if !did.IsValidDid(msg.Data.RecipientDid) {
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

func (msg MsgWithdrawFunds) GetSignerDid() did.Did { return msg.Data.RecipientDid }
func (msg MsgWithdrawFunds) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgWithdrawFunds) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgWithdrawFunds) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(b)
}
