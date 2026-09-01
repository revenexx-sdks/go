package models

import (
    "encoding/json"
    "errors"
)

// OrderVocabularySummary One vocabulary, named and titled but without its
// values.
type OrderVocabularySummary struct {
    // Either one string, or a map of locale to string ({"en": …, "de": …}).
    Description string `json:"description"`
    // Vocabulary name, unique within the app.
    Name string `json:"name"`
    // Either one string, or a map of locale to string ({"en": …, "de": …}).
    Title string `json:"title"`

    // Used by Decode() method
    data []byte
}

func (model OrderVocabularySummary) New(data []byte) *OrderVocabularySummary {
    model.data = data
    return &model
}

func (model *OrderVocabularySummary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}