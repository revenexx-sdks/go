package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValueExternalUserIdIsMirrorManagedAndIgnored
// Model
type ContactUpdateRequest struct {
    // 
    Email string `json:"email"`
    // 
    FirstName string `json:"first_name"`
    // The primary contact of its organization.
    IsPrimary bool `json:"is_primary"`
    // 
    LastName string `json:"last_name"`
    // BCP 47, e.g. de-DE
    Locale string `json:"locale"`
    // Owning organization — membership is mirrored to the platform team.
    OrganizationId string `json:"organization_id"`
    // 
    Phone string `json:"phone"`
    // Default 'buyer' — also the team role on the platform mirror.
    Role string `json:"role"`
    // Default 'invited' on create.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model ContactUpdateRequest) New(data []byte) *ContactUpdateRequest {
    model.data = data
    return &model
}

func (model *ContactUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}