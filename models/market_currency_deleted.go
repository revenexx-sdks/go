package models

import (
    "encoding/json"
    "errors"
)

// MarketCurrencyDeleted Confirmation that the currency of a market is gone.
// The row itself is not returned — read it before deleting if you need it.
type MarketCurrencyDeleted struct {
    // Always true — a row that was not there is a 404, not a false.
    Deleted bool `json:"deleted"`
    // The id of the row that was deleted.
    Id string `json:"id"`

    // Used by Decode() method
    data []byte
}

func (model MarketCurrencyDeleted) New(data []byte) *MarketCurrencyDeleted {
    model.data = data
    return &model
}

func (model *MarketCurrencyDeleted) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}