package contracts

import (
	"bytes"
	"encoding/json"
)

type WasmMsgMint struct {
	Mint Mint `json:"mint"`
}

// / Mint a new NFT, can only be called by the contract minter
type Mint struct {
	/// Unique ID of the NFT
	TokenId string `json:"token_id"`
	/// The owner of the newly minter NFT
	Owner string `json:"owner"`
	/// Universal resource identifier for this NFT
	/// Should point to a JSON file that conforms to the ERC721
	/// Metadata JSON Schema
	TokenUri string `json:"token_uri"`
	/// Any custom extension used by this contract
	Extension json.RawMessage `json:"extension"`
}

// / Transfer is a base message to move a token to another account without triggering actions
type WasmMsgTransferNft struct {
	TransferNft TransferNft `json:"transfer_nft"`
}
type TransferNft struct {
	TokenId   string `json:"token_id"`
	Recipient string `json:"recipient"`
}

type WasmMsgInitiateNftContract struct {
	InstantiateMsg InitiateNftContract `json:"instantiate_msg"`
}
type InitiateNftContract struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	/// The minter is the only one who can create new NFTs.
	Minter string `json:"minter"`
}

// / Allows operator to transfer / send the token from the owner's account.
// / If expiration is set, then this allowance has a time/height limit
type WasmMsgApprove struct {
	ApproveNftTransfer ApproveNftTransfer `json:"approve"`
}
type ApproveNftTransfer struct {
	Spender string `json:"spender"`
	TokenId string `json:"token_id"`
	Expires string `json:"expires,omitempty"`
}

// / Return the owner of the given token, error if token does not exist
// #[returns(cw721::OwnerOfResponse)]
type WasmQueryOwnerOf struct {
	OwnerOf OwnerOf `json:"owner_of"`
}
type OwnerOf struct {
	TokenId string `json:"token_id"`
}

type OwnerOfResponse struct {
	Owner     string   `json:"owner"`
	Approvals []string `json:"approvals"`
}

func Marshal(value interface{}) ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(value); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}
