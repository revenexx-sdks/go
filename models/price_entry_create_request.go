package models

import (
    "encoding/json"
    "errors"
)

// PriceEntryCreateRequest An entry needs an identity: 'product_id' or 'sku'.
type PriceEntryCreateRequest struct {
    // Free-form bag: whatever JSON object you write round-trips exactly, and this
    // app never reads it. Its keys are yours.
    Metadata interface{} `json:"metadata"`
    // Default 'standard'; 'on_request' is the explicit no-price marker — it
    // STOPS resolution for this item on this list and answers "price on request"
    // even where a cheaper list exists.
    PriceType string `json:"price_type"`
    // The product this rung prices. An entry needs product_id or sku — the row
    // CHECK enforces it.
    ProductId string `json:"product_id"`
    // Tier threshold (Staffelpreis): this price applies from this quantity
    // upwards (default 1). The rungs of one item are the entries sharing its
    // identity; the highest threshold at or below the requested quantity wins.
    QuantityMin float64 `json:"quantity_min"`
    // The article number this rung prices (alternative to product_id). Matched
    // exactly on resolve — never normalised or case-folded.
    Sku string `json:"sku"`
    // Unit of measure the price is per — free text, neither validated nor
    // converted here. A resolve call’s `quantity` is counted in it.
    Unit string `json:"unit"`
    // Price for ONE unit of `unit`, in the LIST’s currency and on the LIST’s
    // tax basis — a decimal amount in major units (19.90), never minor
    // units/cents. Stored at 4 decimals and echoed back exactly as sent (default
    // 0).
    UnitPrice float64 `json:"unit_price"`
    // Start of this entry’s own validity (ISO 8601) — how a promo price is
    // expressed: a second rung, live only for its window. null = open-ended.
    ValidFrom string `json:"valid_from"`
    // End of this entry’s own validity; null = open-ended. Outside it the rung
    // is skipped and the ladder resolves as if it were not there.
    ValidUntil string `json:"valid_until"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntryCreateRequest) New(data []byte) *PriceEntryCreateRequest {
    model.data = data
    return &model
}

func (model *PriceEntryCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}