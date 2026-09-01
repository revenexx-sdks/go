package models

import (
    "encoding/json"
    "errors"
)

// CollectionField One field in a collection schema.
type CollectionField struct {
    // Whether the field can be faceted on.
    Facet bool `json:"facet"`
    // 
    Index bool `json:"index"`
    // 
    Name string `json:"name"`
    // 
    Optional bool `json:"optional"`
    // 
    Sort bool `json:"sort"`
    // Typesense field type, e.g. `string`, `int64`, `string[]`, `object`.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model CollectionField) New(data []byte) *CollectionField {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *CollectionField) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}