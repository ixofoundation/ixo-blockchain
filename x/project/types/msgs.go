package types

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	didexported "github.com/ixofoundation/ixo-blockchain/x/did/exported"
	didtypes "github.com/ixofoundation/ixo-blockchain/x/did/types"
	ixotypes "github.com/ixofoundation/ixo-blockchain/x/ixo/types"
	"github.com/spf13/viper"
)

const (
	TypeMsgCreateProject       = "create-project"
	TypeMsgUpdateProjectStatus = "update-project-status"
	TypeMsgCreateAgent         = "create-agent"
	TypeMsgUpdateAgent         = "update-agent"
	TypeMsgCreateClaim         = "create-claim"
	TypeMsgCreateEvaluation    = "create-evaluation"
	TypeMsgWithdrawFunds       = "withdraw-funds"
	TypeMsgUpdateProjectDoc    = "update-project-doc"

	MsgCreateProjectTotalFee       = int64(1000000)
	MsgCreateProjectTransactionFee = int64(10000)
	// Project funding will be totalFee - transactionFee = 990000
)

var (
	_ ixotypes.IxoMsg = &MsgCreateProject{}
	_ ixotypes.IxoMsg = &MsgUpdateProjectStatus{}
	_ ixotypes.IxoMsg = &MsgCreateAgent{}
	_ ixotypes.IxoMsg = &MsgUpdateAgent{}
	_ ixotypes.IxoMsg = &MsgCreateClaim{}
	_ ixotypes.IxoMsg = &MsgCreateEvaluation{}
	_ ixotypes.IxoMsg = &MsgWithdrawFunds{}
	_ ixotypes.IxoMsg = &MsgUpdateProjectDoc{}
)

//type MsgCreateProject struct {
//	TxHash     string          `json:"txHash" yaml:"txHash"`
//	SenderDid  didexported.Did         `json:"senderDid" yaml:"senderDid"`
//	ProjectDid didexported.Did         `json:"projectDid" yaml:"projectDid"`
//	PubKey     string          `json:"pubKey" yaml:"pubKey"`
//	Data       json.RawMessage `json:"data" yaml:"data"`
//}

func NewMsgCreateProject(senderDid didexported.Did, projectData json.RawMessage,
	projectDid didexported.Did, pubKey string) *MsgCreateProject {
	return &MsgCreateProject{
		TxHash:     "",
		SenderDid:  senderDid,
		ProjectDid: projectDid,
		PubKey:     pubKey,
		Data:       projectData,
	}
}

func (msg MsgCreateProject) ToStdSignMsg(fee int64) legacytx.StdSignMsg {
	chainID := viper.GetString(flags.FlagChainID)
	accNum, accSeq := uint64(0), uint64(0)
	stdFee := legacytx.NewStdFee(0, sdk.NewCoins(sdk.NewCoin(
		ixotypes.IxoNativeToken, sdk.NewInt(fee))))
	memo := viper.GetString(flags.FlagMemo)

	return legacytx.StdSignMsg{
		ChainID:       chainID,
		AccountNumber: accNum,
		Sequence:      accSeq,
		Fee:           stdFee,
		Msgs:          []sdk.Msg{&msg},
		Memo:          memo,
	}
}

func (msg MsgCreateProject) Type() string { return TypeMsgCreateProject }

func (msg MsgCreateProject) Route() string { return RouterKey }

func (msg MsgCreateProject) ValidateBasic() error {
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
		return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "failed to unmarshal project data map")
	}

	// Check that project DID matches the PubKey
	unprefixedDid := didexported.UnprefixedDid(msg.ProjectDid)
	expectedUnprefixedDid := didexported.UnprefixedDidFromPubKey(msg.PubKey)
	if unprefixedDid != expectedUnprefixedDid {
		return sdkerrors.Wrapf(didtypes.ErrDidPubKeyMismatch,
			"did not deducable from pubKey; expected: %s received: %s",
			expectedUnprefixedDid, unprefixedDid)
	}

	return nil
}

func (msg MsgCreateProject) GetSignerDid() didexported.Did { return msg.ProjectDid }
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

//func (msg MsgCreateProject) GetPubKey() string { return msg.PubKey }

func (msg MsgCreateProject) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

//type MsgUpdateProjectStatus struct {
//	TxHash     string                 `json:"txHash" yaml:"txHash"`
//	SenderDid  didexported.Did                `json:"senderDid" yaml:"senderDid"`
//	ProjectDid didexported.Did                `json:"projectDid" yaml:"projectDid"`
//	Data       UpdateProjectStatusDoc `json:"data" yaml:"data"`
//}

