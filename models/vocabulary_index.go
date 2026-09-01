package models

import (
    "encoding/json"
    "errors"
)

// VocabularyIndex model.
type VocabularyIndex struct {
    // This app's name — the part before the dot in the qualified id
    // `customers.<name>`.
    App string `json:"app"`
    // Every vocabulary this app publishes, without their values.
    Vocabularies []interface{} `json:"vocabularies"`

    // Used by Decode() method
    data []byte
}

func (model VocabularyIndex) New(data []byte) *VocabularyIndex {
    model.data = data
    return &model
}

func (model *VocabularyIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}