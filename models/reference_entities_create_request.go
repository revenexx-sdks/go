package models

import (
    "encoding/json"
    "errors"
)

// Model
type ReferenceEntitiesCreateRequest struct {
    // 
    Code string `json:"code"`
    // 
    Image string `json:"image"`
    // 
    Labels interface{} `json:"labels"`

    // Used by Decode() method
    data []byte
}

func (model ReferenceEntitiesCreateRequest) New(data []byte) *ReferenceEntitiesCreateRequest {
    model.data = data
    return &model
}

func (model *ReferenceEntitiesCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}