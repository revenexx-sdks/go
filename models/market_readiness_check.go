package models

import (
    "encoding/json"
    "errors"
)

// MarketReadinessCheck One question asked of the market, its verdict, and how
// much the answer costs.
type MarketReadinessCheck struct {
    // One sentence naming what was found and, for a warning, what covers for it.
    Detail string `json:"detail"`
    // Which question. 'locales' — is there a language to render in?
    // 'currencies' — is the base currency registered and marked default?
    // 'tax_classes' — is there a rate to tax with? 'tax_basis' —
    // informational, restating whether stored prices are gross or net.
    Id string `json:"id"`
    // Whether this check passed. A false with severity `info` cannot occur —
    // the informational check always passes.
    Ok bool `json:"ok"`
    // What a failure costs. 'blocking' — the market cannot trade. 'warning' —
    // degraded but serviceable, and `detail` names what covers for it. 'info' —
    // a fact worth reporting that is never a failure. The severity is not fixed
    // per check: no locales is blocking without a tenant fallback_locale and a
    // warning with one.
    Severity string `json:"severity"`

    // Used by Decode() method
    data []byte
}

func (model MarketReadinessCheck) New(data []byte) *MarketReadinessCheck {
    model.data = data
    return &model
}

func (model *MarketReadinessCheck) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}