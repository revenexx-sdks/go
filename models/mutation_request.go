package models

import (
    "encoding/json"
    "errors"
)

// Model
type MutationRequest struct {
    // 
    Langcode string `json:"langcode"`
    // 
    Payload interface{} `json:"payload"`
    // Mutation plugin id (add, move, delete, duplicate, update_field_value, ...).
    Plugin string `json:"plugin"`

    // Used by Decode() method
    data []byte
}

func (model MutationRequest) New(data []byte) *MutationRequest {
    model.data = data
    return &model
}

func (model *MutationRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}