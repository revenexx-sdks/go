package models

import (
    "encoding/json"
    "errors"
)

// OrderVocabularyIndex model.
type OrderVocabularyIndex struct {
    // This app's name — the part before the dot in the qualified id.
    App string `json:"app"`
    // Every vocabulary this app publishes, without its values — fetch one with
    // GET /orders/vocabularies/{name}.
    Vocabularies []OrderVocabularySummary `json:"vocabularies"`

    // Used by Decode() method
    data []byte
}

func (model OrderVocabularyIndex) New(data []byte) *OrderVocabularyIndex {
    model.data = data
    return &model
}

func (model *OrderVocabularyIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}