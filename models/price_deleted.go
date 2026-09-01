package models

import (
    "encoding/json"
    "errors"
)

// PriceDeleted The row is gone. Deleting a price list cascades to its
// entries.
type PriceDeleted struct {
    // Always true — a row that was not there answers 404 instead.
    Deleted bool `json:"deleted"`
    // The row that was removed.
    Id string `json:"id"`

    // Used by Decode() method
    data []byte
}

func (model PriceDeleted) New(data []byte) *PriceDeleted {
    model.data = data
    return &model
}

func (model *PriceDeleted) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}