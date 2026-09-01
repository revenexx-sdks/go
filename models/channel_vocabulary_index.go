package models

import (
    "encoding/json"
    "errors"
)

// ChannelVocabularyIndex model.
type ChannelVocabularyIndex struct {
    // The app that owns this vocabulary.
    App string `json:"app"`
    // Every vocabulary this app owns, alphabetically: statuses, types,
    // unassigned-visibility. Names only — fetch the values with GET
    // /channels/vocabularies/{name}.
    Vocabularies []ChannelVocabularyRef `json:"vocabularies"`

    // Used by Decode() method
    data []byte
}

func (model ChannelVocabularyIndex) New(data []byte) *ChannelVocabularyIndex {
    model.data = data
    return &model
}

func (model *ChannelVocabularyIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}