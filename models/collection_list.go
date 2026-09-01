package models

import (
    "encoding/json"
    "errors"
)

// CollectionList model.
type CollectionList struct {
    // Public collection names the tenant owns. These are the values accepted for
    // the `collection` path parameter.
    Collections []string `json:"collections"`

    // Used by Decode() method
    data []byte
}

func (model CollectionList) New(data []byte) *CollectionList {
    model.data = data
    return &model
}

func (model *CollectionList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}