package models

import (
    "encoding/json"
    "errors"
)

// MarketBackfillRequest The path id is the market being REPAIRED; `source` is
// the market to copy from (a uuid or a market code). The three flags default
// to true.
type MarketBackfillRequest struct {
    // Take the source's traded currencies for codes this market does not already
    // carry. Default true.
    Currencies bool `json:"currencies"`
    // Take the source's locales for codes this market does not already carry.
    // Default true.
    Locales bool `json:"locales"`
    // The market to copy the missing pieces FROM — a uuid or a market code.
    // Must not be the market in the path. Pick a market that is already right;
    // nothing about it is changed.
    Source string `json:"source"`
    // Take the source's tax classes for codes this market does not already carry.
    // An existing code keeps ITS rate — a backfill never re-rates a class the
    // merchant already set. Default true.
    TaxClasses bool `json:"tax_classes"`

    // Used by Decode() method
    data []byte
}

func (model MarketBackfillRequest) New(data []byte) *MarketBackfillRequest {
    model.data = data
    return &model
}

func (model *MarketBackfillRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}