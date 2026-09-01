package models

import (
    "encoding/json"
    "errors"
)

// MarketsPage Where in the result set this answer sits. `limit` and `offset`
// are the values that were APPLIED, not the ones that were asked for — the
// data plane clamps rather than refuses, so an out-of-range or unparseable
// value comes back corrected here instead of as a 400.
type MarketsPage struct {
    // True when `offset + returned < total`, i.e. another page exists. Cheaper to
    // branch on than comparing the three numbers yourself.
    HasMore bool `json:"hasMore"`
    // Page size actually applied. A request over 200 is clamped to 200, one under
    // 1 (or one that is not a number) to the 50-row default.
    Limit int `json:"limit"`
    // Row offset actually applied. A negative offset is clamped to 0.
    Offset int `json:"offset"`
    // Rows in `items` on this page. Lower than `limit` on the last page.
    Returned int `json:"returned"`
    // Rows matching the filter across ALL pages, ignoring limit and offset —
    // the number to paginate against.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model MarketsPage) New(data []byte) *MarketsPage {
    model.data = data
    return &model
}

func (model *MarketsPage) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}