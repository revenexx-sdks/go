package models

import (
    "encoding/json"
    "errors"
)

// Model
type Greeting struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Locale string `json:"locale"`
    // 
    Message string `json:"message"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    Name string `json:"name"`
    // 
    TenantId string `json:"tenant_id"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Greeting) New(data []byte) *Greeting {
    model.data = data
    return &model
}

func (model *Greeting) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}