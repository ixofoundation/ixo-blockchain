package cw20

import (
	"bytes"
	"encoding/json"
	"time"
)

func Marshal(value interface{}) ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(value); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type Expiration struct {
	AtHeight string
	AtTime   time.Time
	Never    interface{}
}

// -----------------------------------------------------------------------------
// MESSAGES
// -----------------------------------------------------------------------------
type InstantiateMsg struct {
	/// The minter is the only one who can create new tokens.
	Minter string `json:"minter"`
}

type WasmTransfer struct {
	Transfer Transfer `json:"transfer"`
}

// / Transfer is a base message to move tokens to another account without triggering actions
type Transfer struct {
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

type WasmBurn struct {
	Burn Burn `json:"burn"`
}

// / Burn is a base message to destroy tokens forever
type Burn struct {
	Amount string `json:"amount"`
}

type WasmSend struct {
	Send Send `json:"send"`
}

// / Send is a base message to transfer tokens to a contract and trigger an action
// / on the receiving contract.
type Send struct {
	Contract string `json:"contract"`
	Amount   string `json:"amount"`
	Msg      []byte `json:"msg"`
}

type WasmIncreaseAllowance struct {
	IncreaseAllowance IncreaseAllowance `json:"increase_allowance"`
}

// / Only with "approval" extension. Allows spender to access an additional amount tokens
// / from the owner's (env.sender) account. If expires is Some(), overwrites current allowance
// / expiration with this one.
type IncreaseAllowance struct {
	Spender string `json:"spender"`
	Amount  string `json:"amount"`
	Expires Expiration
}

type WasmDecreaseAllowance struct {
	DecreaseAllowance DecreaseAllowance `json:"decrease_allowance"`
}

// / Only with "approval" extension. Lowers the spender's access of tokens
// / from the owner's (env.sender) account by amount. If expires is Some(), overwrites current
// / allowance expiration with this one.
type DecreaseAllowance struct {
	Spender string `json:"spender"`
	Amount  string `json:"amount"`
	Expires Expiration
}

type WasmTransferFrom struct {
	TransferFrom TransferFrom `json:"transfer_from"`
}

// / Only with "approval" extension. Transfers amount tokens from owner -> recipient
// / if `env.sender` has sufficient pre-approval.
type TransferFrom struct {
	Owner     string `json:"owner"`
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

type WasmSendFrom struct {
	SendFrom SendFrom `json:"send_from"`
}

// / Only with "approval" extension. Sends amount tokens from owner -> contract
// / if `env.sender` has sufficient pre-approval.
type SendFrom struct {
	Owner    string `json:"owner"`
	Contract string `json:"contract"`
	Amount   string `json:"amount"`
	Msg      []byte `json:"msg"`
}

type WasmBurnFrom struct {
	BurnFrom BurnFrom `json:"burn_from"`
}

// / Only with "approval" extension. Destroys tokens forever
type BurnFrom struct {
	Owner  string `json:"owner"`
	Amount string `json:"amount"`
}

type WasmMint struct {
	Mint Mint `json:"mint"`
}

// / Only with the "mintable" extension. If authorized, creates amount new tokens
// / and adds to the recipient balance.
type Mint struct {
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

// -----------------------------------------------------------------------------
// QUERIES
// -----------------------------------------------------------------------------
type WasmBalance struct {
	Balance Balance `json:"balance"`
}

// / Returns the current balance of the given address, 0 if unset.
// / Return type: BalanceResponse.
type Balance struct {
	Address string `json:"address"`
}

type WasmTokenInfo struct {
	TokenInfo TokenInfo `json:"token_info"`
}

// / Returns metadata on the contract - name, decimals, supply, etc.
// / Return type: TokenInfoResponse.
type TokenInfo struct{}

type WasmAllowance struct {
	Allowance Allowance `json:"allowance"`
}

// / Only with "allowance" extension.
// / Returns how much spender can use from owner account, 0 if unset.
// / Return type: AllowanceResponse.
type Allowance struct {
	Owner   string `json:"owner"`
	Spender string `json:"spender"`
}
