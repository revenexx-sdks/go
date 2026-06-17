package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValueExternalTeamIdIsMirrorManagedAndIgnored
// Model
type OrganizationUpdateRequest struct {
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

func (model OrganizationUpdateRequest) New(data []byte) *OrganizationUpdateRequest {
    model.data = data
    return &model
}

func (model *OrganizationUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}