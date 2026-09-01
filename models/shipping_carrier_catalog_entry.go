package models

import (
    "encoding/json"
    "errors"
)

// ShippingCarrierCatalogEntry One carrier this app knows the facts for,
// exactly as it would be created.
type ShippingCarrierCatalogEntry struct {
    // The code the seeded row would carry, and the code a method's `carrier` text
    // has to match to resolve to it.
    Code string `json:"code"`
    // The countries this carrier serves. ISO 3166-1 alpha-2 codes; null or an
    // empty array means no restriction. Compared upper-cased, so a lower-case
    // entry still matches. Declared as an array rather than the bare object a
    // jsonb column derives to — this one is always a list.
    Countries []string `json:"countries"`
    // This carrier's own daily pickup cut-off, HH:MM in 24-hour form, UTC.
    // Overrides the tenant's cutoff_time for methods on this carrier — one
    // shop-wide time cannot be both DHL's 16:00 and a forwarder's 12:00. Null or
    // the empty string means this carrier declares none; any other shape is a
    // 400, because a cut-off the estimator cannot read is a delivery promise
    // silently computed without one.
    CutoffTime string `json:"cutoff_time"`
    // Transit time upper bound, in calendar days from the ship date.
    EtaDaysMax int `json:"eta_days_max"`
    // Transit time lower bound, in calendar days from the ship date — inherited
    // by any method on this carrier that states no ETA of its own.
    EtaDaysMin int `json:"eta_days_min"`
    // Days needed to make a consignment ready for THIS carrier, added to the ship
    // date before the transit days. Overrides the tenant's handling_days.
    HandlingDays int `json:"handling_days"`
    // Localized display names the seed would carry. A flat map keyed by locale
    // — the Cockpit falls back to `en`. Null means the row has no translations
    // and every client shows the untranslated column instead.
    Labels interface{} `json:"labels"`
    // The display name the seeded row would carry. An existing row keeps the
    // merchant's own name — the seed never writes over one.
    Name string `json:"name"`
    // Whether a fresh install starts with this carrier. False means this app
    // knows how to describe it but only creates it when asked.
    Seeded bool `json:"seeded"`
    // Service-level code the seeded row carries — one of the tenant's own
    // values.
    ServiceLevel string `json:"service_level"`
    // Tracking page URL with {tracking_code} where the number goes; {postal_code}
    // and {country} are also substituted, URL-encoded. Null for a carrier with no
    // public tracking page.
    TrackingUrlTemplate string `json:"tracking_url_template"`

    // Used by Decode() method
    data []byte
}

func (model ShippingCarrierCatalogEntry) New(data []byte) *ShippingCarrierCatalogEntry {
    model.data = data
    return &model
}

func (model *ShippingCarrierCatalogEntry) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}