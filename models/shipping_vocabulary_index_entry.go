package models

import (
    "encoding/json"
    "errors"
)

// ShippingVocabularyIndexEntry One vocabulary, named and titled.
type ShippingVocabularyIndexEntry struct {
    // What the vocabulary is for. Either one string or a locale map keyed by
    // locale (e.g. {en, de}) — curated copy carries the map, a value falling
    // back to its own key carries the string.
    Description string `json:"description"`
    // The part after the dot in the qualified id — what GET
    // /shipping/vocabularies/{name} takes.
    Name string `json:"name"`
    // What the vocabulary is called. Either one string or a locale map keyed by
    // locale (e.g. {en, de}) — curated copy carries the map, a value falling
    // back to its own key carries the string.
    Title string `json:"title"`

    // Used by Decode() method
    data []byte
}

func (model ShippingVocabularyIndexEntry) New(data []byte) *ShippingVocabularyIndexEntry {
    model.data = data
    return &model
}

func (model *ShippingVocabularyIndexEntry) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}