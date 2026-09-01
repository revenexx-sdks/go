package models

import (
    "encoding/json"
    "errors"
)

// SeedRequest A theme's starting content. Both lists are optional; sending
// neither is a no-op.
type SeedRequest struct {
    // The menus to create. One with no key or no label is reported under
    // `skipped`.
    Menus []interface{} `json:"menus"`
    // The pages to create. One that has no `slug` or no `title` is reported under
    // `skipped` rather than refused, so one bad entry never loses the rest.
    Pages []interface{} `json:"pages"`

    // Used by Decode() method
    data []byte
}

func (model SeedRequest) New(data []byte) *SeedRequest {
    model.data = data
    return &model
}

func (model *SeedRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}