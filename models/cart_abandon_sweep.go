package models

import (
    "encoding/json"
    "errors"
)

// CartAbandonSweep The first sweep: active carts nobody has touched since
// their market's window become abandoned. Nothing else in the platform ever
// stamps abandoned_at, so without this the abandonment funnel is empty by
// construction rather than empty because nobody abandons carts.
type CartAbandonSweep struct {
    // Carts actually marked. 0 on a dry run — see `found`.
    Abandoned int `json:"abandoned"`
    // The abandon_after_minutes of the TENANT baseline — what a cart in no
    // market ran on. 0 disables the sweep. Carts in a market were each held
    // against their own market's window, which may differ from this.
    AfterMinutes float64 `json:"after_minutes"`
    // This pass looked at as many carts as one pass looks at, so there may be
    // more behind them. The rest go on the next tick, oldest first — a backlog
    // is visible here rather than merely slow.
    Capped bool `json:"capped"`
    // The carts this sweep touched, so a merchant can look at them before or
    // after.
    CartIds []string `json:"cart_ids"`
    // Carts untouched since this instant were swept — the BASELINE cutoff. A
    // run no longer has one cutoff, because each cart was held against its own
    // market's clock; this is the one unassigned carts ran on.
    Cutoff string `json:"cutoff"`
    // At least one window in force (the baseline, or some market's). False means
    // every applicable window was 0 and nothing was even considered.
    Enabled bool `json:"enabled"`
    // Carts past their window. On a dry run this is the whole answer —
    // `abandoned` stays 0.
    Found int `json:"found"`
    // The market codes this pass came across, so an operator can see whose
    // windows were actually in play. Empty when no examined cart belongs to a
    // market.
    Markets []string `json:"markets"`

    // Used by Decode() method
    data []byte
}

func (model CartAbandonSweep) New(data []byte) *CartAbandonSweep {
    model.data = data
    return &model
}

func (model *CartAbandonSweep) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}