package models

import (
    "encoding/json"
    "errors"
)

// Model
type MarketCurrency struct {
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    IsDefault bool `json:"is_default"`
    // 
    MarketId string `json:"market_id"`
    // 
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model MarketCurrency) New(data []byte) *MarketCurrency {
    model.data = data
    return &model
}

func (model *MarketCurrency) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}