func NewMsgUpdateProjectStatus(senderDid didexported.Did, updateProjectStatusDoc UpdateProjectStatusDoc, projectDid didexported.Did) *MsgUpdateProjectStatus {
	return &MsgUpdateProjectStatus{
		TxHash:     "",
		SenderDid:  senderDid,
		ProjectDid: projectDid,
		Data:       updateProjectStatusDoc,
	}
}

func (msg MsgUpdateProjectStatus) Type() string  { return TypeMsgUpdateProjectStatus }
func (msg MsgUpdateProjectStatus) Route() string { return RouterKey }

func (msg MsgUpdateProjectStatus) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// TODO: perform some checks on the Data (of type UpdateProjectStatusDoc)

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.ProjectDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "project DID is invalid")
	} else if !didtypes.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender DID is invalid")
	}

	// IsValidProgressionFrom checked by the handler

	return nil
}

func (msg MsgUpdateProjectStatus) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUpdateProjectStatus) GetSignerDid() didexported.Did { return msg.ProjectDid }
func (msg MsgUpdateProjectStatus) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

//type MsgCreateAgent struct {
//	TxHash     string         `json:"txHash" yaml:"txHash"`
//	SenderDid  didexported.Did        `json:"senderDid" yaml:"senderDid"`
//	ProjectDid didexported.Did        `json:"projectDid" yaml:"projectDid"`
//	Data       CreateAgentDoc `json:"data" yaml:"data"`
//}

func NewMsgCreateAgent(txHash string, senderDid didexported.Did, createAgentDoc CreateAgentDoc, projectDid didexported.Did) *MsgCreateAgent {
	return &MsgCreateAgent{
		ProjectDid: projectDid,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createAgentDoc,
	}
}

func (msg MsgCreateAgent) Type() string  { return TypeMsgCreateAgent }
func (msg MsgCreateAgent) Route() string { return RouterKey }
func (msg MsgCreateAgent) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// TODO: perform some checks on the Data (of type CreateAgentDoc)

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.ProjectDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "project DID is invalid")
	} else if !didtypes.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender DID is invalid")
	} else if !didtypes.IsValidDid(msg.Data.AgentDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "agent DID is invalid")
	}

	return nil
}

func (msg MsgCreateAgent) GetSignerDid() didexported.Did { return msg.ProjectDid }
func (msg MsgCreateAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreateAgent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateAgent) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}

//type MsgUpdateAgent struct {
//	TxHash     string         `json:"txHash" yaml:"txHash"`
//	SenderDid  didexported.Did        `json:"senderDid" yaml:"senderDid"`
//	ProjectDid didexported.Did        `json:"projectDid" yaml:"projectDid"`
//	Data       UpdateAgentDoc `json:"data" yaml:"data"`
//}

func NewMsgUpdateAgent(txHash string, senderDid didexported.Did, updateAgentDoc UpdateAgentDoc, projectDid didexported.Did) *MsgUpdateAgent {
	return &MsgUpdateAgent{
		ProjectDid: projectDid,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       updateAgentDoc,
	}
}

func (msg MsgUpdateAgent) Type() string  { return TypeMsgUpdateAgent }
func (msg MsgUpdateAgent) Route() string { return RouterKey }
func (msg MsgUpdateAgent) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// TODO: perform some checks on the Data (of type UpdateAgentDoc)

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.ProjectDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "project DID is invalid")
	} else if !didtypes.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender DID is invalid")
	} else if !didtypes.IsValidDid(msg.Data.Did) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "agent DID is invalid")
	}

	return nil
}

func (msg MsgUpdateAgent) GetSignerDid() didexported.Did { return msg.ProjectDid }
func (msg MsgUpdateAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgUpdateAgent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUpdateAgent) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(b)
}

//type MsgCreateClaim struct {
//	TxHash     string         `json:"txHash" yaml:"txHash"`
//	SenderDid  didexported.Did        `json:"senderDid" yaml:"senderDid"`
//	ProjectDid didexported.Did        `json:"projectDid" yaml:"projectDid"`
//	Data       CreateClaimDoc `json:"data" yaml:"data"`
//}

func NewMsgCreateClaim(txHash string, senderDid didexported.Did, createClaimDoc CreateClaimDoc, projectDid didexported.Did) *MsgCreateClaim {
	return &MsgCreateClaim{
		ProjectDid: projectDid,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createClaimDoc,
	}
}

func (msg MsgCreateClaim) Type() string  { return TypeMsgCreateClaim }
func (msg MsgCreateClaim) Route() string { return RouterKey }

func (msg MsgCreateClaim) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Data.ClaimId, "ClaimID"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.Data.ClaimTemplateId, "ClaimTemplateID"); !valid {
		return err
	}

	// TODO: perform some additional checks on the Data (of type CreateClaimDoc)

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.ProjectDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "project DID is invalid")
	} else if !didtypes.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender DID is invalid")
	}

	return nil
}

