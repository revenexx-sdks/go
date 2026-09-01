package models

import (
    "encoding/json"
    "errors"
)

// CartVocabularyIndex model.
type CartVocabularyIndex struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // Every vocabulary this app publishes, without its values — enough to build
    // a menu, and one call per vocabulary to fill it.
    Vocabularies []CartVocabularyRef `json:"vocabularies"`

    // Used by Decode() method
    data []byte
}

func (model CartVocabularyIndex) New(data []byte) *CartVocabularyIndex {
    model.data = data
    return &model
}

func (model *CartVocabularyIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}