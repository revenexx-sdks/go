package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type AddressUpdateRequest struct {
    // 
    City string `json:"city"`
    // 
    Company string `json:"company"`
    // Owning contact (personal address).
    ContactId string `json:"contact_id"`
    // ISO 3166-1 alpha-2 code.
    Country string `json:"country"`
    // The default address of its owner and type.
    IsDefault bool `json:"is_default"`
    // Recipient name.
    Name string `json:"name"`
    // Owning organization (company address).
    OrganizationId string `json:"organization_id"`
    // 
    Phone string `json:"phone"`
    // 
    Region string `json:"region"`
    // 
    Street string `json:"street"`
    // 
    Street2 string `json:"street2"`
    // Default 'shipping'.
    Type string `json:"type"`
    // 
    Zip string `json:"zip"`

    // Used by Decode() method
    data []byte
}

func (model AddressUpdateRequest) New(data []byte) *AddressUpdateRequest {
    model.data = data
    return &model
}

func (model *AddressUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}