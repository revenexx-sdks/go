package models

import (
    "encoding/json"
    "errors"
)

// MarketCurrency One currency a market accepts, as opposed to the single base
// currency on the market row that its prices are quoted in. The base currency
// must be registered here or the market cannot serve.
type MarketCurrency struct {
    // ISO 4217 code, unique per market — one entry in the set of currencies
    // this market TRADES in, as opposed to the single base currency on the market
    // row that its prices are quoted in. The base currency must appear here or
    // the market cannot serve; clone and backfill register it for you.
    Code string `json:"code"`
    // When the currency was registered on this market. Set by the database; never
    // writable.
    CreatedAt string `json:"created_at"`
    // Primary key of this currency registration. The currency is named by `code`
    // everywhere else.
    Id string `json:"id"`
    // The currency offered first to a buyer who states no preference. At most one
    // per market, and it should be the market's base currency — readiness
    // reports it as a warning when it is not.
    IsDefault bool `json:"is_default"`
    // The market this currency belongs to. Filled from the route path on write
    // and never read out of the body; ON DELETE CASCADE, so deleting the market
    // deletes this row.
    MarketId string `json:"market_id"`
    // Sort position among this market's currencies, ascending, default 0 — the
    // order a currency switcher lists them in.
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model MarketCurrency) New(data []byte) *MarketCurrency {
    model.data = data
    return &model
}

func (model *MarketCurrency) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}