func (msg MsgCreateClaim) GetSignerDid() didexported.Did { return msg.ProjectDid }
func (msg MsgCreateClaim) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreateClaim) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateClaim) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(b)
}

//type MsgCreateEvaluation struct {
//	TxHash     string              `json:"txHash" yaml:"txHash"`
//	SenderDid  didexported.Did             `json:"senderDid" yaml:"senderDid"`
//	ProjectDid didexported.Did             `json:"projectDid" yaml:"projectDid"`
//	Data       CreateEvaluationDoc `json:"data" yaml:"data"`
//}

func NewMsgCreateEvaluation(txHash string, senderDid didexported.Did, createEvaluationDoc CreateEvaluationDoc, projectDid didexported.Did) *MsgCreateEvaluation {
	return &MsgCreateEvaluation{
		ProjectDid: projectDid,
		TxHash:     txHash,
		SenderDid:  senderDid,
		Data:       createEvaluationDoc,
	}
}

func (msg MsgCreateEvaluation) Type() string  { return TypeMsgCreateEvaluation }
func (msg MsgCreateEvaluation) Route() string { return RouterKey }

func (msg MsgCreateEvaluation) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// TODO: perform some checks on the Data (of type CreateEvaluationDoc)

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.ProjectDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "project DID is invalid")
	} else if !didtypes.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender DID is invalid")
	}

	return nil
}

func (msg MsgCreateEvaluation) GetSignerDid() didexported.Did { return msg.ProjectDid }
func (msg MsgCreateEvaluation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgCreateEvaluation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateEvaluation) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(b)
}

//type MsgWithdrawFunds struct {
//	SenderDid didexported.Did          `json:"senderDid" yaml:"senderDid"`
//	Data      WithdrawFundsDoc `json:"data" yaml:"data"`
//}

func NewMsgWithdrawFunds(senderDid didexported.Did, data WithdrawFundsDoc) *MsgWithdrawFunds {
	return &MsgWithdrawFunds{
		SenderDid: senderDid,
		Data:      data,
	}
}

func (msg MsgWithdrawFunds) Type() string  { return TypeMsgWithdrawFunds }
func (msg MsgWithdrawFunds) Route() string { return RouterKey }

func (msg MsgWithdrawFunds) ValidateBasic() error {
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
	if !didtypes.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender DID is invalid")
	} else if !didtypes.IsValidDid(msg.Data.ProjectDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "project DID is invalid")
	} else if !didtypes.IsValidDid(msg.Data.RecipientDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "recipient DID is invalid")
	}

	// Check that the sender is also the recipient
	if msg.SenderDid != msg.Data.RecipientDid {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender did must match recipient did")
	}

	// Check that amount is positive
	if !msg.Data.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "amount should be positive")
	}

	return nil
}

func (msg MsgWithdrawFunds) GetSignerDid() didexported.Did { return msg.Data.RecipientDid }
func (msg MsgWithdrawFunds) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}

func (msg MsgWithdrawFunds) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdrawFunds) String() string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func NewMsgUpdateProjectDoc(senderDid didexported.Did, projectData json.RawMessage, projectDid didexported.Did) *MsgUpdateProjectDoc {
	return &MsgUpdateProjectDoc{
		TxHash:     "",
		SenderDid:  senderDid,
		ProjectDid: projectDid,
		Data:       projectData,
	}
}

func (msg MsgUpdateProjectDoc) Type() string  { return TypeMsgUpdateProjectDoc }
func (msg MsgUpdateProjectDoc) Route() string { return RouterKey }

func (msg MsgUpdateProjectDoc) ValidateBasic() error {
	// Check that not empty
	if valid, err := CheckNotEmpty(msg.ProjectDid, "ProjectDid"); !valid {
		return err
	} else if valid, err := CheckNotEmpty(msg.SenderDid, "SenderDid"); !valid {
		return err
	}

	// Check that data marshallable to map[string]json.RawMessage
	var dataMap ProjectDataMap
	err := json.Unmarshal(msg.Data, &dataMap)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "failed to unmarshal project data map")
	}

	// Check that DIDs valid
	if !didtypes.IsValidDid(msg.ProjectDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "project DID is invalid")
	} else if !didtypes.IsValidDid(msg.SenderDid) {
		return sdkerrors.Wrap(didtypes.ErrInvalidDid, "sender DID is invalid")
	}

	return nil
}

func (msg MsgUpdateProjectDoc) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUpdateProjectDoc) GetSignerDid() didexported.Did { return msg.ProjectDid }
func (msg MsgUpdateProjectDoc) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{nil} // not used in signature verification in ixo AnteHandler
}