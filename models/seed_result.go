package models

import (
    "encoding/json"
    "errors"
)

// SeedResult What was created and what was already there. Nothing is ever
// overwritten, so a non-empty `skipped` is the normal answer to a second run.
type SeedResult struct {
    // The menu half of the run.
    Menus interface{} `json:"menus"`
    // The page half of the run.
    Pages interface{} `json:"pages"`

    // Used by Decode() method
    data []byte
}

func (model SeedResult) New(data []byte) *SeedResult {
    model.data = data
    return &model
}

func (model *SeedResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}