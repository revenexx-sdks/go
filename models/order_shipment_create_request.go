package models

import (
    "encoding/json"
    "errors"
)

// OrderShipmentCreateRequest Book what went out. Every field is optional: an
// empty body ships every position that still has an open quantity, in full,
// on a delivery note number drawn from the tenant's delivery range — which
// is the whole payload for the common case.
type OrderShipmentCreateRequest struct {
    // Who is carrying it, in the merchant's own words. Free text — this app
    // neither validates it nor knows the carrier's API.
    Carrier string `json:"carrier"`
    // Free-form data for the caller — the warehouse system's own reference for
    // this handover. Stored and returned untouched.
    Metadata interface{} `json:"metadata"`
    // The DELIVERY NOTE number — drawn from the tenant's delivery range, unique
    // per tenant, and a different series from the order number. A caller may
    // supply its own when the number is issued by the warehouse system instead.
    // Drawn from the 'delivery' range when omitted; supply one only when the
    // number is issued elsewhere.
    Number string `json:"number"`
    // What this shipment carries. Omitted = every position with an open quantity,
    // in full. GET /orders/{id}/shippable answers exactly the budget each one is
    // guarded against.
    Positions []OrderShipmentPosition `json:"positions"`
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

func (model OrderShipmentCreateRequest) New(data []byte) *OrderShipmentCreateRequest {
    model.data = data
    return &model
}

func (model *OrderShipmentCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}