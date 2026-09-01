package models

import (
    "encoding/json"
    "errors"
)

// Table Table
type Table struct {
    // Table creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Table ID.
    Id string `json:"$id"`
    // Table permissions. Each entry is a permission string: an action wrapping a
    // role, e.g. `read("any")`, `update("user:abc")`, `delete("team:abc/owner")`.
    // Actions are `read`, `create`, `update`, `delete` and the aggregate `write`
    // (= create + update + delete); the role inside the quotes takes the form
    // described under “Role strings” in this document's introduction.
    Permissions []string `json:"$permissions"`
    // Table update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Maximum row size in bytes. Returns 0 when no limit applies.
    BytesMax int `json:"bytesMax"`
    // Currently used row size in bytes based on defined columns.
    BytesUsed int `json:"bytesUsed"`
    // Table columns.
    Columns []interface{} `json:"columns"`
    // Database ID.
    DatabaseId string `json:"databaseId"`
    // Table enabled. Can be 'enabled' or 'disabled'. When disabled, the table is
    // inaccessible to users, but remains accessible to Server SDKs using API
    // keys.
    Enabled bool `json:"enabled"`
    // Table indexes.
    Indexes []ColumnIndex `json:"indexes"`
    // Table name.
    Name string `json:"name"`
    // Whether row-level permissions are enabled. When it is, each record's own
    // `$permissions` are enforced on top of the container's.
    RowSecurity bool `json:"rowSecurity"`

    // Used by Decode() method
    data []byte
}

func (model Table) New(data []byte) *Table {
    model.data = data
    return &model
}

func (model *Table) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}