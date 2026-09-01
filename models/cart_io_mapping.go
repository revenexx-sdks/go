package models

import (
    "encoding/json"
    "errors"
)

// CartIoMapping Baseline-IO-compatible column mapping. An empty object (or
// null) is identity: the full canonical shape, every field under its own
// name.
type CartIoMapping struct {
    // Renames, in order. On export the row is narrowed to these columns; on
    // import a column that is not listed is ignored. Omit or leave empty for
    // identity.
    Columns []CartIoMappingColumn `json:"columns"`
    // Fields that identify a line in the payload — what the bundled quick-order
    // template sets to ['sku'].
    Keys []string `json:"keys"`

    // Used by Decode() method
    data []byte
}

func (model CartIoMapping) New(data []byte) *CartIoMapping {
    model.data = data
    return &model
}

func (model *CartIoMapping) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}