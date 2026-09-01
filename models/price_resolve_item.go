package models

import (
    "encoding/json"
    "errors"
)

// PriceResolveItem Identify by 'product_id' or 'sku' — an item without
// identity resolves to on_request with a per-item error rather than failing
// the call.
type PriceResolveItem struct {
    // Product to price.
    ProductId string `json:"product_id"`
    // Requested quantity, counted in the entry’s `unit`. It picks the tier (the
    // highest `quantity_min` at or below it) and multiplies into `line_total`.
    // Default 1; a non-positive value falls back to 1.
    Quantity float64 `json:"quantity"`
    // SKU to price (alternative to product_id). Matched exactly against the
    // entries’ own `sku`.
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model PriceResolveItem) New(data []byte) *PriceResolveItem {
    model.data = data
    return &model
}

func (model *PriceResolveItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}