package models

import (
    "encoding/json"
    "errors"
)

// Model
type Categories struct {
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Labels interface{} `json:"labels"`
    // 
    ParentId string `json:"parent_id"`
    // 
    Path string `json:"path"`
    // 
    Position int `json:"position"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    Values interface{} `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model Categories) New(data []byte) *Categories {
    model.data = data
    return &model
}

func (model *Categories) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}