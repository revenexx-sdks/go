package models

import (
    "encoding/json"
    "errors"
)

// Model
type InventoryAdjustRequest struct {
    // The corrections — quantities are SIGNED deltas (at most 200).
    Items []InventoryAdjustItem `json:"items"`
    // Adjusted location (default 'main').
    LocationCode string `json:"location_code"`
    // Mandatory audit reason — every adjustment is a ledger row.
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model InventoryAdjustRequest) New(data []byte) *InventoryAdjustRequest {
    model.data = data
    return &model
}

func (model *InventoryAdjustRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}