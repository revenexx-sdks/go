package models

import (
    "encoding/json"
    "errors"
)

// FormsVocabulary model.
type FormsVocabulary struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // The set is exhaustive.
    Closed bool `json:"closed"`
    // The tone a value nobody gave one falls back to — what a badge looks like
    // for a status that was added to the CHECK constraint before anyone styled
    // it.
    DefaultTone string `json:"default_tone"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // Vocabulary name, unique within the app.
    Name string `json:"name"`
    // Parsed from the CHECK constraint.
    Source string `json:"source"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`
    // Every permitted value, in constraint order — which is the order a select
    // should offer them in, because it is the lifecycle order.
    Values []FormsVocabularyValue `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model FormsVocabulary) New(data []byte) *FormsVocabulary {
    model.data = data
    return &model
}

func (model *FormsVocabulary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}