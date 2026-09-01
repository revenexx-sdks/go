package models

import (
    "encoding/json"
    "errors"
)

// InventoryAvailabilityItem One item to check: 'product_id' or 'sku'.
// Checking is free of consequence — it books nothing and holds nothing.
type InventoryAvailabilityItem struct {
    // The product to move, as the products app knows it. Give this OR `sku` —
    // an item that names neither is answered 400. Matching is exact: a stock row
    // keyed by SKU is not found by product id.
    ProductId string `json:"product_id"`
    // How many are wanted. It only decides `orderable`; the on_hand / reserved /
    // available figures come back whatever it is. Omit it (or send null) to ask
    // "is this sellable at all?", which is a check against 1.
    Quantity float64 `json:"quantity"`
    // The article number to move, when the item has no product id. Give this OR
    // `product_id`.
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model InventoryAvailabilityItem) New(data []byte) *InventoryAvailabilityItem {
    model.data = data
    return &model
}

func (model *InventoryAvailabilityItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}