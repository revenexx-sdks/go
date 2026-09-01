package models

import (
    "encoding/json"
    "errors"
)

// PriceVocabulary One closed value set with the words a human reads for it
// — so a UI never keeps its own copy of an enum this app enforces.
type PriceVocabulary struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // Always true here: the values come from a CHECK constraint, so the list is
    // exhaustive and a value outside it is stale data rather than a missing
    // label.
    Closed bool `json:"closed"`
    // The tone a value that carries none falls back to.
    DefaultTone string `json:"default_tone"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // Vocabulary name, unique within the app.
    Name string `json:"name"`
    // Where the values came from. 'schema' = a CHECK constraint in this app's own
    // schema.json.
    Source string `json:"source"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`
    // Every permitted value, in CHECK-constraint order — which is the order an
    // author wrote and the order a select should offer.
    Values []PriceVocabularyValue `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model PriceVocabulary) New(data []byte) *PriceVocabulary {
    model.data = data
    return &model
}

func (model *PriceVocabulary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}