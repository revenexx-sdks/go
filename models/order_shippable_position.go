package models

import (
    "encoding/json"
    "errors"
)

// OrderShippablePosition One order position with the quantity that may still
// be shipped, and the three numbers that quantity is made of. Every position
// of the order is here, including the ones with nothing left open — a
// dialog needs to show a fully shipped line as fully shipped, not omit it.
type OrderShippablePosition struct {
    // The article name as it stood at place-time, frozen. Falls back to the sku
    // when the caller sent none — a position always reads as something.
    Name string `json:"name"`
    // The position, by the id a positions[] payload names it with. This is what
    // POST /orders/{id}/ship expects — copy it, do not construct it.
    OrderItemId string `json:"order_item_id"`
    // The line number a human reads, and what the order is sorted by. Numbered in
    // steps of the range's position_step (10, 20, 30) unless the caller set it
    // explicitly — the gap is what lets a line be inserted later without
    // renumbering.
    Position int `json:"position"`
    // The catalog product this line was taken from (the products app). Null on a
    // custom line, and it stays a reference — the position keeps working after
    // the product is retired.
    ProductId string `json:"product_id"`
    // How much was ORDERED on this position. Unchanged by anything that happens
    // afterwards.
    Quantity float64 `json:"quantity"`
    // How much was cancelled and will never go out.
    QuantityCancelled float64 `json:"quantity_cancelled"`
    // quantity − shipped − cancelled: the budget POST /orders/{id}/ship
    // guards this position against, and the largest quantity it will accept. Zero
    // means the line is done.
    QuantityOpen float64 `json:"quantity_open"`
    // How much has already gone out.
    QuantityShipped float64 `json:"quantity_shipped"`
    // The article number as it stood at place-time, frozen with the rest of the
    // line. The value an ERP and a warehouse both join on, and the one field a
    // picker reads. Null only on a line that never had one.
    Sku string `json:"sku"`
    // The unit the quantity is counted in — piece, metre, kilogram, package.
    // Free text as the catalog carries it; this app does no conversion.
    Unit string `json:"unit"`

    // Used by Decode() method
    data []byte
}

func (model OrderShippablePosition) New(data []byte) *OrderShippablePosition {
    model.data = data
    return &model
}

func (model *OrderShippablePosition) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}