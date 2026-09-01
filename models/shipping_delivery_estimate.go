package models

import (
    "encoding/json"
    "errors"
)

// ShippingDeliveryEstimate The delivery window a checkout can print. Calendar
// days, cut-off evaluated in UTC (send `at` to control the instant).
type ShippingDeliveryEstimate struct {
    // Whether the cut-off had passed at evaluation time, costing a day.
    CutoffPassed bool `json:"cutoff_passed"`
    // The cut-off applied (HH:MM, UTC), or null when none is configured — the
    // carrier's own when it declares one, else the market's `cutoff_time`
    // setting.
    CutoffTime string `json:"cutoff_time"`
    // ship_date + eta_days_min.
    Earliest string `json:"earliest"`
    // The tenant's handling_days setting, as applied.
    HandlingDays int `json:"handling_days"`
    // ship_date + eta_days_max.
    Latest string `json:"latest"`
    // The day the parcel leaves — today plus handling days, plus one when the
    // cut-off has passed.
    ShipDate string `json:"ship_date"`

    // Used by Decode() method
    data []byte
}

func (model ShippingDeliveryEstimate) New(data []byte) *ShippingDeliveryEstimate {
    model.data = data
    return &model
}

func (model *ShippingDeliveryEstimate) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}