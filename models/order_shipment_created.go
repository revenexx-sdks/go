package models

import (
    "encoding/json"
    "errors"
)

// OrderShipmentCreated What the booking produced: the new shipment with the
// quantities it took, and the order as it now stands.
type OrderShipmentCreated struct {
    // The order after the booking: fulfillment_status is re-derived from the
    // positions, and status may have moved to in_fulfillment or (depending on the
    // tenant's auto_complete_on) completed.
    Order Order `json:"order"`
    // The shipment that was created, WITH the position quantities it booked —
    // the only place a caller learns which quantities actually went out when the
    // positions were defaulted.
    Shipment OrderShipment `json:"shipment"`

    // Used by Decode() method
    data []byte
}

func (model OrderShipmentCreated) New(data []byte) *OrderShipmentCreated {
    model.data = data
    return &model
}

func (model *OrderShipmentCreated) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}