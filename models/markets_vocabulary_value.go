package models

import (
    "encoding/json"
    "errors"
)

// MarketsVocabularyValue One permitted value, with the copy and the badge
// tone a client renders it as.
type MarketsVocabularyValue struct {
    // Either one string, or a map of locale to string ({"en": …, "de": …}).
    Description string `json:"description"`
    // A terminal state nothing moves out of.
    Final bool `json:"final"`
    // The value as stored in the column.
    Key string `json:"key"`
    // Either one string, or a map of locale to string ({"en": …, "de": …}).
    Title string `json:"title"`
    // Semantic badge tone — the client decides what it looks like.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model MarketsVocabularyValue) New(data []byte) *MarketsVocabularyValue {
    model.data = data
    return &model
}

func (model *MarketsVocabularyValue) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}