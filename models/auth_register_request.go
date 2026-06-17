package models

import (
    "encoding/json"
    "errors"
)

// Model
type AuthRegisterRequest struct {
    // 
    Email string `json:"email"`
    // 
    FirstName string `json:"first_name"`
    // 
    LastName string `json:"last_name"`
    // BCP 47, e.g. de-DE
    Locale string `json:"locale"`
    // Join an existing organization.
    OrganizationId string `json:"organization_id"`
    // Found a new organization; the contact becomes its admin.
    OrganizationName string `json:"organization_name"`
    // 
    Password string `json:"password"`

    // Used by Decode() method
    data []byte
}

func (model AuthRegisterRequest) New(data []byte) *AuthRegisterRequest {
    model.data = data
    return &model
}

func (model *AuthRegisterRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}