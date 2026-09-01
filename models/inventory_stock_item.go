package models

import (
    "encoding/json"
    "errors"
)

// InventoryStockItem One item and how much of it: 'product_id' or 'sku', plus
// a positive quantity.
type InventoryStockItem struct {
    // The product to move, as the products app knows it. Give this OR `sku` —
    // an item that names neither is answered 400. Matching is exact: a stock row
    // keyed by SKU is not found by product id.
    ProductId string `json:"product_id"`
    // How many units this booking moves. Always POSITIVE here — the direction
    // is the route (receive adds, reserve holds, restock returns), not the sign.
    // Zero or a negative number is answered 400; a signed correction is what POST
    // /inventories/adjust is for.
    Quantity float64 `json:"quantity"`
    // The article number to move, when the item has no product id. Give this OR
    // `product_id`.
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model InventoryStockItem) New(data []byte) *InventoryStockItem {
    model.data = data
    return &model
}

func (model *InventoryStockItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}