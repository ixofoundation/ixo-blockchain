package cw20

import (
	"bytes"
	"encoding/json"
)

type InstantiateMarketingInfo struct {
	Project     string `json:"project"`
	Description string `json:"description"`
	Marketing   string `json:"marketing"`
	// Logo        Logo
}

type InstantiateMsg struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint32 `json:"decimals"`
	// InitialBalances []Cw20Coin
	// Mint            MinterResponse
	Marketing InstantiateMarketingInfo
}

func (m InstantiateMsg) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}

type Mint struct {
	Recipient string `json:"recipient"`
	Amount    uint   `json:"amount"`
}

func (m Mint) Marshal() ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(m); err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}
