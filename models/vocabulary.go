package models

import (
    "encoding/json"
    "errors"
)

// Vocabulary model.
type Vocabulary struct {
    // This app's name — the part before the dot in the qualified id.
    App string `json:"app"`
    // True when the values are the complete permitted set. For a CHECK-backed
    // vocabulary the constraint guarantees it; for a table-backed one the app
    // refuses a value outside the rows, and for `locales` outside the configured
    // list — the same guarantee by three mechanisms.
    Closed bool `json:"closed"`
    // The tone an unlabelled value gets.
    DefaultTone string `json:"default_tone"`
    // A plain string, or a locale map keyed by language tag ({ "en": …, "de":
    // … }). Read the requested tag, fall back to `en`. A curated label is a
    // map; a value nobody labelled is humanized into a plain string.
    Description interface{} `json:"description"`
    // The vocabulary this is.
    Name string `json:"name"`
    // 'schema' — a CHECK constraint owns the set. 'table' — the tenant's own
    // rows do. 'defaults' — a table-backed set the tenant never wrote down,
    // answered from the built-ins. 'tenant' — the merchant configured the
    // values through a setting (locales).
    Source string `json:"source"`
    // A plain string, or a locale map keyed by language tag ({ "en": …, "de":
    // … }). Read the requested tag, fall back to `en`. A curated label is a
    // map; a value nobody labelled is humanized into a plain string.
    Title interface{} `json:"title"`
    // Every permitted value, in the order a select should offer them.
    Values []interface{} `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model Vocabulary) New(data []byte) *Vocabulary {
    model.data = data
    return &model
}

func (model *Vocabulary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}