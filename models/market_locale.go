package models

import (
    "encoding/json"
    "errors"
)

// Model
type MarketLocale struct {
    // 
    Code string `json:"code"`
    // 
    Country string `json:"country"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    IsDefault bool `json:"is_default"`
    // 
    Language string `json:"language"`
    // 
    MarketId string `json:"market_id"`
    // 
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model MarketLocale) New(data []byte) *MarketLocale {
    model.data = data
    return &model
}

func (model *MarketLocale) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}