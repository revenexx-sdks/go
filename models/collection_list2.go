package models

import (
    "encoding/json"
    "errors"
)

// CollectionList2 Collections List
type CollectionList2 struct {
    // List of collections.
    Collections []Collection2 `json:"collections"`
    // Total number of collections that matched your query.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model CollectionList2) New(data []byte) *CollectionList2 {
    model.data = data
    return &model
}

func (model *CollectionList2) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}