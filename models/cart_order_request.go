package models

import (
    "encoding/json"
    "errors"
)

// CartOrderRequest model.
type CartOrderRequest struct {
    // The order number this cart becomes, in order management's own numbering.
    // Stored on the cart — filtering on it is how anyone gets from an order
    // back to the cart behind it — and it is also the reference the stock
    // reservation is booked under. Omit it and the cart id is used for the
    // reservation instead.
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