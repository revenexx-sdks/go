package models

import (
    "encoding/json"
    "errors"
)

// DeliveryPageRef Just enough of a published page to link to it. The block
// tree is not here — fetch it with `GET /pages/delivery/page`.
type DeliveryPageRef struct {
    // The page type, so a sitemap can group or a picker can filter.
    Bundle string `json:"bundle"`
    // The page id, usable as `?id=` on the delivery route.
    Id string `json:"id"`
    // The path segment to build the URL from. `null` for a page reachable only by
    // id, which a sitemap should skip.
    Slug string `json:"slug"`
    // The page title in its source language — this projection is not
    // language-resolved.
    Title string `json:"title"`

    // Used by Decode() method
    data []byte
}

func (model DeliveryPageRef) New(data []byte) *DeliveryPageRef {
    model.data = data
    return &model
}

func (model *DeliveryPageRef) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}