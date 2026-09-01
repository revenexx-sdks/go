package models

import (
    "encoding/json"
    "errors"
)

// InventoryVocabularyIndex model.
type InventoryVocabularyIndex struct {
    // This app's name — the part before the dot in a qualified vocabulary id
    // such as `inventories.movement-types`.
    App string `json:"app"`
    // Every vocabulary this app publishes, WITHOUT its values — the index a
    // client reads to discover them. Fetch the values with GET
    // /inventories/vocabularies/{name}.
    Vocabularies []interface{} `json:"vocabularies"`

    // Used by Decode() method
    data []byte
}

func (model InventoryVocabularyIndex) New(data []byte) *InventoryVocabularyIndex {
    model.data = data
    return &model
}

func (model *InventoryVocabularyIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}