package models

import (
    "encoding/json"
    "errors"
)

// Model
type MarketTaxClass struct {
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    IsDefault bool `json:"is_default"`
    // 
    Labels interface{} `json:"labels"`
    // 
    MarketId string `json:"market_id"`
    // 
    Name string `json:"name"`
    // 
    Position int `json:"position"`
    // 
    Rate float64 `json:"rate"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model MarketTaxClass) New(data []byte) *MarketTaxClass {
    model.data = data
    return &model
}

func (model *MarketTaxClass) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}