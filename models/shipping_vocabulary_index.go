package models

import (
    "encoding/json"
    "errors"
)

// ShippingVocabularyIndex model.
type ShippingVocabularyIndex struct {
    // The app that owns these vocabularies — the part before the dot in a
    // qualified id.
    App string `json:"app"`
    // Every vocabulary this app publishes, without its values. Names only: fetch
    // one to get the set.
    Vocabularies []ShippingVocabularyIndexEntry `json:"vocabularies"`

    // Used by Decode() method
    data []byte
}

func (model ShippingVocabularyIndex) New(data []byte) *ShippingVocabularyIndex {
    model.data = data
    return &model
}

func (model *ShippingVocabularyIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}