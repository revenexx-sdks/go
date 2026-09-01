package models

import (
    "encoding/json"
    "errors"
)

// InventoryAdjustRequest model.
type InventoryAdjustRequest struct {
    // The corrections, at most 200 in one call — a stocktake, breakage,
    // shrinkage. Quantities are SIGNED deltas, not new balances.
    Items []InventoryAdjustItem `json:"items"`
    // Which location is being corrected. Omitted, the `default_location_code`
    // setting decides. A correction is per location: the same SKU in two
    // warehouses is two corrections.
    LocationCode string `json:"location_code"`
    // Inline single-item form: the product to move, instead of a one-entry
    // `items` array. The two forms are equivalent — nothing downstream knows
    // which arrived.
    ProductId string `json:"product_id"`
    // Inline single-item form: the SIGNED correction (negative writes stock off,
    // positive finds it). Non-zero.
    Quantity float64 `json:"quantity"`
    // Why the stock is being corrected — this is the audit trail a stocktake
    // leaves behind. Owed unless `movement_reason_required` is 'none' (its
    // default, 'adjustments', asks for one exactly here); missing where it is
    // owed, the call is 400.
    Reason string `json:"reason"`
    // Inline single-item form: the article number to move (instead of
    // `product_id`).
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model InventoryAdjustRequest) New(data []byte) *InventoryAdjustRequest {
    model.data = data
    return &model
}

func (model *InventoryAdjustRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}