package models

import (
    "encoding/json"
    "errors"
)

// Model
type Contact struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    Email string `json:"email"`
    // 
    ExternalUserId string `json:"external_user_id"`
    // 
    FirstName string `json:"first_name"`
    // 
    Id string `json:"id"`
    // 
    IsPrimary bool `json:"is_primary"`
    // 
    LastName string `json:"last_name"`
    // 
    Locale string `json:"locale"`
    // 
    OrganizationId string `json:"organization_id"`
    // 
    Phone string `json:"phone"`
    // 
    Role string `json:"role"`
    // 
    Status string `json:"status"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Contact) New(data []byte) *Contact {
    model.data = data
    return &model
}

func (model *Contact) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}