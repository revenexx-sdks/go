package models

import (
    "encoding/json"
    "errors"
)

// ANewMatrixTierFromValuePriceOfTheMethodInThePath Model
type ShippingRateTierCreateRequest struct {
    // Tier threshold (default 0) — the tier with the highest from_value at or
    // below the measured value wins.
    FromValue float64 `json:"from_value"`
    // Sort order (default 0; bulk replace derives it from the array index).
    Position int `json:"position"`
    // Price of this tier (default 0).
    Price float64 `json:"price"`

    // Used by Decode() method
    data []byte
}

func (model ShippingRateTierCreateRequest) New(data []byte) *ShippingRateTierCreateRequest {
    model.data = data
    return &model
}

func (model *ShippingRateTierCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}