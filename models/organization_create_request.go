package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrganizationCreateRequest struct {
    // Company name — mirrored to the platform team.
    Name string `json:"name"`
    // Free-form organization settings.
    Settings interface{} `json:"settings"`
    // Default 'active'.
    Status string `json:"status"`
    // 
    VatId string `json:"vat_id"`

    // Used by Decode() method
    data []byte
}

func (model OrganizationCreateRequest) New(data []byte) *OrganizationCreateRequest {
    model.data = data
    return &model
}

func (model *OrganizationCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}