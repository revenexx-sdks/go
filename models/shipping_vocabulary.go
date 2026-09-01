package models

import (
    "encoding/json"
    "errors"
)

// ShippingVocabulary model.
type ShippingVocabulary struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // The set is exhaustive, so a value outside it is stale data rather than a
    // missing label. True either way — what differs is who may extend it.
    Closed bool `json:"closed"`
    // The badge colour a value that names none falls back to.
    DefaultTone string `json:"default_tone"`
    // What the vocabulary is for. Either one string or a locale map keyed by
    // locale (e.g. {en, de}) — curated copy carries the map, a value falling
    // back to its own key carries the string.
    Description string `json:"description"`
    // The vocabulary name — the part after the dot in the qualified id.
    Name string `json:"name"`
    // 'schema' — the values are a CHECK constraint's, so the served set IS the
    // enforced set. 'table' — the values are the tenant's own rows, read per
    // request.
    Source string `json:"source"`
    // What the vocabulary is called. Either one string or a locale map keyed by
    // locale (e.g. {en, de}) — curated copy carries the map, a value falling
    // back to its own key carries the string.
    Title string `json:"title"`
    // Every permitted value, in the order a select should offer them —
    // constraint order for a schema vocabulary, `position` for a table one.
    Values []ShippingVocabularyValue `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model ShippingVocabulary) New(data []byte) *ShippingVocabulary {
    model.data = data
    return &model
}

func (model *ShippingVocabulary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}