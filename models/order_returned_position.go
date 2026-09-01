package models

import (
    "encoding/json"
    "errors"
)

// OrderReturnedPosition One position quantity registered for return.
type OrderReturnedPosition struct {
    // The order item this quantity was booked against — an id out of the same
    // order, never another one.
    OrderItemId string `json:"order_item_id"`
    // The quantity booked on that position, in the position's own unit. Three
    // decimal places, so 0.5 m of cable is a real booking.
    Quantity float64 `json:"quantity"`
    // Whether this quantity is reported for restocking when the return completes.
    // Restocking itself stays an explicit inventories.restock call by the
    // orchestrator — this app never writes another app's stock.
    Restock bool `json:"restock"`

    // Used by Decode() method
    data []byte
}

func (model OrderReturnedPosition) New(data []byte) *OrderReturnedPosition {
    model.data = data
    return &model
}

func (model *OrderReturnedPosition) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}