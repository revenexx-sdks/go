package models

import (
    "encoding/json"
    "errors"
)

// Document Model
type Document struct {
    // Collection ID.
    CollectionId string `json:"$collectionId"`
    // Document creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Database ID.
    DatabaseId string `json:"$databaseId"`
    // Document ID.
    Id string `json:"$id"`
    // Document permissions. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    Permissions []string `json:"$permissions"`
    // Document automatically incrementing ID.
    Sequence int `json:"$sequence"`
    // Document update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`

    // Used by Decode() method
    data []byte
}

func (model Document) New(data []byte) *Document {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *Document) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}