package models

import (
    "encoding/json"
    "errors"
)

// Suppression model.
type Suppression struct {
    // 
    Address string `json:"address"`
    // 
    AddressHash string `json:"address_hash"`
    // 
    Channel string `json:"channel"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    ExpiresAt string `json:"expires_at"`
    // 
    Id string `json:"id"`
    // 
    Note string `json:"note"`
    // 
    Reason string `json:"reason"`
    // 
    Scope string `json:"scope"`
    // 
    Source string `json:"source"`
    // 
    TenantId string `json:"tenant_id"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Suppression) New(data []byte) *Suppression {
    model.data = data
    return &model
}

func (model *Suppression) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}