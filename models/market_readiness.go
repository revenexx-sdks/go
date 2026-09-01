package models

import (
    "encoding/json"
    "errors"
)

// MarketReadiness Can this market actually trade? `ready` is false only when
// a BLOCKING check failed — no currency to quote in, no tax class to tax
// with. Warnings are degraded-but-serviceable.
type MarketReadiness struct {
    // Ids of the checks that failed BLOCKING — the market cannot do the job at
    // all until each is fixed. Empty exactly when `ready` is true.
    Blocking []string `json:"blocking"`
    // Every check that ran, passed or failed, in a fixed order: locales,
    // currencies, tax_classes, tax_basis. `blocking` and `warnings` are the
    // failures from this list by id; this is where the reason lives.
    Checks []MarketReadinessCheck `json:"checks"`
    // `blocking` is empty. Deliberately not "every check passed": a market with
    // one locale and no default flag on it is serviceable, and a verdict that
    // cried wolf about that would be ignored on the day it mattered.
    Ready bool `json:"ready"`
    // true when the market's status is 'active'. An active market that is not
    // ready is live and broken — that combination is the one worth an alert.
    Serving bool `json:"serving"`
    // Ids of the checks that failed as WARNINGS — degraded but serviceable,
    // because something else covers for them. A missing locale is only a warning
    // while the tenant declares a fallback_locale.
    Warnings []string `json:"warnings"`

    // Used by Decode() method
    data []byte
}

func (model MarketReadiness) New(data []byte) *MarketReadiness {
    model.data = data
    return &model
}

func (model *MarketReadiness) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}