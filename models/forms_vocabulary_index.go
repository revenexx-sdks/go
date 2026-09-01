package models

import (
    "encoding/json"
    "errors"
)

// FormsVocabularyIndex model.
type FormsVocabularyIndex struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // Every vocabulary this app publishes, without its values — enough to build
    // a menu, not enough to fill a select. Fetch one by name for that.
    Vocabularies []FormsVocabularySummary `json:"vocabularies"`

    // Used by Decode() method
    data []byte
}

func (model FormsVocabularyIndex) New(data []byte) *FormsVocabularyIndex {
    model.data = data
    return &model
}

func (model *FormsVocabularyIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}