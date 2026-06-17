package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderHoldRequest struct {
    // Why the order is blocked (shown on the shipping guard).
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model OrderHoldRequest) New(data []byte) *OrderHoldRequest {
    model.data = data
    return &model
}

func (model *OrderHoldRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}