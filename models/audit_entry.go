package models

import (
    "encoding/json"
    "errors"
)

// AuditEntry model.
type AuditEntry struct {
    // 
    Action string `json:"action"`
    // 
    Changes []interface{} `json:"changes"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    ResourceId string `json:"resource_id"`
    // 
    ResourceKey string `json:"resource_key"`
    // 
    ResourceType string `json:"resource_type"`
    // 
    Subject string `json:"subject"`
    // 
    TenantId string `json:"tenant_id"`

    // Used by Decode() method
    data []byte
}

func (model AuditEntry) New(data []byte) *AuditEntry {
    model.data = data
    return &model
}

func (model *AuditEntry) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}