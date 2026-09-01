package models

import (
    "encoding/json"
    "errors"
)

// PriceListRef The price list this answer came out of — enough to link to
// it or to explain the number to a merchant ("this came from the dealer
// list").
type PriceListRef struct {
    // The list’s unique per-tenant code.
    Code string `json:"code"`
    // The list, by id — the same id `GET /prices/lists/{id}` takes.
    Id string `json:"id"`

    // Used by Decode() method
    data []byte
}

func (model PriceListRef) New(data []byte) *PriceListRef {
    model.data = data
    return &model
}

func (model *PriceListRef) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}