package models

import (
    "encoding/json"
    "errors"
)

// ChannelVocabulary model.
type ChannelVocabulary struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // Always true: the set is exhaustive at this moment, so a value outside it is
    // stale data rather than a missing label. For a table-backed vocabulary that
    // is a statement about now, not forever — the tenant may add to it.
    Closed bool `json:"closed"`
    // The tone a value that carries none falls back to.
    DefaultTone string `json:"default_tone"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // Vocabulary name, unique within the app.
    Name string `json:"name"`
    // Who owns the value set. 'schema' = a CHECK constraint in this app's own
    // schema.json; 'table' = the tenant's own rows.
    Source string `json:"source"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`
    // Every permitted value, in author order — the order a select should offer,
    // not alphabetical. For a CHECK-backed vocabulary that is the constraint's
    // own order; for the table-backed `types` it is the tenant's `position`
    // order.
    Values []ChannelVocabularyValue `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model ChannelVocabulary) New(data []byte) *ChannelVocabulary {
    model.data = data
    return &model
}

func (model *ChannelVocabulary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}