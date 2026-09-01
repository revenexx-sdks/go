package models

import (
    "encoding/json"
    "errors"
)

// PageRevisionRef One publication of this page, without the snapshot — who
// published, when, and under what name.
type PageRevisionRef struct {
    // When this revision was published.
    CreatedAt string `json:"created_at"`
    // The user id that published.
    CreatedBy string `json:"created_by"`
    // That user's display name, copied in at publish time so the history stays
    // readable after the user is gone.
    CreatedByName string `json:"created_by_name"`
    // The revision id. A page's `published_revision_id` points at one of these,
    // and it is the only thing delivery reads.
    Id string `json:"id"`
    // What this publication was called, e.g. "Autumn campaign". It is what turns
    // the history into a list of changes rather than a list of timestamps.
    Label string `json:"label"`
    // The page this revision belongs to.
    PageId string `json:"page_id"`

    // Used by Decode() method
    data []byte
}

func (model PageRevisionRef) New(data []byte) *PageRevisionRef {
    model.data = data
    return &model
}

func (model *PageRevisionRef) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}