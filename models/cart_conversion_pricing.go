package models

import (
    "encoding/json"
    "errors"
)

// CartConversionPricing How price_snapshot_mode settled the two prices every
// line carries.
type CartConversionPricing struct {
    // Lines in the cart when it converted.
    Lines int `json:"lines"`
    // Lines the mode had to rewrite because snapshot and unit_price disagreed —
    // repriced in 'snapshot' mode, re-snapshotted in 'live' mode. A line whose
    // snapshot carries no readable price is never touched in either mode.
    LinesChanged int `json:"lines_changed"`
    // The tenant's price_snapshot_mode, as it ran. 'snapshot' books the order on
    // the price the buyer was shown; 'live' books it on the line's current
    // unit_price and rewrites the snapshot to agree, so the frozen line never
    // claims a price nobody was charged.
    Mode string `json:"mode"`
    // The cart's frozen subtotal, and what the order is booked on.
    SubtotalAfter float64 `json:"subtotal_after"`
    // The cart's subtotal as it stood before the mode was applied. Compare it
    // with subtotal_after and 'why is the order €4 off the cart' is answered by
    // the response instead of by an argument.
    SubtotalBefore float64 `json:"subtotal_before"`

    // Used by Decode() method
    data []byte
}

func (model CartConversionPricing) New(data []byte) *CartConversionPricing {
    model.data = data
    return &model
}

func (model *CartConversionPricing) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}