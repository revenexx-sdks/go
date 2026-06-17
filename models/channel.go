package models

import (
    "encoding/json"
    "errors"
)

// Model
type Channel struct {
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    IsDefault bool `json:"is_default"`
    // 
    Labels interface{} `json:"labels"`
    // 
    Name string `json:"name"`
    // 
    Position int `json:"position"`
    // 
    Status string `json:"status"`
    // 
    Type string `json:"type"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Channel) New(data []byte) *Channel {
    model.data = data
    return &model
}

func (model *Channel) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}