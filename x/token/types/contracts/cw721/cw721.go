package cw721

import (
	"bytes"
	"encoding/json"
)

type Mint struct {
	TokenId   string          `json:"token_id"`
	Owner     string          `json:"owner"`
	TokenUri  string          `json:"token_uri"`
	Extension json.RawMessage `json:"extension"`
}

func (m Mint) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type WasmMsgMint struct {
	Mint Mint `json:"mint"`
}

func (m WasmMsgMint) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type TransferNft struct {
	TokenId   string `json:"token_id"`
	Recipient string `json:"recipient"`
}

func (m TransferNft) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type WasmMsgTransferNft struct {
	TransferNft TransferNft `json:"transfer_nft"`
}

func (m WasmMsgTransferNft) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type InitiateNftContract struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Minter string `json:"minter"`
}

func (m InitiateNftContract) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type WasmMsgInitiateNftContract struct {
	InstantiateMsg InitiateNftContract `json:"instantiate_msg"`
}

func (m WasmMsgInitiateNftContract) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type ApproveNftTransfer struct {
	Spender string `json:"spender"`
	TokenId string `json:"token_id"`
	Expires string `json:"expires,omitempty"`
}

func (m ApproveNftTransfer) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type WasmMsgApprove struct {
	ApproveNftTransfer ApproveNftTransfer `json:"approve"`
}

func (m WasmMsgApprove) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type QueryNft struct {
	Sender  string `json:"sender"`
	TokenId string `json:"token_id"`
	Expires string `json:"expires,omitempty"`
}

func (m QueryNft) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type WasmMsgQueryNft struct {
	QueryNft QueryNft `json:"approve"`
}

func (m WasmMsgQueryNft) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}
