package models

import (
    "encoding/json"
    "errors"
)

// PriceAdjustPreviewRow One entry, before and after — the row a
// confirmation dialog shows.
type PriceAdjustPreviewRow struct {
    // The price entry this row is about.
    Id string `json:"id"`
    // After rounding and ending snapping, in the same currency and on the same
    // basis. Never negative: below the lowest candidate ending it clamps to it.
    NewUnitPrice float64 `json:"new_unit_price"`
    // The product it prices — null when the entry is identified by SKU.
    ProductId string `json:"product_id"`
    // Which rung of the ladder this is.
    QuantityMin float64 `json:"quantity_min"`
    // The SKU it prices — null when the entry is identified by product id.
    Sku string `json:"sku"`
    // Before the change, in the list’s currency and on its tax basis.
    UnitPrice float64 `json:"unit_price"`

    // Used by Decode() method
    data []byte
}

func (model PriceAdjustPreviewRow) New(data []byte) *PriceAdjustPreviewRow {
    model.data = data
    return &model
}

func (model *PriceAdjustPreviewRow) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}