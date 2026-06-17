package models

import (
    "encoding/json"
    "errors"
)

// Model
type CartOrderRequest struct {
    // External order reference from order management.
    OrderRef string `json:"order_ref"`

    // Used by Decode() method
    data []byte
}

func (model CartOrderRequest) New(data []byte) *CartOrderRequest {
    model.data = data
    return &model
}

func (model *CartOrderRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}