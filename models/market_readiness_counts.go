package models

import (
    "encoding/json"
    "errors"
)

// MarketReadinessCounts How much of a market this market actually is. All
// three at zero is a market that is a row and nothing else — the state two
// of the three live markets on the platform were left in, and the reason
// /clone and /backfill exist.
type MarketReadinessCounts struct {
    // Traded currencies registered on this market.
    Currencies int `json:"currencies"`
    // Locales registered on this market.
    Locales int `json:"locales"`
    // Tax classes registered on this market.
    TaxClasses int `json:"tax_classes"`

    // Used by Decode() method
    data []byte
}

func (model MarketReadinessCounts) New(data []byte) *MarketReadinessCounts {
    model.data = data
    return &model
}

func (model *MarketReadinessCounts) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}