package models

import (
    "encoding/json"
    "errors"
)

// MarketTaxClassDeleted Confirmation that the tax class of a market is gone.
// The row itself is not returned — read it before deleting if you need it.
type MarketTaxClassDeleted struct {
    // Always true — a row that was not there is a 404, not a false.
    Deleted bool `json:"deleted"`
    // The id of the row that was deleted.
    Id string `json:"id"`
    // False when the cross-app usage question could not be asked (shipping not
    // installed, or unreachable) — the row was deleted without that guarantee.
    UsageChecked bool `json:"usage_checked"`

    // Used by Decode() method
    data []byte
}

func (model MarketTaxClassDeleted) New(data []byte) *MarketTaxClassDeleted {
    model.data = data
    return &model
}

func (model *MarketTaxClassDeleted) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}