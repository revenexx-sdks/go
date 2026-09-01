package models

import (
    "encoding/json"
    "errors"
)

// OrderDeleted The row is gone. Deleting is not idempotent here: a second
// call answers 404, because the row no longer resolves.
type OrderDeleted struct {
    // Always true — a failed delete is a status code, not a false here.
    Deleted bool `json:"deleted"`
    // The id of the row that was deleted, echoed back.
    Id string `json:"id"`

    // Used by Decode() method
    data []byte
}

func (model OrderDeleted) New(data []byte) *OrderDeleted {
    model.data = data
    return &model
}

func (model *OrderDeleted) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}