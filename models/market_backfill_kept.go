package models

import (
    "encoding/json"
    "errors"
)

// MarketBackfillKept What this market already held BEFORE the repair, per
// collection — the rows that were left exactly as the merchant left them.
type MarketBackfillKept struct {
    // Traded currencies this market already held, untouched.
    Currencies int `json:"currencies"`
    // Locales this market already held, untouched.
    Locales int `json:"locales"`
    // Tax classes this market already held, untouched.
    TaxClasses int `json:"tax_classes"`

    // Used by Decode() method
    data []byte
}

func (model MarketBackfillKept) New(data []byte) *MarketBackfillKept {
    model.data = data
    return &model
}

func (model *MarketBackfillKept) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}