package ixo1155

import (
	"bytes"
	"encoding/json"
	"time"
)

type InstantiateMsg struct {
	/// The minter is the only one who can create new tokens.
	Minter string `json:"minter"`
}

// Must be [Token_id, Amount, Uri]
type Batch []string

type Expiration struct {
	AtHeight string
	AtTime   time.Time
	Never    interface{}
}

type WasmSendFrom struct {
	SendFrom SendFrom `json:"send_from"`
}

// SendFrom / SendFrom is a base message to move tokens
// / if `env.sender` is the owner or has sufficient pre-approval.
type SendFrom struct {
	From string `json:"from"`
	/// If `to` is not contract `msg` should be `None`
	To       string `json:"to"`
	Token_id string `json:"token_id"`
	Value    string `json:"value"`
	/// `None` means don't call the receiver interface
	Msg []byte `json:"msg"`
}

type WasmBatchSendFrom struct {
	BatchSendFrom BatchSendFrom `json:"batch_send_from"`
}

// BatchSendFrom / BatchSendFrom is a base message to move multiple types of tokens in batch
// / if `env.sender` is the owner or has sufficient pre-approval.
type BatchSendFrom struct {
	From string `json:"from"`
	/// if `to` is not contract `msg` should be `None`
	To    string  `json:"to"`
	Batch []Batch `json:"batch"`
	/// `None` means don't call the receiver interface
	Msg []byte `json:"msg"`
}

type WasmMsgMint struct {
	Mint Mint `json:"mint"`
}

// Mint / Mint is a base message to mint tokens.
type Mint struct {
	/// If `to` is not contract `msg` should be `None`
	To      string `json:"to"`
	TokenId string `json:"token_id"`
	Value   string `json:"value"`
	Uri     string `json:"uri"`
	/// `None` means don't call the receiver interface
	Msg []byte `json:"msg,omitempty"`
}

type WasmMsgBatchMint struct {
	BatchMint BatchMint `json:"batch_mint"`
}

// BatchMint / BatchMint is a base message to mint multiple types of tokens in batch.
type BatchMint struct {
	/// If `to` is not contract `msg` should be `None`
	To    string  `json:"to"`
	Batch []Batch `json:"batch"`
	/// `None` means don't call the receiver interface
	Msg []byte `json:"msg"`
}

type WasmMsgBurn struct {
	Burn Burn `json:"burn"`
}

// Burn / Burn is a base message to burn tokens.
type Burn struct {
	From     string `json:"from"`
	Token_id string `json:"token_id"`
	Value    string `json:"value"`
}

type WasmMsgBatchBurn struct {
	BatchBurn BatchBurn `json:"batch_burn"`
}

// BatchBurn / BatchBurn is a base message to burn multiple types of tokens in batch.
type BatchBurn struct {
	From  string  `json:"from"`
	Batch []Batch `json:"batch"`
}

type WasmMsgApproveAll struct {
	ApproveAll ApproveAll `json:"approve_all"`
}

// ApproveAll / Allows operator to transfer / send any token from the owner's account.
// / If expiration is set then this allowance has a time/height limit
type ApproveAll struct {
	Operator string     `json:"operator"`
	Expires  Expiration `json:"expires"`
}

type WasmMsgRevokeAll struct {
	RevokeAll RevokeAll `json:"revoke_all"`
}

// RevokeAll / Remove previously granted ApproveAll permission
type RevokeAll struct {
	Operator string `json:"operator"`
}

//queries

type WasmMsgBalance struct {
	Balance Balance `json:"balance"`
}

// Balance #[returns(BalanceResponse)]
type Balance struct {
	Owner    string `json:"owner"`
	Token_id string `json:"token_id"`
}

type WasmMsgBatchBalance struct {
	BatchBalance BatchBalance `json:"batch_balance"`
}

// BatchBalance BatchBalance / Returns the current balance of the given address for a batch of tokens 0 if unset.
// #[returns(BatchBalanceResponse)]
type BatchBalance struct {
	Owner     string   `json:"owner"`
	Token_ids []string `json:"token_ids"`
}

type WasmMsgApprovedForAll struct {
	ApprovedForAll ApprovedForAll `json:"approved_for_all"`
}

// ApprovedForAll / List all operators that can access all of the owner's tokens.
// #[returns(ApprovedForAllResponse)]
type ApprovedForAll struct {
	Owner string `json:"owner"`
	/// unset or false will filter out expired approvals you must set to true to see them
	Include_expired bool   `json:"include_expired"`
	Start_after     string `json:"start_after"`
	Limit           string `json:"limit"`
}

type WasmMsgIsApprovedForAll struct {
	IsApprovedForAll IsApprovedForAll `json:"is_approved_for_all"`
}

// IsApprovedForAll / Query approved status `owner` granted to `operator`.
// #[returns(IsApprovedForAllResponse)]
type IsApprovedForAll struct {
	Owner    string `json:"owner"`
	Operator string `json:"operator"`
}

type WasmMsgTokenInfo struct {
	TokenInfo TokenInfo `json:"token_info"`
}

// TokenInfo / With MetaData Extension.
// / Query metadata of token
// #[returns(TokenInfoResponse)]
type TokenInfo struct {
	Token_id string `json:"token_id"`
}

type WasmMsgTokens struct {
	Tokens Tokens `json:"tokens"`
}

// Tokens / With Enumerable extension.
// / Returns all tokens owned by the given address [] if unset.
// #[returns(TokensResponse)]
type Tokens struct {
	Owner       string `json:"owner"`
	Start_after string `json:"start_after"`
	Limit       string `json:"limit"`
}

type WasmMsgAllTokens struct {
	AllTokens AllTokens `json:"all_tokens"`
}

// AllTokens / With Enumerable extension.
// / Requires pagination. Lists all token_ids controlled by the contract.
// #[returns(TokensResponse)]
type AllTokens struct {
	Start_after string `json:"start_after"`
	Limit       string `json:"limit"`
}

func Marshal(value interface{}) ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(value); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}
