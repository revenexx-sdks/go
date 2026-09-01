package models

import (
    "encoding/json"
    "errors"
)

// DeliveryBlock One block, ready to render: props resolved for the requested
// language, library references already expanded, scheduled blocks already
// filtered out.
type DeliveryBlock struct {
    // The block type. This is what a theme switches its component on.
    Bundle string `json:"bundle"`
    // Nested blocks keyed by the field they sit in — `{ "columns": [...] }`.
    // Empty object on a leaf block.
    Children interface{} `json:"children"`
    // The theme fragment to render instead of a props-driven component.
    // Theme-defined, like a bundle.
    FragmentName string `json:"fragmentName"`
    // The library item this block came from, or `null`. Its content is already
    // inlined above — this is for cache invalidation and editor links, not for
    // a second fetch.
    LibraryItemId string `json:"libraryItemId"`
    // Display options for this block, as a flat `option key → value` map.
    Options interface{} `json:"options"`
    // The block's field values for the requested language, source values already
    // overlaid with that language's overrides. Theme-defined keys.
    Props interface{} `json:"props"`
    // The block uuid — stable across publishes, so it is safe to use as a
    // render key or an anchor.
    Uuid string `json:"uuid"`

    // Used by Decode() method
    data []byte
}

func (model DeliveryBlock) New(data []byte) *DeliveryBlock {
    model.data = data
    return &model
}

func (model *DeliveryBlock) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}