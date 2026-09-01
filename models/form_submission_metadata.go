package models

import (
    "encoding/json"
    "errors"
)

// FormSubmissionMetadata Free-form metadata, plus what this app stamped on at
// insert. The recipient is resolved ONCE, here, because this row is the
// payload of `form.submitted` — a workflow reads the address off the event
// instead of re-resolving a form's settings that may since have changed.
type FormSubmissionMetadata struct {
    // The resolved notification recipient, or null when neither the form nor the
    // tenant names one.
    NotifyEmail string `json:"notify_email"`
    // Which of the two configured recipients won: the form's own, or the tenant
    // setting.
    NotifySource string `json:"notify_source"`
    // Present only on a submission the honeypot caught: 'honeypot'.
    SpamReason string `json:"spam_reason"`

    // Used by Decode() method
    data []byte
}

func (model FormSubmissionMetadata) New(data []byte) *FormSubmissionMetadata {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *FormSubmissionMetadata) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}