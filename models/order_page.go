package models

import (
    "encoding/json"
    "errors"
)

// OrderPage Where this answer sits in the whole result set.
type OrderPage struct {
    // Whether another page exists after this one (offset + returned < total). The
    // one field a "load more" button should read.
    HasMore bool `json:"hasMore"`
    // The page size that was applied. A requested limit above 200 is CLAMPED to
    // 200 rather than refused, so this is the number to believe, not the one you
    // sent.
    Limit int `json:"limit"`
    // The row offset that was applied.
    Offset int `json:"offset"`
    // How many rows are in `items` right here — less than `limit` on the last
    // page.
    Returned int `json:"returned"`
    // How many rows match the filter in total, ignoring limit and offset. This is
    // what a page count is computed from.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model OrderPage) New(data []byte) *OrderPage {
    model.data = data
    return &model
}

func (model *OrderPage) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}