package ixo1155

import "time"

type TokenId string
type Batch struct {
	Token_id TokenId
	Amt      uint64
}
type Expiration struct {
	AtHeight uint64
	AtTime   time.Time
	Never    interface{}
}

// SendFrom / SendFrom is a base message to move tokens
// / if `env.sender` is the owner or has sufficient pre-approval.
type SendFrom struct {
	From string
	/// If `to` is not contract `msg` should be `None`
	To       string
	Token_id TokenId
	Value    uint64
	/// `None` means don't call the receiver interface
	Msg []byte
}

// BatchSendFrom / BatchSendFrom is a base message to move multiple types of tokens in batch
// / if `env.sender` is the owner or has sufficient pre-approval.
type BatchSendFrom struct {
	From string
	/// if `to` is not contract `msg` should be `None`
	To    string
	Batch Batch
	/// `None` means don't call the receiver interface
	Msg []byte
}

// Mint / Mint is a base message to mint tokens.
type Mint struct {
	/// If `to` is not contract `msg` should be `None`
	To       string
	Token_id TokenId
	Value    uint64
	/// `None` means don't call the receiver interface
	Msg []byte
}

// BatchMint / BatchMint is a base message to mint multiple types of tokens in batch.
type BatchMint struct {
	/// If `to` is not contract `msg` should be `None`
	To    string
	Batch Batch
	/// `None` means don't call the receiver interface
	Msg []byte
}

// Burn / Burn is a base message to burn tokens.
type Burn struct {
	From     string
	Token_id TokenId
	Value    uint64
}

// BatchBurn / BatchBurn is a base message to burn multiple types of tokens in batch.
type BatchBurn struct {
	From  string
	Batch Batch
}

// ApproveAll / Allows operator to transfer / send any token from the owner's account.
// / If expiration is set then this allowance has a time/height limit
type ApproveAll struct {
	Operator string
	Expires  Expiration
}

// RevokeAll / Remove previously granted ApproveAll permission
type RevokeAll struct{ operator string }

//queries

// Balance #[returns(BalanceResponse)]
type Balance struct {
	Owner    string
	Token_id TokenId
}

// BatchBalance BatchBalance / Returns the current balance of the given address for a batch of tokens 0 if unset.
// #[returns(BatchBalanceResponse)]
type BatchBalance struct {
	Owner     string
	Token_ids []TokenId
}

// ApprovedForAll / List all operators that can access all of the owner's tokens.
// #[returns(ApprovedForAllResponse)]
type ApprovedForAll struct {
	Owner string
	/// unset or false will filter out expired approvals you must set to true to see them
	Include_expired bool
	Start_after     string
	Limit           uint32
}

// IsApprovedForAll / Query approved status `owner` granted to `operator`.
// #[returns(IsApprovedForAllResponse)]
type IsApprovedForAll struct {
	Owner    string
	Operator string
}

// TokenInfo / With MetaData Extension.
// / Query metadata of token
// #[returns(TokenInfoResponse)]
type TokenInfo struct{ Token_id TokenId }

// Tokens / With Enumerable extension.
// / Returns all tokens owned by the given address [] if unset.
// #[returns(TokensResponse)]
type Tokens struct {
	Owner       string
	Start_after string
	Limit       uint32
}

// AllTokens / With Enumerable extension.
// / Requires pagination. Lists all token_ids controlled by the contract.
// #[returns(TokensResponse)]
type AllTokens struct {
	Start_after string
	Limit       uint32
}
