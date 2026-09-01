package models

import (
    "encoding/json"
    "errors"
)

// OrderNumberRangeCreateRequest Number pattern: '{prefix}{counter padded to
// padding}{suffix}'.
type OrderNumberRangeCreateRequest struct {
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
    // answers 409). Defaults to 0, so the first number drawn is step.
    Counter int `json:"counter"`
    // Free-form data for the caller. This app stores it and returns it, and reads
    // nothing out of it.
    Metadata interface{} `json:"metadata"`
    // How wide the counter is written, zero-padded: 6 makes 123 into 000123. 0
    // writes the bare number. Widening it later does not renumber what was
    // already drawn. Defaults to 6.
    Padding int `json:"padding"`
    // The gap between the position numbers of a new order: 10 numbers the lines
    // 10, 20, 30 — room to slot a line in between later without renumbering the
    // rest. Read from the ORDER range only. Defaults to 10.
    PositionStep int `json:"position_step"`
    // Literal text in front of the counter: 'ORD-' turns counter 123 into
    // ORD-000123. Empty by default. Defaults to ''.
    Prefix string `json:"prefix"`
    // How far the counter moves per draw. 1 is consecutive numbering; a larger
    // step is what a merchant chooses who does not want their order volume
    // readable off an invoice. Defaults to 1.
    Step int `json:"step"`
    // Literal text after the counter — a market or year marker on merchants who
    // number that way. Empty by default, which is what most of them use. Defaults
    // to ''.
    Suffix string `json:"suffix"`

    // Used by Decode() method
    data []byte
}

func (model OrderNumberRangeCreateRequest) New(data []byte) *OrderNumberRangeCreateRequest {
    model.data = data
    return &model
}

func (model *OrderNumberRangeCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}