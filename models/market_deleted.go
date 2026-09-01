package models

import (
    "encoding/json"
    "errors"
)

// MarketDeleted Confirmation that the market is gone. The row itself is not
// returned — read it before deleting if you need it.
type MarketDeleted struct {
    // Always true — a row that was not there is a 404, not a false.
    Deleted bool `json:"deleted"`
    // The id of the row that was deleted.
    Id string `json:"id"`

    // Used by Decode() method
    data []byte
}

func (model MarketDeleted) New(data []byte) *MarketDeleted {
    model.data = data
    return &model
}

func (model *MarketDeleted) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}