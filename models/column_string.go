package models

import (
    "encoding/json"
    "errors"
)

// ColumnString Model
type ColumnString struct {
    // Column creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Column update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Is column an array?
    Array bool `json:"array"`
    // Defines whether this column is encrypted or not.
    Encrypt bool `json:"encrypt"`
    // Error message. Displays error generated on failure of creating or deleting
    // an column.
    Error string `json:"error"`
    // Column Key.
    Key string `json:"key"`
    // Is column required?
    Required bool `json:"required"`
    // Column size.
    Size int `json:"size"`
    // Column status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Column type.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model ColumnString) New(data []byte) *ColumnString {
    model.data = data
    return &model
}

func (model *ColumnString) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}