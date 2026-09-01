package models

import (
    "encoding/json"
    "errors"
)

// Binding model.
type Binding struct {
    // 
    Channel string `json:"channel"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Enabled bool `json:"enabled"`
    // 
    EventTopic string `json:"event_topic"`
    // 
    FallbackOrder int `json:"fallback_order"`
    // 
    Id string `json:"id"`
    // 
    Locale string `json:"locale"`
    // 
    Recipient string `json:"recipient"`
    // 
    TemplateKey string `json:"template_key"`
    // 
    TenantId string `json:"tenant_id"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Binding) New(data []byte) *Binding {
    model.data = data
    return &model
}

func (model *Binding) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}