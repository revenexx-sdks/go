package models

import (
    "encoding/json"
    "errors"
)

// ShippingTrackingCarrier The carrier row that owns the URL format,
// identified so the caller can show who is carrying the parcel without a
// second read. Resolved whatever its status — a retired carrier still
// answers here.
type ShippingTrackingCarrier struct {
    // Stable carrier code, unique per tenant (e.g. dhl, dpd, gls). A method whose
    // `carrier` text equals this code resolves to this carrier — that is the
    // migration path off the free-text field. Deliberately no slug pattern: the
    // column asks only for a non-empty string, and a contract stricter than the
    // implementation would refuse codes merchants already keep.
    Code string `json:"code"`
    // Row id, assigned by the database on insert.
    Id string `json:"id"`
    // Display name, for the line that reads "shipped with …".
    Name string `json:"name"`
    // The class of service this row represents (default 'standard'), as a CODE
    // into the tenant's own service levels (GET /shipping/service-levels). One
    // row is one class: a carrier selling both a parcel and an express product is
    // two rows. Deliberately not an enum here — the set is the merchant's, so a
    // fixed list in this contract would make the gateway reject a level they
    // created. A code the tenant does not keep is a 400 naming the codes they do.
    ServiceLevel string `json:"service_level"`
    // Whether this carrier may be quoted (default 'active'). Anything else
    // excludes every method that ships with it from POST /shipping/rates, with a
    // reason. Tracking links are NOT gated on it — a retired carrier's old
    // shipments stay resolvable. Reported here so a UI can mark a link as
    // belonging to a carrier nobody quotes any more.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model ShippingTrackingCarrier) New(data []byte) *ShippingTrackingCarrier {
    model.data = data
    return &model
}

func (model *ShippingTrackingCarrier) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}