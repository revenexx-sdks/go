package models

import (
    "encoding/json"
    "errors"
)

// Model
type AttributeOptionsCreateRequest struct {
    // 
    AttributeId string `json:"attribute_id"`
    // 
    Code string `json:"code"`
    // 
    Labels interface{} `json:"labels"`
    // 
    Position int `json:"position"`
    // 
    Swatch interface{} `json:"swatch"`

    // Used by Decode() method
    data []byte
}

func (model AttributeOptionsCreateRequest) New(data []byte) *AttributeOptionsCreateRequest {
    model.data = data
    return &model
}

func (model *AttributeOptionsCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}