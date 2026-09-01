package models

import (
    "encoding/json"
    "errors"
)

// MarketCurrencyFilter The exact-column filters this call applied, echoed
// back. Every value is the raw query string, never the column's own type:
// `?is_default=true` comes back as `"true"`. A `?column=value` naming a
// column this entity does not have is DROPPED rather than refused — the
// call answers 200 with the unfiltered list, and the key missing from here is
// the only way to find out.
type MarketCurrencyFilter struct {
    // The `code` filter as it arrived, verbatim. Present only when the call sent
    // it.
    Code string `json:"code"`
    // The `created_at` filter as it arrived, verbatim. Present only when the call
    // sent it. Any form the database accepts as a timestamp, including a bare
    // date.
    CreatedAt string `json:"created_at"`
    // The `id` filter as it arrived, verbatim. Present only when the call sent
    // it.
    Id string `json:"id"`
    // The `is_default` filter as it arrived, verbatim. Present only when the call
    // sent it.
    IsDefault string `json:"is_default"`
    // The owning market, taken from the route path. ALWAYS present, and always
    // the path's market — a `?market_id=` in the query is overwritten by it
    // rather than honoured, so this is never the value a caller sent.
    MarketId string `json:"market_id"`
    // The `position` filter as it arrived, verbatim. Present only when the call
    // sent it.
    Position string `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model MarketCurrencyFilter) New(data []byte) *MarketCurrencyFilter {
    model.data = data
    return &model
}

func (model *MarketCurrencyFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}