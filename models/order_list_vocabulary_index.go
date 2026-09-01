package models

import (
    "encoding/json"
    "errors"
)

// OrderListVocabularyIndex model.
type OrderListVocabularyIndex struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // Every vocabulary this app publishes, without its values — the values are
    // one call further down, at GET /orderlists/vocabularies/{name}.
    Vocabularies []interface{} `json:"vocabularies"`

    // Used by Decode() method
    data []byte
}

func (model OrderListVocabularyIndex) New(data []byte) *OrderListVocabularyIndex {
    model.data = data
    return &model
}

func (model *OrderListVocabularyIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}