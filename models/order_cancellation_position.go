package models

import (
    "encoding/json"
    "errors"
)

// OrderCancellationPosition One position quantity this cancellation removed.
type OrderCancellationPosition struct {
    // The order item this quantity was booked against — an id out of the same
    // order, never another one.
    OrderItemId string `json:"order_item_id"`
    // The quantity booked on that position, in the position's own unit. Three
    // decimal places, so 0.5 m of cable is a real booking.
    Quantity float64 `json:"quantity"`

    // Used by Decode() method
    data []byte
}

func (model OrderCancellationPosition) New(data []byte) *OrderCancellationPosition {
    model.data = data
    return &model
}

func (model *OrderCancellationPosition) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}