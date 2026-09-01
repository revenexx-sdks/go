package models

import (
    "encoding/json"
    "errors"
)

// PriceEntriesLadderRequest The quantity ladder (Staffelpreise) for ONE item,
// generated instead of typed: a price at the first tier and a discount
// compounded per tier. Identify the item with 'product_id' or 'sku'.
type PriceEntriesLadderRequest struct {
    // Price for ONE unit at the FIRST tier, in the list’s currency and on the
    // list’s tax basis — a decimal amount in major units (19.90), never minor
    // units/cents.
    BasePrice float64 `json:"base_price"`
    // Discount applied per tier, COMPOUNDED down the ladder rather than off the
    // base price: 5 gives 19.90 / 18.91 / 17.96. Default 0.
    DiscountPercent float64 `json:"discount_percent"`
    // The item the ladder prices.
    ProductId string `json:"product_id"`
    // Tier thresholds, ascending — an array of numbers or a comma-separated
    // string ('1, 10, 50'). Duplicates are collapsed and the set is sorted.
    // Default [1, 10, 50], at most 50 tiers.
    Quantities []float64 `json:"quantities"`
    // Default true: the item's existing entries in this list are removed first,
    // so the ladder IS the ladder. false appends.
    Replace bool `json:"replace"`
    // Ending the computed prices snap to (nearest match). Omit to use the
    // tenant's bulk_adjust_rounding setting.
    Rounding string `json:"rounding"`
    // The item the ladder prices (alternative to product_id).
    Sku string `json:"sku"`
    // Unit of measure carried onto every generated tier. Free text, neither
    // validated nor converted.
    Unit string `json:"unit"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntriesLadderRequest) New(data []byte) *PriceEntriesLadderRequest {
    model.data = data
    return &model
}

func (model *PriceEntriesLadderRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}