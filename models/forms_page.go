package models

import (
    "encoding/json"
    "errors"
)

// FormsPage Where this page sits in the result set. Everything needed to
// fetch the next one is here, so a client never has to guess whether it has
// seen everything.
type FormsPage struct {
    // True while `offset + returned < total`: another page follows, at `offset +
    // returned`.
    HasMore bool `json:"hasMore"`
    // The page size that was applied — the `limit` parameter after clamping to
    // 1…200, or 50 when none was given.
    Limit int `json:"limit"`
    // How many matching rows were skipped before this page.
    Offset int `json:"offset"`
    // How many rows are in `items` — below `limit` exactly on the last page.
    Returned int `json:"returned"`
    // How many rows match the filter in total, ignoring the page. This is the
    // number to show a merchant; `returned` is only what fitted.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model FormsPage) New(data []byte) *FormsPage {
    model.data = data
    return &model
}

func (model *FormsPage) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}