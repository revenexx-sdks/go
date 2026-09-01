package models

import (
    "encoding/json"
    "errors"
)

// TenantConfig model.
type TenantConfig struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    DefaultLocale string `json:"default_locale"`
    // 
    Defaults []interface{} `json:"defaults"`
    // 
    DeliveryReporting []interface{} `json:"delivery_reporting"`
    // 
    Locales []interface{} `json:"locales"`
    // 
    Product string `json:"product"`
    // 
    ProvisionedAt string `json:"provisioned_at"`
    // 
    QuietHours []interface{} `json:"quiet_hours"`
    // 
    Quotas []interface{} `json:"quotas"`
    // 
    RetentionDays int `json:"retention_days"`
    // 
    SupportEmail string `json:"support_email"`
    // 
    TenantId string `json:"tenant_id"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model TenantConfig) New(data []byte) *TenantConfig {
    model.data = data
    return &model
}

func (model *TenantConfig) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}