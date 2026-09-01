package models

import (
    "encoding/json"
    "errors"
)

// PriceVocabularyRef One vocabulary, named and titled — fetch its values
// with GET /prices/vocabularies/{name}.
type PriceVocabularyRef struct {
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

func (model PriceVocabularyRef) New(data []byte) *PriceVocabularyRef {
    model.data = data
    return &model
}

func (model *PriceVocabularyRef) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}