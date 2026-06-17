package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderEvent struct {
    // 
    Actor string `json:"actor"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Name string `json:"name"`
    // 
    OrderId string `json:"order_id"`
    // 
    Payload interface{} `json:"payload"`

    // Used by Decode() method
    data []byte
}

func (model OrderEvent) New(data []byte) *OrderEvent {
    model.data = data
    return &model
}

func (model *OrderEvent) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}