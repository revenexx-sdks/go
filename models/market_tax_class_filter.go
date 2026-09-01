package models

import (
    "encoding/json"
    "errors"
)

// MarketTaxClassFilter The exact-column filters this call applied, echoed
// back. Every value is the raw query string, never the column's own type:
// `?is_default=true` comes back as `"true"`. A `?column=value` naming a
// column this entity does not have is DROPPED rather than refused — the
// call answers 200 with the unfiltered list, and the key missing from here is
// the only way to find out.
type MarketTaxClassFilter struct {
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
    // The `labels` filter as it arrived, verbatim. Present only when the call
    // sent it.
    Labels string `json:"labels"`
    // The owning market, taken from the route path. ALWAYS present, and always
    // the path's market — a `?market_id=` in the query is overwritten by it
    // rather than honoured, so this is never the value a caller sent.
    MarketId string `json:"market_id"`
    // The `name` filter as it arrived, verbatim. Present only when the call sent
    // it.
    Name string `json:"name"`
    // The `position` filter as it arrived, verbatim. Present only when the call
    // sent it.
    Position string `json:"position"`
    // The `rate` filter as it arrived, verbatim. Present only when the call sent
    // it.
    Rate string `json:"rate"`
    // The `updated_at` filter as it arrived, verbatim. Present only when the call
    // sent it. Any form the database accepts as a timestamp, including a bare
    // date.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model MarketTaxClassFilter) New(data []byte) *MarketTaxClassFilter {
    model.data = data
    return &model
}

func (model *MarketTaxClassFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}