package models

import (
    "encoding/json"
    "errors"
)

// OrderShipment One handover to a carrier — a delivery note. An order has
// as many of these as it took to get the goods out; each carries the position
// quantities it booked.
type OrderShipment struct {
    // Who is carrying it, in the merchant's own words. Free text — this app
    // neither validates it nor knows the carrier's API.
    Carrier string `json:"carrier"`
    // When the shipment was booked here, which is not necessarily when it left
    // — that is shipped_at.
    CreatedAt string `json:"created_at"`
    // Primary key of the shipment.
    Id string `json:"id"`
    // The booked position quantities of this shipment.
    Items []OrderShipmentItem `json:"items"`
    // Free-form data for the caller — the warehouse system's own reference for
    // this handover. Stored and returned untouched.
    Metadata interface{} `json:"metadata"`
    // The DELIVERY NOTE number — drawn from the tenant's delivery range, unique
    // per tenant, and a different series from the order number. A caller may
    // supply its own when the number is issued by the warehouse system instead.
    Number string `json:"number"`
    // The order this shipment belongs to. Deleting the order deletes its
    // shipments.
    OrderId string `json:"order_id"`
    // When the goods actually left. Defaults to now, and a caller may backdate it
    // — a shipment booked on Monday for a Friday handover says Friday.
    ShippedAt string `json:"shipped_at"`
    // The consignment number the carrier issued. Free text: every carrier formats
    // it differently and this app stores whatever it is given.
    TrackingCode string `json:"tracking_code"`
    // Where a human can follow the parcel. Supplied by the caller — this app
    // does not build it, because only the caller knows the carrier's tracking
    // address.
    TrackingUrl string `json:"tracking_url"`

    // Used by Decode() method
    data []byte
}

func (model OrderShipment) New(data []byte) *OrderShipment {
    model.data = data
    return &model
}

func (model *OrderShipment) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}