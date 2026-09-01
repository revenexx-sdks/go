package models

import (
    "encoding/json"
    "errors"
)

// FormDefaultsResult model.
type FormDefaultsResult struct {
    // Slugs this call created. On a tenant that has had the app installed for
    // more than a moment this is empty — the sample form is seeded on
    // `app.installed`.
    Created []string `json:"created"`
    // Slugs that were already there and were left alone. Nothing about them was
    // overwritten — a form the merchant has edited stays edited.
    Existing []string `json:"existing"`

    // Used by Decode() method
    data []byte
}

func (model FormDefaultsResult) New(data []byte) *FormDefaultsResult {
    model.data = data
    return &model
}

func (model *FormDefaultsResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}