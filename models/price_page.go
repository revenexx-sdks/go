package models

import (
    "encoding/json"
    "errors"
)

// PricePage Where this page sits in the full result set. Rows beyond `limit`
// are not returned and are not lost — ask for the next page with `offset`.
type PricePage struct {
    // true when `offset + returned < total` — there is another page to fetch.
    HasMore bool `json:"hasMore"`
    // Page size actually applied — the `limit` you sent, clamped to 1…200
    // (default 50).
    Limit int `json:"limit"`
    // Row offset actually applied (default 0).
    Offset int `json:"offset"`
    // Rows in `items` on this page.
    Returned int `json:"returned"`
    // Rows matching the filter across all pages, not just this one.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model PricePage) New(data []byte) *PricePage {
    model.data = data
    return &model
}

func (model *PricePage) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}