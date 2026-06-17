package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderCommentCreateRequest struct {
    // 
    Author string `json:"author"`
    // 
    Body string `json:"body"`
    // Default 'internal'.
    Visibility string `json:"visibility"`

    // Used by Decode() method
    data []byte
}

func (model OrderCommentCreateRequest) New(data []byte) *OrderCommentCreateRequest {
    model.data = data
    return &model
}

func (model *OrderCommentCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}