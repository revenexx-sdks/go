package models

import (
    "encoding/json"
    "errors"
)

// Model
type AttributeGroupsCreateRequest struct {
    // 
    Code string `json:"code"`
    // 
    Labels interface{} `json:"labels"`
    // 
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model AttributeGroupsCreateRequest) New(data []byte) *AttributeGroupsCreateRequest {
    model.data = data
    return &model
}

func (model *AttributeGroupsCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}