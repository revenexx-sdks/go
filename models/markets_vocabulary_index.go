package models

import (
    "encoding/json"
    "errors"
)

// MarketsVocabularyIndex Every closed value set this app owns, by name —
// enough to build a menu of them without fetching each one.
type MarketsVocabularyIndex struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // Every vocabulary this app publishes, named and titled but without its
    // values — fetch one by name for those.
    Vocabularies []MarketsVocabularySummary `json:"vocabularies"`

    // Used by Decode() method
    data []byte
}

func (model MarketsVocabularyIndex) New(data []byte) *MarketsVocabularyIndex {
    model.data = data
    return &model
}

func (model *MarketsVocabularyIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}