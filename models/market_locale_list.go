package models

import (
    "encoding/json"
    "errors"
)

// MarketLocaleList One page of locales of a market, the page it sits on, and
// the filters that produced it.
type MarketLocaleList struct {
    // The exact-column filters this call applied, echoed back. Every value is the
    // raw query string, never the column's own type: `?is_default=true` comes
    // back as `"true"`. A `?column=value` naming a column this entity does not
    // have is DROPPED rather than refused — the call answers 200 with the
    // unfiltered list, and the key missing from here is the only way to find out.
    Filter MarketLocaleFilter `json:"filter"`
    // The locales of a market on this page, in `order` — by `position`
    // ascending unless the call asked otherwise.
    Items []MarketLocale `json:"items"`
    // Where in the result set this answer sits. `limit` and `offset` are the
    // values that were APPLIED, not the ones that were asked for — the data
    // plane clamps rather than refuses, so an out-of-range or unparseable value
    // comes back corrected here instead of as a 400.
    Page MarketsPage `json:"page"`

    // Used by Decode() method
    data []byte
}

func (model MarketLocaleList) New(data []byte) *MarketLocaleList {
    model.data = data
    return &model
}

func (model *MarketLocaleList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}