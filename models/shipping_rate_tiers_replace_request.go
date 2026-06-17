package models

import (
    "encoding/json"
    "errors"
)

// Model
type ShippingRateTiersReplaceRequest struct {
    // The complete new tier set (set semantics) — positions are derived from
    // the array order.
    Tiers []ShippingRateTierReplaceItem `json:"tiers"`

    // Used by Decode() method
    data []byte
}

func (model ShippingRateTiersReplaceRequest) New(data []byte) *ShippingRateTiersReplaceRequest {
    model.data = data
    return &model
}

func (model *ShippingRateTiersReplaceRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}