package models

import (
    "encoding/json"
    "errors"
)

// Model
type ReferenceEntities struct {
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Image string `json:"image"`
    // 
    Labels interface{} `json:"labels"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model ReferenceEntities) New(data []byte) *ReferenceEntities {
    model.data = data
    return &model
}

func (model *ReferenceEntities) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}