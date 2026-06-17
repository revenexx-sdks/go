package models

import (
    "encoding/json"
    "errors"
)

// Model
type Location struct {
    // 
    Address interface{} `json:"address"`
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Enabled bool `json:"enabled"`
    // 
    Id string `json:"id"`
    // 
    Labels interface{} `json:"labels"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    Name string `json:"name"`
    // 
    Priority int `json:"priority"`
    // 
    Type string `json:"type"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Location) New(data []byte) *Location {
    model.data = data
    return &model
}

func (model *Location) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}