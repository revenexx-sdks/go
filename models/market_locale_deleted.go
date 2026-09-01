package models

import (
    "encoding/json"
    "errors"
)

// MarketLocaleDeleted Confirmation that the locale of a market is gone. The
// row itself is not returned — read it before deleting if you need it.
type MarketLocaleDeleted struct {
    // Always true — a row that was not there is a 404, not a false.
    Deleted bool `json:"deleted"`
    // The id of the row that was deleted.
    Id string `json:"id"`

    // Used by Decode() method
    data []byte
}

func (model MarketLocaleDeleted) New(data []byte) *MarketLocaleDeleted {
    model.data = data
    return &model
}

func (model *MarketLocaleDeleted) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}