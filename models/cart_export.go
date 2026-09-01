package models

import (
    "encoding/json"
    "errors"
)

// CartExport model.
type CartExport struct {
    // The export itself. For json: `{ "cart": { name, status, currency,
    // channel_id, item_count, subtotal }, "items": [ … ] }` — exactly what
    // carts.import takes back, so an export round-trips. For csv: the lines as a
    // CSV string, header first, with jsonb columns serialized as JSON text.
    // Deliberately untyped, because a profile's mapping renames the columns and
    // that mapping is the caller's own.
    Content string `json:"content"`
    // A suggested download name, built as `cart-<cart id>.<format>`. Nothing is
    // stored under it; it is there so a browser download has a name that says
    // which cart it is.
    Filename string `json:"filename"`
    // The format that ran — the profile's, or the ad-hoc one.
    Format string `json:"format"`

    // Used by Decode() method
    data []byte
}

func (model CartExport) New(data []byte) *CartExport {
    model.data = data
    return &model
}

func (model *CartExport) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}