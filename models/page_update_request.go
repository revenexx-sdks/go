package models

import (
    "encoding/json"
    "errors"
)

// PageUpdateRequest Partial update — only title, slug, status, meta and
// bundle are applied; other keys are ignored. The page's CONTENT is never
// edited here: blocks change through the editor's mutation log.
type PageUpdateRequest struct {
    // The page type. Changing it changes which template the theme renders.
    Bundle string `json:"bundle"`
    // The page's metadata bag. Replaced wholesale, not merged.
    Meta interface{} `json:"meta"`
    // The path segment the storefront routes it under. Sending a slug another
    // live page holds answers 409; sending null makes the page unreachable by
    // path.
    Slug string `json:"slug"`
    // The lifecycle status. Setting `published` here does NOT publish content —
    // delivery still needs a revision, which only `POST
    // /pages/editor/{page_id}/publish` writes.
    Status string `json:"status"`
    // The page title in its source language.
    Title string `json:"title"`

    // Used by Decode() method
    data []byte
}

func (model PageUpdateRequest) New(data []byte) *PageUpdateRequest {
    model.data = data
    return &model
}

func (model *PageUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}