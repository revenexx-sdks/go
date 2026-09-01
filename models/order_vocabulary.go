package models

import (
    "encoding/json"
    "errors"
)

// OrderVocabulary model.
type OrderVocabulary struct {
    // This app's name — the part before the dot in the qualified id.
    App string `json:"app"`
    // True when the values are the complete permitted set — always, since the
    // routes enforce the ones the schema does not.
    Closed bool `json:"closed"`
    // The tone an unlabelled value gets.
    DefaultTone string `json:"default_tone"`
    // Either one string, or a map of locale to string ({"en": …, "de": …}).
    Description string `json:"description"`
    // Which vocabulary this is — echoed from the path, and the part after the
    // dot in the qualified id.
    Name string `json:"name"`
    // Who enforces the set: 'schema' = a CHECK constraint, 'app' = the routes.
    Source string `json:"source"`
    // Either one string, or a map of locale to string ({"en": …, "de": …}).
    Title string `json:"title"`
    // Every permitted value, in CONSTRAINT order — which for a status is
    // lifecycle order, so a client can render them as a sequence without knowing
    // one.
    Values []OrderVocabularyValue `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model OrderVocabulary) New(data []byte) *OrderVocabulary {
    model.data = data
    return &model
}

func (model *OrderVocabulary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}