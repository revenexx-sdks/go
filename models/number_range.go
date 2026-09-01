package models

import (
    "encoding/json"
    "errors"
)

// NumberRange A counter that issues human-readable numbers, one per series:
// orders, delivery notes, returns. The format is {prefix}{counter padded to
// padding}{suffix}, and drawing a number moves the counter.
type NumberRange struct {
    // The sales channel this range was created for, as a label. It does NOT
    // select the range: a draw finds the range by `code` alone, and the unique
    // index (tenant, code) means one code is one range per tenant — so an order
    // on another channel draws from the same range this one names. Null on the
    // three seeded ranges, which is every tenant-wide range.
    ChannelId string `json:"channel_id"`
    // Which counter this is, in the app's own words: 'order' numbers orders,
    // 'delivery' numbers delivery notes, 'return' numbers returns. Unique per
    // tenant, and the value the order_number_range_code /
    // delivery_number_range_code / return_number_range_code settings point at —
    // a setting naming a code no range carries is the 422 'number_range_missing'.
    Code string `json:"code"`
    // The last number DRAWN — state, not configuration. The next draw is
    // counter + step and writes the new value back, so moving this forward skips
    // numbers and moving it back re-issues them (and the unique index then
    // answers 409).
    Counter int `json:"counter"`
    // When the range was created.
    CreatedAt string `json:"created_at"`
    // Primary key of the number range.
    Id string `json:"id"`
    // Free-form data for the caller. This app stores it and returns it, and reads
    // nothing out of it.
    Metadata interface{} `json:"metadata"`
    // How wide the counter is written, zero-padded: 6 makes 123 into 000123. 0
    // writes the bare number. Widening it later does not renumber what was
    // already drawn.
    Padding int `json:"padding"`
    // The gap between the position numbers of a new order: 10 numbers the lines
    // 10, 20, 30 — room to slot a line in between later without renumbering the
    // rest. Read from the ORDER range only.
    PositionStep int `json:"position_step"`
    // Literal text in front of the counter: 'ORD-' turns counter 123 into
    // ORD-000123. Empty by default.
    Prefix string `json:"prefix"`
    // How far the counter moves per draw. 1 is consecutive numbering; a larger
    // step is what a merchant chooses who does not want their order volume
    // readable off an invoice.
    Step int `json:"step"`
    // Literal text after the counter — a market or year marker on merchants who
    // number that way. Empty by default, which is what most of them use.
    Suffix string `json:"suffix"`
    // When the range last changed — which includes every single number draw,
    // because a draw writes the counter.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model NumberRange) New(data []byte) *NumberRange {
    model.data = data
    return &model
}

func (model *NumberRange) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}