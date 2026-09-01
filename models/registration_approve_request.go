package models

import (
    "encoding/json"
    "errors"
)

// RegistrationApproveRequest No required fields — send {}.
type RegistrationApproveRequest struct {
    // Who approved it — recorded on the contact and carried in the event. Free
    // text (operator id or email); this app does not resolve it.
    DecidedBy string `json:"decided_by"`

    // Used by Decode() method
    data []byte
}

func (model RegistrationApproveRequest) New(data []byte) *RegistrationApproveRequest {
    model.data = data
    return &model
}

func (model *RegistrationApproveRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}