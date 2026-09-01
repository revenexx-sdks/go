package models

import (
    "encoding/json"
    "errors"
)

// RegistrationRejectRequest model.
type RegistrationRejectRequest struct {
    // Who rejected it — recorded on the contact and carried in the event.
    DecidedBy string `json:"decided_by"`
    // Why the application was declined. Always stored on the contact. It only
    // reaches the APPLICANT when the tenant's registration_reason_disclosed
    // setting is on — the event payload then carries it, and so does the 403
    // the login answers.
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model RegistrationRejectRequest) New(data []byte) *RegistrationRejectRequest {
    model.data = data
    return &model
}

func (model *RegistrationRejectRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}