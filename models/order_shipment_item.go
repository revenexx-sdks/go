package models

import (
    "encoding/json"
    "errors"
)

// OrderShipmentItem One line of a delivery note: how much of one order
// position went out in one shipment.
type OrderShipmentItem struct {
    // When the booking was written.
    CreatedAt string `json:"created_at"`
    // Primary key of the booked position line.
    Id string `json:"id"`
    // Which order position went out. Always a position of the same order as the
    // shipment.
    OrderItemId string `json:"order_item_id"`
    // How much of that position this shipment carried. The sum of these over all
    // shipments is the position's quantity_shipped.
    Quantity float64 `json:"quantity"`
    // The shipment this booking belongs to. Deleting the shipment deletes it.
    ShipmentId string `json:"shipment_id"`

    // Used by Decode() method
    data []byte
}

func (model OrderShipmentItem) New(data []byte) *OrderShipmentItem {
    model.data = data
    return &model
}

func (model *OrderShipmentItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}