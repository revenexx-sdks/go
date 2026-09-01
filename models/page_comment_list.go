package models

import (
    "encoding/json"
    "errors"
)

// PageCommentList Every comment of the page, roots and replies flat in one
// list, oldest first — the editor builds the threads from `parentUuid`.
// Every write route answers this same full list rather than the row it
// changed.
type PageCommentList struct {
    // The page's comments, oldest first.
    Items []PageCommentItem `json:"items"`

    // Used by Decode() method
    data []byte
}

func (model PageCommentList) New(data []byte) *PageCommentList {
    model.data = data
    return &model
}

func (model *PageCommentList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}