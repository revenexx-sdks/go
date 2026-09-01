package models

import (
    "encoding/json"
    "errors"
)

// ShippingVocabularyValue model.
type ShippingVocabularyValue struct {
    // What the value means. Either one string or a locale map keyed by locale
    // (e.g. {en, de}) — curated copy carries the map, a value falling back to
    // its own key carries the string.
    Description string `json:"description"`
    // Table-backed only: localized descriptions, keyed by locale.
    Descriptions interface{} `json:"descriptions"`
    // weight-units only: kilograms per unit. A weight vocabulary without it is a
    // list of names you cannot convert with.
    Factor float64 `json:"factor"`
    // The value ends the lifecycle.
    Final bool `json:"final"`
    // weight-units only: the unit every other factor is expressed in.
    IsBase bool `json:"is_base"`
    // Table-backed only: the value a caller falls back to, so a client can mark
    // it without reading the settings as well.
    IsDefault bool `json:"is_default"`
    // Table-backed only: seeded on install. Still renameable and retirable.
    IsSystem bool `json:"is_system"`
    // The value as the database stores it — what a column carries and what a
    // filter matches. The only field a machine should compare on.
    Key string `json:"key"`
    // Table-backed only: localized titles, keyed by locale. Absent for a
    // vocabulary whose values come from a CHECK constraint — those carry their
    // copy in `title` instead.
    Labels interface{} `json:"labels"`
    // What a person reads. Falls back to a humanized key. Either one string or a
    // locale map keyed by locale (e.g. {en, de}) — curated copy carries the
    // map, a value falling back to its own key carries the string.
    Title string `json:"title"`
    // Semantic badge colour. The client owns what each tone looks like.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model ShippingVocabularyValue) New(data []byte) *ShippingVocabularyValue {
    model.data = data
    return &model
}

func (model *ShippingVocabularyValue) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}