package models

import (
    "encoding/json"
    "errors"
)

// MarketBackfillAdded Child rows copied in from the source, per collection
// — only codes this market did not already carry. Zero everywhere on a
// second run: the call is idempotent.
type MarketBackfillAdded struct {
    // Traded currencies added from the source market.
    Currencies int `json:"currencies"`
    // Locales added from the source market.
    Locales int `json:"locales"`
    // Tax classes added from the source market.
    TaxClasses int `json:"tax_classes"`

    // Used by Decode() method
    data []byte
}

func (model MarketBackfillAdded) New(data []byte) *MarketBackfillAdded {
    model.data = data
    return &model
}

func (model *MarketBackfillAdded) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}