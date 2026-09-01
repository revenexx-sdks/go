package models

import (
    "encoding/json"
    "errors"
)

// StockLevelsFilter The exact-column filters this call was understood to
// carry, verbatim as they arrived. A query parameter that is not a column of
// `stock_levels` — a typo, a filter another entity has, `?q=` — is
// DROPPED and cannot appear here, and the list comes back unfiltered. This
// object is the only way to tell that apart from "nothing matched".
type StockLevelsFilter struct {
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?location_id=` value this call was understood to carry.
    LocationId string `json:"location_id"`
    // The literal `?metadata=` value this call was understood to carry.
    Metadata string `json:"metadata"`
    // The literal `?on_hand=` value this call was understood to carry.
    OnHand string `json:"on_hand"`
    // The literal `?product_id=` value this call was understood to carry.
    ProductId string `json:"product_id"`
    // The literal `?reorder_point=` value this call was understood to carry.
    ReorderPoint string `json:"reorder_point"`
    // The literal `?reserved=` value this call was understood to carry.
    Reserved string `json:"reserved"`
    // The literal `?sku=` value this call was understood to carry.
    Sku string `json:"sku"`
    // The literal `?updated_at=` value this call was understood to carry.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model StockLevelsFilter) New(data []byte) *StockLevelsFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *StockLevelsFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}