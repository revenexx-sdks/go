package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderComment struct {
    // 
    Author string `json:"author"`
    // 
    Body string `json:"body"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    OrderId string `json:"order_id"`
    // 
    Visibility string `json:"visibility"`

    // Used by Decode() method
    data []byte
}

func (model OrderComment) New(data []byte) *OrderComment {
    model.data = data
    return &model
}

func (model *OrderComment) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}