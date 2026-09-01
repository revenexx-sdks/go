package models

import (
    "encoding/json"
    "errors"
)

// MarketsVocabulary One closed value set this app owns, parsed out of the
// CHECK constraint in schema.json — the served set IS the enforced set.
// `closed: true` means a client may treat anything outside `values` as stale
// data.
type MarketsVocabulary struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // Always true here: the values come from a CHECK constraint, so the list is
    // exhaustive.
    Closed bool `json:"closed"`
    // The tone a value that carries none falls back to.
    DefaultTone string `json:"default_tone"`
    // Either one string, or a map of locale to string ({"en": …, "de": …}).
    Description string `json:"description"`
    // Vocabulary name, unique within the app.
    Name string `json:"name"`
    // Where the values came from. 'schema' = a CHECK constraint in this app's own
    // schema.json.
    Source string `json:"source"`
    // Either one string, or a map of locale to string ({"en": …, "de": …}).
    Title string `json:"title"`
    // Every value the column may hold, in the order the CHECK constraint lists
    // them — which is the order a select box should offer them in. Exhaustive,
    // because `closed` is true.
    Values []MarketsVocabularyValue `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model MarketsVocabulary) New(data []byte) *MarketsVocabulary {
    model.data = data
    return &model
}

func (model *MarketsVocabulary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}