package models

import (
    "encoding/json"
    "errors"
)

// Model
type InventoryReleaseRequest struct {
    // The order whose active reservations are released.
    OrderRef string `json:"order_ref"`

    // Used by Decode() method
    data []byte
}

func (model InventoryReleaseRequest) New(data []byte) *InventoryReleaseRequest {
    model.data = data
    return &model
}

func (model *InventoryReleaseRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}