package models

import (
    "encoding/json"
    "errors"
)

// PageBlockTree The block and everything under it, serialized. This is the
// payload: every page that references the item renders THIS tree, so editing
// it here changes every placement at once.
type PageBlockTree struct {
    // The block type — `hero`, `text`, `teaser`, whatever the active theme
    // defines. It decides which component renders it and which props it carries.
    Bundle string `json:"bundle"`
    // Nested blocks, keyed by the field they sit in — `{ "content": [...],
    // "buttons": [...] }`. Absent on a leaf block.
    Children interface{} `json:"children"`
    // The theme fragment this block renders instead of a props-driven component,
    // or `null` for an ordinary block. Theme-defined, like a bundle.
    FragmentName string `json:"fragment_name"`
    // blökkli display options for this block, as a flat `option key → value`
    // map (variant, spacing, background). Theme-defined, set by the
    // `update_options` mutation.
    Options interface{} `json:"options"`
    // The block's field values in the page's SOURCE language, as a flat `field
    // name → value` map. The field names are the theme's; this app stores and
    // replays them without reading one.
    Props interface{} `json:"props"`
    // Per-language overrides of `props`, keyed by langcode: `{ "en": { "title":
    // "About us" } }`. A field missing for a language falls back to `props`,
    // which is why a half-translated page still renders.
    PropsI18n interface{} `json:"props_i18n"`

    // Used by Decode() method
    data []byte
}

func (model PageBlockTree) New(data []byte) *PageBlockTree {
    model.data = data
    return &model
}

func (model *PageBlockTree) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}