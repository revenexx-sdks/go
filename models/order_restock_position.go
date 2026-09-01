package models

import (
    "encoding/json"
    "errors"
)

// OrderRestockPosition One quantity to put back into stock, named the way the
// inventories app wants it: by product, by sku, and how much.
type OrderRestockPosition struct {
    // The catalog product to restock. Null on a custom line, which is why `sku`
    // is carried alongside it.
    ProductId string `json:"product_id"`
    // How much came back on this position, in the position's own unit.
    Quantity float64 `json:"quantity"`
    // The article number to restock — the key a warehouse actually books
    // against.
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model OrderRestockPosition) New(data []byte) *OrderRestockPosition {
    model.data = data
    return &model
}

func (model *OrderRestockPosition) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}