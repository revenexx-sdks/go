package models

import (
    "encoding/json"
    "errors"
)

// FormSubmissionPruneSample One row the sweep would delete, shown so a
// merchant can recognise what is at stake before turning the preview off.
// Three columns only — never the submitted data.
type FormSubmissionPruneSample struct {
    // When it arrived — the age this sweep is judging it on.
    CreatedAt string `json:"created_at"`
    // The form's slug as it stood when this submission arrived, copied onto the
    // row: the inbox filters by form without a join, and a submission still says
    // which form collected it after that form has been renamed. It does not
    // outlive a DELETED form — the foreign key cascades and takes the
    // submission with it. On a write the body's value WINS; omit it and the
    // form's own slug is copied in.
    FormSlug string `json:"form_slug"`
    // The submission that would be deleted. Fetch it with GET
    // /v1/forms/submissions/{id} to see what it holds.
    Id string `json:"id"`

    // Used by Decode() method
    data []byte
}

func (model FormSubmissionPruneSample) New(data []byte) *FormSubmissionPruneSample {
    model.data = data
    return &model
}

func (model *FormSubmissionPruneSample) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}