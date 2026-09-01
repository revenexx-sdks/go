package models

import (
    "encoding/json"
    "errors"
)

// MarketLocale One language a market is rendered in, and one key its
// translations are stored under. A market may register several; one of them
// is the default a storefront falls back to.
type MarketLocale struct {
    // Locale code, language-COUNTRY — the language a storefront renders this
    // market in, and the key a translation is stored under. Unique per market.
    // The app's own seeded value is the tenant's `fallback_locale` setting, whose
    // declared default is de-DE.
    Code string `json:"code"`
    // ISO 3166-1 alpha-2 country code — the region half of `code`. It is a
    // spelling of the language, not a shipping destination: a market may register
    // de-AT without trading in Austria.
    Country string `json:"country"`
    // When the locale was registered on this market. Set by the database; never
    // writable.
    CreatedAt string `json:"created_at"`
    // Primary key of this locale registration. The locale is named by `code`
    // everywhere else.
    Id string `json:"id"`
    // The locale a storefront renders this market in when the request asks for
    // none. At most one per market; where none carries the flag the first by
    // position is used, and `default_locale.source` on the context says which of
    // the two happened.
    IsDefault bool `json:"is_default"`
    // ISO 639-1 language code — the language half of `code`, stored separately
    // so a client can group markets by language without parsing.
    Language string `json:"language"`
    // The market this locale belongs to. Filled from the route path on write and
    // never read out of the body; ON DELETE CASCADE, so deleting the market
    // deletes this row.
    MarketId string `json:"market_id"`
    // Sort position among this market's locales, ascending, default 0 — and the
    // tie-break that picks a default when no locale is flagged.
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model MarketLocale) New(data []byte) *MarketLocale {
    model.data = data
    return &model
}

func (model *MarketLocale) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}