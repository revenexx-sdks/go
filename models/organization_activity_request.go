package models

import (
    "encoding/json"
    "errors"
)

// OrganizationActivityRequest model.
type OrganizationActivityRequest struct {
    // Who logged it (operator id or email). Free text; this app does not resolve
    // it.
    Actor string `json:"actor"`
    // The person dealt with. Must be a contact of this organization.
    ContactId string `json:"contact_id"`
    // What happened. 'system' is deliberately NOT accepted — those rows are the
    // registration decision trail and are written by the approve/reject routes.
    // Default 'note'.
    Kind string `json:"kind"`
    // The long form. Stored inside the event payload as `note`, not as a column
    // of its own.
    Note string `json:"note"`
    // When it actually happened. Defaults to now — a call logged on Monday
    // about Friday should say Friday.
    OccurredAt string `json:"occurred_at"`
    // One line a person can scan in a timeline. Required — an entry nobody can
    // read at a glance is not worth the row.
    Subject string `json:"subject"`

    // Used by Decode() method
    data []byte
}

func (model OrganizationActivityRequest) New(data []byte) *OrganizationActivityRequest {
    model.data = data
    return &model
}

func (model *OrganizationActivityRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}