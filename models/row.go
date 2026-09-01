package models

import (
    "encoding/json"
    "errors"
)

// Row Row
type Row struct {
    // Row creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Database ID.
    DatabaseId string `json:"$databaseId"`
    // Row ID.
    Id string `json:"$id"`
    // Row permissions. Each entry is a permission string: an action wrapping a
    // role, e.g. `read("any")`, `update("user:abc")`, `delete("team:abc/owner")`.
    // Actions are `read`, `create`, `update`, `delete` and the aggregate `write`
    // (= create + update + delete); the role inside the quotes takes the form
    // described under “Role strings” in this document's introduction.
    Permissions []string `json:"$permissions"`
    // Row automatically incrementing ID.
    Sequence int `json:"$sequence"`
    // Table ID.
    TableId string `json:"$tableId"`
    // Row update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`

    // Used by Decode() method
    data []byte
}

func (model Row) New(data []byte) *Row {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *Row) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}