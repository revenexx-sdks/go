package models

import (
    "encoding/json"
    "errors"
)

// PagesVocabularyIndex Which vocabularies this app publishes.
type PagesVocabularyIndex struct {
    // Always 'pages' — the first half of the qualified id a client holds.
    App string `json:"app"`
    // One entry per vocabulary, without its values.
    Vocabularies []PagesVocabularyRef `json:"vocabularies"`

    // Used by Decode() method
    data []byte
}

func (model PagesVocabularyIndex) New(data []byte) *PagesVocabularyIndex {
    model.data = data
    return &model
}

func (model *PagesVocabularyIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}