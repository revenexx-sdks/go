package models

import (
    "encoding/json"
    "errors"
)

// ContactInviteResponse model.
type ContactInviteResponse struct {
    // Who was invited.
    ContactId string `json:"contact_id"`
    // Always true when this answers — a failure to send is a 502, not a false
    // here.
    Invited bool `json:"invited"`
    // The company they were invited into.
    OrganizationId string `json:"organization_id"`

    // Used by Decode() method
    data []byte
}

func (model ContactInviteResponse) New(data []byte) *ContactInviteResponse {
    model.data = data
    return &model
}

func (model *ContactInviteResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}