package models

import (
    "encoding/json"
    "errors"
)

// Model
type PaymentProvider struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    Credentials interface{} `json:"credentials"`
    // 
    Enabled bool `json:"enabled"`
    // 
    Id string `json:"id"`
    // 
    Name string `json:"name"`
    // 
    Options interface{} `json:"options"`
    // 
    Provider string `json:"provider"`
    // 
    TestMode bool `json:"test_mode"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    WebhookSecret string `json:"webhook_secret"`

    // Used by Decode() method
    data []byte
}

func (model PaymentProvider) New(data []byte) *PaymentProvider {
    model.data = data
    return &model
}

func (model *PaymentProvider) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}