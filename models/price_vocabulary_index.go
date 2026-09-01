package models

import (
    "encoding/json"
    "errors"
)

// PriceVocabularyIndex What this app publishes, without the values — one
// fetch a UI can cache and then pull only the vocabularies it renders.
type PriceVocabularyIndex struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // Every vocabulary this app owns, sorted by name.
    Vocabularies []PriceVocabularyRef `json:"vocabularies"`

    // Used by Decode() method
    data []byte
}

func (model PriceVocabularyIndex) New(data []byte) *PriceVocabularyIndex {
    model.data = data
    return &model
}

func (model *PriceVocabularyIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}