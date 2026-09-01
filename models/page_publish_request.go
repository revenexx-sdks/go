package models

import (
    "encoding/json"
    "errors"
)

// PagePublishRequest What to record about this publication.
type PagePublishRequest struct {
    // Publish despite violations. Without it a page with unresolved violations
    // answers 422 and nothing is written.
    Force bool `json:"force"`
    // What to call this publication in the page's history — "Autumn campaign"
    // rather than a timestamp.
    Label string `json:"label"`

    // Used by Decode() method
    data []byte
}

func (model PagePublishRequest) New(data []byte) *PagePublishRequest {
    model.data = data
    return &model
}

func (model *PagePublishRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}