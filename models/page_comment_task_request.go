package models

import (
    "encoding/json"
    "errors"
)

// PageCommentTaskRequest Which checkbox to flip.
type PageCommentTaskRequest struct {
    // The task item to toggle, counted in document order from 0. A comment with
    // fewer tasks than that answers 400, and so does anything that is not a whole
    // number at or above 0.
    TaskIndex int `json:"taskIndex"`

    // Used by Decode() method
    data []byte
}

func (model PageCommentTaskRequest) New(data []byte) *PageCommentTaskRequest {
    model.data = data
    return &model
}

func (model *PageCommentTaskRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}