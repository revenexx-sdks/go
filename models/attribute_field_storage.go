package models

import (
    "encoding/json"
    "errors"
)

// AttributeFieldStorage Where the value lives. Absent on an app whose custom
// fields are plain columns — then the field name IS the column.
type AttributeFieldStorage struct {
    // Which scope bucket this attribute writes to, implied by
    // localizable/scopable.
    Bucket string `json:"bucket"`
    // The jsonb column holding the values (`attribute_values`).
    Column string `json:"column"`
    // The exact key path for the requested context, or null when the request
    // named no locale/channel and the bucket needs one. Null means: read-only
    // until a context is chosen.
    Path []string `json:"path"`

    // Used by Decode() method
    data []byte
}

func (model AttributeFieldStorage) New(data []byte) *AttributeFieldStorage {
    model.data = data
    return &model
}

func (model *AttributeFieldStorage) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}