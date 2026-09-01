package models

import (
    "encoding/json"
    "errors"
)

// PageCommentUpdateRequest The new body. Nothing else about a comment is
// editable.
type PageCommentUpdateRequest struct {
    // The comment, as editor HTML. Replaces the old body completely.
    Body string `json:"body"`

    // Used by Decode() method
    data []byte
}

func (model PageCommentUpdateRequest) New(data []byte) *PageCommentUpdateRequest {
    model.data = data
    return &model
}

func (model *PageCommentUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}