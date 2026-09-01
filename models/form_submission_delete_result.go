package models

import (
    "encoding/json"
    "errors"
)

// FormSubmissionDeleteResult model.
type FormSubmissionDeleteResult struct {
    // Always true — the row is gone. A submission that was not there answers
    // 404 instead, so this is never false.
    Deleted bool `json:"deleted"`
    // The submission that was removed, echoed from the path.
    Id string `json:"id"`

    // Used by Decode() method
    data []byte
}

func (model FormSubmissionDeleteResult) New(data []byte) *FormSubmissionDeleteResult {
    model.data = data
    return &model
}

func (model *FormSubmissionDeleteResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}