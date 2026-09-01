package models

import (
    "encoding/json"
    "errors"
)

// PagesVocabulary One vocabulary and every value it permits.
type PagesVocabulary struct {
    // Always 'pages'.
    App string `json:"app"`
    // The set is exhaustive, so a value outside it is stale data rather than a
    // missing label.
    Closed bool `json:"closed"`
    // The badge colour a value nobody toned falls back to.
    DefaultTone string `json:"default_tone"`
    // What the set is for, or null. A plain string, or a locale map keyed by
    // language tag ({ "en": …, "de": … }). Read the requested tag, fall back
    // to `en`.
    Description interface{} `json:"description"`
    // The vocabulary name, echoed.
    Name string `json:"name"`
    // Always 'schema' — the values are parsed from the column's CHECK
    // constraint, which is why the served set cannot drift from the enforced one.
    Source string `json:"source"`
    // What this set of values is called. A plain string, or a locale map keyed by
    // language tag ({ "en": …, "de": … }). Read the requested tag, fall back
    // to `en`.
    Title interface{} `json:"title"`
    // Every permitted value, in the order the constraint lists them — which is
    // the order a select should offer.
    Values []PagesVocabularyValue `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model PagesVocabulary) New(data []byte) *PagesVocabulary {
    model.data = data
    return &model
}

func (model *PagesVocabulary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}