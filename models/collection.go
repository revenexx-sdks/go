package models

import (
    "encoding/json"
    "errors"
)

// Collection Model
type Collection struct {
    // Collection creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Collection ID.
    Id string `json:"$id"`
    // Collection permissions. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    Permissions []string `json:"$permissions"`
    // Collection update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Collection attributes.
    Attributes []map[string]any `json:"attributes"`
    // Maximum document size in bytes. Returns 0 when no limit applies.
    BytesMax int `json:"bytesMax"`
    // Currently used document size in bytes based on defined attributes.
    BytesUsed int `json:"bytesUsed"`
    // Database ID.
    DatabaseId string `json:"databaseId"`
    // Whether document-level permissions are enabled. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    DocumentSecurity bool `json:"documentSecurity"`
    // Collection enabled. Can be 'enabled' or 'disabled'. When disabled, the
    // collection is inaccessible to users, but remains accessible to Server SDKs
    // using API keys.
    Enabled bool `json:"enabled"`
    // Collection indexes.
    Indexes []Index `json:"indexes"`
    // Collection name.
    Name string `json:"name"`

    // Used by Decode() method
    data []byte
}

func (model Collection) New(data []byte) *Collection {
    model.data = data
    return &model
}

func (model *Collection) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}