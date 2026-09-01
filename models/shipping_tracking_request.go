package models

import (
    "encoding/json"
    "errors"
)

// ShippingTrackingRequest One parcel, resolved into a tracking link by the
// carrier that owns the URL format.
type ShippingTrackingRequest struct {
    // Carrier code (what an order shipment already stores) or the carrier row id
    // — a value matching the uuid form is read as the id, anything else as a
    // code, case-insensitively. Must name a carrier THIS tenant keeps; one that
    // does not is a 404.
    Carrier string `json:"carrier"`
    // Destination ISO 3166-1 alpha-2 code — only needed by a template that
    // names {country}. Upper-cased before substitution.
    Country string `json:"country"`
    // Destination postcode — only needed by a template that names
    // {postal_code}.
    PostalCode string `json:"postal_code"`
    // The carrier's tracking number. Required by every template that names
    // {tracking_code}, which is all of them in the shipped catalog. URL-encoded
    // before substitution, so a code with a space or a slash cannot reshape the
    // link.
    TrackingCode string `json:"tracking_code"`

    // Used by Decode() method
    data []byte
}

func (model ShippingTrackingRequest) New(data []byte) *ShippingTrackingRequest {
    model.data = data
    return &model
}

func (model *ShippingTrackingRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}