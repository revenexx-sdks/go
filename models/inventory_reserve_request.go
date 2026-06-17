package models

import (
    "encoding/json"
    "errors"
)

// Model
type InventoryReserveRequest struct {
    // Optional reservation expiry.
    ExpiresAt string `json:"expires_at"`
    // The items to reserve — all-or-nothing (at most 200).
    Items []InventoryStockItem `json:"items"`
    // The order this reservation belongs to.
    OrderRef string `json:"order_ref"`

    // Used by Decode() method
    data []byte
}

func (model InventoryReserveRequest) New(data []byte) *InventoryReserveRequest {
    model.data = data
    return &model
}

func (model *InventoryReserveRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}