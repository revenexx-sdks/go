package models

import (
    "encoding/json"
    "errors"
)

// FormsVocabularyValue model.
type FormsVocabularyValue struct {
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // The value ends the lifecycle.
    Final bool `json:"final"`
    // The value as the database stores and enforces it.
    Key string `json:"key"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`
    // Semantic badge colour. The client owns what each tone looks like.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model FormsVocabularyValue) New(data []byte) *FormsVocabularyValue {
    model.data = data
    return &model
}

func (model *FormsVocabularyValue) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}