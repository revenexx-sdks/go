package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderCancellation struct {
    // 
    CancelledBy string `json:"cancelled_by"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    OrderId string `json:"order_id"`
    // 
    Positions interface{} `json:"positions"`
    // 
    Reason string `json:"reason"`
    // 
    Scope string `json:"scope"`

    // Used by Decode() method
    data []byte
}

func (model OrderCancellation) New(data []byte) *OrderCancellation {
    model.data = data
    return &model
}

func (model *OrderCancellation) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}