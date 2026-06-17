package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderItemsCancelRequest struct {
    // Acting user/system.
    CancelledBy string `json:"cancelled_by"`
    // 
    Positions []OrderCancelPosition `json:"positions"`
    // 
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model OrderItemsCancelRequest) New(data []byte) *OrderItemsCancelRequest {
    model.data = data
    return &model
}

func (model *OrderItemsCancelRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}