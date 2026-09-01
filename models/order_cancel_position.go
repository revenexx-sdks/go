package models

import (
    "encoding/json"
    "errors"
)

// OrderCancelPosition A position quantity to cancel — guarded against the
// open (unshipped, uncancelled) quantity.
type OrderCancelPosition struct {
    // The order item (position) to act on. Read the ids from GET /orders/{id}
    // (items[].id) or GET /orders/{id}/shippable (positions[].order_item_id) —
    // an id this order does not carry is a 400.
    OrderItemId string `json:"order_item_id"`
    // Defaults to the full remaining quantity of the position.
    Quantity float64 `json:"quantity"`

    // Used by Decode() method
    data []byte
}

func (model OrderCancelPosition) New(data []byte) *OrderCancelPosition {
    model.data = data
    return &model
}

func (model *OrderCancelPosition) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}