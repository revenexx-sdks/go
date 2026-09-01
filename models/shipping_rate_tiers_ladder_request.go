package models

import (
    "encoding/json"
    "errors"
)

// ShippingRateTiersLadderRequest An evenly-stepped tier table. Tiers are
// generated at from_value, from_value+step, … up to to_value; each costs
// step_price more than the one before.
type ShippingRateTiersLadderRequest struct {
    // Price of the first tier.
    BasePrice float64 `json:"base_price"`
    // First tier threshold (default 0), in the method's matrix measure.
    FromValue float64 `json:"from_value"`
    // Replace the whole table (default true) or append to it.
    Replace bool `json:"replace"`
    // Distance between two tiers. Must be > 0.
    Step float64 `json:"step"`
    // Added to each subsequent tier (default 0). A negative value is allowed as
    // long as no tier ends up below 0.
    StepPrice float64 `json:"step_price"`
    // Last tier threshold. The final tier keeps applying above it — a matrix
    // has no upper bound. Must be >= from_value.
    ToValue float64 `json:"to_value"`

    // Used by Decode() method
    data []byte
}

func (model ShippingRateTiersLadderRequest) New(data []byte) *ShippingRateTiersLadderRequest {
    model.data = data
    return &model
}

func (model *ShippingRateTiersLadderRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}