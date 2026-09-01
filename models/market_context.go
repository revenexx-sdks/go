package models

import (
    "encoding/json"
    "errors"
)

// MarketContext The whole of one market: the row, its three collections, and
// the four resolved answers a client would otherwise have to work out for
// itself.
type MarketContext struct {
    // Every currency this market trades in, in position order. Capped at 200. The
    // market's own base currency should be among them; readiness reports it as
    // blocking when it is not.
    Currencies []MarketCurrency `json:"currencies"`
    // The locale a storefront should render this market in. `source` names where
    // it came from: 'market' (a locale flagged is_default), 'market_first' (no
    // flag — first by position) or 'tenant_fallback' (the market registers
    // none; the tenant's fallback_locale setting answered).
    DefaultLocale MarketDefaultLocale `json:"default_locale"`
    // How this tenant keys its translations, resolved rather than named: the key
    // a client WRITES and the order it READS, per locale. Emitting the resolved
    // answer is the point — a client handed only the setting names
    // re-implements the policy and gets it subtly different, which is how a label
    // editor came to ask for de-DE while the row held de.
    LocalePolicy MarketLocalePolicy `json:"locale_policy"`
    // Every locale this market registers, in position order. Capped at 200. Empty
    // is a real answer — read `default_locale` before assuming a language.
    Locales []MarketLocale `json:"locales"`
    // A distinct business context within a tenant — a country, a region, or a
    // storefront segment such as B2C vs B2B — with its own base currency,
    // locales, traded currencies and tax classes. A market is also the platform's
    // `market` SCOPE dimension: every other commerce app slices its data by one,
    // keyed on this row's `code`. A market is never just this row: it needs at
    // least one locale, one currency and one tax class before it can serve, which
    // is what /readiness measures and what /clone and /backfill build.
    Market Market `json:"market"`
    // Whether a stored price in this market is NET or GROSS — the market layer
    // of an answer the prices app also holds. A price list's own tax_basis wins
    // over this; `tax_basis: null` with `source: 'unset'` means this market
    // declares nothing and the reader must fall through to the tenant's own
    // default.
    Pricing MarketPricing `json:"pricing"`
    // Can this market actually trade? `ready` is false only when a BLOCKING check
    // failed — no currency to quote in, no tax class to tax with. Warnings are
    // degraded-but-serviceable.
    Readiness MarketReadiness `json:"readiness"`
    // Every tax class of this market with its rate, in position order. Capped at
    // 200. This is the rate table other apps resolve a line against, by code.
    TaxClasses []MarketTaxClass `json:"tax_classes"`

    // Used by Decode() method
    data []byte
}

func (model MarketContext) New(data []byte) *MarketContext {
    model.data = data
    return &model
}

func (model *MarketContext) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}