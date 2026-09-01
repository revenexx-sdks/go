package models

import (
    "encoding/json"
    "errors"
)

// MarketCloneSeeded Rows this call added that were copied from nowhere,
// because the new market would otherwise have been left unable to trade: the
// tenant `fallback_locale` when neither market had a locale, and the base
// currency when it is not in the copied set. Zero on both is the normal,
// healthy answer — it means nothing had to be invented.
type MarketCloneSeeded struct {
    // 1 when the market's own base currency was registered because the copied set
    // did not contain it; 0 otherwise.
    Currencies int `json:"currencies"`
    // 1 when the tenant's fallback_locale was written as this market's only
    // locale, marked default; 0 otherwise.
    Locales int `json:"locales"`

    // Used by Decode() method
    data []byte
}

func (model MarketCloneSeeded) New(data []byte) *MarketCloneSeeded {
    model.data = data
    return &model
}

func (model *MarketCloneSeeded) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}