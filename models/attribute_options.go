package models

import (
    "encoding/json"
    "errors"
)

// Model
type AttributeOptions struct {
    // 
    AttributeId string `json:"attribute_id"`
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Labels interface{} `json:"labels"`
    // 
    Position int `json:"position"`
    // 
    Swatch interface{} `json:"swatch"`

    // Used by Decode() method
    data []byte
}

func (model AttributeOptions) New(data []byte) *AttributeOptions {
    model.data = data
    return &model
}

func (model *AttributeOptions) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}