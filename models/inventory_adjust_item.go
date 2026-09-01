package models

import (
    "encoding/json"
    "errors"
)

// InventoryAdjustItem One item and its SIGNED correction: 'product_id' or
// 'sku', plus a non-zero delta.
type InventoryAdjustItem struct {
    // The product to move, as the products app knows it. Give this OR `sku` —
    // an item that names neither is answered 400. Matching is exact: a stock row
    // keyed by SKU is not found by product id.
    ProductId string `json:"product_id"`
    // The SIGNED correction to `on_hand`: −3 writes off three, +3 finds three.
    // It is a delta, not the new balance. Zero is refused (400) because a
    // correction of nothing is a mistake, not a booking — the rule is the
    // handler's, not a database CHECK, which is why it is stated here rather than
    // declared as a bound.
    Quantity float64 `json:"quantity"`
    // The article number to move, when the item has no product id. Give this OR
    // `product_id`.
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model InventoryAdjustItem) New(data []byte) *InventoryAdjustItem {
    model.data = data
    return &model
}

func (model *InventoryAdjustItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}