package models

import (
    "encoding/json"
    "errors"
)

// OrderListVocabulary model.
type OrderListVocabulary struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // The set is exhaustive: a value outside it is stale data, not a missing
    // label.
    Closed bool `json:"closed"`
    // The badge colour a value carries when it names none of its own.
    DefaultTone string `json:"default_tone"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // Vocabulary name, unique within the app.
    Name string `json:"name"`
    // 'schema' — a CHECK constraint owns the set; 'table' — the tenant's own
    // rows do.
    Source string `json:"source"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`
    // Every permitted value, in the order a select should offer them.
    Values []OrderListVocabularyValue `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model OrderListVocabulary) New(data []byte) *OrderListVocabulary {
    model.data = data
    return &model
}

func (model *OrderListVocabulary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}