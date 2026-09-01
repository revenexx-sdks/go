package models

import (
    "encoding/json"
    "errors"
)

// MarketCloneRequest The path id is the SOURCE market (a uuid or a market
// code). Everything the new market does not inherit is here. The copy flags
// default to true; `is_default` is never copied, and the new market always
// gets its own base currency registered and marked default.
type MarketCloneRequest struct {
    // Code of the NEW market (unique per tenant).
    Code string `json:"code"`
    // Copy the source's traded currencies. Default true. The new market's own
    // base currency is registered and marked default either way.
    CopyCurrencies bool `json:"copy_currencies"`
    // Copy the source's locales. Default true. False leaves the new market with
    // no language of its own, so the tenant fallback_locale is seeded instead —
    // it is never left with none.
    CopyLocales bool `json:"copy_locales"`
    // Copy the source's tax classes, rates and all. Default true. False leaves
    // the market unable to tax anything, which readiness reports as blocking.
    CopyTaxClasses bool `json:"copy_tax_classes"`
    // Base currency of the new market (ISO 4217). Defaults to the source
    // market's, and is registered and marked default on the new one either way.
    Currency string `json:"currency"`
    // Display name of the new market. Defaults to its code.
    Name string `json:"name"`
    // Status of the new market. Defaults to 'active'; clone it 'inactive' to
    // build it out before it serves anyone.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model MarketCloneRequest) New(data []byte) *MarketCloneRequest {
    model.data = data
    return &model
}

func (model *MarketCloneRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}