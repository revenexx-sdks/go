package models

import (
    "encoding/json"
    "errors"
)

// ColumnURL Model
type ColumnUrl struct {
    // Column creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Column update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Is column an array?
    Array bool `json:"array"`
    // Error message. Displays error generated on failure of creating or deleting
    // an column.
    Error string `json:"error"`
    // String format.
    Format string `json:"format"`
    // Column Key.
    Key string `json:"key"`
    // Is column required?
    Required bool `json:"required"`
    // Column status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Column type.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model ColumnUrl) New(data []byte) *ColumnUrl {
    model.data = data
    return &model
}

func (model *ColumnUrl) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}