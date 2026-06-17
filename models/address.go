package models

import (
    "encoding/json"
    "errors"
)

// Model
type Address struct {
    // 
    City string `json:"city"`
    // 
    Company string `json:"company"`
    // 
    ContactId string `json:"contact_id"`
    // 
    Country string `json:"country"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    IsDefault bool `json:"is_default"`
    // 
    Name string `json:"name"`
    // 
    OrganizationId string `json:"organization_id"`
    // 
    Phone string `json:"phone"`
    // 
    Region string `json:"region"`
    // 
    Street string `json:"street"`
    // 
    Street2 string `json:"street2"`
    // 
    Type string `json:"type"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    Zip string `json:"zip"`

    // Used by Decode() method
    data []byte
}

func (model Address) New(data []byte) *Address {
    model.data = data
    return &model
}

func (model *Address) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}