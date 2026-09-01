package models

import (
    "encoding/json"
    "errors"
)

// MarketCloneResult What was built. `copied` and `seeded` account for every
// child row that now exists, and `readiness` is the verdict on the result —
// so the call that made the market also tells you whether it finished the
// job.
type MarketCloneResult struct {
    // Child rows copied from the source, per collection. A flag left false is a
    // zero here, and so is a source that had none of that kind.
    Copied MarketCloneCopied `json:"copied"`
    // A distinct business context within a tenant — a country, a region, or a
    // storefront segment such as B2C vs B2B — with its own base currency,
    // locales, traded currencies and tax classes. A market is also the platform's
    // `market` SCOPE dimension: every other commerce app slices its data by one,
    // keyed on this row's `code`. A market is never just this row: it needs at
    // least one locale, one currency and one tax class before it can serve, which
    // is what /readiness measures and what /clone and /backfill build.
    Market Market `json:"market"`
    // Can this market actually trade? `ready` is false only when a BLOCKING check
    // failed — no currency to quote in, no tax class to tax with. Warnings are
    // degraded-but-serviceable.
    Readiness MarketReadiness `json:"readiness"`
    // Rows this call added that were copied from nowhere, because the new market
    // would otherwise have been left unable to trade: the tenant
    // `fallback_locale` when neither market had a locale, and the base currency
    // when it is not in the copied set. Zero on both is the normal, healthy
    // answer — it means nothing had to be invented.
    Seeded MarketCloneSeeded `json:"seeded"`
    // The market that was read from, resolved — so a caller who passed a code
    // back gets the uuid, and one who passed a uuid gets the code the rest of the
    // platform stores.
    Source MarketRef `json:"source"`

    // Used by Decode() method
    data []byte
}

func (model MarketCloneResult) New(data []byte) *MarketCloneResult {
    model.data = data
    return &model
}

func (model *MarketCloneResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}