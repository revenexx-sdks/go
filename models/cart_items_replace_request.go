package models

import (
    "encoding/json"
    "errors"
)

// Model
type CartItemsReplaceRequest struct {
    // The complete new item set (set semantics).
    Items []CartItemCreateRequest `json:"items"`

    // Used by Decode() method
    data []byte
}

func (model CartItemsReplaceRequest) New(data []byte) *CartItemsReplaceRequest {
    model.data = data
    return &model
}

func (model *CartItemsReplaceRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}