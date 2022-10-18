package contracts

import (
	"bytes"
	"encoding/json"
)

type Mint struct {
	TokenId  string `json:"token_id"`
	Owner    string `json:"owner"`
	TokenUrl string `json:"token_url"`
}

type WasmMintNftMessage struct {
	Mint Mint `json:"mint"`
}

func (m Mint) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

func (m WasmMintNftMessage) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}
