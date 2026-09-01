package models

import (
    "encoding/json"
    "errors"
)

// CategoryRecomputeRequest Omit the body entirely to resume an unfinished
// pass, or start a fresh one when the last completed.
type CategoryRecomputeRequest struct {
    // The `cursor` a previous call returned, to continue that pass. Send `null`
    // explicitly to restart from the beginning; omit the field to let the app
    // decide (resume if a pass is in flight, otherwise start fresh). Anything
    // that is not a string or null is a 400.
    Cursor string `json:"cursor"`

    // Used by Decode() method
    data []byte
}

func (model CategoryRecomputeRequest) New(data []byte) *CategoryRecomputeRequest {
    model.data = data
    return &model
}

func (model *CategoryRecomputeRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}