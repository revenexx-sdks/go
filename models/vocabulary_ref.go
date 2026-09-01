package models

import (
    "encoding/json"
    "errors"
)

// VocabularyRef model.
type VocabularyRef struct {
    // A plain string, or a locale map keyed by language tag ({ "en": …, "de":
    // … }). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // The name to pass to `GET /products/vocabularies/{name}`.
    Name string `json:"name"`
    // A plain string, or a locale map keyed by language tag ({ "en": …, "de":
    // … }). Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`

    // Used by Decode() method
    data []byte
}

func (model VocabularyRef) New(data []byte) *VocabularyRef {
    model.data = data
    return &model
}

func (model *VocabularyRef) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}