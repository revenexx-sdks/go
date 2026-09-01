package models

import (
    "encoding/json"
    "errors"
)

// OrderShipmentPosition A position quantity to ship — guarded against the
// open quantity.
type OrderShipmentPosition struct {
    // The order item (position) to act on. Read the ids from GET /orders/{id}
    // (items[].id) or GET /orders/{id}/shippable (positions[].order_item_id) —
    // an id this order does not carry is a 400.
    OrderItemId string `json:"order_item_id"`
    // Defaults to the full remaining quantity of the position.
    Quantity float64 `json:"quantity"`

    // Used by Decode() method
    data []byte
}

func (model OrderShipmentPosition) New(data []byte) *OrderShipmentPosition {
    model.data = data
    return &model
}

func (model *OrderShipmentPosition) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}