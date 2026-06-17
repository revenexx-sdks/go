package models

import (
    "encoding/json"
    "errors"
)

// AnEntryNeedsAnIdentityProductIdOrSkuEveryOtherFieldIsNormalizedToItsDefaultWhenNullOmitted
// Model
type PriceEntryReplaceItem struct {
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // Default 'standard'; 'on_request' is the explicit no-price marker — it
    // stops resolution and answers "price on request".
    PriceType string `json:"price_type"`
    // Priced product.
    ProductId string `json:"product_id"`
    // Tier threshold (Staffelpreis): this price applies from this quantity
    // (default 1).
    QuantityMin float64 `json:"quantity_min"`
    // Priced SKU (alternative to product_id).
    Sku string `json:"sku"`
    // 
    Unit string `json:"unit"`
    // Per-unit price (default 0).
    UnitPrice float64 `json:"unit_price"`
    // Per-entry validity start (promo prices).
    ValidFrom string `json:"valid_from"`
    // Per-entry validity end.
    ValidUntil string `json:"valid_until"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntryReplaceItem) New(data []byte) *PriceEntryReplaceItem {
    model.data = data
    return &model
}

func (model *PriceEntryReplaceItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}