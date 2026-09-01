package models

import (
    "encoding/json"
    "errors"
)

// PaymentVocabulary One enum this app owns, with every permitted value.
type PaymentVocabulary struct {
    // The app that owns this vocabulary — always `payments` here. Together with
    // `name` it forms the platform-wide key `payments.statuses`.
    App string `json:"app"`
    // True when the set comes from a CHECK constraint and is therefore exhaustive
    // — a client may treat anything outside it as stale data rather than a
    // missing label.
    Closed bool `json:"closed"`
    // The tone a permitted value nobody labelled falls back to, so every value is
    // renderable.
    DefaultTone string `json:"default_tone"`
    // What this set of values is about. A plain string, or a locale map keyed by
    // language tag ({ "en": …, "de": … }). Read the requested tag, fall back
    // to `en`.
    Description interface{} `json:"description"`
    // The vocabulary name, as it appears in the URL.
    Name string `json:"name"`
    // Where the values come from. `schema` means they were parsed out of the
    // CHECK constraint, so what is served is what the database enforces.
    Source string `json:"source"`
    // The vocabulary's own label, for a filter heading or a column title. A plain
    // string, or a locale map keyed by language tag ({ "en": …, "de": … }).
    // Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`
    // Every permitted value, in constraint order — which is the lifecycle order
    // an author wrote, and the order a select should offer.
    Values []PaymentVocabularyValue `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model PaymentVocabulary) New(data []byte) *PaymentVocabulary {
    model.data = data
    return &model
}

func (model *PaymentVocabulary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}