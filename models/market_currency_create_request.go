package models

import (
    "encoding/json"
    "errors"
)

// MarketCurrencyCreateRequest The owning market comes from the route path
// ('market_id').
type MarketCurrencyCreateRequest struct {
    // ISO 4217 code, unique per market — one entry in the set of currencies
    // this market TRADES in, as opposed to the single base currency on the market
    // row that its prices are quoted in. The base currency must appear here or
    // the market cannot serve; clone and backfill register it for you.
    Code string `json:"code"`
    // The currency offered first to a buyer who states no preference. At most one
    // per market, and it should be the market's base currency — readiness
    // reports it as a warning when it is not.
    IsDefault bool `json:"is_default"`
    // Sort position among this market's currencies, ascending, default 0 — the
    // order a currency switcher lists them in.
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model MarketCurrencyCreateRequest) New(data []byte) *MarketCurrencyCreateRequest {
    model.data = data
    return &model
}

func (model *MarketCurrencyCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}