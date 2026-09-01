package models

import (
    "encoding/json"
    "errors"
)

// CartVocabularyRef model.
type CartVocabularyRef struct {
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // Vocabulary name, unique within the app.
    Name string `json:"name"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`

    // Used by Decode() method
    data []byte
}

func (model CartVocabularyRef) New(data []byte) *CartVocabularyRef {
    model.data = data
    return &model
}

func (model *CartVocabularyRef) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}