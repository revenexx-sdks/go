package models

import (
    "encoding/json"
    "errors"
)

// AMatrixTierOfTheNewSetFromValuePriceNullFallsBackTo0PositionDerivesFromTheArrayOrder
// Model
type ShippingRateTierReplaceItem struct {
    // Tier threshold (default 0) — the tier with the highest from_value at or
    // below the measured value wins.
    FromValue float64 `json:"from_value"`
    // Ignored — derived from the array index.
    Position int `json:"position"`
    // Price of this tier (default 0).
    Price float64 `json:"price"`

    // Used by Decode() method
    data []byte
}

func (model ShippingRateTierReplaceItem) New(data []byte) *ShippingRateTierReplaceItem {
    model.data = data
    return &model
}

func (model *ShippingRateTierReplaceItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}