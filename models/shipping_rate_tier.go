package models

import (
    "encoding/json"
    "errors"
)

// Model
type ShippingRateTier struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    FromValue float64 `json:"from_value"`
    // 
    Id string `json:"id"`
    // 
    MethodId string `json:"method_id"`
    // 
    Position int `json:"position"`
    // 
    Price float64 `json:"price"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model ShippingRateTier) New(data []byte) *ShippingRateTier {
    model.data = data
    return &model
}

func (model *ShippingRateTier) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}